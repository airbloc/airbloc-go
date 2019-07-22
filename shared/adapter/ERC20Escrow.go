// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.NewKeyedTransactor
	_ = types.HexToID
	_ = common.Big1
	_ = ethTypes.BloomLookup
	_ = event.NewSubscription
)

// ERC20EscrowABI is the input ABI used to generate the binding from.
const ERC20EscrowABI = "{\"Constructor\":{\"Name\":\"\",\"Const\":false,\"Inputs\":[{\"Name\":\"exchangeContract\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":null},\"Methods\":{\"convert\":{\"Name\":\"convert\",\"Const\":true,\"Inputs\":[{\"Name\":\"sign\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":4,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"args\",\"Type\":{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"getTransactSelector\":{\"Name\":\"getTransactSelector\",\"Const\":true,\"Inputs\":[],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":4,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"transact\":{\"Name\":\"transact\",\"Const\":false,\"Inputs\":[{\"Name\":\"token\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"amount\",\"Type\":{\"Elem\":null,\"Kind\":22,\"Type\":{},\"Size\":256,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]}},\"Events\":{}}"

// ERC20Escrow is an auto generated Go binding around an Ethereum contract.
type ERC20Escrow struct {
	Address               common.Address
	ERC20EscrowCaller     // Read-only binding to the contract
	ERC20EscrowTransactor // Write-only binding to the contract
	ERC20EscrowFilterer   // Log filterer for contract events
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

// NewERC20Escrow creates a new instance of ERC20Escrow, bound to a specific deployed contract.
func NewERC20Escrow(address common.Address, backend bind.ContractBackend) (*ERC20Escrow, error) {
	contract, err := bindERC20Escrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Escrow{
		Address:               address,
		ERC20EscrowCaller:     ERC20EscrowCaller{contract: contract},
		ERC20EscrowTransactor: ERC20EscrowTransactor{contract: contract},
		ERC20EscrowFilterer:   ERC20EscrowFilterer{contract: contract},
	}, nil
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
func (_ERC20Escrow *ERC20EscrowRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.ERC20EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Escrow *ERC20EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
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

// NewERC20EscrowCaller creates a new read-only instance of ERC20Escrow, bound to a specific deployed contract.
func NewERC20EscrowCaller(address common.Address, caller bind.ContractCaller) (*ERC20EscrowCaller, error) {
	contract, err := bindERC20Escrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20EscrowCaller{contract: contract}, nil
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

// NewERC20EscrowTransactor creates a new write-only instance of ERC20Escrow, bound to a specific deployed contract.
func NewERC20EscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20EscrowTransactor, error) {
	contract, err := bindERC20Escrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20EscrowTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Escrow *ERC20EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Escrow *ERC20EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.contract.Transact(opts, method, params...)
}

// ERC20EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewERC20EscrowFilterer creates a new log filterer instance of ERC20Escrow, bound to a specific deployed contract.
func NewERC20EscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20EscrowFilterer, error) {
	contract, err := bindERC20Escrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20EscrowFilterer{contract: contract}, nil
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.ContractList["ERC20Escrow"] = (&ERC20Escrow{}).new
	blockchain.RegisterSelector("0x0bd9e0f8", "transact(address,uint256,bytes8)")

}

// bindERC20Escrow binds a generic wrapper to an already deployed contract.
func bindERC20Escrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ERC20Escrow *ERC20Escrow) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewERC20Escrow(address, backend)
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (_ERC20Escrow *ERC20EscrowCaller) Convert(opts *bind.CallOpts, sign [4]byte, args []byte, offerId [8]byte) ([]byte, error) {
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
func (_ERC20Escrow *ERC20EscrowSession) Convert(sign [4]byte, args []byte, offerId [8]byte) ([]byte, error) {
	return _ERC20Escrow.Contract.Convert(&_ERC20Escrow.CallOpts, sign, args, offerId)
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (_ERC20Escrow *ERC20EscrowCallerSession) Convert(sign [4]byte, args []byte, offerId [8]byte) ([]byte, error) {
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
func (_ERC20Escrow *ERC20EscrowTransactor) Transact(opts *bind.TransactOpts, token common.Address, amount *big.Int, offerId [8]byte) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.contract.Transact(opts, "transact", token, amount, offerId)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (_ERC20Escrow *ERC20EscrowSession) Transact(token common.Address, amount *big.Int, offerId [8]byte) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.Transact(&_ERC20Escrow.TransactOpts, token, amount, offerId)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (_ERC20Escrow *ERC20EscrowTransactorSession) Transact(token common.Address, amount *big.Int, offerId [8]byte) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.Transact(&_ERC20Escrow.TransactOpts, token, amount, offerId)
}
