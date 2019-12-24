package p2p

import (
	"reflect"

	"github.com/airbloc/airbloc-go/network/p2p/message"
	"github.com/airbloc/logger"

	"github.com/perlin-network/noise"
	"github.com/pkg/errors"
)

type aggregatedMessage struct {
	message message.Message
	opcode  noise.Opcode
}

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
				case <-peer.Context().Done():
					msg, _ := noise.MessageFromOpcode(opcode) // must not be nil
					node.logger().Debug(
						"Message aggregator for message {} stopped",
						reflect.TypeOf(msg).String(),
						logger.Attrs{"peer-address": (<-peer.GetAddressAsync()).Hex()})
					return
				case receivedMessage := <-peer.Receive(opcode):
					if _, ok := receivedMessage.(message.Message); ok {
						aggregatedMessageChan <- struct {
							message message.Message
							opcode  noise.Opcode
						}{
							message: receivedMessage,
							opcode:  opcode,
						}
					}
				}
			}
		}(opcode)
	}
}

func MessageHandler(node Node, peer Peer) {
	var (
		workerWaitgroup                                = peer.messageWorkerWaitGroup()
		aggregatedMessageChan <-chan aggregatedMessage = peer.aggregatedMessageChannnel()
	)

	workerWaitgroup.Add(1)
	defer workerWaitgroup.Done()

	for {
		select {
		case <-peer.Context().Done():
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

			err := node.handle(receivedMessage.opcode, receivedMessage.message, peer.Peer)
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
