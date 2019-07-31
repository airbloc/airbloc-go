// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.NewKeyedTransactor
	_ = types.HexToID
	_ = common.Big1
	_ = ethTypes.BloomLookup
	_ = event.NewSubscription
)

// ControllerRegistryABI is the input ABI used to generate the binding from.
const ControllerRegistryABI = "{\"Constructor\":{\"Name\":\"\",\"Const\":false,\"Inputs\":null,\"Outputs\":null},\"Methods\":{\"exists\":{\"Name\":\"exists\",\"Const\":true,\"Inputs\":[{\"Name\":\"controller\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"get\":{\"Name\":\"get\",\"Const\":true,\"Inputs\":[{\"Name\":\"controller\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":25,\"Type\":{},\"Size\":0,\"T\":6,\"TupleElems\":[{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":22,\"Type\":{},\"Size\":256,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null}],\"TupleRawNames\":[\"controller\",\"usersCount\"]},\"Indexed\":false}]},\"isOwner\":{\"Name\":\"isOwner\",\"Const\":true,\"Inputs\":[],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"owner\":{\"Name\":\"owner\",\"Const\":true,\"Inputs\":[],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"register\":{\"Name\":\"register\",\"Const\":false,\"Inputs\":[{\"Name\":\"controllerAddr\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"renounceOwnership\":{\"Name\":\"renounceOwnership\",\"Const\":false,\"Inputs\":[],\"Outputs\":[]},\"transferOwnership\":{\"Name\":\"transferOwnership\",\"Const\":false,\"Inputs\":[{\"Name\":\"newOwner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]}},\"Events\":{\"OwnershipTransferred\":{\"Name\":\"OwnershipTransferred\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"previousOwner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"newOwner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true}]},\"Registration\":{\"Name\":\"Registration\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"controller\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true}]},\"Unregistration\":{\"Name\":\"Unregistration\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"controller\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true}]}}}"

// ControllerRegistry is an auto generated Go binding around an Ethereum contract.
type ControllerRegistry struct {
	address                      common.Address
	txHash                       common.Hash
	createdAt                    *big.Int
	ControllerRegistryCaller     // Read-only binding to the contract
	ControllerRegistryTransactor // Write-only binding to the contract
	ControllerRegistryFilterer   // Log filterer for contract events
}

// Address is getter method of ControllerRegistry.address
func (_ControllerRegistry *ControllerRegistry) Address() common.Address {
	return _ControllerRegistry.address
}

// TxHash is getter method of ControllerRegistry.txHash
func (_ControllerRegistry *ControllerRegistry) TxHash() common.Hash {
	return _ControllerRegistry.txHash
}

// CreatedAt is getter method of ControllerRegistry.createdAt
func (_ControllerRegistry *ControllerRegistry) CreatedAt() *big.Int {
	return _ControllerRegistry.createdAt
}

// ControllerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerRegistrySession struct {
	Contract     *ControllerRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ControllerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRegistryRaw struct {
	Contract *ControllerRegistry // Generic contract binding to access the raw methods on
}

// NewControllerRegistry creates a new instance of ControllerRegistry, bound to a specific deployed contract.
func NewControllerRegistry(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (*ControllerRegistry, error) {
	contract, err := bindControllerRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistry{
		address:                      address,
		txHash:                       txHash,
		createdAt:                    createdAt,
		ControllerRegistryCaller:     ControllerRegistryCaller{contract: contract},
		ControllerRegistryTransactor: ControllerRegistryTransactor{contract: contract},
		ControllerRegistryFilterer:   ControllerRegistryFilterer{contract: contract},
	}, nil
}

// bindControllerRegistry binds a generic wrapper to an already deployed contract.
func bindControllerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ControllerRegistry *ControllerRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ControllerRegistry.Contract.ControllerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ControllerRegistry *ControllerRegistryRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.ControllerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ControllerRegistry *ControllerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.ControllerRegistryTransactor.contract.Transact(opts, method, params...)
}

// ControllerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerRegistryCallerSession struct {
	Contract *ControllerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ControllerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerRegistryCallerRaw struct {
	Contract *ControllerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NewControllerRegistryCaller creates a new read-only instance of ControllerRegistry, bound to a specific deployed contract.
func NewControllerRegistryCaller(address common.Address, caller bind.ContractCaller) (*ControllerRegistryCaller, error) {
	contract, err := bindControllerRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryCaller{contract: contract}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ControllerRegistry *ControllerRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ControllerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// ControllerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerRegistryTransactorSession struct {
	Contract     *ControllerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ControllerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerRegistryTransactorRaw struct {
	Contract *ControllerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewControllerRegistryTransactor creates a new write-only instance of ControllerRegistry, bound to a specific deployed contract.
func NewControllerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerRegistryTransactor, error) {
	contract, err := bindControllerRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ControllerRegistry *ControllerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ControllerRegistry *ControllerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.contract.Transact(opts, method, params...)
}

// ControllerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewControllerRegistryFilterer creates a new log filterer instance of ControllerRegistry, bound to a specific deployed contract.
func NewControllerRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerRegistryFilterer, error) {
	contract, err := bindControllerRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryFilterer{contract: contract}, nil
}

//go:generate mockgen -source controller_registry.go -destination ./mocks/mock_controller_registry.go -package mocks IControllerRegistryManager,IControllerRegistryContract
type IControllerRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	Exists(controller common.Address) (bool, error)
	Get(controller common.Address) (types.DataController, error)
	IsOwner() (bool, error)
	Owner() (common.Address, error)

	// Transact methods
	Register(ctx context.Context, controllerAddr common.Address) error
	RenounceOwnership(ctx context.Context) error
	TransferOwnership(ctx context.Context, newOwner common.Address) error

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerRegistryOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)

	FilterRegistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryRegistrationIterator, error)
	WatchRegistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryRegistration, controller []common.Address) (event.Subscription, error)

	FilterUnregistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryUnregistrationIterator, error)
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
	client   blockchain.TxClient
	contract *ControllerRegistry
	ControllerRegistryFilterer
}

// Address is getter method of ControllerRegistry.address
func (c *ControllerRegistryContract) Address() common.Address {
	return c.contract.Address()
}

// TxHash is getter method of ControllerRegistry.txHash
func (c *ControllerRegistryContract) TxHash() common.Hash {
	return c.contract.TxHash()
}

// CreatedAt is getter method of ControllerRegistry.createdAt
func (c *ControllerRegistryContract) CreatedAt() *big.Int {
	return c.contract.CreatedAt()
}

