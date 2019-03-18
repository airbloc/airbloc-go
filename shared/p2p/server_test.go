package p2p

import (
	"context"
	"encoding/hex"
	"github.com/airbloc/airbloc-go/test/utils"
	"github.com/airbloc/logger"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"time"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

const (
	numOfPingPongServers = 10
	bootNodeTimeout      = 2 * time.Second
)

func setupTestPeers(t *testing.T, numPeers int) (keys []*key.Key, peers []Server, teardown func()) {
	ctx, stopBootNode := context.WithCancel(context.Background())

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
		for _, peer := range peers {
			peer.Stop()
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
	logOutput := logger.NewStandardOutput(os.Stdout, "*", "*")
	logOutput.ColorsEnabled = true
	logger.SetLogger(logOutput)
}

func TestAirblocServer_Publish(t *testing.T) {
	_, servers, teardown := setupTestPeers(t, numOfPingPongServers)
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
	keys, servers, teardown := setupTestPeers(t, 2)
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

	err := alice.Send(ctx, pingMsg, "ping", bob.Host().ID())
	require.NoError(t, err)

	timeout := time.After(5 * time.Second)

	select {
	case bobReceivedAddress := <-waitForBob:
		assert.Equal(t, aliceAddress, bobReceivedAddress)

	case <-timeout:
		assert.Fail(t, "Timeout")
	}
}
