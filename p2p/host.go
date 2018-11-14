package p2p

import (
	"bufio"
	"context"
	"io"
	"log"

	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/golang/protobuf/proto"
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
	host host.Host
}

func (h *AirblocHost) ID() peer.ID {
	return h.host.ID()
}

func (h *AirblocHost) BootInfo() (peerstore.PeerInfo, error) {
	iaddr, err := multiaddr.NewMultiaddr("/ipfs/" + h.ID().Pretty())
	if err != nil {
		return peerstore.PeerInfo{}, err
	}
	bootinfo, err := peerstore.InfoFromP2pAddr(h.Addrs()[1].Encapsulate(iaddr))
	return *bootinfo, err
}

func (h *AirblocHost) Addrs() []multiaddr.Multiaddr {
	return h.host.Addrs()
}

func (h *AirblocHost) Peerstore() peerstore.Peerstore {
	return h.host.Peerstore()
}

func (h *AirblocHost) Mux() *multistream.MultistreamMuxer {
	return h.host.Mux()
}

func (h *AirblocHost) Network() net.Network {
	return h.host.Network()
}

func (h *AirblocHost) ConnManager() ifconnmgr.ConnManager {
	return h.host.ConnManager()
}

func (h AirblocHost) Connect(ctx context.Context, pi peerstore.PeerInfo) error {
	return h.host.Connect(ctx, pi)
}

func (h *AirblocHost) RegisterHandler(pid protocol.ID, handler func(p2p.Message)) {
	h.host.SetStreamHandler(pid, func(stream net.Stream) {
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

func (h *AirblocHost) UnregisterHandler(pid protocol.ID) {
	h.host.RemoveStreamHandler(pid)
}

func (h *AirblocHost) Send(ctx context.Context, message *p2p.Message, p peer.ID, pids ...protocol.ID) error {
	if err := h.host.Connect(ctx, peerstore.PeerInfo{ID: p}); err != nil {
		return errors.Wrap(err, "stream error : failed to connect peer")
	}

	stream, err := h.host.NewStream(ctx, p, pids...)
	if err != nil {
		return errors.Wrap(err, "stream error : failed to create stream")
	}
	defer stream.Close()

	message.From = []byte(h.host.ID())
	raw, err := proto.Marshal(message)
	if err != nil {
		return errors.Wrap(err, "stream error : failed to marshal proto message")
	}

	if _, err := stream.Write(raw); err != nil {
		return errors.Wrap(err, "stream error : failed to write data to stream")
	}
	return nil
}

func (h *AirblocHost) Publish(ctx context.Context, message *p2p.Message, pids ...protocol.ID) error {
	peerStore := h.Peerstore()
	for _, peerID := range peerStore.PeersWithAddrs() {
		err := h.Send(ctx, message, peerID, pids...)
		if err != nil {
			return errors.Wrap(err, "publish error : failed to publish message")
		}
	}
	return nil
}
