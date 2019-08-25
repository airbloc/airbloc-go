package p2p

import (
	"context"
	"time"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/logger"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/klaytn/klaytn/common"
	peer "github.com/libp2p/go-libp2p-peer"
	"github.com/pkg/errors"
)

// Lookup querys address on P2P network and returns peer ID matching with the address.
func Lookup(ctx context.Context, server Server, addr common.Address, ackTimeout time.Duration) (peer.ID, error) {
	ctx, cancel := context.WithTimeout(ctx, ackTimeout)
	defer cancel()

	// 1. prepare for listening lookup ACKs
	waitForAck := make(chan peer.ID)
	ackHandler := func(_ Server, ctx context.Context, msg *IncomingMessage) {
		waitForAck <- msg.Sender
	}
	server.SubscribeTopic("/lookup/ack", &pb.LookupAck{}, ackHandler)
	defer server.UnsubscribeTopic("/lookup/ack")

	// 2. publish the lookup message
	if err := server.Publish(&pb.Lookup{}, "/lookup/"+addr.Hex()); err != nil {
		return "", errors.Wrap(err, "failed to publish lookup message")
	}

	// 3. wait for ACKs
	for {
		select {
		case id := <-waitForAck:
			// 4.1. return peer.ID matching with given address
			addrOfId, _ := addrFromID(id)
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

	addr, err := addrFromID(server.getHost().ID())
	if err != nil {
		return errors.Wrap(err, "invalid peer ID: is your libp2p.Identity SECP256k1?")
	}
	nodeAddr := addr.Hex()

	server.SubscribeTopic("/lookup/"+nodeAddr, &empty.Empty{}, func(_ Server, ctx context.Context, msg *IncomingMessage) {
		// send ACK (empty message)
		if err := server.Send(ctx, &pb.LookupAck{}, "/lookup/ack", msg.Sender); err != nil {
			log.Error("error: unable to send handshake reply: %s", err.Error(), msg.Sender.Loggable())
		}
	})
	return nil
}
