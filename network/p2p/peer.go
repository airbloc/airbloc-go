package p2p

import (
	"sync"
	"time"

	"github.com/airbloc/airbloc-go/network/p2p/handshake/identity"

	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise"
	"github.com/pkg/errors"
)

const (
	KeyPeerContext                    = "abl.ctx"
	KeyPeerAggregatedMessageChannel   = "abl.ch.msg.aggr"
	KeyPeerWaitGroupMessageWorker     = "abl.wg.msg.wkr"
	KeyPeerWaitGroupMessageAggregator = "abl.wg.msg.aggr"
)

type Peer struct {
	*noise.Peer
}

func (peer Peer) Context() Context {
	return peer.Get(KeyPeerContext).(Context)
}

func (peer Peer) GetAddressAsync() <-chan common.Address {
	peerAddressChan := make(chan common.Address, 1)
	go func() {
		timeout := time.After(2 * time.Second)
		for {
			select {
			case <-timeout:
				peerAddressChan <- common.Address{}
				return
			default:
				peerAddress := peer.Get(identity.KeyAddress)
				if peerAddress != nil {
					peerAddressChan <- peerAddress.(common.Address)
					return
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	return peerAddressChan
}

func (peer Peer) GetAddress() (common.Address, error) {
	peerAddress := <-peer.GetAddressAsync()
	if peerAddress == (common.Address{}) {
		return common.Address{}, errors.New("cannot get peer's address")
	}
	return peerAddress, nil
}

func (peer Peer) aggregatedMessageChannnel() chan aggregatedMessage {
	return peer.Get(KeyPeerAggregatedMessageChannel).(chan aggregatedMessage)
}

func (peer Peer) messageWorkerWaitGroup() *sync.WaitGroup {
	return peer.Get(KeyPeerWaitGroupMessageWorker).(*sync.WaitGroup)
}

func (peer Peer) messageAggregatorWaitGroup() *sync.WaitGroup {
	return peer.Get(KeyPeerWaitGroupMessageAggregator).(*sync.WaitGroup)
}
