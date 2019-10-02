// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"

	blockchain "github.com/airbloc/airbloc-go/shared/blockchain"
	types "github.com/airbloc/airbloc-go/shared/types"
	platform "github.com/klaytn/klaytn"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// ConsentsABI is the input ABI used to generate the binding from.
const ConsentsABI = "[{\"inputs\":[{\"name\":\"accountReg\",\"type\":\"address\"},{\"name\":\"appReg\",\"type\":\"address\"},{\"name\":\"controllerReg\",\"type\":\"address\"},{\"name\":\"dataTypeReg\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"dataType\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"Consented\",\"signature\":\"0x8599a1c756b9cd519b80b172f29a03b19082bf7df728da8456cbcab9eeaba8e3\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"components\":[{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"consentData\",\"type\":\"tuple\"}],\"name\":\"consent\",\"outputs\":[],\"payable\":false,\"signature\":\"0xcd4dc804\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"components\":[{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"consentData\",\"type\":\"tuple[]\"}],\"name\":\"consentMany\",\"outputs\":[],\"payable\":false,\"signature\":\"0xdd43ad05\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"components\":[{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"consentData\",\"type\":\"tuple\"}],\"name\":\"consentByController\",\"outputs\":[],\"payable\":false,\"signature\":\"0xf573f89a\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"components\":[{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"consentData\",\"type\":\"tuple[]\"}],\"name\":\"consentManyByController\",\"outputs\":[],\"payable\":false,\"signature\":\"0xae6d5034\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"components\":[{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"consentData\",\"type\":\"tuple\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"modifyConsentByController\",\"outputs\":[],\"payable\":false,\"signature\":\"0x0bfec389\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"components\":[{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"consentData\",\"type\":\"tuple[]\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"modifyConsentManyByController\",\"outputs\":[],\"payable\":false,\"signature\":\"0xe031b1cf\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x50615985\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"isAllowedAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x7cdda67c\",\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Consents is an auto generated Go binding around an Ethereum contract.
type Consents struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int

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

// ConsentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsentsCaller struct {
	contract *blockchain.BoundContract // Generic contract wrapper for the low level calls
}

// ConsentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsentsTransactor struct {
	contract *blockchain.BoundContract // Generic contract wrapper for the low level calls
}

// ConsentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsentsFilterer struct {
	contract *blockchain.BoundContract // Generic contract wrapper for the low level calls
}

// ConsentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsentsSession struct {
	Contract     *Consents               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts blockchain.TransactOpts // Transaction auth options to use throughout this session
}

// ConsentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsentsCallerSession struct {
	Contract *ConsentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ConsentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsentsTransactorSession struct {
	Contract     *ConsentsTransactor     // Generic contract transactor binding to set the session for
	TransactOpts blockchain.TransactOpts // Transaction auth options to use throughout this session
}

// ConsentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsentsRaw struct {
	Contract *Consents // Generic contract binding to access the raw methods on
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
func (_Consents *ConsentsRaw) Transfer(opts *blockchain.TransactOpts) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ConsentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consents *ConsentsRaw) Transact(opts *blockchain.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ConsentsTransactor.contract.Transact(opts, method, params...)
}

// ConsentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsentsCallerRaw struct {
	Contract *ConsentsCaller // Generic read-only contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consents *ConsentsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Consents.Contract.contract.Call(opts, result, method, params...)
}

// ConsentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsentsTransactorRaw struct {
	Contract *ConsentsTransactor // Generic write-only contract binding to access the raw methods on
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consents *ConsentsTransactorRaw) Transfer(opts *blockchain.TransactOpts) (*chainTypes.Transaction, error) {
	return _Consents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consents *ConsentsTransactorRaw) Transact(opts *blockchain.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _Consents.Contract.contract.Transact(opts, method, params...)
}

