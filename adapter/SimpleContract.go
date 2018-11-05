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
const SimpleContractABI = "[{\"inputs\":[{\"name\":\"_exchange\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"open\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SimpleContractBin is the compiled bytecode used for deploying new contracts.
const SimpleContractBin = `0x60806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063688e8391146100515780636d55224814610099575b600080fd5b34801561005d57600080fd5b50610097600480360381019080803577ffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291905050506100d4565b005b6100d2600480360381019080803577ffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610780565b005b60008060008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663107f04b4896040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff19168152602001915050606060405180830381600087803b1580156101a557600080fd5b505af11580156101b9573d6000803e3d6000fd5b505050506040513d60608110156101cf57600080fd5b810190808051906020019092919080519060200190929190805190602001909291905050509650965096503073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614151561029d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f74207468697320636f6e747261637400000000000000000000000000000081525060200191505060405180910390fd5b8673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161493508573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161492508380156103095750825b151561037d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73686f756c64206861766520617574686f72697479000000000000000000000081525060200191505060405180910390fd5b600160008977ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1680156104395750600160008977ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160019054906101000a900460ff165b156106a1576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663688e8391896040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b15801561050457600080fd5b505af1158015610518573d6000803e3d6000fd5b505050506040513d602081101561052e57600080fd5b81019080805190602001909291905050509150600260008977ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff19168152602001908152602001600020549050600160008977ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff19168152602001908152602001600020600080820160006101000a81549060ff02191690556000820160016101000a81549060ff02191690555050811515610654578673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561064e573d6000803e3d6000fd5b5061069c565b8573ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561069a573d6000803e3d6000fd5b505b610776565b83156107105760018060008a77ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160006101000a81548160ff021916908315150217905550610775565b60018060008a77ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160016101000a81548160ff0219169083151502179055505b5b5050505050505050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663107f04b4846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff19168152602001915050606060405180830381600087803b15801561084957600080fd5b505af115801561085d573d6000803e3d6000fd5b505050506040513d606081101561087357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919050505092505091508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610940576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73686f756c64206861766520617574686f72697479000000000000000000000081525060200191505060405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156109e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f74207468697320636f6e747261637400000000000000000000000000000081525060200191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d552248846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff19168152602001915050600060405180830381600087803b158015610aa957600080fd5b505af1158015610abd573d6000803e3d6000fd5b5050505034600260008577ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002081905550604080519081016040528060001515815260200160001515815250600160008577ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548160ff0219169083151502179055509050505050505600a165627a7a7230582034547de5584f490ac66be645a5d6ab207cb2b147f200294767921cc2fd76f7bf0029`

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
