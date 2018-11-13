package api

import (
	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/collections"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/common"
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
	schemaId, err := ablCommon.IDFromString(req.SchemaId)
	if err != nil {
		return nil, errors.Wrap(err, "invalid schema ID")
	}

	hash, err := api.collections.Register(ctx, &collections.Collection{
		AppId:    common.HexToHash(req.AppId),
		SchemaId: schemaId,
		Policy: &collections.IncentivePolicy{
			DataProducer:  req.Policy.DataProducer,
			DataProcessor: req.Policy.DataProcessor,
			DataRelayer:   req.Policy.DataRelayer,
			DataSource:    req.Policy.DataSource,
		},
	})
	return &CreateCollectionResponse{
		CollectionId: hash.Hex(),
	}, err
}

// TODO after localdb integration
func (api *API) List(ctx context.Context, req *ListCollectionRequest) (*ListCollectionResponse, error) {
	return nil, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterCollectionServer(service.GrpcServer, api)
}
