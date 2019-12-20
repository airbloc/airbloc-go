package p2p

import (
	"context"
	"crypto/rand"
	"log"
	"testing"
	"time"

	"github.com/klaytn/klaytn/blockchain/types"

	"github.com/airbloc/airbloc-go/network/p2p/message"
	"github.com/airbloc/airbloc-go/network/p2p/message/users"
	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise/skademlia"

	"github.com/perlin-network/noise"

	perlinLog "github.com/perlin-network/noise/log"
	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/require"

	"github.com/airbloc/airbloc-go/account"
	"github.com/klaytn/klaytn/crypto"
)

var _ = perlinLog.Error()

func newAirblocNode(ctx context.Context, port uint16) (Node, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return Node{}, err
	}
	node, err := NewNode(ctx, "0.0.0.0", port, nil, account.NewKeyedAccount(key))
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

func newAirblocNodes(ctx context.Context, count int, initializer func(nodes []Node, index int, node Node) error) (Nodes, error) {
	nodes := make([]Node, count)
	for i := 0; i < count; i += 1 {
		node, err := newAirblocNode(ctx, 0)
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
	log.SetFlags(log.Lshortfile)

	testContext, cancel := context.WithCancel(context.Background())
	defer func() {
		time.Sleep(1 * time.Second)
		log.Println("closing ==================================================")
		cancel()
		time.Sleep(1 * time.Second)
	}()

	log.Println("==================================================")

	bootstrapNodeInitializer := func(nodes []Node, index int, node Node) error {
		node.Start()
		return nil
	}
	bootstrapNodes, err := newAirblocNodes(testContext, 1, bootstrapNodeInitializer)
	require.NoError(t, err)
	for index, node := range bootstrapNodes {
		err = node.Bootstrap(bootstrapNodes[index+1:].Addresses()...)
		require.NoError(t, err)
	}

	time.Sleep(1 * time.Second)
	log.Println("==================================================")

	nodeInitializer := func(nodes []Node, index int, node Node) error {
		node.RegisterHandler(message.OpcodeUsersSignUpRequest, func(context context.Context, message noise.Message, peer *noise.Peer) error {
			node.logger().Debug("Request received from {}!!, sending response", (<-Peer{peer}.GetAddressAsync()).Hex())
			//return peer.SendMessage(users.SignUpRequest{
			//	IdentityHash: common.HexToHash("0xdeadbeefdeadbeef"),
			//})
			return nil
		})
		node.RegisterHandler(message.OpcodeUsersSignUpResponse, func(context context.Context, message noise.Message, peer *noise.Peer) error {
			resp := message.(users.SignUpResponse)
			node.logger().Debug("Response received from {}!! value: {}", (<-Peer{peer}.GetAddressAsync()).Hex(), resp.Tx.Hash.Hex())
			return nil
		})
		return nil
	}

	nodes, err := newAirblocNodes(testContext, 3, nodeInitializer)
	require.NoError(t, err)

	for _, node := range nodes {
		node.StartWithInitialNodes(bootstrapNodes.Addresses()...)
	}

	time.Sleep(1 * time.Second)
	log.Println("==================================================")

	noise.DebugOpcodes()

	log.Println("==================================================")

	skademlia.Broadcast(nodes[0].Node, users.SignUpRequest{
		IdentityHash: common.HexToHash("0xdeadbeefdeadbeef"),
	})

	time.Sleep(1 * time.Second)
	log.Println("==================================================")

	skademlia.Broadcast(nodes[1].Node, users.SignUpResponse{
		Tx: struct {
			Hash    common.Hash    `json:"hash"`
			Receipt *types.Receipt `json:"receipt"`
		}{
			Hash:    common.HexToHash("0xbeefdeadbeefdead"),
			Receipt: types.NewReceipt(1, common.HexToHash("0xbeefdeadbeefdead"), 1612),
		},
	})
}

func TestSignature(t *testing.T) {
	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	var payload []byte
	_, err = rand.Read(payload)
	require.NoError(t, err)

	sig, err := crypto.Sign(payload, privKey)
	assert.NoError(t, err)

	pubKey, err := crypto.SigToPub(payload, sig)
	assert.NoError(t, err)
	_ = pubKey
	//crypto.VerifySignature(pubKey)
}
