package exchange

import (
	"net"

	"github.com/airbloc/airbloc-go/common"
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
func (s *API) Order(ctx context.Context, req *OrderRequest) (*common.Hash, error) {
	return nil, nil
}

func (s *API) Settle(ctx context.Context, req *SettleMessage) (*common.Result, error) {
	return nil, nil
}

func (s *API) Reject(ctx context.Context, id *OrderId) (*common.Result, error) {
	return nil, nil
}

func (s *API) CloseOrder(ctx context.Context, id *OrderId) (*common.Result, error) {
	return nil, nil
}

func NewAPI(conn net.Conn) (*API, error) {
	api := &API{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterExchangeServer(api.server, api)
	return api, nil
}
