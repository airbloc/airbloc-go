package api

import (
	"github.com/airbloc/airbloc-go/api"
	commonApi "github.com/airbloc/airbloc-go/common/api"
	"github.com/airbloc/airbloc-go/exchange"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type API struct {
	exchange *exchange.Exchange
}

func New(backend api.Backend) (api.API, error) {
	ex, err := exchange.New(backend.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Exchange API")
	}
	return &API{ex}, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterExchangeServer(service.GrpcServer, api)
}

// TODO
func (api *API) Order(ctx context.Context, req *OrderRequest) (*commonApi.Hash, error) {
	return nil, nil
}

func (api *API) Settle(ctx context.Context, id *commonApi.Hash) (*commonApi.Result, error) {
	return nil, nil
}

func (api *API) Reject(ctx context.Context, id *OrderIdRequest) (*commonApi.Result, error) {
	return nil, nil
}

func (api *API) CloseOrder(ctx context.Context, id *OrderIdRequest) (*commonApi.Result, error) {
	return nil, nil
}
