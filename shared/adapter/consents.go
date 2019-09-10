// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
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

// ConsentsABI is the input ABI used to generate the binding from.
const ConsentsABI = "[{\"inputs\":[{\"name\":\"accountReg\",\"type\":\"address\"},{\"name\":\"appReg\",\"type\":\"address\"},{\"name\":\"controllerReg\",\"type\":\"address\"},{\"name\":\"dataTypeReg\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"dataType\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"Consented\",\"signature\":\"0x8599a1c756b9cd519b80b172f29a03b19082bf7df728da8456cbcab9eeaba8e3\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"consent\",\"outputs\":[],\"payable\":false,\"signature\":\"0x40561d8d\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"consentByController\",\"outputs\":[],\"payable\":false,\"signature\":\"0x2f928d24\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allowed\",\"type\":\"bool\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"modifyConsentByController\",\"outputs\":[],\"payable\":false,\"signature\":\"0xda5b0f50\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x50615985\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"isAllowedAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x7cdda67c\",\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Consents is an auto generated Go binding around an Ethereum contract.
type Consents struct {
	address            common.Address
	txHash             common.Hash
	createdAt          *big.Int
	ConsentsCaller     // Read-only binding to the contract
	ConsentsTransactor // Write-only binding to the contract
	ConsentsFilterer   // Log filterer for contract events
}

// Address is getter method of Consents.address
func (_Consents *Consents) Address() common.Address {
	return _Consents.address
}

// TxHash is getter method of Consents.txHash
func (_Consents *Consents) TxHash() common.Hash {
	return _Consents.txHash
}

// CreatedAt is getter method of Consents.createdAt
func (_Consents *Consents) CreatedAt() *big.Int {
	return _Consents.createdAt
}

// ConsentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsentsSession struct {
	Contract     *Consents         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsentsRaw struct {
	Contract *Consents // Generic contract binding to access the raw methods on
}

// NewConsents creates a new instance of Consents, bound to a specific deployed contract.
func NewConsents(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (*Consents, error) {
	contract, err := bindConsents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Consents{
		address:            address,
		txHash:             txHash,
		createdAt:          createdAt,
		ConsentsCaller:     ConsentsCaller{contract: contract},
		ConsentsTransactor: ConsentsTransactor{contract: contract},
		ConsentsFilterer:   ConsentsFilterer{contract: contract},
	}, nil
}

// bindConsents binds a generic wrapper to an already deployed contract.
func bindConsents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consents *ConsentsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Consents.Contract.ConsentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consents *ConsentsRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ConsentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consents *ConsentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ConsentsTransactor.contract.Transact(opts, method, params...)
}

// ConsentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsentsCallerSession struct {
	Contract *ConsentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ConsentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsentsCallerRaw struct {
	Contract *ConsentsCaller // Generic read-only contract binding to access the raw methods on
}

// NewConsentsCaller creates a new read-only instance of Consents, bound to a specific deployed contract.
func NewConsentsCaller(address common.Address, caller bind.ContractCaller) (*ConsentsCaller, error) {
	contract, err := bindConsents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsentsCaller{contract: contract}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consents *ConsentsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Consents.Contract.contract.Call(opts, result, method, params...)
}

// ConsentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsentsTransactorSession struct {
	Contract     *ConsentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ConsentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsentsTransactorRaw struct {
	Contract *ConsentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsentsTransactor creates a new write-only instance of Consents, bound to a specific deployed contract.
func NewConsentsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsentsTransactor, error) {
	contract, err := bindConsents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsentsTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consents *ConsentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _Consents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consents *ConsentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _Consents.Contract.contract.Transact(opts, method, params...)
}

// ConsentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewConsentsFilterer creates a new log filterer instance of Consents, bound to a specific deployed contract.
func NewConsentsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsentsFilterer, error) {
	contract, err := bindConsents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsentsFilterer{contract: contract}, nil
}

//go:generate mockgen -source consents.go -destination ./mocks/mock_consents.go -package mocks IConsentsManager,IConsentsContract
type IConsentsManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error)
	IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error)

	// Transact methods
	Consent(ctx context.Context, appName string, action uint8, dataType string, allowed bool) error
	ConsentByController(ctx context.Context, userId types.ID, appName string, action uint8, dataType string, allowed bool) error
	ModifyConsentByController(ctx context.Context, userId types.ID, appName string, action uint8, dataType string, allowed bool, passwordSignature []byte) error

	FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appAddr []common.Address) (*ConsentsConsentedIterator, error)
	WatchConsented(opts *bind.WatchOpts, sink chan<- *ConsentsConsented, action []uint8, userId []types.ID, appAddr []common.Address) (event.Subscription, error)
}

type IConsentsContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int
	Filterer() ConsentsFilterer

	IConsentsCalls
	IConsentsTransacts
	IConsentsEvents
}

// ConsentsContract is contract wrapper struct
type ConsentsContract struct {
	client   blockchain.TxClient
	contract *Consents
	ConsentsFilterer
}

// Address is getter method of Consents.address
func (c *ConsentsContract) Address() common.Address {
	return c.contract.Address()
}

// TxHash is getter method of Consents.txHash
func (c *ConsentsContract) TxHash() common.Hash {
	return c.contract.TxHash()
}

// CreatedAt is getter method of Consents.createdAt
func (c *ConsentsContract) CreatedAt() *big.Int {
	return c.contract.CreatedAt()
}

// Filterer is getter method of Consents.ConsentsFilterer
func (c *ConsentsContract) Filterer() ConsentsFilterer {
	return c.ConsentsFilterer
}

// NewConsentsContract makes new *ConsentsContract struct
func NewConsentsContract(client blockchain.TxClient) IConsentsContract {
	contract := client.GetContract(&Consents{}).(*Consents)
	return &ConsentsContract{
		client:           client,
		contract:         contract,
		ConsentsFilterer: contract.ConsentsFilterer,
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("Consents", (&Consents{}).new)
	blockchain.RegisterSelector("0x40561d8d", "consent(string,uint8,string,bool)")
	blockchain.RegisterSelector("0x2f928d24", "consentByController(bytes8,string,uint8,string,bool)")
	blockchain.RegisterSelector("0xda5b0f50", "modifyConsentByController(bytes8,string,uint8,string,bool,bytes)")
}

func (_Consents *Consents) new(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (interface{}, error) {
	return NewConsents(address, txHash, createdAt, backend)
}

type IConsentsCalls interface {
	IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error)
	IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error)
}

// IsAllowed is a free data retrieval call binding the contract method 0x50615985.
//
// Solidity: function isAllowed(bytes8 userId, string appName, uint8 action, string dataType) constant returns(bool)
func (c *ConsentsContract) IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	return c.contract.IsAllowed(nil, userId, appName, action, dataType)
}

// IsAllowed is a free data retrieval call binding the contract method 0x50615985.
//
// Solidity: function isAllowed(bytes8 userId, string appName, uint8 action, string dataType) constant returns(bool)
func (_Consents *ConsentsCaller) IsAllowed(opts *bind.CallOpts, userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _Consents.contract.Call(opts, out, "isAllowed", userId, appName, action, dataType)
	return *ret0, err
}

// IsAllowed is a free data retrieval call binding the contract method 0x50615985.
//
// Solidity: function isAllowed(bytes8 userId, string appName, uint8 action, string dataType) constant returns(bool)
func (_Consents *ConsentsSession) IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	return _Consents.Contract.IsAllowed(&_Consents.CallOpts, userId, appName, action, dataType)
}

// IsAllowed is a free data retrieval call binding the contract method 0x50615985.
//
// Solidity: function isAllowed(bytes8 userId, string appName, uint8 action, string dataType) constant returns(bool)
func (_Consents *ConsentsCallerSession) IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	return _Consents.Contract.IsAllowed(&_Consents.CallOpts, userId, appName, action, dataType)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x7cdda67c.
//
// Solidity: function isAllowedAt(bytes8 userId, string appName, uint8 action, string dataType, uint256 blockNumber) constant returns(bool)
func (c *ConsentsContract) IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	return c.contract.IsAllowedAt(nil, userId, appName, action, dataType, blockNumber)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x7cdda67c.
//
// Solidity: function isAllowedAt(bytes8 userId, string appName, uint8 action, string dataType, uint256 blockNumber) constant returns(bool)
func (_Consents *ConsentsCaller) IsAllowedAt(opts *bind.CallOpts, userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _Consents.contract.Call(opts, out, "isAllowedAt", userId, appName, action, dataType, blockNumber)
	return *ret0, err
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x7cdda67c.
//
// Solidity: function isAllowedAt(bytes8 userId, string appName, uint8 action, string dataType, uint256 blockNumber) constant returns(bool)
func (_Consents *ConsentsSession) IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	return _Consents.Contract.IsAllowedAt(&_Consents.CallOpts, userId, appName, action, dataType, blockNumber)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x7cdda67c.
//
// Solidity: function isAllowedAt(bytes8 userId, string appName, uint8 action, string dataType, uint256 blockNumber) constant returns(bool)
func (_Consents *ConsentsCallerSession) IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	return _Consents.Contract.IsAllowedAt(&_Consents.CallOpts, userId, appName, action, dataType, blockNumber)
}

