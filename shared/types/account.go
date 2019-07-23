package types

import "github.com/ethereum/go-ethereum/common"

type Account struct {
	Owner         common.Address
	Status        uint8
	Controller    common.Address
	PasswordProof common.Address
}
