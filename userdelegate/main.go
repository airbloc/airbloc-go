package userdelegate

import (
	"context"
	"fmt"
	"log"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	"github.com/airbloc/airbloc-go/p2p/common"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/userdelegate"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
)

func newServer(memdb localdb.Database, k *key.Key, port int, bootnode bool, bootinfos ...peerstore.PeerInfo) p2p.Server {
	addr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	server, err := p2p.NewAirblocServer(memdb, k, addr, bootnode, bootinfos)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	return server
}

func Main() {
	memdb := localdb.NewMemDB()
	k, err := key.Generate()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	// setup node
	addr, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/9100")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	info, err := peerstore.InfoFromP2pAddr(addr)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	server := newServer(memdb, k, 9100, false, *info)
	server.Start()
	defer server.Stop()

	server.SubscribeTopic("dauth-allow-deadbeefdeadbeef", &pb.DAuthRequest{}, authHandler(true))
	server.SubscribeTopic("dauth-deny-deadbeefdeadbeef", &pb.DAuthRequest{}, authHandler(false))
}

func authHandler(allow bool) func(server p2p.Server, context context.Context, message common.Message) {
	return func(server p2p.Server, context context.Context, message common.Message) {
	}
}
