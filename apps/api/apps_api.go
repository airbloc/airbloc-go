package api

import (
	"context"

	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/apps"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type API struct {
	apps *apps.Manager
}

func New(backend api.Backend) (api.API, error) {
	appsManager, err := apps.NewManager(backend.Client())
	return &API{appsManager}, err
}

func (api *API) NewOwner(ctx context.Context, req *NewOwnerRequest) (*Response, error) {
	appId, err := ablCommon.IDFromString(req.GetAppId())
	if err != nil {
		return nil, errors.Wrap(err, "api : invalid app ID")
	}
	address := ethCommon.HexToAddress(req.GetOwner())

	res, err := api.apps.NewOwner(ctx, appId, address)
	if err != nil {
		return nil, errors.Wrap(err, "api : failed to set new owner")
	}

	return &Response{
		Success: res,
		Message: "",
	}, nil
}
func (api *API) CheckOwner(ctx context.Context, req *CheckOwnerRequest) (*Response, error) {
	appId, err := ablCommon.IDFromString(req.GetAppId())
	if err != nil {
		return nil, errors.Wrap(err, "api : invalid app ID")
	}
	address := ethCommon.HexToAddress(req.GetOwner())

	res, err := api.apps.CheckOwner(ctx, appId, address)
	if err != nil {
		return nil, errors.Wrap(err, "api : failed to check owner")
	}

	return &Response{
		Success: res,
		Message: "",
	}, nil
}
func (api *API) Register(ctx context.Context, req *RegisterRequest) (*Response, error) {
	appId, err := api.apps.Register(ctx, req.GetName())
	if err != nil {
		return nil, errors.Wrap(err, "api : failed to register app")
	}
	return &Response{
		Success: true,
		Message: appId.String(),
	}, nil
}
func (api *API) Unregister(ctx context.Context, req *UnregisterRequest) (*Response, error) {
	appId, err := ablCommon.IDFromString(req.GetAppId())
	if err != nil {
		return nil, errors.Wrap(err, "api : invalid app ID")
	}

	res, err := api.apps.Unregister(ctx, appId)
	if err != nil {
		return nil, errors.Wrap(err, "api : failed to unregister app")
	}

	return &Response{
		Success: res,
		Message: "",
	}, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterAppServer(service.GrpcServer, api)
}
