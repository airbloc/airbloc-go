package p2p

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
	"testing"
	"time"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/test/utils"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-pubsub"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	numOfPingPongServers = 10
	bootNodeTimeout      = 3 * time.Second
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

func NewTestServer(
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

func (s *testServer) setContext(ctx context.Context) { s.ctx = ctx }
func (s *testServer) getHost() host.Host             { return s.host }

func connect(t *testing.T, a, b host.Host) {
	err := b.Connect(context.Background(), peerstore.PeerInfo{
		ID:    a.ID(),
		Addrs: a.Addrs(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func connectAll(t *testing.T, hosts []Server) {
	for i, a := range hosts {
		for j, b := range hosts {
			if i == j {
				continue
			}

			connect(t, a.getHost(), b.getHost())
		}
	}
}

func setupTestPeers(t *testing.T, numPeers int) (keys []*key.Key, peers []Server, teardown func()) {
	var err error
	keys = make([]*key.Key, numPeers)
	peers = make([]Server, numPeers)
	for i := 0; i < numPeers; i++ {
		addr := multiaddr.StringCast("/ip4/127.0.0.1/tcp/" + strconv.Itoa(testutils.ReservePort()))
		keys[i], _ = key.Generate()
		peers[i] = NewTestServer(t, keys[i], addr)
		require.NoError(t, err)
	}

	connectAll(t, peers)

	teardown = func() {
		for _, p := range peers {
			p.Stop()
		}
	}
	return
}

func setupAirblocPeers(t *testing.T, numPeers int) (keys []*key.Key, peers []Server, teardown func()) {
	ctx, stopBootNode := context.WithCancel(context.Background())

	// wait for bootnode to be prepared
	time.Sleep(bootNodeTimeout)

	// launch bootnode for DHT
	addr := multiaddr.StringCast("/ip4/127.0.0.1/tcp/" + strconv.Itoa(testutils.ReservePort()))
	k, _ := key.Generate()
	bootInfo, err := StartBootstrapServer(ctx, k, addr)
	require.NoError(t, err)

	// wait for bootnode to be prepared
	time.Sleep(bootNodeTimeout)

	keys = make([]*key.Key, numPeers)
	peers = make([]Server, numPeers)
	for i := 0; i < numPeers; i++ {
		addr := multiaddr.StringCast("/ip4/127.0.0.1/tcp/" + strconv.Itoa(testutils.ReservePort()))
		keys[i], _ = key.Generate()
		peers[i], err = NewAirblocServer(keys[i], addr, []peerstore.PeerInfo{bootInfo}, Options{EnableMDNS: false})
		require.NoError(t, err)

		err = peers[i].Start()
		require.NoError(t, err)
	}

	teardown = func() {
		for _, p := range peers {
			p.Stop()
		}
		stopBootNode()
		time.Sleep(bootNodeTimeout)
	}
	return
}

var (
	pingMsg = &pb.TestPing{Message: "Ping"}
	pongMsg = &pb.TestPong{Message: "Pong"}
)

func init() {
	log.SetFlags(log.Lshortfile)
	logOutput := logger.NewStandardOutput(os.Stdout, "*", "*")
	logOutput.ColorsEnabled = true
	logger.SetLogger(logOutput)
}

func TestAirblocServer_Publish(t *testing.T) {
	_, servers, teardown := setupAirblocPeers(t, numOfPingPongServers)
	defer teardown()

	pingWaits := testutils.NewTimeoutWaitGroup(5 * time.Second)
	for _, s := range servers {
		s.SubscribeTopic("ping", pingMsg, func(_ Server, _ context.Context, msg *IncomingMessage) {
			log.Println("Ping", msg.Sender.Pretty(), msg.Payload.String())
			pingWaits.Done()

			err := s.Send(context.TODO(), pongMsg, "pong", msg.Sender)
			require.NoError(t, err)
		})
		time.Sleep(100 * time.Millisecond)
	}

	// decrease one, since it skips broadcasted msg from myself
	pingWaits.Add(numOfPingPongServers - 1)

	// first server will receive all pongs
	pongWaits := testutils.NewTimeoutWaitGroup(5 * time.Second)
	servers[0].SubscribeTopic("pong", pongMsg, func(_ Server, _ context.Context, msg *IncomingMessage) {
		log.Println("Pong", msg.Sender.Pretty(), msg.Payload.String())
		pongWaits.Done()
	})
	pongWaits.Add(numOfPingPongServers - 1)
	time.Sleep(1000 * time.Millisecond)

	// publish ping
	err := servers[0].Publish(pingMsg, "ping")
	require.NoError(t, err)
	require.NoError(t, pingWaits.Wait(), "ping timeout")
	require.NoError(t, pongWaits.Wait(), "pong timeout")
}

func TestAirblocServer_Send(t *testing.T) {
	keys, servers, teardown := setupAirblocPeers(t, 2)
	defer teardown()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	alice, bob := servers[0], servers[1]
	aliceAddress := keys[0].EthereumAddress.Hex()
	log.Printf("Alice address : %s\n", aliceAddress)
	log.Printf("Alice pubkey : %s\n", hex.EncodeToString(crypto.CompressPubkey(&keys[1].PublicKey)))

	// bob listens to alice, try to recover alice's address
	waitForBob := make(chan string, 1)
	bob.SubscribeTopic("ping", &pb.TestPing{}, func(s Server, ctx context.Context, message *IncomingMessage) {
		waitForBob <- message.SenderAddr.Hex()
	})

	err := alice.Send(ctx, pingMsg, "ping", bob.getHost().ID())
	require.NoError(t, err)

	timeout := time.After(5 * time.Second)

	select {
	case bobReceivedAddress := <-waitForBob:
		assert.Equal(t, aliceAddress, bobReceivedAddress)

	case <-timeout:
		assert.Fail(t, "Timeout")
	}
}
