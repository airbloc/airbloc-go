package types

import (
	"math/big"

	"github.com/klaytn/klaytn/common"
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

// ConsentData is bind of Consents.ConsentData
type ConsentData struct {
	Action   uint8
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
