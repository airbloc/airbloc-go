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

// SimpleContractABI is the input ABI used to generate the binding from.
const SimpleContractABI = "[{\"inputs\":[{\"name\":\"_exchange\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"}],\"name\":\"open\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"}],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SimpleContractBin is the compiled bytecode used for deploying new contracts.
const SimpleContractBin = `0x60806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806339c79e0c14610051578063d3b3f73a14610082575b600080fd5b34801561005d57600080fd5b5061008060048036038101908080356000191690602001909291905050506100a6565b005b6100a4600480360381019080803560001916906020019092919050505061057e565b005b6000806000806000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630b0235be886040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050600060405180830381600087803b15801561014757600080fd5b505af115801561015b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250606081101561018557600080fd5b81019080805190602001909291908051906020019092919080516401000000008111156101b157600080fd5b828101905060208101848111156101c757600080fd5b81518560208202830111640100000000821117156101e457600080fd5b505092919050505050955095508573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161493508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614925083801561025d5750825b15156102d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73686f756c64206861766520617574686f72697479000000000000000000000081525060200191505060405180910390fd5b60016000886000191660001916815260200190815260200160002060000160009054906101000a900460ff168015610331575060016000886000191660001916815260200190815260200160002060000160019054906101000a900460ff165b156104fc576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339c79e0c886040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082600019166000191681526020019150506040805180830381600087803b1580156103cd57600080fd5b505af11580156103e1573d6000803e3d6000fd5b505050506040513d60408110156103f757600080fd5b81019080805190602001909291908051906020019092919050505091509150811515610469578573ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610463573d6000803e3d6000fd5b506104b1565b8473ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156104af573d6000803e3d6000fd5b505b600160008860001916600019168152602001908152602001600020600080820160006101000a81549060ff02191690556000820160016101000a81549060ff02191690555050610575565b831561053d576001806000896000191660001916815260200190815260200160002060000160006101000a81548160ff021916908315150217905550610574565b6001806000896000191660001916815260200190815260200160002060000160016101000a81548160ff0219169083151502179055505b5b50505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635778472a836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050608060405180830381600087803b15801561061857600080fd5b505af115801561062c573d6000803e3d6000fd5b505050506040513d608081101561064257600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919050505050505090508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610719576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73686f756c64206861766520617574686f72697479000000000000000000000081525060200191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639185d08d8360006040519080825280602002602001820160405280156107875781602001602082028038833980820191505090505b503460006040518563ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180856000191660001916815260200180602001848152602001838152602001828103825285818151815260200191508051906020019060200280838360005b838110156108135780820151818401526020810190506107f8565b5050505090500195505050505050600060405180830381600087803b15801561083b57600080fd5b505af115801561084f573d6000803e3d6000fd5b5050505050505600a165627a7a723058207ef482c85edb2eb72359fab8fc44b9d533830caa17da57177ebae69bacc5e88f0029`

// DeploySimpleContract deploys a new Ethereum contract, binding an instance of SimpleContract to it.
func DeploySimpleContract(auth *bind.TransactOpts, backend bind.ContractBackend, _exchange common.Address) (common.Address, *types.Transaction, *SimpleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleContractBin), backend, _exchange)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleContract{SimpleContractCaller: SimpleContractCaller{contract: contract}, SimpleContractTransactor: SimpleContractTransactor{contract: contract}, SimpleContractFilterer: SimpleContractFilterer{contract: contract}}, nil
}

// SimpleContract is an auto generated Go binding around an Ethereum contract.
type SimpleContract struct {
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

// NewSimpleContract creates a new instance of SimpleContract, bound to a specific deployed contract.
func NewSimpleContract(address common.Address, backend bind.ContractBackend) (*SimpleContract, error) {
	contract, err := bindSimpleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleContract{SimpleContractCaller: SimpleContractCaller{contract: contract}, SimpleContractTransactor: SimpleContractTransactor{contract: contract}, SimpleContractFilterer: SimpleContractFilterer{contract: contract}}, nil
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

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(_offerId bytes32) returns()
func (_SimpleContract *SimpleContractTransactor) Close(opts *bind.TransactOpts, _offerId [32]byte) (*types.Transaction, error) {
	return _SimpleContract.contract.Transact(opts, "close", _offerId)
}

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(_offerId bytes32) returns()
func (_SimpleContract *SimpleContractSession) Close(_offerId [32]byte) (*types.Transaction, error) {
	return _SimpleContract.Contract.Close(&_SimpleContract.TransactOpts, _offerId)
}

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(_offerId bytes32) returns()
func (_SimpleContract *SimpleContractTransactorSession) Close(_offerId [32]byte) (*types.Transaction, error) {
	return _SimpleContract.Contract.Close(&_SimpleContract.TransactOpts, _offerId)
}

// Open is a paid mutator transaction binding the contract method 0xd3b3f73a.
//
// Solidity: function open(_offerId bytes32) returns()
func (_SimpleContract *SimpleContractTransactor) Open(opts *bind.TransactOpts, _offerId [32]byte) (*types.Transaction, error) {
	return _SimpleContract.contract.Transact(opts, "open", _offerId)
}

// Open is a paid mutator transaction binding the contract method 0xd3b3f73a.
//
// Solidity: function open(_offerId bytes32) returns()
func (_SimpleContract *SimpleContractSession) Open(_offerId [32]byte) (*types.Transaction, error) {
	return _SimpleContract.Contract.Open(&_SimpleContract.TransactOpts, _offerId)
}

// Open is a paid mutator transaction binding the contract method 0xd3b3f73a.
//
// Solidity: function open(_offerId bytes32) returns()
func (_SimpleContract *SimpleContractTransactorSession) Open(_offerId [32]byte) (*types.Transaction, error) {
	return _SimpleContract.Contract.Open(&_SimpleContract.TransactOpts, _offerId)
}
