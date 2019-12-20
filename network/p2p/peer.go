package p2p

import (
	"reflect"
	"sync"
	"time"

	"github.com/airbloc/logger"

	"github.com/airbloc/airbloc-go/network/p2p/message"

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

type aggregatedMessage struct {
	message noise.Message
	opcode  noise.Opcode
}

func (peer Peer) context() Context {
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

func RunMessageAggregator(node Node, peer Peer) {
	var (
		aggregatorWaitGroup                            = peer.messageAggregatorWaitGroup()
		aggregatedMessageChan chan<- aggregatedMessage = peer.aggregatedMessageChannnel()
	)

	aggregatorWaitGroup.Add(len(message.Opcodes))
	for _, opcode := range message.Opcodes {
		go func(opcode noise.Opcode) {
			defer aggregatorWaitGroup.Done()
			for {
				select {
				case <-peer.context().Done():
					return
				case receivedMessage := <-peer.Receive(opcode):
					aggregatedMessageChan <- struct {
						message noise.Message
						opcode  noise.Opcode
					}{
						message: receivedMessage,
						opcode:  opcode,
					}
				}
			}
		}(opcode)
	}
}

func RunMessageWorker(node Node, peer Peer) {
	var (
		workerWaitgroup                                = peer.messageWorkerWaitGroup()
		aggregatedMessageChan <-chan aggregatedMessage = peer.aggregatedMessageChannnel()
	)

	workerWaitgroup.Add(1)
	go func() {
		defer workerWaitgroup.Done()

		for {
			select {
			case <-peer.context().Done():
				return
			case receivedMessage := <-aggregatedMessageChan:
				node.logger().Debug("Message received", logger.Attrs{
					"opcode": uint8(receivedMessage.opcode),
					"type":   reflect.TypeOf(receivedMessage.message).String(),
				})
				handler, err := node.GetHandler(receivedMessage.opcode)
				if err != nil {
					node.logger().Error("Cannot get handler", logger.Attrs{
						"error": logger.Attrs{
							"type":    reflect.TypeOf(errors.Cause(err)).String(),
							"message": err.Error(),
						},
						"message": reflect.TypeOf(receivedMessage.message).String(),
					})
					continue
				}

				err = handler(node.context(), receivedMessage.message, peer.Peer)
				if err != nil {
					node.logger().Error("Handler error occurred", logger.Attrs{
						"error": logger.Attrs{
							"type":    reflect.TypeOf(errors.Cause(err)).String(),
							"message": err.Error(),
						},
						"message": reflect.TypeOf(receivedMessage.message).String(),
					})
					continue
				}
			}
		}
	}()
}
