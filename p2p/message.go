package p2p

import (
	"context"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-peerstore"
)

type Message struct {
	ctx  context.Context
	Data proto.Message
	Info peerstore.PeerInfo
}

func msgType(msg proto.Message) reflect.Type {
	return reflect.ValueOf(msg).Elem().Type()
}