// NewManager makes new *Manager struct
func NewControllerRegistryContract(client blockchain.TxClient) IControllerRegistryContract {
	contract := client.GetContract(&ControllerRegistry{}).(*ControllerRegistry)
	return &ControllerRegistryContract{
		client:                     client,
		contract:                   contract,
		ControllerRegistryFilterer: contract.ControllerRegistryFilterer,
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("ControllerRegistry", (&ControllerRegistry{}).new)
	blockchain.RegisterSelector("0x4420e486", "register(address)")
	blockchain.RegisterSelector("0x715018a6", "renounceOwnership()")
	blockchain.RegisterSelector("0xf2fde38b", "transferOwnership(address)")
}

func (_ControllerRegistry *ControllerRegistry) new(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (interface{}, error) {
	return NewControllerRegistry(address, txHash, createdAt, backend)
}

type IControllerRegistryCalls interface {
	Exists(controller common.Address) (bool, error)
	Get(controller common.Address) (types.DataController, error)
	IsOwner() (bool, error)
	Owner() (common.Address, error)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (c *ControllerRegistryContract) Exists(controller common.Address) (bool, error) {
	return c.contract.Exists(nil, controller)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (_ControllerRegistry *ControllerRegistryCaller) Exists(opts *bind.CallOpts, controller common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _ControllerRegistry.contract.Call(opts, out, "exists", controller)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (_ControllerRegistry *ControllerRegistrySession) Exists(controller common.Address) (bool, error) {
	return _ControllerRegistry.Contract.Exists(&_ControllerRegistry.CallOpts, controller)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (_ControllerRegistry *ControllerRegistryCallerSession) Exists(controller common.Address) (bool, error) {
	return _ControllerRegistry.Contract.Exists(&_ControllerRegistry.CallOpts, controller)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (c *ControllerRegistryContract) Get(controller common.Address) (types.DataController, error) {
	return c.contract.Get(nil, controller)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (_ControllerRegistry *ControllerRegistryCaller) Get(opts *bind.CallOpts, controller common.Address) (types.DataController, error) {
	ret := new(types.DataController)

	out := ret
	err := _ControllerRegistry.contract.Call(opts, out, "get", controller)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (_ControllerRegistry *ControllerRegistrySession) Get(controller common.Address) (types.DataController, error) {
	return _ControllerRegistry.Contract.Get(&_ControllerRegistry.CallOpts, controller)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (_ControllerRegistry *ControllerRegistryCallerSession) Get(controller common.Address) (types.DataController, error) {
	return _ControllerRegistry.Contract.Get(&_ControllerRegistry.CallOpts, controller)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (c *ControllerRegistryContract) IsOwner() (bool, error) {
	return c.contract.IsOwner(nil)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ControllerRegistry *ControllerRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _ControllerRegistry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ControllerRegistry *ControllerRegistrySession) IsOwner() (bool, error) {
	return _ControllerRegistry.Contract.IsOwner(&_ControllerRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ControllerRegistry *ControllerRegistryCallerSession) IsOwner() (bool, error) {
	return _ControllerRegistry.Contract.IsOwner(&_ControllerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (c *ControllerRegistryContract) Owner() (common.Address, error) {
	return c.contract.Owner(nil)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ControllerRegistry *ControllerRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := &[]interface{}{ret0}
	err := _ControllerRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ControllerRegistry *ControllerRegistrySession) Owner() (common.Address, error) {
	return _ControllerRegistry.Contract.Owner(&_ControllerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ControllerRegistry *ControllerRegistryCallerSession) Owner() (common.Address, error) {
	return _ControllerRegistry.Contract.Owner(&_ControllerRegistry.CallOpts)
}

type IControllerRegistryTransacts interface {
	Register(ctx context.Context, controllerAddr common.Address) (*ethTypes.Receipt, error)
	RenounceOwnership(ctx context.Context) (*ethTypes.Receipt, error)
	TransferOwnership(ctx context.Context, newOwner common.Address) (*ethTypes.Receipt, error)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (c *ControllerRegistryContract) Register(ctx context.Context, controllerAddr common.Address) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Register(c.client.Account(), controllerAddr)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (_ControllerRegistry *ControllerRegistryTransactor) Register(opts *bind.TransactOpts, controllerAddr common.Address) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.contract.Transact(opts, "register", controllerAddr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (_ControllerRegistry *ControllerRegistrySession) Register(controllerAddr common.Address) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.Register(&_ControllerRegistry.TransactOpts, controllerAddr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (_ControllerRegistry *ControllerRegistryTransactorSession) Register(controllerAddr common.Address) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.Register(&_ControllerRegistry.TransactOpts, controllerAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (c *ControllerRegistryContract) RenounceOwnership(ctx context.Context) (*ethTypes.Receipt, error) {
	tx, err := c.contract.RenounceOwnership(c.client.Account())
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ControllerRegistry *ControllerRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ControllerRegistry *ControllerRegistrySession) RenounceOwnership() (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.RenounceOwnership(&_ControllerRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ControllerRegistry *ControllerRegistryTransactorSession) RenounceOwnership() (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.RenounceOwnership(&_ControllerRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (c *ControllerRegistryContract) TransferOwnership(ctx context.Context, newOwner common.Address) (*ethTypes.Receipt, error) {
	tx, err := c.contract.TransferOwnership(c.client.Account(), newOwner)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ControllerRegistry *ControllerRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ControllerRegistry *ControllerRegistrySession) TransferOwnership(newOwner common.Address) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.TransferOwnership(&_ControllerRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ControllerRegistry *ControllerRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*ethTypes.Transaction, error) {
	return _ControllerRegistry.Contract.TransferOwnership(&_ControllerRegistry.TransactOpts, newOwner)
}

type IControllerRegistryEvents interface {
	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerRegistryOwnershipTransferredIterator, error)
	ParseOwnershipTransferredFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryOwnershipTransferred, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)

	FilterRegistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryRegistrationIterator, error)
	ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryRegistration, error)
	WatchRegistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryRegistration, controller []common.Address) (event.Subscription, error)

	FilterUnregistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryUnregistrationIterator, error)
	ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryUnregistration, error)
	WatchUnregistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryUnregistration, controller []common.Address) (event.Subscription, error)
}

// ControllerRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ControllerRegistry contract.
type ControllerRegistryOwnershipTransferredIterator struct {
	Event *ControllerRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ControllerRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ControllerRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ControllerRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ControllerRegistry contract.
type ControllerRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           ethTypes.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryOwnershipTransferredIterator{contract: _ControllerRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// FilterOwnershipTransferred parses the event from given transaction receipt.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (manager *ControllerRegistryContract) ParseOwnershipTransferredFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryOwnershipTransferred, error) {
	return manager.contract.ParseOwnershipTransferredFromReceipt(receipt)
}

// FilterOwnershipTransferred parses the event from given transaction receipt.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseOwnershipTransferredFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryOwnershipTransferred, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0") {
			event := new(ControllerRegistryOwnershipTransferred)
			if err := _ControllerRegistry.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OwnershipTransferred event not found")
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRegistryOwnershipTransferred)
				if err := _ControllerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ControllerRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the ControllerRegistry contract.
type ControllerRegistryRegistrationIterator struct {
	Event *ControllerRegistryRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ControllerRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRegistryRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ControllerRegistryRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ControllerRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRegistryRegistration represents a Registration event raised by the ControllerRegistry contract.
type ControllerRegistryRegistration struct {
	Controller common.Address
	Raw        ethTypes.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryRegistrationIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.FilterLogs(opts, "Registration", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryRegistrationIterator{contract: _ControllerRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(address indexed controller)
func (manager *ControllerRegistryContract) ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryRegistration, error) {
	return manager.contract.ParseRegistrationFromReceipt(receipt)
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryRegistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524") {
			event := new(ControllerRegistryRegistration)
			if err := _ControllerRegistry.contract.UnpackLog(event, "Registration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Registration event not found")
}

// WatchRegistration is a free log subscription operation binding the contract event 0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryRegistration, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.WatchLogs(opts, "Registration", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRegistryRegistration)
				if err := _ControllerRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ControllerRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the ControllerRegistry contract.
type ControllerRegistryUnregistrationIterator struct {
	Event *ControllerRegistryUnregistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ControllerRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRegistryUnregistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ControllerRegistryUnregistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ControllerRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRegistryUnregistration represents a Unregistration event raised by the ControllerRegistry contract.
type ControllerRegistryUnregistration struct {
	Controller common.Address
	Raw        ethTypes.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryUnregistrationIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.FilterLogs(opts, "Unregistration", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryUnregistrationIterator{contract: _ControllerRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(address indexed controller)
func (manager *ControllerRegistryContract) ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryUnregistration, error) {
	return manager.contract.ParseUnregistrationFromReceipt(receipt)
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*ControllerRegistryUnregistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9") {
			event := new(ControllerRegistryUnregistration)
			if err := _ControllerRegistry.contract.UnpackLog(event, "Unregistration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Unregistration event not found")
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryUnregistration, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.WatchLogs(opts, "Unregistration", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRegistryUnregistration)
				if err := _ControllerRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
