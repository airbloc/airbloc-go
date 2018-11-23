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
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
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

const (
	ProtocolName    = "airbloc"
	ProtocolVersion = "1.0.0"
)

type AirblocServer struct {
	// controller
	mutex  *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	// network
	id      cid.Cid
	pid     common.Pid
	host    Host
	dht     *kaddht.IpfsDHT
	nodekey *key.Key

	// database
	db localdb.Database

	// topic - handlers
	types    map[string]reflect.Type
	topics   map[reflect.Type]string
	handlers map[reflect.Type]TopicHandler
}

func NewServer(
	localdb localdb.Database,
	nodekey *key.Key,
	addr multiaddr.Multiaddr,
	bootnode bool,
	bootinfos []peerstore.PeerInfo,
) (Server, error) {
	privKey, err := nodekey.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, err
	}

	pid, err := common.NewPid(ProtocolName, ProtocolVersion)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate pid")
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := &AirblocServer{
		ctx:     ctx,
		cancel:  cancel,
		mutex:   new(sync.Mutex),
		pid:     pid,
		nodekey: nodekey,

		db:       localdb,
		types:    make(map[string]reflect.Type),
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

	h = routedhost.Wrap(h, server.dht)
	server.host = &AirblocHost{
		BasicHost: BasicHost{h},
		limit:     20,
	}

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

	idVal := int32(pb.CID_AIRBLOC)

	v1b := cid.V1Builder{
		Codec:  uint64(idVal),
		MhType: multihash.KECCAK_256,
	}

	server.id, err = v1b.Sum([]byte(pb.CID_name[idVal]))
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
	s.host.RegisterProtocol(s.pid, func(message common.ProtoMessage) {
		typ, ok := s.types[message.GetTopic()]
		if !ok {
			log.Println("unregistered topic")
			return
		}

		topic := s.topics[typ]
		handler := s.handlers[typ]
		if topic != message.Topic {
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

func (s *AirblocServer) SubscribeTopic(topic string, msg proto.Message, handler TopicHandler) error {
	typ := common.MessageType(msg)

	s.mutex.Lock()
	s.types[topic] = typ
	s.topics[typ] = topic
	s.handlers[typ] = handler
	s.mutex.Unlock()

	return nil
}

func (s *AirblocServer) UnsubscribeTopic(topic string) error {
	msgType := s.types[topic]

	s.mutex.Lock()
	delete(s.types, topic)
	delete(s.topics, msgType)
	delete(s.handlers, msgType)
	s.mutex.Unlock()

	return nil
}

func (s *AirblocServer) Send(ctx context.Context, msg proto.Message, topic string, p peer.ID) error {
	payload, err := common.NewProtoMessage(msg, topic)
	if err != nil {
		return errors.Wrap(err, "send error")
	}
	if err := payload.Sign(s.nodekey); err != nil {
		return errors.Wrap(err, "send error")
	}
	return s.host.Send(ctx, *payload, p, s.pid)
}

func (s *AirblocServer) Publish(ctx context.Context, msg proto.Message, topic string) error {
	payload, err := common.NewProtoMessage(msg, topic)
	if err != nil {
		return errors.Wrap(err, "publish error")
	}
	if err := payload.Sign(s.nodekey); err != nil {
		return errors.Wrap(err, "send error")
	}
	return s.host.Publish(ctx, *payload, s.pid)
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
