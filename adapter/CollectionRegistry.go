// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CollectionRegistryABI is the input ABI used to generate the binding from.
const CollectionRegistryABI = "[{\"inputs\":[{\"name\":\"_appReg\",\"type\":\"address\"},{\"name\":\"_schemaReg\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_colectionId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_colectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_appId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_schemaId\",\"type\":\"bytes8\"}],\"name\":\"Unregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_collectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"Allowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_collectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"Denied\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes8\"},{\"name\":\"_schemaId\",\"type\":\"bytes8\"},{\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_uid\",\"type\":\"bytes32\"}],\"name\":\"checkAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CollectionRegistryBin is the compiled bytecode used for deploying new contracts.
const CollectionRegistryBin = `0x6080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631a0919dc81146100875780631fed449f146100a1578063399e0792146100bc5780634fe929c2146100e85780638eaa6ac014610103578063a3b42cba1461013d578063d42e715514610168575b600080fd5b34801561009357600080fd5b5061009f600435610183565b005b3480156100ad57600080fd5b5061009f600435602435610390565b3480156100c857600080fd5b506100d46004356103e5565b604080519115158252519081900360200190f35b3480156100f457600080fd5b5061009f600435602435610409565b34801561010f57600080fd5b5061011b60043561045b565b60408051928352600160c060020a031990911660208301528051918290030190f35b34801561014957600080fd5b5061009f600160c060020a031960043581169060243516604435610493565b34801561017457600080fd5b506100d46004356024356106f5565b61018b610933565b600080548382526002602090815260408084205481517f672b7beb00000000000000000000000000000000000000000000000000000000815260c060020a909102600160c060020a0319166004820152336024820152905173ffffffffffffffffffffffffffffffffffffffff9093169363672b7beb93604480840194939192918390030190829087803b15801561022257600080fd5b505af1158015610236573d6000803e3d6000fd5b505050506040513d602081101561024c57600080fd5b505115156102ca576040805160e560020a62461bcd02815260206004820152602160248201527f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686960448201527f7000000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b50600081815260026020818152604080842081516060810183528154600160c060020a031960c060020a80830282168452680100000000000000008304028116838701908152855180870187526001860180548252868a018054838b0152868901929092528b8b52989097526fffffffffffffffffffffffffffffffff199092169093559486905592859055925182519151929490841693919091169185917f0231d91ceaa0291166e678222375c22a49172b2641f9003dcdebd63e1160cc7091a45050565b6000828152600260209081526040808320848452600301909152808220805460ff1916600117905551829184917fa22515132971f50f788da1869934dff79436ef9486db69a7a59731a5fb61689d9190a35050565b60009081526002602052604090205460c060020a02600160c060020a031916151590565b6000828152600260209081526040808320848452600301909152808220805460ff1916905551829184917f4d28190f1b112cd85f6380723c76d76de35cb4a80b5ab017d01f320b25db009e9190a35050565b600080600061046984610719565b5460c060020a808202600160c060020a031916966801000000000000000090920402945092505050565b60008054604080517f672b7beb000000000000000000000000000000000000000000000000000000008152600160c060020a031987166004820152336024820152905173ffffffffffffffffffffffffffffffffffffffff9092169163672b7beb9160448082019260209290919082900301818787803b15801561051657600080fd5b505af115801561052a573d6000803e3d6000fd5b505050506040513d602081101561054057600080fd5b505115156105be576040805160e560020a62461bcd02815260206004820152602160248201527f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686960448201527f7000000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60408051600160c060020a0319808716602080840191909152908616602883015282516010818403018152603090920192839052815191929182918401908083835b6020831061061f5780518252601f199092019160209182019101610600565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061065984848461072a565b6000828152600260208181526040808420855181548785015160c060020a9081900468010000000000000000026fffffffffffffffff0000000000000000199190930467ffffffffffffffff1990921691909117161781559481015180516001870155909101519390910192909255905182917f10906fae603eebfac53ddc0f103bee8a044dd7643c425c7a90f921dfa15ef62c91a250505050565b60009182526002602090815260408084209284526003909201905290205460ff1690565b600090815260026020526040902090565b610732610933565b600154604080517f97e4fea7000000000000000000000000000000000000000000000000000000008152600160c060020a031986166004820152905173ffffffffffffffffffffffffffffffffffffffff909216916397e4fea7916024808201926020929091908290030181600087803b1580156107af57600080fd5b505af11580156107c3573d6000803e3d6000fd5b505050506040513d60208110156107d957600080fd5b50511515610831576040805160e560020a62461bcd02815260206004820152601b60248201527f676976656e20736368656d6120646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b610844600160c060020a031984166103e5565b151561089a576040805160e560020a62461bcd02815260206004820152601960248201527f636f6c6c656374696f6e20616c72656164792065786973747300000000000000604482015290519081900360640190fd5b6060604051908101604052808577ffffffffffffffffffffffffffffffffffffffffffffffff191681526020018477ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200160408051908101604052808581526020016109158668056bc75e2d6310000061092190919063ffffffff16565b90529052949350505050565b60008282111561092d57fe5b50900390565b6040805160808101825260008082526020820152908101610952610957565b905290565b6040805180820190915260008082526020820152905600a165627a7a72305820ebbcb261213d9655625d3c50fdd3da110dcf99a45e50cae608d450f568e88edc0029`

