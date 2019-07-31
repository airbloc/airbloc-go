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

// DataTypeRegistryABI is the input ABI used to generate the binding from.
const DataTypeRegistryABI = "{\"Constructor\":{\"Name\":\"\",\"Const\":false,\"Inputs\":null,\"Outputs\":null},\"Methods\":{\"exists\":{\"Name\":\"exists\",\"Const\":true,\"Inputs\":[{\"Name\":\"name\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"get\":{\"Name\":\"get\",\"Const\":true,\"Inputs\":[{\"Name\":\"name\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":25,\"Type\":{},\"Size\":0,\"T\":6,\"TupleElems\":[{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null}],\"TupleRawNames\":[\"name\",\"owner\",\"schemaHash\"]},\"Indexed\":false}]},\"isOwner\":{\"Name\":\"isOwner\",\"Const\":true,\"Inputs\":[{\"Name\":\"name\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"owner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"register\":{\"Name\":\"register\",\"Const\":false,\"Inputs\":[{\"Name\":\"name\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"schemaHash\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"unregister\":{\"Name\":\"unregister\",\"Const\":false,\"Inputs\":[{\"Name\":\"name\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]}},\"Events\":{\"Registration\":{\"Name\":\"Registration\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"name\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"Unregistration\":{\"Name\":\"Unregistration\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"name\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]}}}"

// DataTypeRegistry is an auto generated Go binding around an Ethereum contract.
type DataTypeRegistry struct {
	address                    common.Address
	txHash                     common.Hash
	createdAt                  *big.Int
	DataTypeRegistryCaller     // Read-only binding to the contract
	DataTypeRegistryTransactor // Write-only binding to the contract
	DataTypeRegistryFilterer   // Log filterer for contract events
}

// Address is getter method of DataTypeRegistry.address
func (_DataTypeRegistry *DataTypeRegistry) Address() common.Address {
	return _DataTypeRegistry.address
}

// TxHash is getter method of DataTypeRegistry.txHash
func (_DataTypeRegistry *DataTypeRegistry) TxHash() common.Hash {
	return _DataTypeRegistry.txHash
}

// CreatedAt is getter method of DataTypeRegistry.createdAt
func (_DataTypeRegistry *DataTypeRegistry) CreatedAt() *big.Int {
	return _DataTypeRegistry.createdAt
}

// DataTypeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataTypeRegistrySession struct {
	Contract     *DataTypeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataTypeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataTypeRegistryRaw struct {
	Contract *DataTypeRegistry // Generic contract binding to access the raw methods on
}

// NewDataTypeRegistry creates a new instance of DataTypeRegistry, bound to a specific deployed contract.
func NewDataTypeRegistry(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (*DataTypeRegistry, error) {
	contract, err := bindDataTypeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistry{
		address:                    address,
		txHash:                     txHash,
		createdAt:                  createdAt,
		DataTypeRegistryCaller:     DataTypeRegistryCaller{contract: contract},
		DataTypeRegistryTransactor: DataTypeRegistryTransactor{contract: contract},
		DataTypeRegistryFilterer:   DataTypeRegistryFilterer{contract: contract},
	}, nil
}

// bindDataTypeRegistry binds a generic wrapper to an already deployed contract.
func bindDataTypeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataTypeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTypeRegistry *DataTypeRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataTypeRegistry.Contract.DataTypeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTypeRegistry *DataTypeRegistryRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.Contract.DataTypeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTypeRegistry *DataTypeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.Contract.DataTypeRegistryTransactor.contract.Transact(opts, method, params...)
}

// DataTypeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataTypeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataTypeRegistryCallerSession struct {
	Contract *DataTypeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DataTypeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataTypeRegistryCallerRaw struct {
	Contract *DataTypeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NewDataTypeRegistryCaller creates a new read-only instance of DataTypeRegistry, bound to a specific deployed contract.
func NewDataTypeRegistryCaller(address common.Address, caller bind.ContractCaller) (*DataTypeRegistryCaller, error) {
	contract, err := bindDataTypeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryCaller{contract: contract}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTypeRegistry *DataTypeRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataTypeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// DataTypeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTypeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTypeRegistryTransactorSession struct {
	Contract     *DataTypeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DataTypeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTypeRegistryTransactorRaw struct {
	Contract *DataTypeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataTypeRegistryTransactor creates a new write-only instance of DataTypeRegistry, bound to a specific deployed contract.
func NewDataTypeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTypeRegistryTransactor, error) {
	contract, err := bindDataTypeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTypeRegistry *DataTypeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTypeRegistry *DataTypeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.Contract.contract.Transact(opts, method, params...)
}

// DataTypeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataTypeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewDataTypeRegistryFilterer creates a new log filterer instance of DataTypeRegistry, bound to a specific deployed contract.
func NewDataTypeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DataTypeRegistryFilterer, error) {
	contract, err := bindDataTypeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryFilterer{contract: contract}, nil
}

//go:generate mockgen -source data_type_registry.go -destination ./mocks/mock_data_type_registry.go -package mocks IDataTypeRegistryManager,IDataTypeRegistryContract
type IDataTypeRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	Exists(name string) (bool, error)
	Get(name string) (types.DataType, error)
	IsOwner(name string, owner common.Address) (bool, error)

	// Transact methods
	Register(ctx context.Context, name string, schemaHash common.Hash) error
	Unregister(ctx context.Context, name string) error

	FilterRegistration(opts *bind.FilterOpts) (*DataTypeRegistryRegistrationIterator, error)
	WatchRegistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryRegistration) (event.Subscription, error)

	FilterUnregistration(opts *bind.FilterOpts) (*DataTypeRegistryUnregistrationIterator, error)
	WatchUnregistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryUnregistration) (event.Subscription, error)
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
	client   blockchain.TxClient
	contract *DataTypeRegistry
	DataTypeRegistryFilterer
}

