package airbloc

import (
	"context"
	"log"

	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	"github.com/ipfs/go-cid"
	"github.com/jbenet/goprocess"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p/p2p/host/routed"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
)

type Node struct {
	host host.Host
	dht  *dht.IpfsDHT

	peerch map[uint64]<-chan peerstore.PeerInfo
	proc   goprocess.Process
	cfg    NodeConfig
}

// NewNode configure node's specific functions and returns it
// TODO: replace key input to account pointer (if account package finished)
func NewNode(
	ctx context.Context,
	cfg NodeConfig,
) (p2p.Node, error) {
	node, err := newNode(ctx, cfg.Addr, cfg.Identity)
	if err != nil {
		return nil, err
	}

	if cfg.DHTBootstrap != nil {
		bootCfg := cfg.DHTBootstrap

		node.proc, err = node.dht.BootstrapWithConfig(bootCfg.BootstrapConfig)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to enable Bootstrap mode")
		}
	}

	if cfg.DHTClient != nil {
		clientCfg := cfg.DHTClient

		for _, bootstrapAddr := range clientCfg.BootstrapAddrs {
			peerinfo, err := peerstore.InfoFromP2pAddr(bootstrapAddr)
			if err != nil {
				return nil, errors.Wrap(err, "Failed to get peerinfo")
			}
			if err := node.host.Connect(ctx, *peerinfo); err != nil {
				return nil, errors.Wrap(err, "Failed to connect bootstrap node")
			} else {
				log.Println("Connect successfully", peerinfo.ID.Pretty())
			}
		}
	}

	return node, err
}

// newNode makes libp2p & kad-dht host and pack into Node struct and returns it.
func newNode(
	ctx context.Context,
	addr ma.Multiaddr,
	identity *key.Key,
) (*Node, error) {
	privKey, err := identity.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to Derive Libp2p KeyPair")
	}

	ablHost, err := libp2p.New(ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(addr))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get Libp2p host")
	}

	ablDHT, err := dht.New(ctx, ablHost)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get Kad-DHT host")
	}

	return &Node{
		host:   ablHost,
		dht:    ablDHT,
		peerch: make(map[uint64]<-chan peerstore.PeerInfo, len(cids)),
	}, nil
}

func (node *Node) Start() error {
	ctx := context.Background()

	clientCfg := node.cfg.DHTClient
	for _, code := range cids {
		v1b := cid.V1Builder{
			Codec:  code,
			MhType: multihash.KECCAK_256,
		}
		id, err := v1b.Sum([]byte(""))
		if err != nil {
			return err
		}

		err = node.dht.Provide(ctx, id, true)
		if err != nil {
			return err
		}

		node.peerch[code] = node.dht.FindProvidersAsync(
			context.Background(),
			id, clientCfg.MaxProvider,
		)
	}

	routedhost.Wrap()

	return nil
}
func (node *Node) Send()    {}
func (node *Node) Receive() {}
func (node *Node) Stop() {
	node.proc.Close()
	node.dht.Close()
	node.host.Close()
}
