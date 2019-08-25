package types

import (
	"math/big"

	"github.com/klaytn/klaytn/common"
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

// AccountStatusList is list of AccountStatus
var AccountStatusList = map[uint8]AccountStatus{
	uint8(AccountStatusNone):      AccountStatusNone,
	uint8(AccountStatusTemporary): AccountStatusTemporary,
	uint8(AccountStatusCreated):   AccountStatusCreated,
}

// Account is bind of Accounts.Account
type Account struct {
	Owner         common.Address
	Status        AccountStatus
	Controller    common.Address
	PasswordProof common.Address
}

// App is bind of AppRegistry.App
type App struct {
	Name  string
	Owner common.Address
	Addr  common.Address
}

// ConsentActionTypes is bind of Consents.ActionTypes
type ConsentActionTypes uint8

const (
	// ConsentActionCollection is bind of Consents.ActionTypes.Collectionn
	ConsentActionCollection = ConsentActionTypes(iota)
	// ConsentActionExchange is bind of Consents.ActionTypes.Exchange
	ConsentActionExchange
)

// ConsentActionList is list of ConsentAction
var ConsentActionList = map[uint8]ConsentActionTypes{
	uint8(ConsentActionCollection): ConsentActionCollection,
	uint8(ConsentActionExchange):   ConsentActionExchange,
}

// ConsentData is bind of Consents.ConsentData
type ConsentData struct {
	Action   ConsentActionTypes
	DataType string
	Allow    bool
}

// DataController is bind of ControllerRegistry.DataController
type DataController struct {
	Controller common.Address
	UsersCount *big.Int
}

// DataType is bind of DataTypeRegistry.DataType
type DataType struct {
	Name       string
	Owner      common.Address
	SchemaHash common.Hash
}

// Offer is bind of ExchangeLib.Offer
type Offer struct {
	Provider string
	Consumer common.Address
	DataIds  []DataId
	At       *big.Int
	Until    *big.Int
	Escrow   struct {
		Addr common.Address
		Sign [4]byte
		Args []byte
	}
	Status uint8
}
