package p2p

import (
	"context"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/logger"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"sync"
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
	iaddr, err := multiaddr.NewMultiaddr(h.Addrs()[1].String() + "/p2p/" + h.ID().Pretty())
	if err != nil {
		return peerstore.PeerInfo{}, err
	}
	bootinfo, err := peerstore.InfoFromP2pAddr(iaddr)
	return *bootinfo, err
}

// This code is borrowed from the go-ipfs bootstrap process
func bootstrapConnect(ctx context.Context, ph host.Host, peers []peerstore.PeerInfo) error {
	log := logger.New("bootstrapper")
	errs := make(chan error, len(peers))
	var wg sync.WaitGroup
	for _, p := range peers {

		// performed asynchronously because when performed synchronously, if
		// one `Connect` call hangs, subsequent calls are more likely to
		// fail/abort due to an expiring context.
		// Also, performed asynchronously for dial speed.

		wg.Add(1)
		go func(p peerstore.PeerInfo) {
			defer wg.Done()
			log.Info("{} bootstrapping to {}", ph.ID(), p.ID)

			// ph.Peerstore().AddAddrs(p.ID, p.Addrs, peerstore.PermanentAddrTTL)
			if err := ph.Connect(ctx, p); err != nil {
				log.Error("failed to bootstrap with {}", err, p.ID)
				errs <- err
				return
			}
			// log.Debug("bootstrapped with {}", p.ID)
		}(p)
	}
	wg.Wait()

	// our failure condition is when no connection attempt succeeded.
	// So drain the errs channel, counting the results.
	close(errs)
	count := 0
	var err error
	for err = range errs {
		if err != nil {
			count++
		}
	}
	if count == len(peers) {
		return errors.Wrap(err, "failed to bootstrap")
	}
	return nil
}
