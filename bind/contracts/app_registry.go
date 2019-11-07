package contracts

import (
	"context"
	"errors"
	"math/big"
	"strings"

	ablbind "github.com/airbloc/airbloc-go/bind"
	types "github.com/airbloc/airbloc-go/bind/types"
	platform "github.com/klaytn/klaytn"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// AppRegistryABI is the input ABI used to generate the binding from.
const (
	AppRegistryAddress   = "0x719c7110898438a78d7f74Ce6Cc90DD69Fa320a0"
	AppRegistryTxHash    = "0x428fc7085318656f9f56bc2deadaf6b1065726e55cabfd9662c432dce3b092b3"
	AppRegistryCreatedAt = "0x0000000000000000000000000000000000000000000000000000000000823872"
	AppRegistryABI       = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"Registration\",\"signature\":\"0x0d8d636375a5c89a44d886dc1bd7257c82dbf9d475396c77cdbf443158ecf4e8\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"Unregistration\",\"signature\":\"0x03adf6d1cf18f8d8f64f7dbe8bde608e0d3fbca9079aa3cb3498715ef807bde9\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"AppOwnerTransferred\",\"signature\":\"0xbf3f214e451e16a835d0833b12209f2928a822c65ee68cce51eb31338747e3df\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"signature\":\"0xf2c298be\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"signature\":\"0x6598a1ae\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x693ec85e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x261a323e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xbde1eee7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAppOwner\",\"outputs\":[],\"payable\":false,\"signature\":\"0x1a9dff9f\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
)

// AppRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppRegistryCaller interface {
	Exists(
		ctx context.Context,
		appName string,
	) (
		bool,
		error,
	)
	Get(
		ctx context.Context,
		appName string,
	) (
		types.App,
		error,
	)
	IsOwner(
		ctx context.Context,
		appName string,
		owner common.Address,
	) (
		bool,
		error,
	)
}

type appRegistryCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (_AppRegistry *appRegistryCaller) Exists(ctx context.Context, appName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _AppRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "exists", appName)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns(types.App)
func (_AppRegistry *appRegistryCaller) Get(ctx context.Context, appName string) (types.App, error) {
	var (
		ret0 = new(types.App)
	)
	out := ret0

	err := _AppRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "get", appName)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (_AppRegistry *appRegistryCaller) IsOwner(ctx context.Context, appName string, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _AppRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "isOwner", appName, owner)
	return *ret0, err
}

// AppRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppRegistryTransactor interface {
	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
	) (*chainTypes.Receipt, error)
	TransferAppOwner(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
		newOwner common.Address,
	) (*chainTypes.Receipt, error)
	Unregister(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
	) (*chainTypes.Receipt, error)
}

type appRegistryTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
	backend  ablbind.ContractBackend
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (_AppRegistry *appRegistryTransactor) Register(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
) (*chainTypes.Receipt, error) {
	tx, err := _AppRegistry.contract.Transact(opts, "register", appName)
	if err != nil {
		return nil, err
	}
	return _AppRegistry.backend.WaitMined(ctx, tx)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (_AppRegistry *appRegistryTransactor) TransferAppOwner(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
	newOwner common.Address,
) (*chainTypes.Receipt, error) {
	tx, err := _AppRegistry.contract.Transact(opts, "transferAppOwner", appName, newOwner)
	if err != nil {
		return nil, err
	}
	return _AppRegistry.backend.WaitMined(ctx, tx)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (_AppRegistry *appRegistryTransactor) Unregister(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
) (*chainTypes.Receipt, error) {
	tx, err := _AppRegistry.contract.Transact(opts, "unregister", appName)
	if err != nil {
		return nil, err
	}
	return _AppRegistry.backend.WaitMined(ctx, tx)
}

type AppRegistryEvents interface {
	AppRegistryEventFilterer
	AppRegistryEventParser
	AppRegistryEventWatcher
}

// AppRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppRegistryEventFilterer interface {
	// Filterer
	FilterAppOwnerTransferred(
		opts *bind.FilterOpts,
		appAddr []common.Address, oldOwner []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRegistration(
		opts *bind.FilterOpts,
		appAddr []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterUnregistration(
		opts *bind.FilterOpts,
		appAddr []common.Address,
	) (ablbind.EventIterator, error)
}

type AppRegistryEventParser interface {
	// Parser
	ParseAppOwnerTransferred(log chainTypes.Log) (*AppRegistryAppOwnerTransferred, error)
	ParseAppOwnerTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryAppOwnerTransferred, error)

	// Parser
	ParseRegistration(log chainTypes.Log) (*AppRegistryRegistration, error)
	ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRegistration, error)

	// Parser
	ParseUnregistration(log chainTypes.Log) (*AppRegistryUnregistration, error)
	ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryUnregistration, error)
}

