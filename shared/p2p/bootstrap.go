package p2p

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
)

// StartBootstrapServer launches bootstrap node,
// which is an essential node in network to initialize Kademila DHT node discovery.
func StartBootstrapServer(ctx context.Context, nodekey *key.Key, addr multiaddr.Multiaddr) (info peerstore.PeerInfo, err error) {
	privKey, err := nodekey.DeriveLibp2pKeyPair()
	if err != nil {
		return
	}
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(addr),
	)
	if err != nil {
		return
	}
	dht, err := kaddht.New(ctx, h)
	if err != nil {
		return
	}
	if err = dht.Bootstrap(ctx); err != nil {
		return
	}
	info, err = getBootNodeInfo(h)
	return
}

// getBootInfo generates bootnode's providing address info and returns it
func getBootNodeInfo(h host.Host) (peerstore.PeerInfo, error) {
	iaddr, err := multiaddr.NewMultiaddr("/ipfs/" + h.ID().Pretty())
	if err != nil {
		return peerstore.PeerInfo{}, err
	}
	bootinfo, err := peerstore.InfoFromP2pAddr(h.Addrs()[1].Encapsulate(iaddr))
	return *bootinfo, err
}
