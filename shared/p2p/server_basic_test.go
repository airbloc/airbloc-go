package p2p

import (
	"context"
	"fmt"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/logger"
	"github.com/golang/protobuf/proto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-pubsub"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"reflect"
	"sync"
	"testing"
)

type testServer struct {
	// controller
	lock   *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	// network
	id      cid.Cid
	host    host.Host
	nodekey *key.Key

	pubsub *pubsub.PubSub

	// for pubsub
	types        map[string]reflect.Type
	handlers     map[string]TopicHandler
	unsubscriber map[string]context.CancelFunc

	// log
	log *logger.Logger
}

func NewBasicServer(
	t *testing.T,
	nodekey *key.Key,
	addr multiaddr.Multiaddr,
) Server {
	privKey, err := nodekey.DeriveLibp2pKeyPair()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	server := &testServer{
		ctx:     ctx,
		cancel:  cancel,
		lock:    new(sync.Mutex),
		nodekey: nodekey,

		types:        make(map[string]reflect.Type),
		handlers:     make(map[string]TopicHandler),
		unsubscriber: make(map[string]context.CancelFunc),
	}

	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(addr),
		libp2p.DisableRelay(),
	)
	require.NoError(t, err)

	dht, err := kaddht.New(ctx, h)
	require.NoError(t, err)

	server.host = rhost.Wrap(h, dht)
	server.log = logger.New("p2p-" + h.ID().String())
	server.pubsub, err = pubsub.NewGossipSub(ctx, h)
	require.NoError(t, err)

	server.log.Info("Initialized", logger.Attrs{
		"protocol":   fmt.Sprintf("%s %s", ProtocolName, ProtocolVersion),
		"on address": addr.String(),
	})
	return server
}

// handleMessage receives all messages to server.
func (s *testServer) handleMessage(message []byte, from peer.ID, topic string) {
	typ, ok := s.types[topic]
	if !ok {
		s.log.Error("Unknown topic: {}", topic)
		return
	}
	handler := s.handlers[topic]

	msg, err := newIncomingMessage(message, typ, from)
	if err != nil {
		s.log.Error("Failed to make message", err)
		return
	}

	timer := s.log.Timer()
	handler(s, s.ctx, msg)
	timer.End("Received message", logger.Attrs{
		"from":  msg.SenderAddr.String(),
		"topic": topic,
	})
}

// api backend interfaces
func (s *testServer) Start() error { return nil }
func (s *testServer) Stop()        { s.cancel() }

func (s *testServer) SubscribeTopic(topic string, typeSample proto.Message, handler TopicHandler) {
	s.log.Debug("subscribed to {topic}", logger.Attrs{"topic": topic})

	// infer type from given sample
	typ := reflect.ValueOf(typeSample).Elem().Type()

	// subscription finishes when the context is cancelled
	subCtx, cancel := context.WithCancel(s.ctx)

	s.lock.Lock()
	s.types[topic] = typ
	s.handlers[topic] = handler
	s.unsubscriber[topic] = cancel
	s.lock.Unlock()

	sub, err := s.pubsub.Subscribe(topic)
	if err != nil {
		s.log.Wtf("error: failed to subscribe to topic {}", err, topic)
		return
	}
	go func() {
		defer sub.Cancel()
		defer s.log.Recover(logger.Attrs{"topic": topic})
		for {
			pubsubMsg, err := sub.Next(subCtx)
			if err == context.Canceled {
				return
			}
			if err != nil {
				s.log.Error("failed to get message from {topic}", err, logger.Attrs{"topic": topic})
				continue
			}
			if pubsubMsg == nil || pubsubMsg.GetFrom() == s.host.ID() {
				continue
			}
			s.handleMessage(pubsubMsg.GetData(), pubsubMsg.GetFrom(), topic)
		}
	}()

	// we need to also register a stream handler,
	// for case of direct connection via Send().
	s.host.SetStreamHandler(pidFrom(topic), func(stream net.Stream) {
		defer s.log.Recover(logger.Attrs{"topic": topic})

		msg, err := readDirectMessageFrom(stream)
		if err != nil {
			s.log.Error("failed to read message from {topic}", err, logger.Attrs{"topic": topic})
			return
		}
		s.handleMessage(msg.GetPayload(), peer.ID(msg.GetFrom()), topic)
	})
}

func (s *testServer) UnsubscribeTopic(topic string) {
	s.host.RemoveStreamHandler(pidFrom(topic))
	s.unsubscriber[topic]()

	s.lock.Lock()
	delete(s.types, topic)
	delete(s.handlers, topic)
	delete(s.unsubscriber, topic)
	s.lock.Unlock()
}

func (s *testServer) Send(ctx context.Context, payload proto.Message, topic string, p peer.ID) error {
	s.log.Info("Sending P2P message", logger.Attrs{
		"topic": topic,
		"id":    p.Pretty(),
	})
	msg, err := marshalDirectMessage(s, payload)
	if err != nil {
		return errors.Wrap(err, "send error")
	}
	stream, err := s.host.NewStream(ctx, p, pidFrom(topic))
	if err != nil {
		return errors.Wrap(err, "failed to open stream")
	}
	defer stream.Close()
	_, err = stream.Write(msg)
	return err
}

func (s *testServer) Publish(payload proto.Message, topic string) error {
	s.log.Info("Broadcasting P2P message", logger.Attrs{"topic": topic})
	data, err := proto.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to encapsulate outgoing message")
	}
	return s.pubsub.Publish(topic, data)
}

// for test
func (s *testServer) setContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *testServer) getHost() host.Host {
	return s.host
}
