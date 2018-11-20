package p2p

import (
	"context"

	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-peerstore"
)

type TopicRegistry interface {
	RegisterTopic(string, proto.Message, TopicHandler) error
	UnregisterTopic(string) error
}

type PeerManager interface {
	Discovery()
	clearPeer()
	updatePeer()
}

type Server interface {
	// api backend
	api.Service

	// server interfaces
	PeerManager
	TopicRegistry

	// host interfaces
	Sender
	ProtocolRegistry

	// database interfaces
	LocalDB() localdb.Database
	MetaDB() metadb.Database

	// for test
	setContext(context.Context)
	getHost() Host
	bootInfo() (peerstore.PeerInfo, error)
}
