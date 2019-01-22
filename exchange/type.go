package exchange

import (
	"github.com/ethereum/go-ethereum/common"
)

type Escrow struct {
	Addr common.Address
	Sign [4]byte
	Args []byte
}

type Offer struct {
	From    common.Address
	To      common.Address
	DataIds [][32]byte
	Escrow  *Escrow
	Status  uint8
}

type OfferCompact struct {
	From   common.Address
	To     common.Address
	Escrow common.Address
}
