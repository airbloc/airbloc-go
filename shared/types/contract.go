package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AccountStatus uint8

const (
	AccountStatusNone = AccountStatus(iota)
	AccountStatusTemporary
	AccountStatusCreated
)

// Account is bind of Accounts.Account
type Account struct {
	Owner         common.Address "json:\"Owner\""
	Status        AccountStatus  "json:\"Status\""
	Controller    common.Address "json:\"Controller\""
	PasswordProof common.Address "json:\"PasswordProof\""
}

// App is bind of AppRegistry.App
type App struct {
	Name       string         "json:\"name\""
	Owner      common.Address "json:\"owner\""
	HashedName common.Hash    "json:\"hashedName\""
}

type ConsentActionTypes uint8

const (
	ConsentActionCollection = ConsentActionTypes(iota)
	ConsentActionExchange
)

// DataController is bind of ControllerRegistry.DataController
type DataController struct {
	Controller common.Address "json:\"controller\""
	UsersCount *big.Int       "json:\"usersCount\""
}

// DataType is bind of DataTypeRegistry.DataType
type DataType struct {
	Name       string         "json:\"name\""
	Owner      common.Address "json:\"owner\""
	SchemaHash common.Hash    "json:\"schemaHash\""
}

// Offer is bind of ExchangeLib.Offer
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
