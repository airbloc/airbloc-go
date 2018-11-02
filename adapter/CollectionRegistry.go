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
const CollectionRegistryABI = "[{\"inputs\":[{\"name\":\"_appReg\",\"type\":\"address\"},{\"name\":\"_schemaReg\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_colectionId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_colectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_appId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_schemaId\",\"type\":\"bytes32\"}],\"name\":\"Unregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_collectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"Allowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_collectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"Denied\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes32\"},{\"name\":\"_schemaId\",\"type\":\"bytes32\"},{\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"checkAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CollectionRegistryBin is the compiled bytecode used for deploying new contracts.
const CollectionRegistryBin = `0x608060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631a0919dc146100885780631fed449f146100b9578063399e0792146100f85780634073c0cc146101415780634fe929c21461018a5780638eaa6ac0146101c9578063d42e715514610225575b600080fd5b34801561009457600080fd5b506100b7600480360381019080803560001916906020019092919050505061027c565b005b3480156100c557600080fd5b506100f660048036038101908080356000191690602001909291908035600019169060200190929190505050610544565b005b34801561010457600080fd5b5061012760048036038101908080356000191690602001909291905050506105ce565b604051808215151515815260200191505060405180910390f35b34801561014d57600080fd5b506101886004803603810190808035600019169060200190929190803560001916906020019092919080359060200190929190505050610604565b005b34801561019657600080fd5b506101c7600480360381019080803560001916906020019092919080356000191690602001909291905050506108e2565b005b3480156101d557600080fd5b506101f86004803603810190808035600019169060200190929190505050610963565b60405180836000191660001916815260200182600019166000191681526020019250505060405180910390f35b34801561023157600080fd5b5061026260048036038101908080356000191690602001909291908035600019169060200190929190505050610987565b604051808215151515815260200191505060405180910390f35b610284610c3c565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634caf58a360026000856000191660001916815260200190815260200160002060000154336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561036e57600080fd5b505af1158015610382573d6000803e3d6000fd5b505050506040513d602081101561039857600080fd5b81019080805190602001909291905050501515610443576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686981526020017f700000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60026000836000191660001916815260200190815260200160002060606040519081016040529081600082015460001916600019168152602001600182015460001916600019168152602001600282016040805190810160405290816000820154815260200160018201548152505081525050905060026000836000191660001916815260200190815260200160002060008082016000905560018201600090556002820160008082016000905560018201600090555050505080602001516000191681600001516000191683600019167f50caa0042a588c64a0d615b147438223a7af5d184c29dbf5f9398ba82d56a7eb60405160405180910390a45050565b60016002600084600019166000191681526020019081526020016000206004016000836000191660001916815260200190815260200160002060006101000a81548160ff021916908315150217905550806000191682600019167fa22515132971f50f788da1869934dff79436ef9486db69a7a59731a5fb61689d60405160405180910390a35050565b60008060010260001916600260008460001916600019168152602001908152602001600020600001546000191614159050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634caf58a385336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156106d257600080fd5b505af11580156106e6573d6000803e3d6000fd5b505050506040513d60208110156106fc57600080fd5b810190808051906020019092919050505015156107a7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686981526020017f700000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b83836040516020018083600019166000191681526020018260001916600019168152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310151561081457805182526020820191506020810190506020830392506107ef565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061084e8484846109d6565b60026000836000191660001916815260200190815260200160002060008201518160000190600019169055602082015181600101906000191690556040820151816002016000820151816000015560208201518160010155505090505080600019167f10906fae603eebfac53ddc0f103bee8a044dd7643c425c7a90f921dfa15ef62c60405160405180910390a250505050565b6002600083600019166000191681526020019081526020016000206004016000826000191660001916815260200190815260200160002060006101000a81549060ff0219169055806000191682600019167f4d28190f1b112cd85f6380723c76d76de35cb4a80b5ab017d01f320b25db009e60405160405180910390a35050565b600080600061097184610bff565b9050806000015481600101549250925050915091565b60006002600084600019166000191681526020019081526020016000206004016000836000191660001916815260200190815260200160002060009054906101000a900460ff16905092915050565b6109de610c3c565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663399e0792846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015610a7757600080fd5b505af1158015610a8b573d6000803e3d6000fd5b505050506040513d6020811015610aa157600080fd5b81019080805190602001909291905050501515610b26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c696420736368656d6100000000000000000000000000000000000081525060200191505060405180910390fd5b610b2f836105ce565b1515610ba3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f636f6c6c656374696f6e20616c7265616479206578697374730000000000000081525060200191505060405180910390fd5b60606040519081016040528085600019168152602001846000191681526020016040805190810160405280858152602001610bf08668056bc75e2d63100000610c2390919063ffffffff16565b81525081525090509392505050565b60006002600083600019166000191681526020019081526020016000209050919050565b6000828211151515610c3157fe5b818303905092915050565b6080604051908101604052806000801916815260200160008019168152602001610c64610c6a565b81525090565b6040805190810160405280600081526020016000815250905600a165627a7a72305820f70a6c17651477ce19922981b7155182bf279a8fa7f4e21699c958031d9fcd750029`

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
	CollectionId [32]byte
	Uid          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowed is a free log retrieval operation binding the contract event 0xa22515132971f50f788da1869934dff79436ef9486db69a7a59731a5fb61689d.
