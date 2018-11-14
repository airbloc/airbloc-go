package p2p

import (
	"context"
	"sync"

	"reflect"

	"log"

	"github.com/airbloc/airbloc-go/key"
	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
)

type Server struct {
	// controller
	mutex  *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	// discovery
	id   cid.Cid
	host AirblocHost
	dht  *kaddht.IpfsDHT

	// topics
	types    map[p2p.Topic]reflect.Type
	topics   map[reflect.Type]string
	handlers map[reflect.Type]Handler
}

func NewServer(
	p2pId p2p.CID,
	identity *key.Key,
	addr multiaddr.Multiaddr,
	bootnode bool,
	bootinfos []peerstore.PeerInfo,
) (*Server, error) {
	privKey, err := identity.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := &Server{
		ctx:    ctx,
		cancel: cancel,
		mutex:  new(sync.Mutex),

		types:    make(map[p2p.Topic]reflect.Type),
		topics:   make(map[reflect.Type]string),
		handlers: make(map[reflect.Type]Handler),
	}

	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(addr),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	server.dht, err = kaddht.New(ctx, h)
	if err != nil {
		cancel()
		return nil, err
	}
	server.host = AirblocHost{routedhost.Wrap(h, server.dht)}

	if bootnode {
		if err := server.dht.Bootstrap(ctx); err != nil {
			cancel()
			return nil, errors.Wrap(err, "server error : failed to launch bootstrap node")
		}
	} else {
		if len(bootinfos) < 1 {
			cancel()
			return nil, errors.New("server error : input bootinfos should contains at least 1 element")
		}

		for _, bootinfo := range bootinfos {
			if err := h.Connect(ctx, bootinfo); err != nil {
				cancel()
				return nil, errors.Wrap(err, "server error : failed to connect bootstrap node")
			}
		}
	}

	idVal := int32(p2pId)

	v1b := cid.V1Builder{
		Codec:  uint64(idVal),
		MhType: multihash.KECCAK_256,
	}

	server.id, err = v1b.Sum([]byte(p2p.CID_name[idVal]))
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "server error : failed to generate cid")
	}

	return server, nil
}

// api backend interfaces
func (s *Server) Start() error {
	// start discovery
	s.host.RegisterHandler("/airbloc", s.handler)
	go s.peerWorker()
	return nil
}

func (s *Server) Stop() {
	s.cancel()
}

func (s *Server) handler(message p2p.Message) {
	typ, ok := s.types[message.Topic]
	if !ok {
		log.Println("unregistered topic")
		return
	}

	topic := s.topics[typ]
	handler := s.handlers[typ]
	if topic != message.Topic.String() {
		log.Println("message and topic mismatch")
		return
	}

	msg, ok := reflect.New(typ).Interface().(proto.Message)
	if !ok {
		log.Println("message is not protobuf message")
		return
	}

	if err := proto.Unmarshal(message.GetData(), msg); err != nil {
		log.Println("failed to unmarshal data :", err)
		return
	}

	pMsg := Message{
		ctx:  s.ctx,
		Data: msg,
		Info: peerstore.PeerInfo{ID: peer.ID(message.GetFrom())},
	}

	handler(pMsg)
}

func (s *Server) RegisterTopic(topic string, message proto.Message, handler Handler) error {
	val, ok := p2p.Topic_value[topic]
	if !ok {
		return errors.New("topic already registered")
	}
	typ := msgType(message)

	s.mutex.Lock()
	s.types[p2p.Topic(val)] = typ
	s.topics[typ] = topic
	s.handlers[typ] = handler
	s.mutex.Unlock()

	return nil
}

func (s *Server) UnregisterTopic(topic string) error {
	val, ok := p2p.Topic_value[topic]
	if !ok {
		return errors.New("invalid topic")
	}
	msgType := s.types[p2p.Topic(val)]

	s.mutex.Lock()
	delete(s.types, p2p.Topic(val))
	delete(s.topics, msgType)
	delete(s.handlers, msgType)
	s.mutex.Unlock()

	return nil
}