// IsAllowed is a free data retrieval call binding the contract method 0x50615985.
//
// Solidity: function isAllowed(bytes8 userId, string appName, uint8 action, string dataType) constant returns(bool)
func (_Consents *ConsentsCaller) IsAllowed(opts *bind.CallOpts, userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
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
func (_Consents *ConsentsCaller) IsAllowedAt(opts *bind.CallOpts, userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
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

// Consent is a paid mutator transaction binding the contract method 0xcd4dc804.
//
// Solidity: function consent(string appName, types.ConsentData consentData) returns()
func (_Consents *ConsentsTransactor) Consent(opts *blockchain.TransactOpts, appName string, consentData types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "consent", appName, consentData)
}

// Consent is a paid mutator transaction binding the contract method 0xcd4dc804.
//
// Solidity: function consent(string appName, types.ConsentData consentData) returns()
func (_Consents *ConsentsSession) Consent(appName string, consentData types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.Contract.Consent(&_Consents.TransactOpts, appName, consentData)
}

// Consent is a paid mutator transaction binding the contract method 0xcd4dc804.
//
// Solidity: function consent(string appName, types.ConsentData consentData) returns()
func (_Consents *ConsentsTransactorSession) Consent(appName string, consentData types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.Contract.Consent(&_Consents.TransactOpts, appName, consentData)
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf573f89a.
//
// Solidity: function consentByController(bytes8 userId, string appName, types.ConsentData consentData) returns()
func (_Consents *ConsentsTransactor) ConsentByController(opts *blockchain.TransactOpts, userId types.ID, appName string, consentData types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "consentByController", userId, appName, consentData)
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf573f89a.
//
// Solidity: function consentByController(bytes8 userId, string appName, types.ConsentData consentData) returns()
func (_Consents *ConsentsSession) ConsentByController(userId types.ID, appName string, consentData types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ConsentByController(&_Consents.TransactOpts, userId, appName, consentData)
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf573f89a.
//
// Solidity: function consentByController(bytes8 userId, string appName, types.ConsentData consentData) returns()
func (_Consents *ConsentsTransactorSession) ConsentByController(userId types.ID, appName string, consentData types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ConsentByController(&_Consents.TransactOpts, userId, appName, consentData)
}

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, []types.ConsentData consentData) returns()
func (_Consents *ConsentsTransactor) ConsentMany(opts *blockchain.TransactOpts, appName string, consentData []types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "consentMany", appName, consentData)
}

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, []types.ConsentData consentData) returns()
func (_Consents *ConsentsSession) ConsentMany(appName string, consentData []types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ConsentMany(&_Consents.TransactOpts, appName, consentData)
}

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, []types.ConsentData consentData) returns()
func (_Consents *ConsentsTransactorSession) ConsentMany(appName string, consentData []types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ConsentMany(&_Consents.TransactOpts, appName, consentData)
}

// ConsentManyByController is a paid mutator transaction binding the contract method 0xae6d5034.
//
// Solidity: function consentManyByController(bytes8 userId, string appName, []types.ConsentData consentData) returns()
func (_Consents *ConsentsTransactor) ConsentManyByController(opts *blockchain.TransactOpts, userId types.ID, appName string, consentData []types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "consentManyByController", userId, appName, consentData)
}

// ConsentManyByController is a paid mutator transaction binding the contract method 0xae6d5034.
//
// Solidity: function consentManyByController(bytes8 userId, string appName, []types.ConsentData consentData) returns()
func (_Consents *ConsentsSession) ConsentManyByController(userId types.ID, appName string, consentData []types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ConsentManyByController(&_Consents.TransactOpts, userId, appName, consentData)
}

