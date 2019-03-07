package p2p

import (
	"context"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-peerstore"
	mdns "github.com/libp2p/go-libp2p/p2p/discovery"
	"time"
)

// Discovery interval for multicast DNS querying.
var discoveryInterval = 1 * time.Minute

// mDNSTag is the name of the mDNS service.
var mDNSTag = mdns.ServiceTag

// startmDNSDiscovery supports discovery via multicast DNS peer discovery.
func startmDNSDiscovery(ctx context.Context, h host.Host) error {
	mdnsService, err := mdns.NewMdnsService(ctx, h, discoveryInterval, mDNSTag)
	if err != nil {
		return err
	}

	mdnsService.RegisterNotifee(&discovery{
		ctx:  ctx,
		host: h,
		log:  logger.New("p2p-mdns"),
	})
	return nil
}

// Discovery implements mDNS notifee interface.
type discovery struct {
	ctx  context.Context
	host host.Host
	log  *logger.Logger
}

// HandlePeerFound registers the peer with the host.
func (d *discovery) HandlePeerFound(pi peerstore.PeerInfo) {
	d.host.Peerstore().AddAddrs(pi.ID, pi.Addrs, peerstore.PermanentAddrTTL)
	if err := d.host.Connect(d.ctx, pi); err != nil {
		d.log.Error("Warning: Failed to connect to new peer {id}", err, logger.Attrs{
			"id": pi.ID.String(),
		})
	}
	log.Debug("Peers are now {peerCount}", logger.Attrs{
		"peers": d.host.Peerstore().Peers(),
	})
}
