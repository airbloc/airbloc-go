package p2p

import (
	"time"

	"github.com/airbloc/airbloc-go/network/p2p/handshake/identity"
	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise"
)

func GetPeerAddressAsync(peer *noise.Peer) <-chan common.Address {
	peerAddressChan := make(chan common.Address, 1)
	go func() {
		for peer.Get(identity.KeyAddress) == nil {
			time.Sleep(100 * time.Millisecond)
		}
		peerAddressChan <- peer.Get(identity.KeyAddress).(common.Address)
	}()
	return peerAddressChan
}

func GetPeerAddress(peer *noise.Peer) common.Address {
	return <-GetPeerAddressAsync(peer)
}
