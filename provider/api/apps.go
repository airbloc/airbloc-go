package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/shared/apps"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
)

type AppsAPI struct {
	apps *apps.Manager
}

func NewAppsAPI(backend service.Backend) (api.API, error) {
	appsManager := apps.NewManager(backend.Client())
	return &AppsAPI{appsManager}, nil
}

func (api *AppsAPI) NewOwner(ctx context.Context, req *pb.NewOwnerRequest) (*empty.Empty, error) {
	appId, err := types.HexToID(req.GetAppId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid app ID: %s", req.GetAppId())
	}
	address := ethCommon.HexToAddress(req.GetOwner())

	_, err = api.apps.TransferOwnership(ctx, appId, address)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (api *AppsAPI) CheckOwner(ctx context.Context, req *pb.CheckOwnerRequest) (*pb.CheckOwnerResult, error) {
	appId, err := types.HexToID(req.GetAppId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid app ID: %s", req.GetAppId())
	}
	address := ethCommon.HexToAddress(req.GetOwner())

	ok, err := api.apps.IsOwner(ctx, appId, address)
	if err != nil {
		return nil, err
	}
	return &pb.CheckOwnerResult{Ok: ok}, nil
}

func (api *AppsAPI) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResult, error) {
	appId, err := api.apps.Register(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResult{
		AppId: appId.Hex(),
	}, nil
}
func (api *AppsAPI) Unregister(ctx context.Context, req *pb.UnregisterRequest) (*empty.Empty, error) {
	appId, err := types.HexToID(req.GetAppId())
	if err != nil {
		return nil, errors.Wrap(err, "api : invalid app ID")
	}

	_, err = api.apps.Unregister(ctx, appId)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (api *AppsAPI) AttachToAPI(service *api.Service) {
	pb.RegisterAppsServer(service.GrpcServer, api)
}
