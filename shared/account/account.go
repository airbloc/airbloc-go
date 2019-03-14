package account

import (
	"github.com/airbloc/airbloc-go/shared/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type Status uint8

const (
	StatusNone = Status(iota)
	StatusTemporary
	StatusCreated
)

type Account struct {
	ID            types.ID
	Status        Status
	Owner         ethCommon.Address
	Proxy         ethCommon.Address
	Delegate      ethCommon.Address
	PasswordProof ethCommon.Address
}
