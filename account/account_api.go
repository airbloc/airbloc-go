package account

import (
	"context"
	"github.com/airbloc/airbloc-go/api"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type API struct {
	manager *Manager
}

func NewAPI(backend *api.AirblocBackend) (api.API, error) {
	manager, err := NewManager(backend.Ethclient, ethCommon.Address{})
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
		AccountId: &ablCommon.Hash{Hash: id[:]},
	}, err
}
