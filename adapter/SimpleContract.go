// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
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

// SimpleContractBin is the compiled bytecode used for deploying new contracts.
const SimpleContractBin = `0x60806040526004361061004b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663688e839181146100505780636d55224814610074575b600080fd5b34801561005c57600080fd5b50610072600160c060020a031960043516610089565b005b610072600160c060020a031960043516610419565b60008054604080517f107f04b4000000000000000000000000000000000000000000000000000000008152600160c060020a0319851660048201529051839283928392839283928392600160a060020a039092169163107f04b49160248082019260609290919082900301818787803b15801561010557600080fd5b505af1158015610119573d6000803e3d6000fd5b505050506040513d606081101561012f57600080fd5b50805160208201516040909201519098509096509450600160a060020a03851630146101a5576040805160e560020a62461bcd02815260206004820152601160248201527f6e6f74207468697320636f6e7472616374000000000000000000000000000000604482015290519081900360640190fd5b33600160a060020a038089168214955087161492508380156101c45750825b151561021a576040805160e560020a62461bcd02815260206004820152601560248201527f73686f756c64206861766520617574686f726974790000000000000000000000604482015290519081900360640190fd5b600160c060020a0319881660009081526001602052604090205460ff1680156102615750600160c060020a03198816600090815260016020526040902054610100900460ff165b156103b65760008054604080517f688e8391000000000000000000000000000000000000000000000000000000008152600160c060020a03198c1660048201529051600160a060020a039092169263688e8391926024808401936020939083900390910190829087803b1580156102d757600080fd5b505af11580156102eb573d6000803e3d6000fd5b505050506040513d602081101561030157600080fd5b5051600160c060020a031989166000908152600260209081526040808320546001909252909120805461ffff19169055909250905081151561037957604051600160a060020a0388169082156108fc029083906000818181858888f19350505050158015610373573d6000803e3d6000fd5b506103b1565b604051600160a060020a0387169082156108fc029083906000818181858888f193505050501580156103af573d6000803e3d6000fd5b505b61040f565b83156103e857600160c060020a031988166000908152600160208190526040909120805460ff1916909117905561040f565b600160c060020a031988166000908152600160205260409020805461ff0019166101001790555b5050505050505050565b60008054604080517f107f04b4000000000000000000000000000000000000000000000000000000008152600160c060020a03198516600482015290518392600160a060020a03169163107f04b491602480830192606092919082900301818787803b15801561048857600080fd5b505af115801561049c573d6000803e3d6000fd5b505050506040513d60608110156104b257600080fd5b508051604090910151909250905033600160a060020a03831614610520576040805160e560020a62461bcd02815260206004820152601560248201527f73686f756c64206861766520617574686f726974790000000000000000000000604482015290519081900360640190fd5b600160a060020a0381163014610580576040805160e560020a62461bcd02815260206004820152601160248201527f6e6f74207468697320636f6e7472616374000000000000000000000000000000604482015290519081900360640190fd5b60008054604080517f6d552248000000000000000000000000000000000000000000000000000000008152600160c060020a0319871660048201529051600160a060020a0390921692636d5522489260248084019382900301818387803b1580156105ea57600080fd5b505af11580156105fe573d6000803e3d6000fd5b505050600160c060020a0319909316600081815260026020908152604080832034905580518082018252838152808301848152948452600190925290912090518154925115156101000261ff001991151560ff1990941693909317169190911790555050505600a165627a7a723058203cd362e05ee8713ce37f656cebe014897133c4f22cec1a1f11edb70b7857078f0029`

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
