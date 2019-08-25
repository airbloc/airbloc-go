// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"math/big"

	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
)

// ERC20Escrow is an auto generated Go binding around an Ethereum contract.
type ERC20Escrow struct {
	address               common.Address
	txHash                common.Hash
	createdAt             *big.Int
	ERC20EscrowCaller     // Read-only binding to the contract
	ERC20EscrowTransactor // Write-only binding to the contract
	ERC20EscrowFilterer   // Log filterer for contract events
}

// Address is getter method of ERC20Escrow.address
func (_ERC20Escrow *ERC20Escrow) Address() common.Address {
	return _ERC20Escrow.address
}

// TxHash is getter method of ERC20Escrow.txHash
func (_ERC20Escrow *ERC20Escrow) TxHash() common.Hash {
	return _ERC20Escrow.txHash
}

// CreatedAt is getter method of ERC20Escrow.createdAt
func (_ERC20Escrow *ERC20Escrow) CreatedAt() *big.Int {
	return _ERC20Escrow.createdAt
}

// ERC20EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20EscrowSession struct {
	Contract     *ERC20Escrow      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20EscrowRaw struct {
	Contract *ERC20Escrow // Generic contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Escrow *ERC20EscrowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Escrow.Contract.ERC20EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Escrow *ERC20EscrowRaw) Transfer(opts *bind.TransactOpts) (*klayTypes.Transaction, error) {
	return _ERC20Escrow.Contract.ERC20EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Escrow *ERC20EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*klayTypes.Transaction, error) {
	return _ERC20Escrow.Contract.ERC20EscrowTransactor.contract.Transact(opts, method, params...)
}

// ERC20EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20EscrowCallerSession struct {
	Contract *ERC20EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20EscrowCallerRaw struct {
	Contract *ERC20EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Escrow *ERC20EscrowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Escrow.Contract.contract.Call(opts, result, method, params...)
}

// ERC20EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20EscrowTransactorSession struct {
	Contract     *ERC20EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20EscrowTransactorRaw struct {
	Contract *ERC20EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Escrow *ERC20EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*klayTypes.Transaction, error) {
	return _ERC20Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Escrow *ERC20EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*klayTypes.Transaction, error) {
	return _ERC20Escrow.Contract.contract.Transact(opts, method, params...)
}

// ERC20EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (_ERC20Escrow *ERC20EscrowCaller) Convert(opts *bind.CallOpts, sign [4]byte, args []byte, offerId types.ID) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := &[]interface{}{ret0}
	err := _ERC20Escrow.contract.Call(opts, out, "convert", sign, args, offerId)
	return *ret0, err
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (_ERC20Escrow *ERC20EscrowSession) Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error) {
	return _ERC20Escrow.Contract.Convert(&_ERC20Escrow.CallOpts, sign, args, offerId)
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (_ERC20Escrow *ERC20EscrowCallerSession) Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error) {
	return _ERC20Escrow.Contract.Convert(&_ERC20Escrow.CallOpts, sign, args, offerId)
}

// GetTransactSelector is a free data retrieval call binding the contract method 0xc0a79b5b.
//
// Solidity: function getTransactSelector() constant returns(bytes4)
func (_ERC20Escrow *ERC20EscrowCaller) GetTransactSelector(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := &[]interface{}{ret0}
	err := _ERC20Escrow.contract.Call(opts, out, "getTransactSelector")
	return *ret0, err
}

// GetTransactSelector is a free data retrieval call binding the contract method 0xc0a79b5b.
//
// Solidity: function getTransactSelector() constant returns(bytes4)
func (_ERC20Escrow *ERC20EscrowSession) GetTransactSelector() ([4]byte, error) {
	return _ERC20Escrow.Contract.GetTransactSelector(&_ERC20Escrow.CallOpts)
}

// GetTransactSelector is a free data retrieval call binding the contract method 0xc0a79b5b.
//
// Solidity: function getTransactSelector() constant returns(bytes4)
func (_ERC20Escrow *ERC20EscrowCallerSession) GetTransactSelector() ([4]byte, error) {
	return _ERC20Escrow.Contract.GetTransactSelector(&_ERC20Escrow.CallOpts)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (_ERC20Escrow *ERC20EscrowTransactor) Transact(opts *bind.TransactOpts, token common.Address, amount *big.Int, offerId types.ID) (*klayTypes.Transaction, error) {
	return _ERC20Escrow.contract.Transact(opts, "transact", token, amount, offerId)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (_ERC20Escrow *ERC20EscrowSession) Transact(token common.Address, amount *big.Int, offerId types.ID) (*klayTypes.Transaction, error) {
	return _ERC20Escrow.Contract.Transact(&_ERC20Escrow.TransactOpts, token, amount, offerId)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (_ERC20Escrow *ERC20EscrowTransactorSession) Transact(token common.Address, amount *big.Int, offerId types.ID) (*klayTypes.Transaction, error) {
	return _ERC20Escrow.Contract.Transact(&_ERC20Escrow.TransactOpts, token, amount, offerId)
}
