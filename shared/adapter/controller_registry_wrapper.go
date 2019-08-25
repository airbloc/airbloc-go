// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source controller_registry_wrapper.go -destination ./mocks/mock_controller_registry.go -package mocks IControllerRegistryManager,IControllerRegistryContract
type IControllerRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IControllerRegistryCalls

	// Transact methods
	Register(ctx context.Context, controllerAddr common.Address) error
	RenounceOwnership(ctx context.Context) error
	TransferOwnership(ctx context.Context, newOwner common.Address) error

	// Event methods
	IControllerRegistryFilterer
	IControllerRegistryWatcher
}

type IControllerRegistryCalls interface {
	Exists(controller common.Address) (bool, error)
	Get(controller common.Address) (types.DataController, error)
	IsOwner() (bool, error)
	Owner() (common.Address, error)
}

type IControllerRegistryTransacts interface {
	Register(ctx context.Context, controllerAddr common.Address) (*klayTypes.Receipt, error)
	RenounceOwnership(ctx context.Context) (*klayTypes.Receipt, error)
	TransferOwnership(ctx context.Context, newOwner common.Address) (*klayTypes.Receipt, error)
}

type IControllerRegistryEvents interface {
	IControllerRegistryFilterer
	IControllerRegistryParser
	IControllerRegistryWatcher
}

type IControllerRegistryFilterer interface {
	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerRegistryOwnershipTransferredIterator, error)
	FilterRegistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryRegistrationIterator, error)
	FilterUnregistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryUnregistrationIterator, error)
}

type IControllerRegistryParser interface {
	ParseOwnershipTransferredFromReceipt(receipt *klayTypes.Receipt) (*ControllerRegistryOwnershipTransferred, error)
	ParseRegistrationFromReceipt(receipt *klayTypes.Receipt) (*ControllerRegistryRegistration, error)
	ParseUnregistrationFromReceipt(receipt *klayTypes.Receipt) (*ControllerRegistryUnregistration, error)
}

type IControllerRegistryWatcher interface {
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	WatchRegistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryRegistration, controller []common.Address) (event.Subscription, error)
	WatchUnregistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryUnregistration, controller []common.Address) (event.Subscription, error)
}

type IControllerRegistryContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IControllerRegistryCalls
	IControllerRegistryTransacts
	IControllerRegistryEvents
}

// Manager is contract wrapper struct
type ControllerRegistryContract struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int
	client    blockchain.TxClient

	ControllerRegistryCaller
	ControllerRegistryFilterer
	ControllerRegistryTransactor
}

// Address is getter method of Accounts.address
func (c *ControllerRegistryContract) Address() common.Address {
	return c.address
}

// TxHash is getter method of Accounts.txHash
func (c *ControllerRegistryContract) TxHash() common.Hash {
	return c.txHash
}

// CreatedAt is getter method of Accounts.createdAt
func (c *ControllerRegistryContract) CreatedAt() *big.Int {
	return c.createdAt
}

func newControllerRegistryContract(address common.Address, txHash common.Hash, createdAt *big.Int, parsedABI abi.ABI, backend bind.ContractBackend) interface{} {
	contract := bind.NewBoundContract(address, parsedABI, backend, backend, backend)

	return &ControllerRegistryContract{
		address:   address,
		txHash:    txHash,
		createdAt: createdAt,
		client:    backend.(blockchain.TxClient),

		ControllerRegistryCaller:     ControllerRegistryCaller{contract: contract},
		ControllerRegistryTransactor: ControllerRegistryTransactor{contract: contract},
		ControllerRegistryFilterer:   ControllerRegistryFilterer{contract: contract},
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("ControllerRegistry", newControllerRegistryContract)
	blockchain.RegisterSelector("0x4420e486", "register(address)")
	blockchain.RegisterSelector("0x715018a6", "renounceOwnership()")
	blockchain.RegisterSelector("0xf2fde38b", "transferOwnership(address)")
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (c *ControllerRegistryContract) Exists(controller common.Address) (bool, error) {
	return c.ControllerRegistryCaller.Exists(nil, controller)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (c *ControllerRegistryContract) Get(controller common.Address) (types.DataController, error) {
	return c.ControllerRegistryCaller.Get(nil, controller)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (c *ControllerRegistryContract) IsOwner() (bool, error) {
	return c.ControllerRegistryCaller.IsOwner(nil)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (c *ControllerRegistryContract) Owner() (common.Address, error) {
	return c.ControllerRegistryCaller.Owner(nil)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (c *ControllerRegistryContract) Register(ctx context.Context, controllerAddr common.Address) (*klayTypes.Receipt, error) {
	tx, err := c.ControllerRegistryTransactor.Register(c.client.Account(), controllerAddr)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (c *ControllerRegistryContract) RenounceOwnership(ctx context.Context) (*klayTypes.Receipt, error) {
	tx, err := c.ControllerRegistryTransactor.RenounceOwnership(c.client.Account())
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (c *ControllerRegistryContract) TransferOwnership(ctx context.Context, newOwner common.Address) (*klayTypes.Receipt, error) {
	tx, err := c.ControllerRegistryTransactor.TransferOwnership(c.client.Account(), newOwner)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
