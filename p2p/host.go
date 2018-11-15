package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p-interface-connmgr"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-protocol"
	"github.com/multiformats/go-multistream"
)

type Host interface {
	// network interfaces
	Mux() *multistream.MultistreamMuxer
	Network() net.Network
	ConnManager() ifconnmgr.ConnManager

	// peer information
	PID() protocol.ID
	PeerInfo() peerstore.PeerInfo
	BootInfo() (peerstore.PeerInfo, error)
	Peerstore() peerstore.Peerstore

	// host interfaces
	Connect(context.Context, peerstore.PeerInfo) error
	ProtocolRegistry
	Sender
}
