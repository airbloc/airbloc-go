package api

import (
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/shared/collections"
	"github.com/airbloc/airbloc-go/shared/schemas"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CollectionsAPI struct {
	collections *collections.Manager
	schemas     *schemas.Schemas
}

func NewCollectionsAPI(backend service.Backend) (api.API, error) {
	return &CollectionsAPI{
		collections: collections.NewManager(backend.Client()),
		schemas:     schemas.New(backend.MetaDatabase(), backend.Client()),
	}, nil
}

func (api *CollectionsAPI) Create(ctx context.Context, req *pb.CreateCollectionRequest) (*pb.CreateCollectionResponse, error) {
	schemaId, err := types.HexToID(req.GetSchemaId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid schema ID: %s", req.GetSchemaId())
	}

	appId, err := types.HexToID(req.GetAppId())
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

func (api *CollectionsAPI) Get(ctx context.Context, req *pb.GetCollectionRequest) (*pb.GetCollectionResult, error) {
	collectionId, err := types.HexToID(req.GetCollectionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid app ID: %s", req.GetCollectionId())
	}
	collection, err := api.collections.Get(collectionId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve collection %s", collectionId.Hex())
	}
	schema, err := api.schemas.Get(collection.Schema.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve schema %s", collection.Schema.Id.Hex())
	}

	return &pb.GetCollectionResult{
		Id:   req.GetCollectionId(),
		Name: req.GetCollectionId(), // TODO: return collection name
		Schema: &pb.GetCollectionResult_Schema{
			Id:     schema.Id.Hex(),
			Name:   schema.Name,
			Schema: schema.Schema,
		},
		Policy: &pb.Policy{
			DataOwner:    collection.IncentivePolicy.DataOwner,
			DataProvider: collection.IncentivePolicy.DataProvider,
		},
	}, nil
}

func (api *CollectionsAPI) List(ctx context.Context, req *pb.ListCollectionRequest) (*pb.ListCollectionResponse, error) {
	appId, err := types.HexToID(req.GetAppId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid app ID: %s", req.GetAppId())
	}

	collectionIds, err := api.collections.ListID(ctx, appId)
	if err != nil {
		return nil, err
	}

	response := &pb.ListCollectionResponse{Total: int32(len(collectionIds))}
	for _, id := range collectionIds {
		collection, err := api.Get(ctx, &pb.GetCollectionRequest{CollectionId: id.Hex()})
		if err != nil {
			return nil, err
		}
		response.Collections = append(response.Collections, collection)
	}
	return response, nil
}

func (api *CollectionsAPI) AttachToAPI(service *api.Service) {
	pb.RegisterCollectionServer(service.GrpcServer, api)
}
