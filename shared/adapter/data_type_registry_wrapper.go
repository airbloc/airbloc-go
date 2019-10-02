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

//go:generate mockgen -source data_type_registry_wrapper.go -destination ./mocks/mock_data_type_registry.go -package mocks IDataTypeRegistryManager,IDataTypeRegistryContract
type IDataTypeRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IDataTypeRegistryCalls

	// Transact methods
	Register(
		ctx context.Context,
		opts *blockchain.TransactOpts,
		name string,
		schemaHash common.Hash,
	) error

	Unregister(
		ctx context.Context,
		opts *blockchain.TransactOpts,
		name string,
	) error

	// Event methods
	IDataTypeRegistryFilterer
	IDataTypeRegistryWatcher
}

type IDataTypeRegistryCalls interface {
	Exists(
		name string,
	) (
		bool,
		error,
	)
	Get(
		name string,
	) (
		types.DataType,
		error,
	)
	IsOwner(
		name string,
		owner common.Address,
	) (
		bool,
		error,
	)
}

type IDataTypeRegistryTransacts interface {
	Register(
		ctx context.Context,
		opts *blockchain.TransactOpts,
		name string,
		schemaHash common.Hash,
	) (*chainTypes.Receipt, error)
	Unregister(
		ctx context.Context,
		opts *blockchain.TransactOpts,
		name string,
	) (*chainTypes.Receipt, error)
}

type IDataTypeRegistryEvents interface {
	IDataTypeRegistryFilterer
	IDataTypeRegistryParser
	IDataTypeRegistryWatcher
}

type IDataTypeRegistryFilterer interface {
	FilterRegistration(
		opts *bind.FilterOpts,

	) (*DataTypeRegistryRegistrationIterator, error)
	FilterUnregistration(
		opts *bind.FilterOpts,

	) (*DataTypeRegistryUnregistrationIterator, error)
}

type IDataTypeRegistryParser interface {
	ParseRegistration(log chainTypes.Log) (*DataTypeRegistryRegistration, error)
	ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*DataTypeRegistryRegistration, error)
	ParseUnregistration(log chainTypes.Log) (*DataTypeRegistryUnregistration, error)
	ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*DataTypeRegistryUnregistration, error)
}

type IDataTypeRegistryWatcher interface {
	WatchRegistration(
		opts *bind.WatchOpts,
		sink chan<- *DataTypeRegistryRegistration,

	) (event.Subscription, error)
	WatchUnregistration(
		opts *bind.WatchOpts,
		sink chan<- *DataTypeRegistryUnregistration,

	) (event.Subscription, error)
}

type IDataTypeRegistryContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IDataTypeRegistryCalls
	IDataTypeRegistryTransacts
	IDataTypeRegistryEvents
}

// Manager is contract wrapper struct
type DataTypeRegistryContract struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int
	client    blockchain.TxClient

	DataTypeRegistryCaller
	DataTypeRegistryFilterer
	DataTypeRegistryTransactor
}

// Address is getter method of Accounts.address
func (c *DataTypeRegistryContract) Address() common.Address {
	return c.address
}

// TxHash is getter method of Accounts.txHash
func (c *DataTypeRegistryContract) TxHash() common.Hash {
	return c.txHash
}

// CreatedAt is getter method of Accounts.createdAt
func (c *DataTypeRegistryContract) CreatedAt() *big.Int {
	return c.createdAt
}

func newDataTypeRegistryContract(address common.Address, txHash common.Hash, createdAt *big.Int, parsedABI abi.ABI, backend bind.ContractBackend) interface{} {
	contract := blockchain.NewBoundContract(address, parsedABI, backend, backend, backend)

	return &DataTypeRegistryContract{
		address:   address,
		txHash:    txHash,
		createdAt: createdAt,
		client:    backend.(blockchain.TxClient),

		DataTypeRegistryCaller:     DataTypeRegistryCaller{contract: contract},
		DataTypeRegistryTransactor: DataTypeRegistryTransactor{contract: contract},
		DataTypeRegistryFilterer:   DataTypeRegistryFilterer{contract: contract},
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("DataTypeRegistry", newDataTypeRegistryContract)
	blockchain.RegisterSelector("0x656afdee", "register(string,bytes32)")
	blockchain.RegisterSelector("0x6598a1ae", "unregister(string)")
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (c *DataTypeRegistryContract) Exists(
	name string,
) (

	bool,
	error,
) {
	return c.DataTypeRegistryCaller.Exists(nil, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (c *DataTypeRegistryContract) Get(
	name string,
) (

	types.DataType,
	error,
) {
	return c.DataTypeRegistryCaller.Get(nil, name)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (c *DataTypeRegistryContract) IsOwner(
	name string,
	owner common.Address,
) (

	bool,
	error,
) {
	return c.DataTypeRegistryCaller.IsOwner(nil, name, owner)
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (c *DataTypeRegistryContract) Register(
	ctx context.Context,
	opts *blockchain.TransactOpts,
	name string,
	schemaHash common.Hash,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &blockchain.TransactOpts{TxType: chainTypes.TxTypeSmartContractExecution}
	}

	tx, err := c.DataTypeRegistryTransactor.Register(c.client.Account(ctx, opts), name, schemaHash)

	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (c *DataTypeRegistryContract) Unregister(
	ctx context.Context,
	opts *blockchain.TransactOpts,
	name string,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &blockchain.TransactOpts{TxType: chainTypes.TxTypeSmartContractExecution}
	}

	tx, err := c.DataTypeRegistryTransactor.Unregister(c.client.Account(ctx, opts), name)

	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
