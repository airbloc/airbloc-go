package contracts

import (
	"errors"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	types "github.com/airbloc/airbloc-go/shared/adapter/types"
	platform "github.com/klaytn/klaytn"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// ControllerRegistryABI is the input ABI used to generate the binding from.
const (
	ControllerRegistryAddress   = "0x23FB993D540a16a3c2337039353bf4cEa6F9EB6B"
	ControllerRegistryTxHash    = "0x729230160035ac58558939906b0fdda903189be259ffd9e1113d5fb79e8eb106"
	ControllerRegistryCreatedAt = "0x000000000000000000000000000000000000000000000000000000000063b9e4"
	ControllerRegistryABI       = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"signature\":\"0x715018a6\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"0x8da5cb5b\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x8f32d59b\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"signature\":\"0xf2fde38b\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"Registration\",\"signature\":\"0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"Unregistration\",\"signature\":\"0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"signature\":\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"controllerAddr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"signature\":\"0x4420e486\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"usersCount\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0xc2bc2efc\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xf6a3d24e\",\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

// ControllerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerRegistryCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewControllerRegistryCaller(contract *ablbind.BoundContract) ControllerRegistryCaller {
	return ControllerRegistryCaller{contract: contract}
}

// ControllerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerRegistryTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewControllerRegistryTransactor(contract *ablbind.BoundContract) ControllerRegistryTransactor {
	return ControllerRegistryTransactor{contract: contract}
}

// ControllerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerRegistryFilterer struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewControllerRegistryFilterer(contract *ablbind.BoundContract) ControllerRegistryFilterer {
	return ControllerRegistryFilterer{contract: contract}
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

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (_ControllerRegistry *ControllerRegistryTransactor) Register(opts *ablbind.TransactOpts, controllerAddr common.Address) (*chainTypes.Transaction, error) {
	return _ControllerRegistry.contract.Transact(opts, "register", controllerAddr)
}

// ControllerRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ControllerRegistry contract.
type ControllerRegistryOwnershipTransferredIterator struct {
	Evt *ControllerRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ControllerRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ControllerRegistryOwnershipTransferred)
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
		it.Evt = new(ControllerRegistryOwnershipTransferred)
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
func (it *ControllerRegistryOwnershipTransferredIterator) Event() interface{} {
	return it.Evt
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
func (_ControllerRegistry *ControllerRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (ablbind.EventIterator, error) {

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
				evt := new(ControllerRegistryOwnershipTransferred)
				if err := _ControllerRegistry.contract.UnpackLog(evt, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseOwnershipTransferred(log chainTypes.Log) (*ControllerRegistryOwnershipTransferred, error) {
	evt := new(ControllerRegistryOwnershipTransferred)
	if err := _ControllerRegistry.contract.UnpackLog(evt, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterOwnershipTransferred parses the event from given transaction receipt.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseOwnershipTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryOwnershipTransferred, error) {
	var evts []*ControllerRegistryOwnershipTransferred
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0") {
			evt, err := _ControllerRegistry.ParseOwnershipTransferred(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("OwnershipTransferred event not found")
	}
	return evts, nil
}

// ControllerRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the ControllerRegistry contract.
type ControllerRegistryRegistrationIterator struct {
	Evt *ControllerRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *ControllerRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ControllerRegistryRegistration)
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
		it.Evt = new(ControllerRegistryRegistration)
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
func (it *ControllerRegistryRegistrationIterator) Event() interface{} {
	return it.Evt
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
func (_ControllerRegistry *ControllerRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, controller []common.Address) (ablbind.EventIterator, error) {

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
				evt := new(ControllerRegistryRegistration)
				if err := _ControllerRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
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

// ParseRegistration is a log parse operation binding the contract event 0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseRegistration(log chainTypes.Log) (*ControllerRegistryRegistration, error) {
	evt := new(ControllerRegistryRegistration)
	if err := _ControllerRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryRegistration, error) {
	var evts []*ControllerRegistryRegistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x478f5152d8fc568db3f8de9fb402fd9d98a1a7541ecfe434e59cf574fbfc5524") {
			evt, err := _ControllerRegistry.ParseRegistration(*log)
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

// ControllerRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the ControllerRegistry contract.
type ControllerRegistryUnregistrationIterator struct {
	Evt *ControllerRegistryUnregistration // Event containing the contract specifics and raw log

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
func (it *ControllerRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ControllerRegistryUnregistration)
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
		it.Evt = new(ControllerRegistryUnregistration)
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
func (it *ControllerRegistryUnregistrationIterator) Event() interface{} {
	return it.Evt
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
func (_ControllerRegistry *ControllerRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts, controller []common.Address) (ablbind.EventIterator, error) {

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
				evt := new(ControllerRegistryUnregistration)
				if err := _ControllerRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
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

// ParseUnregistration is a log parse operation binding the contract event 0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseUnregistration(log chainTypes.Log) (*ControllerRegistryUnregistration, error) {
	evt := new(ControllerRegistryUnregistration)
	if err := _ControllerRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(address indexed controller)
func (_ControllerRegistry *ControllerRegistryFilterer) ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*ControllerRegistryUnregistration, error) {
	var evts []*ControllerRegistryUnregistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x2171d18d6eaa5385a17d6cacd86394726517e8399c558ab99acf728be83f5bb9") {
			evt, err := _ControllerRegistry.ParseUnregistration(*log)
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
