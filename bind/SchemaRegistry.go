// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SchemaRegistryABI is the input ABI used to generate the binding from.
const SchemaRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SchemaRegistryBin is the compiled bytecode used for deploying new contracts.
const SchemaRegistryBin = `0x608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631a0919dc1461005c578063399e07921461008d578063e1fa8e84146100d6575b600080fd5b34801561006857600080fd5b5061008b6004803603810190808035600019169060200190929190505050610107565b005b34801561009957600080fd5b506100bc6004803603810190808035600019169060200190929190505050610224565b604051808215151515815260200191505060405180910390f35b3480156100e257600080fd5b506101056004803603810190808035600019169060200190929190505050610297565b005b3373ffffffffffffffffffffffffffffffffffffffff16600080836000191660001916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156101e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6f6e6c79206f776e65722063616e20646f20746869730000000000000000000081525060200191505060405180910390fd5b600080826000191660001916815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550565b60008073ffffffffffffffffffffffffffffffffffffffff16600080846000191660001916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b33600080836000191660001916815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820353d0e0ab39ba04eab3ab375abbb86402d91523d6d44864e9ceba795abc350a30029`

// DeploySchemaRegistry deploys a new Ethereum contract, binding an instance of SchemaRegistry to it.
func DeploySchemaRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SchemaRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(SchemaRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SchemaRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SchemaRegistry{SchemaRegistryCaller: SchemaRegistryCaller{contract: contract}, SchemaRegistryTransactor: SchemaRegistryTransactor{contract: contract}, SchemaRegistryFilterer: SchemaRegistryFilterer{contract: contract}}, nil
}

// SchemaRegistry is an auto generated Go binding around an Ethereum contract.
type SchemaRegistry struct {
	SchemaRegistryCaller     // Read-only binding to the contract
	SchemaRegistryTransactor // Write-only binding to the contract
	SchemaRegistryFilterer   // Log filterer for contract events
}

// SchemaRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SchemaRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SchemaRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SchemaRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SchemaRegistrySession struct {
	Contract     *SchemaRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SchemaRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SchemaRegistryCallerSession struct {
	Contract *SchemaRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SchemaRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SchemaRegistryTransactorSession struct {
	Contract     *SchemaRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SchemaRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SchemaRegistryRaw struct {
	Contract *SchemaRegistry // Generic contract binding to access the raw methods on
}

// SchemaRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SchemaRegistryCallerRaw struct {
	Contract *SchemaRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SchemaRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SchemaRegistryTransactorRaw struct {
	Contract *SchemaRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSchemaRegistry creates a new instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistry(address common.Address, backend bind.ContractBackend) (*SchemaRegistry, error) {
	contract, err := bindSchemaRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistry{SchemaRegistryCaller: SchemaRegistryCaller{contract: contract}, SchemaRegistryTransactor: SchemaRegistryTransactor{contract: contract}, SchemaRegistryFilterer: SchemaRegistryFilterer{contract: contract}}, nil
}

// NewSchemaRegistryCaller creates a new read-only instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryCaller(address common.Address, caller bind.ContractCaller) (*SchemaRegistryCaller, error) {
	contract, err := bindSchemaRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryCaller{contract: contract}, nil
}

// NewSchemaRegistryTransactor creates a new write-only instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SchemaRegistryTransactor, error) {
	contract, err := bindSchemaRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryTransactor{contract: contract}, nil
}

// NewSchemaRegistryFilterer creates a new log filterer instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SchemaRegistryFilterer, error) {
	contract, err := bindSchemaRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryFilterer{contract: contract}, nil
}

// bindSchemaRegistry binds a generic wrapper to an already deployed contract.
func bindSchemaRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SchemaRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SchemaRegistry *SchemaRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SchemaRegistry.Contract.SchemaRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SchemaRegistry *SchemaRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.SchemaRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SchemaRegistry *SchemaRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.SchemaRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SchemaRegistry *SchemaRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SchemaRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SchemaRegistry *SchemaRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SchemaRegistry *SchemaRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(id bytes32) constant returns(bool)
func (_SchemaRegistry *SchemaRegistryCaller) Check(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SchemaRegistry.contract.Call(opts, out, "check", id)
	return *ret0, err
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(id bytes32) constant returns(bool)
func (_SchemaRegistry *SchemaRegistrySession) Check(id [32]byte) (bool, error) {
	return _SchemaRegistry.Contract.Check(&_SchemaRegistry.CallOpts, id)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(id bytes32) constant returns(bool)
func (_SchemaRegistry *SchemaRegistryCallerSession) Check(id [32]byte) (bool, error) {
	return _SchemaRegistry.Contract.Check(&_SchemaRegistry.CallOpts, id)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(id bytes32) returns()
func (_SchemaRegistry *SchemaRegistryTransactor) Register(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SchemaRegistry.contract.Transact(opts, "register", id)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(id bytes32) returns()
func (_SchemaRegistry *SchemaRegistrySession) Register(id [32]byte) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Register(&_SchemaRegistry.TransactOpts, id)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(id bytes32) returns()
func (_SchemaRegistry *SchemaRegistryTransactorSession) Register(id [32]byte) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Register(&_SchemaRegistry.TransactOpts, id)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(id bytes32) returns()
func (_SchemaRegistry *SchemaRegistryTransactor) Unregister(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SchemaRegistry.contract.Transact(opts, "unregister", id)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(id bytes32) returns()
func (_SchemaRegistry *SchemaRegistrySession) Unregister(id [32]byte) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Unregister(&_SchemaRegistry.TransactOpts, id)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(id bytes32) returns()
func (_SchemaRegistry *SchemaRegistryTransactorSession) Unregister(id [32]byte) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Unregister(&_SchemaRegistry.TransactOpts, id)
}
