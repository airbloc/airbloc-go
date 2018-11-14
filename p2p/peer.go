package p2p

import (
	"log"
	"time"

	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
)

func (s *Server) peer(id peer.ID) (peerstore.PeerInfo, error) {
	oinfo, err := s.dht.FindPeer(s.ctx, id)
	if err != nil {
		return peerstore.PeerInfo{}, err
	}

	pinfo := s.host.Peerstore().PeerInfo(id)
	for _, oaddr := range oinfo.Addrs {
		exists := false
		for _, paddr := range pinfo.Addrs {
			if oaddr.Equal(paddr) {
				exists = true
			}
		}
		if !exists {
			s.host.Peerstore().AddAddr(id, oaddr, peerstore.TempAddrTTL)
		}
	}

	return s.host.Peerstore().PeerInfo(id), nil
}

func (s *Server) peerWorker() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			idch, err := s.dht.GetClosestPeers(s.ctx, s.id.KeyString())
			if s.ctx.Err() != nil {
				log.Println("context error:", err)
				return
			}

			if err != nil {
				log.Println("failed to get closest peers:", err)
				return
			}

			for id := range idch {
				err = s.host.Connect(s.ctx, peerstore.PeerInfo{ID: id})
				if err != nil {
					log.Println("failed to connect peer:", err)
					return
				}
			}
		}
	}
}
