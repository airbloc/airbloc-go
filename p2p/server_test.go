package p2p

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/airbloc/airbloc-go/logger"
	"log"
	"os"
	"testing"

	"github.com/airbloc/airbloc-go/p2p/common"
	"github.com/ethereum/go-ethereum/crypto"

	"time"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/key"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

var testBootNodeAddr = "/ip4/18.179.100.185/tcp/9100/ipfs/16Uiu2HAm3wQccsTNbbAo3rHy6Mj1i5ia7cyMEJ7WesKoSzsnB1rr"

const Size = 50

var (
	pingMsg *pb.TestPing
	keys    []*key.Key
	addrs   []multiaddr.Multiaddr
)

func init() {
	logger.Setup(os.Stdout, "*", "*")

	pingMsg = &pb.TestPing{Message: "Hello"}

	for i := 1; i <= Size+1; i++ {
		privKey, _ := key.Generate()
		addrStr := fmt.Sprintf("/ip4/127.0.0.1/tcp/24%02d", i)
		addr, _ := multiaddr.NewMultiaddr(addrStr)
		keys = append(keys, privKey)
		addrs = append(addrs, addr)
	}
}

func makeBasicServer(ctx context.Context, index int, bootnode bool, bootinfos ...peerstore.PeerInfo) (Server, error) {
	server, err := NewAirblocServer(localdb.NewMemDB(), keys[index], addrs[index], bootnode, bootinfos)
	if err != nil {
		return nil, err
	}
	server.setContext(ctx)
	return server, nil
}

func TestNewServer(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	bootnode, err := makeBasicServer(ctx, 0, true)
	assert.NoError(t, err)

	bootinfo, err := bootnode.BootInfo()
	assert.NoError(t, err)

	servers := make([]Server, Size)
	servers[0] = bootnode

	for i := 1; i < Size; i++ {
		server, err := makeBasicServer(ctx, i, false, bootinfo)
		assert.NoError(t, err)
		server.Start()

		server.SubscribeTopic("ping", &pb.TestPing{}, testPingHandler)
		server.SubscribeTopic("pong", &pb.TestPong{}, testPongHandler)

		servers[i] = server
	}

	err = servers[Size/2].Publish(ctx, pingMsg, "ping")
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)
}

func TestAirblocHost_Publish(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bootnode, err := makeBasicServer(ctx, 0, true)
	assert.NoError(t, err)

	bootinfo, err := bootnode.BootInfo()
	assert.NoError(t, err)

	// make alice and bob
	alice, err := makeBasicServer(ctx, 1, false, bootinfo)
	alice.Start()
	assert.NoError(t, err)

	aliceAddress := keys[1].EthereumAddress.Hex()
	log.Printf("Alice address : %s\n", aliceAddress)
	log.Printf("Alice pubkey : %s\n", hex.EncodeToString(crypto.CompressPubkey(&keys[1].PublicKey)))

	bob, err := makeBasicServer(ctx, 2, false, bootinfo)
	bob.Start()
	assert.NoError(t, err)

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

func TestBasicHost_Connect(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ma, err := multiaddr.NewMultiaddr(testBootNodeAddr)
	assert.NoError(t, err)

	bootinfo, err := peerstore.InfoFromP2pAddr(ma)
	assert.NoError(t, err)

	s1, err := makeBasicServer(ctx, 1, false, *bootinfo)
	s1.Start()
	assert.NoError(t, err)
	s2, err := makeBasicServer(ctx, 2, false, *bootinfo)
	s2.Start()
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	s1Host := s1.getHost()
	s2Host := s2.getHost()

	for _, peer := range s1Host.Peerstore().Peers() {
		log.Println("s1 :", s1Host.Peerstore().PeerInfo(peer))
	}
	for _, peer := range s2Host.Peerstore().Peers() {
		log.Println("s2 :", s1Host.Peerstore().PeerInfo(peer))
	}
}
