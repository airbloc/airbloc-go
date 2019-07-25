package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type DataController struct {
	Controller common.Address "json:\"controller\""
	UsersCount *big.Int       "json:\"usersCount\""
}
