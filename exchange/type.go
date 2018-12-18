package exchange

import (
	"github.com/ethereum/go-ethereum/common"
)

type Escrow struct {
	Addr      common.Address
	OpenSign  [4]byte
	OpenArgs  []byte
	CloseSign [4]byte
	CloseArgs []byte
}

type Offer struct {
	From     common.Address
	To       common.Address
	DataIds  [][20]byte
	Escrow   *Escrow
	Status   uint8
	Reverted bool
}

type OfferCompact struct {
	From     common.Address
	To       common.Address
	Escrow   common.Address
	Reverted bool
}
