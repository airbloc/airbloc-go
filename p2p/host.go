package p2p

import (
	"context"

	"github.com/airbloc/airbloc-go/p2p/common"
	"github.com/libp2p/go-libp2p-interface-connmgr"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multistream"
)

type ProtocolRegistry interface {
	RegisterProtocol(common.Pid, ProtocolHandler, ...ProtocolAdapter)
	UnregisterProtocol(common.Pid)
}

type Sender interface {
	Send(context.Context, common.ProtoMessage, peer.ID, ...common.Pid) error
	Publish(context.Context, common.ProtoMessage, ...common.Pid) error
}

type Host interface {
	// network interfaces
	Mux() *multistream.MultistreamMuxer
	Network() net.Network
	ConnManager() ifconnmgr.ConnManager

	// peer information
	PeerInfo() peerstore.PeerInfo
	BootInfo() (peerstore.PeerInfo, error)
	Peerstore() peerstore.Peerstore

	// host interfaces
	Connect(context.Context, peerstore.PeerInfo) error
	ProtocolRegistry
	Sender
}
