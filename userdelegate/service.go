package userdelegate

import (
	"context"
	"encoding/base64"
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/apps"
	"github.com/airbloc/airbloc-go/collections"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/p2p"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/azer/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrDelegationNotAllowed = errors.New("the user haven't been designated you as a delegate.")
)

type Service struct {
	accountIds []ablCommon.ID
	p2p        p2p.Server
	isRunning  bool

	// node identity information
	id   string
	addr ethCommon.Address

	// managers for blockchain interaction
	apps        *apps.Manager
	dauth       *dauth.Manager
	accounts    *account.Manager
	collections *collections.Collections

	log *logger.Logger
}

func NewService(backend node.Backend) (node.Service, error) {
	var accountIds []ablCommon.ID
	for _, accIdStr := range backend.Config().UserDelegate.AccountIds {
		accountId, err := ablCommon.HexToID(accIdStr)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid account ID: %s", accIdStr)
		}
		accountIds = append(accountIds, accountId)
	}
	key := backend.Kms().NodeKey()
	nodeId := base64.StdEncoding.EncodeToString(crypto.FromECDSAPub(&key.PublicKey))
	return &Service{
		accountIds:  accountIds,
		p2p:         backend.P2P(),
		addr:        key.EthereumAddress,
		id:          nodeId,
		apps:        apps.NewManager(backend.Client()),
		dauth:       dauth.NewManager(backend.Client()),
		accounts:    account.NewManager(backend.Client()),
		collections: collections.New(backend.Client()),
		log:         logger.New("userdelegate"),
	}, nil
}

func (service *Service) Sync(ctx context.Context) error {
	accounts := service.accounts.GetContract()

	proxyAddress := []ethCommon.Address{service.addr}
	options := &bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	}
	events, err := accounts.FilterTemporaryCreated(options, proxyAddress, [][32]byte{})
	if err != nil {
		return errors.Wrap(err, "failed to scan events in Accounts")
	}
	for events.Next() {
		accountId := ablCommon.ID(events.Event.AccountId)
		service.AddUser(accountId)
	}
	if events.Error() != nil {
		return errors.Wrap(events.Error(), "failed to iterate over events in Accounts")
	}
	return nil
}

// AddUser adds a user to the delegated user list,
// therefore manage
func (service *Service) AddUser(accountId ablCommon.ID) error {
	// you can be delegate of a user after the user designate you as a delegate.
	if isDelegate, err := service.accounts.IsDelegateOf(service.addr, accountId); err != nil {
		return errors.Wrapf(err, "failed to call Accounts.IsDelegateOf")
	} else if !isDelegate {
		return ErrDelegationNotAllowed
	}
	service.accountIds = append(service.accountIds, accountId)
	return nil
}

// HasUser returns true if given user account ID is registered.
func (service *Service) HasUser(accountId ablCommon.ID) bool {
	for _, id := range service.accountIds {
		if id == accountId {
			return true
		}
	}
	return false
}

func (service *Service) Start() error {
	service.log.Info("Starting service...")

	ctx, cancelSync := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelSync()

	timer := service.log.Timer()
	if err := service.Sync(ctx); err != nil {
		return err
	}
	timer.End("%d accounts have been scanned", len(service.accountIds))

	// register p2p RPC handlers
	rpc := p2p.NewRPC(service.p2p)
	rpc.Handle("dauth-allow", &pb.DAuthRequest{}, &pb.DAuthResponse{}, service.createDAuthHandler(true))
	rpc.Handle("dauth-deny", &pb.DAuthRequest{}, &pb.DAuthResponse{}, service.createDAuthHandler(false))
	rpc.Handle("dauth-signup", &pb.DAuthSignUpRequest{}, &pb.DAuthSignUpResponse{}, service.signUpHandler)

	service.isRunning = true

	service.log.Info("User Delegate ID=%s", service.id)
	return nil
}

func (service *Service) createDAuthHandler(allow bool) p2p.RPCHandler {
	return func(ctx context.Context, from p2p.SenderInfo, req proto.Message) (proto.Message, error) {
		request, _ := req.(*pb.DAuthRequest)

		accountId, err := ablCommon.HexToID(request.GetAccountId())
		if err != nil {
			return nil, errors.Wrapf(err, "invalid account ID %s", request.GetAccountId())
		}
		collectionId, err := ablCommon.HexToID(request.CollectionId)
		if err != nil {
			return nil, errors.Wrapf(err, "Invalid collection ID %s", request.GetCollectionId())
		}

		// check that the given user is registered
		if !service.HasUser(accountId) {
			return nil, errors.Errorf("user %s is not registered", accountId.Hex())
		}

		// the message sender should be the data provider (the collection's owner)
		if ok, err := service.isCollectionOwner(ctx, collectionId, from.Addr); err != nil {
			return nil, errors.Wrap(err, "failed to retrieve collection owner")
		} else if !ok {
			return nil, errors.Wrapf(err, "The address %s is not a data provider.", collectionId.Hex())
		}

		if allow {
			err = service.dauth.AllowByDelegate(collectionId, accountId)
		} else {
			err = service.dauth.DenyByDelegate(collectionId, accountId)
		}
		if err != nil {
			return nil, errors.Wrap(err, "failed to modify DAuth settings")
		}
		return &pb.DAuthResponse{}, nil
	}
}

func (service *Service) signUpHandler(ctx context.Context, from p2p.SenderInfo, req proto.Message) (proto.Message, error) {
	request, _ := req.(*pb.DAuthSignUpRequest)

	identityHash := ethCommon.HexToHash(request.GetIdentityHash())
	accountId, err := service.accounts.CreateTemporary(identityHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create temporary account")
	}

	service.log.Info("Created account %s by request from the data provider %s", accountId.Hex(), from.Addr.Hex())
	service.AddUser(accountId)

	return &pb.DAuthSignUpResponse{
		AccountId: accountId.Hex(),
	}, nil
}

// isCollectionOwner checks that the P2P message sender is
// same with the owner of the collection (data provider, app owner).
func (service *Service) isCollectionOwner(ctx context.Context, collectionId ablCommon.ID, senderAddr ethCommon.Address) (bool, error) {
	collection, err := service.collections.Get(collectionId)
	if err != nil {
		return false, errors.Wrap(err, "unable to retrieve collection")
	}
	return service.apps.IsOwner(ctx, collection.AppId, senderAddr)
}

func (service *Service) Stop() {
	service.log.Info("Stopping...")
	service.isRunning = false
}
