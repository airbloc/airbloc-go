// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AppRegistryABI is the input ABI used to generate the binding from.
const AppRegistryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes32\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"newOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"checkOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes32\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes32\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AppRegistryBin is the compiled bytecode used for deploying new contracts.
const AppRegistryBin = `0x60806040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631a0919dc146100935780632534fa5d146100c4578063399e0792146101155780634caf58a31461015e578063715018a6146101c75780638da5cb5b146101de578063e1fa8e8414610235578063f2fde38b14610266575b600080fd5b34801561009f57600080fd5b506100c260048036038101908080356000191690602001909291905050506102a9565b005b3480156100d057600080fd5b506101136004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610354565b005b34801561012157600080fd5b506101446004803603810190808035600019169060200190929190505050610459565b604051808215151515815260200191505060405180910390f35b34801561016a57600080fd5b506101ad6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061047d565b604051808215151515815260200191505060405180910390f35b3480156101d357600080fd5b506101dc6104e2565b005b3480156101ea57600080fd5b506101f36105e4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561024157600080fd5b506102646004803603810190808035600019169060200190929190505050610609565b005b34801561027257600080fd5b506102a7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061068a565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561030457600080fd5b6001600082600019166000191681526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055505050565b61035e823361047d565b15156103f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686981526020017f700000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8060016000846000191660001916815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000806001026000191661046c836106f1565b600001546000191614159050919050565b60008173ffffffffffffffffffffffffffffffffffffffff1661049f846106f1565b60010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561053d57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61061281610715565b6001600083600019166000191681526020019081526020016000206000820151816000019060001916905560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106e557600080fd5b6106ee816107d0565b50565b60006001600083600019166000191681526020019081526020016000209050919050565b61071d6108ca565b61072682610459565b151561079a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f61707020616c726561647920657869737473000000000000000000000000000081525060200191505060405180910390fd5b6040805190810160405280836000191681526020013373ffffffffffffffffffffffffffffffffffffffff168152509050919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561080c57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b604080519081016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905600a165627a7a723058208f021979223cdbfadc0d1bda36d75a6d8ac8acb6f4ceb3b033d307ef4305231e0029`

// DeployAppRegistry deploys a new Ethereum contract, binding an instance of AppRegistry to it.
func DeployAppRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AppRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(AppRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AppRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AppRegistry{AppRegistryCaller: AppRegistryCaller{contract: contract}, AppRegistryTransactor: AppRegistryTransactor{contract: contract}, AppRegistryFilterer: AppRegistryFilterer{contract: contract}}, nil
}

// AppRegistry is an auto generated Go binding around an Ethereum contract.
type AppRegistry struct {
	AppRegistryCaller     // Read-only binding to the contract
	AppRegistryTransactor // Write-only binding to the contract
	AppRegistryFilterer   // Log filterer for contract events
}

// AppRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppRegistrySession struct {
	Contract     *AppRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppRegistryCallerSession struct {
	Contract *AppRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AppRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppRegistryTransactorSession struct {
	Contract     *AppRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AppRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRegistryRaw struct {
	Contract *AppRegistry // Generic contract binding to access the raw methods on
}

// AppRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppRegistryCallerRaw struct {
	Contract *AppRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AppRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppRegistryTransactorRaw struct {
	Contract *AppRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppRegistry creates a new instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistry(address common.Address, backend bind.ContractBackend) (*AppRegistry, error) {
	contract, err := bindAppRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppRegistry{AppRegistryCaller: AppRegistryCaller{contract: contract}, AppRegistryTransactor: AppRegistryTransactor{contract: contract}, AppRegistryFilterer: AppRegistryFilterer{contract: contract}}, nil
}

// NewAppRegistryCaller creates a new read-only instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryCaller(address common.Address, caller bind.ContractCaller) (*AppRegistryCaller, error) {
	contract, err := bindAppRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppRegistryCaller{contract: contract}, nil
}

// NewAppRegistryTransactor creates a new write-only instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AppRegistryTransactor, error) {
	contract, err := bindAppRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppRegistryTransactor{contract: contract}, nil
}

// NewAppRegistryFilterer creates a new log filterer instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AppRegistryFilterer, error) {
	contract, err := bindAppRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppRegistryFilterer{contract: contract}, nil
}

// bindAppRegistry binds a generic wrapper to an already deployed contract.
func bindAppRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
func (_AppRegistry *AppRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppRegistry *AppRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AppRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppRegistry *AppRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppRegistry.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_appId bytes32) constant returns(bool)
func (_AppRegistry *AppRegistryCaller) Check(opts *bind.CallOpts, _appId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "check", _appId)
	return *ret0, err
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_appId bytes32) constant returns(bool)
func (_AppRegistry *AppRegistrySession) Check(_appId [32]byte) (bool, error) {
	return _AppRegistry.Contract.Check(&_AppRegistry.CallOpts, _appId)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_appId bytes32) constant returns(bool)
func (_AppRegistry *AppRegistryCallerSession) Check(_appId [32]byte) (bool, error) {
	return _AppRegistry.Contract.Check(&_AppRegistry.CallOpts, _appId)
}

