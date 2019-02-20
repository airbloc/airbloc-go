package serverapi

import (
	"context"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/user"
	"github.com/airbloc/airbloc-go/warehouse/service"
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

func (api *UserAPI) GetData(ctx context.Context, req *pb.UserId) (*pb.GetDataReponse, error) {
	userId, err := common.HexToID(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert userId to common.Id format : *v", err)
	}

	userData, err := api.manager.GetData(ctx, userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info : %v", err)
	}

	_ = userData

	return nil, nil
}

func (api *UserAPI) GetDataIds(ctx context.Context, req *pb.UserId) (*pb.GetDataIdsResponse, error) {
	userId, err := common.HexToID(req.GetUserId())
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
				Id:         dataId.Id,
				IngestedAt: dataId.IngestedAt,
			}
		}
	}
	return &pb.GetDataIdsResponse{Collections: collections}, nil
}
