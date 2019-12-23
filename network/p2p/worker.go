package p2p

import (
	"reflect"

	"github.com/airbloc/airbloc-go/network/p2p/message"
	"github.com/airbloc/logger"

	"github.com/perlin-network/noise"
	"github.com/pkg/errors"
)

func MessageAggregator(node Node, peer Peer) {
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
					msg, _ := noise.MessageFromOpcode(opcode) // must not be nil
					node.logger().Debug(
						"Message aggregator for message {} stopped",
						reflect.TypeOf(msg).String(),
						logger.Attrs{"peer-address": (<-peer.GetAddressAsync()).Hex()})
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

func MessageHandler(node Node, peer Peer) func() {
	var (
		workerWaitgroup                                = peer.messageWorkerWaitGroup()
		aggregatedMessageChan <-chan aggregatedMessage = peer.aggregatedMessageChannnel()
	)

	workerWaitgroup.Add(1)
	return func() {
		defer workerWaitgroup.Done()

		for {
			select {
			case <-peer.context().Done():
				node.logger().Debug("Message handle worker stopped")
				return
			case receivedMessage := <-aggregatedMessageChan:
				if receivedMessage.message == nil {
					node.logger().Info("Nil message {}", receivedMessage)
					continue
				}
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
	}
}
