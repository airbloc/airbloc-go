package p2p

import (
	"reflect"

	"github.com/airbloc/logger"
	"github.com/perlin-network/noise"
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
	nodeEventHandler{node}.register(node.node)
}

func (handler nodeEventHandler) register(node *noise.Node) {
	node.OnListenerError(handler.onListenerErrorHandler())
	node.OnPeerInit(handler.onPeerInitHandler())
	node.OnPeerDisconnected(handler.onPeerDisconnectedHandler())
}

func (handler nodeEventHandler) onListenerErrorHandler() noise.OnErrorCallback {
	return func(_ *noise.Node, err error) error {
		node := handler.node
		node.log.Error("Listener error occurred", logger.Attrs{
			"type":    reflect.TypeOf(err).String(),
			"message": err.Error(),
		})
		return err
	}
}

func (handler nodeEventHandler) onPeerInitHandler() noise.OnPeerInitCallback {
	return func(node *noise.Node, peer *noise.Peer) error {
		registerPeerEventHandler(handler.node, peer)

		//aggregatedMessages := make(chan noise.Message)
		//for _, opcode := range message.Opcodes {
		//
		//}

		return nil
	}
}

func (handler nodeEventHandler) onPeerDisconnectedHandler() noise.OnPeerDisconnectCallback {
	return func(node *noise.Node, peer *noise.Peer) error {
		return nil
	}
}
