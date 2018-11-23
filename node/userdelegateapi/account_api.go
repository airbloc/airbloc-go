package userdelegateapi

import (
	"context"

	"github.com/airbloc/airbloc-go/account"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/userdelegate"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type AccountAPI struct {
	manager *account.Manager
}

func NewAccountAPI(backend node.Backend) (node.API, error) {
	manager, err := account.NewManager(backend.Client())
	return &AccountAPI{manager}, err
}

func (api *AccountAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterAccountServer(service.GrpcServer, api)
}

func (api *AccountAPI) Create(
	ctx context.Context,
	req *pb.AccountCreateRequest,
) (*pb.AccountCreateResponse, error) {
	address := ethCommon.BytesToAddress(req.GetAddress())
	id, err := api.manager.CreateUsingProxy(address, req.GetPasswordSignature())
	return &pb.AccountCreateResponse{AccountId: id.String()}, err
}

func (api *AccountAPI) Get(ctx context.Context, req *pb.AccountGetRequest) (*pb.AccountGetResponse, error) {
	accountId, err := ablCommon.IDFromString(req.GetAccountId())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse account ID: %s", req.GetAccountId())
	}

	acc, err := api.manager.Get(accountId)
	if err != nil {
		return nil, err
	}

	return &pb.AccountGetResponse{
		AccountId:     acc.ID.String(),
		OwnerAddress:  acc.Owner.Bytes(),
		Status:        pb.AccountGetResponse_Status(acc.Status),
		ProxyAddress:  acc.Proxy.Bytes(),
		PasswordProof: acc.PasswordProof.Bytes(),
	}, nil
}

func (api *AccountAPI) GetByIdentity(
	ctx context.Context,
	req *pb.AccountGetByIdentityRequest,
) (*pb.AccountGetResponse, error) {
	acc, err := api.manager.GetByIdentity(req.GetIdentity())
	if err != nil {
		return nil, err
	}

	return &pb.AccountGetResponse{
		AccountId:     acc.ID.String(),
		OwnerAddress:  acc.Owner.Bytes(),
		Status:        pb.AccountGetResponse_Status(acc.Status),
		ProxyAddress:  acc.Proxy.Bytes(),
		PasswordProof: acc.PasswordProof.Bytes(),
	}, nil
}

func (api *AccountAPI) TestPassword(ctx context.Context, req *pb.TestPasswordRequest) (*pb.TestPasswordResponse, error) {
	exists, err := api.manager.TestPassword(
		ethCommon.BytesToHash(req.GetMessageHash()),
		req.GetSignature())

	return &pb.TestPasswordResponse{Exists: exists}, err
}
