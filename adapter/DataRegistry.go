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

// DataRegistryABI is the input ABI used to generate the binding from.
const DataRegistryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_accounts\",\"type\":\"address\"},{\"name\":\"_collections\",\"type\":\"address\"},{\"name\":\"_smt\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"BundleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"Punished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"name\":\"usersRoot\",\"type\":\"bytes32\"},{\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"registerBundle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"name\":\"bundleIndex\",\"type\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"name\":\"bundleIndex\",\"type\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"isMyDataIncluded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isCollectionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataRegistryBin is the compiled bytecode used for deploying new contracts.
const DataRegistryBin = `0x6080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631ac30935811461008757806345ab73d1146100c5578063715018a61461014d5780638da5cb5b14610162578063bde66f2c14610193578063d8684912146101c9578063f2fde38b146101f7575b600080fd5b34801561009357600080fd5b506100c360048035600160c060020a031916906024803567ffffffffffffffff1691604435918201910135610218565b005b3480156100d157600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610139948235600160c060020a031916946024803567ffffffffffffffff169536959460649492019190819084018382808284375094975061059d9650505050505050565b604080519115158252519081900360200190f35b34801561015957600080fd5b506100c3610824565b34801561016e57600080fd5b50610177610890565b60408051600160a060020a039092168252519081900360200190f35b34801561019f57600080fd5b506100c360048035600160c060020a0319169060248035916044359160643590810191013561089f565b3480156101d557600080fd5b50610139600160c060020a031960043516600160a060020a0360243516610a06565b34801561020357600080fd5b506100c3600160a060020a0360043516610a1a565b600160c060020a031984166000908152600160205260408120805490919081908190600160a060020a03161515610299576040805160e560020a62461bcd02815260206004820152601a60248201527f436f6c6c656374696f6e20646f6573206e6f742065786973742e000000000000604482015290519081900360640190fd5b60028401805467ffffffffffffffff89169081106102b357fe5b906000526020600020906005020192506102d288338560030154610a3d565b1515610374576040805160e560020a62461bcd02815260206004820152604c60248201527f596f752068617665206265656e20616c6c6f77656420746f20636f6c6c65637460448201527f20746865206461746120617420746861742074696d652e20576879206973206960648201527f7420612070726f626c656d3f0000000000000000000000000000000000000000608482015290519081900360a40190fd5b600254604080517fe0b490f70000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163e0b490f7916024808201926020929091908290030181600087803b1580156103da57600080fd5b505af11580156103ee573d6000803e3d6000fd5b505050506040513d602081101561040457600080fd5b50516004805485546040517ff586df6500000000000000000000000000000000000000000000000000000000815260019381018481526024820183905267ffffffffffffffff78010000000000000000000000000000000000000000000000009096049586166044830152608060648301908152608483018c9052959750600160a060020a039093169463f586df65949388928d928d929160a40184848082843782019150509650505050505050602060405180830381600087803b1580156104cc57600080fd5b505af11580156104e0573d6000803e3d6000fd5b505050506040513d60208110156104f657600080fd5b50519050801561054357835460408051600160a060020a039092168252517ff6a2a2bc3297e42d6b873d907a8cd2699857e3d3700babd53b7061a3b4de60949181900360200190a1610593565b6040805160e560020a62461bcd02815260206004820152600c60248201527f50726f6f66206661696c65640000000000000000000000000000000000000000604482015290519081900360640190fd5b5050505050505050565b600160c060020a031983166000908152600160205260408120805482908190600160a060020a0316151561061b576040805160e560020a62461bcd02815260206004820152601a60248201527f436f6c6c656374696f6e20646f6573206e6f742065786973742e000000000000604482015290519081900360640190fd5b600254604080517fe0b490f70000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163e0b490f7916024808201926020929091908290030181600087803b15801561068157600080fd5b505af1158015610695573d6000803e3d6000fd5b505050506040513d60208110156106ab57600080fd5b5051600284018054780100000000000000000000000000000000000000000000000090920493509067ffffffffffffffff88169081106106e757fe5b60009182526020808320600590920290910154600480546040517ff586df6500000000000000000000000000000000000000000000000000000000815260019281018381526024820185905267ffffffffffffffff891660448301526080606483019081528c5160848401528c51959850600160a060020a039093169663f586df6596949589958b958f95909360a49091019291860191908190849084905b8381101561079e578181015183820152602001610786565b50505050905090810190601f1680156107cb5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156107ed57600080fd5b505af1158015610801573d6000803e3d6000fd5b505050506040513d602081101561081757600080fd5b5051979650505050505050565b600054600160a060020a0316331461083b57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600160c060020a03198516600090815260016020526040812080549091908190600160a060020a0316151561091e576040805160e560020a62461bcd02815260206004820152601a60248201527f436f6c6c656374696f6e20646f6573206e6f742065786973742e000000000000604482015290519081900360640190fd5b86825560018201869055610936600283018686610b31565b5050600282810180546001808201808455600093845260209093208554600584029091019081558582015481830155858501805493958794929361098c9385830193929181161561010002600019011604610baf565b50600382810154908201556004918201549101805467ffffffffffffffff191667ffffffffffffffff92831617905560408051918416825251600160c060020a03198b1692507fab7212f2e313639f22d9f8d95bc067b9289814ce97d8136a08e37d239023b1a39181900360200190a25050505050505050565b6000610a13838343610a3d565b9392505050565b600054600160a060020a03163314610a3157600080fd5b610a3a81610ab4565b50565b600160c060020a031983166000908152600160208181526040808420600160a060020a0387168552909201905281205460ff168015610aac5750600160c060020a031984166000908152600160208181526040808420600160a060020a03881685528301909152909120015482115b949350505050565b600160a060020a0381161515610ac957600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b725782800160ff19823516178555610b9f565b82800160010185558215610b9f579182015b82811115610b9f578235825591602001919060010190610b84565b50610bab929150610c24565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610be85780548555610b9f565b82800160010185558215610b9f57600052602060002091601f016020900482015b82811115610b9f578254825591600101919060010190610c09565b610c3e91905b80821115610bab5760008155600101610c2a565b905600a165627a7a7230582084ed27242cd538d63c11ab45783269e815d4a0e869f3fd1de2e405c292c7ce690029`

