// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CollectionRegistryABI is the input ABI used to generate the binding from.
const CollectionRegistryABI = "[{\"inputs\":[{\"name\":\"_appReg\",\"type\":\"address\"},{\"name\":\"_schemaReg\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes32\"},{\"name\":\"_schemaId\",\"type\":\"bytes32\"},{\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"checkAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CollectionRegistryBin is the compiled bytecode used for deploying new contracts.
const CollectionRegistryBin = `0x60806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631a0919dc14610072578063399e0792146100a35780634073c0cc146100ec5780638eaa6ac014610135578063d42e715514610191575b600080fd5b34801561007e57600080fd5b506100a160048036038101908080356000191690602001909291905050506101e8565b005b3480156100af57600080fd5b506100d260048036038101908080356000191690602001909291905050506103ef565b604051808215151515815260200191505060405180910390f35b3480156100f857600080fd5b506101336004803603810190808035600019169060200190929190803560001916906020019092919080359060200190929190505050610425565b005b34801561014157600080fd5b506101646004803603810190808035600019169060200190929190505050610633565b60405180836000191660001916815260200182600019166000191681526020019250505060405180910390f35b34801561019d57600080fd5b506101ce60048036038101908080356000191690602001909291908035600019169060200190929190505050610657565b604051808215151515815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634caf58a360026000846000191660001916815260200190815260200160002060000154336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156102d257600080fd5b505af11580156102e6573d6000803e3d6000fd5b505050506040513d60208110156102fc57600080fd5b810190808051906020019092919050505015156103a7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686981526020017f700000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60026000826000191660001916815260200190815260200160002060008082016000905560018201600090556002820160008082016000905560018201600090555050505050565b60008060010260001916600260008460001916600019168152602001908152602001600020600001546000191614159050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634caf58a384336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156104f157600080fd5b505af1158015610505573d6000803e3d6000fd5b505050506040513d602081101561051b57600080fd5b810190808051906020019092919050505015156105c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686981526020017f700000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6105d18383836106a6565b600260008460001916600019168152602001908152602001600020600082015181600001906000191690556020820151816001019060001916905560408201518160020160008201518160000155602082015181600101555050905050505050565b6000806000610641846108cf565b9050806000015481600101549250925050915091565b60006002600084600019166000191681526020019081526020016000206004016000836000191660001916815260200190815260200160002060009054906101000a900460ff16905092915050565b6106ae61090c565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663399e0792846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561074757600080fd5b505af115801561075b573d6000803e3d6000fd5b505050506040513d602081101561077157600080fd5b810190808051906020019092919050505015156107f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c696420736368656d6100000000000000000000000000000000000081525060200191505060405180910390fd5b6107ff836103ef565b1515610873576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f636f6c6c656374696f6e20616c7265616479206578697374730000000000000081525060200191505060405180910390fd5b606060405190810160405280856000191681526020018460001916815260200160408051908101604052808581526020016108c08668056bc75e2d631000006108f390919063ffffffff16565b81525081525090509392505050565b60006002600083600019166000191681526020019081526020016000209050919050565b600082821115151561090157fe5b818303905092915050565b608060405190810160405280600080191681526020016000801916815260200161093461093a565b81525090565b6040805190810160405280600081526020016000815250905600a165627a7a72305820a4f0c7d2bba5946e8fe27dec8cf84c8fb684c2dda7a69e3abb1b264cc92a60590029`

// DeployCollectionRegistry deploys a new Ethereum contract, binding an instance of CollectionRegistry to it.
func DeployCollectionRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _appReg common.Address, _schemaReg common.Address) (common.Address, *types.Transaction, *CollectionRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CollectionRegistryBin), backend, _appReg, _schemaReg)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CollectionRegistry{CollectionRegistryCaller: CollectionRegistryCaller{contract: contract}, CollectionRegistryTransactor: CollectionRegistryTransactor{contract: contract}, CollectionRegistryFilterer: CollectionRegistryFilterer{contract: contract}}, nil
}

