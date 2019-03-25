package p2p

import (
	"context"
	"fmt"
	"path"
	"reflect"
	"sync"
	"time"

	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/logger"
	"github.com/golang/protobuf/proto"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	kadopts "github.com/libp2p/go-libp2p-kad-dht/opts"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-protocol"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

const (
	ProtocolName    = "airbloc"
	ProtocolVersion = "1.0.0"
)

func pidFrom(topic string) protocol.ID {
	return protocol.ID(path.Join("/", ProtocolName, ProtocolVersion, "topic", topic))
}

type TopicHandler func(Server, context.Context, *IncomingMessage)

type Server interface {
	Start() error
	Stop()

	// server interfaces
	SubscribeTopic(string, proto.Message, TopicHandler)
	UnsubscribeTopic(string)

	// sender interfaces
	Send(context.Context, proto.Message, string, peer.ID) error
	Publish(proto.Message, string) error

	getHost() host.Host
	setContext(context.Context)
}

type Options struct {
	EnableMDNS          bool
	EnableLibP2PLogging bool
}

type AirblocServer struct {
	opt Options

	// controller
	lock   *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	// network
	id        cid.Cid
	host      host.Host
	dht       *kaddht.IpfsDHT
	nodekey   *key.Key
	bootinfos []peerstore.PeerInfo

	pubsub *pubsub.PubSub

	// for pubsub
	types        map[string]reflect.Type
	handlers     map[string]TopicHandler
	unsubscriber map[string]context.CancelFunc

	// log
	log *logger.Logger
}

func NewAirblocServer(
	nodekey *key.Key,
	addr multiaddr.Multiaddr,
	bootinfos []peerstore.PeerInfo,
	opts ...Options,
) (Server, error) {
	privKey, err := nodekey.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := &AirblocServer{
		ctx:       ctx,
		cancel:    cancel,
		lock:      new(sync.Mutex),
		nodekey:   nodekey,
		bootinfos: bootinfos,

		types:        make(map[string]reflect.Type),
		handlers:     make(map[string]TopicHandler),
		unsubscriber: make(map[string]context.CancelFunc),

		// log: logger.New("p2p"),
	}

	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(addr),
		libp2p.DisableRelay(),
	)
	if err != nil {
		return nil, err
	}
	server.log = logger.New("p2p-" + h.ID().String())

	ds := dsync.MutexWrap(datastore.NewMapDatastore())
	server.dht, err = kaddht.New(ctx, h, kadopts.Datastore(ds))
	if err != nil {
		return nil, err
	}
	server.host = routedhost.Wrap(h, server.dht)

	server.pubsub, err = pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize pubsub")
	}

	// to prevent breaking change
	if len(opts) >= 1 {
		server.opt = opts[0]
	} else {
		server.opt = Options{
			EnableMDNS:          true,
			EnableLibP2PLogging: false,
		}
	}

	server.log.Info("Initialized", logger.Attrs{
		"protocol":   fmt.Sprintf("%s %s", ProtocolName, ProtocolVersion),
		"on address": addr.String(),
	})
	return server, nil
}

// handleMessage receives all messages to server.
func (s *AirblocServer) handleMessage(message []byte, from peer.ID, topic string) {
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
func (s *AirblocServer) Start() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to bootstrap nodes for initializing DHT
	if len(s.bootinfos) > 0 {
		if err := bootstrapConnect(ctx, s.host, s.bootinfos); err != nil {
			return err
		}
		if err := s.dht.Bootstrap(s.ctx); err != nil {
			return errors.Wrap(err, "failed to bootstrap")
		}
	} else {
		s.log.Info("Warning: no bootstrap nodes are given. only local peers will be available.")
	}

	if s.opt.EnableMDNS {
		if err := startmDNSDiscovery(s.ctx, s.host); err != nil {
			return errors.Wrap(err, "could not start peer discovery via mDNS")
		}
	}
	return nil
}

func (s *AirblocServer) Stop() {
	s.cancel()
}

func (s *AirblocServer) SubscribeTopic(topic string, typeSample proto.Message, handler TopicHandler) {
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

func (s *AirblocServer) UnsubscribeTopic(topic string) {
	s.host.RemoveStreamHandler(pidFrom(topic))
	s.unsubscriber[topic]()

	s.lock.Lock()
	delete(s.types, topic)
	delete(s.handlers, topic)
	delete(s.unsubscriber, topic)
	s.lock.Unlock()
}

func (s *AirblocServer) Send(ctx context.Context, payload proto.Message, topic string, p peer.ID) error {
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

func (s *AirblocServer) Publish(payload proto.Message, topic string) error {
	s.log.Info("Broadcasting P2P message", logger.Attrs{"topic": topic})
	data, err := proto.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to encapsulate outgoing message")
	}
	return s.pubsub.Publish(topic, data)
}

// for test
func (s *AirblocServer) setContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *AirblocServer) getHost() host.Host {
	return s.host
}
