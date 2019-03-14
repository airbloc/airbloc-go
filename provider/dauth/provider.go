package dauth

import (
	"context"
	"crypto/ecdsa"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/airbloc-go/shared/account"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/p2p"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Provider is a P2P client for data provider (server) nodes to do DAuth.
// it does not interact with blockchain directly; it just requests to designated user delegate nodes.
type Provider struct {
	providerId *ecdsa.PublicKey
	accounts   *account.Manager
	p2pRpc     p2p.RPC
	log        *logger.Logger
}

// NewProviderClient creates DAuth client for data provider (server) nodes.
func NewProviderClient(kms key.Manager, client blockchain.TxClient, p2pServer p2p.Server) *Provider {
	accounts := account.NewManager(client)
	return &Provider{
		providerId: &kms.NodeKey().PublicKey,
		accounts:   accounts,
		p2pRpc:     p2p.NewRPC(p2pServer),
		log:        logger.New("dauth"),
	}
}

// SignIn creates a new account if there is no account corresponding to the identity (e.g. email),
// otherwise it returns the ID of the existing account.
func (client *Provider) SignIn(ctx context.Context, identity string, userDelegate common.Address) (types.ID, error) {
	acc, err := client.accounts.GetByIdentity(identity)
	if err != nil {
		if err == account.ErrNoAccount {
			client.log.Info("No account for %s. creating new one...", identity)
			return client.SignUp(ctx, identity, userDelegate)
		}
		return types.ID{}, errors.Wrap(err, "unable to call Accounts.GetByIdentity")
	}
	client.log.Info("SignIn(\"%s\"): Already signed up.", identity, logger.Attrs{
		"id":     acc.ID.Hex(),
		"status": acc.Status,
		"proxy":  acc.Proxy.Hex(),
	})
	return acc.ID, nil
}

// SignUp requests user delegate to create new temporary account using given identity data.
func (client *Provider) SignUp(ctx context.Context, identity string, userDelegate common.Address) (types.ID, error) {
	identityHash := client.accounts.HashIdentity(identity)
	req := &pb.DAuthSignUpRequest{
		IdentityHash: identityHash.Hex(),
	}

	// request to user delegate through Airbloc network
	reply := new(pb.DAuthSignUpResponse)
	if _, err := client.p2pRpc.Invoke(ctx, userDelegate, "dauth-signup", req, reply); err != nil {
		return types.ID{}, err
	}
	return types.HexToID(reply.GetAccountId())
}

func (client *Provider) Allow(ctx context.Context, collectionId types.ID, accountId types.ID) error {
	return client.sendDauthRequest(ctx, collectionId, accountId, "allow")
}

func (client *Provider) Deny(ctx context.Context, collectionId types.ID, accountId types.ID) error {
	return client.sendDauthRequest(ctx, collectionId, accountId, "deny")
}

func (client *Provider) sendDauthRequest(ctx context.Context, collectionId types.ID, accountId types.ID, typ string) error {
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
