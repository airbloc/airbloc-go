package p2p

import (
	"context"
	"testing"
	"time"

	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/network/p2p/message"
	"github.com/airbloc/airbloc-go/network/p2p/message/users"
	"github.com/airbloc/logger"

	"github.com/klaytn/klaytn/crypto"
	"github.com/perlin-network/noise"
	perlinLog "github.com/perlin-network/noise/log"
	"github.com/perlin-network/noise/skademlia"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var _ = perlinLog.Error()

func newAirblocNode(port uint16) (Node, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return Node{}, err
	}
	node, err := NewNode("0.0.0.0", port, nil, account.NewKeyedAccount(key))
	if err != nil {
		return Node{}, err
	}
	return node, nil
}

type Nodes []Node

func (nodes Nodes) Addresses() (addresses []string) {
	for _, node := range nodes {
		addresses = append(addresses, node.ExternalAddress())
	}
	return
}

func (nodes Nodes) Close() {
	for _, node := range nodes {
		node.Stop()
	}
}

type l struct{}

func (l) Init()             {}
func (l) Write(*logger.Log) {}

func newAirblocNodes(count int, initializer func(nodes []Node, index int, node Node) error) (Nodes, error) {
	nodes := make([]Node, count)
	for i := 0; i < count; i += 1 {
		node, err := newAirblocNode(0)
		if err != nil {
			return nil, errors.Wrap(err, "making new airbloc node")
		}
		if initializer != nil {
			err = initializer(nodes, i, node)
			if err != nil {
				return nil, errors.Wrap(err, "initializing airbloc node")
			}
		}
		nodes[i] = node
	}
	return nodes, nil
}

func TestAirblocNode(t *testing.T) {
	perlinLog.Disable()
	logger.SetLogger(l{})

	testContext, cancel := context.WithCancel(context.Background())
	defer func() { cancel() }()

	bootstrapNodeInitializer := func(nodes []Node, index int, node Node) error {
		node.Start(testContext)
		return nil
	}
	bootstrapNodes, err := newAirblocNodes(1, bootstrapNodeInitializer)
	require.NoError(t, err)
	for index, node := range bootstrapNodes {
		err = node.Bootstrap(bootstrapNodes[index+1:].Addresses()...)
		require.NoError(t, err)
	}

	time.Sleep(1 * time.Second)

	nodeInitializer := func(nodes []Node, index int, node Node) error {
		node.Set(KeyNodeLogger, logger.New("abl-p2p"+"/"+node.account().Address().Hex()))
		node.RegisterHandler(message.OpcodeUsersSignUpRequest, func(context context.Context, message message.Message, peer *noise.Peer) error {
			reqMsg := message.(*users.SignUpRequest)

			signature, err := node.account().SignMessage(reqMsg.IdentityHash.Bytes())
			if err != nil {
				return errors.Wrap(err, "failed to sign response message")
			}

			return peer.SendMessage(users.SignUpResponse{
				MessageId: reqMsg.MessageId,
				TxHash:    reqMsg.IdentityHash,
				Sign:      signature,
			})
		})
		node.RegisterHandler(message.OpcodeUsersSignUpResponse, func(context context.Context, message message.Message, peer *noise.Peer) error {
			peerAddr := <-Peer{peer}.GetAddressAsync()
			respMsg := message.(*users.SignUpResponse)

			pubKey, err := crypto.SigToPub(respMsg.TxHash.Bytes(), respMsg.Sign)
			if err != nil {
				return errors.Wrap(err, "failed to derive pubkey from signature")
			}
			if crypto.PubkeyToAddress(*pubKey) != peerAddr {
				return errors.Wrap(err, "failed to validate response")
			}

			return nil
		})
		return nil
	}

	nodes, err := newAirblocNodes(3, nodeInitializer)
	require.NoError(t, err)

	for _, node := range nodes {
		require.NoError(t, node.StartWithInitialNodes(testContext, bootstrapNodes.Addresses()...))
	}
	time.Sleep(1 * time.Second)

	skademlia.Broadcast(nodes[0].Node, users.SignUpRequest{IdentityHash: nodes[0].account().Address().Hash()})
}
