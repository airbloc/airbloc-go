// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"

	types "github.com/airbloc/airbloc-go/shared/types"
	platform "github.com/klaytn/klaytn"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// AppRegistryABI is the input ABI used to generate the binding from.
const AppRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"Registration\",\"signature\":\"0x0d8d636375a5c89a44d886dc1bd7257c82dbf9d475396c77cdbf443158ecf4e8\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"Unregistration\",\"signature\":\"0x03adf6d1cf18f8d8f64f7dbe8bde608e0d3fbca9079aa3cb3498715ef807bde9\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"AppOwnerTransferred\",\"signature\":\"0xbf3f214e451e16a835d0833b12209f2928a822c65ee68cce51eb31338747e3df\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"signature\":\"0xf2c298be\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"signature\":\"0x6598a1ae\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x693ec85e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x261a323e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xbde1eee7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAppOwner\",\"outputs\":[],\"payable\":false,\"signature\":\"0x1a9dff9f\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AppRegistry is an auto generated Go binding around an Ethereum contract.
type AppRegistry struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int

	AppRegistryCaller     // Read-only binding to the contract
	AppRegistryTransactor // Write-only binding to the contract
	AppRegistryFilterer   // Log filterer for contract events
}

// Address is getter method of AppRegistry.address
func (_AppRegistry *AppRegistry) Address() common.Address {
	return _AppRegistry.address
}

// TxHash is getter method of AppRegistry.txHash
func (_AppRegistry *AppRegistry) TxHash() common.Hash {
	return _AppRegistry.txHash
}

// CreatedAt is getter method of AppRegistry.createdAt
func (_AppRegistry *AppRegistry) CreatedAt() *big.Int {
	return _AppRegistry.createdAt
}

// AppRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppRegistrySession struct {
	Contract     *AppRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppRegistryCallerSession struct {
	Contract *AppRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AppRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppRegistryTransactorSession struct {
	Contract     *AppRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AppRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRegistryRaw struct {
	Contract *AppRegistry // Generic contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppRegistry *AppRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AppRegistry.Contract.AppRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppRegistry *AppRegistryRaw) Transfer(opts *bind.TransactOpts) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transact(opts, method, params...)
}

// AppRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppRegistryCallerRaw struct {
	Contract *AppRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppRegistry *AppRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AppRegistry.Contract.contract.Call(opts, result, method, params...)
}

// AppRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppRegistryTransactorRaw struct {
	Contract *AppRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppRegistry *AppRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (_AppRegistry *AppRegistryCaller) Exists(opts *bind.CallOpts, appName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "exists", appName)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (_AppRegistry *AppRegistrySession) Exists(appName string) (bool, error) {
	return _AppRegistry.Contract.Exists(&_AppRegistry.CallOpts, appName)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (_AppRegistry *AppRegistryCallerSession) Exists(appName string) (bool, error) {
	return _AppRegistry.Contract.Exists(&_AppRegistry.CallOpts, appName)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns(types.App)
func (_AppRegistry *AppRegistryCaller) Get(opts *bind.CallOpts, appName string) (types.App, error) {
	var (
		ret0 = new(types.App)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "get", appName)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns(types.App)
func (_AppRegistry *AppRegistrySession) Get(appName string) (types.App, error) {
	return _AppRegistry.Contract.Get(&_AppRegistry.CallOpts, appName)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns(types.App)
func (_AppRegistry *AppRegistryCallerSession) Get(appName string) (types.App, error) {
	return _AppRegistry.Contract.Get(&_AppRegistry.CallOpts, appName)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (_AppRegistry *AppRegistryCaller) IsOwner(opts *bind.CallOpts, appName string, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "isOwner", appName, owner)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (_AppRegistry *AppRegistrySession) IsOwner(appName string, owner common.Address) (bool, error) {
	return _AppRegistry.Contract.IsOwner(&_AppRegistry.CallOpts, appName, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (_AppRegistry *AppRegistryCallerSession) IsOwner(appName string, owner common.Address) (bool, error) {
	return _AppRegistry.Contract.IsOwner(&_AppRegistry.CallOpts, appName, owner)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (_AppRegistry *AppRegistryTransactor) Register(opts *bind.TransactOpts, appName string) (*chainTypes.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "register", appName)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (_AppRegistry *AppRegistrySession) Register(appName string) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.Register(&_AppRegistry.TransactOpts, appName)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (_AppRegistry *AppRegistryTransactorSession) Register(appName string) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.Register(&_AppRegistry.TransactOpts, appName)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (_AppRegistry *AppRegistryTransactor) TransferAppOwner(opts *bind.TransactOpts, appName string, newOwner common.Address) (*chainTypes.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "transferAppOwner", appName, newOwner)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (_AppRegistry *AppRegistrySession) TransferAppOwner(appName string, newOwner common.Address) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.TransferAppOwner(&_AppRegistry.TransactOpts, appName, newOwner)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (_AppRegistry *AppRegistryTransactorSession) TransferAppOwner(appName string, newOwner common.Address) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.TransferAppOwner(&_AppRegistry.TransactOpts, appName, newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (_AppRegistry *AppRegistryTransactor) Unregister(opts *bind.TransactOpts, appName string) (*chainTypes.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "unregister", appName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (_AppRegistry *AppRegistrySession) Unregister(appName string) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.Unregister(&_AppRegistry.TransactOpts, appName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (_AppRegistry *AppRegistryTransactorSession) Unregister(appName string) (*chainTypes.Transaction, error) {
	return _AppRegistry.Contract.Unregister(&_AppRegistry.TransactOpts, appName)
}

// AppRegistryAppOwnerTransferredIterator is returned from FilterAppOwnerTransferred and is used to iterate over the raw logs and unpacked data for AppOwnerTransferred events raised by the AppRegistry contract.
type AppRegistryAppOwnerTransferredIterator struct {
	Event *AppRegistryAppOwnerTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppRegistryAppOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryAppOwnerTransferred)
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
		it.Event = new(AppRegistryAppOwnerTransferred)
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
func (it *AppRegistryAppOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryAppOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryAppOwnerTransferred represents a AppOwnerTransferred event raised by the AppRegistry contract.
type AppRegistryAppOwnerTransferred struct {
	AppAddr  common.Address
	AppName  string
	OldOwner common.Address
	NewOwner common.Address
	Raw      chainTypes.Log // Blockchain specific contextual infos
}

// FilterAppOwnerTransferred is a free log retrieval operation binding the contract event 0xbf3f214e451e16a835d0833b12209f2928a822c65ee68cce51eb31338747e3df.
//
// Solidity: event AppOwnerTransferred(address indexed appAddr, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *AppRegistryFilterer) FilterAppOwnerTransferred(opts *bind.FilterOpts, appAddr []common.Address, oldOwner []common.Address) (*AppRegistryAppOwnerTransferredIterator, error) {

	var appAddrRule []interface{}
	for _, appAddrItem := range appAddr {
		appAddrRule = append(appAddrRule, appAddrItem)
	}

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "AppOwnerTransferred", appAddrRule, oldOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryAppOwnerTransferredIterator{contract: _AppRegistry.contract, event: "AppOwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchAppOwnerTransferred is a free log subscription operation binding the contract event 0xbf3f214e451e16a835d0833b12209f2928a822c65ee68cce51eb31338747e3df.
//
// Solidity: event AppOwnerTransferred(address indexed appAddr, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *AppRegistryFilterer) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AppRegistryAppOwnerTransferred, appAddr []common.Address, oldOwner []common.Address) (event.Subscription, error) {

	var appAddrRule []interface{}
	for _, appAddrItem := range appAddr {
		appAddrRule = append(appAddrRule, appAddrItem)
	}

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "AppOwnerTransferred", appAddrRule, oldOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryAppOwnerTransferred)
				if err := _AppRegistry.contract.UnpackLog(event, "AppOwnerTransferred", log); err != nil {
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

// ParseAppOwnerTransferred is a log parse operation binding the contract event 0xbf3f214e451e16a835d0833b12209f2928a822c65ee68cce51eb31338747e3df.
//
// Solidity: event AppOwnerTransferred(address indexed appAddr, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *AppRegistryFilterer) ParseAppOwnerTransferred(log chainTypes.Log) (*AppRegistryAppOwnerTransferred, error) {
	event := new(AppRegistryAppOwnerTransferred)
	if err := _AppRegistry.contract.UnpackLog(event, "AppOwnerTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterAppOwnerTransferred parses the event from given transaction receipt.
//
// Solidity: event AppOwnerTransferred(address indexed appAddr, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *AppRegistryFilterer) ParseAppOwnerTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryAppOwnerTransferred, error) {
	var events []*AppRegistryAppOwnerTransferred
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xbf3f214e451e16a835d0833b12209f2928a822c65ee68cce51eb31338747e3df") {
			event, err := _AppRegistry.ParseAppOwnerTransferred(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("AppOwnerTransferred event not found")
	}
	return events, nil
}

// AppRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the AppRegistry contract.
type AppRegistryRegistrationIterator struct {
	Event *AppRegistryRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryRegistration)
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
		it.Event = new(AppRegistryRegistration)
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
func (it *AppRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRegistration represents a Registration event raised by the AppRegistry contract.
type AppRegistryRegistration struct {
	AppAddr common.Address
	AppName string
	Raw     chainTypes.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0x0d8d636375a5c89a44d886dc1bd7257c82dbf9d475396c77cdbf443158ecf4e8.
//
// Solidity: event Registration(address indexed appAddr, string appName)
func (_AppRegistry *AppRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, appAddr []common.Address) (*AppRegistryRegistrationIterator, error) {

	var appAddrRule []interface{}
	for _, appAddrItem := range appAddr {
		appAddrRule = append(appAddrRule, appAddrItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Registration", appAddrRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRegistrationIterator{contract: _AppRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// WatchRegistration is a free log subscription operation binding the contract event 0x0d8d636375a5c89a44d886dc1bd7257c82dbf9d475396c77cdbf443158ecf4e8.
//
// Solidity: event Registration(address indexed appAddr, string appName)
func (_AppRegistry *AppRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *AppRegistryRegistration, appAddr []common.Address) (event.Subscription, error) {

	var appAddrRule []interface{}
	for _, appAddrItem := range appAddr {
		appAddrRule = append(appAddrRule, appAddrItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Registration", appAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryRegistration)
				if err := _AppRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
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

// ParseRegistration is a log parse operation binding the contract event 0x0d8d636375a5c89a44d886dc1bd7257c82dbf9d475396c77cdbf443158ecf4e8.
//
// Solidity: event Registration(address indexed appAddr, string appName)
func (_AppRegistry *AppRegistryFilterer) ParseRegistration(log chainTypes.Log) (*AppRegistryRegistration, error) {
	event := new(AppRegistryRegistration)
	if err := _AppRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(address indexed appAddr, string appName)
func (_AppRegistry *AppRegistryFilterer) ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRegistration, error) {
	var events []*AppRegistryRegistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x0d8d636375a5c89a44d886dc1bd7257c82dbf9d475396c77cdbf443158ecf4e8") {
			event, err := _AppRegistry.ParseRegistration(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("Registration event not found")
	}
	return events, nil
}

// AppRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the AppRegistry contract.
type AppRegistryUnregistrationIterator struct {
	Event *AppRegistryUnregistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryUnregistration)
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
		it.Event = new(AppRegistryUnregistration)
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
func (it *AppRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryUnregistration represents a Unregistration event raised by the AppRegistry contract.
type AppRegistryUnregistration struct {
	AppAddr common.Address
	AppName string
	Raw     chainTypes.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x03adf6d1cf18f8d8f64f7dbe8bde608e0d3fbca9079aa3cb3498715ef807bde9.
//
// Solidity: event Unregistration(address indexed appAddr, string appName)
func (_AppRegistry *AppRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts, appAddr []common.Address) (*AppRegistryUnregistrationIterator, error) {

	var appAddrRule []interface{}
	for _, appAddrItem := range appAddr {
		appAddrRule = append(appAddrRule, appAddrItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Unregistration", appAddrRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryUnregistrationIterator{contract: _AppRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x03adf6d1cf18f8d8f64f7dbe8bde608e0d3fbca9079aa3cb3498715ef807bde9.
//
// Solidity: event Unregistration(address indexed appAddr, string appName)
func (_AppRegistry *AppRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *AppRegistryUnregistration, appAddr []common.Address) (event.Subscription, error) {

	var appAddrRule []interface{}
	for _, appAddrItem := range appAddr {
		appAddrRule = append(appAddrRule, appAddrItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Unregistration", appAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryUnregistration)
				if err := _AppRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
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

// ParseUnregistration is a log parse operation binding the contract event 0x03adf6d1cf18f8d8f64f7dbe8bde608e0d3fbca9079aa3cb3498715ef807bde9.
//
// Solidity: event Unregistration(address indexed appAddr, string appName)
func (_AppRegistry *AppRegistryFilterer) ParseUnregistration(log chainTypes.Log) (*AppRegistryUnregistration, error) {
	event := new(AppRegistryUnregistration)
	if err := _AppRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(address indexed appAddr, string appName)
func (_AppRegistry *AppRegistryFilterer) ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryUnregistration, error) {
	var events []*AppRegistryUnregistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x03adf6d1cf18f8d8f64f7dbe8bde608e0d3fbca9079aa3cb3498715ef807bde9") {
			event, err := _AppRegistry.ParseUnregistration(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("Unregistration event not found")
	}
	return events, nil
}
