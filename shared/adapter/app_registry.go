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

// AppRegistryABI is the input ABI used to generate the binding from.
const AppRegistryABI = "{\"Constructor\":{\"Name\":\"\",\"Const\":false,\"Inputs\":null,\"Outputs\":null},\"Methods\":{\"exists\":{\"Name\":\"exists\",\"Const\":true,\"Inputs\":[{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"get\":{\"Name\":\"get\",\"Const\":true,\"Inputs\":[{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":25,\"Type\":{},\"Size\":0,\"T\":6,\"TupleElems\":[{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null}],\"TupleRawNames\":[\"name\",\"owner\",\"hashedName\"]},\"Indexed\":false}]},\"isOwner\":{\"Name\":\"isOwner\",\"Const\":true,\"Inputs\":[{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"owner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"register\":{\"Name\":\"register\",\"Const\":false,\"Inputs\":[{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"transferAppOwner\":{\"Name\":\"transferAppOwner\",\"Const\":false,\"Inputs\":[{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"newOwner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"unregister\":{\"Name\":\"unregister\",\"Const\":false,\"Inputs\":[{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]}},\"Events\":{\"AppOwnerTransferred\":{\"Name\":\"AppOwnerTransferred\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"hashedAppName\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"oldOwner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"newOwner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"Registration\":{\"Name\":\"Registration\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"hashedAppName\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"Unregistration\":{\"Name\":\"Unregistration\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"hashedAppName\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]}}}"

// AppRegistry is an auto generated Go binding around an Ethereum contract.
type AppRegistry struct {
	Address               common.Address
	AppRegistryCaller     // Read-only binding to the contract
	AppRegistryTransactor // Write-only binding to the contract
	AppRegistryFilterer   // Log filterer for contract events
}

// AppRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppRegistrySession struct {
	Contract     *AppRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRegistryRaw struct {
	Contract *AppRegistry // Generic contract binding to access the raw methods on
}

// NewAppRegistry creates a new instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistry(address common.Address, backend bind.ContractBackend) (*AppRegistry, error) {
	contract, err := bindAppRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppRegistry{
		Address:               address,
		AppRegistryCaller:     AppRegistryCaller{contract: contract},
		AppRegistryTransactor: AppRegistryTransactor{contract: contract},
		AppRegistryFilterer:   AppRegistryFilterer{contract: contract},
	}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppRegistry *AppRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AppRegistry.Contract.AppRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppRegistry *AppRegistryRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transact(opts, method, params...)
}

// AppRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppRegistryCallerSession struct {
	Contract *AppRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AppRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppRegistryCallerRaw struct {
	Contract *AppRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NewAppRegistryCaller creates a new read-only instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryCaller(address common.Address, caller bind.ContractCaller) (*AppRegistryCaller, error) {
	contract, err := bindAppRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppRegistryCaller{contract: contract}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppRegistry *AppRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AppRegistry.Contract.contract.Call(opts, result, method, params...)
}

// AppRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppRegistryTransactorSession struct {
	Contract     *AppRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AppRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppRegistryTransactorRaw struct {
	Contract *AppRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppRegistryTransactor creates a new write-only instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AppRegistryTransactor, error) {
	contract, err := bindAppRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppRegistryTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppRegistry *AppRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.contract.Transact(opts, method, params...)
}

// AppRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewAppRegistryFilterer creates a new log filterer instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AppRegistryFilterer, error) {
	contract, err := bindAppRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppRegistryFilterer{contract: contract}, nil
}

//go:generate mockgen -source app_registry.go -destination ./mocks/mock_app_registry.go -package mocks IAppRegistryManager,IAppRegistryContract
type IAppRegistryManager interface {
	// Call methods
	Exists(appName string) (bool, error)
	Get(appName string) (types.App, error)
	IsOwner(appName string, owner common.Address) (bool, error)

	// Transact methods
	Register(ctx context.Context, appName string) error
	TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) error
	Unregister(ctx context.Context, appName string) error
}

type IAppRegistryContract interface {
	// Call methods
	Exists(appName string) (bool, error)
	Get(appName string) (types.App, error)
	IsOwner(appName string, owner common.Address) (bool, error)

	// Transact methods
	Register(ctx context.Context, appName string) (*ethTypes.Receipt, error)
	TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) (*ethTypes.Receipt, error)
	Unregister(ctx context.Context, appName string) (*ethTypes.Receipt, error)

	// Event parser
	ParseAppOwnerTransferredFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryAppOwnerTransferred, error)
	ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryRegistration, error)
	ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryUnregistration, error)
}

