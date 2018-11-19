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

var (
	ErrInvalidPassword = errors.New("invalid password.")
)

type Client struct {
	manager api.AccountClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		manager: api.NewAccountClient(conn),
	}
}

func (client *Client) Create(walletAddress ethCommon.Address, password string) (*Session, error) {
	identity := crypto.Keccak256Hash(walletAddress.Bytes())
	priv := key.DeriveFromPassword(identity, password)

	sig, err := crypto.Sign(identity[:], priv.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create signature by password")
	}

	request := &api.AccountCreateRequest{
		Address:           walletAddress.Bytes(),
		PasswordSignature: sig,
	}

	response, err := client.manager.Create(context.Background(), request)
	if err != nil {
		return nil, errors.Wrap(err, "RPC call failed")
	}
	accountId, err := ablCommon.IDFromString(response.GetAccountId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid ID returned from the server: %s", response.GetAccountId())
	}
	return &Session{
		AccountId:     accountId,
		WalletAddress: walletAddress,
		Key:           priv,
	}, nil
}

func (client *Client) LogIn(identity string, password string) (*Session, error) {
	request := &api.AccountGetByIdentityRequest{
		Identity: identity,
	}
	response, err := client.manager.GetByIdentity(context.Background(), request)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to log in: %s", identity)
	}
	accountId, err := ablCommon.IDFromString(response.GetAccountId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid ID returned from the server: %s", response.GetAccountId())
	}
	session := newSession(accountId, ethCommon.BytesToAddress(response.GetOwnerAddress()), password)

	// generate test signature
	identityHash := crypto.Keccak256Hash([]byte(identity))
	sig, err := session.Sign(identityHash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to sign using password key")
	}

	// test the signature to check whether the password is correct
	testReq := &api.TestPasswordRequest{
		MessageHash: identityHash[:],
		Signature:   sig,
	}
	testResp, err := client.manager.TestPassword(context.Background(), testReq)
	if err != nil {
		return nil, errors.Wrapf(err, "RPC call TestPassword failed")
	} else if !testResp.Exists {
		return nil, ErrInvalidPassword
	}
	return session, nil
}
