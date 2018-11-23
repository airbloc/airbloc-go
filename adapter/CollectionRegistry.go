// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"
	"strings"

	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = ablCommon.IDFromString
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CollectionRegistryABI is the input ABI used to generate the binding from.
const CollectionRegistryABI = "[{\"inputs\":[{\"name\":\"_accounts\",\"type\":\"address\"},{\"name\":\"_appReg\",\"type\":\"address\"},{\"name\":\"_schemaReg\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registrar\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"appId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"collectionId\",\"type\":\"bytes8\"}],\"name\":\"Registration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"appId\",\"type\":\"bytes8\"}],\"name\":\"Unregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"Allowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"Denied\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes8\"},{\"name\":\"_schemaId\",\"type\":\"bytes8\"},{\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"appId\",\"type\":\"bytes8\"},{\"name\":\"schemaId\",\"type\":\"bytes8\"},{\"name\":\"incentiveRatioSelf\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"allowByPassword\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"denyByPassword\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"name\":\"user\",\"type\":\"bytes8\"}],\"name\":\"isCollectionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"name\":\"user\",\"type\":\"bytes8\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"isCollectionAllowedAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CollectionRegistry is an auto generated Go binding around an Ethereum contract.
type CollectionRegistry struct {
	Address                      common.Address
	CollectionRegistryCaller     // Read-only binding to the contract
	CollectionRegistryTransactor // Write-only binding to the contract
	CollectionRegistryFilterer   // Log filterer for contract events
}

// CollectionRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectionRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectionRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectionRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectionRegistrySession struct {
	Contract     *CollectionRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CollectionRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectionRegistryCallerSession struct {
	Contract *CollectionRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CollectionRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectionRegistryTransactorSession struct {
	Contract     *CollectionRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CollectionRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectionRegistryRaw struct {
	Contract *CollectionRegistry // Generic contract binding to access the raw methods on
}

// CollectionRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectionRegistryCallerRaw struct {
	Contract *CollectionRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CollectionRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectionRegistryTransactorRaw struct {
	Contract *CollectionRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

type CollectionRegistryAuth struct {
	AuthorizedAt *big.Int
	IsAllowed    bool
}

type CollectionRegistryCollection struct {
	AppId            ablCommon.ID
	DataCollectionOf map[ablCommon.ID]CollectionRegistryAuth
	Policy           CollectionRegistryIncentivePolicy
	SchemaId         ablCommon.ID
}

type CollectionRegistryIncentivePolicy struct {
	Owner *big.Int
	Self  *big.Int
}

func init() {
	blockchain.ContractList["CollectionRegistry"] = (&CollectionRegistry{}).new
}

// NewCollectionRegistry creates a new instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistry(address common.Address, backend bind.ContractBackend) (*CollectionRegistry, error) {
	contract, err := bindCollectionRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistry{
		Address:                      address,
		CollectionRegistryCaller:     CollectionRegistryCaller{contract: contract},
		CollectionRegistryTransactor: CollectionRegistryTransactor{contract: contract},
		CollectionRegistryFilterer:   CollectionRegistryFilterer{contract: contract},
	}, nil
}

// NewCollectionRegistryCaller creates a new read-only instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryCaller(address common.Address, caller bind.ContractCaller) (*CollectionRegistryCaller, error) {
	contract, err := bindCollectionRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryCaller{contract: contract}, nil
}

// NewCollectionRegistryTransactor creates a new write-only instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectionRegistryTransactor, error) {
	contract, err := bindCollectionRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryTransactor{contract: contract}, nil
}

// NewCollectionRegistryFilterer creates a new log filterer instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectionRegistryFilterer, error) {
	contract, err := bindCollectionRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryFilterer{contract: contract}, nil
}

