package warehouse

import (
	"context"
	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/pkg/errors"
)

type API struct {
	warehouse *DataWarehouse
}

func NewAPI(airbloc *api.AirblocBackend) (api.API, error) {
	config := airbloc.Config.Warehouse

	supportedProtocols := []protocol.Protocol{
		protocol.NewHttpProtocol(config.Http.Timeout, config.Http.MaxConnsPerHost),
		protocol.NewHttpsProtocol(config.Http.Timeout, config.Http.MaxConnsPerHost),
	}

	var defaultStorage storage.Storage
	if config.DefaultStorage == "local" {
		defaultStorage = storage.NewLocalStorage(
			config.LocalStorage.SavePath,
			config.LocalStorage.Endpoint)
	} else {
		return nil, errors.Errorf("unknown storage type: %s", config.DefaultStorage)
	}

	warehouse := New(airbloc.Kms, defaultStorage, supportedProtocols)
	return &API{warehouse}, nil
}

func (api *API) StoreBundle(stream Warehouse_StoreBundleServer) error {
	return nil
}

func (api *API) StoreEncryptedBundle(stream Warehouse_StoreEncryptedBundleServer) error {
	return nil
}

func (api *API) DeleteBundle(context context.Context, request *DeleteBundleRequest) (*DeleteBundleResult, error) {
	return nil, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterWarehouseServer(service.GrpcServer, api)
}