// DeployDataRegistry deploys a new Ethereum contract, binding an instance of DataRegistry to it.
func DeployDataRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _accounts common.Address, _collections common.Address, _smt common.Address) (common.Address, *types.Transaction, *DataRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(DataRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataRegistryBin), backend, _accounts, _collections, _smt)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataRegistry{DataRegistryCaller: DataRegistryCaller{contract: contract}, DataRegistryTransactor: DataRegistryTransactor{contract: contract}, DataRegistryFilterer: DataRegistryFilterer{contract: contract}}, nil
}

// DataRegistry is an auto generated Go binding around an Ethereum contract.
type DataRegistry struct {
	DataRegistryCaller     // Read-only binding to the contract
	DataRegistryTransactor // Write-only binding to the contract
	DataRegistryFilterer   // Log filterer for contract events
}

// DataRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataRegistrySession struct {
	Contract     *DataRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataRegistryCallerSession struct {
	Contract *DataRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DataRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataRegistryTransactorSession struct {
	Contract     *DataRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DataRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataRegistryRaw struct {
	Contract *DataRegistry // Generic contract binding to access the raw methods on
}

// DataRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataRegistryCallerRaw struct {
	Contract *DataRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DataRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataRegistryTransactorRaw struct {
	Contract *DataRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataRegistry creates a new instance of DataRegistry, bound to a specific deployed contract.
func NewDataRegistry(address common.Address, backend bind.ContractBackend) (*DataRegistry, error) {
	contract, err := bindDataRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataRegistry{DataRegistryCaller: DataRegistryCaller{contract: contract}, DataRegistryTransactor: DataRegistryTransactor{contract: contract}, DataRegistryFilterer: DataRegistryFilterer{contract: contract}}, nil
}

// NewDataRegistryCaller creates a new read-only instance of DataRegistry, bound to a specific deployed contract.
func NewDataRegistryCaller(address common.Address, caller bind.ContractCaller) (*DataRegistryCaller, error) {
	contract, err := bindDataRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataRegistryCaller{contract: contract}, nil
}

// NewDataRegistryTransactor creates a new write-only instance of DataRegistry, bound to a specific deployed contract.
func NewDataRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DataRegistryTransactor, error) {
	contract, err := bindDataRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataRegistryTransactor{contract: contract}, nil
}

// NewDataRegistryFilterer creates a new log filterer instance of DataRegistry, bound to a specific deployed contract.
func NewDataRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DataRegistryFilterer, error) {
	contract, err := bindDataRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataRegistryFilterer{contract: contract}, nil
}

// bindDataRegistry binds a generic wrapper to an already deployed contract.
func bindDataRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataRegistry *DataRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataRegistry.Contract.DataRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataRegistry *DataRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataRegistry.Contract.DataRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataRegistry *DataRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataRegistry.Contract.DataRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataRegistry *DataRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataRegistry *DataRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataRegistry *DataRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xd8684912.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user address) constant returns(bool)
func (_DataRegistry *DataRegistryCaller) IsCollectionAllowed(opts *bind.CallOpts, collectionId [8]byte, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DataRegistry.contract.Call(opts, out, "isCollectionAllowed", collectionId, user)
	return *ret0, err
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xd8684912.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user address) constant returns(bool)
func (_DataRegistry *DataRegistrySession) IsCollectionAllowed(collectionId [8]byte, user common.Address) (bool, error) {
	return _DataRegistry.Contract.IsCollectionAllowed(&_DataRegistry.CallOpts, collectionId, user)
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xd8684912.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user address) constant returns(bool)
func (_DataRegistry *DataRegistryCallerSession) IsCollectionAllowed(collectionId [8]byte, user common.Address) (bool, error) {
	return _DataRegistry.Contract.IsCollectionAllowed(&_DataRegistry.CallOpts, collectionId, user)
}

