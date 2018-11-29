package serverapi

import (
	"github.com/airbloc/airbloc-go/data/datamanager"
	"github.com/airbloc/airbloc-go/node"
	commonpb "github.com/airbloc/airbloc-go/proto/rpc/v1"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"golang.org/x/net/context"
)

type DataAPI struct {
	manager *datamanager.Manager
}

func NewDataAPI(backend node.Backend) (node.API, error) {
	manager := datamanager.NewManager(backend.Kms(), backend.LocalDatabase(), backend.Client())
	return &DataAPI{manager}, nil
}

func (api *DataAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterDataServer(service.GrpcServer, api)
}

func (api *DataAPI) Get(ctx context.Context, dataId *pb.DataId) (*pb.DataResult, error) {
	return nil, nil
}

func (api *DataAPI) BatchGet(ctx context.Context, batchId *pb.BatchRequest) (*pb.BatchGetResult, error) {
	return nil, nil
}

func (api *DataAPI) SetPermission(ctx context.Context, req *pb.SetDataPermissionRequest) (*commonpb.Result, error) {
	return nil, nil
}

func (api *DataAPI) SetPermissionBatch(ctx context.Context, req *pb.SetBatchDataPermissionRequest) (*commonpb.Results, error) {
	return nil, nil
}

func (api *DataAPI) Delete(ctx context.Context, dataId *pb.DataId) (*commonpb.Result, error) {
	return nil, nil
}

func (api *DataAPI) DeleteBatch(ctx context.Context, batchId *pb.BatchRequest) (*commonpb.Results, error) {
	return nil, nil
}

func (api *DataAPI) Select(stream pb.Data_SelectServer) error {
	return nil
}

func (api *DataAPI) Release(ctx context.Context, batchId *pb.BatchRequest) (*commonpb.Result, error) {
	return nil, nil
}
