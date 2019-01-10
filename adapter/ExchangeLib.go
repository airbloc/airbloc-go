// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"
	"strings"

	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/blockchain/bind"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	_ = bind.NewKeyedTransactor
	_ = ablCommon.HexToID
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangeLibABI is the input ABI used to generate the binding from.
const ExchangeLibABI = "[]"

// ExchangeLib is an auto generated Go binding around an Ethereum contract.
type ExchangeLib struct {
	Address               common.Address
	ExchangeLibCaller     // Read-only binding to the contract
	ExchangeLibTransactor // Write-only binding to the contract
	ExchangeLibFilterer   // Log filterer for contract events
}

// ExchangeLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeLibSession struct {
	Contract     *ExchangeLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeLibCallerSession struct {
	Contract *ExchangeLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExchangeLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeLibTransactorSession struct {
	Contract     *ExchangeLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExchangeLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeLibRaw struct {
	Contract *ExchangeLib // Generic contract binding to access the raw methods on
}

// ExchangeLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeLibCallerRaw struct {
	Contract *ExchangeLibCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeLibTransactorRaw struct {
	Contract *ExchangeLibTransactor // Generic write-only contract binding to access the raw methods on
}

//
//	type ExchangeLibOfferStatus int8
//
//	const (
//		ExchangeLibOfferStatus_NEUTRAL ExchangeLibOfferStatus = 0
//		ExchangeLibOfferStatus_PENDING ExchangeLibOfferStatus = 1
//		ExchangeLibOfferStatus_REJECTED ExchangeLibOfferStatus = 3
//		ExchangeLibOfferStatus_SETTLED ExchangeLibOfferStatus = 2
//
//	)
//
//	var ExchangeLibOfferStatus_name = map[int8]string{
//		0: "NEUTRAL",
//		1: "PENDING",
//		3: "REJECTED",
//		2: "SETTLED",
//
//	}
//
//	var ExchangeLibOfferStatus_value = map[string]int8{
//		"NEUTRAL": 0,
//		"PENDING": 1,
//		"REJECTED": 3,
//		"SETTLED": 2,
//
//	}
//
//
//
//	type Escrow struct {
//		Addr	common.Address
//		Args	[]byte
//		Sign	[4]byte
//
//	}
//
//	type Offer struct {
//		DataIds
//		Escrow	Escrow
//		From	common.Address
//		Reverted	bool
//		Status	ExchangeLibOfferStatus
//		To	common.Address
//
//	}
//
//	type Orderbook struct {
//		Orders	map[ablCommon.ID]Offer
//
//	}
//

func init() {
	// convenient hacks for blockchain.Client
	blockchain.ContractList["ExchangeLib"] = (&ExchangeLib{}).new

}

// NewExchangeLib creates a new instance of ExchangeLib, bound to a specific deployed contract.
func NewExchangeLib(address common.Address, backend bind.ContractBackend) (*ExchangeLib, error) {
	contract, err := bindExchangeLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeLib{
		Address:               address,
		ExchangeLibCaller:     ExchangeLibCaller{contract: contract},
		ExchangeLibTransactor: ExchangeLibTransactor{contract: contract},
		ExchangeLibFilterer:   ExchangeLibFilterer{contract: contract},
	}, nil
}

// NewExchangeLibCaller creates a new read-only instance of ExchangeLib, bound to a specific deployed contract.
func NewExchangeLibCaller(address common.Address, caller bind.ContractCaller) (*ExchangeLibCaller, error) {
	contract, err := bindExchangeLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeLibCaller{contract: contract}, nil
}

// NewExchangeLibTransactor creates a new write-only instance of ExchangeLib, bound to a specific deployed contract.
func NewExchangeLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeLibTransactor, error) {
	contract, err := bindExchangeLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeLibTransactor{contract: contract}, nil
}

// NewExchangeLibFilterer creates a new log filterer instance of ExchangeLib, bound to a specific deployed contract.
func NewExchangeLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeLibFilterer, error) {
	contract, err := bindExchangeLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeLibFilterer{contract: contract}, nil
}

// bindExchangeLib binds a generic wrapper to an already deployed contract.
func bindExchangeLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ExchangeLib *ExchangeLib) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewExchangeLib(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeLib *ExchangeLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeLib.Contract.ExchangeLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeLib *ExchangeLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeLib.Contract.ExchangeLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeLib *ExchangeLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeLib.Contract.ExchangeLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeLib *ExchangeLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeLib *ExchangeLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeLib *ExchangeLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeLib.Contract.contract.Transact(opts, method, params...)
}
