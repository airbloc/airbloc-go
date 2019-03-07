package api

import (
	"context"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/shared/node"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/airbloc-go/shared/user"
	"github.com/airbloc/airbloc-go/shared/warehouse/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserAPI struct {
	manager *user.Manager
}

func NewUserAPI(backend node.Backend) (node.API, error) {
	manager := user.NewManager(
		backend.Kms(),
		backend.MetaDatabase(),
		backend.Client(),
		backend.GetService("warehouse").(*warehouseservice.Service).GetManager(),
	)
	return &UserAPI{manager}, nil
}

func (api *UserAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterUserServer(service.GrpcServer, api)
}

func (api *UserAPI) GetData(ctx context.Context, req *pb.DataRequest) (*pb.GetDataReponse, error) {
	userId, err := types.HexToID(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert userId to common.Id format : *v", err)
	}

	usersData, err := api.manager.GetData(ctx, userId, req.GetFrom())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info : %v", err)
	}

	collections := make([]*pb.GetDataReponse_Collection, len(usersData))
	for i, collection := range usersData {
		collections[i] = new(pb.GetDataReponse_Collection)
		collections[i].Id = collection.CollectionId
		collections[i].Data = make([]*pb.GetDataReponse_Data, len(collection.Data))

		for j, userData := range collection.Data {
			collections[i].Data[j] = new(pb.GetDataReponse_Data)
			collections[i].Data[j] = &pb.GetDataReponse_Data{
				CollectedAt: userData.CollectedAt,
				IngestedAt:  userData.IngestedAt,
				Payload:     userData.Payload,
			}
		}
	}

	return &pb.GetDataReponse{Collections: collections}, nil
}

func (api *UserAPI) GetDataIds(ctx context.Context, req *pb.DataIdRequest) (*pb.GetDataIdsResponse, error) {
	userId, err := types.HexToID(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert userId to common.ID format : %v", err)
	}

	userInfoes, err := api.manager.GetDataIds(ctx, userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info : %v", err)
	}

	collections := make([]*pb.GetDataIdsResponse_Collection, len(userInfoes))
	for i, collection := range userInfoes {
		collections[i] = &pb.GetDataIdsResponse_Collection{
			CollectionId: collection.CollectionId,
			DataInfoes:   make([]*pb.GetDataIdsResponse_DataInfo, len(collection.DataIds)),
		}

		for j, dataId := range collection.DataIds {
			collections[i].DataInfoes[j] = &pb.GetDataIdsResponse_DataInfo{
				Id:          dataId.String(),
				CollectedAt: dataId.CollectedAt,
				IngestedAt:  collection.IngestedAt,
			}
		}
	}
	return &pb.GetDataIdsResponse{Collections: collections}, nil
}
