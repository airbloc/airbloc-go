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
const ConsentsABI = "{\"Constructor\":{\"Name\":\"\",\"Const\":false,\"Inputs\":[{\"Name\":\"accountReg\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"appReg\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"controllerReg\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataTypeReg\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":null},\"Methods\":{\"consent\":{\"Name\":\"consent\",\"Const\":false,\"Inputs\":[{\"Name\":\"action\",\"Type\":{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataType\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"allowed\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"consentByController\":{\"Name\":\"consentByController\",\"Const\":false,\"Inputs\":[{\"Name\":\"action\",\"Type\":{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"userId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataType\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"allowed\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"isAllowed\":{\"Name\":\"isAllowed\",\"Const\":true,\"Inputs\":[{\"Name\":\"action\",\"Type\":{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"userId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataType\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"isAllowedAt\":{\"Name\":\"isAllowedAt\",\"Const\":true,\"Inputs\":[{\"Name\":\"action\",\"Type\":{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"userId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataType\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"blockNumber\",\"Type\":{\"Elem\":null,\"Kind\":22,\"Type\":{},\"Size\":256,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"modifyConsentByController\":{\"Name\":\"modifyConsentByController\",\"Const\":false,\"Inputs\":[{\"Name\":\"action\",\"Type\":{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"userId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataType\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"allowed\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"passwordSignature\",\"Type\":{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]}},\"Events\":{\"Consented\":{\"Name\":\"Consented\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"action\",\"Type\":{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"userId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"app\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"appName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataType\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"allowed\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]}}}"

// Consents is an auto generated Go binding around an Ethereum contract.
type Consents struct {
	Address            common.Address
	ConsentsCaller     // Read-only binding to the contract
	ConsentsTransactor // Write-only binding to the contract
	ConsentsFilterer   // Log filterer for contract events
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
func NewConsents(address common.Address, backend bind.ContractBackend) (*Consents, error) {
	contract, err := bindConsents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Consents{
		Address:            address,
		ConsentsCaller:     ConsentsCaller{contract: contract},
		ConsentsTransactor: ConsentsTransactor{contract: contract},
		ConsentsFilterer:   ConsentsFilterer{contract: contract},
	}, nil
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

type ConsentsManager interface {
	// Pure/View methods
	IsAllowed(action uint8, userId types.ID, appName string, dataType string) (bool, error)
	IsAllowedAt(action uint8, userId types.ID, appName string, dataType string, blockNumber *big.Int) (bool, error)

	// Other methods
	Consent(ctx context.Context, action uint8, appName string, dataType string, allowed bool) error
	ConsentByController(ctx context.Context, action uint8, userId types.ID, appName string, dataType string, allowed bool) error
	ModifyConsentByController(ctx context.Context, action uint8, userId types.ID, appName string, dataType string, allowed bool, passwordSignature []byte) error
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.ContractList["Consents"] = (&Consents{}).new
	blockchain.RegisterSelector("0xbecae241", "consent(uint8,string,string,bool)")
	blockchain.RegisterSelector("0xf92519d8", "consentByController(uint8,bytes8,string,string,bool)")
	blockchain.RegisterSelector("0xedf2ef20", "modifyConsentByController(uint8,bytes8,string,string,bool,bytes)")
}

// bindConsents binds a generic wrapper to an already deployed contract.
func bindConsents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Consents *Consents) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewConsents(address, backend)
}

// IsAllowed is a free data retrieval call binding the contract method 0xa1d2bbf5.
//
// Solidity: function isAllowed(uint8 action, bytes8 userId, string appName, string dataType) constant returns(bool)
func (_Consents *ConsentsCaller) IsAllowed(opts *bind.CallOpts, action uint8, userId types.ID, appName string, dataType string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _Consents.contract.Call(opts, out, "isAllowed", action, userId, appName, dataType)
	return *ret0, err
}

// IsAllowed is a free data retrieval call binding the contract method 0xa1d2bbf5.
//
// Solidity: function isAllowed(uint8 action, bytes8 userId, string appName, string dataType) constant returns(bool)
func (_Consents *ConsentsSession) IsAllowed(action uint8, userId types.ID, appName string, dataType string) (bool, error) {
	return _Consents.Contract.IsAllowed(&_Consents.CallOpts, action, userId, appName, dataType)
}

