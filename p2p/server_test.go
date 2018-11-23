package p2p

import (
	"context"
	"fmt"
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

func makeBasicServer(index int, bootnode bool, bootinfos ...peerstore.PeerInfo) (Server, error) {
	return NewServer(localdb.NewMemDB(), keys[index], addrs[index], bootnode, bootinfos)
}

func TestNewServer(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	bootnode, err := makeBasicServer(0, true)
	assert.NoError(t, err)

	bootinfo, err := bootnode.bootInfo()
	assert.NoError(t, err)

	servers := make([]Server, Size)
	servers[0] = bootnode

	for i := 1; i < Size; i++ {
		server, err := makeBasicServer(i, false, bootinfo)
		assert.NoError(t, err)
		server.Start()
		server.setContext(ctx)

		// ping
		server.SubscribeTopic("ping", &pb.TestPing{}, Ping)

		// pong
		server.SubscribeTopic("pong", &pb.TestPong{}, Pong)

		servers[i] = server
	}

	err = servers[Size/2].Publish(ctx, pingMsg, "ping")
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)
}
