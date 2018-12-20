package serverapi

import (
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CollectionsAPI struct {
	collections *collections.Collections
}

func NewCollectionsAPI(backend node.Backend) (node.API, error) {
	collectionManager := collections.New(backend.Client())
	return &CollectionsAPI{collectionManager}, nil
}

func (api *CollectionsAPI) Create(ctx context.Context, req *pb.CreateCollectionRequest) (*pb.CreateCollectionResponse, error) {
	schemaId, err := common.HexToID(req.GetSchemaId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid schema ID: %s", req.GetSchemaId())
	}

	appId, err := common.HexToID(req.GetAppId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid app ID: %s", req.GetAppId())
	}

	collection := collections.NewCollection(
		appId,
		schemaId,
		collections.IncentivePolicy{
			DataProvider:  req.Policy.DataProvider,
			DataProcessor: req.Policy.DataProcessor,
			DataRelayer:   req.Policy.DataRelayer,
			DataOwner:     req.Policy.DataOwner,
		},
	)
	collectionId, err := api.collections.Register(ctx, collection)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCollectionResponse{
		CollectionId: collectionId.Hex(),
	}, nil
}

func (api *CollectionsAPI) List(ctx context.Context, req *pb.ListCollectionRequest) (*pb.ListCollectionResponse, error) {
	appId, err := common.HexToID(req.GetAppId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid app ID: %s", req.GetAppId())
	}

	cols, err := api.collections.List(ctx, appId)
	if err != nil {
		return nil, err
	}

	response := &pb.ListCollectionResponse{Total: int32(len(cols))}
	for _, col := range cols {
		collection := &pb.ListCollectionResponse_Collection{
			Id:       col.Id.Hex(),
			SchemaId: col.Schema.Id.Hex(),
		}
		response.Collections = append(response.Collections, collection)
	}
	return response, nil
}

func (api *CollectionsAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterCollectionServer(service.GrpcServer, api)
}
