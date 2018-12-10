// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"
	"strings"

	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
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
	_ = ablCommon.HexToID
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"escrowIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"toIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fromIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"OfferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"OfferPresented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"OfferSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"OfferRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"_offeror\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_offeree\",\"type\":\"address\"}],\"name\":\"Receipt\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offeror\",\"type\":\"address\"},{\"name\":\"_offeree\",\"type\":\"address\"},{\"name\":\"_escrow\",\"type\":\"address\"},{\"name\":\"_sign\",\"type\":\"bytes4\"},{\"name\":\"_args\",\"type\":\"bytes\"},{\"name\":\"_dataIds\",\"type\":\"bytes16[]\"}],\"name\":\"prepare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"},{\"name\":\"_dataIds\",\"type\":\"bytes16[]\"}],\"name\":\"addDataIds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"close\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offeror\",\"type\":\"address\"}],\"name\":\"getReceiptsByOfferor\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offeree\",\"type\":\"address\"}],\"name\":\"getReceiptsByOfferee\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_escrow\",\"type\":\"address\"}],\"name\":\"getReceiptsByEscrow\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"getOfferCompact\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"getOffer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes16[]\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	Address            common.Address
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

func init() {
	// convenient hacks for blockchain.Client
	blockchain.ContractList["Exchange"] = (&Exchange{}).new
	blockchain.RegisterSelector("0xe6d16fcb", "addDataIds(bytes8,bytes16[])")
	blockchain.RegisterSelector("0x688e8391", "close(bytes8)")
	blockchain.RegisterSelector("0x0cf833fb", "order(bytes8)")
	blockchain.RegisterSelector("0xf35f380b", "prepare(address,address,address,bytes4,bytes,bytes16[])")
	blockchain.RegisterSelector("0x6622e153", "reject(bytes8)")
	blockchain.RegisterSelector("0xa60d9b5f", "settle(bytes8)")

}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{
		Address:            address,
		ExchangeCaller:     ExchangeCaller{contract: contract},
		ExchangeTransactor: ExchangeTransactor{contract: contract},
		ExchangeFilterer:   ExchangeFilterer{contract: contract},
	}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Exchange *Exchange) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewExchange(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// EscrowIndex is a free data retrieval call binding the contract method 0x016e1077.
//
// Solidity: function escrowIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeCaller) EscrowIndex(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "escrowIndex", arg0, arg1)
	return *ret0, err
}

// EscrowIndex is a free data retrieval call binding the contract method 0x016e1077.
//
// Solidity: function escrowIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeSession) EscrowIndex(arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	return _Exchange.Contract.EscrowIndex(&_Exchange.CallOpts, arg0, arg1)
}

// EscrowIndex is a free data retrieval call binding the contract method 0x016e1077.
//
// Solidity: function escrowIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeCallerSession) EscrowIndex(arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	return _Exchange.Contract.EscrowIndex(&_Exchange.CallOpts, arg0, arg1)
}

// FromIndex is a free data retrieval call binding the contract method 0xd32029fa.
//
// Solidity: function fromIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeCaller) FromIndex(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "fromIndex", arg0, arg1)
	return *ret0, err
}

// FromIndex is a free data retrieval call binding the contract method 0xd32029fa.
//
// Solidity: function fromIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeSession) FromIndex(arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	return _Exchange.Contract.FromIndex(&_Exchange.CallOpts, arg0, arg1)
}

// FromIndex is a free data retrieval call binding the contract method 0xd32029fa.
//
// Solidity: function fromIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeCallerSession) FromIndex(arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	return _Exchange.Contract.FromIndex(&_Exchange.CallOpts, arg0, arg1)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(_offerId bytes8) constant returns(address, address, bytes16[], address, bytes4, bytes, uint8, bool)
func (_Exchange *ExchangeCaller) GetOffer(opts *bind.CallOpts, _offerId [8]byte) (common.Address, common.Address, [][16]byte, common.Address, [4]byte, []byte, uint8, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new([][16]byte)
		ret3 = new(common.Address)
		ret4 = new([4]byte)
		ret5 = new([]byte)
		ret6 = new(uint8)
		ret7 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
	}
	err := _Exchange.contract.Call(opts, out, "getOffer", _offerId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, err
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(_offerId bytes8) constant returns(address, address, bytes16[], address, bytes4, bytes, uint8, bool)
func (_Exchange *ExchangeSession) GetOffer(_offerId [8]byte) (common.Address, common.Address, [][16]byte, common.Address, [4]byte, []byte, uint8, bool, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, _offerId)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(_offerId bytes8) constant returns(address, address, bytes16[], address, bytes4, bytes, uint8, bool)
func (_Exchange *ExchangeCallerSession) GetOffer(_offerId [8]byte) (common.Address, common.Address, [][16]byte, common.Address, [4]byte, []byte, uint8, bool, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, _offerId)
}

