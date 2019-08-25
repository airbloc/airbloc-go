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

//go:generate mockgen -source app_registry_wrapper.go -destination ./mocks/mock_app_registry.go -package mocks IAppRegistryManager,IAppRegistryContract
type IAppRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IAppRegistryCalls

	// Transact methods
	Register(ctx context.Context, appName string) error
	TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) error
	Unregister(ctx context.Context, appName string) error

	// Event methods
	IAppRegistryFilterer
	IAppRegistryWatcher
}

type IAppRegistryCalls interface {
	Exists(appName string) (bool, error)
	Get(appName string) (types.App, error)
	IsOwner(appName string, owner common.Address) (bool, error)
}

type IAppRegistryTransacts interface {
	Register(ctx context.Context, appName string) (*klayTypes.Receipt, error)
	TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) (*klayTypes.Receipt, error)
	Unregister(ctx context.Context, appName string) (*klayTypes.Receipt, error)
}

type IAppRegistryEvents interface {
	IAppRegistryFilterer
	IAppRegistryParser
	IAppRegistryWatcher
}

type IAppRegistryFilterer interface {
	FilterAppOwnerTransferred(opts *bind.FilterOpts, appAddr []common.Address, oldOwner []common.Address) (*AppRegistryAppOwnerTransferredIterator, error)
	FilterRegistration(opts *bind.FilterOpts, appAddr []common.Address) (*AppRegistryRegistrationIterator, error)
	FilterUnregistration(opts *bind.FilterOpts, appAddr []common.Address) (*AppRegistryUnregistrationIterator, error)
}

type IAppRegistryParser interface {
	ParseAppOwnerTransferredFromReceipt(receipt *klayTypes.Receipt) (*AppRegistryAppOwnerTransferred, error)
	ParseRegistrationFromReceipt(receipt *klayTypes.Receipt) (*AppRegistryRegistration, error)
	ParseUnregistrationFromReceipt(receipt *klayTypes.Receipt) (*AppRegistryUnregistration, error)
}

type IAppRegistryWatcher interface {
	WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AppRegistryAppOwnerTransferred, appAddr []common.Address, oldOwner []common.Address) (event.Subscription, error)
	WatchRegistration(opts *bind.WatchOpts, sink chan<- *AppRegistryRegistration, appAddr []common.Address) (event.Subscription, error)
	WatchUnregistration(opts *bind.WatchOpts, sink chan<- *AppRegistryUnregistration, appAddr []common.Address) (event.Subscription, error)
}

type IAppRegistryContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IAppRegistryCalls
	IAppRegistryTransacts
	IAppRegistryEvents
}

// Manager is contract wrapper struct
type AppRegistryContract struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int
	client    blockchain.TxClient

	AppRegistryCaller
	AppRegistryFilterer
	AppRegistryTransactor
}

// Address is getter method of Accounts.address
func (c *AppRegistryContract) Address() common.Address {
	return c.address
}

// TxHash is getter method of Accounts.txHash
func (c *AppRegistryContract) TxHash() common.Hash {
	return c.txHash
}

// CreatedAt is getter method of Accounts.createdAt
func (c *AppRegistryContract) CreatedAt() *big.Int {
	return c.createdAt
}

func newAppRegistryContract(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (*AppRegistryContract, error) {
	contract, err := newAppRegistry(address, txHash, createdAt, backend)
	if err != nil {
		return nil, err
	}

	return &AppRegistryContract{
		address:   address,
		txHash:    txHash,
		createdAt: createdAt,
		client:    backend.(blockchain.TxClient),

		AppRegistryCaller:     contract.AppRegistryCaller,
		AppRegistryFilterer:   contract.AppRegistryFilterer,
		AppRegistryTransactor: contract.AppRegistryTransactor,
	}, nil
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("AppRegistry", (&AppRegistry{}).new)
	blockchain.RegisterSelector("0xf2c298be", "register(string)")
	blockchain.RegisterSelector("0x1a9dff9f", "transferAppOwner(string,address)")
	blockchain.RegisterSelector("0x6598a1ae", "unregister(string)")
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (c *AppRegistryContract) Exists(appName string) (bool, error) {
	return c.AppRegistryCaller.Exists(nil, appName)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,address))
func (c *AppRegistryContract) Get(appName string) (types.App, error) {
	return c.AppRegistryCaller.Get(nil, appName)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (c *AppRegistryContract) IsOwner(appName string, owner common.Address) (bool, error) {
	return c.AppRegistryCaller.IsOwner(nil, appName, owner)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (c *AppRegistryContract) Register(ctx context.Context, appName string) (*klayTypes.Receipt, error) {
	tx, err := c.AppRegistryTransactor.Register(c.client.Account(), appName)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (c *AppRegistryContract) TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) (*klayTypes.Receipt, error) {
	tx, err := c.AppRegistryTransactor.TransferAppOwner(c.client.Account(), appName, newOwner)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (c *AppRegistryContract) Unregister(ctx context.Context, appName string) (*klayTypes.Receipt, error) {
	tx, err := c.AppRegistryTransactor.Unregister(c.client.Account(), appName)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
