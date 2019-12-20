package p2p

import (
	"context"
	"encoding/hex"
	"sync"

	"github.com/airbloc/airbloc-go/network/p2p/message"

	"github.com/klaytn/klaytn/common"

	"github.com/airbloc/airbloc-go/network/p2p/handshake/identity"

	"github.com/perlin-network/noise/cipher/aead"
	"github.com/perlin-network/noise/handshake/ecdh"

	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/logger"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/protocol"
	"github.com/perlin-network/noise/skademlia"
	"github.com/perlin-network/noise/transport"
	"github.com/pkg/errors"
)

const (
	KeyNodeAccount   = "abl.acc"
	KeyNodeContext   = "abl.ctx"
	KeyNodeHandlers  = "abl.handlers"
	KeyNodeLogger    = "abl.logger"
	KeyNodePeerStore = "abl.peerstore"

	ecdhHandshakeMessage = ".airbloc_node_handshake"
)

type Node struct{ *noise.Node }

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
	return node, nil
}

func NewNode(
	parentContext context.Context,
	host string, port uint16, tp transport.Layer,
	account account.Account,
) (Node, error) {
	n, err := newNode(host, port, tp)
	if err != nil {
		return Node{}, err
	}
	node := Node{n}

	ecdhBlock := ecdh.New()
	ecdhBlock.WithHandshakeMessage(ecdhHandshakeMessage)

	p := protocol.New()
	p.Register(ecdhBlock)
	p.Register(aead.New())
	p.Register(skademlia.New())
	p.Register(identity.New(account))
	p.Enforce(node.Node)

	node.Set(KeyNodeAccount, account)
	node.Set(KeyNodeContext, newContext(context.WithCancel(parentContext)))
	node.Set(KeyNodeHandlers, new(sync.Map))
	// TODO - remove postfix address
	node.Set(KeyNodeLogger, logger.New("abl-p2p"+"/"+account.Address().Hex()))
	node.Set(KeyNodePeerStore, new(sync.Map))

	registerNodeEventHandler(node)

	return node, nil
}

func (n Node) account() account.Account {
	return n.Get(KeyNodeAccount).(account.Account)
}

func (n Node) context() Context {
	return n.Get(KeyNodeContext).(Context)
}

func (n Node) handlers() *sync.Map {
	return n.Get(KeyNodeHandlers).(*sync.Map)
}

func (n Node) logger() *logger.Logger {
	return n.Get(KeyNodeLogger).(*logger.Logger)
}

func (n Node) peerstore() *sync.Map {
	return n.Get(KeyNodePeerStore).(*sync.Map)
}

func (n Node) RegisterHandler(opcode noise.Opcode, handler message.HandlerFunc) {
	n.handlers().Store(opcode, handler)
}

func (n Node) GetHandler(opcode noise.Opcode) (message.HandlerFunc, error) {
	handler, exist := n.handlers().Load(opcode)
	if !exist {
		return nil, errors.New("handler that matches given opcode does not registered")
	}
	return handler.(message.HandlerFunc), nil
}

func (n Node) RegisterPeer(address common.Address, peer *noise.Peer) (exist bool) {
	_, exist = n.peerstore().LoadOrStore(address, peer)
	if !exist {
		n.logger().Debug("Peer registered, address is {}", address.Hex())
	} else {
		n.logger().Debug("Peer duplicated, disconnecting {}", address.Hex())
	}
	return
}

func (n Node) UnregisterPeer(address common.Address) {
	if _, exist := n.peerstore().Load(address); exist {
		n.peerstore().Delete(address)
		n.logger().Debug("Peer unregistered, address is {}", address.Hex())
	}
}

func (n Node) Start() {
	// listen
	go n.Listen()

	// start killer
	go func() {
		<-n.context().Done()
		n.peerstore().Range(func(key, value interface{}) bool {
			value.(*noise.Peer).Disconnect()
			return true
		})
		n.Kill()
		n.logger().Info("Node service closing...")
	}()
	n.logger().Info("Node listening at {}", n.ExternalAddress())
}

func (n Node) Bootstrap(nodeAddresses ...string) error {
	for _, nodeAddress := range nodeAddresses {
		peer, err := n.Dial(nodeAddress)
		if err != nil {
			n.Stop()
			return errors.Wrapf(err, "dialing initial node, address: %s", nodeAddress)
		}
		skademlia.WaitUntilAuthenticated(peer)
	}

	peers := skademlia.FindNode(n.Node, protocol.NodeID(n.Node).(skademlia.ID), skademlia.BucketSize(), 8)
	peerAddresses := make([]string, len(peers))
	for index, peer := range peers {
		peerAddresses[index] = hex.EncodeToString(peer.PublicKey())
	}
	n.logger().Info("Bootstrapped with {} peers", len(peers), logger.Attrs{"peers": peerAddresses})
	return nil
}

func (n Node) StartWithInitialNodes(nodeAddresses ...string) error {
	n.Start()
	return n.Bootstrap(nodeAddresses...)
}

func (n Node) Stop() {
	n.context().Cancel()
}
