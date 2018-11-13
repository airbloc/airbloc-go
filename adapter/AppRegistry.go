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

// AppRegistryABI is the input ABI used to generate the binding from.
const AppRegistryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"apps\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"appId\",\"type\":\"bytes8\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes8\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"newOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes8\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"checkOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes8\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes8\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AppRegistryBin is the compiled bytecode used for deploying new contracts.
const AppRegistryBin = `0x6080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663260a818e811461009d578063398bc4e8146100c1578063672b7beb146100f7578063715018a6146101255780638da5cb5b1461013a578063a856fe781461016b578063e66e46c014610199578063f2c298be1461024c578063f2fde38b146102a5575b600080fd5b3480156100a957600080fd5b506100bf600160c060020a0319600435166102c6565b005b3480156100cd57600080fd5b506100e3600160c060020a031960043516610451565b604080519115158252519081900360200190f35b34801561010357600080fd5b506100e3600160c060020a031960043516600160a060020a0360243516610472565b34801561013157600080fd5b506100bf61049c565b34801561014657600080fd5b5061014f610508565b60408051600160a060020a039092168252519081900360200190f35b34801561017757600080fd5b506100bf600160c060020a031960043516600160a060020a0360243516610517565b3480156101a557600080fd5b506101bb600160c060020a0319600435166105ff565b604051808060200183600160a060020a0316600160a060020a03168152602001828103825284818151815260200191508051906020019080838360005b838110156102105781810151838201526020016101f8565b50505050905090810190601f16801561023d5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561025857600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100bf9436949293602493928401919081908401838280828437509497506106aa9650505050505050565b3480156102b157600080fd5b506100bf600160a060020a03600435166108b3565b60006102d182610451565b151561033e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f41707020646f6573206e6f742065786973742e00000000000000000000000000604482015290519081900360640190fd5b600160008377ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160405180828054600181600116156101000203166002900480156103e55780601f106103c35761010080835404028352918201916103e5565b820191906000526020600020905b8154815290600101906020018083116103d1575b50506040805191829003909120600081815260026020908152838220805460ff19169055600160c060020a031988168252600190529182209094509250905061042e8282610a1a565b50600101805473ffffffffffffffffffffffffffffffffffffffff191690555050565b60008061045d836108d6565b60010154600160a060020a0316141592915050565b600081600160a060020a0316610487846108d6565b60010154600160a060020a0316149392505050565b600054600160a060020a031633146104b357600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b6105218233610472565b15156105b457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f6f6e6c79206f776e65722063616e207472616e73666572206f776e657273686960448201527f7000000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160c060020a031991909116600090815260016020819052604090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b60016020818152600092835260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529283918301828280156106915780601f1061066657610100808354040283529160200191610691565b820191906000526020600020905b81548152906001019060200180831161067457829003601f168201915b50505060019093015491925050600160a060020a031682565b600080826040518082805190602001908083835b602083106106dd5780518252601f1990920191602091820191016106be565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008181526002909252929020549195505060ff1615915061078a905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f417070206e616d6520616c7265616479206578697374732e0000000000000000604482015290519081900360640190fd5b6000828152600260205260409020805460ff191660011790556107ad82336108f1565b600160c060020a03198116600090815260016020908152604090912085519293506107dc929091860190610a5e565b50600160c060020a03198116600090815260016020818152604092839020909101805473ffffffffffffffffffffffffffffffffffffffff1916331790559051845185928291908401908083835b602083106108495780518252601f19909201916020918201910161082a565b51815160209384036101000a600019018019909216911617905260408051929094018290038220600160c060020a03198816835293519395507f81fd67a36836c34ada265e9e349edf3799560c1e24a2149ba1ff6c427179f59194509083900301919050a2505050565b600054600160a060020a031633146108ca57600080fd5b6108d38161099d565b50565b600160c060020a031916600090815260016020526040902090565b604080516c01000000000000000000000000600160a060020a0384160260208083019190915243603483015260548083018690528351808403909101815260749092019283905281516000938392909182918401908083835b602083106109695780518252601f19909201916020918201910161094a565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120979650505050505050565b600160a060020a03811615156109b257600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b50805460018160011615610100020316600290046000825580601f10610a4057506108d3565b601f0160209004906000526020600020908101906108d39190610adc565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a9f57805160ff1916838001178555610acc565b82800160010185558215610acc579182015b82811115610acc578251825591602001919060010190610ab1565b50610ad8929150610adc565b5090565b610af691905b80821115610ad85760008155600101610ae2565b905600a165627a7a723058200b2f14f88bc841cf2272af180127db4141217713531755526c3592f3d0a43fb90029`

