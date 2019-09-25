package types

import (
	"math/big"

	"github.com/klaytn/klaytn/common"
)

// TODO: refactor type structure

const (
	// AccountStatusNone is bind of Accounts.Status.None
	AccountStatusNone uint8 = iota
	// AccountStatusTemporary is bind of Accounts.Status.Temporary
	AccountStatusTemporary
	// AccountStatusCreated is bind of Accounts.Status.Created
	AccountStatusCreated
)

// Account is bind of Accounts.Account
type Account struct {
	Owner         common.Address
	Status        uint8
	Controller    common.Address
	PasswordProof common.Address
}

// App is bind of AppRegistry.App
type App struct {
	Name  string
	Owner common.Address
	Addr  common.Address
}

const (
	// ConsentActionCollection is bind of Consents.ActionTypes.Collectionn
	ConsentActionCollection uint8 = iota
	// ConsentActionExchange is bind of Consents.ActionTypes.Exchange
	ConsentActionExchange
)

// ConsentData is bind of Consents.ConsentData
type ConsentData struct {
	Action   uint8  `json:"action" binding:"required"`
	DataType string `json:"data_type" binding:"required"`
	Allow    bool   `json:"allow" binding:"required"`
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

// Escrow is bind of ExchangeLib.Escrow
type Escrow struct {
	Addr common.Address
	Sign [4]byte
	Args []byte
}

// Offer is bind of ExchangeLib.Offer
type Offer struct {
	Provider string
	Consumer common.Address
	DataIds  []DataId
	At       *big.Int
	Until    *big.Int
	Escrow   Escrow
	Status   uint8
}
