package serverapi

import (
	"context"
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
)

type AccountsAPI struct {
	accounts *account.Manager
}

func NewAccountsAPI(backend node.Backend) (node.API, error) {
	return &AccountsAPI{
		accounts: account.NewManager(backend.Client()),
	}, nil
}

func (api *AccountsAPI) Exists(ctx context.Context, req *pb.AccountExistsRequest) (*pb.AccountExistsResponse, error) {
	acc, err := api.accounts.GetByIdentity(req.GetIdentity())
	if err == account.ErrNoAccount {
		return &pb.AccountExistsResponse{Exists: false}, nil
	} else if err != nil {
		return nil, err
	}
	return &pb.AccountExistsResponse{
		Exists:    true,
		AccountId: acc.ID.Hex(),
	}, nil
}

func (api *AccountsAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterAccountsServer(service.GrpcServer, api)
}
