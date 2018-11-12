package p2p

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
)

var (
	topic = "frostornge"
	point cid.Cid
	keys  []crypto.PrivKey
	addrs []multiaddr.Multiaddr
)

type Host struct {
	ctx context.Context
	host.Host
	dht *kaddht.IpfsDHT
	sub *pubsub.PubSub
}

func init() {
	for i := 1; i < 10; i++ {
		privKey, _, _ := crypto.GenerateEd25519Key(rand.Reader)
		addrStr := fmt.Sprintf("/ip4/127.0.0.1/tcp/247%d", i)
		addr, _ := multiaddr.NewMultiaddr(addrStr)
		keys = append(keys, privKey)
		addrs = append(addrs, addr)
	}

	v1b := cid.V1Builder{
		Codec:  cid.Raw,
		MhType: multihash.KECCAK_256,
	}
	point, _ = v1b.Sum([]byte(topic))
}

func makeBasicHost(
	t *testing.T,
	ctx context.Context,
	index int,
	bootnode bool,
	bootinfos []peerstore.PeerInfo,
) *Host {
	h, err := libp2p.New(ctx, libp2p.Identity(keys[index]), libp2p.ListenAddrs(addrs[index]))
	assert.NoError(t, err)
	log.Println("host", index, "ID:", h.ID().Pretty())

	dht, err := kaddht.New(ctx, h)
	assert.NoError(t, err)
	h = routedhost.Wrap(h, dht)

	if bootnode {
		assert.NoError(t, dht.Bootstrap(ctx))
	} else {
		for _, bootinfo := range bootinfos {
			assert.NoError(t, h.Connect(ctx, bootinfo))
		}
	}

	sub, err := pubsub.NewFloodSub(ctx, h)
	assert.NoError(t, err)

	return &Host{
		ctx:  ctx,
		Host: h,
		dht:  dht,
		sub:  sub,
	}
}

func (h *Host) subscribe(t *testing.T, topic string) {
	s, err := h.sub.Subscribe(topic)
	assert.NoError(t, err)

	for _, peer := range h.Peerstore().Peers() {
		if strings.EqualFold(h.ID().Pretty(), peer.Pretty()) {
			continue
		}
		err := h.Connect(h.ctx, peerstore.PeerInfo{ID: peer})
		if err != nil {
			log.Println(h.ID().Pretty(), peer.Pretty(), err)
		}
		assert.NoError(t, err)
	}

	go func() {
		defer s.Cancel()
		for {
			msg, err := s.Next(h.ctx)
			if h.ctx.Err() != nil {
				log.Println("context error")
				return
			}
			if err != nil {
				log.Println(h.ID().Pretty(), err)
			}
			log.Println(h.ID().Pretty(), string(msg.GetData()))
		}
	}()
}

func TestNewServer(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	host1 := makeBasicHost(t, ctx, 1, true, nil)
	iaddr, err := multiaddr.NewMultiaddr("/ipfs/" + host1.ID().Pretty())
	bootinfo, err := peerstore.InfoFromP2pAddr(host1.Addrs()[1].Encapsulate(iaddr))
	assert.NoError(t, err)
	host2 := makeBasicHost(t, ctx, 2, false, []peerstore.PeerInfo{*bootinfo})
	host3 := makeBasicHost(t, ctx, 3, false, []peerstore.PeerInfo{*bootinfo})
	host4 := makeBasicHost(t, ctx, 4, false, []peerstore.PeerInfo{*bootinfo})

	assert.NoError(t, host2.dht.Provide(ctx, point, true))
	assert.NoError(t, host3.dht.Provide(ctx, point, true))
	assert.NoError(t, host4.dht.Provide(ctx, point, true))

	host2.subscribe(t, topic)
	host3.subscribe(t, topic)
	host4.subscribe(t, topic)

	for _, peer := range host2.Peerstore().Peers() {
		log.Println(peer.Pretty(), host2.Peerstore().PeerInfo(peer).Addrs)
	}

	host2.sub.Publish(topic, []byte("Hello"))

	host3.Close()
	host2.sub.Publish(topic, []byte("ReHi!"))
	time.Sleep(20 * time.Second)
	host2.sub.Publish(topic, []byte("ReHi!"))

	for _, peer := range host2.Peerstore().Peers() {
		log.Println(peer.Pretty(), host2.Peerstore().PeerInfo(peer).Addrs)
	}
}
