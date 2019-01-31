package serverapi

import (
	"github.com/airbloc/airbloc-go/common"
	dataType "github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/data/datamanager"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type DataAPI struct {
	manager *datamanager.Manager
}

func NewDataAPI(backend node.Backend) (node.API, error) {
	manager := datamanager.NewManager(
		backend.Kms(),
		backend.P2P(),
		backend.LocalDatabase(),
		backend.Client(),
		backend.GetService("warehouse").(*warehouse.Service).GetManager())
	return &DataAPI{manager}, nil
}

func (api *DataAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterDataServer(service.GrpcServer, api)
}

func (api *DataAPI) makeDataResult(bundle *dataType.Bundle, data *common.Data) *pb.DataResult {
	return &pb.DataResult{
		CollectionId: bundle.Collection.Hex(),
		OwnerUserAid: data.UserId.Hex(),
		IngestedAt:   bundle.IngestedAt.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)),
		Payload:      data.Payload,
	}
}

func (api *DataAPI) Get(ctx context.Context, dataId *pb.DataId) (*pb.DataResult, error) {
	bundle, data, err := api.manager.Get(dataId.DataId)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to get data %s", dataId.DataId)
	}
	return api.makeDataResult(bundle, data), nil
}

func (api *DataAPI) GetBatch(ctx context.Context, batchId *pb.BatchRequest) (*pb.GetBatchResult, error) {
	batchManager := api.manager.Batches()
	batchInfo, err := batchManager.Get(batchId.BatchId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get batchId from manager %s", batchId.BatchId)
	}

	batch, err := api.manager.GetBatch(batchInfo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get batch data from given batchInfo %s", batchId.BatchId)
	}

	batchResult := make([]*pb.DataResult, len(batch))
	index := 1
	for bundleInfo, bundle := range batch {
		for _, data := range bundle {
			batchResult[index] = api.makeDataResult(bundleInfo, data)
			index++
		}
	}
	return &pb.GetBatchResult{Data: batchResult}, nil
}

func (api *DataAPI) SetPermission(ctx context.Context, req *pb.SetDataPermissionRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented method")
}

func (api *DataAPI) SetPermissionBatch(ctx context.Context, req *pb.SetBatchDataPermissionRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented method")
}

func (api *DataAPI) Delete(ctx context.Context, dataId *pb.DataId) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented method")
}

func (api *DataAPI) DeleteBatch(ctx context.Context, batchId *pb.BatchRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented method")
}

func (api *DataAPI) Select(stream pb.Data_SelectServer) error {
	return nil
}

func (api *DataAPI) Release(ctx context.Context, batchId *pb.BatchRequest) (*empty.Empty, error) {
	return nil, nil
}
