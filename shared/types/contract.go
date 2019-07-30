package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// AccountStatus is bind of Accounts.Status
type AccountStatus uint8

const (
	// AccountStatusNone is bind of Accounts.Status.None
	AccountStatusNone = AccountStatus(iota)
	// AccountStatusTemporary is bind of Accounts.Status.Temporary
	AccountStatusTemporary
	// AccountStatusCreated is bind of Accounts.Status.Created
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
	Name  string         "json:\"name\""
	Owner common.Address "json:\"owner\""
	Addr  common.Address "json:\"hashedName\""
}

// ConsentActionTypes is bind of Consents.ActionTypes
type ConsentActionTypes uint8

const (
	// ConsentActionCollection is bind of Consents.ActionTypes.Collectionn
	ConsentActionCollection = ConsentActionTypes(iota)
	// ConsentActionExchange is bind of Consents.ActionTypes.Exchange
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
