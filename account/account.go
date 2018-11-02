package account

import (
	"crypto/ecdsa"

	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	Opts *bind.TransactOpts
	key  *ecdsa.PrivateKey
}

func NewAccount(key *ecdsa.PrivateKey) *Account {
	return &Account{
		Opts: bind.NewKeyedTransactor(key),
		key:  key,
	}
}
