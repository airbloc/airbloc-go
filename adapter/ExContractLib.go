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

// ExContractLibABI is the input ABI used to generate the binding from.
const ExContractLibABI = "[]"

// ExContractLib is an auto generated Go binding around an Ethereum contract.
type ExContractLib struct {
	Address                 common.Address
	ExContractLibCaller     // Read-only binding to the contract
	ExContractLibTransactor // Write-only binding to the contract
	ExContractLibFilterer   // Log filterer for contract events
}

// ExContractLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExContractLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExContractLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExContractLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExContractLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExContractLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExContractLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExContractLibSession struct {
	Contract     *ExContractLib    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExContractLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExContractLibCallerSession struct {
	Contract *ExContractLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ExContractLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExContractLibTransactorSession struct {
	Contract     *ExContractLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ExContractLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExContractLibRaw struct {
	Contract *ExContractLib // Generic contract binding to access the raw methods on
}

// ExContractLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExContractLibCallerRaw struct {
	Contract *ExContractLibCaller // Generic read-only contract binding to access the raw methods on
}

// ExContractLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExContractLibTransactorRaw struct {
	Contract *ExContractLibTransactor // Generic write-only contract binding to access the raw methods on
}

func init() {
	blockchain.ContractList["ExContractLib"] = (&ExContractLib{}).new
}

// NewExContractLib creates a new instance of ExContractLib, bound to a specific deployed contract.
func NewExContractLib(address common.Address, backend bind.ContractBackend) (*ExContractLib, error) {
	contract, err := bindExContractLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExContractLib{
		Address:                 address,
		ExContractLibCaller:     ExContractLibCaller{contract: contract},
		ExContractLibTransactor: ExContractLibTransactor{contract: contract},
		ExContractLibFilterer:   ExContractLibFilterer{contract: contract},
	}, nil
}

// NewExContractLibCaller creates a new read-only instance of ExContractLib, bound to a specific deployed contract.
func NewExContractLibCaller(address common.Address, caller bind.ContractCaller) (*ExContractLibCaller, error) {
	contract, err := bindExContractLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExContractLibCaller{contract: contract}, nil
}

// NewExContractLibTransactor creates a new write-only instance of ExContractLib, bound to a specific deployed contract.
func NewExContractLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ExContractLibTransactor, error) {
	contract, err := bindExContractLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExContractLibTransactor{contract: contract}, nil
}

// NewExContractLibFilterer creates a new log filterer instance of ExContractLib, bound to a specific deployed contract.
func NewExContractLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ExContractLibFilterer, error) {
	contract, err := bindExContractLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExContractLibFilterer{contract: contract}, nil
}

// bindExContractLib binds a generic wrapper to an already deployed contract.
func bindExContractLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExContractLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ExContractLib *ExContractLib) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewExContractLib(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExContractLib *ExContractLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExContractLib.Contract.ExContractLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExContractLib *ExContractLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExContractLib.Contract.ExContractLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExContractLib *ExContractLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExContractLib.Contract.ExContractLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExContractLib *ExContractLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExContractLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExContractLib *ExContractLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExContractLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExContractLib *ExContractLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExContractLib.Contract.contract.Transact(opts, method, params...)
}
