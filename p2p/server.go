package p2p

import (
	"context"
	"github.com/libp2p/go-libp2p-peer"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/node"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-peerstore"
)

type TopicRegistry interface {
	SubscribeTopic(string, proto.Message, TopicHandler) error
	UnsubscribeTopic(string) error
}

type PeerManager interface {
	Discovery()
	clearPeer()
	updatePeer()
}

type Server interface {
	// api backend
	node.Service

	// server interfaces
	PeerManager
	TopicRegistry

	// sender interfaces
	Send(context.Context, proto.Message, string, peer.ID) error
	Publish(context.Context, proto.Message, string) error

	// database interfaces
	LocalDB() localdb.Database
	MetaDB() metadb.Database

	// for test
	setContext(context.Context)
	getHost() Host
	bootInfo() (peerstore.PeerInfo, error)
}
