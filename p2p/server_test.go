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
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

var (
	pongMsg []byte
	pingMsg []byte
	keys    []*key.Key
	addrs   []multiaddr.Multiaddr
)

func init() {
	pongMsg, _ = proto.Marshal(&p2p.TestPing{Message: "World!"})
	pingMsg, _ = proto.Marshal(&p2p.TestPing{Message: "Hello"})

	for i := 1; i < 12; i++ {
		privKey, _ := key.Generate()
		addrStr := fmt.Sprintf("/ip4/127.0.0.1/tcp/24%02d", i)
		log.Println(addrStr)
		addr, _ := multiaddr.NewMultiaddr(addrStr)
		keys = append(keys, privKey)
		addrs = append(addrs, addr)
	}
}

func makeBasicServer(index int, bootnode bool, bootinfos ...peerstore.PeerInfo) (*Server, error) {
	return NewServer(
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

	size := 10
	servers := make([]*Server, size)
	for i := 0; i < size; i++ {
		server, err := makeBasicServer(i+1, false, bootinfo)
		assert.NoError(t, err)
		server.Start()
		server.ctx = ctx

		// ping
		server.RegisterTopic(p2p.Topic_TEST_PING.String(), &p2p.TestPing{}, func(message Message) {
			log.Println(message.Info.ID.Pretty(), message.Data.String())
			server.host.Send(message.ctx, &p2p.Message{
				Topic: p2p.Topic_TEST_PONG,
				Data:  pongMsg,
			}, message.Info.ID, "/airbloc")
		})

		// pong
		server.RegisterTopic(p2p.Topic_TEST_PONG.String(), &p2p.TestPong{}, func(message Message) {
			log.Println(message.Info.ID.Pretty(), message.Data.String())
		})

		servers[i] = server
	}

	err = servers[0].host.Send(ctx, &p2p.Message{
		Topic: p2p.Topic_TEST_PING,
		Data:  pingMsg,
	}, servers[1].host.ID(), "/airbloc")
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
}
