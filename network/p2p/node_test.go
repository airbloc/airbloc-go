package p2p

import (
	"context"
	"crypto/rand"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/airbloc/airbloc-go/network/p2p/message/users"
	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise/skademlia"

	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/require"

	"github.com/airbloc/airbloc-go/account"
	"github.com/klaytn/klaytn/crypto"
)

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
		addresses = append(addresses, node.node.ExternalAddress())
	}
	return
}

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
	//perlinLog.Disable()
	log.SetFlags(log.Llongfile)

	testContext, cancel := context.WithCancel(context.Background())
	defer func() {
		time.Sleep(1 * time.Second)
		log.Println("closing ==================================================")
		cancel()
		time.Sleep(1 * time.Second)
	}()

	bootstrapNodeInitializer := func(nodes []Node, index int, node Node) error {
		node.Start(testContext)
		return nil
	}
	bootstrapNodes, err := newAirblocNodes(2, bootstrapNodeInitializer)
	require.NoError(t, err)
	for index, node := range bootstrapNodes {
		err = node.Bootstrap(bootstrapNodes[index+1:].Addresses()...)
		require.NoError(t, err)
	}

	nodes, err := newAirblocNodes(7, nil)
	require.NoError(t, err)

	nodeWaitGroup := new(sync.WaitGroup)
	for _, node := range nodes {
		nodeWaitGroup.Add(1)
		go func(node Node) {
			defer nodeWaitGroup.Done()
			err = node.StartWithInitialNodes(testContext, bootstrapNodes.Addresses()...)
			assert.NoError(t, err)
		}(node)
	}
	nodeWaitGroup.Wait()

	skademlia.Broadcast(nodes[0].node, users.SignUpRequest{
		IdentityHash: common.HexToHash("0xdeadbeefdeadbeef"),
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
