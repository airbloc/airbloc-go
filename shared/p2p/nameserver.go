package p2p

import (
	"context"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"time"

	p2pCommon "github.com/airbloc/airbloc-go/shared/p2p/common"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/pkg/errors"
)

// Lookup querys address on P2P network and returns peer ID matching with the address.
func Lookup(ctx context.Context, server Server, addr common.Address, ackTimeout time.Duration) (peer.ID, error) {
	ctx, cancel := context.WithTimeout(ctx, ackTimeout)
	defer cancel()

	// 1. prepare for listening lookup ACKs
	waitForAck := make(chan peer.ID)
	ackHandler := func(_ Server, ctx context.Context, msg p2pCommon.Message) {
		waitForAck <- msg.SenderInfo.ID
	}
	server.SubscribeTopic("/lookup/ack", &empty.Empty{}, ackHandler)
	defer server.UnsubscribeTopic("/lookup/ack")

	// 2. publish the lookup message
	if err := server.Publish(ctx, &empty.Empty{}, "/lookup"); err != nil {
		return "", errors.Wrap(err, "failed to publish lookup message")
	}

	// 3. wait for ACKs
	for {
		select {
		case id := <-waitForAck:
			// 4.1. return peer.ID matching with given address
			addrOfId, _ := p2pCommon.AddrFromID(id)
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

	addr, err := p2pCommon.AddrFromID(server.getHost().ID())
	if err != nil {
		return errors.Wrap(err, "invalid peer ID")
	}
	nodeAddr := addr.Hex()

	server.SubscribeTopic("/lookup", &pb.Lookup{}, func(_ Server, ctx context.Context, msg p2pCommon.Message) {
		req, _ := msg.Data.(*pb.Lookup)
		if req.Address == nodeAddr {
			// send ACK (empty message)
			if err := server.Send(ctx, &pb.LookupAck{}, "/lookup/ack", msg.SenderInfo.ID); err != nil {
				log.Error("error: unable to send handshake reply: %s", err.Error(), msg.SenderInfo.ID.Loggable())
			}
		}
	})
	return nil
}
