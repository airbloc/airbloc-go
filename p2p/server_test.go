package p2p

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/key"
	airbloc "github.com/airbloc/airbloc-go/p2p/airbloc"
	"github.com/airbloc/airbloc-go/p2p/common"
	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

const Size = 50

var (
	pongMsg common.ProtoMessage
	pingMsg common.ProtoMessage
	keys    []*key.Key
	addrs   []multiaddr.Multiaddr
)

func init() {
	pong, _ := proto.Marshal(&p2p.TestPing{Message: "World!"})
	pongMsg = common.ProtoMessage{
		Message: p2p.Message{
			Topic: p2p.Topic_TEST_PONG,
			Data:  pong,
		},
	}
	ping, _ := proto.Marshal(&p2p.TestPing{Message: "Hello"})
	pingMsg = common.ProtoMessage{
		Message: p2p.Message{
			Topic: p2p.Topic_TEST_PING,
			Data:  ping,
		},
	}

	for i := 1; i <= Size+1; i++ {
		privKey, _ := key.Generate()
		addrStr := fmt.Sprintf("/ip4/127.0.0.1/tcp/24%02d", i)
		addr, _ := multiaddr.NewMultiaddr(addrStr)
		keys = append(keys, privKey)
		addrs = append(addrs, addr)
	}
}

func makeBasicServer(index int, bootnode bool, bootinfos ...peerstore.PeerInfo) (Server, error) {
	return airbloc.NewServer(localdb.NewMemDB(), keys[index], addrs[index], bootnode, bootinfos)
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

	pid, err := common.NewPid("airbloc", "0.0.1")
	assert.NoError(t, err)

	for i := 1; i < Size; i++ {
		server, err := makeBasicServer(i, false, bootinfo)
		assert.NoError(t, err)
		server.Start()
		server.setContext(ctx)

		// ping
		server.RegisterTopic(p2p.Topic_TEST_PING.String(), &p2p.TestPing{}, testPingHandler)

		// pong
		server.RegisterTopic(p2p.Topic_TEST_PONG.String(), &p2p.TestPong{}, testPongHandler)

		servers[i] = server
	}

	err = servers[Size/2].Publish(ctx, pingMsg, pid)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)
}
