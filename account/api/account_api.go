package api

import (
	"context"

	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/api"
	commonApi "github.com/airbloc/airbloc-go/common/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type API struct {
	manager *account.Manager
}

func New(backend *api.AirblocBackend) (api.API, error) {
	manager, err := account.NewManager(backend.Ethclient, common.Address{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create AccountManager")
	}
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
