package account

import (
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/key"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type Status int8

const (
	StatusNone = Status(iota)
	StatusTemporary
	StatusCreated
)

type Account struct {
	ID          ablCommon.ID
	Status      Status
	Owner       ethCommon.Address
	Proxy       ethCommon.Address
	PasswordSig []byte

	k *key.Key
}
