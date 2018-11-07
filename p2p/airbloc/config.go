package airbloc

import (
	"github.com/airbloc/airbloc-go/key"
	"github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
)

type DHTClientConfig struct {
	BootstrapAddrs []ma.Multiaddr
	MaxProvider    int
}

type DHTBootstrapConfig struct {
	dht.BootstrapConfig
}

type NodeConfig struct {
	Addr     ma.Multiaddr
	Identity *key.Key

	DHTClient    *DHTClientConfig
	DHTBootstrap *DHTBootstrapConfig
}
