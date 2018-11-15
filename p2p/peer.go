package p2p

import (
	"time"

	"log"
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

func (s *AirblocServer) peerWorker() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	s.refreshPeer()
	for {
		select {
		case <-ticker.C:
			s.refreshPeer()
		}
	}
}

func (s *AirblocServer) refreshPeer() {
	s.clearPeer()
	s.updatePeer()
}

func (s *AirblocServer) clearPeer() {
	peerStore := s.host.Peerstore()
	for _, peerID := range peerStore.PeersWithAddrs() {
		peerStore.ClearAddrs(peerID)
	}
}

func (s *AirblocServer) updatePeer() {
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
		info, err := s.dht.FindPeer(s.ctx, id)
		if err != nil {
			log.Println("failed to find peer", id.Pretty(), ":", err)
			return
		}
		s.host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.TempAddrTTL)
	}
}
