package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Offer struct {
	Provider string
	Consumer common.Address
	DataIds  [][20]byte
	At       *big.Int
	Until    *big.Int
	Escrow   struct {
		Addr common.Address
		Sign [4]byte
		Args []byte
	}
	Status uint8
}
