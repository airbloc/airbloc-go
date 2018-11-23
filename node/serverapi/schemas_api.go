package serverapi

import (
	"encoding/json"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/schemas"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type SchemaAPI struct {
	schemas *schemas.Schemas
}

func NewSchemaAPI(backend node.Backend) (node.API, error) {
	schemaManager := schemas.New(backend.MetaDatabase(), backend.Client())
	return &SchemaAPI{schemaManager}, nil
}

func (api *SchemaAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterSchemaServer(service.GrpcServer, api)
}

// TODO
func (api *SchemaAPI) Create(ctx context.Context, req *pb.CreateSchemaRequest) (*pb.CreateSchemaResult, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(req.Schema), &data)
	if err != nil {
		return nil, err
	}

	id, err := api.schemas.Register(req.Name, data)
	if err == schemas.ErrNameExists {
		return &pb.CreateSchemaResult{Exists: true}, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to register schema")
	}

	return &pb.CreateSchemaResult{
		Id:     id.String(),
		Exists: false,
	}, nil
}
