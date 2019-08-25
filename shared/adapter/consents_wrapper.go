// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source consents_wrapper.go -destination ./mocks/mock_consents.go -package mocks IConsentsManager,IConsentsContract
type IConsentsManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IConsentsCalls

	// Transact methods
	Consent(ctx context.Context, appName string, consentData types.ConsentData) error
	ConsentByController(ctx context.Context, userId types.ID, appName string, consentData types.ConsentData) error
	ConsentMany(ctx context.Context, appName string, consentData []types.ConsentData) error
	ConsentManyByController(ctx context.Context, userId types.ID, appName string, consentData []types.ConsentData) error
	ModifyConsentByController(ctx context.Context, userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) error

	// Event methods
	IConsentsFilterer
	IConsentsWatcher
}

type IConsentsCalls interface {
	IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error)
	IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error)
}

type IConsentsTransacts interface {
	Consent(ctx context.Context, appName string, consentData types.ConsentData) (*klayTypes.Receipt, error)
	ConsentByController(ctx context.Context, userId types.ID, appName string, consentData types.ConsentData) (*klayTypes.Receipt, error)
	ConsentMany(ctx context.Context, appName string, consentData []types.ConsentData) (*klayTypes.Receipt, error)
	ConsentManyByController(ctx context.Context, userId types.ID, appName string, consentData []types.ConsentData) (*klayTypes.Receipt, error)
	ModifyConsentByController(ctx context.Context, userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) (*klayTypes.Receipt, error)
}

type IConsentsEvents interface {
	IConsentsFilterer
	IConsentsParser
	IConsentsWatcher
}

type IConsentsFilterer interface {
	FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appAddr []common.Address) (*ConsentsConsentedIterator, error)
}

type IConsentsParser interface {
	ParseConsentedFromReceipt(receipt *klayTypes.Receipt) (*ConsentsConsented, error)
}

type IConsentsWatcher interface {
	WatchConsented(opts *bind.WatchOpts, sink chan<- *ConsentsConsented, action []uint8, userId []types.ID, appAddr []common.Address) (event.Subscription, error)
}

type IConsentsContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IConsentsCalls
	IConsentsTransacts
	IConsentsEvents
}

// Manager is contract wrapper struct
type ConsentsContract struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int
	client    blockchain.TxClient

	ConsentsCaller
	ConsentsFilterer
	ConsentsTransactor
}

// Address is getter method of Accounts.address
func (c *ConsentsContract) Address() common.Address {
	return c.address
}

// TxHash is getter method of Accounts.txHash
func (c *ConsentsContract) TxHash() common.Hash {
	return c.txHash
}

// CreatedAt is getter method of Accounts.createdAt
func (c *ConsentsContract) CreatedAt() *big.Int {
	return c.createdAt
}

func newConsentsContract(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (*ConsentsContract, error) {
	contract, err := newConsents(address, txHash, createdAt, backend)
	if err != nil {
		return nil, err
	}

	return &ConsentsContract{
		address:   address,
		txHash:    txHash,
		createdAt: createdAt,
		client:    backend.(blockchain.TxClient),

		ConsentsCaller:     contract.ConsentsCaller,
		ConsentsFilterer:   contract.ConsentsFilterer,
		ConsentsTransactor: contract.ConsentsTransactor,
	}, nil
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("Consents", (&Consents{}).new)
	blockchain.RegisterSelector("0xcd4dc804", "consent(string,(uint8,string,bool))")
	blockchain.RegisterSelector("0xf573f89a", "consentByController(bytes8,string,(uint8,string,bool))")
	blockchain.RegisterSelector("0xdd43ad05", "consentMany(string,(uint8,string,bool)[])")
	blockchain.RegisterSelector("0xae6d5034", "consentManyByController(bytes8,string,(uint8,string,bool)[])")
	blockchain.RegisterSelector("0x0bfec389", "modifyConsentByController(bytes8,string,(uint8,string,bool),bytes)")
}

// IsAllowed is a free data retrieval call binding the contract method 0x50615985.
//
// Solidity: function isAllowed(bytes8 userId, string appName, uint8 action, string dataType) constant returns(bool)
func (c *ConsentsContract) IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	return c.ConsentsCaller.IsAllowed(nil, userId, appName, action, dataType)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x7cdda67c.
//
// Solidity: function isAllowedAt(bytes8 userId, string appName, uint8 action, string dataType, uint256 blockNumber) constant returns(bool)
func (c *ConsentsContract) IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	return c.ConsentsCaller.IsAllowedAt(nil, userId, appName, action, dataType, blockNumber)
}

// Consent is a paid mutator transaction binding the contract method 0xcd4dc804.
//
// Solidity: function consent(string appName, (uint8,string,bool) consentData) returns()
func (c *ConsentsContract) Consent(ctx context.Context, appName string, consentData types.ConsentData) (*klayTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.Consent(c.client.Account(), appName, consentData)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf573f89a.
//
// Solidity: function consentByController(bytes8 userId, string appName, (uint8,string,bool) consentData) returns()
func (c *ConsentsContract) ConsentByController(ctx context.Context, userId types.ID, appName string, consentData types.ConsentData) (*klayTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ConsentByController(c.client.Account(), userId, appName, consentData)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, (uint8,string,bool)[] consentData) returns()
func (c *ConsentsContract) ConsentMany(ctx context.Context, appName string, consentData []types.ConsentData) (*klayTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ConsentMany(c.client.Account(), appName, consentData)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ConsentManyByController is a paid mutator transaction binding the contract method 0xae6d5034.
//
// Solidity: function consentManyByController(bytes8 userId, string appName, (uint8,string,bool)[] consentData) returns()
func (c *ConsentsContract) ConsentManyByController(ctx context.Context, userId types.ID, appName string, consentData []types.ConsentData) (*klayTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ConsentManyByController(c.client.Account(), userId, appName, consentData)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0x0bfec389.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, (uint8,string,bool) consentData, bytes passwordSignature) returns()
func (c *ConsentsContract) ModifyConsentByController(ctx context.Context, userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) (*klayTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ModifyConsentByController(c.client.Account(), userId, appName, consentData, passwordSignature)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
