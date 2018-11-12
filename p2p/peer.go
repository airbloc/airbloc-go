package p2p

import (
	"log"
	"time"

	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

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
