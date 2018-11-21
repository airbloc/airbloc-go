// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"
	"strings"

	"github.com/airbloc/airbloc-go/blockchain"
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

// ExContractABI is the input ABI used to generate the binding from.
const ExContractABI = "[]"

// ExContract is an auto generated Go binding around an Ethereum contract.
type ExContract struct {
	Address              common.Address
	ExContractCaller     // Read-only binding to the contract
	ExContractTransactor // Write-only binding to the contract
	ExContractFilterer   // Log filterer for contract events
}

// ExContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExContractSession struct {
	Contract     *ExContract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExContractCallerSession struct {
	Contract *ExContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExContractTransactorSession struct {
	Contract     *ExContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExContractRaw struct {
	Contract *ExContract // Generic contract binding to access the raw methods on
}

// ExContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExContractCallerRaw struct {
	Contract *ExContractCaller // Generic read-only contract binding to access the raw methods on
}

// ExContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExContractTransactorRaw struct {
	Contract *ExContractTransactor // Generic write-only contract binding to access the raw methods on
}

func init() {
	blockchain.ContractList["ExContract"] = &ExContract{}.New
}

// NewExContract creates a new instance of ExContract, bound to a specific deployed contract.
func NewExContract(address common.Address, backend bind.ContractBackend) (*ExContract, error) {
	contract, err := bindExContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExContract{
		Address:              address,
		ExContractCaller:     ExContractCaller{contract: contract},
		ExContractTransactor: ExContractTransactor{contract: contract},
		ExContractFilterer:   ExContractFilterer{contract: contract},
	}, nil
}

// NewExContractCaller creates a new read-only instance of ExContract, bound to a specific deployed contract.
func NewExContractCaller(address common.Address, caller bind.ContractCaller) (*ExContractCaller, error) {
	contract, err := bindExContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExContractCaller{contract: contract}, nil
}

// NewExContractTransactor creates a new write-only instance of ExContract, bound to a specific deployed contract.
func NewExContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ExContractTransactor, error) {
	contract, err := bindExContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExContractTransactor{contract: contract}, nil
}

// NewExContractFilterer creates a new log filterer instance of ExContract, bound to a specific deployed contract.
func NewExContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ExContractFilterer, error) {
	contract, err := bindExContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExContractFilterer{contract: contract}, nil
}

// bindExContract binds a generic wrapper to an already deployed contract.
func bindExContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ExContract *ExContract) New(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewExContract(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExContract *ExContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExContract.Contract.ExContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExContract *ExContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExContract.Contract.ExContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExContract *ExContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExContract.Contract.ExContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExContract *ExContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExContract *ExContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExContract *ExContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExContract.Contract.contract.Transact(opts, method, params...)
}
