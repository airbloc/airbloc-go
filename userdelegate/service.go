package userdelegate

import (
	"context"
	"fmt"
	"log"

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
	"github.com/pkg/errors"
)

var (
	ErrDelegationNotAllowed = errors.New("the user haven't been designated you as a delegate.")
)

type Service struct {
	accountIds []ablCommon.ID
	p2p        p2p.Server
	selfAddr   ethCommon.Address
	isRunning  bool

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
	return &Service{
		accountIds:  accountIds,
		p2p:         backend.P2P(),
		selfAddr:    backend.Kms().NodeKey().EthereumAddress,
		apps:        apps.NewManager(backend.Client()),
		dauth:       dauth.NewManager(backend.Client()),
		accounts:    account.NewManager(backend.Client()),
		collections: collections.New(backend.LocalDatabase(), backend.MetaDatabase(), backend.Client()),
		log:         logger.New("userdelegate"),
	}, nil
}

// AddUser adds a user to the delegated user list,
// therefore manage
func (service *Service) AddUser(accountId ablCommon.ID) error {
	// you can be delegate of a user after the user designate you as a delegate.
	if isDelegate, err := service.accounts.IsDelegateOf(service.selfAddr, accountId); err != nil {
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
	select {}
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
			log.Println("error: Invalid topic.")
			return
		}

		collectionId, err := ablCommon.IDFromString(request.CollectionId)
		if err != nil {
			log.Println("error: Invalid Collection ID", collectionId, err.Error())
			return
		}

		// the message sender should be the data provider (the collection's owner)
		if ok, err := service.isCollectionOwner(ctx, collectionId, message.SenderAddr); err != nil {
			log.Println("error: Failed to retrieve collection owner", err.Error())
			return
		} else if !ok {
			log.Println("error: The address", message.SenderAddr.Hex(), "is not a data provider.")
			return
		}

		if allow {
			err = service.dauth.AllowByDelegate(collectionId, accountId)
		} else {
			err = service.dauth.DenyByDelegate(collectionId, accountId)
		}
		if err != nil {
			log.Println("error: Failed to modify DAuth settings: ", err.Error())
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
			log.Println("error: Failed to send response to data provider:", err.Error())
		}
	}
}

func (service *Service) signUpHandler(server p2p.Server, ctx context.Context, message p2pcommon.Message) {
	request, ok := message.Data.(*pb.DAuthSignUpRequest)
	if !ok {
		log.Println("error: Invalid topic.")
		return
	}

	identityHash := ethCommon.HexToHash(request.GetIdentityHash())
	accountId, err := service.accounts.CreateTemporary(identityHash)
	if err != nil {
		log.Println("error: Failed to create temporary account:", err.Error())
	}

	service.log.Info("Created account %s by request from the data provider %s",
		accountId.String(), message.SenderAddr.Hex())
	service.AddUser(accountId)

	response := &pb.DAuthSignUpResponse{
		UserId: accountId.String(),
	}
	if err = server.Send(context.Background(), response, "dauth-signup-response", message.SenderInfo.ID); err != nil {
		log.Println("error: Failed to send response to data provider:", err.Error())
	}
}

// isCollectionOwner checks that the P2P message sender is
// same with the owner of the collection (data provider, app owner).
func (service *Service) isCollectionOwner(ctx context.Context, collectionId ablCommon.ID, senderAddr ethCommon.Address) (bool, error) {
	collection, err := service.collections.Get(collectionId)
	if err != nil {
		return false, errors.Wrap(err, "unable to retrieve collection")
	}
	return service.apps.CheckOwner(ctx, collection.AppId, senderAddr)
}

func (service *Service) Stop() {
	service.log.Info("Stopping...")
	service.isRunning = false
}
