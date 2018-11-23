package serverapi

import (
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type CollectionsAPI struct {
	collections *collections.Collections
}

func NewCollectionsAPI(backend node.Backend) (node.API, error) {
	collectionManager, err := collections.New(
		backend.LocalDatabase(),
		backend.MetaDatabase(),
		backend.Client(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "collection api : failed to create collection API")
	}
	return &CollectionsAPI{collectionManager}, err
}

func (api *CollectionsAPI) Create(ctx context.Context, req *pb.CreateCollectionRequest) (*pb.CreateCollectionResponse, error) {
	schemaId, err := common.IDFromString(req.GetSchemaId())
	if err != nil {
		return nil, errors.Wrap(err, "invalid schema ID")
	}

	appId, err := common.IDFromString(req.GetAppId())
	if err != nil {
		return nil, errors.Wrap(err, "invalid app ID")
	}

	collection := &collections.Collection{
		AppId:    appId,
		SchemaId: schemaId,
		Policy: &collections.IncentivePolicy{
			DataProvider:  req.Policy.DataProvider,
			DataProcessor: req.Policy.DataProcessor,
			DataRelayer:   req.Policy.DataRelayer,
			DataOwner:     req.Policy.DataOwner,
		},
	}
	collectionId, err := api.collections.Register(ctx, collection)

	return &pb.CreateCollectionResponse{
		CollectionId: collectionId.String(),
	}, err
}

// TODO after localdb integration
func (api *CollectionsAPI) List(ctx context.Context, req *pb.ListCollectionRequest) (*pb.ListCollectionResponse, error) {
	return nil, nil
}

func (api *CollectionsAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterCollectionServer(service.GrpcServer, api)
}
