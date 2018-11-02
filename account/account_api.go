package account

import (
	"net"

	"context"

	"google.golang.org/grpc"
)

type API struct {
	conn    net.Conn
	server  *grpc.Server
	service *Service
}

func (api *API) Create(ctx context.Context, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	id, err := api.service.Create(ctx)
	return &AccountCreateResponse{
		AccountId: id[:],
	}, err
}

func NewAPI(conn net.Conn, service *Service) (*API, error) {
	api := &API{
		conn:    conn,
		server:  grpc.NewServer(),
		service: service,
	}
	RegisterAccountServer(api.server, api)
	return api, nil
}
