package exchange

import (
	"github.com/airbloc/airbloc-go/api"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type API struct {
	exchange *Exchange
}

func NewAPI(backend *api.AirblocBackend) (api.API, error) {
	exchange, err := New(backend.Ethclient, ethCommon.Address{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Exchange API")
	}
	return &API{exchange}, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterExchangeServer(service.GrpcServer, api)
}

// TODO
func (api *API) Order(ctx context.Context, req *OrderRequest) (*ablCommon.Hash, error) {
	return nil, nil
}

func (api *API) Settle(ctx context.Context, req *SettleMessage) (*ablCommon.Result, error) {
	return nil, nil
}

func (api *API) Reject(ctx context.Context, id *OrderId) (*ablCommon.Result, error) {
	return nil, nil
}

func (api *API) CloseOrder(ctx context.Context, id *OrderId) (*ablCommon.Result, error) {
	return nil, nil
}