// bindCollectionRegistry binds a generic wrapper to an already deployed contract.
func bindCollectionRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_CollectionRegistry *CollectionRegistry) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewCollectionRegistry(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionRegistry *CollectionRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CollectionRegistry.Contract.CollectionRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionRegistry *CollectionRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.CollectionRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionRegistry *CollectionRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.CollectionRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionRegistry *CollectionRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CollectionRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionRegistry *CollectionRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionRegistry *CollectionRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) Exists(opts *bind.CallOpts, _id [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "exists", _id)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) Exists(_id [8]byte) (bool, error) {
	return _CollectionRegistry.Contract.Exists(&_CollectionRegistry.CallOpts, _id)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) Exists(_id [8]byte) (bool, error) {
	return _CollectionRegistry.Contract.Exists(&_CollectionRegistry.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x47ba65d2.
//
// Solidity: function get(_id bytes8) constant returns(appId bytes8, schemaId bytes8, incentiveRatioSelf uint256)
func (_CollectionRegistry *CollectionRegistryCaller) Get(opts *bind.CallOpts, _id [8]byte) (struct {
	AppId              [8]byte
	SchemaId           [8]byte
	IncentiveRatioSelf *big.Int
}, error) {
	ret := new(struct {
		AppId              [8]byte
		SchemaId           [8]byte
		IncentiveRatioSelf *big.Int
	})
	out := ret
	err := _CollectionRegistry.contract.Call(opts, out, "get", _id)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0x47ba65d2.
//
// Solidity: function get(_id bytes8) constant returns(appId bytes8, schemaId bytes8, incentiveRatioSelf uint256)
func (_CollectionRegistry *CollectionRegistrySession) Get(_id [8]byte) (struct {
	AppId              [8]byte
	SchemaId           [8]byte
	IncentiveRatioSelf *big.Int
}, error) {
	return _CollectionRegistry.Contract.Get(&_CollectionRegistry.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x47ba65d2.
//
// Solidity: function get(_id bytes8) constant returns(appId bytes8, schemaId bytes8, incentiveRatioSelf uint256)
func (_CollectionRegistry *CollectionRegistryCallerSession) Get(_id [8]byte) (struct {
	AppId              [8]byte
	SchemaId           [8]byte
	IncentiveRatioSelf *big.Int
}, error) {
	return _CollectionRegistry.Contract.Get(&_CollectionRegistry.CallOpts, _id)
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xf8907491.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) IsCollectionAllowed(opts *bind.CallOpts, collectionId [8]byte, user [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "isCollectionAllowed", collectionId, user)
	return *ret0, err
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xf8907491.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) IsCollectionAllowed(collectionId [8]byte, user [8]byte) (bool, error) {
	return _CollectionRegistry.Contract.IsCollectionAllowed(&_CollectionRegistry.CallOpts, collectionId, user)
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xf8907491.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) IsCollectionAllowed(collectionId [8]byte, user [8]byte) (bool, error) {
	return _CollectionRegistry.Contract.IsCollectionAllowed(&_CollectionRegistry.CallOpts, collectionId, user)
}

// IsCollectionAllowedAt is a free data retrieval call binding the contract method 0xf53fb0cb.
//
// Solidity: function isCollectionAllowedAt(collectionId bytes8, user bytes8, blockNumber uint256) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) IsCollectionAllowedAt(opts *bind.CallOpts, collectionId [8]byte, user [8]byte, blockNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "isCollectionAllowedAt", collectionId, user, blockNumber)
	return *ret0, err
}

// IsCollectionAllowedAt is a free data retrieval call binding the contract method 0xf53fb0cb.
//
// Solidity: function isCollectionAllowedAt(collectionId bytes8, user bytes8, blockNumber uint256) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) IsCollectionAllowedAt(collectionId [8]byte, user [8]byte, blockNumber *big.Int) (bool, error) {
	return _CollectionRegistry.Contract.IsCollectionAllowedAt(&_CollectionRegistry.CallOpts, collectionId, user, blockNumber)
}

// IsCollectionAllowedAt is a free data retrieval call binding the contract method 0xf53fb0cb.
//
// Solidity: function isCollectionAllowedAt(collectionId bytes8, user bytes8, blockNumber uint256) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) IsCollectionAllowedAt(collectionId [8]byte, user [8]byte, blockNumber *big.Int) (bool, error) {
	return _CollectionRegistry.Contract.IsCollectionAllowedAt(&_CollectionRegistry.CallOpts, collectionId, user, blockNumber)
}

// Allow is a paid mutator transaction binding the contract method 0x4c9b30b4.
//
// Solidity: function allow(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Allow(opts *bind.TransactOpts, _id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "allow", _id)
}

// Allow is a paid mutator transaction binding the contract method 0x4c9b30b4.
//
// Solidity: function allow(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistrySession) Allow(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Allow(&_CollectionRegistry.TransactOpts, _id)
}

// Allow is a paid mutator transaction binding the contract method 0x4c9b30b4.
//
// Solidity: function allow(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Allow(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Allow(&_CollectionRegistry.TransactOpts, _id)
}

// AllowByPassword is a paid mutator transaction binding the contract method 0x4a91ee2a.
//
// Solidity: function allowByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) AllowByPassword(opts *bind.TransactOpts, _id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "allowByPassword", _id, passwordSignature)
}

