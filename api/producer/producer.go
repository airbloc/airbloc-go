package producer

import (
	"net"

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
func (s *Service) AddData(req Producer_AddDataServer) error {
	return nil
}

func NewService(conn net.Conn) (*Service, error) {
	service := &Service{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterProducerServer(service.server, service)
	return service, nil
}