// DeployAppRegistry deploys a new Ethereum contract, binding an instance of AppRegistry to it.
func DeployAppRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AppRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(AppRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AppRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AppRegistry{AppRegistryCaller: AppRegistryCaller{contract: contract}, AppRegistryTransactor: AppRegistryTransactor{contract: contract}, AppRegistryFilterer: AppRegistryFilterer{contract: contract}}, nil
}

// AppRegistry is an auto generated Go binding around an Ethereum contract.
type AppRegistry struct {
	AppRegistryCaller     // Read-only binding to the contract
	AppRegistryTransactor // Write-only binding to the contract
	AppRegistryFilterer   // Log filterer for contract events
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

// AppRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppRegistryCallerRaw struct {
	Contract *AppRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AppRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppRegistryTransactorRaw struct {
	Contract *AppRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppRegistry creates a new instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistry(address common.Address, backend bind.ContractBackend) (*AppRegistry, error) {
	contract, err := bindAppRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppRegistry{AppRegistryCaller: AppRegistryCaller{contract: contract}, AppRegistryTransactor: AppRegistryTransactor{contract: contract}, AppRegistryFilterer: AppRegistryFilterer{contract: contract}}, nil
}

// NewAppRegistryCaller creates a new read-only instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryCaller(address common.Address, caller bind.ContractCaller) (*AppRegistryCaller, error) {
	contract, err := bindAppRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppRegistryCaller{contract: contract}, nil
}

// NewAppRegistryTransactor creates a new write-only instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AppRegistryTransactor, error) {
	contract, err := bindAppRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppRegistryTransactor{contract: contract}, nil
}

// NewAppRegistryFilterer creates a new log filterer instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AppRegistryFilterer, error) {
	contract, err := bindAppRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppRegistryFilterer{contract: contract}, nil
}

// bindAppRegistry binds a generic wrapper to an already deployed contract.
func bindAppRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
func (_AppRegistry *AppRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppRegistry *AppRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AppRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppRegistry *AppRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppRegistry.Contract.contract.Transact(opts, method, params...)
}

// Apps is a free data retrieval call binding the contract method 0xe66e46c0.
//
// Solidity: function apps( bytes8) constant returns(name string, owner address)
func (_AppRegistry *AppRegistryCaller) Apps(opts *bind.CallOpts, arg0 [8]byte) (struct {
	Name  string
	Owner common.Address
}, error) {
	ret := new(struct {
		Name  string
		Owner common.Address
	})
	out := ret
	err := _AppRegistry.contract.Call(opts, out, "apps", arg0)
	return *ret, err
}

// Apps is a free data retrieval call binding the contract method 0xe66e46c0.
//
// Solidity: function apps( bytes8) constant returns(name string, owner address)
func (_AppRegistry *AppRegistrySession) Apps(arg0 [8]byte) (struct {
	Name  string
	Owner common.Address
}, error) {
	return _AppRegistry.Contract.Apps(&_AppRegistry.CallOpts, arg0)
}

// Apps is a free data retrieval call binding the contract method 0xe66e46c0.
//
// Solidity: function apps( bytes8) constant returns(name string, owner address)
func (_AppRegistry *AppRegistryCallerSession) Apps(arg0 [8]byte) (struct {
	Name  string
	Owner common.Address
}, error) {
	return _AppRegistry.Contract.Apps(&_AppRegistry.CallOpts, arg0)
}

// Check is a free data retrieval call binding the contract method 0x398bc4e8.
//
// Solidity: function check(_appId bytes8) constant returns(bool)
func (_AppRegistry *AppRegistryCaller) Check(opts *bind.CallOpts, _appId [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "check", _appId)
	return *ret0, err
}

