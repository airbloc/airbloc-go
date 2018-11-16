package account

import (
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type Status uint8

const (
	StatusNone = Status(iota)
	StatusTemporary
	StatusCreated
)

type Account struct {
	ID            ablCommon.ID
	Status        Status
	Owner         ethCommon.Address
	Proxy         ethCommon.Address
	PasswordProof ethCommon.Address
}
