package serverapi

import (
	"context"

	"github.com/airbloc/airbloc-go/node"

	"github.com/airbloc/airbloc-go/apps"
	ablCommon "github.com/airbloc/airbloc-go/common"
	commonpb "github.com/airbloc/airbloc-go/proto/rpc/v1"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type AppsAPI struct {
	apps *apps.Manager
}

func NewAppsAPI(backend node.Backend) (node.API, error) {
	appsManager := apps.NewManager(backend.Client())
	return &AppsAPI{appsManager}, nil
}

func (api *AppsAPI) NewOwner(ctx context.Context, req *pb.NewOwnerRequest) (*commonpb.Result, error) {
	appId, err := ablCommon.IDFromString(req.GetAppId())
	if err != nil {
		return nil, errors.Wrap(err, "api : invalid app ID")
	}
	address := ethCommon.HexToAddress(req.GetOwner())

	_, err = api.apps.NewOwner(ctx, appId, address)
	if err != nil {
		return nil, errors.Wrap(err, "api : failed to set new owner")
	}
	return &commonpb.Result{}, nil
}

func (api *AppsAPI) CheckOwner(ctx context.Context, req *pb.CheckOwnerRequest) (*pb.CheckOwnerResult, error) {
	appId, err := ablCommon.IDFromString(req.GetAppId())
	if err != nil {
		return nil, errors.Wrap(err, "api : invalid app ID")
	}
	address := ethCommon.HexToAddress(req.GetOwner())

	ok, err := api.apps.CheckOwner(ctx, appId, address)
	if err != nil {
		return nil, errors.Wrap(err, "api : failed to check owner")
	}

	return &pb.CheckOwnerResult{Ok: ok}, nil
}

func (api *AppsAPI) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResult, error) {
	appId, err := api.apps.Register(ctx, req.GetName())
	if err != nil {
		return nil, errors.Wrap(err, "api : failed to register app")
	}
	return &pb.RegisterResult{
		AppId: appId.String(),
	}, nil
}
func (api *AppsAPI) Unregister(ctx context.Context, req *pb.UnregisterRequest) (*commonpb.Result, error) {
	appId, err := ablCommon.IDFromString(req.GetAppId())
	if err != nil {
		return nil, errors.Wrap(err, "api : invalid app ID")
	}

	_, err = api.apps.Unregister(ctx, appId)
	if err != nil {
		return nil, errors.Wrap(err, "api : failed to unregister app")
	}

	return &commonpb.Result{}, nil
}

func (api *AppsAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterAppsServer(service.GrpcServer, api)
}