// ConsentManyByController is a paid mutator transaction binding the contract method 0xae6d5034.
//
// Solidity: function consentManyByController(bytes8 userId, string appName, []types.ConsentData consentData) returns()
func (_Consents *ConsentsTransactorSession) ConsentManyByController(userId types.ID, appName string, consentData []types.ConsentData) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ConsentManyByController(&_Consents.TransactOpts, userId, appName, consentData)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0x0bfec389.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, types.ConsentData consentData, bytes passwordSignature) returns()
func (_Consents *ConsentsTransactor) ModifyConsentByController(opts *blockchain.TransactOpts, userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) (*chainTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "modifyConsentByController", userId, appName, consentData, passwordSignature)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0x0bfec389.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, types.ConsentData consentData, bytes passwordSignature) returns()
func (_Consents *ConsentsSession) ModifyConsentByController(userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ModifyConsentByController(&_Consents.TransactOpts, userId, appName, consentData, passwordSignature)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0x0bfec389.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, types.ConsentData consentData, bytes passwordSignature) returns()
func (_Consents *ConsentsTransactorSession) ModifyConsentByController(userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ModifyConsentByController(&_Consents.TransactOpts, userId, appName, consentData, passwordSignature)
}

// ModifyConsentManyByController is a paid mutator transaction binding the contract method 0xe031b1cf.
//
// Solidity: function modifyConsentManyByController(bytes8 userId, string appName, []types.ConsentData consentData, bytes passwordSignature) returns()
func (_Consents *ConsentsTransactor) ModifyConsentManyByController(opts *blockchain.TransactOpts, userId types.ID, appName string, consentData []types.ConsentData, passwordSignature []byte) (*chainTypes.Transaction, error) {
	return _Consents.contract.Transact(opts, "modifyConsentManyByController", userId, appName, consentData, passwordSignature)
}

// ModifyConsentManyByController is a paid mutator transaction binding the contract method 0xe031b1cf.
//
// Solidity: function modifyConsentManyByController(bytes8 userId, string appName, []types.ConsentData consentData, bytes passwordSignature) returns()
func (_Consents *ConsentsSession) ModifyConsentManyByController(userId types.ID, appName string, consentData []types.ConsentData, passwordSignature []byte) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ModifyConsentManyByController(&_Consents.TransactOpts, userId, appName, consentData, passwordSignature)
}

// ModifyConsentManyByController is a paid mutator transaction binding the contract method 0xe031b1cf.
//
// Solidity: function modifyConsentManyByController(bytes8 userId, string appName, []types.ConsentData consentData, bytes passwordSignature) returns()
func (_Consents *ConsentsTransactorSession) ModifyConsentManyByController(userId types.ID, appName string, consentData []types.ConsentData, passwordSignature []byte) (*chainTypes.Transaction, error) {
	return _Consents.Contract.ModifyConsentManyByController(&_Consents.TransactOpts, userId, appName, consentData, passwordSignature)
}

// ConsentsConsentedIterator is returned from FilterConsented and is used to iterate over the raw logs and unpacked data for Consented events raised by the Consents contract.
type ConsentsConsentedIterator struct {
	Event *ConsentsConsented // Event containing the contract specifics and raw log

	contract *blockchain.BoundContract // Generic contract to use for unpacking event data
	event    string                    // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
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
	Raw      chainTypes.Log // Blockchain specific contextual infos
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

// ParseConsented is a log parse operation binding the contract event 0x8599a1c756b9cd519b80b172f29a03b19082bf7df728da8456cbcab9eeaba8e3.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, address indexed appAddr, string appName, string dataType, bool allowed)
func (_Consents *ConsentsFilterer) ParseConsented(log chainTypes.Log) (*ConsentsConsented, error) {
	event := new(ConsentsConsented)
	if err := _Consents.contract.UnpackLog(event, "Consented", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterConsented parses the event from given transaction receipt.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, address indexed appAddr, string appName, string dataType, bool allowed)
func (_Consents *ConsentsFilterer) ParseConsentedFromReceipt(receipt *chainTypes.Receipt) ([]*ConsentsConsented, error) {
	var events []*ConsentsConsented
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8599a1c756b9cd519b80b172f29a03b19082bf7df728da8456cbcab9eeaba8e3") {
			event, err := _Consents.ParseConsented(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("Consented event not found")
	}
	return events, nil
}