type AppRegistryEventWatcher interface {
	// Watcher
	WatchAppOwnerTransferred(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryAppOwnerTransferred,
		appAddr []common.Address, oldOwner []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchRegistration(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryRegistration,
		appAddr []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchUnregistration(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryUnregistration,
		appAddr []common.Address,
	) (event.Subscription, error)
}

type appRegistryEvents struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryAppOwnerTransferredIterator is returned from FilterAppOwnerTransferred and is used to iterate over the raw logs and unpacked data for AppOwnerTransferred events raised by the AppRegistry contract.
type AppRegistryAppOwnerTransferredIterator struct {
	Evt *AppRegistryAppOwnerTransferred // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

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
			it.Evt = new(AppRegistryAppOwnerTransferred)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(AppRegistryAppOwnerTransferred)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryAppOwnerTransferredIterator) Event() interface{} {
	return it.Evt
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
func (_AppRegistry *appRegistryEvents) FilterAppOwnerTransferred(opts *bind.FilterOpts, appAddr []common.Address, oldOwner []common.Address) (ablbind.EventIterator, error) {

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
func (_AppRegistry *appRegistryEvents) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AppRegistryAppOwnerTransferred, appAddr []common.Address, oldOwner []common.Address) (event.Subscription, error) {

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
				evt := new(AppRegistryAppOwnerTransferred)
				if err := _AppRegistry.contract.UnpackLog(evt, "AppOwnerTransferred", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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
func (_AppRegistry *appRegistryEvents) ParseAppOwnerTransferred(log chainTypes.Log) (*AppRegistryAppOwnerTransferred, error) {
	evt := new(AppRegistryAppOwnerTransferred)
	if err := _AppRegistry.contract.UnpackLog(evt, "AppOwnerTransferred", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseAppOwnerTransferredFromReceipt parses the event from given transaction receipt.
//
// Solidity: event AppOwnerTransferred(address indexed appAddr, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *appRegistryEvents) ParseAppOwnerTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryAppOwnerTransferred, error) {
	var evts []*AppRegistryAppOwnerTransferred
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xbf3f214e451e16a835d0833b12209f2928a822c65ee68cce51eb31338747e3df") {
			evt, err := _AppRegistry.ParseAppOwnerTransferred(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("AppOwnerTransferred event not found")
	}
	return evts, nil
}

// AppRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the AppRegistry contract.
type AppRegistryRegistrationIterator struct {
	Evt *AppRegistryRegistration // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

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
			it.Evt = new(AppRegistryRegistration)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(AppRegistryRegistration)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryRegistrationIterator) Event() interface{} {
	return it.Evt
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
func (_AppRegistry *appRegistryEvents) FilterRegistration(opts *bind.FilterOpts, appAddr []common.Address) (ablbind.EventIterator, error) {

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
func (_AppRegistry *appRegistryEvents) WatchRegistration(opts *bind.WatchOpts, sink chan<- *AppRegistryRegistration, appAddr []common.Address) (event.Subscription, error) {

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
				evt := new(AppRegistryRegistration)
				if err := _AppRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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
func (_AppRegistry *appRegistryEvents) ParseRegistration(log chainTypes.Log) (*AppRegistryRegistration, error) {
	evt := new(AppRegistryRegistration)
	if err := _AppRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRegistrationFromReceipt parses the event from given transaction receipt.
//
// Solidity: event Registration(address indexed appAddr, string appName)
func (_AppRegistry *appRegistryEvents) ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRegistration, error) {
	var evts []*AppRegistryRegistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x0d8d636375a5c89a44d886dc1bd7257c82dbf9d475396c77cdbf443158ecf4e8") {
			evt, err := _AppRegistry.ParseRegistration(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("Registration event not found")
	}
	return evts, nil
}

// AppRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the AppRegistry contract.
type AppRegistryUnregistrationIterator struct {
	Evt *AppRegistryUnregistration // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

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
			it.Evt = new(AppRegistryUnregistration)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(AppRegistryUnregistration)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryUnregistrationIterator) Event() interface{} {
	return it.Evt
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
func (_AppRegistry *appRegistryEvents) FilterUnregistration(opts *bind.FilterOpts, appAddr []common.Address) (ablbind.EventIterator, error) {

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
func (_AppRegistry *appRegistryEvents) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *AppRegistryUnregistration, appAddr []common.Address) (event.Subscription, error) {

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
				evt := new(AppRegistryUnregistration)
				if err := _AppRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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
func (_AppRegistry *appRegistryEvents) ParseUnregistration(log chainTypes.Log) (*AppRegistryUnregistration, error) {
	evt := new(AppRegistryUnregistration)
	if err := _AppRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseUnregistrationFromReceipt parses the event from given transaction receipt.
//
// Solidity: event Unregistration(address indexed appAddr, string appName)
func (_AppRegistry *appRegistryEvents) ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryUnregistration, error) {
	var evts []*AppRegistryUnregistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x03adf6d1cf18f8d8f64f7dbe8bde608e0d3fbca9079aa3cb3498715ef807bde9") {
			evt, err := _AppRegistry.ParseUnregistration(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("Unregistration event not found")
	}
	return evts, nil
}

// Manager is contract wrapper struct
type AppRegistryContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	AppRegistryCaller
	AppRegistryTransactor
	AppRegistryEvents
}

func NewAppRegistryContract(backend ablbind.ContractBackend) (*AppRegistryContract, error) {
	deployment, exist := backend.Deployment("AppRegistry")
	if !exist {
		evmABI, err := abi.JSON(strings.NewReader(AppRegistryABI))
		if err != nil {
			return nil, err
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(AppRegistryAddress),
			common.HexToHash(AppRegistryTxHash),
			new(big.Int).SetBytes(common.HexToHash(AppRegistryCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, "AppRegistry", backend)

	contract := &AppRegistryContract{
		Deployment: deployment,
		client:     backend,

		AppRegistryCaller:     &appRegistryCaller{base},
		AppRegistryTransactor: &appRegistryTransactor{base, backend},
		AppRegistryEvents:     &appRegistryEvents{base},
	}

	return contract, nil
}