// GetOfferCompact is a free data retrieval call binding the contract method 0x5bcb452a.
//
// Solidity: function getOfferCompact(_offerId bytes8) constant returns(address, address, address, bool)
func (_Exchange *ExchangeCaller) GetOfferCompact(opts *bind.CallOpts, _offerId [8]byte) (common.Address, common.Address, common.Address, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
		ret3 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Exchange.contract.Call(opts, out, "getOfferCompact", _offerId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetOfferCompact is a free data retrieval call binding the contract method 0x5bcb452a.
//
// Solidity: function getOfferCompact(_offerId bytes8) constant returns(address, address, address, bool)
func (_Exchange *ExchangeSession) GetOfferCompact(_offerId [8]byte) (common.Address, common.Address, common.Address, bool, error) {
	return _Exchange.Contract.GetOfferCompact(&_Exchange.CallOpts, _offerId)
}

// GetOfferCompact is a free data retrieval call binding the contract method 0x5bcb452a.
//
// Solidity: function getOfferCompact(_offerId bytes8) constant returns(address, address, address, bool)
func (_Exchange *ExchangeCallerSession) GetOfferCompact(_offerId [8]byte) (common.Address, common.Address, common.Address, bool, error) {
	return _Exchange.Contract.GetOfferCompact(&_Exchange.CallOpts, _offerId)
}

// GetReceiptsByEscrow is a free data retrieval call binding the contract method 0xd7e1989e.
//
// Solidity: function getReceiptsByEscrow(_escrow address) constant returns(bytes8[])
func (_Exchange *ExchangeCaller) GetReceiptsByEscrow(opts *bind.CallOpts, _escrow common.Address) ([][8]byte, error) {
	var (
		ret0 = new([][8]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getReceiptsByEscrow", _escrow)
	return *ret0, err
}

// GetReceiptsByEscrow is a free data retrieval call binding the contract method 0xd7e1989e.
//
// Solidity: function getReceiptsByEscrow(_escrow address) constant returns(bytes8[])
func (_Exchange *ExchangeSession) GetReceiptsByEscrow(_escrow common.Address) ([][8]byte, error) {
	return _Exchange.Contract.GetReceiptsByEscrow(&_Exchange.CallOpts, _escrow)
}

// GetReceiptsByEscrow is a free data retrieval call binding the contract method 0xd7e1989e.
//
// Solidity: function getReceiptsByEscrow(_escrow address) constant returns(bytes8[])
func (_Exchange *ExchangeCallerSession) GetReceiptsByEscrow(_escrow common.Address) ([][8]byte, error) {
	return _Exchange.Contract.GetReceiptsByEscrow(&_Exchange.CallOpts, _escrow)
}

// GetReceiptsByOfferee is a free data retrieval call binding the contract method 0x17329f19.
//
// Solidity: function getReceiptsByOfferee(_offeree address) constant returns(bytes8[])
func (_Exchange *ExchangeCaller) GetReceiptsByOfferee(opts *bind.CallOpts, _offeree common.Address) ([][8]byte, error) {
	var (
		ret0 = new([][8]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getReceiptsByOfferee", _offeree)
	return *ret0, err
}

// GetReceiptsByOfferee is a free data retrieval call binding the contract method 0x17329f19.
//
// Solidity: function getReceiptsByOfferee(_offeree address) constant returns(bytes8[])
func (_Exchange *ExchangeSession) GetReceiptsByOfferee(_offeree common.Address) ([][8]byte, error) {
	return _Exchange.Contract.GetReceiptsByOfferee(&_Exchange.CallOpts, _offeree)
}

// GetReceiptsByOfferee is a free data retrieval call binding the contract method 0x17329f19.
//
// Solidity: function getReceiptsByOfferee(_offeree address) constant returns(bytes8[])
func (_Exchange *ExchangeCallerSession) GetReceiptsByOfferee(_offeree common.Address) ([][8]byte, error) {
	return _Exchange.Contract.GetReceiptsByOfferee(&_Exchange.CallOpts, _offeree)
}

// GetReceiptsByOfferor is a free data retrieval call binding the contract method 0xfaaf7393.
//
// Solidity: function getReceiptsByOfferor(_offeror address) constant returns(bytes8[])
func (_Exchange *ExchangeCaller) GetReceiptsByOfferor(opts *bind.CallOpts, _offeror common.Address) ([][8]byte, error) {
	var (
		ret0 = new([][8]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getReceiptsByOfferor", _offeror)
	return *ret0, err
}

// GetReceiptsByOfferor is a free data retrieval call binding the contract method 0xfaaf7393.
//
// Solidity: function getReceiptsByOfferor(_offeror address) constant returns(bytes8[])
func (_Exchange *ExchangeSession) GetReceiptsByOfferor(_offeror common.Address) ([][8]byte, error) {
	return _Exchange.Contract.GetReceiptsByOfferor(&_Exchange.CallOpts, _offeror)
}

// GetReceiptsByOfferor is a free data retrieval call binding the contract method 0xfaaf7393.
//
// Solidity: function getReceiptsByOfferor(_offeror address) constant returns(bytes8[])
func (_Exchange *ExchangeCallerSession) GetReceiptsByOfferor(_offeror common.Address) ([][8]byte, error) {
	return _Exchange.Contract.GetReceiptsByOfferor(&_Exchange.CallOpts, _offeror)
}

// ToIndex is a free data retrieval call binding the contract method 0x4dae267e.
//
// Solidity: function toIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeCaller) ToIndex(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "toIndex", arg0, arg1)
	return *ret0, err
}

// ToIndex is a free data retrieval call binding the contract method 0x4dae267e.
//
// Solidity: function toIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeSession) ToIndex(arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	return _Exchange.Contract.ToIndex(&_Exchange.CallOpts, arg0, arg1)
}

// ToIndex is a free data retrieval call binding the contract method 0x4dae267e.
//
// Solidity: function toIndex( address,  uint256) constant returns(bytes8)
func (_Exchange *ExchangeCallerSession) ToIndex(arg0 common.Address, arg1 *big.Int) ([8]byte, error) {
	return _Exchange.Contract.ToIndex(&_Exchange.CallOpts, arg0, arg1)
}

// AddDataIds is a paid mutator transaction binding the contract method 0xe6d16fcb.
//
// Solidity: function addDataIds(_offerId bytes8, _dataIds bytes16[]) returns()
func (_Exchange *ExchangeTransactor) AddDataIds(opts *bind.TransactOpts, _offerId [8]byte, _dataIds [][16]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addDataIds", _offerId, _dataIds)
}

// AddDataIds is a paid mutator transaction binding the contract method 0xe6d16fcb.
//
// Solidity: function addDataIds(_offerId bytes8, _dataIds bytes16[]) returns()
func (_Exchange *ExchangeSession) AddDataIds(_offerId [8]byte, _dataIds [][16]byte) (*types.Transaction, error) {
	return _Exchange.Contract.AddDataIds(&_Exchange.TransactOpts, _offerId, _dataIds)
}

// AddDataIds is a paid mutator transaction binding the contract method 0xe6d16fcb.
//
// Solidity: function addDataIds(_offerId bytes8, _dataIds bytes16[]) returns()
func (_Exchange *ExchangeTransactorSession) AddDataIds(_offerId [8]byte, _dataIds [][16]byte) (*types.Transaction, error) {
	return _Exchange.Contract.AddDataIds(&_Exchange.TransactOpts, _offerId, _dataIds)
}

// Close is a paid mutator transaction binding the contract method 0x688e8391.
//
// Solidity: function close(_offerId bytes8) returns(bool)
func (_Exchange *ExchangeTransactor) Close(opts *bind.TransactOpts, _offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "close", _offerId)
}

// Close is a paid mutator transaction binding the contract method 0x688e8391.
//
// Solidity: function close(_offerId bytes8) returns(bool)
func (_Exchange *ExchangeSession) Close(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Close(&_Exchange.TransactOpts, _offerId)
}

// Close is a paid mutator transaction binding the contract method 0x688e8391.
//
// Solidity: function close(_offerId bytes8) returns(bool)
func (_Exchange *ExchangeTransactorSession) Close(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Close(&_Exchange.TransactOpts, _offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(_offerId bytes8) returns()
func (_Exchange *ExchangeTransactor) Order(opts *bind.TransactOpts, _offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "order", _offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(_offerId bytes8) returns()
func (_Exchange *ExchangeSession) Order(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, _offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(_offerId bytes8) returns()
func (_Exchange *ExchangeTransactorSession) Order(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, _offerId)
}

// Prepare is a paid mutator transaction binding the contract method 0xf35f380b.
//
// Solidity: function prepare(_offeror address, _offeree address, _escrow address, _sign bytes4, _args bytes, _dataIds bytes16[]) returns()
func (_Exchange *ExchangeTransactor) Prepare(opts *bind.TransactOpts, _offeror common.Address, _offeree common.Address, _escrow common.Address, _sign [4]byte, _args []byte, _dataIds [][16]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "prepare", _offeror, _offeree, _escrow, _sign, _args, _dataIds)
}

// Prepare is a paid mutator transaction binding the contract method 0xf35f380b.
//
// Solidity: function prepare(_offeror address, _offeree address, _escrow address, _sign bytes4, _args bytes, _dataIds bytes16[]) returns()
func (_Exchange *ExchangeSession) Prepare(_offeror common.Address, _offeree common.Address, _escrow common.Address, _sign [4]byte, _args []byte, _dataIds [][16]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Prepare(&_Exchange.TransactOpts, _offeror, _offeree, _escrow, _sign, _args, _dataIds)
}

// Prepare is a paid mutator transaction binding the contract method 0xf35f380b.
//
// Solidity: function prepare(_offeror address, _offeree address, _escrow address, _sign bytes4, _args bytes, _dataIds bytes16[]) returns()
func (_Exchange *ExchangeTransactorSession) Prepare(_offeror common.Address, _offeree common.Address, _escrow common.Address, _sign [4]byte, _args []byte, _dataIds [][16]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Prepare(&_Exchange.TransactOpts, _offeror, _offeree, _escrow, _sign, _args, _dataIds)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(_offerId bytes8) returns()
func (_Exchange *ExchangeTransactor) Reject(opts *bind.TransactOpts, _offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "reject", _offerId)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(_offerId bytes8) returns()
func (_Exchange *ExchangeSession) Reject(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Reject(&_Exchange.TransactOpts, _offerId)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(_offerId bytes8) returns()
func (_Exchange *ExchangeTransactorSession) Reject(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Reject(&_Exchange.TransactOpts, _offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(_offerId bytes8) returns()
func (_Exchange *ExchangeTransactor) Settle(opts *bind.TransactOpts, _offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "settle", _offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(_offerId bytes8) returns()
func (_Exchange *ExchangeSession) Settle(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Settle(&_Exchange.TransactOpts, _offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(_offerId bytes8) returns()
func (_Exchange *ExchangeTransactorSession) Settle(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Settle(&_Exchange.TransactOpts, _offerId)
}

// ExchangeOfferPreparedIterator is returned from FilterOfferPrepared and is used to iterate over the raw logs and unpacked data for OfferPrepared events raised by the Exchange contract.
type ExchangeOfferPreparedIterator struct {
	Event *ExchangeOfferPrepared // Event containing the contract specifics and raw log

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
func (it *ExchangeOfferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferPrepared)
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
		it.Event = new(ExchangeOfferPrepared)
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
func (it *ExchangeOfferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferPrepared represents a OfferPrepared event raised by the Exchange contract.
type ExchangeOfferPrepared struct {
	OfferId [8]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferPrepared is a free log retrieval operation binding the contract event 0x03094c3ac453fecfef914ccb2bc5edb724821280c11d2555ff7a8a5147c16350.
//
// Solidity: e OfferPrepared(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) FilterOfferPrepared(opts *bind.FilterOpts, _offerId [][8]byte) (*ExchangeOfferPreparedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferPrepared", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferPreparedIterator{contract: _Exchange.contract, event: "OfferPrepared", logs: logs, sub: sub}, nil
}

// FilterOfferPrepared parses the event from given transaction receipt.
//
// Solidity: e OfferPrepared(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) ParseOfferPreparedFromReceipt(receipt *types.Receipt) (*ExchangeOfferPrepared, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x03094c3ac453fecfef914ccb2bc5edb724821280c11d2555ff7a8a5147c16350") {
			event := new(ExchangeOfferPrepared)
			if err := _Exchange.contract.UnpackLog(event, "OfferPrepared", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferPrepared event not found")
}

// WatchOfferPrepared is a free log subscription operation binding the contract event 0x03094c3ac453fecfef914ccb2bc5edb724821280c11d2555ff7a8a5147c16350.
//
// Solidity: e OfferPrepared(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) WatchOfferPrepared(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPrepared, _offerId [][8]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferPrepared", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferPrepared)
				if err := _Exchange.contract.UnpackLog(event, "OfferPrepared", log); err != nil {
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

// ExchangeOfferPresentedIterator is returned from FilterOfferPresented and is used to iterate over the raw logs and unpacked data for OfferPresented events raised by the Exchange contract.
type ExchangeOfferPresentedIterator struct {
	Event *ExchangeOfferPresented // Event containing the contract specifics and raw log

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
func (it *ExchangeOfferPresentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferPresented)
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
		it.Event = new(ExchangeOfferPresented)
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
func (it *ExchangeOfferPresentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferPresentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferPresented represents a OfferPresented event raised by the Exchange contract.
type ExchangeOfferPresented struct {
	OfferId [8]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferPresented is a free log retrieval operation binding the contract event 0xa4f335e96e13d24e916185e1047d7819b9e10ef587fa3be9e29a023e5ad2d62a.
//
// Solidity: e OfferPresented(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) FilterOfferPresented(opts *bind.FilterOpts, _offerId [][8]byte) (*ExchangeOfferPresentedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferPresented", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferPresentedIterator{contract: _Exchange.contract, event: "OfferPresented", logs: logs, sub: sub}, nil
}

// FilterOfferPresented parses the event from given transaction receipt.
//
// Solidity: e OfferPresented(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) ParseOfferPresentedFromReceipt(receipt *types.Receipt) (*ExchangeOfferPresented, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xa4f335e96e13d24e916185e1047d7819b9e10ef587fa3be9e29a023e5ad2d62a") {
			event := new(ExchangeOfferPresented)
			if err := _Exchange.contract.UnpackLog(event, "OfferPresented", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferPresented event not found")
}

// WatchOfferPresented is a free log subscription operation binding the contract event 0xa4f335e96e13d24e916185e1047d7819b9e10ef587fa3be9e29a023e5ad2d62a.
//
// Solidity: e OfferPresented(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPresented, _offerId [][8]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferPresented", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferPresented)
				if err := _Exchange.contract.UnpackLog(event, "OfferPresented", log); err != nil {
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

// ExchangeOfferRejectedIterator is returned from FilterOfferRejected and is used to iterate over the raw logs and unpacked data for OfferRejected events raised by the Exchange contract.
type ExchangeOfferRejectedIterator struct {
	Event *ExchangeOfferRejected // Event containing the contract specifics and raw log

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
func (it *ExchangeOfferRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferRejected)
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
		it.Event = new(ExchangeOfferRejected)
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
func (it *ExchangeOfferRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferRejected represents a OfferRejected event raised by the Exchange contract.
type ExchangeOfferRejected struct {
	OfferId [8]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferRejected is a free log retrieval operation binding the contract event 0x780cea1692b166ae033a8fe67c3e7dd9f1d520e3c999d1c59b9fda6f6ed372e5.
//
// Solidity: e OfferRejected(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) FilterOfferRejected(opts *bind.FilterOpts, _offerId [][8]byte) (*ExchangeOfferRejectedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferRejected", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferRejectedIterator{contract: _Exchange.contract, event: "OfferRejected", logs: logs, sub: sub}, nil
}

// FilterOfferRejected parses the event from given transaction receipt.
//
// Solidity: e OfferRejected(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) ParseOfferRejectedFromReceipt(receipt *types.Receipt) (*ExchangeOfferRejected, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x780cea1692b166ae033a8fe67c3e7dd9f1d520e3c999d1c59b9fda6f6ed372e5") {
			event := new(ExchangeOfferRejected)
			if err := _Exchange.contract.UnpackLog(event, "OfferRejected", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferRejected event not found")
}

// WatchOfferRejected is a free log subscription operation binding the contract event 0x780cea1692b166ae033a8fe67c3e7dd9f1d520e3c999d1c59b9fda6f6ed372e5.
//
// Solidity: e OfferRejected(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *ExchangeOfferRejected, _offerId [][8]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferRejected", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferRejected)
				if err := _Exchange.contract.UnpackLog(event, "OfferRejected", log); err != nil {
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

// ExchangeOfferSettledIterator is returned from FilterOfferSettled and is used to iterate over the raw logs and unpacked data for OfferSettled events raised by the Exchange contract.
type ExchangeOfferSettledIterator struct {
	Event *ExchangeOfferSettled // Event containing the contract specifics and raw log

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
func (it *ExchangeOfferSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferSettled)
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
		it.Event = new(ExchangeOfferSettled)
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
func (it *ExchangeOfferSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferSettled represents a OfferSettled event raised by the Exchange contract.
type ExchangeOfferSettled struct {
	OfferId [8]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferSettled is a free log retrieval operation binding the contract event 0x8563cd74c7f85b9decc8d88aa698ad179ef37e8224ee11e8ce270d9e3fe3ce28.
//
// Solidity: e OfferSettled(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) FilterOfferSettled(opts *bind.FilterOpts, _offerId [][8]byte) (*ExchangeOfferSettledIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferSettled", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferSettledIterator{contract: _Exchange.contract, event: "OfferSettled", logs: logs, sub: sub}, nil
}

// FilterOfferSettled parses the event from given transaction receipt.
//
// Solidity: e OfferSettled(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) ParseOfferSettledFromReceipt(receipt *types.Receipt) (*ExchangeOfferSettled, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8563cd74c7f85b9decc8d88aa698ad179ef37e8224ee11e8ce270d9e3fe3ce28") {
			event := new(ExchangeOfferSettled)
			if err := _Exchange.contract.UnpackLog(event, "OfferSettled", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferSettled event not found")
}

// WatchOfferSettled is a free log subscription operation binding the contract event 0x8563cd74c7f85b9decc8d88aa698ad179ef37e8224ee11e8ce270d9e3fe3ce28.
//
// Solidity: e OfferSettled(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferSettled, _offerId [][8]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferSettled", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferSettled)
				if err := _Exchange.contract.UnpackLog(event, "OfferSettled", log); err != nil {
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

// ExchangeReceiptIterator is returned from FilterReceipt and is used to iterate over the raw logs and unpacked data for Receipt events raised by the Exchange contract.
type ExchangeReceiptIterator struct {
	Event *ExchangeReceipt // Event containing the contract specifics and raw log

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
func (it *ExchangeReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeReceipt)
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
		it.Event = new(ExchangeReceipt)
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
func (it *ExchangeReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeReceipt represents a Receipt event raised by the Exchange contract.
type ExchangeReceipt struct {
	OfferId [8]byte
	Offeror common.Address
	Offeree common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceipt is a free log retrieval operation binding the contract event 0xc5289be5f34d9c261429f87950e84b7904ca18a5804dc86a8a3cfc119b5aeff9.
//
// Solidity: e Receipt(_offerId indexed bytes8, _offeror indexed address, _offeree indexed address)
func (_Exchange *ExchangeFilterer) FilterReceipt(opts *bind.FilterOpts, _offerId [][8]byte, _offeror []common.Address, _offeree []common.Address) (*ExchangeReceiptIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}
	var _offerorRule []interface{}
	for _, _offerorItem := range _offeror {
		_offerorRule = append(_offerorRule, _offerorItem)
	}
	var _offereeRule []interface{}
	for _, _offereeItem := range _offeree {
		_offereeRule = append(_offereeRule, _offereeItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Receipt", _offerIdRule, _offerorRule, _offereeRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeReceiptIterator{contract: _Exchange.contract, event: "Receipt", logs: logs, sub: sub}, nil
}

// FilterReceipt parses the event from given transaction receipt.
//
// Solidity: e Receipt(_offerId indexed bytes8, _offeror indexed address, _offeree indexed address)
func (_Exchange *ExchangeFilterer) ParseReceiptFromReceipt(receipt *types.Receipt) (*ExchangeReceipt, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xc5289be5f34d9c261429f87950e84b7904ca18a5804dc86a8a3cfc119b5aeff9") {
			event := new(ExchangeReceipt)
			if err := _Exchange.contract.UnpackLog(event, "Receipt", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Receipt event not found")
}

// WatchReceipt is a free log subscription operation binding the contract event 0xc5289be5f34d9c261429f87950e84b7904ca18a5804dc86a8a3cfc119b5aeff9.
//
// Solidity: e Receipt(_offerId indexed bytes8, _offeror indexed address, _offeree indexed address)
func (_Exchange *ExchangeFilterer) WatchReceipt(opts *bind.WatchOpts, sink chan<- *ExchangeReceipt, _offerId [][8]byte, _offeror []common.Address, _offeree []common.Address) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}
	var _offerorRule []interface{}
	for _, _offerorItem := range _offeror {
		_offerorRule = append(_offerorRule, _offerorItem)
	}
	var _offereeRule []interface{}
	for _, _offereeItem := range _offeree {
		_offereeRule = append(_offereeRule, _offereeItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Receipt", _offerIdRule, _offerorRule, _offereeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeReceipt)
				if err := _Exchange.contract.UnpackLog(event, "Receipt", log); err != nil {
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
