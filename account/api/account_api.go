package api

import (
	"context"
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/api"
	commonApi "github.com/airbloc/airbloc-go/common/api"
)

type API struct {
	manager *account.Manager
}

func New(backend *api.AirblocBackend) (api.API, error) {
	manager := account.NewManager(backend.Ethclient)
	return &API{manager}, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterAccountServer(service.GrpcServer, api)
}

func (api *API) Create(ctx context.Context, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	id, err := api.manager.Create(ctx)
	return &AccountCreateResponse{
		AccountId: &commonApi.Hash{Hash: id[:]},
	}, err
}
