package p2p

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/airbloc/logger"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"time"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

const Size = 5

var (
	pongMsg *pb.TestPing
	pingMsg *pb.TestPing
	keys    []*key.Key
	addrs   []multiaddr.Multiaddr
)

func init() {
	logOutput := logger.NewStandardOutput(os.Stdout, "*", "*")
	logOutput.ColorsEnabled = true
	logger.SetLogger(logOutput)

	pongMsg = &pb.TestPing{Message: "World!"}
	pingMsg = &pb.TestPing{Message: "Hello"}

	for i := 1; i <= Size+1; i++ {
		privKey, _ := key.Generate()
		addrStr := fmt.Sprintf("/ip4/127.0.0.1/tcp/24%02d", i)
		addr, _ := multiaddr.NewMultiaddr(addrStr)
		keys = append(keys, privKey)
		addrs = append(addrs, addr)
	}
}

func makeBasicServer(index int, bootinfos ...peerstore.PeerInfo) (Server, error) {
	server, err := NewAirblocServer(keys[index], addrs[index], bootinfos)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func handlePing(s Server, ctx context.Context, message *IncomingMessage) {
	log.Println("Ping", message.SenderInfo.ID.Pretty(), message.Payload.String())

	s.Send(ctx, &pb.TestPing{Message: "World!"}, "ping", message.SenderInfo.ID)
}

func handlePong(s Server, ctx context.Context, message *IncomingMessage) {
	log.Println("Pong", message.SenderInfo.ID.Pretty(), message.Payload.String())
}

func TestNewServer(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootinfo, err := StartBootstrapServer(ctx, keys[0], addrs[0])
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	servers := make([]Server, Size)
	for i := 1; i < Size; i++ {
		server, err := makeBasicServer(i, bootinfo)
		require.NoError(t, err)
		err = server.Start()
		require.NoError(t, err)
		defer server.Stop()

		server.SubscribeTopic("ping", &pb.TestPing{}, handlePing)
		server.SubscribeTopic("pong", &pb.TestPong{}, handlePong)

		servers[i] = server
	}

	err = servers[Size/2].Publish(ctx, pingMsg, "ping")
	require.NoError(t, err)

	time.Sleep(1 * time.Second)
}

func TestAirblocHost_Publish(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootinfo, err := StartBootstrapServer(ctx, keys[0], addrs[0])
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	// make alice and bob
	alice, err := makeBasicServer(1, bootinfo)
	require.NoError(t, err)
	err = alice.Start()
	require.NoError(t, err)
	defer alice.Stop()

	aliceAddress := keys[1].EthereumAddress.Hex()
	log.Printf("Alice address : %s\n", aliceAddress)
	log.Printf("Alice pubkey : %s\n", hex.EncodeToString(crypto.CompressPubkey(&keys[1].PublicKey)))

	bob, err := makeBasicServer(2, bootinfo)
	require.NoError(t, err)
	err = bob.Start()
	require.NoError(t, err)
	defer bob.Stop()

	time.Sleep(2 * time.Second)

	// bob listens to alice, try to recover alice's address
	waitForBob := make(chan string, 1)
	bob.SubscribeTopic("ping", &pb.TestPing{}, func(s Server, ctx context.Context, message *IncomingMessage) {
		waitForBob <- message.SenderAddr.Hex()
	})

	err = alice.Publish(ctx, pingMsg, "ping")
	require.NoError(t, err)

	timeout := time.After(5 * time.Second)

	select {
	case bobReceivedAddress := <-waitForBob:
		assert.Equal(t, aliceAddress, bobReceivedAddress)

	case <-timeout:
		assert.Fail(t, "Timeout")
	}
}
