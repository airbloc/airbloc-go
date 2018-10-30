package schemas

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Service struct {
	conn   net.Conn
	server *grpc.Server
}

func (s *Service) Close() {
	s.server.Stop()
	s.conn.Close()
}

// TODO
func (s *Service) Create(ctx context.Context, req *CreateSchemaRequest) (*CreateSchemaResult, error) {
	return nil, nil
}

func NewService(conn net.Conn) (*Service, error) {
	service := &Service{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterSchemaServiceServer(service.server, service)
	return service, nil
}
