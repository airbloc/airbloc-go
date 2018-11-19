package ablhost

import (
	"context"

	"github.com/airbloc/airbloc-go/p2p"
	"github.com/airbloc/airbloc-go/p2p/common"
	"github.com/multiformats/go-multistream"
	"github.com/pkg/errors"
)

type Host struct {
	p2p.Host
	// when publishing message, limiting count of peers
	limit int
}

func NewHost(host p2p.Host, limit int) p2p.Host {
	return &Host{
		Host:  host,
		limit: limit,
	}
}

func (h *Host) Publish(ctx context.Context, msg common.ProtoMessage, pids ...common.Pid) error {
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
