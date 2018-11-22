package p2p

import (
	"context"
	"log"

	"github.com/airbloc/airbloc-go/p2p/common"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-interface-connmgr"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multistream"
	"github.com/pkg/errors"
)

type BasicHost struct {
	host host.Host
}

func NewBasicHost(host host.Host) Host {
	return &BasicHost{
		host: host,
	}
}

// Mux returns host's multistreamMuxer
func (h *BasicHost) Mux() *multistream.MultistreamMuxer {
	return h.host.Mux()
}

// Network returns host's network interface
func (h *BasicHost) Network() net.Network {
	return h.host.Network()
}

// ConnManager returns host's connection manager interface
func (h *BasicHost) ConnManager() ifconnmgr.ConnManager {
	return h.host.ConnManager()
}

// PeerInfo generates peerstore.PeerInfo object and returns it
func (h *BasicHost) PeerInfo() peerstore.PeerInfo {
	return peerstore.PeerInfo{ID: h.host.ID(), Addrs: h.host.Addrs()}
}

// BootInfo generates bootnode's providing address info and returns it
func (h *BasicHost) BootInfo() (peerstore.PeerInfo, error) {
	info := h.PeerInfo()
	iaddr, err := multiaddr.NewMultiaddr("/ipfs/" + info.ID.Pretty())
	if err != nil {
		return peerstore.PeerInfo{}, err
	}
	bootinfo, err := peerstore.InfoFromP2pAddr(info.Addrs[1].Encapsulate(iaddr))
	return *bootinfo, err
}

// Peerstore returns host's peerstore
func (h *BasicHost) Peerstore() peerstore.Peerstore {
	return h.host.Peerstore()
}

// Connect makes connect with other peer by peerstore.PeerInfo
func (h *BasicHost) Connect(ctx context.Context, pi peerstore.PeerInfo) error {
	return h.host.Connect(ctx, pi)
}

// Protocol Registry
// Register register p2p.Message handler
func (h *BasicHost) RegisterProtocol(
	pid common.Pid,
	handler ProtocolHandler,
	adapters ...ProtocolAdapter,
) {
	h.host.SetStreamHandler(pid.ProtocolID(), func(stream net.Stream) {
		defer stream.Reset()
		msg, err := common.ReadMessage(stream)
		if err != nil {
			log.Println("failed to read message from stream :", err)
			return
		}
		go handler.Handle(adapters...)(msg)
	})
}

// Unregister unregister handler
func (h *BasicHost) UnregisterProtocol(pid common.Pid) {
	h.host.RemoveStreamHandler(pid.ProtocolID())
}

func (h *BasicHost) Send(ctx context.Context, msg common.ProtoMessage, id peer.ID, pids ...common.Pid) error {
	stream, err := h.host.NewStream(ctx, id, common.Pids(pids).ProtocolID()...)
	if err != nil {
		return errors.Wrap(err, "stream error : failed to create stream")
	}
	defer stream.Close()

	msg.From = []byte(h.host.ID())
	msg.Protocol = []byte(stream.Protocol())
	if err := msg.WriteMessage(stream); err != nil {
		return errors.Wrap(err, "stream error : failed to send message to stream")
	}
	return nil
}

func (h *BasicHost) Publish(ctx context.Context, msg common.ProtoMessage, pids ...common.Pid) error {
	for _, peerID := range h.Peerstore().PeersWithAddrs() {
		err := h.Send(ctx, msg, peerID, pids...)
		if errors.Cause(err) == multistream.ErrNotSupported {
			continue
		}

		if err != nil {
			return errors.Wrap(err, "publish error : failed to publish message")
		}
	}
	return nil
}
