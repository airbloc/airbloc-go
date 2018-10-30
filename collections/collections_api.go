package collections

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Service struct {
	conn   net.Conn
	server *grpc.Server
}

func (cs *Service) Close() {
	cs.server.Stop()
	cs.conn.Close()
}

// TODO
func (cs *Service) Create(ctx context.Context, req *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	return nil, nil
}

func (cs *Service) List(ctx context.Context, req *ListCollectionRequest) (*ListCollectionResponse, error) {
	return nil, nil
}

func NewService(conn net.Conn) (*Service, error) {
	service := &Service{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterCollectionServer(service.server, service)
	return service, nil
}
