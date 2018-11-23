package p2p

import (
	"context"

	"github.com/airbloc/airbloc-go/p2p/common"
)

type ProtocolAdapter func(ProtocolHandler) ProtocolHandler
type ProtocolHandler func(common.ProtoMessage)
type TopicHandler func(Server, context.Context, common.Message)

func (h ProtocolHandler) handle(sh ProtocolHandler, adapters ...ProtocolAdapter) ProtocolHandler {
	for i := len(adapters)/2 - 1; i >= 0; i-- {
		opp := len(adapters) - 1 - i
		adapters[i], adapters[opp] = adapters[opp], adapters[i]
	}
	for _, adapter := range adapters {
		sh = adapter(sh)
	}
	return sh
}

func (h ProtocolHandler) Handle(adapters ...ProtocolAdapter) ProtocolHandler {
	var sh ProtocolHandler = func(msg common.ProtoMessage) {
		h(msg)
	}
	return h.handle(sh, adapters...)
}
