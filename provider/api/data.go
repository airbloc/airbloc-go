package api

import (
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/shared/data/datamanager"
	"github.com/airbloc/airbloc-go/shared/node"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/airbloc-go/shared/warehouse/service"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type DataAPI struct {
	manager *datamanager.Manager
}

func NewDataAPI(backend node.Backend) (node.API, error) {
	manager := datamanager.NewManager(
		backend.Kms(),
		backend.P2P(),
		backend.MetaDatabase(),
		backend.LocalDatabase(),
		backend.Client(),
		backend.GetService("warehouse").(*warehouseservice.Service).GetManager())
	return &DataAPI{manager}, nil
}

func (api *DataAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterDataServer(service.GrpcServer, api)
}

func (api *DataAPI) Get(ctx context.Context, dataId *pb.DataId) (*pb.DataResult, error) {
	res, err := api.manager.Get(dataId.DataId)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to get data %s", dataId.DataId)
	}

	return &pb.DataResult{
		CollectionId: res.CollectionId.Hex(),
		UserId:       res.UserId.Hex(),
		IngestedAt:   res.IngestedAt.Timestamp(),
		Payload:      res.Payload,
	}, nil
}

func (api *DataAPI) GetBatch(ctx context.Context, batchId *pb.BatchRequest) (*pb.GetBatchResult, error) {
	batchManager := api.manager.Batches()
	batchInfo, err := batchManager.Get(batchId.BatchId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get batchId from manager %s", batchId.BatchId)
	}

	res, err := api.manager.GetBatch(batchInfo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get batch data from given batchInfo %s", batchId.BatchId)
	}

	batchResult := make([]*pb.DataResult, len(res))
	for i, data := range res {
		batchResult[i] = &pb.DataResult{
			CollectionId: data.CollectionId.Hex(),
			UserId:       data.UserId.Hex(),
			IngestedAt:   data.IngestedAt.Timestamp(),
			Payload:      data.Payload,
		}
	}
	return &pb.GetBatchResult{Data: batchResult}, nil
}

func (api *DataAPI) GetBundleInfo(ctx context.Context, request *pb.BundleInfoRequest) (*pb.BundleInfoResponse, error) {
	bundleId, err := types.HexToID(request.GetBundleId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert bundleId to common.ID format : %v", err)
	}

	bundleInfo, err := api.manager.GetBundleInfo(ctx, bundleId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get bundle info : %v", err)
	}

	return &pb.BundleInfoResponse{
		BundleId:   bundleInfo.Id,
		Uri:        bundleInfo.Uri,
		Provider:   bundleInfo.Provider,
		Collection: bundleInfo.Collection,
		IngestedAt: bundleInfo.IngestedAt,
		DataInfoes: bundleInfo.DataIds,
	}, nil
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
	return status.Error(codes.Unimplemented, "unimplemented method")
}

func (api *DataAPI) Release(ctx context.Context, batchId *pb.BatchRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented method")
}
