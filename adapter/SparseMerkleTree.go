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

// SparseMerkleTreeABI is the input ABI used to generate the binding from.
const SparseMerkleTreeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"defaultHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"leafID\",\"type\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"checkMembership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"getRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SparseMerkleTreeBin is the compiled bytecode used for deploying new contracts.
const SparseMerkleTreeBin = `0x6080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166348419ad8811461005b5780636859274b146100855780637652ce3b14610103575b600080fd5b34801561006757600080fd5b5061007360043561016b565b60408051918252519081900360200190f35b34801561009157600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526100ef948235946024803567ffffffffffffffff169536959460649492019190819084018382808284375094975061017f9650505050505050565b604080519115158252519081900360200190f35b34801561010f57600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261007395833567ffffffffffffffff169536956044949193909101919081908401838280828437509497506101979650505050505050565b6000816041811061017857fe5b0154905081565b60008061018c8484610197565b909414949350505050565b600080600080600080602060088851038115156101b057fe5b061580156101c15750610808875111155b15156101cc57600080fd5b5050505060208301516001906008907801000000000000000000000000000000000000000000000000900460005b604081101561037c57600182161515610223576000816041811061021a57fe5b01549450610244565b6020830192508261ffff1687511015151561023d57600080fd5b8287015194505b6001881615156102d357604080516020808201879052818301889052825180830384018152606090920192839052815191929182918401908083835b6020831061029f5780518252601f199092019160209182019101610280565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209350610354565b604080516020808201889052818301879052825180830384018152606090920192839052815191929182918401908083835b602083106103245780518252601f199092019160209182019101610305565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902093505b600267ffffffffffffffff8316049150600267ffffffffffffffff89160497506001016101fa565b509196955050505050505600a165627a7a723058209a93ae1fef53df7718b7e222a21a3b4572a71d65abe5569d8a876ef922f8af860029`

// DeploySparseMerkleTree deploys a new Ethereum contract, binding an instance of SparseMerkleTree to it.
func DeploySparseMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SparseMerkleTree, error) {
	parsed, err := abi.JSON(strings.NewReader(SparseMerkleTreeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SparseMerkleTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SparseMerkleTree{SparseMerkleTreeCaller: SparseMerkleTreeCaller{contract: contract}, SparseMerkleTreeTransactor: SparseMerkleTreeTransactor{contract: contract}, SparseMerkleTreeFilterer: SparseMerkleTreeFilterer{contract: contract}}, nil
}

// SparseMerkleTree is an auto generated Go binding around an Ethereum contract.
type SparseMerkleTree struct {
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

// NewSparseMerkleTree creates a new instance of SparseMerkleTree, bound to a specific deployed contract.
func NewSparseMerkleTree(address common.Address, backend bind.ContractBackend) (*SparseMerkleTree, error) {
	contract, err := bindSparseMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SparseMerkleTree{SparseMerkleTreeCaller: SparseMerkleTreeCaller{contract: contract}, SparseMerkleTreeTransactor: SparseMerkleTreeTransactor{contract: contract}, SparseMerkleTreeFilterer: SparseMerkleTreeFilterer{contract: contract}}, nil
}

// NewSparseMerkleTreeCaller creates a new read-only instance of SparseMerkleTree, bound to a specific deployed contract.
func NewSparseMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*SparseMerkleTreeCaller, error) {
	contract, err := bindSparseMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SparseMerkleTreeCaller{contract: contract}, nil
}

// NewSparseMerkleTreeTransactor creates a new write-only instance of SparseMerkleTree, bound to a specific deployed contract.
func NewSparseMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*SparseMerkleTreeTransactor, error) {
	contract, err := bindSparseMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SparseMerkleTreeTransactor{contract: contract}, nil
}

// NewSparseMerkleTreeFilterer creates a new log filterer instance of SparseMerkleTree, bound to a specific deployed contract.
func NewSparseMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*SparseMerkleTreeFilterer, error) {
	contract, err := bindSparseMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SparseMerkleTreeFilterer{contract: contract}, nil
}

// bindSparseMerkleTree binds a generic wrapper to an already deployed contract.
func bindSparseMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SparseMerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
