package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/provider/schemas"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"golang.org/x/net/context"
)

type SchemaAPI struct {
	schemas *schemas.Schemas
}

func NewSchemaAPI(backend service.Backend) (api.API, error) {
	schemaManager := schemas.New(backend.MetaDatabase(), backend.Client())
	return &SchemaAPI{schemaManager}, nil
}

func (api *SchemaAPI) AttachToAPI(service *api.Service) {
	pb.RegisterSchemaServer(service.GrpcServer, api)
}

// TODO
func (api *SchemaAPI) Create(ctx context.Context, req *pb.CreateSchemaRequest) (*pb.CreateSchemaResult, error) {
	schema, err := schemas.NewSchema(req.GetName(), req.GetSchema())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid JSON schema: %s", err.Error())
	}

	id, err := api.schemas.Register(schema)
	if err == schemas.ErrNameExists {
		return &pb.CreateSchemaResult{Exists: true}, nil
	}
	if err != nil {
		return nil, err
	}
	return &pb.CreateSchemaResult{
		Id:     id.Hex(),
		Exists: false,
	}, nil
}
