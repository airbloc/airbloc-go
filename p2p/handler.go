package p2p

import (
	"context"

	"github.com/airbloc/airbloc-go/p2p/common"
	"github.com/libp2p/go-libp2p-net"
)

type StreamAdapter func(StreamHandler) StreamHandler
type StreamHandler func(net.Stream)
type ProtocolHandler func(common.ProtoMessage)
type TopicHandler func(Server, context.Context, common.Message)

func (h StreamHandler) handle(sh StreamHandler, adapters ...StreamAdapter) StreamHandler {
	for i := len(adapters)/2 - 1; i >= 0; i-- {
		opp := len(adapters) - 1 - i
		adapters[i], adapters[opp] = adapters[opp], adapters[i]
	}
	for _, adapter := range adapters {
		sh = adapter(sh)
	}
	return sh
}

func (h StreamHandler) Handle(adapters ...StreamAdapter) StreamHandler {
	var sh StreamHandler = func(stream net.Stream) {
		h(stream)
	}
	return h.handle(sh, adapters...)
}