// IsAllowed is a free data retrieval call binding the contract method 0xa1d2bbf5.
//
// Solidity: function isAllowed(uint8 action, bytes8 userId, string appName, string dataType) constant returns(bool)
func (_Consents *ConsentsCallerSession) IsAllowed(action uint8, userId types.ID, appName string, dataType string) (bool, error) {
	return _Consents.Contract.IsAllowed(&_Consents.CallOpts, action, userId, appName, dataType)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x118642e1.
//
// Solidity: function isAllowedAt(uint8 action, bytes8 userId, string appName, string dataType, uint256 blockNumber) constant returns(bool)
func (_Consents *ConsentsCaller) IsAllowedAt(opts *bind.CallOpts, action uint8, userId types.ID, appName string, dataType string, blockNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _Consents.contract.Call(opts, out, "isAllowedAt", action, userId, appName, dataType, blockNumber)
	return *ret0, err
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x118642e1.
//
// Solidity: function isAllowedAt(uint8 action, bytes8 userId, string appName, string dataType, uint256 blockNumber) constant returns(bool)
func (_Consents *ConsentsSession) IsAllowedAt(action uint8, userId types.ID, appName string, dataType string, blockNumber *big.Int) (bool, error) {
	return _Consents.Contract.IsAllowedAt(&_Consents.CallOpts, action, userId, appName, dataType, blockNumber)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x118642e1.
//
// Solidity: function isAllowedAt(uint8 action, bytes8 userId, string appName, string dataType, uint256 blockNumber) constant returns(bool)
func (_Consents *ConsentsCallerSession) IsAllowedAt(action uint8, userId types.ID, appName string, dataType string, blockNumber *big.Int) (bool, error) {
	return _Consents.Contract.IsAllowedAt(&_Consents.CallOpts, action, userId, appName, dataType, blockNumber)
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (_Consents *ConsentsTransactor) Consent(opts *bind.TransactOpts, action uint8, appName string, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "consent", action, appName, dataType, allowed)
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (_Consents *ConsentsSession) Consent(action uint8, appName string, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.Contract.Consent(&_Consents.TransactOpts, action, appName, dataType, allowed)
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (_Consents *ConsentsTransactorSession) Consent(action uint8, appName string, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.Contract.Consent(&_Consents.TransactOpts, action, appName, dataType, allowed)
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf92519d8.
//
// Solidity: function consentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed) returns()
func (_Consents *ConsentsTransactor) ConsentByController(opts *bind.TransactOpts, action uint8, userId types.ID, appName string, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "consentByController", action, userId, appName, dataType, allowed)
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf92519d8.
//
// Solidity: function consentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed) returns()
func (_Consents *ConsentsSession) ConsentByController(action uint8, userId types.ID, appName string, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ConsentByController(&_Consents.TransactOpts, action, userId, appName, dataType, allowed)
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf92519d8.
//
// Solidity: function consentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed) returns()
func (_Consents *ConsentsTransactorSession) ConsentByController(action uint8, userId types.ID, appName string, dataType string, allowed bool) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ConsentByController(&_Consents.TransactOpts, action, userId, appName, dataType, allowed)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xedf2ef20.
//
// Solidity: function modifyConsentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed, bytes passwordSignature) returns()
func (_Consents *ConsentsTransactor) ModifyConsentByController(opts *bind.TransactOpts, action uint8, userId types.ID, appName string, dataType string, allowed bool, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "modifyConsentByController", action, userId, appName, dataType, allowed, passwordSignature)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xedf2ef20.
//
// Solidity: function modifyConsentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed, bytes passwordSignature) returns()
func (_Consents *ConsentsSession) ModifyConsentByController(action uint8, userId types.ID, appName string, dataType string, allowed bool, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ModifyConsentByController(&_Consents.TransactOpts, action, userId, appName, dataType, allowed, passwordSignature)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xedf2ef20.
//
// Solidity: function modifyConsentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed, bytes passwordSignature) returns()
func (_Consents *ConsentsTransactorSession) ModifyConsentByController(action uint8, userId types.ID, appName string, dataType string, allowed bool, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Consents.Contract.ModifyConsentByController(&_Consents.TransactOpts, action, userId, appName, dataType, allowed, passwordSignature)
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
	App      [32]byte
	AppName  string
	DataType string
	Allowed  bool
	Raw      ethTypes.Log // Blockchain specific contextual infos
}

// FilterConsented is a free log retrieval operation binding the contract event 0xd0bd2a4b9fcbb6eee35bf0e8d542816e1d5244740220e033fff96b0abd805fac.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes32 indexed app, string appName, string dataType, bool allowed)
func (_Consents *ConsentsFilterer) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, app [][32]byte) (*ConsentsConsentedIterator, error) {

	var actionRule []interface{}
	for _, actionItem := range action {
		actionRule = append(actionRule, actionItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _Consents.contract.FilterLogs(opts, "Consented", actionRule, userIdRule, appRule)
	if err != nil {
		return nil, err
	}
	return &ConsentsConsentedIterator{contract: _Consents.contract, event: "Consented", logs: logs, sub: sub}, nil
}

// FilterConsented parses the event from given transaction receipt.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes32 indexed app, string appName, string dataType, bool allowed)
func (_Consents *ConsentsFilterer) ParseConsentedFromReceipt(receipt *ethTypes.Receipt) (*ConsentsConsented, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xd0bd2a4b9fcbb6eee35bf0e8d542816e1d5244740220e033fff96b0abd805fac") {
			event := new(ConsentsConsented)
			if err := _Consents.contract.UnpackLog(event, "Consented", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Consented event not found")
}

// WatchConsented is a free log subscription operation binding the contract event 0xd0bd2a4b9fcbb6eee35bf0e8d542816e1d5244740220e033fff96b0abd805fac.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes32 indexed app, string appName, string dataType, bool allowed)
func (_Consents *ConsentsFilterer) WatchConsented(opts *bind.WatchOpts, sink chan<- *ConsentsConsented, action []uint8, userId []types.ID, app [][32]byte) (event.Subscription, error) {

	var actionRule []interface{}
	for _, actionItem := range action {
		actionRule = append(actionRule, actionItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _Consents.contract.WatchLogs(opts, "Consented", actionRule, userIdRule, appRule)
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
