package p2p

import (
	"context"
	"github.com/libp2p/go-libp2p-protocol"
	"log"

	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/pkg/errors"
)

type MessageHandler func(RawMessage)

type PubSubHost struct {
	host.Host
	pid     protocol.ID
	handler MessageHandler
}

func WrapPubSubHost(host host.Host, pid Pid, handler MessageHandler) *PubSubHost {
	h := &PubSubHost{
		Host:    host,
		pid:     pid.ProtocolID(),
		handler: handler,
	}
	host.SetStreamHandler(h.pid, h.handleIncoming)
	return h
}

func (h *PubSubHost) handleIncoming(stream net.Stream) {
	defer stream.Reset()
	msg, err := ReadRawMessage(stream)
	if err != nil {
		log.Println("failed to read message from stream :", err)
		return
	}
	h.handler(msg)
}

func (h *PubSubHost) Send(ctx context.Context, msg RawMessage, id peer.ID) error {
	err := h.Connect(ctx, peerstore.PeerInfo{ID: id})
	if err != nil {
		return errors.Wrapf(err, "unable to connect to %s", id.String())
	}

	stream, err := h.NewStream(ctx, id, h.pid)
	if err != nil {
		return errors.Wrap(err, "failed to open stream")
	}
	defer stream.Close()

	if err := msg.WriteTo(stream); err != nil {
		return errors.Wrap(err, "failed to send message to stream")
	}
	return nil
}
