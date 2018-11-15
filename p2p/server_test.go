package p2p

import (
	"context"
	"fmt"
	"log"
	"testing"

	"time"

	"github.com/airbloc/airbloc-go/key"
	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

const Size = 50

var (
	pid     = protocol.ID("/airbloc")
	pongMsg p2p.Message
	pingMsg p2p.Message
	keys    []*key.Key
	addrs   []multiaddr.Multiaddr
)

func init() {
	pong, _ := proto.Marshal(&p2p.TestPing{Message: "World!"})
	pongMsg = p2p.Message{
		Topic: p2p.Topic_TEST_PONG,
		Data:  pong,
	}
	ping, _ := proto.Marshal(&p2p.TestPing{Message: "Hello"})
	pingMsg = p2p.Message{
		Topic: p2p.Topic_TEST_PING,
		Data:  ping,
	}

	for i := 1; i <= Size+1; i++ {
		privKey, _ := key.Generate()
		addrStr := fmt.Sprintf("/ip4/127.0.0.1/tcp/24%02d", i)
		addr, _ := multiaddr.NewMultiaddr(addrStr)
		keys = append(keys, privKey)
		addrs = append(addrs, addr)
	}
}

func makeBasicServer(index int, bootnode bool, bootinfos ...peerstore.PeerInfo) (*AirblocServer, error) {
	return NewServer(
		"airbloc",
		"0.0.1",
		p2p.CID_AIRBLOC, keys[index],
		addrs[index], bootnode, bootinfos)
}

func TestNewServer(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	bootnode, err := makeBasicServer(0, true)
	assert.NoError(t, err)

	bootinfo, err := bootnode.host.BootInfo()
	assert.NoError(t, err)

	servers := make([]*AirblocServer, Size)
	servers[0] = bootnode

	for i := 1; i < Size; i++ {
		server, err := makeBasicServer(i, false, bootinfo)
		assert.NoError(t, err)
		server.Start()
		server.ctx = ctx

		// ping
		server.RegisterTopic(p2p.Topic_TEST_PING.String(), &p2p.TestPing{}, func(message Message) {
			log.Println("Recevied", message.Info.ID.Pretty(), message.Data.String())
			server.host.Send(message.ctx, pongMsg, message.Info.ID)
		})

		// pong
		server.RegisterTopic(p2p.Topic_TEST_PONG.String(), &p2p.TestPong{}, func(message Message) {
			log.Println("Received", message.Info.ID.Pretty(), message.Data.String())
		})

		servers[i] = server
	}

	err = servers[Size/2].host.Publish(ctx, pingMsg)
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
}
