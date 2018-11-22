// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"
	"strings"

	"github.com/airbloc/airbloc-go/blockchain"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
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
	_ = ethCommon.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SparseMerkleTreeABI is the input ABI used to generate the binding from.
const SparseMerkleTreeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"defaultHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"leafID\",\"type\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"checkMembership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"getRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SparseMerkleTree is an auto generated Go binding around an Ethereum contract.
type SparseMerkleTree struct {
	Address                    ethCommon.Address
	SparseMerkleTreeCaller     // Read-only binding to the contract
	SparseMerkleTreeTransactor // Write-only binding to the contract
	SparseMerkleTreeFilterer   // Log filterer for contract events
}

// SparseMerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SparseMerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparseMerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SparseMerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparseMerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SparseMerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparseMerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SparseMerkleTreeSession struct {
	Contract     *SparseMerkleTree // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SparseMerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SparseMerkleTreeCallerSession struct {
	Contract *SparseMerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SparseMerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SparseMerkleTreeTransactorSession struct {
	Contract     *SparseMerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SparseMerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SparseMerkleTreeRaw struct {
	Contract *SparseMerkleTree // Generic contract binding to access the raw methods on
}

// SparseMerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SparseMerkleTreeCallerRaw struct {
	Contract *SparseMerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// SparseMerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SparseMerkleTreeTransactorRaw struct {
	Contract *SparseMerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

func init() {
	blockchain.ContractList["SparseMerkleTree"] = (&SparseMerkleTree{}).new
}

// NewSparseMerkleTree creates a new instance of SparseMerkleTree, bound to a specific deployed contract.
func NewSparseMerkleTree(address ethCommon.Address, backend bind.ContractBackend) (*SparseMerkleTree, error) {
	contract, err := bindSparseMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SparseMerkleTree{
		Address:                    address,
		SparseMerkleTreeCaller:     SparseMerkleTreeCaller{contract: contract},
		SparseMerkleTreeTransactor: SparseMerkleTreeTransactor{contract: contract},
		SparseMerkleTreeFilterer:   SparseMerkleTreeFilterer{contract: contract},
	}, nil
}

// NewSparseMerkleTreeCaller creates a new read-only instance of SparseMerkleTree, bound to a specific deployed contract.
func NewSparseMerkleTreeCaller(address ethCommon.Address, caller bind.ContractCaller) (*SparseMerkleTreeCaller, error) {
	contract, err := bindSparseMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SparseMerkleTreeCaller{contract: contract}, nil
}

// NewSparseMerkleTreeTransactor creates a new write-only instance of SparseMerkleTree, bound to a specific deployed contract.
func NewSparseMerkleTreeTransactor(address ethCommon.Address, transactor bind.ContractTransactor) (*SparseMerkleTreeTransactor, error) {
	contract, err := bindSparseMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SparseMerkleTreeTransactor{contract: contract}, nil
}

// NewSparseMerkleTreeFilterer creates a new log filterer instance of SparseMerkleTree, bound to a specific deployed contract.
func NewSparseMerkleTreeFilterer(address ethCommon.Address, filterer bind.ContractFilterer) (*SparseMerkleTreeFilterer, error) {
	contract, err := bindSparseMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SparseMerkleTreeFilterer{contract: contract}, nil
}

// bindSparseMerkleTree binds a generic wrapper to an already deployed contract.
func bindSparseMerkleTree(address ethCommon.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SparseMerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_SparseMerkleTree *SparseMerkleTree) new(address ethCommon.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewSparseMerkleTree(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SparseMerkleTree *SparseMerkleTreeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SparseMerkleTree.Contract.SparseMerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SparseMerkleTree *SparseMerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SparseMerkleTree.Contract.SparseMerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SparseMerkleTree *SparseMerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SparseMerkleTree.Contract.SparseMerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SparseMerkleTree *SparseMerkleTreeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SparseMerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SparseMerkleTree *SparseMerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SparseMerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SparseMerkleTree *SparseMerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SparseMerkleTree.Contract.contract.Transact(opts, method, params...)
}

// CheckMembership is a free data retrieval call binding the contract method 0x6859274b.
//
// Solidity: function checkMembership(root bytes32, leafID uint64, proof bytes) constant returns(bool)
func (_SparseMerkleTree *SparseMerkleTreeCaller) CheckMembership(opts *bind.CallOpts, root [32]byte, leafID uint64, proof []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SparseMerkleTree.contract.Call(opts, out, "checkMembership", root, leafID, proof)
	return *ret0, err
}

// CheckMembership is a free data retrieval call binding the contract method 0x6859274b.
//
// Solidity: function checkMembership(root bytes32, leafID uint64, proof bytes) constant returns(bool)
func (_SparseMerkleTree *SparseMerkleTreeSession) CheckMembership(root [32]byte, leafID uint64, proof []byte) (bool, error) {
	return _SparseMerkleTree.Contract.CheckMembership(&_SparseMerkleTree.CallOpts, root, leafID, proof)
}

// CheckMembership is a free data retrieval call binding the contract method 0x6859274b.
//
// Solidity: function checkMembership(root bytes32, leafID uint64, proof bytes) constant returns(bool)
func (_SparseMerkleTree *SparseMerkleTreeCallerSession) CheckMembership(root [32]byte, leafID uint64, proof []byte) (bool, error) {
	return _SparseMerkleTree.Contract.CheckMembership(&_SparseMerkleTree.CallOpts, root, leafID, proof)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes( uint256) constant returns(bytes32)
func (_SparseMerkleTree *SparseMerkleTreeCaller) DefaultHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SparseMerkleTree.contract.Call(opts, out, "defaultHashes", arg0)
	return *ret0, err
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes( uint256) constant returns(bytes32)
func (_SparseMerkleTree *SparseMerkleTreeSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _SparseMerkleTree.Contract.DefaultHashes(&_SparseMerkleTree.CallOpts, arg0)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes( uint256) constant returns(bytes32)
func (_SparseMerkleTree *SparseMerkleTreeCallerSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _SparseMerkleTree.Contract.DefaultHashes(&_SparseMerkleTree.CallOpts, arg0)
}

// GetRoot is a free data retrieval call binding the contract method 0x7652ce3b.
//
// Solidity: function getRoot(index uint64, proof bytes) constant returns(bytes32)
func (_SparseMerkleTree *SparseMerkleTreeCaller) GetRoot(opts *bind.CallOpts, index uint64, proof []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SparseMerkleTree.contract.Call(opts, out, "getRoot", index, proof)
	return *ret0, err
}

// GetRoot is a free data retrieval call binding the contract method 0x7652ce3b.
//
// Solidity: function getRoot(index uint64, proof bytes) constant returns(bytes32)
func (_SparseMerkleTree *SparseMerkleTreeSession) GetRoot(index uint64, proof []byte) ([32]byte, error) {
	return _SparseMerkleTree.Contract.GetRoot(&_SparseMerkleTree.CallOpts, index, proof)
}

// GetRoot is a free data retrieval call binding the contract method 0x7652ce3b.
//
// Solidity: function getRoot(index uint64, proof bytes) constant returns(bytes32)
func (_SparseMerkleTree *SparseMerkleTreeCallerSession) GetRoot(index uint64, proof []byte) ([32]byte, error) {
	return _SparseMerkleTree.Contract.GetRoot(&_SparseMerkleTree.CallOpts, index, proof)
}