// DeployCollectionRegistry deploys a new Ethereum contract, binding an instance of CollectionRegistry to it.
func DeployCollectionRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _appReg common.Address, _schemaReg common.Address) (common.Address, *types.Transaction, *CollectionRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CollectionRegistryBin), backend, _appReg, _schemaReg)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CollectionRegistry{CollectionRegistryCaller: CollectionRegistryCaller{contract: contract}, CollectionRegistryTransactor: CollectionRegistryTransactor{contract: contract}, CollectionRegistryFilterer: CollectionRegistryFilterer{contract: contract}}, nil
}

// CollectionRegistry is an auto generated Go binding around an Ethereum contract.
type CollectionRegistry struct {
	CollectionRegistryCaller     // Read-only binding to the contract
	CollectionRegistryTransactor // Write-only binding to the contract
	CollectionRegistryFilterer   // Log filterer for contract events
}

// CollectionRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectionRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectionRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectionRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectionRegistrySession struct {
	Contract     *CollectionRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CollectionRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectionRegistryCallerSession struct {
	Contract *CollectionRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CollectionRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectionRegistryTransactorSession struct {
	Contract     *CollectionRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CollectionRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectionRegistryRaw struct {
	Contract *CollectionRegistry // Generic contract binding to access the raw methods on
}

// CollectionRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectionRegistryCallerRaw struct {
	Contract *CollectionRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CollectionRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectionRegistryTransactorRaw struct {
	Contract *CollectionRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollectionRegistry creates a new instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistry(address common.Address, backend bind.ContractBackend) (*CollectionRegistry, error) {
	contract, err := bindCollectionRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistry{CollectionRegistryCaller: CollectionRegistryCaller{contract: contract}, CollectionRegistryTransactor: CollectionRegistryTransactor{contract: contract}, CollectionRegistryFilterer: CollectionRegistryFilterer{contract: contract}}, nil
}

// NewCollectionRegistryCaller creates a new read-only instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryCaller(address common.Address, caller bind.ContractCaller) (*CollectionRegistryCaller, error) {
	contract, err := bindCollectionRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryCaller{contract: contract}, nil
}

// NewCollectionRegistryTransactor creates a new write-only instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectionRegistryTransactor, error) {
	contract, err := bindCollectionRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryTransactor{contract: contract}, nil
}

// NewCollectionRegistryFilterer creates a new log filterer instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectionRegistryFilterer, error) {
	contract, err := bindCollectionRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryFilterer{contract: contract}, nil
}

// bindCollectionRegistry binds a generic wrapper to an already deployed contract.
func bindCollectionRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionRegistry *CollectionRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CollectionRegistry.Contract.CollectionRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionRegistry *CollectionRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.CollectionRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionRegistry *CollectionRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.CollectionRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionRegistry *CollectionRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CollectionRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionRegistry *CollectionRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionRegistry *CollectionRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_id bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) Check(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "check", _id)
	return *ret0, err
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_id bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) Check(_id [32]byte) (bool, error) {
	return _CollectionRegistry.Contract.Check(&_CollectionRegistry.CallOpts, _id)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(_id bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) Check(_id [32]byte) (bool, error) {
	return _CollectionRegistry.Contract.Check(&_CollectionRegistry.CallOpts, _id)
}

// CheckAllowed is a free data retrieval call binding the contract method 0xd42e7155.
//
// Solidity: function checkAllowed(_id bytes32, _uid bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) CheckAllowed(opts *bind.CallOpts, _id [32]byte, _uid [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "checkAllowed", _id, _uid)
	return *ret0, err
}