type IConsentsTransacts interface {
	Consent(ctx context.Context, appName string, action uint8, dataType string, allowed bool) (*ethTypes.Receipt, error)
	ConsentByController(ctx context.Context, userId types.ID, appName string, action uint8, dataType string, allowed bool) (*ethTypes.Receipt, error)
	ModifyConsentByController(ctx context.Context, userId types.ID, appName string, action uint8, dataType string, allowed bool, passwordSignature []byte) (*ethTypes.Receipt, error)
}

// Consent is a paid mutator transaction binding the contract method 0x40561d8d.
//
// Solidity: function consent(string appName, uint8 action, string dataType, bool allowed) returns()
func (c *ConsentsContract) Consent(ctx context.Context, appName string, action uint8, dataType string, allowed bool) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Consent(c.client.Account(), appName, action, dataType, allowed)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Consent is a paid mutator transaction binding the contract method 0x40561d8d.
//
// Solidity: function consent(string appName, uint8 action, string dataType, bool allowed) returns()
func (_Consents *ConsentsTransactor) Consent(opts *bind.TransactOpts, appName string, action uint8, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "consent", appName, action, dataType, allowed)
}

// Consent is a paid mutator transaction binding the contract method 0x40561d8d.
//
// Solidity: function consent(string appName, uint8 action, string dataType, bool allowed) returns()
func (_Consents *ConsentsSession) Consent(appName string, action uint8, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.Contract.Consent(&_Consents.TransactOpts, appName, action, dataType, allowed)
}

// Consent is a paid mutator transaction binding the contract method 0x40561d8d.
//
// Solidity: function consent(string appName, uint8 action, string dataType, bool allowed) returns()
func (_Consents *ConsentsTransactorSession) Consent(appName string, action uint8, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.Contract.Consent(&_Consents.TransactOpts, appName, action, dataType, allowed)
}

// ConsentByController is a paid mutator transaction binding the contract method 0x2f928d24.
//
// Solidity: function consentByController(bytes8 userId, string appName, uint8 action, string dataType, bool allowed) returns()
func (c *ConsentsContract) ConsentByController(ctx context.Context, userId types.ID, appName string, action uint8, dataType string, allowed bool) (*ethTypes.Receipt, error) {
	tx, err := c.contract.ConsentByController(c.client.Account(), userId, appName, action, dataType, allowed)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ConsentByController is a paid mutator transaction binding the contract method 0x2f928d24.
//
// Solidity: function consentByController(bytes8 userId, string appName, uint8 action, string dataType, bool allowed) returns()
func (_Consents *ConsentsTransactor) ConsentByController(opts *bind.TransactOpts, userId types.ID, appName string, action uint8, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "consentByController", userId, appName, action, dataType, allowed)
}

// ConsentByController is a paid mutator transaction binding the contract method 0x2f928d24.
//
// Solidity: function consentByController(bytes8 userId, string appName, uint8 action, string dataType, bool allowed) returns()
func (_Consents *ConsentsSession) ConsentByController(userId types.ID, appName string, action uint8, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ConsentByController(&_Consents.TransactOpts, userId, appName, action, dataType, allowed)
}

// ConsentByController is a paid mutator transaction binding the contract method 0x2f928d24.
//
// Solidity: function consentByController(bytes8 userId, string appName, uint8 action, string dataType, bool allowed) returns()
func (_Consents *ConsentsTransactorSession) ConsentByController(userId types.ID, appName string, action uint8, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ConsentByController(&_Consents.TransactOpts, userId, appName, action, dataType, allowed)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xda5b0f50.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, uint8 action, string dataType, bool allowed, bytes passwordSignature) returns()
func (c *ConsentsContract) ModifyConsentByController(ctx context.Context, userId types.ID, appName string, action uint8, dataType string, allowed bool, passwordSignature []byte) (*ethTypes.Receipt, error) {
	tx, err := c.contract.ModifyConsentByController(c.client.Account(), userId, appName, action, dataType, allowed, passwordSignature)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xda5b0f50.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, uint8 action, string dataType, bool allowed, bytes passwordSignature) returns()
