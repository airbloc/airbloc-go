package api

import (
	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type API struct {
	collections *collections.Collections
}

func New(backend *api.AirblocBackend) (api.API, error) {
	collectionManager, err := collections.New(backend.LocalDatabase, backend.Ethclient)
	return &API{collectionManager}, err
}

func (api *API) Create(ctx context.Context, req *CreateCollectionRequest) (*CreateCollectionResponse, error) {
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

	return &CreateCollectionResponse{
		CollectionId: collectionId.String(),
	}, err
}

// TODO after localdb integration
func (api *API) List(ctx context.Context, req *ListCollectionRequest) (*ListCollectionResponse, error) {
	return nil, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterCollectionServer(service.GrpcServer, api)
}
