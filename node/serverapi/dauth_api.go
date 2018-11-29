package serverapi

import (
	"context"
	"github.com/airbloc/airbloc-go/account"
	"github.com/azer/logger"
	"log"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/node"
	commonpb "github.com/airbloc/airbloc-go/proto/rpc/v1"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/pkg/errors"
)

type DAuthAPI struct {
	accountManager *account.Manager
	dauthManager   *dauth.Manager
}

func NewDAuthAPI(backend node.Backend) (node.API, error) {
	return &DAuthAPI{
		accountManager: account.NewManager(backend.Client()),
		dauthManager:   dauth.NewManager(backend.Client()),
	}, nil
}

func (api *DAuthAPI) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	// check that user has been already signed up
	acc, err := api.accountManager.GetByIdentity(req.GetIdentity())
	if err != nil {
		if err == account.ErrNoAccount {
			api.accountManager.CreateTemporary()

		}
		return nil, errors.Wrap(err, "unable to call Accounts.GetByIdentity")
	}
	api.log.Info("SignIn(\"%s\"): Already signed up.", req.GetIdentity(), logger.Attrs{
		"id":     acc.ID.String(),
		"status": acc.Status,
		"proxy":  acc.Proxy.Hex(),
	})
	return &pb.SignInResponse{
		AccountId: acc.ID.String(),
	}, nil
}

func (api *DAuthAPI) Allow(ctx context.Context, req *pb.DAuthRequest) (*commonpb.Result, error) {
	collectionId, err := common.IDFromString(req.GetCollectionId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid collection ID: %s", req.GetCollectionId())
	}

	// err = api.dauthManager.Allow(collectionId, req.GetPasswordSignature())
	log.Println(collectionId.String())
	return &commonpb.Result{}, err
}

func (api *DAuthAPI) Deny(ctx context.Context, req *pb.DAuthRequest) (*commonpb.Result, error) {
	collectionId, err := common.IDFromString(req.GetCollectionId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid collection ID: %s", req.GetCollectionId())
	}

	log.Println(collectionId.String())
	// err = api.dauthManager.Deny(collectionId, req.GetPasswordSignature())
	return &commonpb.Result{}, err
}

func (api *DAuthAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterDAuthServer(service.GrpcServer, api)
}
