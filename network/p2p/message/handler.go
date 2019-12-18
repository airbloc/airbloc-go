package message

import (
	"context"

	"github.com/perlin-network/noise"
)

type HandlerFunc func(context context.Context, message noise.Message, peer *noise.Peer) error
