package api

import (
	"encoding/json"

	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/schemas"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"golang.org/x/net/context"
)

type API struct {
	schemas *schemas.Schemas
}

func New(backend *api.AirblocBackend) (api.API, error) {
	schemas, err := schemas.New(backend.MetaDatabase, backend.Ethclient, common.Address{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Schemas")
	}
	return &API{schemas}, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterSchemaServer(service.GrpcServer, api)
}

// TODO
func (api *API) Create(ctx context.Context, req *CreateSchemaRequest) (*CreateSchemaResult, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(req.Schema), &data)
	if err != nil {
		return nil, err
	}
	api.schemas.Register(ctx, req.Name, data)
	return nil, nil
}