//
// Solidity: e Allowed(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterAllowed(opts *bind.FilterOpts, _collectionId [][32]byte, _uid [][32]byte) (*CollectionRegistryAllowedIterator, error) {

	var _collectionIdRule []interface{}
	for _, _collectionIdItem := range _collectionId {
		_collectionIdRule = append(_collectionIdRule, _collectionIdItem)
	}
	var _uidRule []interface{}
	for _, _uidItem := range _uid {
		_uidRule = append(_uidRule, _uidItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Allowed", _collectionIdRule, _uidRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryAllowedIterator{contract: _CollectionRegistry.contract, event: "Allowed", logs: logs, sub: sub}, nil
}

// WatchAllowed is a free log subscription operation binding the contract event 0xa22515132971f50f788da1869934dff79436ef9486db69a7a59731a5fb61689d.
//
// Solidity: e Allowed(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchAllowed(opts *bind.WatchOpts, sink chan<- *CollectionRegistryAllowed, _collectionId [][32]byte, _uid [][32]byte) (event.Subscription, error) {

	var _collectionIdRule []interface{}
	for _, _collectionIdItem := range _collectionId {
		_collectionIdRule = append(_collectionIdRule, _collectionIdItem)
	}
	var _uidRule []interface{}
	for _, _uidItem := range _uid {
		_uidRule = append(_uidRule, _uidItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Allowed", _collectionIdRule, _uidRule)
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
	CollectionId [32]byte
	Uid          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDenied is a free log retrieval operation binding the contract event 0x4d28190f1b112cd85f6380723c76d76de35cb4a80b5ab017d01f320b25db009e.
//
// Solidity: e Denied(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterDenied(opts *bind.FilterOpts, _collectionId [][32]byte, _uid [][32]byte) (*CollectionRegistryDeniedIterator, error) {

	var _collectionIdRule []interface{}
	for _, _collectionIdItem := range _collectionId {
		_collectionIdRule = append(_collectionIdRule, _collectionIdItem)
	}
	var _uidRule []interface{}
	for _, _uidItem := range _uid {
		_uidRule = append(_uidRule, _uidItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Denied", _collectionIdRule, _uidRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryDeniedIterator{contract: _CollectionRegistry.contract, event: "Denied", logs: logs, sub: sub}, nil
}

// WatchDenied is a free log subscription operation binding the contract event 0x4d28190f1b112cd85f6380723c76d76de35cb4a80b5ab017d01f320b25db009e.
//
// Solidity: e Denied(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchDenied(opts *bind.WatchOpts, sink chan<- *CollectionRegistryDenied, _collectionId [][32]byte, _uid [][32]byte) (event.Subscription, error) {

	var _collectionIdRule []interface{}
	for _, _collectionIdItem := range _collectionId {
		_collectionIdRule = append(_collectionIdRule, _collectionIdItem)
	}
	var _uidRule []interface{}
	for _, _uidItem := range _uid {
		_uidRule = append(_uidRule, _uidItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Denied", _collectionIdRule, _uidRule)
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

// CollectionRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the CollectionRegistry contract.
type CollectionRegistryRegisteredIterator struct {
	Event *CollectionRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryRegistered)
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
		it.Event = new(CollectionRegistryRegistered)
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
func (it *CollectionRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryRegistered represents a Registered event raised by the CollectionRegistry contract.
type CollectionRegistryRegistered struct {
	ColectionId [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x10906fae603eebfac53ddc0f103bee8a044dd7643c425c7a90f921dfa15ef62c.
//
// Solidity: e Registered(_colectionId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, _colectionId [][32]byte) (*CollectionRegistryRegisteredIterator, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Registered", _colectionIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryRegisteredIterator{contract: _CollectionRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x10906fae603eebfac53ddc0f103bee8a044dd7643c425c7a90f921dfa15ef62c.
//
// Solidity: e Registered(_colectionId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *CollectionRegistryRegistered, _colectionId [][32]byte) (event.Subscription, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Registered", _colectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryRegistered)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// CollectionRegistryUnregisteredIterator is returned from FilterUnregistered and is used to iterate over the raw logs and unpacked data for Unregistered events raised by the CollectionRegistry contract.
type CollectionRegistryUnregisteredIterator struct {
	Event *CollectionRegistryUnregistered // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryUnregistered)
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
		it.Event = new(CollectionRegistryUnregistered)
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
func (it *CollectionRegistryUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryUnregistered represents a Unregistered event raised by the CollectionRegistry contract.
type CollectionRegistryUnregistered struct {
	ColectionId [32]byte
	AppId       [32]byte
	SchemaId    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnregistered is a free log retrieval operation binding the contract event 0x50caa0042a588c64a0d615b147438223a7af5d184c29dbf5f9398ba82d56a7eb.
//
// Solidity: e Unregistered(_colectionId indexed bytes32, _appId indexed bytes32, _schemaId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterUnregistered(opts *bind.FilterOpts, _colectionId [][32]byte, _appId [][32]byte, _schemaId [][32]byte) (*CollectionRegistryUnregisteredIterator, error) {

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

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Unregistered", _colectionIdRule, _appIdRule, _schemaIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryUnregisteredIterator{contract: _CollectionRegistry.contract, event: "Unregistered", logs: logs, sub: sub}, nil
}

// WatchUnregistered is a free log subscription operation binding the contract event 0x50caa0042a588c64a0d615b147438223a7af5d184c29dbf5f9398ba82d56a7eb.
//
// Solidity: e Unregistered(_colectionId indexed bytes32, _appId indexed bytes32, _schemaId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchUnregistered(opts *bind.WatchOpts, sink chan<- *CollectionRegistryUnregistered, _colectionId [][32]byte, _appId [][32]byte, _schemaId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Unregistered", _colectionIdRule, _appIdRule, _schemaIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryUnregistered)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Unregistered", log); err != nil {
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
