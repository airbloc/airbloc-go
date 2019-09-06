package dauth

import (
	"context"
	"crypto/ecdsa"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/p2p"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/pkg/errors"
)

// Client is a P2P client for data provider (server) nodes to do DAuth.
// it does not interact with blockchain directly; it just requests to designated user delegate nodes.
type Client struct {
	providerId *ecdsa.PublicKey
	accounts   adapter.IAccountsManager
	p2pRpc     p2p.RPC
	log        *logger.Logger
}

// NewProviderClient creates DAuth client for data provider (server) nodes.
func NewProviderClient(kms key.Manager, client blockchain.TxClient, p2pServer p2p.Server) *Client {
	return &Client{
		providerId: &kms.NodeKey().PublicKey,
		accounts:   adapter.NewAccountsManager(client),
		p2pRpc:     p2p.NewRPC(p2pServer),
		log:        logger.New("dauth"),
	}
}

func hashIdentity(identity string, saltHash common.Hash) common.Hash {
	identityPreimage := crypto.Keccak256Hash([]byte(identity))
	identityHash := crypto.Keccak256Hash(append(identityPreimage[:], saltHash[:]...))
	return identityHash
}

// SignIn creates a new account if there is no account corresponding to the identity (e.g. email),
// otherwise it returns the ID of the existing account.
func (client *Client) SignIn(
	ctx context.Context,
	identity string,
	controller common.Address,
) (types.ID, error) {
	identityHash := hashIdentity(identity, common.Hash{})
	accountId, err := client.accounts.GetAccountIdByIdentityHash(identityHash)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "unable to call Accounts.GetAccountByIdentityHash")
	}

	exists, err := client.accounts.Exists(accountId)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "unable to call Accounts.Exists")
	}

	if exists {
		acc, err := client.accounts.GetAccount(accountId)
		if err != nil {
			return types.ID{}, errors.Wrap(err, "unable to call Accounts.GetByIdentity")
		}

		client.log.Info("SignIn(\"%s\"): Already signed up.", identity, logger.Attrs{
			"account-id": accountId.Hex(),
			"status":     acc.Status,
			"controller": acc.Controller.Hex(),
		})

		return accountId, nil
	}

	client.log.Info("No account for %s. creating new one...", identity)
	return client.SignUp(ctx, identity, controller)
}

// SignUp requests user delegate to create new temporary account using given identity data.
func (client *Client) SignUp(
	ctx context.Context,
	identity string,
	controller common.Address,
) (types.ID, error) {
	identityHash := hashIdentity(identity, common.Hash{})
	req := &pb.DAuthSignUpRequest{
		IdentityHash: identityHash.Hex(),
	}

	// request to user delegate through Airbloc network
	reply := new(pb.DAuthSignUpResponse)
	if _, err := client.p2pRpc.Invoke(ctx, controller, "dauth-signup", req, reply); err != nil {
		return types.ID{}, err
	}

	return types.HexToID(reply.GetAccountId())
}

// Allow requests data controller to consent at given dataType
func (client *Client) Allow(
	ctx context.Context,
	accountId types.ID,
	dataType string,
	action uint8,
	appName string,
) error {
	return client.sendDauthRequest(ctx, appName, action, dataType, accountId, "allow")
}

// Deny requests data controller to consent at given dataType
func (client *Client) Deny(
	ctx context.Context,
	accountId types.ID,
	dataType string,
	action uint8,
	appName string,
) error {
	return client.sendDauthRequest(ctx, appName, action, dataType, accountId, "deny")
}

func (client *Client) Many(
	ctx context.Context,
	accountId types.ID,
	appName string,
	consentData []types.ConsentData,
) error {
	acc, err := client.accounts.GetAccount(accountId)
	if err != nil {
		return err
	}

	consentRawData := make([]*pb.DAuthManyRequest_ConsentData, len(consentData))
	for index, consent := range consentData {
		consentRawData[index] = &pb.DAuthManyRequest_ConsentData{
			Action:   uint32(consent.Action),
			DataType: consent.DataType,
			Allow:    consent.Allow,
		}
	}

	req := &pb.DAuthManyRequest{
		AppName:     appName,
		AccountId:   accountId.Hex(),
		ConsentData: consentRawData,
	}

	_, err = client.p2pRpc.Invoke(ctx, acc.Controller, "dauth-many", req, &pb.DAuthResponse{})
	if err != nil {
		return errors.Wrapf(err, "failed to publish DAuth many message")
	}
	return nil
}

func (client *Client) sendDauthRequest(
	ctx context.Context,
	appName string,
	action uint8,
	dataType string,
	accountId types.ID,
	messageType string,
) error {
	acc, err := client.accounts.GetAccount(accountId)
	if err != nil {
		return err
	}

	req := &pb.DAuthRequest{
		AccountId: accountId.Hex(),
		DataType:  dataType,
		Action:    uint32(action),
		AppName:   appName,
	}

	_, err = client.p2pRpc.Invoke(ctx, acc.Controller, "dauth-"+messageType, req, &pb.DAuthResponse{})
	if err != nil {
		return errors.Wrapf(err, "failed to publish DAuth %s message", messageType)
	}
	return nil
}
