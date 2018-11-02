package schemas

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type API struct {
	conn   net.Conn
	server *grpc.Server
}

func (s *API) Close() {
	s.server.Stop()
	s.conn.Close()
}

// TODO
func (s *API) Create(ctx context.Context, req *CreateSchemaRequest) (*CreateSchemaResult, error) {
	return nil, nil
}

func NewService(conn net.Conn) (*API, error) {
	service := &API{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterSchemaServer(service.server, service)
	return service, nil
}
