package p2p

import (
	"context"

	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-protocol"
)

type TopicRegistry interface {
	RegisterTopic(string, proto.Message, TopicHandler) error
	UnregisterTopic(string) error
}

type ProtocolRegistry interface {
	RegisterProtocol(protocol.ID, ProtocolHandler)
	UnregisterProtocol(protocol.ID)
}

type Sender interface {
	Send(context.Context, p2p.Message, peer.ID) error
	Publish(context.Context, p2p.Message) error
}