// Check is a free data retrieval call binding the contract method 0x398bc4e8.
//
// Solidity: function check(_appId bytes8) constant returns(bool)
func (_AppRegistry *AppRegistrySession) Check(_appId [8]byte) (bool, error) {
	return _AppRegistry.Contract.Check(&_AppRegistry.CallOpts, _appId)
}

// Check is a free data retrieval call binding the contract method 0x398bc4e8.
//
// Solidity: function check(_appId bytes8) constant returns(bool)
func (_AppRegistry *AppRegistryCallerSession) Check(_appId [8]byte) (bool, error) {
	return _AppRegistry.Contract.Check(&_AppRegistry.CallOpts, _appId)
}

// CheckOwner is a free data retrieval call binding the contract method 0x672b7beb.
//
// Solidity: function checkOwner(_appId bytes8, _owner address) constant returns(bool)
func (_AppRegistry *AppRegistryCaller) CheckOwner(opts *bind.CallOpts, _appId [8]byte, _owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "checkOwner", _appId, _owner)
	return *ret0, err
}

// CheckOwner is a free data retrieval call binding the contract method 0x672b7beb.
//
// Solidity: function checkOwner(_appId bytes8, _owner address) constant returns(bool)
func (_AppRegistry *AppRegistrySession) CheckOwner(_appId [8]byte, _owner common.Address) (bool, error) {
	return _AppRegistry.Contract.CheckOwner(&_AppRegistry.CallOpts, _appId, _owner)
}

