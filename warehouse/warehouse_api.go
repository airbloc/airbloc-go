package warehouse

import (
	"context"
	"net"

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
func (s *API) StoreBundle(stream Warehouse_StoreBundleServer) error {
	return nil
}

func (s *API) StoreEncryptedBundle(stream Warehouse_StoreEncryptedBundleServer) error {
	return nil
}

func (s *API) DeleteBundle(context context.Context, request *DeleteBundleRequest) (*DeleteBundleResult, error) {
	return nil, nil
}

func NewService(conn net.Conn) (*API, error) {
	service := &API{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterWarehouseServer(service.server, service)
	return service, nil
}
