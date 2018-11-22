package api

import (
	"context"

	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/api"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type API struct {
	manager *account.Manager
}

func New(backend api.Backend) (api.API, error) {
	manager, err := account.NewManager(backend.Client())
	if err != nil {
		return nil, errors.Wrap(err, "account api : failed to create account API")
	}
	return &API{manager}, err
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterAccountServer(service.GrpcServer, api)
}

func (api *API) Create(ctx context.Context, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	address := ethCommon.BytesToAddress(req.GetAddress())
	id, err := api.manager.CreateUsingProxy(address, req.GetPasswordSignature())
	return &AccountCreateResponse{AccountId: id.String()}, err
}

func (api *API) Get(ctx context.Context, req *AccountGetRequest) (*AccountGetResponse, error) {
	accountId, err := ablCommon.IDFromString(req.GetAccountId())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse account ID: %s", req.GetAccountId())
	}

	acc, err := api.manager.Get(accountId)
	if err != nil {
		return nil, err
	}

	return &AccountGetResponse{
		AccountId:     acc.ID.String(),
		OwnerAddress:  acc.Owner.Bytes(),
		Status:        AccountGetResponse_Status(acc.Status),
		ProxyAddress:  acc.Proxy.Bytes(),
		PasswordProof: acc.PasswordProof.Bytes(),
	}, nil
}

func (api *API) GetByIdentity(ctx context.Context, req *AccountGetByIdentityRequest) (*AccountGetResponse, error) {
	acc, err := api.manager.GetByIdentity(req.GetIdentity())
	if err != nil {
		return nil, err
	}

	return &AccountGetResponse{
		AccountId:     acc.ID.String(),
		OwnerAddress:  acc.Owner.Bytes(),
		Status:        AccountGetResponse_Status(acc.Status),
		ProxyAddress:  acc.Proxy.Bytes(),
		PasswordProof: acc.PasswordProof.Bytes(),
	}, nil
}

func (api *API) TestPassword(ctx context.Context, req *TestPasswordRequest) (*TestPasswordResponse, error) {
	exists, err := api.manager.TestPassword(
		ethCommon.BytesToHash(req.GetMessageHash()),
		req.GetSignature())

	return &TestPasswordResponse{Exists: exists}, err
}
