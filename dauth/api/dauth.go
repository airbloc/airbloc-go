package api

import (
	"context"

	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/common"
	commonAPI "github.com/airbloc/airbloc-go/common/api"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/pkg/errors"
)

type API struct {
	manager *dauth.Manager
}

func New(backend *api.AirblocBackend) (api.API, error) {
	manager := dauth.NewManager(backend.Ethclient)
	return &API{manager}, nil
}

func (api *API) Allow(ctx context.Context, req *DAuthRequest) (*commonAPI.Result, error) {
	collectionId, err := common.IDFromString(req.GetCollectionId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid collection ID: %s", req.GetCollectionId())
	}

	err = api.manager.Allow(collectionId, req.GetPasswordSignature())
	return &commonAPI.Result{}, err
}

func (api *API) Deny(ctx context.Context, req *DAuthRequest) (*commonAPI.Result, error) {
	collectionId, err := common.IDFromString(req.GetCollectionId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid collection ID: %s", req.GetCollectionId())
	}

	err = api.manager.Deny(collectionId, req.GetPasswordSignature())
	return &commonAPI.Result{}, err
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterDAuthServer(service.GrpcServer, api)
}
