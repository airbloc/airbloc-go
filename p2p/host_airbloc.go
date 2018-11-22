package p2p

import (
	"context"

	"github.com/airbloc/airbloc-go/p2p/common"
	"github.com/multiformats/go-multistream"
	"github.com/pkg/errors"
)

type AirblocHost struct {
	Host
	// when publishing message, limiting count of peers
	limit int
}

func NewAirblocHost(host Host, limit int) Host {
	return &AirblocHost{
		Host:  host,
		limit: limit,
	}
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
