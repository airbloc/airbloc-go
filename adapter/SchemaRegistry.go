// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SchemaRegistryABI is the input ABI used to generate the binding from.
const SchemaRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"schemas\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"Unregistered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SchemaRegistryBin is the compiled bytecode used for deploying new contracts.
const SchemaRegistryBin = `0x60806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663143dc3f88114610071578063260a818e1461009d57806397e4fea7146100c1578063f2c298be146100e3578063f45e6aaf1461013c575b600080fd5b34801561007d57600080fd5b50610089600435610209565b604080519115158252519081900360200190f35b3480156100a957600080fd5b506100bf600160c060020a03196004351661021e565b005b3480156100cd57600080fd5b50610089600160c060020a031960043516610422565b3480156100ef57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100bf9436949293602493928401919081908401838280828437509497506104569650505050505050565b34801561014857600080fd5b5061015e600160c060020a031960043516610658565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101cd5781810151838201526020016101b5565b50505050905090810190601f1680156101fa5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b60016020526000908152604090205460ff1681565b600160c060020a031981166000908152602081905260408120805490919073ffffffffffffffffffffffffffffffffffffffff1633146102bf57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c79206f776e65722063616e20646f207468697300000000000000000000604482015290519081900360640190fd5b8160010160405160200180828054600181600116156101000203166002900480156103215780601f106102ff576101008083540402835291820191610321565b820191906000526020600020905b81548152906001019060200180831161030d575b50509150506040516020818303038152906040526040518082805190602001908083835b602083106103645780518252601f199092019160209182019101610345565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008181526001808452858220805460ff19169055600160c060020a03198b16825292819052938420805473ffffffffffffffffffffffffffffffffffffffff1916815590965094509192506103e69184019050826107cf565b5050604051600160c060020a03198416907f406b9f2601ae72913a540201177507f1183b2a84b763bb9cd282c1c1cf6ad0b990600090a2505050565b600160c060020a03191660009081526020819052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b6000806000836040516020018082805190602001908083835b6020831061048e5780518252601f19909201916020918201910161046f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106104f15780518252601f1990920191602091820191016104d2565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008181526001909252929020549196505060ff1615915061059e905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f54686520736368656d6120616c72656164792065786973747321000000000000604482015290519081900360640190fd5b6105a88333610716565b600160c060020a03198116600090815260208181526040909120805473ffffffffffffffffffffffffffffffffffffffff191633178155865192945092506105f7916001840191870190610816565b50600083815260016020818152604092839020805460ff191690921790915581513381529151600160c060020a03198516927f10e48ff9b9d60a8908f9f0baa8372e4ab117adbe99b34c21451e83e5969470a192908290030190a250505050565b600060208181529181526040908190208054600180830180548551600261010094831615949094026000190190911692909204601f810187900487028301870190955284825273ffffffffffffffffffffffffffffffffffffffff909216949293909283018282801561070c5780601f106106e15761010080835404028352916020019161070c565b820191906000526020600020905b8154815290600101906020018083116106ef57829003601f168201915b5050505050905082565b604080516c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff84160260208083019190915243603483015260548083018690528351808403909101815260749092019283905281516000938392909182918401908083835b6020831061079b5780518252601f19909201916020918201910161077c565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120979650505050505050565b50805460018160011615610100020316600290046000825580601f106107f55750610813565b601f0160209004906000526020600020908101906108139190610894565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061085757805160ff1916838001178555610884565b82800160010185558215610884579182015b82811115610884578251825591602001919060010190610869565b50610890929150610894565b5090565b6108ae91905b80821115610890576000815560010161089a565b905600a165627a7a72305820b6f92cca12ad29b7f705dff28e950e828a33e12c43a153153db5ce4a0ea487d30029`

// DeploySchemaRegistry deploys a new Ethereum contract, binding an instance of SchemaRegistry to it.
func DeploySchemaRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SchemaRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(SchemaRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SchemaRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SchemaRegistry{SchemaRegistryCaller: SchemaRegistryCaller{contract: contract}, SchemaRegistryTransactor: SchemaRegistryTransactor{contract: contract}, SchemaRegistryFilterer: SchemaRegistryFilterer{contract: contract}}, nil
}