// Address is getter method of DataTypeRegistry.address
func (c *DataTypeRegistryContract) Address() common.Address {
	return c.contract.Address()
}

// TxHash is getter method of DataTypeRegistry.txHash
func (c *DataTypeRegistryContract) TxHash() common.Hash {
	return c.contract.TxHash()
}

// CreatedAt is getter method of DataTypeRegistry.createdAt
func (c *DataTypeRegistryContract) CreatedAt() *big.Int {
	return c.contract.CreatedAt()
}

// NewManager makes new *Manager struct
func NewDataTypeRegistryContract(client blockchain.TxClient) IDataTypeRegistryContract {
	contract := client.GetContract(&DataTypeRegistry{}).(*DataTypeRegistry)
	return &DataTypeRegistryContract{
		client:                   client,
		contract:                 contract,
		DataTypeRegistryFilterer: contract.DataTypeRegistryFilterer,
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("DataTypeRegistry", (&DataTypeRegistry{}).new)
	blockchain.RegisterSelector("0x656afdee", "register(string,bytes32)")
	blockchain.RegisterSelector("0x6598a1ae", "unregister(string)")
}

func (_DataTypeRegistry *DataTypeRegistry) new(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (interface{}, error) {
	return NewDataTypeRegistry(address, txHash, createdAt, backend)
}

type IDataTypeRegistryCalls interface {
	Exists(name string) (bool, error)
	Get(name string) (types.DataType, error)
	IsOwner(name string, owner common.Address) (bool, error)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (c *DataTypeRegistryContract) Exists(name string) (bool, error) {
	return c.contract.Exists(nil, name)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (_DataTypeRegistry *DataTypeRegistryCaller) Exists(opts *bind.CallOpts, name string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _DataTypeRegistry.contract.Call(opts, out, "exists", name)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (_DataTypeRegistry *DataTypeRegistrySession) Exists(name string) (bool, error) {
	return _DataTypeRegistry.Contract.Exists(&_DataTypeRegistry.CallOpts, name)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (_DataTypeRegistry *DataTypeRegistryCallerSession) Exists(name string) (bool, error) {
	return _DataTypeRegistry.Contract.Exists(&_DataTypeRegistry.CallOpts, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (c *DataTypeRegistryContract) Get(name string) (types.DataType, error) {
	return c.contract.Get(nil, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (_DataTypeRegistry *DataTypeRegistryCaller) Get(opts *bind.CallOpts, name string) (types.DataType, error) {
	ret := new(types.DataType)

	out := ret
	err := _DataTypeRegistry.contract.Call(opts, out, "get", name)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (_DataTypeRegistry *DataTypeRegistrySession) Get(name string) (types.DataType, error) {
	return _DataTypeRegistry.Contract.Get(&_DataTypeRegistry.CallOpts, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (_DataTypeRegistry *DataTypeRegistryCallerSession) Get(name string) (types.DataType, error) {
	return _DataTypeRegistry.Contract.Get(&_DataTypeRegistry.CallOpts, name)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (c *DataTypeRegistryContract) IsOwner(name string, owner common.Address) (bool, error) {
	return c.contract.IsOwner(nil, name, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (_DataTypeRegistry *DataTypeRegistryCaller) IsOwner(opts *bind.CallOpts, name string, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _DataTypeRegistry.contract.Call(opts, out, "isOwner", name, owner)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (_DataTypeRegistry *DataTypeRegistrySession) IsOwner(name string, owner common.Address) (bool, error) {
	return _DataTypeRegistry.Contract.IsOwner(&_DataTypeRegistry.CallOpts, name, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (_DataTypeRegistry *DataTypeRegistryCallerSession) IsOwner(name string, owner common.Address) (bool, error) {
	return _DataTypeRegistry.Contract.IsOwner(&_DataTypeRegistry.CallOpts, name, owner)
}

type IDataTypeRegistryTransacts interface {
	Register(ctx context.Context, name string, schemaHash common.Hash) (*ethTypes.Receipt, error)
	Unregister(ctx context.Context, name string) (*ethTypes.Receipt, error)
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (c *DataTypeRegistryContract) Register(ctx context.Context, name string, schemaHash common.Hash) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Register(c.client.Account(), name, schemaHash)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (_DataTypeRegistry *DataTypeRegistryTransactor) Register(opts *bind.TransactOpts, name string, schemaHash common.Hash) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.contract.Transact(opts, "register", name, schemaHash)
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (_DataTypeRegistry *DataTypeRegistrySession) Register(name string, schemaHash common.Hash) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.Contract.Register(&_DataTypeRegistry.TransactOpts, name, schemaHash)
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (_DataTypeRegistry *DataTypeRegistryTransactorSession) Register(name string, schemaHash common.Hash) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.Contract.Register(&_DataTypeRegistry.TransactOpts, name, schemaHash)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (c *DataTypeRegistryContract) Unregister(ctx context.Context, name string) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Unregister(c.client.Account(), name)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (_DataTypeRegistry *DataTypeRegistryTransactor) Unregister(opts *bind.TransactOpts, name string) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.contract.Transact(opts, "unregister", name)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (_DataTypeRegistry *DataTypeRegistrySession) Unregister(name string) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.Contract.Unregister(&_DataTypeRegistry.TransactOpts, name)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (_DataTypeRegistry *DataTypeRegistryTransactorSession) Unregister(name string) (*ethTypes.Transaction, error) {
	return _DataTypeRegistry.Contract.Unregister(&_DataTypeRegistry.TransactOpts, name)
}

type IDataTypeRegistryEvents interface {
	FilterRegistration(opts *bind.FilterOpts) (*DataTypeRegistryRegistrationIterator, error)
	ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*DataTypeRegistryRegistration, error)
	WatchRegistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryRegistration) (event.Subscription, error)

	FilterUnregistration(opts *bind.FilterOpts) (*DataTypeRegistryUnregistrationIterator, error)
	ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*DataTypeRegistryUnregistration, error)
	WatchUnregistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryUnregistration) (event.Subscription, error)
}

// DataTypeRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the DataTypeRegistry contract.
type DataTypeRegistryRegistrationIterator struct {
	Event *DataTypeRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *DataTypeRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataTypeRegistryRegistration)
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
		it.Event = new(DataTypeRegistryRegistration)
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
func (it *DataTypeRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataTypeRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataTypeRegistryRegistration represents a Registration event raised by the DataTypeRegistry contract.
type DataTypeRegistryRegistration struct {
	Name string
	Raw  ethTypes.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) FilterRegistration(opts *bind.FilterOpts) (*DataTypeRegistryRegistrationIterator, error) {

	logs, sub, err := _DataTypeRegistry.contract.FilterLogs(opts, "Registration")
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryRegistrationIterator{contract: _DataTypeRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(string name)
func (manager *DataTypeRegistryContract) ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*DataTypeRegistryRegistration, error) {
	return manager.contract.ParseRegistrationFromReceipt(receipt)
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*DataTypeRegistryRegistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948") {
			event := new(DataTypeRegistryRegistration)
			if err := _DataTypeRegistry.contract.UnpackLog(event, "Registration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Registration event not found")
}

// WatchRegistration is a free log subscription operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryRegistration) (event.Subscription, error) {

	logs, sub, err := _DataTypeRegistry.contract.WatchLogs(opts, "Registration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataTypeRegistryRegistration)
				if err := _DataTypeRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
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

// DataTypeRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the DataTypeRegistry contract.
type DataTypeRegistryUnregistrationIterator struct {
	Event *DataTypeRegistryUnregistration // Event containing the contract specifics and raw log

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
func (it *DataTypeRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataTypeRegistryUnregistration)
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
		it.Event = new(DataTypeRegistryUnregistration)
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
func (it *DataTypeRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataTypeRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataTypeRegistryUnregistration represents a Unregistration event raised by the DataTypeRegistry contract.
type DataTypeRegistryUnregistration struct {
	Name string
	Raw  ethTypes.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts) (*DataTypeRegistryUnregistrationIterator, error) {

	logs, sub, err := _DataTypeRegistry.contract.FilterLogs(opts, "Unregistration")
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryUnregistrationIterator{contract: _DataTypeRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(string name)
func (manager *DataTypeRegistryContract) ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*DataTypeRegistryUnregistration, error) {
	return manager.contract.ParseUnregistrationFromReceipt(receipt)
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*DataTypeRegistryUnregistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83") {
			event := new(DataTypeRegistryUnregistration)
			if err := _DataTypeRegistry.contract.UnpackLog(event, "Unregistration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Unregistration event not found")
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryUnregistration) (event.Subscription, error) {

	logs, sub, err := _DataTypeRegistry.contract.WatchLogs(opts, "Unregistration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataTypeRegistryUnregistration)
				if err := _DataTypeRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
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
