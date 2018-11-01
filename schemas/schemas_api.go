package schemas

import (
	"net"

	"encoding/json"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type API struct {
	conn    net.Conn
	service *Service
	server  *grpc.Server
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

func NewAPI(conn net.Conn, service *Service) (*API, error) {
	api := &API{
		conn:    conn,
		service: service,
		server:  grpc.NewServer(),
	}
	RegisterSchemaServer(api.server, api)
	return api, nil
}