// AllowByPassword is a paid mutator transaction binding the contract method 0x4a91ee2a.
//
// Solidity: function allowByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistrySession) AllowByPassword(_id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.AllowByPassword(&_CollectionRegistry.TransactOpts, _id, passwordSignature)
}

// AllowByPassword is a paid mutator transaction binding the contract method 0x4a91ee2a.
//
// Solidity: function allowByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) AllowByPassword(_id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.AllowByPassword(&_CollectionRegistry.TransactOpts, _id, passwordSignature)
}

// Deny is a paid mutator transaction binding the contract method 0x18856694.
//
// Solidity: function deny(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Deny(opts *bind.TransactOpts, _id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "deny", _id)
}

// Deny is a paid mutator transaction binding the contract method 0x18856694.
//
// Solidity: function deny(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistrySession) Deny(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Deny(&_CollectionRegistry.TransactOpts, _id)
}

// Deny is a paid mutator transaction binding the contract method 0x18856694.
//
// Solidity: function deny(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Deny(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Deny(&_CollectionRegistry.TransactOpts, _id)
}

// DenyByPassword is a paid mutator transaction binding the contract method 0x0c9bb7d2.
//
// Solidity: function denyByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) DenyByPassword(opts *bind.TransactOpts, _id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "denyByPassword", _id, passwordSignature)
}

// DenyByPassword is a paid mutator transaction binding the contract method 0x0c9bb7d2.
//
// Solidity: function denyByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistrySession) DenyByPassword(_id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.DenyByPassword(&_CollectionRegistry.TransactOpts, _id, passwordSignature)
}

// DenyByPassword is a paid mutator transaction binding the contract method 0x0c9bb7d2.
//
// Solidity: function denyByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) DenyByPassword(_id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.DenyByPassword(&_CollectionRegistry.TransactOpts, _id, passwordSignature)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Register(opts *bind.TransactOpts, _appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "register", _appId, _schemaId, _ratio)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistrySession) Register(_appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Register(&_CollectionRegistry.TransactOpts, _appId, _schemaId, _ratio)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Register(_appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Register(&_CollectionRegistry.TransactOpts, _appId, _schemaId, _ratio)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Unregister(opts *bind.TransactOpts, _id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "unregister", _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistrySession) Unregister(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Unregister(&_CollectionRegistry.TransactOpts, _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Unregister(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Unregister(&_CollectionRegistry.TransactOpts, _id)
}

// CollectionRegistryAllowedIterator is returned from FilterAllowed and is used to iterate over the raw logs and unpacked data for Allowed events raised by the CollectionRegistry contract.
type CollectionRegistryAllowedIterator struct {
	Event *CollectionRegistryAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollectionRegistryAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryAllowed)
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
		it.Event = new(CollectionRegistryAllowed)
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
func (it *CollectionRegistryAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryAllowed represents a Allowed event raised by the CollectionRegistry contract.
type CollectionRegistryAllowed struct {
	CollectionId [8]byte
	UserId       [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowed is a free log retrieval operation binding the contract event 0x2575002f9c19a89406e73df97a2c23c867221b5aa503bd19f5fdc8798f009306.
//
// Solidity: e Allowed(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterAllowed(opts *bind.FilterOpts, collectionId [][8]byte, userId [][8]byte) (*CollectionRegistryAllowedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Allowed", collectionIdRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryAllowedIterator{contract: _CollectionRegistry.contract, event: "Allowed", logs: logs, sub: sub}, nil
}

// FilterAllowed parses the event from given transaction receipt.
//
// Solidity: e Allowed(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseAllowedFromReceipt(receipt *types.Receipt) (*CollectionRegistryAllowed, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x2575002f9c19a89406e73df97a2c23c867221b5aa503bd19f5fdc8798f009306") {
			event := new(CollectionRegistryAllowed)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Allowed", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Allowed event not found")
}

// WatchAllowed is a free log subscription operation binding the contract event 0x2575002f9c19a89406e73df97a2c23c867221b5aa503bd19f5fdc8798f009306.
//
// Solidity: e Allowed(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchAllowed(opts *bind.WatchOpts, sink chan<- *CollectionRegistryAllowed, collectionId [][8]byte, userId [][8]byte) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Allowed", collectionIdRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryAllowed)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Allowed", log); err != nil {
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

