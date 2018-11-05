package schemas

import (
	"encoding/json"
	"github.com/airbloc/airbloc-go/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"golang.org/x/net/context"
)

type API struct {
	schemas *Schemas
}

func NewAPI(backend *api.AirblocBackend) (api.API, error) {
	schemas, err := New(backend.MetaDatabase, backend.Ethclient, common.Address{})
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
	s.schemas.Register(ctx, req.Name, data)
	return nil, nil
}