// CheckAllowed is a free data retrieval call binding the contract method 0xd42e7155.
//
// Solidity: function checkAllowed(_id bytes32, _uid bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) CheckAllowed(_id [32]byte, _uid [32]byte) (bool, error) {
	return _CollectionRegistry.Contract.CheckAllowed(&_CollectionRegistry.CallOpts, _id, _uid)
}

// CheckAllowed is a free data retrieval call binding the contract method 0xd42e7155.
//
// Solidity: function checkAllowed(_id bytes32, _uid bytes32) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) CheckAllowed(_id [32]byte, _uid [32]byte) (bool, error) {
	return _CollectionRegistry.Contract.CheckAllowed(&_CollectionRegistry.CallOpts, _id, _uid)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_id bytes32) constant returns(bytes32, bytes8)
func (_CollectionRegistry *CollectionRegistryCaller) Get(opts *bind.CallOpts, _id [32]byte) ([32]byte, [8]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([8]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _CollectionRegistry.contract.Call(opts, out, "get", _id)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_id bytes32) constant returns(bytes32, bytes8)
func (_CollectionRegistry *CollectionRegistrySession) Get(_id [32]byte) ([32]byte, [8]byte, error) {
	return _CollectionRegistry.Contract.Get(&_CollectionRegistry.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_id bytes32) constant returns(bytes32, bytes8)
func (_CollectionRegistry *CollectionRegistryCallerSession) Get(_id [32]byte) ([32]byte, [8]byte, error) {
	return _CollectionRegistry.Contract.Get(&_CollectionRegistry.CallOpts, _id)
}

// Allow is a paid mutator transaction binding the contract method 0x1fed449f.
//
// Solidity: function allow(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Allow(opts *bind.TransactOpts, _id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "allow", _id, _uid)
}

// Allow is a paid mutator transaction binding the contract method 0x1fed449f.
//
// Solidity: function allow(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistrySession) Allow(_id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Allow(&_CollectionRegistry.TransactOpts, _id, _uid)
}

// Allow is a paid mutator transaction binding the contract method 0x1fed449f.
//
// Solidity: function allow(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Allow(_id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Allow(&_CollectionRegistry.TransactOpts, _id, _uid)
}

// Deny is a paid mutator transaction binding the contract method 0x4fe929c2.
//
// Solidity: function deny(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Deny(opts *bind.TransactOpts, _id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "deny", _id, _uid)
}

// Deny is a paid mutator transaction binding the contract method 0x4fe929c2.
//
// Solidity: function deny(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistrySession) Deny(_id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Deny(&_CollectionRegistry.TransactOpts, _id, _uid)
}

