package contracts

import (
	"errors"

	ablbind "github.com/airbloc/airbloc-go/bind"
	types "github.com/airbloc/airbloc-go/bind/types"
	platform "github.com/klaytn/klaytn"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// DataTypeRegistryABI is the input ABI used to generate the binding from.
const (
	DataTypeRegistryAddress   = "0x3E6536e6355B760c59d9f7e48f2A06d4dA8E2AC9"
	DataTypeRegistryTxHash    = "0x6963c52a27251fde0305b5f21446d5fdb3e66b878d732f2b0393f97c9c391fe9"
	DataTypeRegistryCreatedAt = "0x000000000000000000000000000000000000000000000000000000000063b9e7"
	DataTypeRegistryABI       = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Registration\",\"signature\":\"0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Unregistration\",\"signature\":\"0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"schemaHash\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"signature\":\"0x656afdee\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"schemaHash\",\"type\":\"bytes32\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x693ec85e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x261a323e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xbde1eee7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"signature\":\"0x6598a1ae\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
)

// DataTypeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataTypeRegistryCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewDataTypeRegistryCaller(contract *ablbind.BoundContract) DataTypeRegistryCaller {
	return DataTypeRegistryCaller{contract: contract}
}

// DataTypeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTypeRegistryTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewDataTypeRegistryTransactor(contract *ablbind.BoundContract) DataTypeRegistryTransactor {
	return DataTypeRegistryTransactor{contract: contract}
}

// DataTypeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataTypeRegistryFilterer struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewDataTypeRegistryFilterer(contract *ablbind.BoundContract) DataTypeRegistryFilterer {
	return DataTypeRegistryFilterer{contract: contract}
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (_DataTypeRegistry *DataTypeRegistryCaller) Exists(opts *bind.CallOpts, name string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DataTypeRegistry.contract.Call(opts, out, "exists", name)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns(types.DataType)
func (_DataTypeRegistry *DataTypeRegistryCaller) Get(opts *bind.CallOpts, name string) (types.DataType, error) {
	var (
		ret0 = new(types.DataType)
	)
	out := ret0
	err := _DataTypeRegistry.contract.Call(opts, out, "get", name)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (_DataTypeRegistry *DataTypeRegistryCaller) IsOwner(opts *bind.CallOpts, name string, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DataTypeRegistry.contract.Call(opts, out, "isOwner", name, owner)
	return *ret0, err
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (_DataTypeRegistry *DataTypeRegistryTransactor) Register(opts *ablbind.TransactOpts, name string, schemaHash common.Hash) (*chainTypes.Transaction, error) {
	return _DataTypeRegistry.contract.Transact(opts, "register", name, schemaHash)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (_DataTypeRegistry *DataTypeRegistryTransactor) Unregister(opts *ablbind.TransactOpts, name string) (*chainTypes.Transaction, error) {
	return _DataTypeRegistry.contract.Transact(opts, "unregister", name)
}

// DataTypeRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the DataTypeRegistry contract.
type DataTypeRegistryRegistrationIterator struct {
	Evt *DataTypeRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *DataTypeRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(DataTypeRegistryRegistration)
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
		it.Evt = new(DataTypeRegistryRegistration)
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
func (it *DataTypeRegistryRegistrationIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataTypeRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataTypeRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataTypeRegistryRegistration represents a Registration event raised by the DataTypeRegistry contract.
type DataTypeRegistryRegistration struct {
	Name string
	Raw  chainTypes.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) FilterRegistration(opts *bind.FilterOpts) (ablbind.EventIterator, error) {

	logs, sub, err := _DataTypeRegistry.contract.FilterLogs(opts, "Registration")
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryRegistrationIterator{contract: _DataTypeRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// WatchRegistration is a free log subscription operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryRegistration) (event.Subscription, error) {

	logs, sub, err := _DataTypeRegistry.contract.WatchLogs(opts, "Registration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(DataTypeRegistryRegistration)
				if err := _DataTypeRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
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

// ParseRegistration is a log parse operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) ParseRegistration(log chainTypes.Log) (*DataTypeRegistryRegistration, error) {
	evt := new(DataTypeRegistryRegistration)
	if err := _DataTypeRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*DataTypeRegistryRegistration, error) {
	var evts []*DataTypeRegistryRegistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948") {
			evt, err := _DataTypeRegistry.ParseRegistration(*log)
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

// DataTypeRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the DataTypeRegistry contract.
type DataTypeRegistryUnregistrationIterator struct {
	Evt *DataTypeRegistryUnregistration // Event containing the contract specifics and raw log

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
func (it *DataTypeRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(DataTypeRegistryUnregistration)
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
		it.Evt = new(DataTypeRegistryUnregistration)
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
func (it *DataTypeRegistryUnregistrationIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataTypeRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataTypeRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataTypeRegistryUnregistration represents a Unregistration event raised by the DataTypeRegistry contract.
type DataTypeRegistryUnregistration struct {
	Name string
	Raw  chainTypes.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts) (ablbind.EventIterator, error) {

	logs, sub, err := _DataTypeRegistry.contract.FilterLogs(opts, "Unregistration")
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryUnregistrationIterator{contract: _DataTypeRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryUnregistration) (event.Subscription, error) {

	logs, sub, err := _DataTypeRegistry.contract.WatchLogs(opts, "Unregistration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(DataTypeRegistryUnregistration)
				if err := _DataTypeRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
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

// ParseUnregistration is a log parse operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) ParseUnregistration(log chainTypes.Log) (*DataTypeRegistryUnregistration, error) {
	evt := new(DataTypeRegistryUnregistration)
	if err := _DataTypeRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *DataTypeRegistryFilterer) ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*DataTypeRegistryUnregistration, error) {
	var evts []*DataTypeRegistryUnregistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83") {
			evt, err := _DataTypeRegistry.ParseUnregistration(*log)
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
