package api

import (
	"encoding/json"
	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/schemas"
	"golang.org/x/net/context"
)

type API struct {
	schemas *schemas.Schemas
}

func New(backend *api.AirblocBackend) (api.API, error) {
	schemaManager := schemas.New(backend.MetaDatabase, backend.Ethclient)
	return &API{schemaManager}, nil
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
