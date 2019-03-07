package p2p

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"github.com/airbloc/airbloc-go/shared/p2p/common"
	"github.com/ethereum/go-ethereum/crypto"

	"time"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

const Size = 50

var (
	pongMsg *pb.TestPing
	pingMsg *pb.TestPing
	keys    []*key.Key
	addrs   []multiaddr.Multiaddr
)

func init() {
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

func makeBasicServer(ctx context.Context, index int, bootinfos ...peerstore.PeerInfo) (Server, error) {
	server, err := NewAirblocServer(keys[index], addrs[index], bootinfos)
	if err != nil {
		return nil, err
	}
	server.setContext(ctx)
	return server, nil
}

func handlePing(s Server, ctx context.Context, message common.Message) {
	log.Println("Ping", message.SenderInfo.ID.Pretty(), message.Data.String())

	s.Send(ctx, &pb.TestPing{Message: "World!"}, "ping", message.SenderInfo.ID)
}

func handlePong(s Server, ctx context.Context, message common.Message) {
	log.Println("Pong", message.SenderInfo.ID.Pretty(), message.Data.String())
}

func TestNewServer(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	bootinfo, err := StartBootstrapServer(ctx, keys[0], addrs[0])
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	servers := make([]Server, Size)
	for i := 1; i < Size; i++ {
		server, err := makeBasicServer(ctx, i, bootinfo)
		assert.NoError(t, err)
		server.Start()

		server.SubscribeTopic("ping", &pb.TestPing{}, handlePing)
		server.SubscribeTopic("pong", &pb.TestPong{}, handlePong)

		servers[i] = server
	}

	err = servers[Size/2].Publish(ctx, pingMsg, "ping")
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)
}

func TestAirblocHost_Publish(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootinfo, err := StartBootstrapServer(ctx, keys[0], addrs[0])
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	// make alice and bob
	alice, err := makeBasicServer(ctx, 1, bootinfo)
	assert.NoError(t, err)
	alice.Start()

	aliceAddress := keys[1].EthereumAddress.Hex()
	log.Printf("Alice address : %s\n", aliceAddress)
	log.Printf("Alice pubkey : %s\n", hex.EncodeToString(crypto.CompressPubkey(&keys[1].PublicKey)))

	bob, err := makeBasicServer(ctx, 2, bootinfo)
	assert.NoError(t, err)
	bob.Start()

	time.Sleep(2 * time.Second)

	// bob listens to alice, try to recover alice's address
	waitForBob := make(chan string, 1)
	bob.SubscribeTopic("ping", &pb.TestPing{}, func(s Server, ctx context.Context, message common.Message) {
		waitForBob <- message.SenderAddr.Hex()
	})

	err = alice.Publish(ctx, pingMsg, "ping")
	assert.NoError(t, err)

	bobReceivedAddress := <-waitForBob
	assert.Equal(t, aliceAddress, bobReceivedAddress)
}
