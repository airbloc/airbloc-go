// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"
	"strings"

	"github.com/airbloc/airbloc-go/blockchain"
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
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SimpleContractABI is the input ABI used to generate the binding from.
const SimpleContractABI = "[{\"inputs\":[{\"name\":\"_exchange\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"open\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SimpleContract is an auto generated Go binding around an Ethereum contract.
type SimpleContract struct {
	Address                  common.Address
	SimpleContractCaller     // Read-only binding to the contract
	SimpleContractTransactor // Write-only binding to the contract
	SimpleContractFilterer   // Log filterer for contract events
}

// SimpleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleContractSession struct {
	Contract     *SimpleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleContractCallerSession struct {
	Contract *SimpleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SimpleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleContractTransactorSession struct {
	Contract     *SimpleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SimpleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleContractRaw struct {
	Contract *SimpleContract // Generic contract binding to access the raw methods on
}

// SimpleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleContractCallerRaw struct {
	Contract *SimpleContractCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleContractTransactorRaw struct {
	Contract *SimpleContractTransactor // Generic write-only contract binding to access the raw methods on
}

func init() {
	blockchain.ContractList["SimpleContract"] = &SimpleContract{}.New
}

// NewSimpleContract creates a new instance of SimpleContract, bound to a specific deployed contract.
func NewSimpleContract(address common.Address, backend bind.ContractBackend) (*SimpleContract, error) {
	contract, err := bindSimpleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleContract{
		Address:                  address,
		SimpleContractCaller:     SimpleContractCaller{contract: contract},
		SimpleContractTransactor: SimpleContractTransactor{contract: contract},
		SimpleContractFilterer:   SimpleContractFilterer{contract: contract},
	}, nil
}

// NewSimpleContractCaller creates a new read-only instance of SimpleContract, bound to a specific deployed contract.
func NewSimpleContractCaller(address common.Address, caller bind.ContractCaller) (*SimpleContractCaller, error) {
	contract, err := bindSimpleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleContractCaller{contract: contract}, nil
}

// NewSimpleContractTransactor creates a new write-only instance of SimpleContract, bound to a specific deployed contract.
func NewSimpleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleContractTransactor, error) {
	contract, err := bindSimpleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleContractTransactor{contract: contract}, nil
}

// NewSimpleContractFilterer creates a new log filterer instance of SimpleContract, bound to a specific deployed contract.
func NewSimpleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleContractFilterer, error) {
	contract, err := bindSimpleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleContractFilterer{contract: contract}, nil
}

// bindSimpleContract binds a generic wrapper to an already deployed contract.
func bindSimpleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_SimpleContract *SimpleContract) New(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewSimpleContract(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleContract *SimpleContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleContract.Contract.SimpleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleContract *SimpleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleContract.Contract.SimpleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleContract *SimpleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleContract.Contract.SimpleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleContract *SimpleContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleContract *SimpleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleContract *SimpleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleContract.Contract.contract.Transact(opts, method, params...)
}

// Close is a paid mutator transaction binding the contract method 0x688e8391.
//
// Solidity: function close(_offerId bytes8) returns()
func (_SimpleContract *SimpleContractTransactor) Close(opts *bind.TransactOpts, _offerId [8]byte) (*types.Transaction, error) {
	return _SimpleContract.contract.Transact(opts, "close", _offerId)
}

// Close is a paid mutator transaction binding the contract method 0x688e8391.
//
// Solidity: function close(_offerId bytes8) returns()
func (_SimpleContract *SimpleContractSession) Close(_offerId [8]byte) (*types.Transaction, error) {
	return _SimpleContract.Contract.Close(&_SimpleContract.TransactOpts, _offerId)
}

// Close is a paid mutator transaction binding the contract method 0x688e8391.
//
// Solidity: function close(_offerId bytes8) returns()
func (_SimpleContract *SimpleContractTransactorSession) Close(_offerId [8]byte) (*types.Transaction, error) {
	return _SimpleContract.Contract.Close(&_SimpleContract.TransactOpts, _offerId)
}

// Open is a paid mutator transaction binding the contract method 0x6d552248.
//
// Solidity: function open(_offerId bytes8) returns()
func (_SimpleContract *SimpleContractTransactor) Open(opts *bind.TransactOpts, _offerId [8]byte) (*types.Transaction, error) {
	return _SimpleContract.contract.Transact(opts, "open", _offerId)
}

// Open is a paid mutator transaction binding the contract method 0x6d552248.
//
// Solidity: function open(_offerId bytes8) returns()
func (_SimpleContract *SimpleContractSession) Open(_offerId [8]byte) (*types.Transaction, error) {
	return _SimpleContract.Contract.Open(&_SimpleContract.TransactOpts, _offerId)
}

// Open is a paid mutator transaction binding the contract method 0x6d552248.
//
// Solidity: function open(_offerId bytes8) returns()
func (_SimpleContract *SimpleContractTransactorSession) Open(_offerId [8]byte) (*types.Transaction, error) {
	return _SimpleContract.Contract.Open(&_SimpleContract.TransactOpts, _offerId)
}
