package warehouse

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
func (s *Service) Store(stream Warehouse_StoreServer) error {
	return nil
}

func (s *Service) StoreEncrypted(stream Warehouse_StoreEncryptedServer) error {
	return nil
}

func NewService(conn net.Conn) (*Service, error) {
	service := &Service{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterWarehouseServer(service.server, service)
	return service, nil
}
