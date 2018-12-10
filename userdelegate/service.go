package userdelegate

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/apps"
	"github.com/airbloc/airbloc-go/collections"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/p2p"
	p2pcommon "github.com/airbloc/airbloc-go/p2p/common"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/azer/logger"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
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
		accountId, err := ablCommon.IDFromString(accIdStr)
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
	if service.isRunning {
		service.registerDAuthHandler(accountId)
	}
	return nil
}

func (service *Service) Start() error {
	service.p2p.SubscribeTopic("dauth-signup", &pb.DAuthSignUpRequest{}, service.signUpHandler)

	for _, accountId := range service.accountIds {
		service.registerDAuthHandler(accountId)
	}
	service.log.Info("Starting service...")
	service.isRunning = true

	service.log.Info("User Delegate ID=%s", service.id)
	return nil
}

func (service *Service) registerDAuthHandler(accountId ablCommon.ID) {
	accId := accountId.String()
	dauthReq := &pb.DAuthRequest{}

	service.p2p.SubscribeTopic("dauth-allow-"+accId, dauthReq, service.createDAuthHandler(accountId, true))
	service.p2p.SubscribeTopic("dauth-deny-"+accId, dauthReq, service.createDAuthHandler(accountId, false))
}

func (service *Service) createDAuthHandler(accountId ablCommon.ID, allow bool) p2p.TopicHandler {
	return func(server p2p.Server, ctx context.Context, message p2pcommon.Message) {
		request, ok := message.Data.(*pb.DAuthRequest)
		if !ok {
			service.log.Error("Topic is mismatched with data format.")
			return
		}

		collectionId, err := ablCommon.IDFromString(request.CollectionId)
		if err != nil {
			service.log.Error("Invalid Collection ID %s: %s", collectionId, err.Error())
			return
		}

		// the message sender should be the data provider (the collection's owner)
		if ok, err := service.isCollectionOwner(ctx, collectionId, message.SenderAddr); err != nil {
			service.log.Error("Failed to retrieve collection owner: %s", err.Error())
			return
		} else if !ok {
			service.log.Error("The address %s is not a data provider.", message.SenderAddr.Hex())
			return
		}

		if allow {
			err = service.dauth.AllowByDelegate(collectionId, accountId)
		} else {
			err = service.dauth.DenyByDelegate(collectionId, accountId)
		}
		if err != nil {
			service.log.Error("Failed to modify DAuth settings: %s", err.Error())
			return
		}

		// respond to the data provider
		response := &pb.DAuthResponse{
			CollectionId: request.GetCollectionId(),
		}
		var topicName string
		if allow {
			topicName = fmt.Sprintf("dauth-allow-%s-response", accountId.String())
		} else {
			topicName = fmt.Sprintf("dauth-deny-%s-response", accountId.String())
		}
		if err := server.Send(ctx, response, topicName, message.SenderInfo.ID); err != nil {
			service.log.Error("Failed to send response to data provider: %s", err.Error())
		}
	}
}

func (service *Service) signUpHandler(server p2p.Server, ctx context.Context, message p2pcommon.Message) {
	request, ok := message.Data.(*pb.DAuthSignUpRequest)
	if !ok {
		service.log.Error("Topic is mismatched with data format.")
		return
	}

	identityHash := ethCommon.HexToHash(request.GetIdentityHash())
	accountId, err := service.accounts.CreateTemporary(identityHash)
	if err != nil {
		service.log.Error("Failed to create temporary account: %s", err.Error())
		return
	}

	service.log.Info("Created account %s by request from the data provider %s",
		accountId.String(), message.SenderAddr.Hex())
	service.AddUser(accountId)

	response := &pb.DAuthSignUpResponse{
		UserId: accountId.String(),
	}
	if err = server.Send(context.Background(), response, "dauth-signup-response", message.SenderInfo.ID); err != nil {
		service.log.Error("Failed to send response to data provider: %s", err.Error())
	}
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
