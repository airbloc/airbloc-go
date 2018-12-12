package p2p

import (
	"context"
	"time"

	p2pcommon "github.com/airbloc/airbloc-go/p2p/common"
	"github.com/azer/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/pkg/errors"
)

// Lookup querys address on P2P network and returns peer ID matching with the address.
func Lookup(ctx context.Context, server Server, addr common.Address, ackTimeout time.Duration) (peer.ID, error) {
	ctx, cancel := context.WithTimeout(ctx, ackTimeout)
	defer cancel()

	topicName := "/lookup/" + addr.String()

	// 1. publish the lookup message
	if err := server.Publish(ctx, &empty.Empty{}, topicName); err != nil {
		return "", errors.Wrap(err, "failed to publish lookup message")
	}

	// 2. listen to lookup ACKs
	waitForAck := make(chan peer.ID)
	ackHandler := func(_ Server, ctx context.Context, msg p2pcommon.Message) {
		waitForAck <- msg.SenderInfo.ID
	}
	if err := server.SubscribeTopic(topicName+"/ack", &empty.Empty{}, ackHandler); err != nil {
		return "", errors.Wrap(err, "failed to listen to lookup ACKs")
	}

	// 3. wait for ACKs
	for {
		select {
		case id := <-waitForAck:
			// 4.1. return peer.ID matching with given address
			addrOfId, _ := p2pcommon.AddrFromID(id)
			if addrOfId == addr {
				return id, nil
			}
		case <-ctx.Done():
			// 4.2. can't find any matching peer in given timeout
			return "", ErrAddressNotFound
		}
	}
}

// StartNameServer subscribes lookup topic and sends ACK when lookup is requested.
func StartNameServer(server Server) error {
	log := logger.New("nameserver")

	addr, err := p2pcommon.AddrFromID(server.getHost().ID())
	if err != nil {
		return errors.Wrap(err, "invalid peer ID")
	}
	topicName := "/lookup/" + addr.String()

	return server.SubscribeTopic("/lookup/"+addr.String(), &empty.Empty{}, func(_ Server, ctx context.Context, msg p2pcommon.Message) {
		// send ACK (empty message)
		if err := server.Send(ctx, &empty.Empty{}, topicName+"/ack", msg.SenderInfo.ID); err != nil {
			log.Error("error: unable to send handshake reply: %s", err.Error(), msg.SenderInfo.ID.Loggable())
		}
	})
}
