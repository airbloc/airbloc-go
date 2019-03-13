package p2p

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-kbucket"
	"reflect"
	"sync"
	"time"

	"github.com/airbloc/logger"

	"github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"

	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	kadopts "github.com/libp2p/go-libp2p-kad-dht/opts"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

const (
	ProtocolName    = "airbloc"
	ProtocolVersion = "1.0.0"
)

type TopicHandler func(Server, context.Context, *IncomingMessage)

type Server interface {
	Start() error
	Stop()

	// server interfaces
	SubscribeTopic(string, proto.Message, TopicHandler)
	UnsubscribeTopic(string)

	Discovery()

	// sender interfaces
	Send(context.Context, proto.Message, string, peer.ID) error
	Publish(context.Context, proto.Message, string) error

	// for test
	Host() host.Host
	setContext(context.Context)
}

type AirblocServer struct {
	// controller
	mutex  *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	// network
	id        cid.Cid
	pid       Pid
	host      *PubSubHost
	dht       *kaddht.IpfsDHT
	nodekey   *key.Key
	bootinfos []peerstore.PeerInfo

	// for pubsub
	pubsub   *pubsub.PubSub
	types    map[string]reflect.Type
	handlers map[string]TopicHandler

	// log
	log *logger.Logger
}

func NewAirblocServer(
	nodekey *key.Key,
	addr multiaddr.Multiaddr,
	bootinfos []peerstore.PeerInfo,
) (Server, error) {
	privKey, err := nodekey.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, err
	}

	pid, err := NewPid(ProtocolName, ProtocolVersion)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate pid")
	}

	ctx, cancel := context.WithCancel(context.Background())

	server := &AirblocServer{
		ctx:       ctx,
		cancel:    cancel,
		mutex:     new(sync.Mutex),
		pid:       pid,
		nodekey:   nodekey,
		bootinfos: bootinfos,

		types:    make(map[string]reflect.Type),
		handlers: make(map[string]TopicHandler),
		log:      logger.New("p2p"),
	}

	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(addr),
	)
	if err != nil {
		return nil, err
	}

	ds := dsync.MutexWrap(datastore.NewMapDatastore())
	server.dht, err = kaddht.New(ctx, h, kadopts.Datastore(ds))
	if err != nil {
		return nil, err
	}

	h = routedhost.Wrap(h, server.dht)
	server.host = WrapPubSubHost(h, pid, server.handleMessage)

	server.pubsub, err = pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize pubsub")
	}

	server.log.Info("Initialized", logger.Attrs{
		"protocol":   fmt.Sprintf("%s %s", ProtocolName, ProtocolVersion),
		"on address": addr.String(),
	})
	return server, nil
}

// handleMessage receives all messages to server.
func (s *AirblocServer) handleMessage(message RawMessage) {
	topic := message.GetTopic()
	typ, ok := s.types[topic]
	if !ok {
		s.log.Error("Unknown topic: {}", message.GetTopic())
		return
	}
	handler := s.handlers[topic]

	msg, err := NewIncomingMessage(message, typ)
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

	if err := startmDNSDiscovery(s.ctx, s.host); err != nil {
		return errors.Wrap(err, "could not start peer discovery via mDNS")
	}
	go s.Discovery()
	return nil
}

func (s *AirblocServer) Stop() {
	s.cancel()
}

// Discovery finds and updates new peer connection every minute.
func (s *AirblocServer) Discovery() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	numOfPeers := 0
	s.updatePeer()
	for {
		select {
		case <-ticker.C:
			found := s.updatePeer()
			if numOfPeers != found {
				s.log.Info("Connected", logger.Attrs{"peers": found})
				numOfPeers = found
			}
		case <-s.ctx.Done():
			return
		}
	}
}

func (s *AirblocServer) clearPeer() {
	peerStore := s.host.Peerstore()
	for _, peerID := range peerStore.PeersWithAddrs() {
		peerStore.ClearAddrs(peerID)
	}
}

func (s *AirblocServer) updatePeer() int {
	idch, err := s.dht.GetClosestPeers(s.ctx, s.id.KeyString())
	if err == kbucket.ErrLookupFailure {
		s.log.Info("Warning: no peer available")
	} else if err != nil {
		s.log.Error("Failed to discovery peers", err)
	}
	s.log.Info("Found {} peers", len(idch))

	found := 0
	for id := range idch {
		info, err := s.dht.FindPeer(s.ctx, id)
		if err != nil {
			s.log.Error("Warning: Peer {id} found but failed to connect", err, logger.Attrs{"id": id.Pretty()})
			continue
		}
		s.host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.TempAddrTTL)
		found++
	}
	return found
}

func (s *AirblocServer) SubscribeTopic(topic string, typeSample proto.Message, handler TopicHandler) {
	// infer type from given sample
	typ := reflect.ValueOf(typeSample).Elem().Type()

	s.mutex.Lock()
	s.types[topic] = typ
	s.handlers[topic] = handler
	s.mutex.Unlock()

	sub, err := s.pubsub.Subscribe(topic)
	if err != nil {
		s.log.Wtf("error: failed to subscribe to topic {}", err, topic)
		return
	}
	go func() {
		defer sub.Cancel()
		defer s.log.Recover(logger.Attrs{"topic": topic})

		for {
			pubsubMsg, err := sub.Next(s.ctx)
			if err == context.Canceled {
				return
			}
			if err != nil {
				s.log.Error("failed to get message from {topic}", err, logger.Attrs{"topic": topic})
				continue
			}
			msg := RawMessage{}
			if err := proto.Unmarshal(pubsubMsg.GetData(), &msg); err != nil {
				s.log.Error("unmarshal error", err, logger.Attrs{"topic": topic})
				continue
			}
			s.handleMessage(msg)
		}
	}()
}

func (s *AirblocServer) UnsubscribeTopic(topic string) {
	s.mutex.Lock()
	delete(s.types, topic)
	delete(s.handlers, topic)
	s.mutex.Unlock()
}

func (s *AirblocServer) Send(ctx context.Context, payload proto.Message, topic string, p peer.ID) error {
	s.log.Info("Sending P2P message", logger.Attrs{
		"topic": topic,
		"id":    p.Pretty(),
	})
	msg, err := MarshalOutgoingMessage(payload, topic, s.host.ID(), s.pid.ProtocolID())
	if err != nil {
		return errors.Wrap(err, "send error")
	}
	return s.host.Send(ctx, *msg, p)
}

func (s *AirblocServer) Publish(ctx context.Context, payload proto.Message, topic string) error {
	s.log.Info("Broadcasting P2P message", logger.Attrs{"topic": topic})
	msg, err := MarshalOutgoingMessage(payload, topic, s.host.ID(), s.pid.ProtocolID())
	if err != nil {
		return errors.Wrap(err, "failed to encapsulate outgoing message")
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "marshal error")
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

func (s *AirblocServer) Host() host.Host {
	return s.host
}
