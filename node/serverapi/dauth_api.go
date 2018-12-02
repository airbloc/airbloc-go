package serverapi

import (
	"context"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/node"
	commonpb "github.com/airbloc/airbloc-go/proto/rpc/v1"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DAuthAPI struct {
	dauthClient *dauth.ProviderClient
}

func NewDAuthAPI(backend node.Backend) (node.API, error) {
	dauthClient := dauth.NewProviderClient(backend.Kms(), backend.Client(), backend.P2P())
	return &DAuthAPI{dauthClient}, nil
}

func (api *DAuthAPI) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	accountId, err := api.dauthClient.SignIn(ctx, req.GetIdentity(), req.GetUserDelegate())
	if err != nil {
		return nil, err
	}
	return &pb.SignInResponse{
		AccountId: accountId.String(),
	}, nil
}

func (api *DAuthAPI) Allow(ctx context.Context, req *pb.DAuthRequest) (*commonpb.Result, error) {
	collectionId, err := common.IDFromString(req.GetCollectionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid collection ID: %s", req.GetCollectionId())
	}
	accountId, err := common.IDFromString(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid account ID: %s", req.GetAccountId())
	}
	err = api.dauthClient.Allow(ctx, collectionId, accountId)
	return &commonpb.Result{}, err
}

func (api *DAuthAPI) Deny(ctx context.Context, req *pb.DAuthRequest) (*commonpb.Result, error) {
	collectionId, err := common.IDFromString(req.GetCollectionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid collection ID: %s", req.GetCollectionId())
	}
	accountId, err := common.IDFromString(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid account ID: %s", req.GetAccountId())
	}
	err = api.dauthClient.Deny(ctx, collectionId, accountId)
	return &commonpb.Result{}, err
}

func (api *DAuthAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterDAuthServer(service.GrpcServer, api)
}
