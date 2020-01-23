package p2p

import (
	"context"
	"encoding/hex"
	"sync"

	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/network/p2p/handshake/identity"
	"github.com/airbloc/airbloc-go/network/p2p/message"
	"github.com/airbloc/logger"

	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/cipher/aead"
	"github.com/perlin-network/noise/handshake/ecdh"
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
	node.Set(KeyNodeHandlers, new(sync.Map))
	// TODO - remove postfix address <- for debug perpose
	node.Set(KeyNodeLogger, logger.New("abl-p2p"+"/"+account.Address().Hex()))
	node.Set(KeyNodePeerStore, new(sync.Map))

	registerNodeEventHandler(node)

	return node, nil
}

//======== account
func (n Node) account() account.Account {
	return n.Get(KeyNodeAccount).(account.Account)
}

//======== handlers
func (n Node) context() Context {
	return n.Get(KeyNodeContext).(Context)
}

//======== handlers
func (n Node) handlers() *sync.Map {
	return n.Get(KeyNodeHandlers).(*sync.Map)
}

func (n Node) handle(opcode noise.Opcode, msg message.Message, peer *noise.Peer) error {
	handler, exist := n.handlers().Load(opcode)
	if !exist {
		return errors.New("handler that matches given opcode does not registered")
	}
	return handler.(HandlerFunc)(n.context(), msg, peer)
}

func (n Node) RegisterHandler(opcode noise.Opcode, handler HandlerFunc) {
	n.handlers().Store(opcode, handler)
}

//======== logger
func (n Node) logger() *logger.Logger {
	return n.Get(KeyNodeLogger).(*logger.Logger)
}

//======== peerstore
func (n Node) peerstore() *sync.Map {
	return n.Get(KeyNodePeerStore).(*sync.Map)
}

func (n Node) registerPeer(address common.Address, peer *noise.Peer) (exist bool) {
	_, exist = n.peerstore().LoadOrStore(address, peer)
	if !exist {
		n.logger().Debug("Peer registered, address is {}", address.Hex())
	} else {
		n.logger().Debug("Peer duplicated, disconnecting {}", address.Hex())
	}
	return
}

func (n Node) unregisterPeer(address common.Address) {
	if _, exist := n.peerstore().Load(address); exist {
		n.peerstore().Delete(address)
		n.logger().Debug("Peer unregistered, address is {}", address.Hex())
	}
}

func (n Node) Start(parentContext context.Context) {
	n.Set(KeyNodeContext, newContext(context.WithCancel(parentContext)))

	// listen
	go n.Listen()

	// start killer
	go func() {
		<-n.context().Done()
		n.logger().Info("Node service closing...")
		n.peerstore().Range(func(key, value interface{}) bool {
			value.(*noise.Peer).Disconnect()
			return true
		})
		n.Kill()
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

func (n Node) StartWithInitialNodes(parentContext context.Context, nodeAddresses ...string) error {
	n.Start(parentContext)
	return n.Bootstrap(nodeAddresses...)
}

// Broadcast broadcasts message to other nodes and returns all peer count, succeeded message count, error
func (n Node) Broadcast(ctx context.Context, message message.Message) (int, error) {
	var (
		errChans     []<-chan error
		successCount = 0

		nodeTbl = skademlia.Table(n.Node)
		nodeID  = protocol.NodeID(n.Node).Hash()
	)

	closestPeers := skademlia.FindClosestPeers(nodeTbl, nodeID, skademlia.BucketSize())
	if len(closestPeers) == 0 {
		return 0, errors.New("there are no peers in network")
	}

	for _, peerID := range closestPeers {
		peer := protocol.Peer(n.Node, peerID)

		if peer == nil {
			continue
		}

		errChans = append(errChans, peer.SendMessageAsync(message))
	}

	for _, ch := range errChans {
		select {
		case <-ctx.Done():
			return successCount, ctx.Err()
		case err := <-ch:
			if err == nil {
				successCount += 1
			}
		}
	}
	return successCount, nil
}

func (n Node) Stop() {
	n.context().Cancel()
}
