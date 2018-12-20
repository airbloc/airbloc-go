package dauth

import (
	"context"
	"crypto/ecdsa"

	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/azer/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// ProviderClient is a P2P client for data provider (server) nodes to do DAuth.
// it does not interact with blockchain directly; it just requests to designated user delegate nodes.
type ProviderClient struct {
	providerId *ecdsa.PublicKey
	accounts   *account.Manager
	p2pRpc     p2p.RPC
	log        *logger.Logger
}

// NewProviderClient creates DAuth client for data provider (server) nodes.
func NewProviderClient(kms key.Manager, client blockchain.TxClient, p2pServer p2p.Server) *ProviderClient {
	accounts := account.NewManager(client)
	return &ProviderClient{
		providerId: &kms.NodeKey().PublicKey,
		accounts:   accounts,
		p2pRpc:     p2p.NewRPC(p2pServer),
		log:        logger.New("dauth"),
	}
}

// SignIn creates a new account if there is no account corresponding to the identity (e.g. email),
// otherwise it returns the ID of the existing account.
func (client *ProviderClient) SignIn(ctx context.Context, identity string, userDelegate common.Address) (ablCommon.ID, error) {
	acc, err := client.accounts.GetByIdentity(identity)
	if err != nil {
		if err == account.ErrNoAccount {
			client.log.Info("No account for %s. creating new one...", identity)
			return client.SignUp(ctx, identity, userDelegate)
		}
		return ablCommon.ID{}, errors.Wrap(err, "unable to call Accounts.GetByIdentity")
	}
	client.log.Info("SignIn(\"%s\"): Already signed up.", identity, logger.Attrs{
		"id":     acc.ID.Hex(),
		"status": acc.Status,
		"proxy":  acc.Proxy.Hex(),
	})
	return acc.ID, nil
}

// SignUp requests user delegate to create new temporary account using given identity data.
func (client *ProviderClient) SignUp(ctx context.Context, identity string, userDelegate common.Address) (ablCommon.ID, error) {
	identityHash := client.accounts.HashIdentity(identity)
	req := &pb.DAuthSignUpRequest{
		IdentityHash: identityHash.Hex(),
	}

	// request to user delegate through Airbloc network
	reply := new(pb.DAuthSignUpResponse)
	if _, err := client.p2pRpc.Invoke(ctx, userDelegate, "dauth-signup", req, reply); err != nil {
		return ablCommon.ID{}, err
	}
	return ablCommon.HexToID(reply.GetAccountId())
}

func (client *ProviderClient) Allow(ctx context.Context, collectionId ablCommon.ID, accountId ablCommon.ID) error {
	return client.sendDauthRequest(ctx, collectionId, accountId, "allow")
}

func (client *ProviderClient) Deny(ctx context.Context, collectionId ablCommon.ID, accountId ablCommon.ID) error {
	return client.sendDauthRequest(ctx, collectionId, accountId, "deny")
}

func (client *ProviderClient) sendDauthRequest(ctx context.Context, collectionId ablCommon.ID, accountId ablCommon.ID, typ string) error {
	acc, err := client.accounts.Get(accountId)
	if err != nil {
		return err
	}

	req := &pb.DAuthRequest{
		AccountId:    accountId.Hex(),
		CollectionId: collectionId.Hex(),
	}
	if _, err := client.p2pRpc.Invoke(ctx, acc.Delegate, "dauth-"+typ, req, &pb.DAuthResponse{}); err != nil {
		return errors.Wrapf(err, "failed to publish DAuth %s message", typ)
	}
	return nil
}