// Manager is contract wrapper struct
type AppRegistryContract struct {
	client   blockchain.TxClient
	contract *AppRegistry
}

// NewManager makes new *Manager struct
func NewAppRegistryContract(client blockchain.TxClient) IAppRegistryContract {
	contract := client.GetContract(&AppRegistry{})
	return &AppRegistryManager{
		client:   client,
		contract: contract.(*AppRegistry),
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.ContractList["AppRegistry"] = (&AppRegistry{}).new
	blockchain.RegisterSelector("0xf2c298be", "register(string)")
	blockchain.RegisterSelector("0x1a9dff9f", "transferAppOwner(string,address)")
	blockchain.RegisterSelector("0x6598a1ae", "unregister(string)")
}

// bindAppRegistry binds a generic wrapper to an already deployed contract.
func bindAppRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_AppRegistry *AppRegistry) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewAppRegistry(address, backend)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (manager *AppRegistryContract) Exists(appName string) (bool, error) {
	return manager.contract.Exists(nil, appName)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (_AppRegistry *AppRegistryCaller) Exists(opts *bind.CallOpts, appName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _AppRegistry.contract.Call(opts, out, "exists", appName)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (_AppRegistry *AppRegistrySession) Exists(appName string) (bool, error) {
	return _AppRegistry.Contract.Exists(&_AppRegistry.CallOpts, appName)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (_AppRegistry *AppRegistryCallerSession) Exists(appName string) (bool, error) {
	return _AppRegistry.Contract.Exists(&_AppRegistry.CallOpts, appName)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (manager *AppRegistryContract) Get(appName string) (types.App, error) {
	return manager.contract.Get(nil, appName)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (_AppRegistry *AppRegistryCaller) Get(opts *bind.CallOpts, appName string) (types.App, error) {
	ret := new(types.App)

	out := ret
	err := _AppRegistry.contract.Call(opts, out, "get", appName)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (_AppRegistry *AppRegistrySession) Get(appName string) (types.App, error) {
	return _AppRegistry.Contract.Get(&_AppRegistry.CallOpts, appName)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (_AppRegistry *AppRegistryCallerSession) Get(appName string) (types.App, error) {
	return _AppRegistry.Contract.Get(&_AppRegistry.CallOpts, appName)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (manager *AppRegistryContract) IsOwner(appName string, owner common.Address) (bool, error) {
	return manager.contract.IsOwner(nil, appName, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (_AppRegistry *AppRegistryCaller) IsOwner(opts *bind.CallOpts, appName string, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _AppRegistry.contract.Call(opts, out, "isOwner", appName, owner)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (_AppRegistry *AppRegistrySession) IsOwner(appName string, owner common.Address) (bool, error) {
	return _AppRegistry.Contract.IsOwner(&_AppRegistry.CallOpts, appName, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (_AppRegistry *AppRegistryCallerSession) IsOwner(appName string, owner common.Address) (bool, error) {
	return _AppRegistry.Contract.IsOwner(&_AppRegistry.CallOpts, appName, owner)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (manager *AppRegistryContract) Register(ctx context.Context, appName string) (*ethTypes.Receipt, error) {
	tx, err := manager.contract.Register(manager.client.Account(), appName)
	if err != nil {
		return nil, err
	}
	return manager.client.WaitMined(ctx, tx)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (_AppRegistry *AppRegistryTransactor) Register(opts *bind.TransactOpts, appName string) (*ethTypes.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "register", appName)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (_AppRegistry *AppRegistrySession) Register(appName string) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.Register(&_AppRegistry.TransactOpts, appName)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (_AppRegistry *AppRegistryTransactorSession) Register(appName string) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.Register(&_AppRegistry.TransactOpts, appName)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (manager *AppRegistryContract) TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) (*ethTypes.Receipt, error) {
	tx, err := manager.contract.TransferAppOwner(manager.client.Account(), appName, newOwner)
	if err != nil {
		return nil, err
	}
	return manager.client.WaitMined(ctx, tx)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (_AppRegistry *AppRegistryTransactor) TransferAppOwner(opts *bind.TransactOpts, appName string, newOwner common.Address) (*ethTypes.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "transferAppOwner", appName, newOwner)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (_AppRegistry *AppRegistrySession) TransferAppOwner(appName string, newOwner common.Address) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.TransferAppOwner(&_AppRegistry.TransactOpts, appName, newOwner)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (_AppRegistry *AppRegistryTransactorSession) TransferAppOwner(appName string, newOwner common.Address) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.TransferAppOwner(&_AppRegistry.TransactOpts, appName, newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (manager *AppRegistryContract) Unregister(ctx context.Context, appName string) (*ethTypes.Receipt, error) {
	tx, err := manager.contract.Unregister(manager.client.Account(), appName)
	if err != nil {
		return nil, err
	}
	return manager.client.WaitMined(ctx, tx)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (_AppRegistry *AppRegistryTransactor) Unregister(opts *bind.TransactOpts, appName string) (*ethTypes.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "unregister", appName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (_AppRegistry *AppRegistrySession) Unregister(appName string) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.Unregister(&_AppRegistry.TransactOpts, appName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (_AppRegistry *AppRegistryTransactorSession) Unregister(appName string) (*ethTypes.Transaction, error) {
	return _AppRegistry.Contract.Unregister(&_AppRegistry.TransactOpts, appName)
}

// AppRegistryAppOwnerTransferredIterator is returned from FilterAppOwnerTransferred and is used to iterate over the raw logs and unpacked data for AppOwnerTransferred events raised by the AppRegistry contract.
type AppRegistryAppOwnerTransferredIterator struct {
	Event *AppRegistryAppOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *AppRegistryAppOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryAppOwnerTransferred)
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
		it.Event = new(AppRegistryAppOwnerTransferred)
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
func (it *AppRegistryAppOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryAppOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryAppOwnerTransferred represents a AppOwnerTransferred event raised by the AppRegistry contract.
type AppRegistryAppOwnerTransferred struct {
	HashedAppName common.Hash
	AppName       string
	OldOwner      common.Address
	NewOwner      common.Address
	Raw           ethTypes.Log // Blockchain specific contextual infos
}

// FilterAppOwnerTransferred is a free log retrieval operation binding the contract event 0x9323f5fe9b72ac1fe704b80d8d53e6538f63d0d041068ef274c35b41d1cbc1de.
//
// Solidity: event AppOwnerTransferred(bytes32 indexed hashedAppName, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *AppRegistryFilterer) FilterAppOwnerTransferred(opts *bind.FilterOpts, hashedAppName []common.Hash, oldOwner []common.Address) (*AppRegistryAppOwnerTransferredIterator, error) {

	var hashedAppNameRule []interface{}
	for _, hashedAppNameItem := range hashedAppName {
		hashedAppNameRule = append(hashedAppNameRule, hashedAppNameItem)
	}

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "AppOwnerTransferred", hashedAppNameRule, oldOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryAppOwnerTransferredIterator{contract: _AppRegistry.contract, event: "AppOwnerTransferred", logs: logs, sub: sub}, nil
}

// FilterAppOwnerTransferred parses the event from given transaction receipt.
//
// Solidity: event AppOwnerTransferred(bytes32 indexed hashedAppName, string appName, address indexed oldOwner, address newOwner)
func (manager *AppRegistryContract) ParseAppOwnerTransferredFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryAppOwnerTransferred, error) {
	return manager.contract.ParseAppOwnerTransferredFromReceipt(receipt)
}

// FilterAppOwnerTransferred parses the event from given transaction receipt.
//
// Solidity: event AppOwnerTransferred(bytes32 indexed hashedAppName, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *AppRegistryFilterer) ParseAppOwnerTransferredFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryAppOwnerTransferred, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x9323f5fe9b72ac1fe704b80d8d53e6538f63d0d041068ef274c35b41d1cbc1de") {
			event := new(AppRegistryAppOwnerTransferred)
			if err := _AppRegistry.contract.UnpackLog(event, "AppOwnerTransferred", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("AppOwnerTransferred event not found")
}

// WatchAppOwnerTransferred is a free log subscription operation binding the contract event 0x9323f5fe9b72ac1fe704b80d8d53e6538f63d0d041068ef274c35b41d1cbc1de.
//
// Solidity: event AppOwnerTransferred(bytes32 indexed hashedAppName, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *AppRegistryFilterer) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AppRegistryAppOwnerTransferred, hashedAppName []common.Hash, oldOwner []common.Address) (event.Subscription, error) {

	var hashedAppNameRule []interface{}
	for _, hashedAppNameItem := range hashedAppName {
		hashedAppNameRule = append(hashedAppNameRule, hashedAppNameItem)
	}

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "AppOwnerTransferred", hashedAppNameRule, oldOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryAppOwnerTransferred)
				if err := _AppRegistry.contract.UnpackLog(event, "AppOwnerTransferred", log); err != nil {
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

// AppRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the AppRegistry contract.
type AppRegistryRegistrationIterator struct {
	Event *AppRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *AppRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryRegistration)
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
		it.Event = new(AppRegistryRegistration)
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
func (it *AppRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRegistration represents a Registration event raised by the AppRegistry contract.
type AppRegistryRegistration struct {
	HashedAppName common.Hash
	AppName       string
	Raw           ethTypes.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xe7e1383b88439b9522e6630da35051999780d58947518c9a3d1620d19b1bc886.
//
// Solidity: event Registration(bytes32 indexed hashedAppName, string appName)
func (_AppRegistry *AppRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*AppRegistryRegistrationIterator, error) {

	var hashedAppNameRule []interface{}
	for _, hashedAppNameItem := range hashedAppName {
		hashedAppNameRule = append(hashedAppNameRule, hashedAppNameItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Registration", hashedAppNameRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRegistrationIterator{contract: _AppRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(bytes32 indexed hashedAppName, string appName)
func (manager *AppRegistryContract) ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryRegistration, error) {
	return manager.contract.ParseRegistrationFromReceipt(receipt)
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(bytes32 indexed hashedAppName, string appName)
func (_AppRegistry *AppRegistryFilterer) ParseRegistrationFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryRegistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xe7e1383b88439b9522e6630da35051999780d58947518c9a3d1620d19b1bc886") {
			event := new(AppRegistryRegistration)
			if err := _AppRegistry.contract.UnpackLog(event, "Registration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Registration event not found")
}

// WatchRegistration is a free log subscription operation binding the contract event 0xe7e1383b88439b9522e6630da35051999780d58947518c9a3d1620d19b1bc886.
//
// Solidity: event Registration(bytes32 indexed hashedAppName, string appName)
func (_AppRegistry *AppRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *AppRegistryRegistration, hashedAppName []common.Hash) (event.Subscription, error) {

	var hashedAppNameRule []interface{}
	for _, hashedAppNameItem := range hashedAppName {
		hashedAppNameRule = append(hashedAppNameRule, hashedAppNameItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Registration", hashedAppNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryRegistration)
				if err := _AppRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
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

// AppRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the AppRegistry contract.
type AppRegistryUnregistrationIterator struct {
	Event *AppRegistryUnregistration // Event containing the contract specifics and raw log

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
func (it *AppRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryUnregistration)
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
		it.Event = new(AppRegistryUnregistration)
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
func (it *AppRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryUnregistration represents a Unregistration event raised by the AppRegistry contract.
type AppRegistryUnregistration struct {
	HashedAppName common.Hash
	AppName       string
	Raw           ethTypes.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0xe2d00f1029a39aa8484ea81b8c2794180aa48eb0a1f28c721acdc94789f1d638.
//
// Solidity: event Unregistration(bytes32 indexed hashedAppName, string appName)
func (_AppRegistry *AppRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*AppRegistryUnregistrationIterator, error) {

	var hashedAppNameRule []interface{}
	for _, hashedAppNameItem := range hashedAppName {
		hashedAppNameRule = append(hashedAppNameRule, hashedAppNameItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Unregistration", hashedAppNameRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryUnregistrationIterator{contract: _AppRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(bytes32 indexed hashedAppName, string appName)
func (manager *AppRegistryContract) ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryUnregistration, error) {
	return manager.contract.ParseUnregistrationFromReceipt(receipt)
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(bytes32 indexed hashedAppName, string appName)
func (_AppRegistry *AppRegistryFilterer) ParseUnregistrationFromReceipt(receipt *ethTypes.Receipt) (*AppRegistryUnregistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xe2d00f1029a39aa8484ea81b8c2794180aa48eb0a1f28c721acdc94789f1d638") {
			event := new(AppRegistryUnregistration)
			if err := _AppRegistry.contract.UnpackLog(event, "Unregistration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Unregistration event not found")
}

// WatchUnregistration is a free log subscription operation binding the contract event 0xe2d00f1029a39aa8484ea81b8c2794180aa48eb0a1f28c721acdc94789f1d638.
//
// Solidity: event Unregistration(bytes32 indexed hashedAppName, string appName)
func (_AppRegistry *AppRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *AppRegistryUnregistration, hashedAppName []common.Hash) (event.Subscription, error) {

	var hashedAppNameRule []interface{}
	for _, hashedAppNameItem := range hashedAppName {
		hashedAppNameRule = append(hashedAppNameRule, hashedAppNameItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Unregistration", hashedAppNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryUnregistration)
				if err := _AppRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
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