// CollectionRegistry is an auto generated Go binding around an Ethereum contract.
type CollectionRegistry struct {
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

// NewCollectionRegistry creates a new instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistry(address common.Address, backend bind.ContractBackend) (*CollectionRegistry, error) {
	contract, err := bindCollectionRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistry{CollectionRegistryCaller: CollectionRegistryCaller{contract: contract}, CollectionRegistryTransactor: CollectionRegistryTransactor{contract: contract}, CollectionRegistryFilterer: CollectionRegistryFilterer{contract: contract}}, nil
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

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_id bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) Check(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "check", _id)
	return *ret0, err
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_id bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) Check(_id [32]byte) (bool, error) {
	return _CollectionRegistry.Contract.Check(&_CollectionRegistry.CallOpts, _id)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_id bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) Check(_id [32]byte) (bool, error) {
	return _CollectionRegistry.Contract.Check(&_CollectionRegistry.CallOpts, _id)
}

// CheckAllowed is a free data retrieval call binding the contract method 0xd42e7155.
//
// Solidity: function checkAllowed(_id bytes32, _uid bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) CheckAllowed(opts *bind.CallOpts, _id [32]byte, _uid [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "checkAllowed", _id, _uid)
	return *ret0, err
}

// CheckAllowed is a free data retrieval call binding the contract method 0xd42e7155.
//
// Solidity: function checkAllowed(_id bytes32, _uid bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) CheckAllowed(_id [32]byte, _uid [32]byte) (bool, error) {
	return _CollectionRegistry.Contract.CheckAllowed(&_CollectionRegistry.CallOpts, _id, _uid)
}

// CheckAllowed is a free data retrieval call binding the contract method 0xd42e7155.
//
// Solidity: function checkAllowed(_id bytes32, _uid bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) CheckAllowed(_id [32]byte, _uid [32]byte) (bool, error) {
	return _CollectionRegistry.Contract.CheckAllowed(&_CollectionRegistry.CallOpts, _id, _uid)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_id bytes32) constant returns(bytes32, bytes32)
func (_CollectionRegistry *CollectionRegistryCaller) Get(opts *bind.CallOpts, _id [32]byte) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _CollectionRegistry.contract.Call(opts, out, "get", _id)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_id bytes32) constant returns(bytes32, bytes32)
func (_CollectionRegistry *CollectionRegistrySession) Get(_id [32]byte) ([32]byte, [32]byte, error) {
	return _CollectionRegistry.Contract.Get(&_CollectionRegistry.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_id bytes32) constant returns(bytes32, bytes32)
func (_CollectionRegistry *CollectionRegistryCallerSession) Get(_id [32]byte) ([32]byte, [32]byte, error) {
	return _CollectionRegistry.Contract.Get(&_CollectionRegistry.CallOpts, _id)
}

// Register is a paid mutator transaction binding the contract method 0x4073c0cc.
//
// Solidity: function register(_appId bytes32, _schemaId bytes32, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Register(opts *bind.TransactOpts, _appId [32]byte, _schemaId [32]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "register", _appId, _schemaId, _ratio)
}

// Register is a paid mutator transaction binding the contract method 0x4073c0cc.
//
// Solidity: function register(_appId bytes32, _schemaId bytes32, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistrySession) Register(_appId [32]byte, _schemaId [32]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Register(&_CollectionRegistry.TransactOpts, _appId, _schemaId, _ratio)
}

// Register is a paid mutator transaction binding the contract method 0x4073c0cc.
//
// Solidity: function register(_appId bytes32, _schemaId bytes32, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Register(_appId [32]byte, _schemaId [32]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Register(&_CollectionRegistry.TransactOpts, _appId, _schemaId, _ratio)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_id bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Unregister(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "unregister", _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_id bytes32) returns()
func (_CollectionRegistry *CollectionRegistrySession) Unregister(_id [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Unregister(&_CollectionRegistry.TransactOpts, _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_id bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Unregister(_id [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Unregister(&_CollectionRegistry.TransactOpts, _id)
}