// IsMyDataIncluded is a free data retrieval call binding the contract method 0x45ab73d1.
//
// Solidity: function isMyDataIncluded(collectionId bytes8, bundleIndex uint64, proof bytes) constant returns(bool)
func (_DataRegistry *DataRegistryCaller) IsMyDataIncluded(opts *bind.CallOpts, collectionId [8]byte, bundleIndex uint64, proof []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DataRegistry.contract.Call(opts, out, "isMyDataIncluded", collectionId, bundleIndex, proof)
	return *ret0, err
}

// IsMyDataIncluded is a free data retrieval call binding the contract method 0x45ab73d1.
//
// Solidity: function isMyDataIncluded(collectionId bytes8, bundleIndex uint64, proof bytes) constant returns(bool)
func (_DataRegistry *DataRegistrySession) IsMyDataIncluded(collectionId [8]byte, bundleIndex uint64, proof []byte) (bool, error) {
	return _DataRegistry.Contract.IsMyDataIncluded(&_DataRegistry.CallOpts, collectionId, bundleIndex, proof)
}

// IsMyDataIncluded is a free data retrieval call binding the contract method 0x45ab73d1.
//
// Solidity: function isMyDataIncluded(collectionId bytes8, bundleIndex uint64, proof bytes) constant returns(bool)
func (_DataRegistry *DataRegistryCallerSession) IsMyDataIncluded(collectionId [8]byte, bundleIndex uint64, proof []byte) (bool, error) {
	return _DataRegistry.Contract.IsMyDataIncluded(&_DataRegistry.CallOpts, collectionId, bundleIndex, proof)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DataRegistry *DataRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DataRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DataRegistry *DataRegistrySession) Owner() (common.Address, error) {
	return _DataRegistry.Contract.Owner(&_DataRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DataRegistry *DataRegistryCallerSession) Owner() (common.Address, error) {
	return _DataRegistry.Contract.Owner(&_DataRegistry.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0x1ac30935.
//
// Solidity: function challenge(collectionId bytes8, bundleIndex uint64, proof bytes) returns()
func (_DataRegistry *DataRegistryTransactor) Challenge(opts *bind.TransactOpts, collectionId [8]byte, bundleIndex uint64, proof []byte) (*types.Transaction, error) {
	return _DataRegistry.contract.Transact(opts, "challenge", collectionId, bundleIndex, proof)
}

// Challenge is a paid mutator transaction binding the contract method 0x1ac30935.
//
// Solidity: function challenge(collectionId bytes8, bundleIndex uint64, proof bytes) returns()
func (_DataRegistry *DataRegistrySession) Challenge(collectionId [8]byte, bundleIndex uint64, proof []byte) (*types.Transaction, error) {
	return _DataRegistry.Contract.Challenge(&_DataRegistry.TransactOpts, collectionId, bundleIndex, proof)
}

// Challenge is a paid mutator transaction binding the contract method 0x1ac30935.
//
// Solidity: function challenge(collectionId bytes8, bundleIndex uint64, proof bytes) returns()
func (_DataRegistry *DataRegistryTransactorSession) Challenge(collectionId [8]byte, bundleIndex uint64, proof []byte) (*types.Transaction, error) {
	return _DataRegistry.Contract.Challenge(&_DataRegistry.TransactOpts, collectionId, bundleIndex, proof)
}

// RegisterBundle is a paid mutator transaction binding the contract method 0xbde66f2c.
//
// Solidity: function registerBundle(collectionId bytes8, usersRoot bytes32, dataHash bytes32, uri string) returns()
func (_DataRegistry *DataRegistryTransactor) RegisterBundle(opts *bind.TransactOpts, collectionId [8]byte, usersRoot [32]byte, dataHash [32]byte, uri string) (*types.Transaction, error) {
	return _DataRegistry.contract.Transact(opts, "registerBundle", collectionId, usersRoot, dataHash, uri)
}

// RegisterBundle is a paid mutator transaction binding the contract method 0xbde66f2c.
//
// Solidity: function registerBundle(collectionId bytes8, usersRoot bytes32, dataHash bytes32, uri string) returns()
func (_DataRegistry *DataRegistrySession) RegisterBundle(collectionId [8]byte, usersRoot [32]byte, dataHash [32]byte, uri string) (*types.Transaction, error) {
	return _DataRegistry.Contract.RegisterBundle(&_DataRegistry.TransactOpts, collectionId, usersRoot, dataHash, uri)
}

// RegisterBundle is a paid mutator transaction binding the contract method 0xbde66f2c.
//
// Solidity: function registerBundle(collectionId bytes8, usersRoot bytes32, dataHash bytes32, uri string) returns()
func (_DataRegistry *DataRegistryTransactorSession) RegisterBundle(collectionId [8]byte, usersRoot [32]byte, dataHash [32]byte, uri string) (*types.Transaction, error) {
	return _DataRegistry.Contract.RegisterBundle(&_DataRegistry.TransactOpts, collectionId, usersRoot, dataHash, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataRegistry *DataRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataRegistry *DataRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _DataRegistry.Contract.RenounceOwnership(&_DataRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataRegistry *DataRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DataRegistry.Contract.RenounceOwnership(&_DataRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DataRegistry *DataRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DataRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DataRegistry *DataRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DataRegistry.Contract.TransferOwnership(&_DataRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DataRegistry *DataRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DataRegistry.Contract.TransferOwnership(&_DataRegistry.TransactOpts, _newOwner)
}

// DataRegistryBundleRegisteredIterator is returned from FilterBundleRegistered and is used to iterate over the raw logs and unpacked data for BundleRegistered events raised by the DataRegistry contract.
type DataRegistryBundleRegisteredIterator struct {
	Event *DataRegistryBundleRegistered // Event containing the contract specifics and raw log

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
func (it *DataRegistryBundleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRegistryBundleRegistered)
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
		it.Event = new(DataRegistryBundleRegistered)
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
func (it *DataRegistryBundleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRegistryBundleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRegistryBundleRegistered represents a BundleRegistered event raised by the DataRegistry contract.
type DataRegistryBundleRegistered struct {
	CollectionId [8]byte
	Index        uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBundleRegistered is a free log retrieval operation binding the contract event 0xab7212f2e313639f22d9f8d95bc067b9289814ce97d8136a08e37d239023b1a3.
//
// Solidity: e BundleRegistered(collectionId indexed bytes8, index uint64)
func (_DataRegistry *DataRegistryFilterer) FilterBundleRegistered(opts *bind.FilterOpts, collectionId [][8]byte) (*DataRegistryBundleRegisteredIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _DataRegistry.contract.FilterLogs(opts, "BundleRegistered", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return &DataRegistryBundleRegisteredIterator{contract: _DataRegistry.contract, event: "BundleRegistered", logs: logs, sub: sub}, nil
}

// WatchBundleRegistered is a free log subscription operation binding the contract event 0xab7212f2e313639f22d9f8d95bc067b9289814ce97d8136a08e37d239023b1a3.
//
// Solidity: e BundleRegistered(collectionId indexed bytes8, index uint64)
func (_DataRegistry *DataRegistryFilterer) WatchBundleRegistered(opts *bind.WatchOpts, sink chan<- *DataRegistryBundleRegistered, collectionId [][8]byte) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _DataRegistry.contract.WatchLogs(opts, "BundleRegistered", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRegistryBundleRegistered)
				if err := _DataRegistry.contract.UnpackLog(event, "BundleRegistered", log); err != nil {
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

// DataRegistryOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the DataRegistry contract.
type DataRegistryOwnershipRenouncedIterator struct {
	Event *DataRegistryOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DataRegistryOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRegistryOwnershipRenounced)
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
		it.Event = new(DataRegistryOwnershipRenounced)
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
func (it *DataRegistryOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRegistryOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRegistryOwnershipRenounced represents a OwnershipRenounced event raised by the DataRegistry contract.
type DataRegistryOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DataRegistry *DataRegistryFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DataRegistryOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DataRegistry.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DataRegistryOwnershipRenouncedIterator{contract: _DataRegistry.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DataRegistry *DataRegistryFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DataRegistryOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DataRegistry.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRegistryOwnershipRenounced)
				if err := _DataRegistry.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DataRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DataRegistry contract.
type DataRegistryOwnershipTransferredIterator struct {
	Event *DataRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DataRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRegistryOwnershipTransferred)
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
		it.Event = new(DataRegistryOwnershipTransferred)
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
func (it *DataRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the DataRegistry contract.
type DataRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DataRegistry *DataRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DataRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DataRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DataRegistryOwnershipTransferredIterator{contract: _DataRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DataRegistry *DataRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DataRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DataRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRegistryOwnershipTransferred)
				if err := _DataRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DataRegistryPunishedIterator is returned from FilterPunished and is used to iterate over the raw logs and unpacked data for Punished events raised by the DataRegistry contract.
type DataRegistryPunishedIterator struct {
	Event *DataRegistryPunished // Event containing the contract specifics and raw log

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
func (it *DataRegistryPunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRegistryPunished)
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
		it.Event = new(DataRegistryPunished)
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
func (it *DataRegistryPunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRegistryPunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRegistryPunished represents a Punished event raised by the DataRegistry contract.
type DataRegistryPunished struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPunished is a free log retrieval operation binding the contract event 0xf6a2a2bc3297e42d6b873d907a8cd2699857e3d3700babd53b7061a3b4de6094.
//
// Solidity: e Punished(provider address)
func (_DataRegistry *DataRegistryFilterer) FilterPunished(opts *bind.FilterOpts) (*DataRegistryPunishedIterator, error) {

	logs, sub, err := _DataRegistry.contract.FilterLogs(opts, "Punished")
	if err != nil {
		return nil, err
	}
	return &DataRegistryPunishedIterator{contract: _DataRegistry.contract, event: "Punished", logs: logs, sub: sub}, nil
}

// WatchPunished is a free log subscription operation binding the contract event 0xf6a2a2bc3297e42d6b873d907a8cd2699857e3d3700babd53b7061a3b4de6094.
//
// Solidity: e Punished(provider address)
func (_DataRegistry *DataRegistryFilterer) WatchPunished(opts *bind.WatchOpts, sink chan<- *DataRegistryPunished) (event.Subscription, error) {

	logs, sub, err := _DataRegistry.contract.WatchLogs(opts, "Punished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRegistryPunished)
				if err := _DataRegistry.contract.UnpackLog(event, "Punished", log); err != nil {
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
