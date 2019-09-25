package p2p

import (
	"context"
	"time"

	"github.com/airbloc/logger"
	host "github.com/libp2p/go-libp2p-host"
	store "github.com/libp2p/go-libp2p-peerstore"
	mdns "github.com/libp2p/go-libp2p/p2p/discovery"
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

// HandlePeerFound registers the peer with the getHost.
func (d *discovery) HandlePeerFound(pi store.PeerInfo) {
	d.host.Peerstore().AddAddrs(pi.ID, pi.Addrs, store.PermanentAddrTTL)
	if err := d.host.Connect(d.ctx, pi); err != nil {
		d.log.Error("Warning: Failed to connect to new peer {id}", err, logger.Attrs{
			"id": pi.ID.String(),
		})
	}
	d.log.Debug("Peers are now {peerCount}", logger.Attrs{
		"peers": d.host.Peerstore().Peers(),
	})
}
