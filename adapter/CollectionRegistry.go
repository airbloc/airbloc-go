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

// CollectionRegistryABI is the input ABI used to generate the binding from.
const CollectionRegistryABI = "[{\"inputs\":[{\"name\":\"_appReg\",\"type\":\"address\"},{\"name\":\"_schemaReg\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_colectionId\",\"type\":\"bytes32\"}],\"name\":\"CollectionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_colectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_appId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_schemaId\",\"type\":\"bytes32\"}],\"name\":\"CollectionUnregistered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes32\"},{\"name\":\"_schemaId\",\"type\":\"bytes32\"},{\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"checkAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CollectionRegistryBin is the compiled bytecode used for deploying new contracts.
const CollectionRegistryBin = `0x608060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631a0919dc146100885780631fed449f146100b9578063399e0792146100f85780634073c0cc146101415780634fe929c21461018a5780638eaa6ac0146101c9578063d42e715514610225575b600080fd5b34801561009457600080fd5b506100b7600480360381019080803560001916906020019092919050505061027c565b005b3480156100c557600080fd5b506100f660048036038101908080356000191690602001909291908035600019169060200190929190505050610544565b005b34801561010457600080fd5b506101276004803603810190808035600019169060200190929190505050610598565b604051808215151515815260200191505060405180910390f35b34801561014d57600080fd5b5061018860048036038101908080356000191690602001909291908035600019169060200190929190803590602001909291905050506105ce565b005b34801561019657600080fd5b506101c7600480360381019080803560001916906020019092919080356000191690602001909291905050506108ac565b005b3480156101d557600080fd5b506101f860048036038101908080356000191690602001909291905050506108f7565b60405180836000191660001916815260200182600019166000191681526020019250505060405180910390f35b34801561023157600080fd5b506102626004803603810190808035600019169060200190929190803560001916906020019092919050505061091b565b604051808215151515815260200191505060405180910390f35b610284610bd0565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634caf58a360026000856000191660001916815260200190815260200160002060000154336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561036e57600080fd5b505af1158015610382573d6000803e3d6000fd5b505050506040513d602081101561039857600080fd5b81019080805190602001909291905050501515610443576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686981526020017f700000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60026000836000191660001916815260200190815260200160002060606040519081016040529081600082015460001916600019168152602001600182015460001916600019168152602001600282016040805190810160405290816000820154815260200160018201548152505081525050905060026000836000191660001916815260200190815260200160002060008082016000905560018201600090556002820160008082016000905560018201600090555050505080602001516000191681600001516000191683600019167f0277701dd53272376b18573e43f76924e4a08d264c0118f902d6b88379cfbfde60405160405180910390a45050565b60016002600084600019166000191681526020019081526020016000206004016000836000191660001916815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008060010260001916600260008460001916600019168152602001908152602001600020600001546000191614159050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634caf58a385336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561069c57600080fd5b505af11580156106b0573d6000803e3d6000fd5b505050506040513d60208110156106c657600080fd5b81019080805190602001909291905050501515610771576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686981526020017f700000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b83836040516020018083600019166000191681526020018260001916600019168152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831015156107de57805182526020820191506020810190506020830392506107b9565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061081884848461096a565b60026000836000191660001916815260200190815260200160002060008201518160000190600019169055602082015181600101906000191690556040820151816002016000820151816000015560208201518160010155505090505080600019167f3474cdf426db51552cef6395e6974df6d1ecce737fc85ee3771d8f3b1dc68d9b60405160405180910390a250505050565b6002600083600019166000191681526020019081526020016000206004016000826000191660001916815260200190815260200160002060006101000a81549060ff02191690555050565b600080600061090584610b93565b9050806000015481600101549250925050915091565b60006002600084600019166000191681526020019081526020016000206004016000836000191660001916815260200190815260200160002060009054906101000a900460ff16905092915050565b610972610bd0565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663399e0792846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015610a0b57600080fd5b505af1158015610a1f573d6000803e3d6000fd5b505050506040513d6020811015610a3557600080fd5b81019080805190602001909291905050501515610aba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c696420736368656d6100000000000000000000000000000000000081525060200191505060405180910390fd5b610ac383610598565b1515610b37576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f636f6c6c656374696f6e20616c7265616479206578697374730000000000000081525060200191505060405180910390fd5b60606040519081016040528085600019168152602001846000191681526020016040805190810160405280858152602001610b848668056bc75e2d63100000610bb790919063ffffffff16565b81525081525090509392505050565b60006002600083600019166000191681526020019081526020016000209050919050565b6000828211151515610bc557fe5b818303905092915050565b6080604051908101604052806000801916815260200160008019168152602001610bf8610bfe565b81525090565b6040805190810160405280600081526020016000815250905600a165627a7a72305820db9fcfba8db2b0e301576674850870725aa905d742310c52a6f791a1328b93490029`

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

// Allow is a paid mutator transaction binding the contract method 0x1fed449f.
//
// Solidity: function allow(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Allow(opts *bind.TransactOpts, _id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "allow", _id, _uid)
}

// Allow is a paid mutator transaction binding the contract method 0x1fed449f.
//
// Solidity: function allow(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistrySession) Allow(_id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Allow(&_CollectionRegistry.TransactOpts, _id, _uid)
}

// Allow is a paid mutator transaction binding the contract method 0x1fed449f.
//
// Solidity: function allow(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Allow(_id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Allow(&_CollectionRegistry.TransactOpts, _id, _uid)
}