// CheckOwner is a free data retrieval call binding the contract method 0x4caf58a3.
//
// Solidity: function checkOwner(_appId bytes32, _owner address) constant returns(bool)
func (_AppRegistry *AppRegistryCaller) CheckOwner(opts *bind.CallOpts, _appId [32]byte, _owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "checkOwner", _appId, _owner)
	return *ret0, err
}

// CheckOwner is a free data retrieval call binding the contract method 0x4caf58a3.
//
// Solidity: function checkOwner(_appId bytes32, _owner address) constant returns(bool)
func (_AppRegistry *AppRegistrySession) CheckOwner(_appId [32]byte, _owner common.Address) (bool, error) {
	return _AppRegistry.Contract.CheckOwner(&_AppRegistry.CallOpts, _appId, _owner)
}

// CheckOwner is a free data retrieval call binding the contract method 0x4caf58a3.
//
// Solidity: function checkOwner(_appId bytes32, _owner address) constant returns(bool)
func (_AppRegistry *AppRegistryCallerSession) CheckOwner(_appId [32]byte, _owner common.Address) (bool, error) {
	return _AppRegistry.Contract.CheckOwner(&_AppRegistry.CallOpts, _appId, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AppRegistry *AppRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AppRegistry *AppRegistrySession) Owner() (common.Address, error) {
	return _AppRegistry.Contract.Owner(&_AppRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AppRegistry *AppRegistryCallerSession) Owner() (common.Address, error) {
	return _AppRegistry.Contract.Owner(&_AppRegistry.CallOpts)
}

// NewOwner is a paid mutator transaction binding the contract method 0x2534fa5d.
//
// Solidity: function newOwner(_appId bytes32, _newOwner address) returns()
func (_AppRegistry *AppRegistryTransactor) NewOwner(opts *bind.TransactOpts, _appId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "newOwner", _appId, _newOwner)
}

// NewOwner is a paid mutator transaction binding the contract method 0x2534fa5d.
//
// Solidity: function newOwner(_appId bytes32, _newOwner address) returns()
func (_AppRegistry *AppRegistrySession) NewOwner(_appId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.NewOwner(&_AppRegistry.TransactOpts, _appId, _newOwner)
}

// NewOwner is a paid mutator transaction binding the contract method 0x2534fa5d.
//
// Solidity: function newOwner(_appId bytes32, _newOwner address) returns()
func (_AppRegistry *AppRegistryTransactorSession) NewOwner(_appId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.NewOwner(&_AppRegistry.TransactOpts, _appId, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(_appId bytes32) returns()
func (_AppRegistry *AppRegistryTransactor) Register(opts *bind.TransactOpts, _appId [32]byte) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "register", _appId)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(_appId bytes32) returns()
func (_AppRegistry *AppRegistrySession) Register(_appId [32]byte) (*types.Transaction, error) {
	return _AppRegistry.Contract.Register(&_AppRegistry.TransactOpts, _appId)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(_appId bytes32) returns()
func (_AppRegistry *AppRegistryTransactorSession) Register(_appId [32]byte) (*types.Transaction, error) {
	return _AppRegistry.Contract.Register(&_AppRegistry.TransactOpts, _appId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppRegistry *AppRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppRegistry *AppRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _AppRegistry.Contract.RenounceOwnership(&_AppRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppRegistry *AppRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AppRegistry.Contract.RenounceOwnership(&_AppRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AppRegistry *AppRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AppRegistry *AppRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.TransferOwnership(&_AppRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AppRegistry *AppRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.TransferOwnership(&_AppRegistry.TransactOpts, _newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_appId bytes32) returns()
func (_AppRegistry *AppRegistryTransactor) Unregister(opts *bind.TransactOpts, _appId [32]byte) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "unregister", _appId)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_appId bytes32) returns()
func (_AppRegistry *AppRegistrySession) Unregister(_appId [32]byte) (*types.Transaction, error) {
	return _AppRegistry.Contract.Unregister(&_AppRegistry.TransactOpts, _appId)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_appId bytes32) returns()
func (_AppRegistry *AppRegistryTransactorSession) Unregister(_appId [32]byte) (*types.Transaction, error) {
	return _AppRegistry.Contract.Unregister(&_AppRegistry.TransactOpts, _appId)
}

// AppRegistryOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the AppRegistry contract.
type AppRegistryOwnershipRenouncedIterator struct {
	Event *AppRegistryOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *AppRegistryOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryOwnershipRenounced)
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
		it.Event = new(AppRegistryOwnershipRenounced)
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
func (it *AppRegistryOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryOwnershipRenounced represents a OwnershipRenounced event raised by the AppRegistry contract.
type AppRegistryOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*AppRegistryOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryOwnershipRenouncedIterator{contract: _AppRegistry.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *AppRegistryOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryOwnershipRenounced)
				if err := _AppRegistry.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// AppRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AppRegistry contract.
type AppRegistryOwnershipTransferredIterator struct {
	Event *AppRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AppRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryOwnershipTransferred)
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
		it.Event = new(AppRegistryOwnershipTransferred)
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
func (it *AppRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the AppRegistry contract.
type AppRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AppRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryOwnershipTransferredIterator{contract: _AppRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AppRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryOwnershipTransferred)
				if err := _AppRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
