package api

import (
	"context"
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/api"
	"github.com/ethereum/go-ethereum/common"
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
	address := common.BytesToAddress(req.GetAddress())
	id, err := api.manager.CreateUsingProxy(address, req.GetPasswordSignature())
	return &AccountCreateResponse{AccountId: id.String()}, err
}
