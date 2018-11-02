package schemas

import (
	"net"

	"encoding/json"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type API struct {
	conn    net.Conn
	server  *grpc.Server
	service *Service
}

func (s *API) Close() {
	s.server.Stop()
	s.conn.Close()
}

// TODO
func (s *API) Create(ctx context.Context, req *CreateSchemaRequest) (*CreateSchemaResult, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(req.Schema), &data)
	if err != nil {
		return nil, err
	}
	s.service.Register(ctx, req.Name, data)
	return nil, nil
}

//func (s *API) Delete(ctx context.Context, req *DeleteSchemaRequest) (*DeleteSchemaResult, error) {
//
//}

func NewAPI(conn net.Conn, service *Service) (*API, error) {
	api := &API{
		conn:    conn,
		server:  grpc.NewServer(),
		service: service,
	}
	RegisterSchemaServer(api.server, api)
	return api, nil
}
