package data

import (
	"github.com/airbloc/airbloc-go/api"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type API struct {
	manager *Manager
}

func NewAPI(backend *api.AirblocBackend) (api.API, error) {
	manager, err := NewManager(backend.Kms, backend.Ethclient, backend.LocalDatabase, ethCommon.Address{})
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

func (api *API) SetPermission(ctx context.Context, req *SetDataPermissionRequest) (*ablCommon.Result, error) {
	return nil, nil
}

func (api *API) SetPermissionBatch(ctx context.Context, req *SetBatchDataPermissionRequest) (*ablCommon.Results, error) {
	return nil, nil
}

func (api *API) Delete(ctx context.Context, dataId *DataId) (*ablCommon.Result, error) {
	return nil, nil
}

func (api *API) DeleteBatch(ctx context.Context, batchId *BatchRequest) (*ablCommon.Results, error) {
	return nil, nil
}

func (api *API) Select(stream Data_SelectServer) error {
	return nil
}

func (api *API) Release(ctx context.Context, batchId *BatchRequest) (*ablCommon.Result, error) {
	return nil, nil
}