func (_Consents *ConsentsTransactor) ModifyConsentByController(opts *bind.TransactOpts, userId types.ID, appName string, action uint8, dataType string, allowed bool, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "modifyConsentByController", userId, appName, action, dataType, allowed, passwordSignature)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xda5b0f50.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, uint8 action, string dataType, bool allowed, bytes passwordSignature) returns()
func (_Consents *ConsentsSession) ModifyConsentByController(userId types.ID, appName string, action uint8, dataType string, allowed bool, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ModifyConsentByController(&_Consents.TransactOpts, userId, appName, action, dataType, allowed, passwordSignature)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xda5b0f50.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, uint8 action, string dataType, bool allowed, bytes passwordSignature) returns()
func (_Consents *ConsentsTransactorSession) ModifyConsentByController(userId types.ID, appName string, action uint8, dataType string, allowed bool, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ModifyConsentByController(&_Consents.TransactOpts, userId, appName, action, dataType, allowed, passwordSignature)
}

type IConsentsEvents interface {
	FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appAddr []common.Address) (*ConsentsConsentedIterator, error)
	ParseConsentedFromReceipt(receipt *ethTypes.Receipt) (*ConsentsConsented, error)
	WatchConsented(opts *bind.WatchOpts, sink chan<- *ConsentsConsented, action []uint8, userId []types.ID, appAddr []common.Address) (event.Subscription, error)
}

// ConsentsConsentedIterator is returned from FilterConsented and is used to iterate over the raw logs and unpacked data for Consented events raised by the Consents contract.
type ConsentsConsentedIterator struct {
	Event *ConsentsConsented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConsentsConsentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsentsConsented)
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
		it.Event = new(ConsentsConsented)
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
func (it *ConsentsConsentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsentsConsentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsentsConsented represents a Consented event raised by the Consents contract.
type ConsentsConsented struct {
	Action   uint8
	UserId   types.ID
	AppAddr  common.Address
	AppName  string
	DataType string
	Allowed  bool
	Raw      ethTypes.Log // Blockchain specific contextual infos
}

// FilterConsented is a free log retrieval operation binding the contract event 0x8599a1c756b9cd519b80b172f29a03b19082bf7df728da8456cbcab9eeaba8e3.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, address indexed appAddr, string appName, string dataType, bool allowed)
func (_Consents *ConsentsFilterer) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appAddr []common.Address) (*ConsentsConsentedIterator, error) {

	var actionRule []interface{}
	for _, actionItem := range action {
		actionRule = append(actionRule, actionItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var appAddrRule []interface{}
	for _, appAddrItem := range appAddr {
		appAddrRule = append(appAddrRule, appAddrItem)
	}

	logs, sub, err := _Consents.contract.FilterLogs(opts, "Consented", actionRule, userIdRule, appAddrRule)
	if err != nil {
		return nil, err
	}
	return &ConsentsConsentedIterator{contract: _Consents.contract, event: "Consented", logs: logs, sub: sub}, nil
}

// FilterConsented parses the event from given transaction receipt.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, address indexed appAddr, string appName, string dataType, bool allowed)
func (manager *ConsentsContract) ParseConsentedFromReceipt(receipt *ethTypes.Receipt) (*ConsentsConsented, error) {
	return manager.contract.ParseConsentedFromReceipt(receipt)
}

// FilterConsented parses the event from given transaction receipt.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, address indexed appAddr, string appName, string dataType, bool allowed)
func (_Consents *ConsentsFilterer) ParseConsentedFromReceipt(receipt *ethTypes.Receipt) (*ConsentsConsented, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8599a1c756b9cd519b80b172f29a03b19082bf7df728da8456cbcab9eeaba8e3") {
			event := new(ConsentsConsented)
			if err := _Consents.contract.UnpackLog(event, "Consented", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Consented event not found")
}

// WatchConsented is a free log subscription operation binding the contract event 0x8599a1c756b9cd519b80b172f29a03b19082bf7df728da8456cbcab9eeaba8e3.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, address indexed appAddr, string appName, string dataType, bool allowed)
func (_Consents *ConsentsFilterer) WatchConsented(opts *bind.WatchOpts, sink chan<- *ConsentsConsented, action []uint8, userId []types.ID, appAddr []common.Address) (event.Subscription, error) {

	var actionRule []interface{}
	for _, actionItem := range action {
		actionRule = append(actionRule, actionItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var appAddrRule []interface{}
	for _, appAddrItem := range appAddr {
		appAddrRule = append(appAddrRule, appAddrItem)
	}

	logs, sub, err := _Consents.contract.WatchLogs(opts, "Consented", actionRule, userIdRule, appAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsentsConsented)
				if err := _Consents.contract.UnpackLog(event, "Consented", log); err != nil {
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
