package serverapi

import (
	"context"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
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
	userDelegateAddr := ethCommon.HexToAddress(req.GetUserDelegate())
	accountId, err := api.dauthClient.SignIn(ctx, req.GetIdentity(), userDelegateAddr)
	if err != nil {
		return nil, err
	}
	return &pb.SignInResponse{
		AccountId: accountId.Hex(),
	}, nil
}

func (api *DAuthAPI) Allow(ctx context.Context, req *pb.DAuthRequest) (*empty.Empty, error) {
	collectionId, err := common.HexToID(req.GetCollectionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid collection ID: %s", req.GetCollectionId())
	}
	accountId, err := common.HexToID(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid account ID: %s", req.GetAccountId())
	}
	err = api.dauthClient.Allow(ctx, collectionId, accountId)
	return &empty.Empty{}, err
}

func (api *DAuthAPI) Deny(ctx context.Context, req *pb.DAuthRequest) (*empty.Empty, error) {
	collectionId, err := common.HexToID(req.GetCollectionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid collection ID: %s", req.GetCollectionId())
	}
	accountId, err := common.HexToID(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid account ID: %s", req.GetAccountId())
	}
	err = api.dauthClient.Deny(ctx, collectionId, accountId)
	return &empty.Empty{}, err
}

func (api *DAuthAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterDAuthServer(service.GrpcServer, api)
	pb.RegisterDAuthHandlerFromEndpoint(context.Background(), service.RestAPIMux, service.Address, []grpc.DialOption{grpc.WithInsecure()})
}
