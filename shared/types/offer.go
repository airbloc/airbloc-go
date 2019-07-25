package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Offer struct {
	Provider string         "json:\"provider\""
	Consumer common.Address "json:\"consumer\""
	DataIds  []DataId       "json:\"dataIds\""
	At       *big.Int       "json:\"at\""
	Until    *big.Int       "json:\"until\""
	Escrow   struct {
		Addr common.Address "json:\"addr\""
		Sign [4]byte        "json:\"sign\""
		Args []byte         "json:\"args\""
	} "json:\"escrow\""
	Status uint8 "json:\"status\""
}