// CheckOwner is a free data retrieval call binding the contract method 0x672b7beb.
//
// Solidity: function checkOwner(_appId bytes8, _owner address) constant returns(bool)
func (_AppRegistry *AppRegistryCallerSession) CheckOwner(_appId [8]byte, _owner common.Address) (bool, error) {
	return _AppRegistry.Contract.CheckOwner(&_AppRegistry.CallOpts, _appId, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AppRegistry *AppRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AppRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AppRegistry *AppRegistrySession) Owner() (common.Address, error) {
	return _AppRegistry.Contract.Owner(&_AppRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AppRegistry *AppRegistryCallerSession) Owner() (common.Address, error) {
	return _AppRegistry.Contract.Owner(&_AppRegistry.CallOpts)
}

// NewOwner is a paid mutator transaction binding the contract method 0xa856fe78.
//
// Solidity: function newOwner(_appId bytes8, _newOwner address) returns()
func (_AppRegistry *AppRegistryTransactor) NewOwner(opts *bind.TransactOpts, _appId [8]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "newOwner", _appId, _newOwner)
}

// NewOwner is a paid mutator transaction binding the contract method 0xa856fe78.
//
// Solidity: function newOwner(_appId bytes8, _newOwner address) returns()
func (_AppRegistry *AppRegistrySession) NewOwner(_appId [8]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.NewOwner(&_AppRegistry.TransactOpts, _appId, _newOwner)
}

// NewOwner is a paid mutator transaction binding the contract method 0xa856fe78.
//
// Solidity: function newOwner(_appId bytes8, _newOwner address) returns()
func (_AppRegistry *AppRegistryTransactorSession) NewOwner(_appId [8]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.NewOwner(&_AppRegistry.TransactOpts, _appId, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_name string) returns()
func (_AppRegistry *AppRegistryTransactor) Register(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "register", _name)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_name string) returns()
func (_AppRegistry *AppRegistrySession) Register(_name string) (*types.Transaction, error) {
	return _AppRegistry.Contract.Register(&_AppRegistry.TransactOpts, _name)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_name string) returns()
func (_AppRegistry *AppRegistryTransactorSession) Register(_name string) (*types.Transaction, error) {
	return _AppRegistry.Contract.Register(&_AppRegistry.TransactOpts, _name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppRegistry *AppRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppRegistry *AppRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _AppRegistry.Contract.RenounceOwnership(&_AppRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppRegistry *AppRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AppRegistry.Contract.RenounceOwnership(&_AppRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AppRegistry *AppRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AppRegistry *AppRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.TransferOwnership(&_AppRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AppRegistry *AppRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.TransferOwnership(&_AppRegistry.TransactOpts, _newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_appId bytes8) returns()
func (_AppRegistry *AppRegistryTransactor) Unregister(opts *bind.TransactOpts, _appId [8]byte) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "unregister", _appId)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_appId bytes8) returns()
func (_AppRegistry *AppRegistrySession) Unregister(_appId [8]byte) (*types.Transaction, error) {
	return _AppRegistry.Contract.Unregister(&_AppRegistry.TransactOpts, _appId)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_appId bytes8) returns()
func (_AppRegistry *AppRegistryTransactorSession) Unregister(_appId [8]byte) (*types.Transaction, error) {
	return _AppRegistry.Contract.Unregister(&_AppRegistry.TransactOpts, _appId)
}

// AppRegistryOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the AppRegistry contract.
type AppRegistryOwnershipRenouncedIterator struct {
	Event *AppRegistryOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *AppRegistryOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryOwnershipRenounced)
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
		it.Event = new(AppRegistryOwnershipRenounced)
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
func (it *AppRegistryOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryOwnershipRenounced represents a OwnershipRenounced event raised by the AppRegistry contract.
type AppRegistryOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*AppRegistryOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryOwnershipRenouncedIterator{contract: _AppRegistry.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// FilterOwnershipRenounced parses the event from given transaction receipt.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) ParseOwnershipRenouncedFromReceipt(receipt *types.Receipt) (*AppRegistryOwnershipRenounced, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820") {
			event := new(AppRegistryOwnershipRenounced)
			if err := _AppRegistry.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OwnershipRenounced event not found")
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *AppRegistryOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryOwnershipRenounced)
				if err := _AppRegistry.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// AppRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AppRegistry contract.
type AppRegistryOwnershipTransferredIterator struct {
	Event *AppRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AppRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryOwnershipTransferred)
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
		it.Event = new(AppRegistryOwnershipTransferred)
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
func (it *AppRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the AppRegistry contract.
type AppRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AppRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryOwnershipTransferredIterator{contract: _AppRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// FilterOwnershipTransferred parses the event from given transaction receipt.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) ParseOwnershipTransferredFromReceipt(receipt *types.Receipt) (*AppRegistryOwnershipTransferred, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0") {
			event := new(AppRegistryOwnershipTransferred)
			if err := _AppRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OwnershipTransferred event not found")
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AppRegistry *AppRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AppRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryOwnershipTransferred)
				if err := _AppRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// AppRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the AppRegistry contract.
type AppRegistryRegisteredIterator struct {
	Event *AppRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *AppRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryRegistered)
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
		it.Event = new(AppRegistryRegistered)
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
func (it *AppRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRegistered represents a Registered event raised by the AppRegistry contract.
type AppRegistryRegistered struct {
	Name  common.Hash
	AppId [8]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x81fd67a36836c34ada265e9e349edf3799560c1e24a2149ba1ff6c427179f591.
//
// Solidity: e Registered(name indexed string, appId bytes8)
func (_AppRegistry *AppRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, name []string) (*AppRegistryRegisteredIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Registered", nameRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRegisteredIterator{contract: _AppRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// FilterRegistered parses the event from given transaction receipt.
//
// Solidity: e Registered(name indexed string, appId bytes8)
func (_AppRegistry *AppRegistryFilterer) ParseRegisteredFromReceipt(receipt *types.Receipt) (*AppRegistryRegistered, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x81fd67a36836c34ada265e9e349edf3799560c1e24a2149ba1ff6c427179f591") {
			event := new(AppRegistryRegistered)
			if err := _AppRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Registered event not found")
}

// WatchRegistered is a free log subscription operation binding the contract event 0x81fd67a36836c34ada265e9e349edf3799560c1e24a2149ba1ff6c427179f591.
//
// Solidity: e Registered(name indexed string, appId bytes8)
func (_AppRegistry *AppRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *AppRegistryRegistered, name []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Registered", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryRegistered)
				if err := _AppRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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
