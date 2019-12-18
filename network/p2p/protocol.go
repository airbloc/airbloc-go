package p2p

import (
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/cipher/aead"
	"github.com/perlin-network/noise/handshake/ecdh"
	"github.com/perlin-network/noise/protocol"
	"github.com/perlin-network/noise/skademlia"
)

const ECDHHandshakeMessage = ".airbloc_node_handshake"

func enforceProtocol(node *noise.Node) {
	ecdhBlock := ecdh.New()
	ecdhBlock.WithHandshakeMessage(ECDHHandshakeMessage)

	p := protocol.New()
	p.Register(ecdhBlock)
	p.Register(aead.New())
	p.Register(skademlia.New())
	p.Enforce(node)
}
