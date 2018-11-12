package p2p

import (
	"context"
	"reflect"
	"sync"

	"log"

	"github.com/airbloc/airbloc-go/key"
	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
)

type Server struct {
	ctx      context.Context
	cancel   context.CancelFunc
	id       cid.Cid
	host     host.Host
	dht      *kaddht.IpfsDHT
	fsub     *pubsub.PubSub
	mutex    *sync.Mutex
	topics   map[reflect.Type]string
	handlers map[reflect.Type]Handler
}

func NewServer(
	p2pId p2p.CID,
	identity *key.Key,
	listenAddr multiaddr.Multiaddr,
	bootnode bool,
	bootinfos []peerstore.PeerInfo,
) (*Server, error) {
	privKey, err := identity.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(listenAddr),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	dht, err := kaddht.New(ctx, h)
	if err != nil {
		cancel()
		return nil, err
	}
	h = routedhost.Wrap(h, dht)

	if bootnode {
		if err := dht.Bootstrap(ctx); err != nil {
			return nil, errors.Wrap(err, "failed to launch bootstrap node")
		}
	} else {
		if len(bootinfos) < 1 {
			return nil, errors.New("input bootinfos should contains at least 1 element")
		}

		for _, bootinfo := range bootinfos {
			if err := h.Connect(ctx, bootinfo); err != nil {
				return nil, errors.Wrap(err, "failed to connect bootstrap node")
			}
		}
	}

	idVal := int32(p2pId)

	v1b := cid.V1Builder{
		Codec:  uint64(idVal),
		MhType: multihash.KECCAK_256,
	}

	id, err := v1b.Sum([]byte(p2p.CID_name[idVal]))
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate cid")
	}

	fsub, err := pubsub.NewFloodSub(ctx, h)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate pubsub instance")
	}

	return &Server{
		ctx:      ctx,
		cancel:   cancel,
		id:       id,
		host:     h,
		dht:      dht,
		fsub:     fsub,
		mutex:    new(sync.Mutex),
		topics:   make(map[reflect.Type]string),
		handlers: make(map[reflect.Type]Handler),
	}, nil
}

// api backend interfaces
func (s *Server) Start() error {
	// start discovery
	go s.peerWorker()
	return nil
}

func (s *Server) Stop() {
	s.cancel()
}

func (s *Server) RegisterTopic(topic p2p.Topic, message proto.Message, handler Handler) error {
	msgType := messageType(message)
	topicName := p2p.Topic_name[int32(topic)]
	_, ok := s.topics[msgType]
	if ok {
		return errors.New("topic already registered")
	}
	s.topics[msgType] = topicName
	s.handlers[msgType] = handler

	sub, err := s.fsub.Subscribe(topicName)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe topic")
	}
	go func() {
		defer sub.Cancel()

		for {
			msg, err := sub.Next(s.ctx)
			if s.ctx.Err() != nil {
				log.Println("context error :", err)
				return
			}

			if err != nil {
				log.Println("failed to read message :", err)
				return
			}

			d, ok := reflect.New(msgType).Interface().(proto.Message)
			if !ok {
				log.Println("received message is not a protobuf message :", msgType)
			} else {
				if err := proto.Unmarshal(msg.Data, d); err != nil {
					log.Println("failed to decode data :", err)
					continue
				}
				handler(peer.ID(msg.GetFrom()), d)
			}
		}
	}()
	return nil
}

func (s *Server) Publish(topic p2p.Topic, message proto.Message) error {
	topicName := p2p.Topic_name[int32(topic)]
	if topicName != s.topics[messageType(message)] {
		return errors.New("invalid message type")
	}

	data, err := proto.Marshal(message)
	if err != nil {
		return errors.Wrap(err, "failed to marshal proto message")
	}

	return s.fsub.Publish(topicName, data)
}
