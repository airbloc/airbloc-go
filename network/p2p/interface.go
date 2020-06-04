package p2p

import "github.com/perlin-network/noise"

type Router interface {
	RegisterHandler(opcode noise.Opcode, handler HandlerFunc)
}
