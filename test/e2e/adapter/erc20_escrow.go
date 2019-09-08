// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"math/big"

	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
)

// Erc20EscrowABI is the input ABI used to generate the binding from.
const Erc20EscrowABI = "[{\"inputs\":[{\"name\":\"exchangeContract\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"sign\",\"type\":\"bytes4\"},{\"name\":\"args\",\"type\":\"bytes\"},{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"signature\":\"0xf8411fa9\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"transact\",\"outputs\":[],\"payable\":false,\"signature\":\"0x0bd9e0f8\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransactSelector\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"signature\":\"0xc0a79b5b\",\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Erc20Escrow is an auto generated Go binding around an Ethereum contract.
type Erc20Escrow struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int

	Erc20EscrowCaller     // Read-only binding to the contract
	Erc20EscrowTransactor // Write-only binding to the contract
	Erc20EscrowFilterer   // Log filterer for contract events
}

// Address is getter method of Erc20Escrow.address
func (_Erc20Escrow *Erc20Escrow) Address() common.Address {
	return _Erc20Escrow.address
}

// TxHash is getter method of Erc20Escrow.txHash
func (_Erc20Escrow *Erc20Escrow) TxHash() common.Hash {
	return _Erc20Escrow.txHash
}

// CreatedAt is getter method of Erc20Escrow.createdAt
func (_Erc20Escrow *Erc20Escrow) CreatedAt() *big.Int {
	return _Erc20Escrow.createdAt
}

// Erc20EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20EscrowSession struct {
	Contract     *Erc20Escrow      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20EscrowCallerSession struct {
	Contract *Erc20EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Erc20EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20EscrowTransactorSession struct {
	Contract     *Erc20EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Erc20EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20EscrowRaw struct {
	Contract *Erc20Escrow // Generic contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Escrow *Erc20EscrowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc20Escrow.Contract.Erc20EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Escrow *Erc20EscrowRaw) Transfer(opts *bind.TransactOpts) (*chainTypes.Transaction, error) {
	return _Erc20Escrow.Contract.Erc20EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Escrow *Erc20EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _Erc20Escrow.Contract.Erc20EscrowTransactor.contract.Transact(opts, method, params...)
}

// Erc20EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20EscrowCallerRaw struct {
	Contract *Erc20EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Escrow *Erc20EscrowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc20Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Erc20EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20EscrowTransactorRaw struct {
	Contract *Erc20EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Escrow *Erc20EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*chainTypes.Transaction, error) {
	return _Erc20Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Escrow *Erc20EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _Erc20Escrow.Contract.contract.Transact(opts, method, params...)
}
