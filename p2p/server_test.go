package p2p

import (
	"context"
	"fmt"
	"github.com/airbloc/airbloc-go/p2p/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"

	"time"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/key"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
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

func makeBasicServer(ctx context.Context, index int, bootnode bool, bootinfos ...peerstore.PeerInfo) (Server, error) {
	server, err := NewAirblocServer(localdb.NewMemDB(), keys[index], addrs[index], bootnode, bootinfos)
	if err != nil {
		return nil, err
	}
	server.setContext(ctx)
	return server, nil
}

func handlePing(s Server, ctx context.Context, message common.Message) {
	log.Println("Ping", message.Info.ID.Pretty(), message.Data.String())

	s.Send(ctx, &pb.TestPing{Message: "World!"}, "ping", message.Info.ID)
}

func handlePong(s Server, ctx context.Context, message common.Message) {
	log.Println("Pong", message.Info.ID.Pretty(), message.Data.String())
}

func TestNewServer(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	bootnode, err := makeBasicServer(ctx, 0, true)
	assert.NoError(t, err)

	bootinfo, err := bootnode.bootInfo()
	assert.NoError(t, err)

	servers := make([]Server, Size)
	servers[0] = bootnode

	for i := 1; i < Size; i++ {
		server, err := makeBasicServer(ctx, i, false, bootinfo)
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

	bootnode, err := makeBasicServer(ctx, 0, true)
	assert.NoError(t, err)

	bootinfo, err := bootnode.bootInfo()
	assert.NoError(t, err)

	// make alice and bob
	alice, err := makeBasicServer(ctx, 1, false, bootinfo)
	assert.NoError(t, err)

	aliceAddress := keys[1].EthereumAddress.Hex()
	log.Printf("Alice address : %s\n", aliceAddress)

	bob, err := makeBasicServer(ctx, 2, false, bootinfo)
	assert.NoError(t, err)

	// start
	alice.Start()
	bob.Start()

	// bob listens to alice, try to recover alice's address
	waitForBob := make(chan string, 1)
	bob.SubscribeTopic("ping", &pb.TestPing{}, func (s Server, ctx context.Context, message common.Message) {
		recoveredAddress := crypto.PubkeyToAddress(*message.Sender)
		waitForBob <- recoveredAddress.Hex()
	})

	// TODO: without Connect(), it's not working
	// err = alice.getHost().Connect(ctx, bob.getHost().PeerInfo())
	// assert.NoError(t, err)

	err = alice.Publish(ctx, pingMsg, "ping")
	assert.NoError(t, err)

	bobReceivedAddress := <-waitForBob
	assert.Equal(t, aliceAddress, bobReceivedAddress)
}
