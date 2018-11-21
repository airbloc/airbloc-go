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

// AccountsABI is the input ABI used to generate the binding from.
const AccountsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"numberOfAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"identityHashToAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"accounts\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"SignUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"TemporaryCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"accountId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"createTemporary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identityPreimage\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"unlockTemporary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"createUsingProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getAccountId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getAccountIdFromSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"isTemporary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Accounts is an auto generated Go binding around an Ethereum contract.
type Accounts struct {
	Address            common.Address
	AccountsCaller     // Read-only binding to the contract
	AccountsTransactor // Write-only binding to the contract
	AccountsFilterer   // Log filterer for contract events
}

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsSession struct {
	Contract     *Accounts         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsCallerSession struct {
	Contract *AccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsTransactorSession struct {
	Contract     *AccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsRaw struct {
	Contract *Accounts // Generic contract binding to access the raw methods on
}

// AccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsCallerRaw struct {
	Contract *AccountsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsTransactorRaw struct {
	Contract *AccountsTransactor // Generic write-only contract binding to access the raw methods on
}

func init() {
	blockchain.ContractList["Accounts"] = &Accounts{}
}

// NewAccounts creates a new instance of Accounts, bound to a specific deployed contract.
func NewAccounts(address common.Address, backend bind.ContractBackend) (*Accounts, error) {
	contract, err := bindAccounts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accounts{
		Address:            address,
		AccountsCaller:     AccountsCaller{contract: contract},
		AccountsTransactor: AccountsTransactor{contract: contract},
		AccountsFilterer:   AccountsFilterer{contract: contract},
	}, nil
}

// NewAccountsCaller creates a new read-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsCaller(address common.Address, caller bind.ContractCaller) (*AccountsCaller, error) {
	contract, err := bindAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsCaller{contract: contract}, nil
}

// NewAccountsTransactor creates a new write-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsTransactor, error) {
	contract, err := bindAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsTransactor{contract: contract}, nil
}

// NewAccountsFilterer creates a new log filterer instance of Accounts, bound to a specific deployed contract.
func NewAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountsFilterer, error) {
	contract, err := bindAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountsFilterer{contract: contract}, nil
}

