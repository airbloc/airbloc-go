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
	_ = ablCommon.IDFromString
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_WHITELISTED\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"OfferPresented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"OfferSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"OfferRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"OfferOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"_offeror\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_offeree\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_reverted\",\"type\":\"bool\"}],\"name\":\"OfferClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offeror\",\"type\":\"address\"},{\"name\":\"_offeree\",\"type\":\"address\"},{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"open\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"close\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes8\"}],\"name\":\"getOffer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

	blockchain.RegisterSelector("0x7b9417c8", "addAddressToWhitelist(address)")

	blockchain.RegisterSelector("0xe2ec6ec3", "addAddressesToWhitelist(address[])")

	blockchain.RegisterSelector("0x688e8391", "close(bytes8)")

	blockchain.RegisterSelector("0x6d552248", "open(bytes8)")

	blockchain.RegisterSelector("0x8221d46f", "order(address,address,address)")

	blockchain.RegisterSelector("0x6622e153", "reject(bytes8)")

	blockchain.RegisterSelector("0x286dd3f5", "removeAddressFromWhitelist(address)")

	blockchain.RegisterSelector("0x24953eaa", "removeAddressesFromWhitelist(address[])")

	blockchain.RegisterSelector("0x715018a6", "renounceOwnership()")

	blockchain.RegisterSelector("0xa60d9b5f", "settle(bytes8)")

	blockchain.RegisterSelector("0xf2fde38b", "transferOwnership(address)")

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

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_Exchange *ExchangeCaller) ROLEWHITELISTED(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "ROLE_WHITELISTED")
	return *ret0, err
}

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_Exchange *ExchangeSession) ROLEWHITELISTED() (string, error) {
	return _Exchange.Contract.ROLEWHITELISTED(&_Exchange.CallOpts)
}

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_Exchange *ExchangeCallerSession) ROLEWHITELISTED() (string, error) {
	return _Exchange.Contract.ROLEWHITELISTED(&_Exchange.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Exchange *ExchangeCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _Exchange.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Exchange *ExchangeSession) CheckRole(_operator common.Address, _role string) error {
	return _Exchange.Contract.CheckRole(&_Exchange.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Exchange *ExchangeCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _Exchange.Contract.CheckRole(&_Exchange.CallOpts, _operator, _role)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(_offerId bytes8) constant returns(address, address, address)
func (_Exchange *ExchangeCaller) GetOffer(opts *bind.CallOpts, _offerId [8]byte) (common.Address, common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Exchange.contract.Call(opts, out, "getOffer", _offerId)
	return *ret0, *ret1, *ret2, err
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(_offerId bytes8) constant returns(address, address, address)
func (_Exchange *ExchangeSession) GetOffer(_offerId [8]byte) (common.Address, common.Address, common.Address, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, _offerId)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(_offerId bytes8) constant returns(address, address, address)
func (_Exchange *ExchangeCallerSession) GetOffer(_offerId [8]byte) (common.Address, common.Address, common.Address, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, _offerId)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Exchange *ExchangeCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Exchange *ExchangeSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Exchange.Contract.HasRole(&_Exchange.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Exchange *ExchangeCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Exchange.Contract.HasRole(&_Exchange.CallOpts, _operator, _role)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(_operator address) constant returns(bool)
func (_Exchange *ExchangeCaller) Whitelist(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "whitelist", _operator)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(_operator address) constant returns(bool)
func (_Exchange *ExchangeSession) Whitelist(_operator common.Address) (bool, error) {
	return _Exchange.Contract.Whitelist(&_Exchange.CallOpts, _operator)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(_operator address) constant returns(bool)
func (_Exchange *ExchangeCallerSession) Whitelist(_operator common.Address) (bool, error) {
	return _Exchange.Contract.Whitelist(&_Exchange.CallOpts, _operator)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_operator address) returns()
func (_Exchange *ExchangeTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addAddressToWhitelist", _operator)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_operator address) returns()
func (_Exchange *ExchangeSession) AddAddressToWhitelist(_operator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAddressToWhitelist(&_Exchange.TransactOpts, _operator)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_operator address) returns()
func (_Exchange *ExchangeTransactorSession) AddAddressToWhitelist(_operator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAddressToWhitelist(&_Exchange.TransactOpts, _operator)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addAddressesToWhitelist", _operators)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeSession) AddAddressesToWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAddressesToWhitelist(&_Exchange.TransactOpts, _operators)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeTransactorSession) AddAddressesToWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAddressesToWhitelist(&_Exchange.TransactOpts, _operators)
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

// Open is a paid mutator transaction binding the contract method 0x6d552248.
//
// Solidity: function open(_offerId bytes8) returns()
func (_Exchange *ExchangeTransactor) Open(opts *bind.TransactOpts, _offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "open", _offerId)
}

// Open is a paid mutator transaction binding the contract method 0x6d552248.
//
// Solidity: function open(_offerId bytes8) returns()
func (_Exchange *ExchangeSession) Open(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Open(&_Exchange.TransactOpts, _offerId)
}

// Open is a paid mutator transaction binding the contract method 0x6d552248.
//
// Solidity: function open(_offerId bytes8) returns()
func (_Exchange *ExchangeTransactorSession) Open(_offerId [8]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Open(&_Exchange.TransactOpts, _offerId)
}

// Order is a paid mutator transaction binding the contract method 0x8221d46f.
//
// Solidity: function order(_offeror address, _offeree address, _contract address) returns()
func (_Exchange *ExchangeTransactor) Order(opts *bind.TransactOpts, _offeror common.Address, _offeree common.Address, _contract common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "order", _offeror, _offeree, _contract)
}

// Order is a paid mutator transaction binding the contract method 0x8221d46f.
//
// Solidity: function order(_offeror address, _offeree address, _contract address) returns()
func (_Exchange *ExchangeSession) Order(_offeror common.Address, _offeree common.Address, _contract common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, _offeror, _offeree, _contract)
}

// Order is a paid mutator transaction binding the contract method 0x8221d46f.
//
// Solidity: function order(_offeror address, _offeree address, _contract address) returns()
func (_Exchange *ExchangeTransactorSession) Order(_offeror common.Address, _offeree common.Address, _contract common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, _offeror, _offeree, _contract)
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

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_operator address) returns()
func (_Exchange *ExchangeTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeAddressFromWhitelist", _operator)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_operator address) returns()
func (_Exchange *ExchangeSession) RemoveAddressFromWhitelist(_operator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAddressFromWhitelist(&_Exchange.TransactOpts, _operator)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_operator address) returns()
func (_Exchange *ExchangeTransactorSession) RemoveAddressFromWhitelist(_operator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAddressFromWhitelist(&_Exchange.TransactOpts, _operator)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeAddressesFromWhitelist", _operators)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeSession) RemoveAddressesFromWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAddressesFromWhitelist(&_Exchange.TransactOpts, _operators)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeTransactorSession) RemoveAddressesFromWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAddressesFromWhitelist(&_Exchange.TransactOpts, _operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Exchange *ExchangeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, _newOwner)
}

// ExchangeOfferClosedIterator is returned from FilterOfferClosed and is used to iterate over the raw logs and unpacked data for OfferClosed events raised by the Exchange contract.
type ExchangeOfferClosedIterator struct {
	Event *ExchangeOfferClosed // Event containing the contract specifics and raw log

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
func (it *ExchangeOfferClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferClosed)
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
		it.Event = new(ExchangeOfferClosed)
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
func (it *ExchangeOfferClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferClosed represents a OfferClosed event raised by the Exchange contract.
type ExchangeOfferClosed struct {
	OfferId  [8]byte
	Offeror  common.Address
	Offeree  common.Address
	Reverted bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferClosed is a free log retrieval operation binding the contract event 0xb576186fa17f96f0991d21a162ff79d8c544b056e64be35c6511d366c4647c14.
//
// Solidity: e OfferClosed(_offerId indexed bytes8, _offeror address, _offeree address, _reverted bool)
func (_Exchange *ExchangeFilterer) FilterOfferClosed(opts *bind.FilterOpts, _offerId [][8]byte) (*ExchangeOfferClosedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferClosed", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferClosedIterator{contract: _Exchange.contract, event: "OfferClosed", logs: logs, sub: sub}, nil
}

// FilterOfferClosed parses the event from given transaction receipt.
//
// Solidity: e OfferClosed(_offerId indexed bytes8, _offeror address, _offeree address, _reverted bool)
func (_Exchange *ExchangeFilterer) ParseOfferClosedFromReceipt(receipt *types.Receipt) (*ExchangeOfferClosed, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xb576186fa17f96f0991d21a162ff79d8c544b056e64be35c6511d366c4647c14") {
			event := new(ExchangeOfferClosed)
			if err := _Exchange.contract.UnpackLog(event, "OfferClosed", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferClosed event not found")
}

// WatchOfferClosed is a free log subscription operation binding the contract event 0xb576186fa17f96f0991d21a162ff79d8c544b056e64be35c6511d366c4647c14.
//
// Solidity: e OfferClosed(_offerId indexed bytes8, _offeror address, _offeree address, _reverted bool)
func (_Exchange *ExchangeFilterer) WatchOfferClosed(opts *bind.WatchOpts, sink chan<- *ExchangeOfferClosed, _offerId [][8]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferClosed", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferClosed)
				if err := _Exchange.contract.UnpackLog(event, "OfferClosed", log); err != nil {
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

// ExchangeOfferOpenedIterator is returned from FilterOfferOpened and is used to iterate over the raw logs and unpacked data for OfferOpened events raised by the Exchange contract.
type ExchangeOfferOpenedIterator struct {
	Event *ExchangeOfferOpened // Event containing the contract specifics and raw log

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
func (it *ExchangeOfferOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferOpened)
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
		it.Event = new(ExchangeOfferOpened)
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
func (it *ExchangeOfferOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferOpened represents a OfferOpened event raised by the Exchange contract.
type ExchangeOfferOpened struct {
	OfferId [8]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferOpened is a free log retrieval operation binding the contract event 0xad95aba0b0916a320123c0424d84ac766fc031e506a6bbce9b4402783b589920.
//
// Solidity: e OfferOpened(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) FilterOfferOpened(opts *bind.FilterOpts, _offerId [][8]byte) (*ExchangeOfferOpenedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferOpened", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferOpenedIterator{contract: _Exchange.contract, event: "OfferOpened", logs: logs, sub: sub}, nil
}

// FilterOfferOpened parses the event from given transaction receipt.
//
// Solidity: e OfferOpened(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) ParseOfferOpenedFromReceipt(receipt *types.Receipt) (*ExchangeOfferOpened, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xad95aba0b0916a320123c0424d84ac766fc031e506a6bbce9b4402783b589920") {
			event := new(ExchangeOfferOpened)
			if err := _Exchange.contract.UnpackLog(event, "OfferOpened", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferOpened event not found")
}

// WatchOfferOpened is a free log subscription operation binding the contract event 0xad95aba0b0916a320123c0424d84ac766fc031e506a6bbce9b4402783b589920.
//
// Solidity: e OfferOpened(_offerId indexed bytes8)
func (_Exchange *ExchangeFilterer) WatchOfferOpened(opts *bind.WatchOpts, sink chan<- *ExchangeOfferOpened, _offerId [][8]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferOpened", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferOpened)
				if err := _Exchange.contract.UnpackLog(event, "OfferOpened", log); err != nil {
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
	OfferId  [8]byte
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferPresented is a free log retrieval operation binding the contract event 0xd090216304141d567e88f9d1c28798912d797b8d0f627d9f2a97d4d5922a1b79.
//
// Solidity: e OfferPresented(_offerId indexed bytes8, _contract address)
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
// Solidity: e OfferPresented(_offerId indexed bytes8, _contract address)
func (_Exchange *ExchangeFilterer) ParseOfferPresentedFromReceipt(receipt *types.Receipt) (*ExchangeOfferPresented, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xd090216304141d567e88f9d1c28798912d797b8d0f627d9f2a97d4d5922a1b79") {
			event := new(ExchangeOfferPresented)
			if err := _Exchange.contract.UnpackLog(event, "OfferPresented", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferPresented event not found")
}

// WatchOfferPresented is a free log subscription operation binding the contract event 0xd090216304141d567e88f9d1c28798912d797b8d0f627d9f2a97d4d5922a1b79.
//
// Solidity: e OfferPresented(_offerId indexed bytes8, _contract address)
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

// ExchangeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Exchange contract.
type ExchangeOwnershipRenouncedIterator struct {
	Event *ExchangeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipRenounced)
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
		it.Event = new(ExchangeOwnershipRenounced)
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
func (it *ExchangeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipRenounced represents a OwnershipRenounced event raised by the Exchange contract.
type ExchangeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Exchange *ExchangeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ExchangeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipRenouncedIterator{contract: _Exchange.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// FilterOwnershipRenounced parses the event from given transaction receipt.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Exchange *ExchangeFilterer) ParseOwnershipRenouncedFromReceipt(receipt *types.Receipt) (*ExchangeOwnershipRenounced, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820") {
			event := new(ExchangeOwnershipRenounced)
			if err := _Exchange.contract.UnpackLog(event, "OwnershipRenounced", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OwnershipRenounced event not found")
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Exchange *ExchangeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipRenounced)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// FilterOwnershipTransferred parses the event from given transaction receipt.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Exchange *ExchangeFilterer) ParseOwnershipTransferredFromReceipt(receipt *types.Receipt) (*ExchangeOwnershipTransferred, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0") {
			event := new(ExchangeOwnershipTransferred)
			if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OwnershipTransferred event not found")
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ExchangeRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the Exchange contract.
type ExchangeRoleAddedIterator struct {
	Event *ExchangeRoleAdded // Event containing the contract specifics and raw log

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
func (it *ExchangeRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRoleAdded)
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
		it.Event = new(ExchangeRoleAdded)
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
func (it *ExchangeRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRoleAdded represents a RoleAdded event raised by the Exchange contract.
type ExchangeRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*ExchangeRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeRoleAddedIterator{contract: _Exchange.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// FilterRoleAdded parses the event from given transaction receipt.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) ParseRoleAddedFromReceipt(receipt *types.Receipt) (*ExchangeRoleAdded, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489") {
			event := new(ExchangeRoleAdded)
			if err := _Exchange.contract.UnpackLog(event, "RoleAdded", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("RoleAdded event not found")
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *ExchangeRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRoleAdded)
				if err := _Exchange.contract.UnpackLog(event, "RoleAdded", log); err != nil {
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

// ExchangeRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the Exchange contract.
type ExchangeRoleRemovedIterator struct {
	Event *ExchangeRoleRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRoleRemoved)
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
		it.Event = new(ExchangeRoleRemoved)
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
func (it *ExchangeRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRoleRemoved represents a RoleRemoved event raised by the Exchange contract.
type ExchangeRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*ExchangeRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeRoleRemovedIterator{contract: _Exchange.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// FilterRoleRemoved parses the event from given transaction receipt.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) ParseRoleRemovedFromReceipt(receipt *types.Receipt) (*ExchangeRoleRemoved, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a") {
			event := new(ExchangeRoleRemoved)
			if err := _Exchange.contract.UnpackLog(event, "RoleRemoved", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("RoleRemoved event not found")
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRoleRemoved)
				if err := _Exchange.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
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
