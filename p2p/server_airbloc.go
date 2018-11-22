package p2p

import (
	"context"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p/common"
	p2pr "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
)

type AirblocServer struct {
	// controller
	mutex  *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	// network
	id   cid.Cid
	host Host
	dht  *kaddht.IpfsDHT

	// database
	db localdb.Database

	// topic - handlers
	types    map[p2pr.Topic]reflect.Type
	topics   map[reflect.Type]string
	handlers map[reflect.Type]TopicHandler
}

func NewAirblocServer(
	localdb localdb.Database,
	identity *key.Key,
	addr multiaddr.Multiaddr,
	bootnode bool,
	bootinfos []peerstore.PeerInfo,
) (Server, error) {
	privKey, err := identity.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := &AirblocServer{
		ctx:    ctx,
		cancel: cancel,
		mutex:  new(sync.Mutex),

		db:       localdb,
		types:    make(map[p2pr.Topic]reflect.Type),
		topics:   make(map[reflect.Type]string),
		handlers: make(map[reflect.Type]TopicHandler),
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

	server.host = NewAirblocHost(NewBasicHost(h), 20)

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

	idVal := int32(p2pr.CID_AIRBLOC)

	v1b := cid.V1Builder{
		Codec:  uint64(idVal),
		MhType: multihash.KECCAK_256,
	}

	server.id, err = v1b.Sum([]byte(p2pr.CID_name[idVal]))
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "server error : failed to generate cid")
	}

	return server, nil
}

// DHT
func (s *AirblocServer) Discovery() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	s.updatePeer()
	for {
		select {
		case <-ticker.C:
			s.clearPeer()
			s.updatePeer()
		}
	}
}

func (s *AirblocServer) clearPeer() {
	peerStore := s.host.Peerstore()
	for _, peerID := range peerStore.PeersWithAddrs() {
		peerStore.ClearAddrs(peerID)
	}
}

func (s *AirblocServer) updatePeer() {
	idch, err := s.dht.GetClosestPeers(s.ctx, s.id.KeyString())
	if s.ctx.Err() != nil {
		log.Println("context error:", err)
		return
	}

	if err != nil {
		log.Println("failed to get closest peers:", err)
		return
	}

	for id := range idch {
		info, err := s.dht.FindPeer(s.ctx, id)
		if err != nil {
			log.Println("failed to find peer", id.Pretty(), ":", err)
			return
		}
		s.host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.TempAddrTTL)
	}
}

// api backend interfaces
func (s *AirblocServer) Start() error {
	pid, err := common.NewPid("airbloc", "0.0.1")
	if err != nil {
		return errors.Wrap(err, "failed to generate pid")
	}

	s.RegisterProtocol(pid, func(message common.ProtoMessage) {
		typ, ok := s.types[message.GetTopic()]
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

		msg, err := message.MakeMessage(s.ctx, typ)
		if err != nil {
			log.Printf("failed to make message : %+v", err)
			return
		}
		handler(s, s.ctx, msg)
	})

	go s.Discovery()
	return nil
}

func (s *AirblocServer) Stop() {
	s.cancel()
}

func (s *AirblocServer) RegisterProtocol(pid common.Pid, handler ProtocolHandler, adapters ...ProtocolAdapter) {
	s.host.RegisterProtocol(pid, handler, adapters...)
}

func (s *AirblocServer) UnregisterProtocol(pid common.Pid) {
	s.host.UnregisterProtocol(pid)
}

func (s *AirblocServer) RegisterTopic(topic string, msg proto.Message, handler TopicHandler) error {
	val, ok := p2pr.Topic_value[topic]
	if !ok {
		return errors.New("topic already registered")
	}
	typ := common.MessageType(msg)

	s.mutex.Lock()
	s.types[p2pr.Topic(val)] = typ
	s.topics[typ] = topic
	s.handlers[typ] = handler
	s.mutex.Unlock()

	return nil
}

func (s *AirblocServer) UnregisterTopic(topic string) error {
	val, ok := p2pr.Topic_value[topic]
	if !ok {
		return errors.New("invalid topic")
	}
	msgType := s.types[p2pr.Topic(val)]

	s.mutex.Lock()
	delete(s.types, p2pr.Topic(val))
	delete(s.topics, msgType)
	delete(s.handlers, msgType)
	s.mutex.Unlock()

	return nil
}

func (s *AirblocServer) Send(ctx context.Context, msg common.ProtoMessage, p peer.ID, pids ...common.Pid) error {
	return s.host.Send(ctx, msg, p, pids...)
}

func (s *AirblocServer) Publish(ctx context.Context, msg common.ProtoMessage, pids ...common.Pid) error {
	return s.host.Publish(ctx, msg, pids...)
}

func (s *AirblocServer) LocalDB() localdb.Database {
	return s.db
}

func (s *AirblocServer) MetaDB() metadb.Database {
	return nil
}

// for test
func (s *AirblocServer) setContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *AirblocServer) getHost() Host {
	return s.host
}

func (s *AirblocServer) bootInfo() (peerstore.PeerInfo, error) {
	return s.host.BootInfo()
}
