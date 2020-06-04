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

// DataTypeRegistryABI is the input ABI used to generate the binding from.
const (
	DataTypeRegistryAddress   = "0x3b8D2133Df2F99B56d3F8c01b7db37181d7B1959"
	DataTypeRegistryTxHash    = "0x25be6e0cfefc47df47c855b02d97216dc5ae1e897babb72fccda2933f881c757"
	DataTypeRegistryCreatedAt = "0x0000000000000000000000000000000000000000000000000000000000d6d8a5"
	DataTypeRegistryABI       = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Registration\",\"signature\":\"0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Unregistration\",\"signature\":\"0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"schemaHash\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"signature\":\"0x656afdee\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"signature\":\"0x6598a1ae\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"schemaHash\",\"type\":\"bytes32\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x693ec85e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x261a323e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xbde1eee7\",\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

// DataTypeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataTypeRegistryCaller interface {
	Exists(
		ctx context.Context,
		name string,
	) (
		bool,
		error,
	)
	Get(
		ctx context.Context,
		name string,
	) (
		types.DataType,
		error,
	)
	IsOwner(
		ctx context.Context,
		name string,
		owner common.Address,
	) (
		bool,
		error,
	)
}

type dataTypeRegistryCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (_DataTypeRegistry *dataTypeRegistryCaller) Exists(ctx context.Context, name string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _DataTypeRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "exists", name)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns(types.DataType)
func (_DataTypeRegistry *dataTypeRegistryCaller) Get(ctx context.Context, name string) (types.DataType, error) {
	var (
		ret0 = new(types.DataType)
	)
	out := ret0

	err := _DataTypeRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "get", name)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (_DataTypeRegistry *dataTypeRegistryCaller) IsOwner(ctx context.Context, name string, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _DataTypeRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "isOwner", name, owner)
	return *ret0, err
}

// DataTypeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTypeRegistryTransactor interface {
	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		name string,
		schemaHash common.Hash,
	) (*chainTypes.Receipt, error)
	Unregister(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		name string,
	) (*chainTypes.Receipt, error)
}

type dataTypeRegistryTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
	backend  ablbind.ContractBackend
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (_DataTypeRegistry *dataTypeRegistryTransactor) Register(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	name string,
	schemaHash common.Hash,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _DataTypeRegistry.contract.Transact(opts, "register", name, schemaHash)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (_DataTypeRegistry *dataTypeRegistryTransactor) Unregister(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	name string,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _DataTypeRegistry.contract.Transact(opts, "unregister", name)
}

type DataTypeRegistryEvents interface {
	DataTypeRegistryEventFilterer
	DataTypeRegistryEventParser
	DataTypeRegistryEventWatcher
}

// DataTypeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataTypeRegistryEventFilterer interface {
	// Filterer
	FilterRegistration(
		opts *bind.FilterOpts,

	) (ablbind.EventIterator, error)

	// Filterer
	FilterUnregistration(
		opts *bind.FilterOpts,

	) (ablbind.EventIterator, error)
}

type DataTypeRegistryEventParser interface {
	// Parser
	ParseRegistration(log chainTypes.Log) (*DataTypeRegistryRegistration, error)
	ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*DataTypeRegistryRegistration, error)

	// Parser
	ParseUnregistration(log chainTypes.Log) (*DataTypeRegistryUnregistration, error)
	ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*DataTypeRegistryUnregistration, error)
}

type DataTypeRegistryEventWatcher interface {
	// Watcher
	WatchRegistration(
		opts *bind.WatchOpts,
		sink chan<- *DataTypeRegistryRegistration,

	) (event.Subscription, error)

	// Watcher
	WatchUnregistration(
		opts *bind.WatchOpts,
		sink chan<- *DataTypeRegistryUnregistration,

	) (event.Subscription, error)
}

type dataTypeRegistryEvents struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
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
func (_DataTypeRegistry *dataTypeRegistryEvents) FilterRegistration(opts *bind.FilterOpts) (ablbind.EventIterator, error) {

	logs, sub, err := _DataTypeRegistry.contract.FilterLogs(opts, "Registration")
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryRegistrationIterator{contract: _DataTypeRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// WatchRegistration is a free log subscription operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *dataTypeRegistryEvents) WatchRegistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryRegistration) (event.Subscription, error) {

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
func (_DataTypeRegistry *dataTypeRegistryEvents) ParseRegistration(log chainTypes.Log) (*DataTypeRegistryRegistration, error) {
	evt := new(DataTypeRegistryRegistration)
	if err := _DataTypeRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRegistrationFromReceipt parses the event from given transaction receipt.
//
// Solidity: event Registration(string name)
func (_DataTypeRegistry *dataTypeRegistryEvents) ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*DataTypeRegistryRegistration, error) {
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
func (_DataTypeRegistry *dataTypeRegistryEvents) FilterUnregistration(opts *bind.FilterOpts) (ablbind.EventIterator, error) {

	logs, sub, err := _DataTypeRegistry.contract.FilterLogs(opts, "Unregistration")
	if err != nil {
		return nil, err
	}
	return &DataTypeRegistryUnregistrationIterator{contract: _DataTypeRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *dataTypeRegistryEvents) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryUnregistration) (event.Subscription, error) {

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
func (_DataTypeRegistry *dataTypeRegistryEvents) ParseUnregistration(log chainTypes.Log) (*DataTypeRegistryUnregistration, error) {
	evt := new(DataTypeRegistryUnregistration)
	if err := _DataTypeRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseUnregistrationFromReceipt parses the event from given transaction receipt.
//
// Solidity: event Unregistration(string name)
func (_DataTypeRegistry *dataTypeRegistryEvents) ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*DataTypeRegistryUnregistration, error) {
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

// Manager is contract wrapper struct
type DataTypeRegistryContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	DataTypeRegistryCaller
	DataTypeRegistryTransactor
	DataTypeRegistryEvents
}

func NewDataTypeRegistryContract(backend ablbind.ContractBackend) (*DataTypeRegistryContract, error) {
	deployment, exist := backend.Deployment("DataTypeRegistry")
	if !exist {
		evmABI, err := abi.JSON(strings.NewReader(DataTypeRegistryABI))
		if err != nil {
			return nil, err
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(DataTypeRegistryAddress),
			common.HexToHash(DataTypeRegistryTxHash),
			new(big.Int).SetBytes(common.HexToHash(DataTypeRegistryCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, "DataTypeRegistry", backend)

	contract := &DataTypeRegistryContract{
		Deployment: deployment,
		client:     backend,

		DataTypeRegistryCaller:     &dataTypeRegistryCaller{base},
		DataTypeRegistryTransactor: &dataTypeRegistryTransactor{base, backend},
		DataTypeRegistryEvents:     &dataTypeRegistryEvents{base},
	}

	return contract, nil
}