// Deny is a paid mutator transaction binding the contract method 0x4fe929c2.
//
// Solidity: function deny(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Deny(opts *bind.TransactOpts, _id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "deny", _id, _uid)
}

// Deny is a paid mutator transaction binding the contract method 0x4fe929c2.
//
// Solidity: function deny(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistrySession) Deny(_id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Deny(&_CollectionRegistry.TransactOpts, _id, _uid)
}

// Deny is a paid mutator transaction binding the contract method 0x4fe929c2.
//
// Solidity: function deny(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Deny(_id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Deny(&_CollectionRegistry.TransactOpts, _id, _uid)
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

// CollectionRegistryCollectionRegisteredIterator is returned from FilterCollectionRegistered and is used to iterate over the raw logs and unpacked data for CollectionRegistered events raised by the CollectionRegistry contract.
type CollectionRegistryCollectionRegisteredIterator struct {
	Event *CollectionRegistryCollectionRegistered // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryCollectionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryCollectionRegistered)
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
		it.Event = new(CollectionRegistryCollectionRegistered)
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
func (it *CollectionRegistryCollectionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryCollectionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryCollectionRegistered represents a CollectionRegistered event raised by the CollectionRegistry contract.
type CollectionRegistryCollectionRegistered struct {
	ColectionId [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCollectionRegistered is a free log retrieval operation binding the contract event 0x3474cdf426db51552cef6395e6974df6d1ecce737fc85ee3771d8f3b1dc68d9b.
//
// Solidity: e CollectionRegistered(_colectionId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterCollectionRegistered(opts *bind.FilterOpts, _colectionId [][32]byte) (*CollectionRegistryCollectionRegisteredIterator, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "CollectionRegistered", _colectionIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryCollectionRegisteredIterator{contract: _CollectionRegistry.contract, event: "CollectionRegistered", logs: logs, sub: sub}, nil
}

// WatchCollectionRegistered is a free log subscription operation binding the contract event 0x3474cdf426db51552cef6395e6974df6d1ecce737fc85ee3771d8f3b1dc68d9b.
//
// Solidity: e CollectionRegistered(_colectionId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchCollectionRegistered(opts *bind.WatchOpts, sink chan<- *CollectionRegistryCollectionRegistered, _colectionId [][32]byte) (event.Subscription, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "CollectionRegistered", _colectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryCollectionRegistered)
				if err := _CollectionRegistry.contract.UnpackLog(event, "CollectionRegistered", log); err != nil {
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

// CollectionRegistryCollectionUnregisteredIterator is returned from FilterCollectionUnregistered and is used to iterate over the raw logs and unpacked data for CollectionUnregistered events raised by the CollectionRegistry contract.
type CollectionRegistryCollectionUnregisteredIterator struct {
	Event *CollectionRegistryCollectionUnregistered // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryCollectionUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryCollectionUnregistered)
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
		it.Event = new(CollectionRegistryCollectionUnregistered)
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
func (it *CollectionRegistryCollectionUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryCollectionUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryCollectionUnregistered represents a CollectionUnregistered event raised by the CollectionRegistry contract.
type CollectionRegistryCollectionUnregistered struct {
	ColectionId [32]byte
	AppId       [32]byte
	SchemaId    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCollectionUnregistered is a free log retrieval operation binding the contract event 0x0277701dd53272376b18573e43f76924e4a08d264c0118f902d6b88379cfbfde.
//
// Solidity: e CollectionUnregistered(_colectionId indexed bytes32, _appId indexed bytes32, _schemaId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterCollectionUnregistered(opts *bind.FilterOpts, _colectionId [][32]byte, _appId [][32]byte, _schemaId [][32]byte) (*CollectionRegistryCollectionUnregisteredIterator, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}
	var _appIdRule []interface{}
	for _, _appIdItem := range _appId {
		_appIdRule = append(_appIdRule, _appIdItem)
	}
	var _schemaIdRule []interface{}
	for _, _schemaIdItem := range _schemaId {
		_schemaIdRule = append(_schemaIdRule, _schemaIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "CollectionUnregistered", _colectionIdRule, _appIdRule, _schemaIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryCollectionUnregisteredIterator{contract: _CollectionRegistry.contract, event: "CollectionUnregistered", logs: logs, sub: sub}, nil
}

// WatchCollectionUnregistered is a free log subscription operation binding the contract event 0x0277701dd53272376b18573e43f76924e4a08d264c0118f902d6b88379cfbfde.
//
// Solidity: e CollectionUnregistered(_colectionId indexed bytes32, _appId indexed bytes32, _schemaId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchCollectionUnregistered(opts *bind.WatchOpts, sink chan<- *CollectionRegistryCollectionUnregistered, _colectionId [][32]byte, _appId [][32]byte, _schemaId [][32]byte) (event.Subscription, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}
	var _appIdRule []interface{}
	for _, _appIdItem := range _appId {
		_appIdRule = append(_appIdRule, _appIdItem)
	}
	var _schemaIdRule []interface{}
	for _, _schemaIdItem := range _schemaId {
		_schemaIdRule = append(_schemaIdRule, _schemaIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "CollectionUnregistered", _colectionIdRule, _appIdRule, _schemaIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryCollectionUnregistered)
				if err := _CollectionRegistry.contract.UnpackLog(event, "CollectionUnregistered", log); err != nil {
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
