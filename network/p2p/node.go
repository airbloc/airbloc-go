package p2p

import (
	"context"
	"encoding/hex"

	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/logger"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/protocol"
	"github.com/perlin-network/noise/skademlia"
	"github.com/perlin-network/noise/transport"
	"github.com/pkg/errors"
)

type Node struct {
	context contextGroup

	node *noise.Node
	acc  account.Account
	log  *logger.Logger
}

func newNode(host string, port uint16, tp transport.Layer) (*noise.Node, error) {
	param := noise.DefaultParams()
	if host != "" {
		param.Host = host
	}
	if port != 0 {
		param.Port = port
	}
	if tp != nil {
		param.Transport = tp
	}
	// TODO: derive key from account
	param.Keys = skademlia.RandomKeys()

	node, err := noise.NewNode(param)
	if err != nil {
		return nil, errors.Wrap(err, "new node")
	}
	enforceProtocol(node)
	return node, nil
}

func NewNode(host string, port uint16, tp transport.Layer, account account.Account) (Node, error) {
	node, err := newNode(host, port, tp)
	if err != nil {
		return Node{}, err
	}

	airblocNode := Node{
		node: node,
		acc:  account,
		log:  logger.New("airbloc-p2p-node"),
	}
	registerNodeEventHandler(airblocNode)

	return airblocNode, nil
}

func (n *Node) Start(parentContext context.Context) {
	// nil check context
	if parentContext == nil {
		parentContext = context.Background()
	}
	n.context = newContextGroup(context.WithCancel(parentContext))

	// listen
	go n.node.Listen()

	// start killer
	go func() {
		<-n.context.Done()
		n.node.Kill()
		n.log.Info("Node service closing...")
	}()
	n.log.Info("Node listening at {}", n.node.ExternalAddress())
}

func (n *Node) Bootstrap(nodeAddresses ...string) error {
	for _, nodeAddress := range nodeAddresses {
		peer, err := n.node.Dial(nodeAddress)
		if err != nil {
			n.Stop()
			return errors.Wrapf(err, "dialing initial node, address: %s", nodeAddress)
		}
		skademlia.WaitUntilAuthenticated(peer)
	}

	peers := skademlia.FindNode(n.node, protocol.NodeID(n.node).(skademlia.ID), skademlia.BucketSize(), 8)
	peerAddresses := make([]string, len(peers))
	for index, peer := range peers {
		peerAddresses[index] = hex.EncodeToString(peer.PublicKey())
	}
	n.log.Info("Bootstrapped with {} peers", len(peers), logger.Attrs{"peers": peerAddresses})
	return nil
}

func (n *Node) StartWithInitialNodes(parentContext context.Context, nodeAddresses ...string) error {
	n.Start(parentContext)
	return n.Bootstrap(nodeAddresses...)
}

func (n Node) Stop() {
	n.context.Cancel()
}
