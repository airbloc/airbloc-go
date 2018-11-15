package p2p

import (
	"bufio"
	"context"
	"io"
	"log"

	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-interface-connmgr"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-protocol"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multistream"
	"github.com/pkg/errors"
)

type AirblocHost struct {
	pid  Pid
	host host.Host
	// when publish message, limiting count of peers
	limit int
}

// PID generates external pid string and returns it
func (h *AirblocHost) PID() protocol.ID {
	return h.pid.ProtocolID()
}

// PeerInfo generates peerstore.PeerInfo object and returns it
func (h *AirblocHost) PeerInfo() peerstore.PeerInfo {
	return peerstore.PeerInfo{
		ID:    h.host.ID(),
		Addrs: h.host.Addrs(),
	}
}

// BootInfo generates bootnode's providing address info and returns it
func (h *AirblocHost) BootInfo() (peerstore.PeerInfo, error) {
	info := h.PeerInfo()
	iaddr, err := multiaddr.NewMultiaddr("/ipfs/" + info.ID.Pretty())
	if err != nil {
		return peerstore.PeerInfo{}, err
	}
	bootinfo, err := peerstore.InfoFromP2pAddr(info.Addrs[1].Encapsulate(iaddr))
	return *bootinfo, err
}

// Peerstore returns host's peerstore
func (h *AirblocHost) Peerstore() peerstore.Peerstore {
	return h.host.Peerstore()
}

// Mux returns host's multistreamMuxer
func (h *AirblocHost) Mux() *multistream.MultistreamMuxer {
	return h.host.Mux()
}

// Network returns host's network interface
func (h *AirblocHost) Network() net.Network {
	return h.host.Network()
}

// ConnManager returns host's connection manager interface
func (h *AirblocHost) ConnManager() ifconnmgr.ConnManager {
	return h.host.ConnManager()
}

// Connect makes connect with other peer by peerstore.PeerInfo
func (h AirblocHost) Connect(ctx context.Context, pi peerstore.PeerInfo) error {
	return h.host.Connect(ctx, pi)
}

// Register register p2p.Message handler
func (h *AirblocHost) RegisterProtocol(handler ProtocolHandler) {
	h.host.SetStreamHandler(h.PID(), func(stream net.Stream) {
		defer stream.Reset()

		var raw []byte
		reader := bufio.NewReader(stream)
		p := make([]byte, 4)
		for {
			n, err := reader.Read(p)
			if err == io.EOF {
				break
			}
			raw = append(raw, p[:n]...)
		}

		msg := p2p.Message{}
		if err := proto.Unmarshal(raw, &msg); err != nil {
			log.Println("marshal error :", err)
			return
		}

		go handler(msg)
	})
}

// Unregister unregister handler
func (h *AirblocHost) UnregisterProtocol() {
	h.host.RemoveStreamHandler(h.PID())
}

func (h *AirblocHost) Send(ctx context.Context, message p2p.Message, p peer.ID) error {
	err := h.host.Connect(ctx, peerstore.PeerInfo{ID: p})
	if err == peerstore.ErrNotFound {

	}
	if err != nil {
		return errors.Wrap(err, "stream error : failed to connect peer")
	}

	stream, err := h.host.NewStream(ctx, p, h.PID())
	if err != nil {
		return errors.Wrap(err, "stream error : failed to create stream")
	}
	defer stream.Close()

	message.From = []byte(h.host.ID())
	raw, err := proto.Marshal(&message)
	if err != nil {
		return errors.Wrap(err, "stream error : failed to marshal proto message")
	}

	if _, err := stream.Write(raw); err != nil {
		return errors.Wrap(err, "stream error : failed to write data to stream")
	}
	return nil
}

func (h *AirblocHost) Publish(ctx context.Context, message p2p.Message) error {
	peerStore := h.Peerstore()

	limit := h.limit
	for _, peerID := range peerStore.PeersWithAddrs() {
		if limit <= 0 {
			break
		}

		err := h.Send(ctx, message, peerID)
		if errors.Cause(err) == multistream.ErrNotSupported {
			log.Println("protocol not supported")
			continue
		}

		if err != nil {
			return errors.Wrap(err, "publish error : failed to publish message")
		}
		limit--
	}
	return nil
}
