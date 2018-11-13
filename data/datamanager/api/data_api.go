package api

import (
	"github.com/airbloc/airbloc-go/api"
	commonApi "github.com/airbloc/airbloc-go/common/api"
	"github.com/airbloc/airbloc-go/data/datamanager"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type API struct {
	manager *datamanager.Manager
}

func New(backend *api.AirblocBackend) (api.API, error) {
	manager, err := datamanager.NewManager(backend.Kms, backend.Ethclient, backend.LocalDatabase, ethCommon.Address{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Data API")
	}
	return &API{manager}, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterDataServer(service.GrpcServer, api)
}

func (api *API) Get(ctx context.Context, dataId *DataId) (*DataResult, error) {
	return nil, nil
}

func (api *API) BatchGet(ctx context.Context, batchId *BatchRequest) (*BatchGetResult, error) {
	return nil, nil
}

func (api *API) SetPermission(ctx context.Context, req *SetDataPermissionRequest) (*commonApi.Result, error) {
	return nil, nil
}

func (api *API) SetPermissionBatch(ctx context.Context, req *SetBatchDataPermissionRequest) (*commonApi.Results, error) {
	return nil, nil
}

func (api *API) Delete(ctx context.Context, dataId *DataId) (*commonApi.Result, error) {
	return nil, nil
}

func (api *API) DeleteBatch(ctx context.Context, batchId *BatchRequest) (*commonApi.Results, error) {
	return nil, nil
}

func (api *API) Select(stream Data_SelectServer) error {
	return nil
}

func (api *API) Release(ctx context.Context, batchId *BatchRequest) (*commonApi.Result, error) {
	return nil, nil
}
