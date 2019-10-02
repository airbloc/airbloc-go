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

// ControllerRegistryABI is the input ABI used to generate the binding from.
const ControllerRegistryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"signature\":\"0x715018a6\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"0x8da5cb5b\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x8f32d59b\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"signature\":\"0xf2fde38b\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"Registration\",\"signature\":\"0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"Unregistration\",\"signature\":\"0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"signature\":\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"controllerAddr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"signature\":\"0x4420e486\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"usersCount\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0xc2bc2efc\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xf6a3d24e\",\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ControllerRegistry is an auto generated Go binding around an Ethereum contract.
type ControllerRegistry struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int

	ControllerRegistryCaller     // Read-only binding to the contract
	ControllerRegistryTransactor // Write-only binding to the contract
	ControllerRegistryFilterer   // Log filterer for contract events
}

// Address is getter method of ControllerRegistry.address
func (_ControllerRegistry *ControllerRegistry) Address() common.Address {
	return _ControllerRegistry.address
}

// TxHash is getter method of ControllerRegistry.txHash
func (_ControllerRegistry *ControllerRegistry) TxHash() common.Hash {
	return _ControllerRegistry.txHash
}

// CreatedAt is getter method of ControllerRegistry.createdAt
func (_ControllerRegistry *ControllerRegistry) CreatedAt() *big.Int {
	return _ControllerRegistry.createdAt
}

// ControllerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerRegistryCaller struct {
	contract *blockchain.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerRegistryTransactor struct {
	contract *blockchain.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerRegistryFilterer struct {
	contract *blockchain.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerRegistrySession struct {
	Contract     *ControllerRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts blockchain.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerRegistryCallerSession struct {
	Contract *ControllerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ControllerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerRegistryTransactorSession struct {
	Contract     *ControllerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts blockchain.TransactOpts       // Transaction auth options to use throughout this session
}

// ControllerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRegistryRaw struct {
	Contract *ControllerRegistry // Generic contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ControllerRegistry *ControllerRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ControllerRegistry.Contract.ControllerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ControllerRegistry *ControllerRegistryRaw) Transfer(opts *blockchain.TransactOpts) (*chainTypes.Transaction, error) {
	return _ControllerRegistry.Contract.ControllerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ControllerRegistry *ControllerRegistryRaw) Transact(opts *blockchain.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _ControllerRegistry.Contract.ControllerRegistryTransactor.contract.Transact(opts, method, params...)
}

// ControllerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerRegistryCallerRaw struct {
	Contract *ControllerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ControllerRegistry *ControllerRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ControllerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// ControllerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerRegistryTransactorRaw struct {
	Contract *ControllerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ControllerRegistry *ControllerRegistryTransactorRaw) Transfer(opts *blockchain.TransactOpts) (*chainTypes.Transaction, error) {
	return _ControllerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ControllerRegistry *ControllerRegistryTransactorRaw) Transact(opts *blockchain.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _ControllerRegistry.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (_ControllerRegistry *ControllerRegistryCaller) Exists(opts *bind.CallOpts, controller common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ControllerRegistry.contract.Call(opts, out, "exists", controller)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (_ControllerRegistry *ControllerRegistrySession) Exists(controller common.Address) (bool, error) {
	return _ControllerRegistry.Contract.Exists(&_ControllerRegistry.CallOpts, controller)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (_ControllerRegistry *ControllerRegistryCallerSession) Exists(controller common.Address) (bool, error) {
	return _ControllerRegistry.Contract.Exists(&_ControllerRegistry.CallOpts, controller)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns(types.DataController)
func (_ControllerRegistry *ControllerRegistryCaller) Get(opts *bind.CallOpts, controller common.Address) (types.DataController, error) {
	var (
		ret0 = new(types.DataController)
	)
	out := ret0
	err := _ControllerRegistry.contract.Call(opts, out, "get", controller)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns(types.DataController)
func (_ControllerRegistry *ControllerRegistrySession) Get(controller common.Address) (types.DataController, error) {
	return _ControllerRegistry.Contract.Get(&_ControllerRegistry.CallOpts, controller)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns(types.DataController)
func (_ControllerRegistry *ControllerRegistryCallerSession) Get(controller common.Address) (types.DataController, error) {
	return _ControllerRegistry.Contract.Get(&_ControllerRegistry.CallOpts, controller)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (_ControllerRegistry *ControllerRegistryTransactor) Register(opts *blockchain.TransactOpts, controllerAddr common.Address) (*chainTypes.Transaction, error) {
	return _ControllerRegistry.contract.Transact(opts, "register", controllerAddr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (_ControllerRegistry *ControllerRegistrySession) Register(controllerAddr common.Address) (*chainTypes.Transaction, error) {
	return _ControllerRegistry.Contract.Register(&_ControllerRegistry.TransactOpts, controllerAddr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (_ControllerRegistry *ControllerRegistryTransactorSession) Register(controllerAddr common.Address) (*chainTypes.Transaction, error) {
	return _ControllerRegistry.Contract.Register(&_ControllerRegistry.TransactOpts, controllerAddr)
}

// ControllerRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ControllerRegistry contract.
type ControllerRegistryOwnershipTransferredIterator struct {
	Event *ControllerRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ControllerRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRegistryOwnershipTransferred)
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
		it.Event = new(ControllerRegistryOwnershipTransferred)
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
func (it *ControllerRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ControllerRegistry contract.
type ControllerRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           chainTypes.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryOwnershipTransferredIterator{contract: _ControllerRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRegistryOwnershipTransferred)
				if err := _ControllerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseOwnershipTransferred(log chainTypes.Log) (*ControllerRegistryOwnershipTransferred, error) {
	event := new(ControllerRegistryOwnershipTransferred)
	if err := _ControllerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterOwnershipTransferred parses the event from given transaction receipt.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseOwnershipTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryOwnershipTransferred, error) {
	var events []*ControllerRegistryOwnershipTransferred
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0") {
			event, err := _ControllerRegistry.ParseOwnershipTransferred(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("OwnershipTransferred event not found")
	}
	return events, nil
}

// ControllerRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the ControllerRegistry contract.
type ControllerRegistryRegistrationIterator struct {
	Event *ControllerRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *ControllerRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRegistryRegistration)
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
		it.Event = new(ControllerRegistryRegistration)
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
func (it *ControllerRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRegistryRegistration represents a Registration event raised by the ControllerRegistry contract.
type ControllerRegistryRegistration struct {
	Controller common.Address
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryRegistrationIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.FilterLogs(opts, "Registration", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryRegistrationIterator{contract: _ControllerRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// WatchRegistration is a free log subscription operation binding the contract event 0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryRegistration, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.WatchLogs(opts, "Registration", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRegistryRegistration)
				if err := _ControllerRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
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

// ParseRegistration is a log parse operation binding the contract event 0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseRegistration(log chainTypes.Log) (*ControllerRegistryRegistration, error) {
	event := new(ControllerRegistryRegistration)
	if err := _ControllerRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryRegistration, error) {
	var events []*ControllerRegistryRegistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524") {
			event, err := _ControllerRegistry.ParseRegistration(*log)
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

// ControllerRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the ControllerRegistry contract.
type ControllerRegistryUnregistrationIterator struct {
	Event *ControllerRegistryUnregistration // Event containing the contract specifics and raw log

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
func (it *ControllerRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRegistryUnregistration)
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
		it.Event = new(ControllerRegistryUnregistration)
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
func (it *ControllerRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRegistryUnregistration represents a Unregistration event raised by the ControllerRegistry contract.
type ControllerRegistryUnregistration struct {
	Controller common.Address
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts, controller []common.Address) (*ControllerRegistryUnregistrationIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.FilterLogs(opts, "Unregistration", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerRegistryUnregistrationIterator{contract: _ControllerRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *ControllerRegistryUnregistration, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ControllerRegistry.contract.WatchLogs(opts, "Unregistration", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRegistryUnregistration)
				if err := _ControllerRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
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

// ParseUnregistration is a log parse operation binding the contract event 0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseUnregistration(log chainTypes.Log) (*ControllerRegistryUnregistration, error) {
	event := new(ControllerRegistryUnregistration)
	if err := _ControllerRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryUnregistration, error) {
	var events []*ControllerRegistryUnregistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9") {
			event, err := _ControllerRegistry.ParseUnregistration(*log)
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
