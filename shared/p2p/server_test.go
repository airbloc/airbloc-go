package p2p

import (
	"context"
	"encoding/hex"
	"github.com/libp2p/go-libp2p-host"
	"github.com/multiformats/go-multiaddr"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/test/utils"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	numOfPingPongServers = 10
	bootNodeTimeout      = 2 * time.Second
)

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

func setupBasicPeers(t *testing.T, numPeers int) (keys []*key.Key, peers []Server, teardown func()) {
	var err error
	keys = make([]*key.Key, numPeers)
	peers = make([]Server, numPeers)
	for i := 0; i < numPeers; i++ {
		addr := multiaddr.StringCast("/ip4/127.0.0.1/tcp/" + strconv.Itoa(testutils.ReservePort()))
		keys[i], _ = key.Generate()
		peers[i] = NewBasicServer(t, keys[i], addr)
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
	time.Sleep(bootNodeTimeout / 2)

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
		time.Sleep(bootNodeTimeout / 2)
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
