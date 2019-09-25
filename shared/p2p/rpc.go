package p2p

import (
	"context"
	"time"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/logger"
	"github.com/golang/protobuf/proto"
	"github.com/klaytn/klaytn/common"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/pkg/errors"
)

var (
	ErrRPCTimeout         = errors.New("rpc response timeout")
	ErrInvalidRPCResponse = errors.New("invalid RPC response type")
	ErrAddressNotFound    = errors.New("address not found")
)

const (
	// DefaultClientTimeout is 1 minute, which is enough to wait a transaction in Ethereum.
	DefaultClientTimeout    = 1 * time.Minute
	DefaultLookupACKTimeout = 5 * time.Second
)

// RPC is both client and server interface for RPC over libp2p network layer,
// which supports calling RPC directly to peer identified by public key (peer.ID)
type RPC interface {
	Invoke(ctx context.Context, to common.Address, method string, args, reply proto.Message) (proto.Message, error)
	Handle(method string, argsType, replyType proto.Message, handler RPCHandler)
}

// RPCHandler is RPC call handler for server-side (receiver).
type RPCHandler func(ctx context.Context, from SenderInfo, req proto.Message) (proto.Message, error)

// SenderInfo is pretty information about RPC caller,
// which also contains an Ethereum address of caller.
type SenderInfo struct {
	Id   peer.ID
	Addr common.Address
}

// rpc is implementation of RPC interface by wrapping p2p.Server.
type rpc struct {
	server Server
	log    *logger.Logger

	// timeouts
	timeout    time.Duration
	ackTimeout time.Duration
}

// NewRPC creates new RPC instance.
func NewRPC(server Server) RPC {
	return &rpc{
		server:     server,
		timeout:    DefaultClientTimeout,
		ackTimeout: DefaultLookupACKTimeout,
		log:        logger.New("p2p-rpc"),
	}
}

// Invoke calls given method in given peer with arguments,
// and returns reply from the peer as a result.
func (r *rpc) Invoke(ctx context.Context, to common.Address, method string, args, reply proto.Message) (proto.Message, error) {
	requestTopic, replyTopic := topicFromMethod(method)

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	// lookup peer ID of given address
	targetId, err := Lookup(ctx, r.server, to, r.ackTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "lookup error")
	}

	// prepare to receive reply
	waitForResponse := make(chan *IncomingMessage)
	responseCallback := func(_ Server, _ context.Context, msg *IncomingMessage) {
		waitForResponse <- msg
	}
	r.server.SubscribeTopic(replyTopic, &pb.RPCResponse{}, responseCallback)

	// send request message
	if err := r.server.Send(ctx, args, requestTopic, targetId); err != nil {
		return nil, errors.Wrapf(err, "failed to invoke %s", method)
	}

	var replyMsg *IncomingMessage
WaitingForReply:
	for {
		select {
		case replyMsg = <-waitForResponse:
			// messages from other senders except the original recipient are ignored.
			if replyMsg.Sender == targetId {
				break WaitingForReply
			}
		case <-ctx.Done():
			return nil, ErrRPCTimeout
		}
	}

	response, ok := replyMsg.Payload.(*pb.RPCResponse)
	if !ok {
		// response type must be pb.RPCResponse
		return nil, ErrInvalidRPCResponse
	}
	if !response.Ok {
		// just return error from the response
		return nil, errors.Wrapf(errors.New(response.GetError()), "failed to invoke RPC %s", method)
	}
	if err := proto.Unmarshal(response.GetSuccessfulReply(), reply); err != nil {
		return nil, errors.Wrap(err, "invalid reply type returned")
	}
	return reply, nil
}

func (r *rpc) Handle(method string, argsType, replyType proto.Message, handler RPCHandler) {
	requestTopic, replyTopic := topicFromMethod(method)

	callback := func(_ Server, ctx context.Context, msg *IncomingMessage) {
		from := SenderInfo{
			Id:   msg.Sender,
			Addr: msg.SenderAddr,
		}
		reply, err := handler(ctx, from, msg.Payload)

		// make response
		response := new(pb.RPCResponse)
		if err != nil {
			response.Ok = false
			response.Error = err.Error()
		} else {
			response.SuccessfulReply, err = proto.Marshal(reply)
			if err != nil {
				response.Ok = false
				response.Error = errors.Wrap(err, "invalid reply").Error()
			} else {
				response.Ok = true
			}
		}

		// send response using reply topic
		if err := r.server.Send(ctx, response, replyTopic, from.Id); err != nil {
			r.log.Error("error: failed to reply: %s", err.Error(), logger.Attrs{
				"from":   from.Id.Pretty(),
				"method": method,
			})
		}
	}
	r.server.SubscribeTopic(requestTopic, argsType, callback)
}

func topicFromMethod(method string) (requestTopic, replyTopic string) {
	requestTopic = "/rpc/" + method
	replyTopic = "/rpc/" + method + "/reply"
	return
}
