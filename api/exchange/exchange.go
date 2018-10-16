package exchange

import (
	"net"

	"github.com/airbloc/airbloc-go/api/common"
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
func (s *Service) Order(ctx context.Context, req *OrderRequest) (*OrderId, error) {
	return nil, nil
}

func (s *Service) Settle(ctx context.Context, id *OrderId) (*SettleResult, error) {
	return nil, nil
}

func (s *Service) Reject(ctx context.Context, id *OrderId) (*common.Result, error) {
	return nil, nil
}

func NewService(conn net.Conn) (*Service, error) {
	service := &Service{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterExchangeServiceServer(service.server, service)
	return service, nil
}
