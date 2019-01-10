// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"
	"strings"

	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/blockchain/bind"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	_ = bind.NewKeyedTransactor
	_ = ablCommon.HexToID
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SchemaRegistryABI is the input ABI used to generate the binding from.
const SchemaRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"schemas\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registrar\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"Registration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"Unregistration\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SchemaRegistry is an auto generated Go binding around an Ethereum contract.
type SchemaRegistry struct {
	Address                  common.Address
	SchemaRegistryCaller     // Read-only binding to the contract
	SchemaRegistryTransactor // Write-only binding to the contract
	SchemaRegistryFilterer   // Log filterer for contract events
}

// SchemaRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SchemaRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SchemaRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SchemaRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SchemaRegistrySession struct {
	Contract     *SchemaRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SchemaRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SchemaRegistryCallerSession struct {
	Contract *SchemaRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SchemaRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SchemaRegistryTransactorSession struct {
	Contract     *SchemaRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SchemaRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SchemaRegistryRaw struct {
	Contract *SchemaRegistry // Generic contract binding to access the raw methods on
}

// SchemaRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SchemaRegistryCallerRaw struct {
	Contract *SchemaRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SchemaRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SchemaRegistryTransactorRaw struct {
	Contract *SchemaRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

//
//
//
//	type Schema struct {
//		Name	string
//		Owner	common.Address
//
//	}
//

func init() {
	// convenient hacks for blockchain.Client
	blockchain.ContractList["SchemaRegistry"] = (&SchemaRegistry{}).new
	blockchain.RegisterSelector("0xf2c298be", "register(string)")
	blockchain.RegisterSelector("0x260a818e", "unregister(bytes8)")

}

// NewSchemaRegistry creates a new instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistry(address common.Address, backend bind.ContractBackend) (*SchemaRegistry, error) {
	contract, err := bindSchemaRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistry{
		Address:                  address,
		SchemaRegistryCaller:     SchemaRegistryCaller{contract: contract},
		SchemaRegistryTransactor: SchemaRegistryTransactor{contract: contract},
		SchemaRegistryFilterer:   SchemaRegistryFilterer{contract: contract},
	}, nil
}

// NewSchemaRegistryCaller creates a new read-only instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryCaller(address common.Address, caller bind.ContractCaller) (*SchemaRegistryCaller, error) {
	contract, err := bindSchemaRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryCaller{contract: contract}, nil
}

// NewSchemaRegistryTransactor creates a new write-only instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SchemaRegistryTransactor, error) {
	contract, err := bindSchemaRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryTransactor{contract: contract}, nil
}

// NewSchemaRegistryFilterer creates a new log filterer instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SchemaRegistryFilterer, error) {
	contract, err := bindSchemaRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryFilterer{contract: contract}, nil
}

// bindSchemaRegistry binds a generic wrapper to an already deployed contract.
func bindSchemaRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SchemaRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_SchemaRegistry *SchemaRegistry) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewSchemaRegistry(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SchemaRegistry *SchemaRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SchemaRegistry.Contract.SchemaRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SchemaRegistry *SchemaRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.SchemaRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SchemaRegistry *SchemaRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.SchemaRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SchemaRegistry *SchemaRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SchemaRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SchemaRegistry *SchemaRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SchemaRegistry *SchemaRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_SchemaRegistry *SchemaRegistryCaller) Exists(opts *bind.CallOpts, _id [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SchemaRegistry.contract.Call(opts, out, "exists", _id)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_SchemaRegistry *SchemaRegistrySession) Exists(_id [8]byte) (bool, error) {
	return _SchemaRegistry.Contract.Exists(&_SchemaRegistry.CallOpts, _id)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_SchemaRegistry *SchemaRegistryCallerSession) Exists(_id [8]byte) (bool, error) {
	return _SchemaRegistry.Contract.Exists(&_SchemaRegistry.CallOpts, _id)
}

// NameExists is a free data retrieval call binding the contract method 0x143dc3f8.
//
// Solidity: function nameExists( bytes32) constant returns(bool)
func (_SchemaRegistry *SchemaRegistryCaller) NameExists(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SchemaRegistry.contract.Call(opts, out, "nameExists", arg0)
	return *ret0, err
}

// NameExists is a free data retrieval call binding the contract method 0x143dc3f8.
//
// Solidity: function nameExists( bytes32) constant returns(bool)
func (_SchemaRegistry *SchemaRegistrySession) NameExists(arg0 [32]byte) (bool, error) {
	return _SchemaRegistry.Contract.NameExists(&_SchemaRegistry.CallOpts, arg0)
}

// NameExists is a free data retrieval call binding the contract method 0x143dc3f8.
//
// Solidity: function nameExists( bytes32) constant returns(bool)
func (_SchemaRegistry *SchemaRegistryCallerSession) NameExists(arg0 [32]byte) (bool, error) {
	return _SchemaRegistry.Contract.NameExists(&_SchemaRegistry.CallOpts, arg0)
}

// Schemas is a free data retrieval call binding the contract method 0xf45e6aaf.
//
// Solidity: function schemas( bytes8) constant returns(owner address, name string)
func (_SchemaRegistry *SchemaRegistryCaller) Schemas(opts *bind.CallOpts, arg0 [8]byte) (struct {
	Owner common.Address
	Name  string
}, error) {
	ret := new(struct {
		Owner common.Address
		Name  string
	})
	out := ret
	err := _SchemaRegistry.contract.Call(opts, out, "schemas", arg0)
	return *ret, err
}

// Schemas is a free data retrieval call binding the contract method 0xf45e6aaf.
//
// Solidity: function schemas( bytes8) constant returns(owner address, name string)
func (_SchemaRegistry *SchemaRegistrySession) Schemas(arg0 [8]byte) (struct {
	Owner common.Address
	Name  string
}, error) {
	return _SchemaRegistry.Contract.Schemas(&_SchemaRegistry.CallOpts, arg0)
}

// Schemas is a free data retrieval call binding the contract method 0xf45e6aaf.
//
// Solidity: function schemas( bytes8) constant returns(owner address, name string)
func (_SchemaRegistry *SchemaRegistryCallerSession) Schemas(arg0 [8]byte) (struct {
	Owner common.Address
	Name  string
}, error) {
	return _SchemaRegistry.Contract.Schemas(&_SchemaRegistry.CallOpts, arg0)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_name string) returns()
func (_SchemaRegistry *SchemaRegistryTransactor) Register(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _SchemaRegistry.contract.Transact(opts, "register", _name)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_name string) returns()
func (_SchemaRegistry *SchemaRegistrySession) Register(_name string) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Register(&_SchemaRegistry.TransactOpts, _name)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_name string) returns()
func (_SchemaRegistry *SchemaRegistryTransactorSession) Register(_name string) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Register(&_SchemaRegistry.TransactOpts, _name)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_SchemaRegistry *SchemaRegistryTransactor) Unregister(opts *bind.TransactOpts, _id [8]byte) (*types.Transaction, error) {
	return _SchemaRegistry.contract.Transact(opts, "unregister", _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_SchemaRegistry *SchemaRegistrySession) Unregister(_id [8]byte) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Unregister(&_SchemaRegistry.TransactOpts, _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_SchemaRegistry *SchemaRegistryTransactorSession) Unregister(_id [8]byte) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Unregister(&_SchemaRegistry.TransactOpts, _id)
}

// SchemaRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the SchemaRegistry contract.
type SchemaRegistryRegistrationIterator struct {
	Event *SchemaRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *SchemaRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SchemaRegistryRegistration)
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
		it.Event = new(SchemaRegistryRegistration)
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
func (it *SchemaRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SchemaRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SchemaRegistryRegistration represents a Registration event raised by the SchemaRegistry contract.
type SchemaRegistryRegistration struct {
	Registrar common.Address
	Id        [8]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0x34b195a2b14d8eac732b11770bab8ed3823e96642f703b1a973f4a2981208d0f.
//
// Solidity: e Registration(registrar indexed address, _id bytes8)
func (_SchemaRegistry *SchemaRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, registrar []common.Address) (*SchemaRegistryRegistrationIterator, error) {

	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _SchemaRegistry.contract.FilterLogs(opts, "Registration", registrarRule)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryRegistrationIterator{contract: _SchemaRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: e Registration(registrar indexed address, _id bytes8)
func (_SchemaRegistry *SchemaRegistryFilterer) ParseRegistrationFromReceipt(receipt *types.Receipt) (*SchemaRegistryRegistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x34b195a2b14d8eac732b11770bab8ed3823e96642f703b1a973f4a2981208d0f") {
			event := new(SchemaRegistryRegistration)
			if err := _SchemaRegistry.contract.UnpackLog(event, "Registration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Registration event not found")
}

// WatchRegistration is a free log subscription operation binding the contract event 0x34b195a2b14d8eac732b11770bab8ed3823e96642f703b1a973f4a2981208d0f.
//
// Solidity: e Registration(registrar indexed address, _id bytes8)
func (_SchemaRegistry *SchemaRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *SchemaRegistryRegistration, registrar []common.Address) (event.Subscription, error) {

	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _SchemaRegistry.contract.WatchLogs(opts, "Registration", registrarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SchemaRegistryRegistration)
				if err := _SchemaRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
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

// SchemaRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the SchemaRegistry contract.
type SchemaRegistryUnregistrationIterator struct {
	Event *SchemaRegistryUnregistration // Event containing the contract specifics and raw log

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
func (it *SchemaRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SchemaRegistryUnregistration)
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
		it.Event = new(SchemaRegistryUnregistration)
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
func (it *SchemaRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SchemaRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SchemaRegistryUnregistration represents a Unregistration event raised by the SchemaRegistry contract.
type SchemaRegistryUnregistration struct {
	Id  [8]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x8c36f878328ed4dfe683ccea03edd5a0c360e665285312270adb9a22592367fb.
//
// Solidity: e Unregistration(_id indexed bytes8)
func (_SchemaRegistry *SchemaRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts, _id [][8]byte) (*SchemaRegistryUnregistrationIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _SchemaRegistry.contract.FilterLogs(opts, "Unregistration", _idRule)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryUnregistrationIterator{contract: _SchemaRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: e Unregistration(_id indexed bytes8)
func (_SchemaRegistry *SchemaRegistryFilterer) ParseUnregistrationFromReceipt(receipt *types.Receipt) (*SchemaRegistryUnregistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8c36f878328ed4dfe683ccea03edd5a0c360e665285312270adb9a22592367fb") {
			event := new(SchemaRegistryUnregistration)
			if err := _SchemaRegistry.contract.UnpackLog(event, "Unregistration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Unregistration event not found")
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x8c36f878328ed4dfe683ccea03edd5a0c360e665285312270adb9a22592367fb.
//
// Solidity: e Unregistration(_id indexed bytes8)
func (_SchemaRegistry *SchemaRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *SchemaRegistryUnregistration, _id [][8]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _SchemaRegistry.contract.WatchLogs(opts, "Unregistration", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SchemaRegistryUnregistration)
				if err := _SchemaRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
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
