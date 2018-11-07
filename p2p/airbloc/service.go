package airbloc

import (
	"context"
	"sync"

	"github.com/airbloc/airbloc-go/key"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	mutex  *sync.Mutex
	host   host.Host
	gsub   *pubsub.PubSub
}

func NewServer(identity *key.Key, listenAddr multiaddr.Multiaddr) (*Server, error) {
	privKey, err := identity.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(listenAddr),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	dht, err := kaddht.New(ctx, h)
	if err != nil {
		cancel()
		return nil, err
	}
	h = routedhost.Wrap(h, dht)

	gsub, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		cancel()
		return nil, err
	}

	return &Server{
		ctx:    ctx,
		cancel: cancel,
		host:   h,
		gsub:   gsub,
		mutex:  new(sync.Mutex),
	}, nil
}

// api backend interfaces
func (s *Server) Start() error {
	return nil
}

func (s *Server) Stop() {
}

func (s *Server) Regsiter() {

}

func (s *Server) Subscribe() {

}

func (s *Server) Broadcast() {

}

func (s *Server) Emit() {

}
