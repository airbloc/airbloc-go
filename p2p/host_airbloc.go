package p2p

import (
	"context"
	"log"

	"github.com/airbloc/airbloc-go/p2p/common"
	"github.com/libp2p/go-libp2p-net"
	"github.com/multiformats/go-multistream"
	"github.com/pkg/errors"
)

type AirblocHost struct {
	BasicHost
	// when publishing message, limiting count of peers
	limit int
}

// Register register p2p.Message handler
func (h *AirblocHost) RegisterProtocol(pid common.Pid, handler ProtocolHandler) {
	h.BasicHost.RegisterProtocol(pid, func(stream net.Stream) {
		defer stream.Reset()
		msg, err := common.ReadMessage(stream)
		if err != nil {
			log.Println("failed to read message from stream :", err)
			return
		}
		go handler(msg)
	})
}

// Unregister unregister handler
func (h *AirblocHost) UnregisterProtocol(pid common.Pid) {
	h.BasicHost.UnregisterProtocol(pid)
}

func (h *AirblocHost) Publish(ctx context.Context, msg common.ProtoMessage, pids ...common.Pid) error {
	limit := h.limit
	for _, peerID := range h.Peerstore().PeersWithAddrs() {
		if limit <= 0 {
			break
		}

		err := h.Send(ctx, msg, peerID, pids...)
		if errors.Cause(err) == multistream.ErrNotSupported {
			continue
		}

		if err != nil {
			return errors.Wrap(err, "publish error : failed to publish message")
		}
		limit--
	}
	return nil
}
