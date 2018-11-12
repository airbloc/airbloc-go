package p2p

import (
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-peer"
)

type Handler func(peer.ID, proto.Message)
