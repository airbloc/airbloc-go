// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"

	blockchain "github.com/airbloc/airbloc-go/shared/blockchain"
	types "github.com/airbloc/airbloc-go/shared/types"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source controller_registry_wrapper.go -destination ./mocks/mock_controller_registry.go -package mocks IControllerRegistryManager,IControllerRegistryContract
type IControllerRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IControllerRegistryCalls

	// Transact methods
	Register(
		ctx context.Context,
		opts *blockchain.TransactOpts,
		controllerAddr common.Address,
	) error

	// Event methods
	IControllerRegistryFilterer
	IControllerRegistryWatcher
}

type IControllerRegistryCalls interface {
	Exists(
		controller common.Address,
	) (
		bool,
		error,
	)
	Get(
		controller common.Address,
	) (
		types.DataController,
		error,
	)
}

type IControllerRegistryTransacts interface {
	Register(
		ctx context.Context,
		opts *blockchain.TransactOpts,
		controllerAddr common.Address,
	) (*chainTypes.Receipt, error)
}

type IControllerRegistryEvents interface {
	IControllerRegistryFilterer
	IControllerRegistryParser
	IControllerRegistryWatcher
}

type IControllerRegistryFilterer interface {
	FilterOwnershipTransferred(
		opts *bind.FilterOpts,
		previousOwner []common.Address,
		newOwner []common.Address,
	) (*ControllerRegistryOwnershipTransferredIterator, error)
	FilterRegistration(
		opts *bind.FilterOpts,
		controller []common.Address,
	) (*ControllerRegistryRegistrationIterator, error)
	FilterUnregistration(
		opts *bind.FilterOpts,
		controller []common.Address,
	) (*ControllerRegistryUnregistrationIterator, error)
}

type IControllerRegistryParser interface {
	ParseOwnershipTransferred(log chainTypes.Log) (*ControllerRegistryOwnershipTransferred, error)
	ParseOwnershipTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryOwnershipTransferred, error)
	ParseRegistration(log chainTypes.Log) (*ControllerRegistryRegistration, error)
	ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryRegistration, error)
	ParseUnregistration(log chainTypes.Log) (*ControllerRegistryUnregistration, error)
	ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryUnregistration, error)
}

type IControllerRegistryWatcher interface {
	WatchOwnershipTransferred(
		opts *bind.WatchOpts,
		sink chan<- *ControllerRegistryOwnershipTransferred,
		previousOwner []common.Address,
		newOwner []common.Address,
	) (event.Subscription, error)
	WatchRegistration(
		opts *bind.WatchOpts,
		sink chan<- *ControllerRegistryRegistration,
		controller []common.Address,
	) (event.Subscription, error)
	WatchUnregistration(
		opts *bind.WatchOpts,
		sink chan<- *ControllerRegistryUnregistration,
		controller []common.Address,
	) (event.Subscription, error)
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
	contract := blockchain.NewBoundContract(address, parsedABI, backend, backend, backend)

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
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (c *ControllerRegistryContract) Exists(
	controller common.Address,
) (

	bool,
	error,
) {
	return c.ControllerRegistryCaller.Exists(nil, controller)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (c *ControllerRegistryContract) Get(
	controller common.Address,
) (

	types.DataController,
	error,
) {
	return c.ControllerRegistryCaller.Get(nil, controller)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (c *ControllerRegistryContract) Register(
	ctx context.Context,
	opts *blockchain.TransactOpts,
	controllerAddr common.Address,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &blockchain.TransactOpts{TxType: chainTypes.TxTypeSmartContractExecution}
	}

	tx, err := c.ControllerRegistryTransactor.Register(c.client.Account(ctx, opts), controllerAddr)

	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