// SchemaRegistry is an auto generated Go binding around an Ethereum contract.
type SchemaRegistry struct {
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

// NewSchemaRegistry creates a new instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistry(address common.Address, backend bind.ContractBackend) (*SchemaRegistry, error) {
	contract, err := bindSchemaRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistry{SchemaRegistryCaller: SchemaRegistryCaller{contract: contract}, SchemaRegistryTransactor: SchemaRegistryTransactor{contract: contract}, SchemaRegistryFilterer: SchemaRegistryFilterer{contract: contract}}, nil
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

// SchemaRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the SchemaRegistry contract.
type SchemaRegistryRegisteredIterator struct {
	Event *SchemaRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *SchemaRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SchemaRegistryRegistered)
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
		it.Event = new(SchemaRegistryRegistered)
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
func (it *SchemaRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SchemaRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SchemaRegistryRegistered represents a Registered event raised by the SchemaRegistry contract.
type SchemaRegistryRegistered struct {
	Id    [8]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x10e48ff9b9d60a8908f9f0baa8372e4ab117adbe99b34c21451e83e5969470a1.
//
// Solidity: e Registered(_id indexed bytes8, owner address)
func (_SchemaRegistry *SchemaRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, _id [][8]byte) (*SchemaRegistryRegisteredIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _SchemaRegistry.contract.FilterLogs(opts, "Registered", _idRule)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryRegisteredIterator{contract: _SchemaRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x10e48ff9b9d60a8908f9f0baa8372e4ab117adbe99b34c21451e83e5969470a1.
//
// Solidity: e Registered(_id indexed bytes8, owner address)
func (_SchemaRegistry *SchemaRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *SchemaRegistryRegistered, _id [][8]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _SchemaRegistry.contract.WatchLogs(opts, "Registered", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SchemaRegistryRegistered)
				if err := _SchemaRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// SchemaRegistryUnregisteredIterator is returned from FilterUnregistered and is used to iterate over the raw logs and unpacked data for Unregistered events raised by the SchemaRegistry contract.
type SchemaRegistryUnregisteredIterator struct {
	Event *SchemaRegistryUnregistered // Event containing the contract specifics and raw log

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
func (it *SchemaRegistryUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SchemaRegistryUnregistered)
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
		it.Event = new(SchemaRegistryUnregistered)
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
func (it *SchemaRegistryUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SchemaRegistryUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SchemaRegistryUnregistered represents a Unregistered event raised by the SchemaRegistry contract.
type SchemaRegistryUnregistered struct {
	Id  [8]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnregistered is a free log retrieval operation binding the contract event 0x406b9f2601ae72913a540201177507f1183b2a84b763bb9cd282c1c1cf6ad0b9.
//
// Solidity: e Unregistered(_id indexed bytes8)
func (_SchemaRegistry *SchemaRegistryFilterer) FilterUnregistered(opts *bind.FilterOpts, _id [][8]byte) (*SchemaRegistryUnregisteredIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _SchemaRegistry.contract.FilterLogs(opts, "Unregistered", _idRule)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryUnregisteredIterator{contract: _SchemaRegistry.contract, event: "Unregistered", logs: logs, sub: sub}, nil
}

// WatchUnregistered is a free log subscription operation binding the contract event 0x406b9f2601ae72913a540201177507f1183b2a84b763bb9cd282c1c1cf6ad0b9.
//
// Solidity: e Unregistered(_id indexed bytes8)
func (_SchemaRegistry *SchemaRegistryFilterer) WatchUnregistered(opts *bind.WatchOpts, sink chan<- *SchemaRegistryUnregistered, _id [][8]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _SchemaRegistry.contract.WatchLogs(opts, "Unregistered", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SchemaRegistryUnregistered)
				if err := _SchemaRegistry.contract.UnpackLog(event, "Unregistered", log); err != nil {
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
