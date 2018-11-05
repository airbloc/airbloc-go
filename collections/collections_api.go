package collections

import (
	"github.com/airbloc/airbloc-go/api"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
)

type API struct {
	collections *Collections
}

func NewAPI(backend *api.AirblocBackend) (api.API, error) {
	collections, err := New(backend.LocalDatabase, backend.Ethclient, nil, common.Address{})
	return &API{collections}, err
}

func (api *API) Create(ctx context.Context, req *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	hash, err := api.collections.Register(ctx, &Collection{
		AppId:    common.HexToHash(req.AppId),
		SchemaId: common.HexToHash(req.SchemaId),
		Policy: &IncentivePolicy{
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
