package p2p

import (
	"context"
	"github.com/airbloc/airbloc-go/shared/p2p/common"
	"github.com/libp2p/go-libp2p-protocol"
	"github.com/multiformats/go-multiaddr"

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

type Host interface {
	// ID returns the (local) peer.ID associated with this Host
	ID() peer.ID

	// network interfaces
	Mux() *multistream.MultistreamMuxer
	Network() net.Network
	ConnManager() ifconnmgr.ConnManager

	// peer information
	PeerInfo() peerstore.PeerInfo
	Peerstore() peerstore.Peerstore

	// sender
	Send(context.Context, common.ProtoMessage, peer.ID, ...common.Pid) error
	Publish(context.Context, common.ProtoMessage, ...common.Pid) error

	// host interfaces
	Connect(context.Context, peerstore.PeerInfo) error
	ProtocolRegistry

	Addrs() []multiaddr.Multiaddr
	SetStreamHandler(pid protocol.ID, handler net.StreamHandler)
	SetStreamHandlerMatch(protocol.ID, func(string) bool, net.StreamHandler)
	RemoveStreamHandler(pid protocol.ID)
	NewStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (net.Stream, error)

	Close() error
}
