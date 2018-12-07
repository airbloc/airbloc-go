package dauth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	"github.com/airbloc/airbloc-go/p2p/common"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/azer/logger"
	libp2pCrypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/pkg/errors"
)

// ProviderClient is a P2P client for data provider (server) nodes to do DAuth.
// it does not interact with blockchain directly; it just requests to designated user delegate nodes.
type ProviderClient struct {
	providerId *ecdsa.PublicKey
	accounts   *account.Manager
	p2p        p2p.Server
	log        *logger.Logger
}

// NewProviderClient creates DAuth client for data provider (server) nodes.
func NewProviderClient(kms key.Manager, client blockchain.TxClient, p2pServer p2p.Server) *ProviderClient {
	accounts := account.NewManager(client)
	return &ProviderClient{
		providerId: &kms.NodeKey().PublicKey,
		accounts:   accounts,
		p2p:        p2pServer,
		log:        logger.New("dauth"),
	}
}

// SignIn creates a new account if there is no account corresponding to the identity (e.g. email),
// otherwise it returns the ID of the existing account.
func (client *ProviderClient) SignIn(ctx context.Context, identity string, userDelegate []byte) (ablCommon.ID, error) {
	acc, err := client.accounts.GetByIdentity(identity)
	if err != nil {
		if err == account.ErrNoAccount {
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
func (client *ProviderClient) SignUp(ctx context.Context, identity string, userDelegate []byte) (ablCommon.ID, error) {
	identityHash := client.accounts.HashIdentity(identity)
	req := &pb.DAuthSignUpRequest{
		IdentityHash: identityHash.Hex(),
	}
	userDelegatePub, err := libp2pCrypto.UnmarshalSecp256k1PublicKey(userDelegate)
	if err != nil {
		return ablCommon.ID{}, errors.Wrapf(err,
			"invalid user delegate public key: %s", hex.EncodeToString(userDelegate))
	}
	destId, err := peer.IDFromPublicKey(userDelegatePub)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err,
			"failed to get libp2p peer ID from given pubkey")
	}

	// request to user delegate through Airbloc network
	if err := client.p2p.Send(ctx, req, "dauth-signup", destId); err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to send dauth-signup message")
	}

	// wait for the response. NOTE THAT this is not fault-torelant yet:
	// see https://github.com/airbloc/airbloc-go/issues/62
	waitForResponse := make(chan string, 1)
	err = client.p2p.SubscribeTopic("dauth-signup-response", &pb.DAuthSignUpResponse{}, func(server p2p.Server, c context.Context, resp common.Message) {
		if !resp.SenderInfo.ID.MatchesPublicKey(userDelegatePub) {
			return
		}
		response, ok := resp.Data.(*pb.DAuthSignUpResponse)
		if !ok {
			client.log.Error("Invalid response returned.")
			return
		}
		waitForResponse <- response.GetUserId()
	})
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to subscribe topic")
	}
	userId := <-waitForResponse
	return ablCommon.HexToID(userId)
}

func (client *ProviderClient) Allow(ctx context.Context, collectionId ablCommon.ID, accountId ablCommon.ID) error {
	return client.sendDauthRequest(ctx, collectionId, accountId, "allow")
}

func (client *ProviderClient) Deny(ctx context.Context, collectionId ablCommon.ID, accountId ablCommon.ID) error {
	return client.sendDauthRequest(ctx, collectionId, accountId, "deny")
}

func (client *ProviderClient) sendDauthRequest(ctx context.Context, collectionId ablCommon.ID, accountId ablCommon.ID, typ string) error {
	req := &pb.DAuthRequest{CollectionId: collectionId.Hex()}
	topicName := fmt.Sprintf("dauth-%s-%s", typ, accountId.Hex())

	if err := client.p2p.Publish(ctx, req, topicName); err != nil {
		return errors.Wrapf(err, "failed to publish DAuth %s message", typ)
	}
	return nil
}
