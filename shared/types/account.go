package types

import "github.com/ethereum/go-ethereum/common"

type Account struct {
	Owner         common.Address "json:\"Owner\""
	Status        uint8          "json:\"Status\""
	Controller    common.Address "json:\"Controller\""
	PasswordProof common.Address "json:\"PasswordProof\""
}
