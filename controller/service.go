package controller

import (
	"context"
	"encoding/base64"
	"fmt"
	"reflect"
	"time"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/p2p"
	serviceLib "github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/golang/protobuf/proto"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/pkg/errors"
)

type MessageTypeError struct {
	msg string
	Typ reflect.Type
}

func (err *MessageTypeError) Error() string { return fmt.Sprintf("%s : %s", err.msg, err.Typ.String()) }

var (
	ErrDelegationNotAllowed = errors.New("the user haven't been designated you as a delegate.")
)

type service struct {
	accountIds []types.ID
	p2p        p2p.Server
	isRunning  bool

	// node identity information
	id   string
	addr common.Address

	// managers for blockchain interaction
	apps      adapter.IAppRegistryManager
	accounts  adapter.IAccountsManager
	consents  adapter.IConsentsManager
	dataTypes adapter.IDataTypeRegistryManager

	log *logger.Logger
}

func NewService(backend serviceLib.Backend) (serviceLib.Service, error) {
	var accountIds []types.ID
	for _, accIdStr := range backend.Config().UserDelegate.AccountIds {
		accountId, err := types.HexToID(accIdStr)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid account ID: %s", accIdStr)
		}
		accountIds = append(accountIds, accountId)
	}

	key := backend.Kms().NodeKey()
	nodeId := base64.StdEncoding.EncodeToString(crypto.FromECDSAPub(&key.PublicKey))
	return &service{
		accountIds: accountIds,
		p2p:        backend.P2P(),

		// node identity
		addr: key.EthereumAddress,
		id:   nodeId,

		// contract
		apps:      adapter.NewAppRegistryManager(backend.Client()),
		accounts:  adapter.NewAccountsManager(backend.Client()),
		consents:  adapter.NewConsentsManager(backend.Client()),
		dataTypes: adapter.NewDataTypeRegistryManager(backend.Client()),

		log: logger.New("controller"),
	}, nil
}

func (service *service) sync(ctx context.Context) (rerr error) {
	proxyAddress := []common.Address{service.addr}
	options := &bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	}
	events, err := service.accounts.FilterTemporaryCreated(options, proxyAddress, []common.Hash{})
	if err != nil {
		return errors.Wrap(err, "failed to scan events in Accounts")
	}
	defer func() {
		err = events.Close()
		if err != nil {
			rerr = err
		}
	}()

	for events.Next() {
		accountId := types.ID(events.Event.AccountId)
		_ = service.addUser(accountId)
	}
	if events.Error() != nil {
		return errors.Wrap(events.Error(), "failed to iterate over events in Accounts")
	}

	return nil
}

// AddUser adds a user to the delegated user list,
// therefore manage
func (service *service) addUser(accountId types.ID) error {
	// you can be delegate of a user after the user designate you as a delegate.
	if isController, err := service.accounts.IsControllerOf(service.addr, accountId); err != nil {
		return errors.Wrap(err, "failed to call accounts.isControllerOf")
	} else if !isController {
		return ErrDelegationNotAllowed
	}
	service.accountIds = append(service.accountIds, accountId)
	return nil
}

// HasUser returns true if given user account ID is registered.
func (service *service) hasUser(accountId types.ID) bool {
	for _, id := range service.accountIds {
		if id == accountId {
			return true
		}
	}
	return false
}

// Start controller service
func (service *service) Start() error {
	service.log.Info("Starting service...")

	ctx, cancelSync := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelSync()

	timer := service.log.Timer()
	if err := service.sync(ctx); err != nil {
		return err
	}
	timer.End("{} accounts have been scanned", len(service.accountIds))

	// register p2p RPC handlers
	rpc := p2p.NewRPC(service.p2p)
	rpc.Handle("dauth-allow", &pb.DAuthRequest{}, &pb.DAuthResponse{}, service.createConsentHandler(1))
	rpc.Handle("dauth-deny", &pb.DAuthRequest{}, &pb.DAuthResponse{}, service.createConsentHandler(0))
	rpc.Handle("dauth-many", &pb.DAuthManyRequest{}, &pb.DAuthResponse{}, service.consentManyHandler)
	rpc.Handle("dauth-signup", &pb.DAuthSignUpRequest{}, &pb.DAuthSignUpResponse{}, service.signUpHandler)

	service.isRunning = true

	service.log.Info("Started Data Controller", logger.Attrs{"id": service.id})
	return nil
}