// Deny is a paid mutator transaction binding the contract method 0x4fe929c2.
//
// Solidity: function deny(_id bytes32, _uid bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Deny(_id [32]byte, _uid [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Deny(&_CollectionRegistry.TransactOpts, _id, _uid)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Register(opts *bind.TransactOpts, _appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "register", _appId, _schemaId, _ratio)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistrySession) Register(_appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Register(&_CollectionRegistry.TransactOpts, _appId, _schemaId, _ratio)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Register(_appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Register(&_CollectionRegistry.TransactOpts, _appId, _schemaId, _ratio)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_id bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Unregister(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "unregister", _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_id bytes32) returns()
func (_CollectionRegistry *CollectionRegistrySession) Unregister(_id [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Unregister(&_CollectionRegistry.TransactOpts, _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x1a0919dc.
//
// Solidity: function unregister(_id bytes32) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Unregister(_id [32]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Unregister(&_CollectionRegistry.TransactOpts, _id)
}

// CollectionRegistryAllowedIterator is returned from FilterAllowed and is used to iterate over the raw logs and unpacked data for Allowed events raised by the CollectionRegistry contract.
type CollectionRegistryAllowedIterator struct {
	Event *CollectionRegistryAllowed // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryAllowed)
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
		it.Event = new(CollectionRegistryAllowed)
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
func (it *CollectionRegistryAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryAllowed represents a Allowed event raised by the CollectionRegistry contract.
type CollectionRegistryAllowed struct {
	CollectionId [32]byte
	Uid          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowed is a free log retrieval operation binding the contract event 0xa22515132971f50f788da1869934dff79436ef9486db69a7a59731a5fb61689d.
//
// Solidity: e Allowed(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterAllowed(opts *bind.FilterOpts, _collectionId [][32]byte, _uid [][32]byte) (*CollectionRegistryAllowedIterator, error) {

	var _collectionIdRule []interface{}
	for _, _collectionIdItem := range _collectionId {
		_collectionIdRule = append(_collectionIdRule, _collectionIdItem)
	}
	var _uidRule []interface{}
	for _, _uidItem := range _uid {
		_uidRule = append(_uidRule, _uidItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Allowed", _collectionIdRule, _uidRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryAllowedIterator{contract: _CollectionRegistry.contract, event: "Allowed", logs: logs, sub: sub}, nil
}

// FilterAllowed parses the event from given transaction receipt.
//
// Solidity: e Allowed(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseAllowedFromReceipt(receipt *types.Receipt) (*CollectionRegistryAllowed, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xa22515132971f50f788da1869934dff79436ef9486db69a7a59731a5fb61689d") {
			event := new(CollectionRegistryAllowed)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Allowed", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Allowed event not found")
}

// WatchAllowed is a free log subscription operation binding the contract event 0xa22515132971f50f788da1869934dff79436ef9486db69a7a59731a5fb61689d.
//
// Solidity: e Allowed(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchAllowed(opts *bind.WatchOpts, sink chan<- *CollectionRegistryAllowed, _collectionId [][32]byte, _uid [][32]byte) (event.Subscription, error) {

	var _collectionIdRule []interface{}
	for _, _collectionIdItem := range _collectionId {
		_collectionIdRule = append(_collectionIdRule, _collectionIdItem)
	}
	var _uidRule []interface{}
	for _, _uidItem := range _uid {
		_uidRule = append(_uidRule, _uidItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Allowed", _collectionIdRule, _uidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryAllowed)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Allowed", log); err != nil {
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

// CollectionRegistryDeniedIterator is returned from FilterDenied and is used to iterate over the raw logs and unpacked data for Denied events raised by the CollectionRegistry contract.
type CollectionRegistryDeniedIterator struct {
	Event *CollectionRegistryDenied // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryDenied)
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
		it.Event = new(CollectionRegistryDenied)
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
func (it *CollectionRegistryDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryDenied represents a Denied event raised by the CollectionRegistry contract.
type CollectionRegistryDenied struct {
	CollectionId [32]byte
	Uid          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDenied is a free log retrieval operation binding the contract event 0x4d28190f1b112cd85f6380723c76d76de35cb4a80b5ab017d01f320b25db009e.
//
// Solidity: e Denied(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterDenied(opts *bind.FilterOpts, _collectionId [][32]byte, _uid [][32]byte) (*CollectionRegistryDeniedIterator, error) {

	var _collectionIdRule []interface{}
	for _, _collectionIdItem := range _collectionId {
		_collectionIdRule = append(_collectionIdRule, _collectionIdItem)
	}
	var _uidRule []interface{}
	for _, _uidItem := range _uid {
		_uidRule = append(_uidRule, _uidItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Denied", _collectionIdRule, _uidRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryDeniedIterator{contract: _CollectionRegistry.contract, event: "Denied", logs: logs, sub: sub}, nil
}

// FilterDenied parses the event from given transaction receipt.
//
// Solidity: e Denied(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseDeniedFromReceipt(receipt *types.Receipt) (*CollectionRegistryDenied, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x4d28190f1b112cd85f6380723c76d76de35cb4a80b5ab017d01f320b25db009e") {
			event := new(CollectionRegistryDenied)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Denied", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Denied event not found")
}

// WatchDenied is a free log subscription operation binding the contract event 0x4d28190f1b112cd85f6380723c76d76de35cb4a80b5ab017d01f320b25db009e.
//
// Solidity: e Denied(_collectionId indexed bytes32, _uid indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchDenied(opts *bind.WatchOpts, sink chan<- *CollectionRegistryDenied, _collectionId [][32]byte, _uid [][32]byte) (event.Subscription, error) {

	var _collectionIdRule []interface{}
	for _, _collectionIdItem := range _collectionId {
		_collectionIdRule = append(_collectionIdRule, _collectionIdItem)
	}
	var _uidRule []interface{}
	for _, _uidItem := range _uid {
		_uidRule = append(_uidRule, _uidItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Denied", _collectionIdRule, _uidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryDenied)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Denied", log); err != nil {
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

// CollectionRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the CollectionRegistry contract.
type CollectionRegistryRegisteredIterator struct {
	Event *CollectionRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryRegistered)
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
		it.Event = new(CollectionRegistryRegistered)
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
func (it *CollectionRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryRegistered represents a Registered event raised by the CollectionRegistry contract.
type CollectionRegistryRegistered struct {
	ColectionId [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x10906fae603eebfac53ddc0f103bee8a044dd7643c425c7a90f921dfa15ef62c.
//
// Solidity: e Registered(_colectionId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, _colectionId [][32]byte) (*CollectionRegistryRegisteredIterator, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Registered", _colectionIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryRegisteredIterator{contract: _CollectionRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// FilterRegistered parses the event from given transaction receipt.
//
// Solidity: e Registered(_colectionId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseRegisteredFromReceipt(receipt *types.Receipt) (*CollectionRegistryRegistered, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x10906fae603eebfac53ddc0f103bee8a044dd7643c425c7a90f921dfa15ef62c") {
			event := new(CollectionRegistryRegistered)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Registered event not found")
}

// WatchRegistered is a free log subscription operation binding the contract event 0x10906fae603eebfac53ddc0f103bee8a044dd7643c425c7a90f921dfa15ef62c.
//
// Solidity: e Registered(_colectionId indexed bytes32)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *CollectionRegistryRegistered, _colectionId [][32]byte) (event.Subscription, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Registered", _colectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryRegistered)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// CollectionRegistryUnregisteredIterator is returned from FilterUnregistered and is used to iterate over the raw logs and unpacked data for Unregistered events raised by the CollectionRegistry contract.
type CollectionRegistryUnregisteredIterator struct {
	Event *CollectionRegistryUnregistered // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryUnregistered)
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
		it.Event = new(CollectionRegistryUnregistered)
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
func (it *CollectionRegistryUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryUnregistered represents a Unregistered event raised by the CollectionRegistry contract.
type CollectionRegistryUnregistered struct {
	ColectionId [32]byte
	AppId       [32]byte
	SchemaId    [8]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnregistered is a free log retrieval operation binding the contract event 0x0231d91ceaa0291166e678222375c22a49172b2641f9003dcdebd63e1160cc70.
//
// Solidity: e Unregistered(_colectionId indexed bytes32, _appId indexed bytes32, _schemaId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterUnregistered(opts *bind.FilterOpts, _colectionId [][32]byte, _appId [][32]byte, _schemaId [][8]byte) (*CollectionRegistryUnregisteredIterator, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}
	var _appIdRule []interface{}
	for _, _appIdItem := range _appId {
		_appIdRule = append(_appIdRule, _appIdItem)
	}
	var _schemaIdRule []interface{}
	for _, _schemaIdItem := range _schemaId {
		_schemaIdRule = append(_schemaIdRule, _schemaIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Unregistered", _colectionIdRule, _appIdRule, _schemaIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryUnregisteredIterator{contract: _CollectionRegistry.contract, event: "Unregistered", logs: logs, sub: sub}, nil
}

// FilterUnregistered parses the event from given transaction receipt.
//
// Solidity: e Unregistered(_colectionId indexed bytes32, _appId indexed bytes32, _schemaId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseUnregisteredFromReceipt(receipt *types.Receipt) (*CollectionRegistryUnregistered, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x0231d91ceaa0291166e678222375c22a49172b2641f9003dcdebd63e1160cc70") {
			event := new(CollectionRegistryUnregistered)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Unregistered", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Unregistered event not found")
}

// WatchUnregistered is a free log subscription operation binding the contract event 0x0231d91ceaa0291166e678222375c22a49172b2641f9003dcdebd63e1160cc70.
//
// Solidity: e Unregistered(_colectionId indexed bytes32, _appId indexed bytes32, _schemaId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchUnregistered(opts *bind.WatchOpts, sink chan<- *CollectionRegistryUnregistered, _colectionId [][32]byte, _appId [][32]byte, _schemaId [][8]byte) (event.Subscription, error) {

	var _colectionIdRule []interface{}
	for _, _colectionIdItem := range _colectionId {
		_colectionIdRule = append(_colectionIdRule, _colectionIdItem)
	}
	var _appIdRule []interface{}
	for _, _appIdItem := range _appId {
		_appIdRule = append(_appIdRule, _appIdItem)
	}
	var _schemaIdRule []interface{}
	for _, _schemaIdItem := range _schemaId {
		_schemaIdRule = append(_schemaIdRule, _schemaIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Unregistered", _colectionIdRule, _appIdRule, _schemaIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryUnregistered)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Unregistered", log); err != nil {
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