// CollectionRegistryDeniedIterator is returned from FilterDenied and is used to iterate over the raw logs and unpacked data for Denied events raised by the CollectionRegistry contract.
type CollectionRegistryDeniedIterator struct {
	Event *CollectionRegistryDenied // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollectionRegistryDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryDenied)
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
		it.Event = new(CollectionRegistryDenied)
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
func (it *CollectionRegistryDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryDenied represents a Denied event raised by the CollectionRegistry contract.
type CollectionRegistryDenied struct {
	CollectionId [8]byte
	UserId       [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDenied is a free log retrieval operation binding the contract event 0x1a57e3d69528db9b16115c4ff4339d855e8468ce95579571daa74bd206d24303.
//
// Solidity: e Denied(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterDenied(opts *bind.FilterOpts, collectionId [][8]byte, userId [][8]byte) (*CollectionRegistryDeniedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Denied", collectionIdRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryDeniedIterator{contract: _CollectionRegistry.contract, event: "Denied", logs: logs, sub: sub}, nil
}

// FilterDenied parses the event from given transaction receipt.
//
// Solidity: e Denied(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseDeniedFromReceipt(receipt *types.Receipt) (*CollectionRegistryDenied, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x1a57e3d69528db9b16115c4ff4339d855e8468ce95579571daa74bd206d24303") {
			event := new(CollectionRegistryDenied)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Denied", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Denied event not found")
}

// WatchDenied is a free log subscription operation binding the contract event 0x1a57e3d69528db9b16115c4ff4339d855e8468ce95579571daa74bd206d24303.
//
// Solidity: e Denied(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchDenied(opts *bind.WatchOpts, sink chan<- *CollectionRegistryDenied, collectionId [][8]byte, userId [][8]byte) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Denied", collectionIdRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryDenied)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Denied", log); err != nil {
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

// CollectionRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the CollectionRegistry contract.
type CollectionRegistryRegistrationIterator struct {
	Event *CollectionRegistryRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollectionRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryRegistration)
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
		it.Event = new(CollectionRegistryRegistration)
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
func (it *CollectionRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryRegistration represents a Registration event raised by the CollectionRegistry contract.
type CollectionRegistryRegistration struct {
	Registrar    common.Address
	AppId        [8]byte
	CollectionId [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xed612afce4032a5821a725a428005a5afc5c47bff7cc9c9b8d0d69e078b133fb.
//
// Solidity: e Registration(registrar indexed address, appId indexed bytes8, collectionId bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, registrar []common.Address, appId [][8]byte) (*CollectionRegistryRegistrationIterator, error) {

	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Registration", registrarRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryRegistrationIterator{contract: _CollectionRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: e Registration(registrar indexed address, appId indexed bytes8, collectionId bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseRegistrationFromReceipt(receipt *types.Receipt) (*CollectionRegistryRegistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xed612afce4032a5821a725a428005a5afc5c47bff7cc9c9b8d0d69e078b133fb") {
			event := new(CollectionRegistryRegistration)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Registration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Registration event not found")
}

// WatchRegistration is a free log subscription operation binding the contract event 0xed612afce4032a5821a725a428005a5afc5c47bff7cc9c9b8d0d69e078b133fb.
//
// Solidity: e Registration(registrar indexed address, appId indexed bytes8, collectionId bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *CollectionRegistryRegistration, registrar []common.Address, appId [][8]byte) (event.Subscription, error) {

	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Registration", registrarRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryRegistration)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
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

// CollectionRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the CollectionRegistry contract.
type CollectionRegistryUnregistrationIterator struct {
	Event *CollectionRegistryUnregistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollectionRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryUnregistration)
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
		it.Event = new(CollectionRegistryUnregistration)
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
func (it *CollectionRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryUnregistration represents a Unregistration event raised by the CollectionRegistry contract.
type CollectionRegistryUnregistration struct {
	CollectionId [8]byte
	AppId        [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x88bf0005675630b29e5b698355f1c09cabdf78e912367fc1850c1d8b33366f2f.
//
// Solidity: e Unregistration(collectionId indexed bytes8, appId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts, collectionId [][8]byte, appId [][8]byte) (*CollectionRegistryUnregistrationIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Unregistration", collectionIdRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryUnregistrationIterator{contract: _CollectionRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: e Unregistration(collectionId indexed bytes8, appId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseUnregistrationFromReceipt(receipt *types.Receipt) (*CollectionRegistryUnregistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x88bf0005675630b29e5b698355f1c09cabdf78e912367fc1850c1d8b33366f2f") {
			event := new(CollectionRegistryUnregistration)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Unregistration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Unregistration event not found")
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x88bf0005675630b29e5b698355f1c09cabdf78e912367fc1850c1d8b33366f2f.
//
// Solidity: e Unregistration(collectionId indexed bytes8, appId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *CollectionRegistryUnregistration, collectionId [][8]byte, appId [][8]byte) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Unregistration", collectionIdRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryUnregistration)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
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