func (service *service) createConsentHandler(allow uint8) p2p.RPCHandler {
	return func(ctx context.Context, from p2p.SenderInfo, req proto.Message) (proto.Message, error) {
		request, ok := req.(*pb.DAuthRequest)
		if !ok {
			return nil, &MessageTypeError{"invalid message type", reflect.TypeOf(req)}
		}

		var (
			appName  = request.GetAppName()
			dataType = request.GetDataType()
		)

		// accountId
		accountId, err := types.HexToID(request.GetAccountId())
		if err != nil {
			return nil, errors.Wrapf(err, "invalid account ID %s", request.GetAccountId())
		}

		// appName
		if appExists, err := service.apps.Exists(appName); err != nil {
			return nil, errors.Wrap(err, "failed to call appRegistry.Exists")
		} else if !appExists {
			return nil, errors.Errorf("app does not exist. appName: %s", appName)
		}

		// dataType
		if dataTypeExists, err := service.dataTypes.Exists(dataType); err != nil {
			return nil, errors.Wrap(err, "failed to call dataTypeRegistry.Exists")
		} else if !dataTypeExists {
			return nil, errors.Errorf("data type does not exist. dataType: %s", dataType)
		}

		// check that the given user is registered
		if !service.hasUser(accountId) {
			return nil, errors.Errorf("user %s is not registered", accountId.Hex())
		}

		if isController, err := service.accounts.IsControllerOf(service.addr, accountId); err != nil {
			return nil, errors.Wrap(err, "failed to call accounts.isControllerOf")
		} else if !isController {
			return nil, ErrDelegationNotAllowed
		}

		consentData := types.ConsentData{
			Action:   uint8(request.GetAction()),
			DataType: dataType,
			Allow:    allow,
		}
		err = service.consents.ConsentByController(ctx, accountId, appName, consentData)
		if err != nil {
			return nil, errors.Wrap(err, "failed to modify DAuth settings")
		}
		return &pb.DAuthResponse{}, nil
	}
}

func (service *service) consentManyHandler(
	ctx context.Context,
	from p2p.SenderInfo,
	req proto.Message,
) (proto.Message, error) {
	request, ok := req.(*pb.DAuthManyRequest)
	if !ok {
		return nil, &MessageTypeError{"invalid message type", reflect.TypeOf(req)}
	}

	appName := request.GetAppName()

	// accountId
	accountId, err := types.HexToID(request.GetAccountId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid account ID %s", request.GetAccountId())
	}

	// appName
	if appExists, err := service.apps.Exists(appName); err != nil {
		return nil, errors.Wrap(err, "failed to call appRegistry.Exists")
	} else if !appExists {
		return nil, errors.Errorf("app does not exist. appName: %s", appName)
	}

	consentData := make([]types.ConsentData, len(request.GetConsentData()))
	for index, consentRawData := range request.GetConsentData() {
		allow := uint8(0)
		if consentRawData.GetAllow() {
			allow = 1
		}
		consentData[index] = types.ConsentData{
			Action:   uint8(consentRawData.GetAction()),
			DataType: consentRawData.GetDataType(),
			Allow:    allow,
		}
	}

	err = service.consents.ConsentManyByController(ctx, accountId, appName, consentData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to modify DAuth settings")
	}
	return &pb.DAuthResponse{}, nil
}

func (service *service) signUpHandler(
	ctx context.Context,
	from p2p.SenderInfo,
	req proto.Message,
) (proto.Message, error) {
	request, ok := req.(*pb.DAuthSignUpRequest)
	if !ok {
		return nil, &MessageTypeError{"invalid message type", reflect.TypeOf(req)}
	}

	identityHash := common.HexToHash(request.GetIdentityHash())
	accountId, err := service.accounts.CreateTemporary(ctx, identityHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create temporary account")
	}

	service.log.Info("Created account {} by request from the data provider {}", accountId.Hex(), from.Addr.Hex())

	// retry 5 times
	for callCount := 0; callCount < 5; callCount++ {
		if err = service.addUser(accountId); err == nil || err == ErrDelegationNotAllowed {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	return &pb.DAuthSignUpResponse{
		AccountId: accountId.Hex(),
	}, nil
}

// Stop controller service
func (service *service) Stop() {
	service.log.Info("Stopping...")
	service.isRunning = false
}
