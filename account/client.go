package account

import (
	"context"
	"github.com/airbloc/airbloc-go/account/api"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/key"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Client struct {
	manager api.AccountClient
}

func NewClient(conn *grpc.ClientConn) (*Client) {
	return &Client{
		manager: api.NewAccountClient(conn),
	}
}

func (client *Client) Create(walletAddress ethCommon.Address, password string) (ablCommon.ID, error) {
	identity := crypto.Keccak256Hash(walletAddress.Bytes())
	priv := key.DeriveFromPassword(identity, password)

	sig, err := crypto.Sign(identity[:], priv.PrivateKey)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to create signature by password")
	}

	request := &api.AccountCreateRequest{
		Address: walletAddress.Bytes(),
		PasswordSignature: sig,
	}

	response, err := client.manager.Create(context.Background(), request)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "RPC call failed")
	}
	return ablCommon.IDFromString(response.GetAccountId())
}