// bindAccounts binds a generic wrapper to an already deployed contract.
func bindAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Accounts *Accounts) New(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewAccounts(address, backend)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.AccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts( bytes8) constant returns(owner address, status uint8, proxy address, passwordProof address)
func (_Accounts *AccountsCaller) Accounts(opts *bind.CallOpts, arg0 [8]byte) (struct {
	Owner         common.Address
	Status        uint8
	Proxy         common.Address
	PasswordProof common.Address
}, error) {
	ret := new(struct {
		Owner         common.Address
		Status        uint8
		Proxy         common.Address
		PasswordProof common.Address
	})
	out := ret
	err := _Accounts.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts( bytes8) constant returns(owner address, status uint8, proxy address, passwordProof address)
func (_Accounts *AccountsSession) Accounts(arg0 [8]byte) (struct {
	Owner         common.Address
	Status        uint8
	Proxy         common.Address
	PasswordProof common.Address
}, error) {
	return _Accounts.Contract.Accounts(&_Accounts.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts( bytes8) constant returns(owner address, status uint8, proxy address, passwordProof address)
func (_Accounts *AccountsCallerSession) Accounts(arg0 [8]byte) (struct {
	Owner         common.Address
	Status        uint8
	Proxy         common.Address
	PasswordProof common.Address
}, error) {
	return _Accounts.Contract.Accounts(&_Accounts.CallOpts, arg0)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(sender address) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountId(opts *bind.CallOpts, sender common.Address) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccountId", sender)
	return *ret0, err
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(sender address) constant returns(bytes8)
func (_Accounts *AccountsSession) GetAccountId(sender common.Address) ([8]byte, error) {
	return _Accounts.Contract.GetAccountId(&_Accounts.CallOpts, sender)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(sender address) constant returns(bytes8)
func (_Accounts *AccountsCallerSession) GetAccountId(sender common.Address) ([8]byte, error) {
	return _Accounts.Contract.GetAccountId(&_Accounts.CallOpts, sender)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(messageHash bytes32, signature bytes) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountIdFromSignature(opts *bind.CallOpts, messageHash [32]byte, signature []byte) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccountIdFromSignature", messageHash, signature)
	return *ret0, err
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(messageHash bytes32, signature bytes) constant returns(bytes8)
func (_Accounts *AccountsSession) GetAccountIdFromSignature(messageHash [32]byte, signature []byte) ([8]byte, error) {
	return _Accounts.Contract.GetAccountIdFromSignature(&_Accounts.CallOpts, messageHash, signature)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(messageHash bytes32, signature bytes) constant returns(bytes8)
func (_Accounts *AccountsCallerSession) GetAccountIdFromSignature(messageHash [32]byte, signature []byte) ([8]byte, error) {
	return _Accounts.Contract.GetAccountIdFromSignature(&_Accounts.CallOpts, messageHash, signature)
}

// IdentityHashToAccount is a free data retrieval call binding the contract method 0x17aba2d3.
//
// Solidity: function identityHashToAccount( bytes32) constant returns(bytes8)
func (_Accounts *AccountsCaller) IdentityHashToAccount(opts *bind.CallOpts, arg0 [32]byte) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "identityHashToAccount", arg0)
	return *ret0, err
}

// IdentityHashToAccount is a free data retrieval call binding the contract method 0x17aba2d3.
//
// Solidity: function identityHashToAccount( bytes32) constant returns(bytes8)
func (_Accounts *AccountsSession) IdentityHashToAccount(arg0 [32]byte) ([8]byte, error) {
	return _Accounts.Contract.IdentityHashToAccount(&_Accounts.CallOpts, arg0)
}

// IdentityHashToAccount is a free data retrieval call binding the contract method 0x17aba2d3.
//
// Solidity: function identityHashToAccount( bytes32) constant returns(bytes8)
func (_Accounts *AccountsCallerSession) IdentityHashToAccount(arg0 [32]byte) ([8]byte, error) {
	return _Accounts.Contract.IdentityHashToAccount(&_Accounts.CallOpts, arg0)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(accountId bytes8) constant returns(bool)
func (_Accounts *AccountsCaller) IsTemporary(opts *bind.CallOpts, accountId [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isTemporary", accountId)
	return *ret0, err
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(accountId bytes8) constant returns(bool)
func (_Accounts *AccountsSession) IsTemporary(accountId [8]byte) (bool, error) {
	return _Accounts.Contract.IsTemporary(&_Accounts.CallOpts, accountId)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(accountId bytes8) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsTemporary(accountId [8]byte) (bool, error) {
	return _Accounts.Contract.IsTemporary(&_Accounts.CallOpts, accountId)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (_Accounts *AccountsCaller) NumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "numberOfAccounts")
	return *ret0, err
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (_Accounts *AccountsSession) NumberOfAccounts() (*big.Int, error) {
	return _Accounts.Contract.NumberOfAccounts(&_Accounts.CallOpts)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (_Accounts *AccountsCallerSession) NumberOfAccounts() (*big.Int, error) {
	return _Accounts.Contract.NumberOfAccounts(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsCallerSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Accounts *AccountsTransactor) Create(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "create")
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Accounts *AccountsSession) Create() (*types.Transaction, error) {
	return _Accounts.Contract.Create(&_Accounts.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Accounts *AccountsTransactorSession) Create() (*types.Transaction, error) {
	return _Accounts.Contract.Create(&_Accounts.TransactOpts)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(identityHash bytes32) returns()
func (_Accounts *AccountsTransactor) CreateTemporary(opts *bind.TransactOpts, identityHash [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createTemporary", identityHash)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(identityHash bytes32) returns()
func (_Accounts *AccountsSession) CreateTemporary(identityHash [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.CreateTemporary(&_Accounts.TransactOpts, identityHash)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(identityHash bytes32) returns()
func (_Accounts *AccountsTransactorSession) CreateTemporary(identityHash [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.CreateTemporary(&_Accounts.TransactOpts, identityHash)
}

// CreateUsingProxy is a paid mutator transaction binding the contract method 0xdb82967c.
//
// Solidity: function createUsingProxy(owner address, passwordSignature bytes) returns()
func (_Accounts *AccountsTransactor) CreateUsingProxy(opts *bind.TransactOpts, owner common.Address, passwordSignature []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createUsingProxy", owner, passwordSignature)
}

// CreateUsingProxy is a paid mutator transaction binding the contract method 0xdb82967c.
//
// Solidity: function createUsingProxy(owner address, passwordSignature bytes) returns()
func (_Accounts *AccountsSession) CreateUsingProxy(owner common.Address, passwordSignature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.CreateUsingProxy(&_Accounts.TransactOpts, owner, passwordSignature)
}

// CreateUsingProxy is a paid mutator transaction binding the contract method 0xdb82967c.
//
// Solidity: function createUsingProxy(owner address, passwordSignature bytes) returns()
func (_Accounts *AccountsTransactorSession) CreateUsingProxy(owner common.Address, passwordSignature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.CreateUsingProxy(&_Accounts.TransactOpts, owner, passwordSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Accounts *AccountsTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Accounts *AccountsSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Accounts *AccountsTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, _newOwner)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(identityPreimage bytes32, newOwner address, passwordSignature bytes) returns()
func (_Accounts *AccountsTransactor) UnlockTemporary(opts *bind.TransactOpts, identityPreimage [32]byte, newOwner common.Address, passwordSignature []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "unlockTemporary", identityPreimage, newOwner, passwordSignature)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(identityPreimage bytes32, newOwner address, passwordSignature bytes) returns()
func (_Accounts *AccountsSession) UnlockTemporary(identityPreimage [32]byte, newOwner common.Address, passwordSignature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.UnlockTemporary(&_Accounts.TransactOpts, identityPreimage, newOwner, passwordSignature)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(identityPreimage bytes32, newOwner address, passwordSignature bytes) returns()
func (_Accounts *AccountsTransactorSession) UnlockTemporary(identityPreimage [32]byte, newOwner common.Address, passwordSignature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.UnlockTemporary(&_Accounts.TransactOpts, identityPreimage, newOwner, passwordSignature)
}

// AccountsOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Accounts contract.
type AccountsOwnershipRenouncedIterator struct {
	Event *AccountsOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *AccountsOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOwnershipRenounced)
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
		it.Event = new(AccountsOwnershipRenounced)
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
func (it *AccountsOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOwnershipRenounced represents a OwnershipRenounced event raised by the Accounts contract.
type AccountsOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Accounts *AccountsFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*AccountsOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOwnershipRenouncedIterator{contract: _Accounts.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// FilterOwnershipRenounced parses the event from given transaction receipt.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Accounts *AccountsFilterer) ParseOwnershipRenouncedFromReceipt(receipt *types.Receipt) (*AccountsOwnershipRenounced, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820") {
			event := new(AccountsOwnershipRenounced)
			if err := _Accounts.contract.UnpackLog(event, "OwnershipRenounced", *log); err != nil {
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
func (_Accounts *AccountsFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *AccountsOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOwnershipRenounced)
				if err := _Accounts.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// AccountsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Accounts contract.
type AccountsOwnershipTransferredIterator struct {
	Event *AccountsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOwnershipTransferred)
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
		it.Event = new(AccountsOwnershipTransferred)
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
func (it *AccountsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOwnershipTransferred represents a OwnershipTransferred event raised by the Accounts contract.
type AccountsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Accounts *AccountsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOwnershipTransferredIterator{contract: _Accounts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// FilterOwnershipTransferred parses the event from given transaction receipt.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Accounts *AccountsFilterer) ParseOwnershipTransferredFromReceipt(receipt *types.Receipt) (*AccountsOwnershipTransferred, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0") {
			event := new(AccountsOwnershipTransferred)
			if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
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
func (_Accounts *AccountsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOwnershipTransferred)
				if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// AccountsSignUpIterator is returned from FilterSignUp and is used to iterate over the raw logs and unpacked data for SignUp events raised by the Accounts contract.
type AccountsSignUpIterator struct {
	Event *AccountsSignUp // Event containing the contract specifics and raw log

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
func (it *AccountsSignUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsSignUp)
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
		it.Event = new(AccountsSignUp)
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
func (it *AccountsSignUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsSignUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsSignUp represents a SignUp event raised by the Accounts contract.
type AccountsSignUp struct {
	Owner     common.Address
	AccountId [8]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignUp is a free log retrieval operation binding the contract event 0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4.
//
// Solidity: e SignUp(owner indexed address, accountId bytes8)
func (_Accounts *AccountsFilterer) FilterSignUp(opts *bind.FilterOpts, owner []common.Address) (*AccountsSignUpIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "SignUp", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsSignUpIterator{contract: _Accounts.contract, event: "SignUp", logs: logs, sub: sub}, nil
}

// FilterSignUp parses the event from given transaction receipt.
//
// Solidity: e SignUp(owner indexed address, accountId bytes8)
func (_Accounts *AccountsFilterer) ParseSignUpFromReceipt(receipt *types.Receipt) (*AccountsSignUp, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4") {
			event := new(AccountsSignUp)
			if err := _Accounts.contract.UnpackLog(event, "SignUp", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("SignUp event not found")
}

// WatchSignUp is a free log subscription operation binding the contract event 0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4.
//
// Solidity: e SignUp(owner indexed address, accountId bytes8)
func (_Accounts *AccountsFilterer) WatchSignUp(opts *bind.WatchOpts, sink chan<- *AccountsSignUp, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "SignUp", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsSignUp)
				if err := _Accounts.contract.UnpackLog(event, "SignUp", log); err != nil {
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

// AccountsTemporaryCreatedIterator is returned from FilterTemporaryCreated and is used to iterate over the raw logs and unpacked data for TemporaryCreated events raised by the Accounts contract.
type AccountsTemporaryCreatedIterator struct {
	Event *AccountsTemporaryCreated // Event containing the contract specifics and raw log

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
func (it *AccountsTemporaryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsTemporaryCreated)
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
		it.Event = new(AccountsTemporaryCreated)
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
func (it *AccountsTemporaryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsTemporaryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsTemporaryCreated represents a TemporaryCreated event raised by the Accounts contract.
type AccountsTemporaryCreated struct {
	Proxy        common.Address
	IdentityHash [32]byte
	AccountId    [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTemporaryCreated is a free log retrieval operation binding the contract event 0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e.
//
// Solidity: e TemporaryCreated(proxy indexed address, identityHash indexed bytes32, accountId bytes8)
func (_Accounts *AccountsFilterer) FilterTemporaryCreated(opts *bind.FilterOpts, proxy []common.Address, identityHash [][32]byte) (*AccountsTemporaryCreatedIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "TemporaryCreated", proxyRule, identityHashRule)
	if err != nil {
		return nil, err
	}
	return &AccountsTemporaryCreatedIterator{contract: _Accounts.contract, event: "TemporaryCreated", logs: logs, sub: sub}, nil
}

// FilterTemporaryCreated parses the event from given transaction receipt.
//
// Solidity: e TemporaryCreated(proxy indexed address, identityHash indexed bytes32, accountId bytes8)
func (_Accounts *AccountsFilterer) ParseTemporaryCreatedFromReceipt(receipt *types.Receipt) (*AccountsTemporaryCreated, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e") {
			event := new(AccountsTemporaryCreated)
			if err := _Accounts.contract.UnpackLog(event, "TemporaryCreated", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("TemporaryCreated event not found")
}

// WatchTemporaryCreated is a free log subscription operation binding the contract event 0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e.
//
// Solidity: e TemporaryCreated(proxy indexed address, identityHash indexed bytes32, accountId bytes8)
func (_Accounts *AccountsFilterer) WatchTemporaryCreated(opts *bind.WatchOpts, sink chan<- *AccountsTemporaryCreated, proxy []common.Address, identityHash [][32]byte) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "TemporaryCreated", proxyRule, identityHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsTemporaryCreated)
				if err := _Accounts.contract.UnpackLog(event, "TemporaryCreated", log); err != nil {
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

// AccountsUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the Accounts contract.
type AccountsUnlockedIterator struct {
	Event *AccountsUnlocked // Event containing the contract specifics and raw log

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
func (it *AccountsUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsUnlocked)
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
		it.Event = new(AccountsUnlocked)
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
func (it *AccountsUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsUnlocked represents a Unlocked event raised by the Accounts contract.
type AccountsUnlocked struct {
	IdentityHash [32]byte
	AccountId    [8]byte
	NewOwner     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0.
//
// Solidity: e Unlocked(identityHash indexed bytes32, accountId indexed bytes8, newOwner address)
func (_Accounts *AccountsFilterer) FilterUnlocked(opts *bind.FilterOpts, identityHash [][32]byte, accountId [][8]byte) (*AccountsUnlockedIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Unlocked", identityHashRule, accountIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountsUnlockedIterator{contract: _Accounts.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// FilterUnlocked parses the event from given transaction receipt.
//
// Solidity: e Unlocked(identityHash indexed bytes32, accountId indexed bytes8, newOwner address)
func (_Accounts *AccountsFilterer) ParseUnlockedFromReceipt(receipt *types.Receipt) (*AccountsUnlocked, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0") {
			event := new(AccountsUnlocked)
			if err := _Accounts.contract.UnpackLog(event, "Unlocked", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Unlocked event not found")
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0.
//
// Solidity: e Unlocked(identityHash indexed bytes32, accountId indexed bytes8, newOwner address)
func (_Accounts *AccountsFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *AccountsUnlocked, identityHash [][32]byte, accountId [][8]byte) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Unlocked", identityHashRule, accountIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsUnlocked)
				if err := _Accounts.contract.UnpackLog(event, "Unlocked", log); err != nil {
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
