package p2p

import (
	"context"
	"reflect"
	"strings"
	"sync"

	"github.com/airbloc/logger"

	"github.com/perlin-network/noise"
	"github.com/pkg/errors"
)

type peerEventHandler struct{ node Node }

func registerPeerEventHandler(node Node, peer *noise.Peer) {
	peerEventHandler{node}.register(peer)
}

func (handler peerEventHandler) register(peer *noise.Peer) {
	peer.OnConnError(handler.OnConnErrorHandler())
}

func (handler peerEventHandler) OnConnErrorHandler() noise.OnPeerErrorCallback {
	return func(node *noise.Node, peer *noise.Peer, err error) error {
		return err
	}
}

type nodeEventHandler struct{ node Node }

func registerNodeEventHandler(node Node) {
	nodeEventHandler{node}.register(node.Node)
}

func (handler nodeEventHandler) register(node *noise.Node) {
	node.OnListenerError(handler.onListenerErrorHandler())
	node.OnPeerInit(handler.onPeerInitHandler())
	node.OnPeerDisconnected(handler.onPeerDisconnectedHandler())
}

func (handler nodeEventHandler) onListenerErrorHandler() noise.OnErrorCallback {
	return func(_ *noise.Node, err error) error {
		if strings.Contains(err.Error(), "use of closed network connection") {
			return nil
		}

		node := handler.node
		node.logger().Error("Listener error occurred", logger.Attrs{
			"type":    reflect.TypeOf(errors.Cause(err)).String(),
			"message": err.Error(),
		})
		return err
	}
}

func (handler nodeEventHandler) onPeerInitHandler() noise.OnPeerInitCallback {
	return func(node *noise.Node, p *noise.Peer) error {
		peer := Peer{p}
		peer.Set(KeyPeerAggregatedMessageChannel, make(chan aggregatedMessage, 5))
		peer.Set(KeyPeerContext, newContext(context.WithCancel(context.Background())))
		peer.Set(KeyPeerWaitGroupMessageWorker, new(sync.WaitGroup))
		peer.Set(KeyPeerWaitGroupMessageAggregator, new(sync.WaitGroup))

		registerPeerEventHandler(handler.node, peer.Peer)

		peerAddress, err := peer.GetAddress()
		if err != nil {
			peer.Disconnect()
			return err
		}
		exist := handler.node.registerPeer(peerAddress, p)
		if exist {
			peer.Disconnect()
			return nil
		}

		go MessageAggregator(handler.node, peer)
		go MessageHandler(handler.node, peer)

		return nil
	}
}

func (handler nodeEventHandler) onPeerDisconnectedHandler() noise.OnPeerDisconnectCallback {
	return func(node *noise.Node, p *noise.Peer) error {
		peer := Peer{p}
		peer.Context().Cancel()

		// wait for message aggregator
		peer.messageAggregatorWaitGroup().Wait()
		close(peer.aggregatedMessageChannnel())

		// wait for message worker
		peer.messageWorkerWaitGroup().Wait()

		peerAddress, err := peer.GetAddress()
		if err == nil {
			handler.node.unregisterPeer(peerAddress)
		}
		return nil
	}
}
