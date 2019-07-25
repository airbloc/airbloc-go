// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
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
	_ = types.HexToID
	_ = common.Big1
	_ = ethTypes.BloomLookup
	_ = event.NewSubscription
)

// AccountsABI is the input ABI used to generate the binding from.
const AccountsABI = "{\"Constructor\":{\"Name\":\"\",\"Const\":false,\"Inputs\":[{\"Name\":\"controllerReg\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":null},\"Methods\":{\"accounts\":{\"Name\":\"accounts\",\"Const\":true,\"Inputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"owner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"status\",\"Type\":{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"controller\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"passwordProof\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"create\":{\"Name\":\"create\",\"Const\":false,\"Inputs\":[],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"createTemporary\":{\"Name\":\"createTemporary\",\"Const\":false,\"Inputs\":[{\"Name\":\"identityHash\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"exists\":{\"Name\":\"exists\",\"Const\":true,\"Inputs\":[{\"Name\":\"accountId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"getAccount\":{\"Name\":\"getAccount\",\"Const\":true,\"Inputs\":[{\"Name\":\"accountId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":25,\"Type\":{},\"Size\":0,\"T\":6,\"TupleElems\":[{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null}],\"TupleRawNames\":[\"owner\",\"status\",\"controller\",\"passwordProof\"]},\"Indexed\":false}]},\"getAccountId\":{\"Name\":\"getAccountId\",\"Const\":true,\"Inputs\":[{\"Name\":\"sender\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"getAccountIdFromSignature\":{\"Name\":\"getAccountIdFromSignature\",\"Const\":true,\"Inputs\":[{\"Name\":\"messageHash\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"signature\",\"Type\":{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"identityHashToAccount\":{\"Name\":\"identityHashToAccount\",\"Const\":true,\"Inputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"isControllerOf\":{\"Name\":\"isControllerOf\",\"Const\":true,\"Inputs\":[{\"Name\":\"sender\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"accountId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"isTemporary\":{\"Name\":\"isTemporary\",\"Const\":true,\"Inputs\":[{\"Name\":\"accountId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"numberOfAccounts\":{\"Name\":\"numberOfAccounts\",\"Const\":true,\"Inputs\":[],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":22,\"Type\":{},\"Size\":256,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"setController\":{\"Name\":\"setController\",\"Const\":false,\"Inputs\":[{\"Name\":\"controller\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"unlockTemporary\":{\"Name\":\"unlockTemporary\",\"Const\":false,\"Inputs\":[{\"Name\":\"identityPreimage\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"newOwner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"passwordSignature\",\"Type\":{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]}},\"Events\":{\"SignUp\":{\"Name\":\"SignUp\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"owner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"accountId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"TemporaryCreated\":{\"Name\":\"TemporaryCreated\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"proxy\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"identityHash\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"accountId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"Unlocked\":{\"Name\":\"Unlocked\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"identityHash\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":32,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"accountId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"newOwner\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]}}}"

// Accounts is an auto generated Go binding around an Ethereum contract.
type Accounts struct {
	Address            common.Address
	AccountsCaller     // Read-only binding to the contract
	AccountsTransactor // Write-only binding to the contract
	AccountsFilterer   // Log filterer for contract events
}

// AccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsSession struct {
	Contract     *Accounts         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsRaw struct {
	Contract *Accounts // Generic contract binding to access the raw methods on
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

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.AccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transact(opts, method, params...)
}

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsCallerSession struct {
	Contract *AccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsCallerRaw struct {
	Contract *AccountsCaller // Generic read-only contract binding to access the raw methods on
}

// NewAccountsCaller creates a new read-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsCaller(address common.Address, caller bind.ContractCaller) (*AccountsCaller, error) {
	contract, err := bindAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsCaller{contract: contract}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.contract.Call(opts, result, method, params...)
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsTransactorSession struct {
	Contract     *AccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsTransactorRaw struct {
	Contract *AccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountsTransactor creates a new write-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsTransactor, error) {
	contract, err := bindAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.contract.Transact(opts, method, params...)
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewAccountsFilterer creates a new log filterer instance of Accounts, bound to a specific deployed contract.
func NewAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountsFilterer, error) {
	contract, err := bindAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountsFilterer{contract: contract}, nil
}

type AccountsManager interface {
	// Pure/View methods
	Accounts(arg0 types.ID) (types.Account, error)
	Exists(accountId types.ID) (bool, error)
	GetAccount(accountId types.ID) (types.Account, error)
	GetAccountId(sender common.Address) (types.ID, error)
	GetAccountIdFromSignature(messageHash [32]byte, signature []byte) (types.ID, error)
	IdentityHashToAccount(arg0 [32]byte) (types.ID, error)
	IsControllerOf(sender common.Address, accountId types.ID) (bool, error)
	IsTemporary(accountId types.ID) (bool, error)
	NumberOfAccounts() (*big.Int, error)

	// Other methods
	Create(ctx context.Context) (types.ID, error)
	CreateTemporary(ctx context.Context, identityHash [32]byte) (types.ID, error)
	SetController(ctx context.Context, controller common.Address) error
	UnlockTemporary(ctx context.Context, identityPreimage [32]byte, newOwner common.Address, passwordSignature []byte) error
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.ContractList["Accounts"] = (&Accounts{}).new
	blockchain.RegisterSelector("0xefc81a8c", "create()")
	blockchain.RegisterSelector("0x56003f0f", "createTemporary(bytes32)")
	blockchain.RegisterSelector("0x92eefe9b", "setController(address)")
	blockchain.RegisterSelector("0x2299219d", "unlockTemporary(bytes32,address,bytes)")
}

// bindAccounts binds a generic wrapper to an already deployed contract.
func bindAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Accounts *Accounts) new(address common.Address, backend bind.ContractBackend) (interface{}, error) {
	return NewAccounts(address, backend)
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts(bytes8 ) constant returns(address owner, uint8 status, address controller, address passwordProof)
func (_Accounts *AccountsCaller) Accounts(opts *bind.CallOpts, arg0 types.ID) (types.Account, error) {
	ret := new(types.Account)

	out := ret
	err := _Accounts.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts(bytes8 ) constant returns(address owner, uint8 status, address controller, address passwordProof)
func (_Accounts *AccountsSession) Accounts(arg0 types.ID) (types.Account, error) {
	return _Accounts.Contract.Accounts(&_Accounts.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts(bytes8 ) constant returns(address owner, uint8 status, address controller, address passwordProof)
func (_Accounts *AccountsCallerSession) Accounts(arg0 types.ID) (types.Account, error) {
	return _Accounts.Contract.Accounts(&_Accounts.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCaller) Exists(opts *bind.CallOpts, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _Accounts.contract.Call(opts, out, "exists", accountId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsSession) Exists(accountId types.ID) (bool, error) {
	return _Accounts.Contract.Exists(&_Accounts.CallOpts, accountId)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCallerSession) Exists(accountId types.ID) (bool, error) {
	return _Accounts.Contract.Exists(&_Accounts.CallOpts, accountId)
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (_Accounts *AccountsCaller) GetAccount(opts *bind.CallOpts, accountId types.ID) (types.Account, error) {
	ret := new(types.Account)

	out := ret
	err := _Accounts.contract.Call(opts, out, "getAccount", accountId)
	return *ret, err
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (_Accounts *AccountsSession) GetAccount(accountId types.ID) (types.Account, error) {
	return _Accounts.Contract.GetAccount(&_Accounts.CallOpts, accountId)
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (_Accounts *AccountsCallerSession) GetAccount(accountId types.ID) (types.Account, error) {
	return _Accounts.Contract.GetAccount(&_Accounts.CallOpts, accountId)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountId(opts *bind.CallOpts, sender common.Address) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := &[]interface{}{ret0}
	err := _Accounts.contract.Call(opts, out, "getAccountId", sender)
	return *ret0, err
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (_Accounts *AccountsSession) GetAccountId(sender common.Address) (types.ID, error) {
	return _Accounts.Contract.GetAccountId(&_Accounts.CallOpts, sender)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (_Accounts *AccountsCallerSession) GetAccountId(sender common.Address) (types.ID, error) {
	return _Accounts.Contract.GetAccountId(&_Accounts.CallOpts, sender)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountIdFromSignature(opts *bind.CallOpts, messageHash [32]byte, signature []byte) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := &[]interface{}{ret0}
	err := _Accounts.contract.Call(opts, out, "getAccountIdFromSignature", messageHash, signature)
	return *ret0, err
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (_Accounts *AccountsSession) GetAccountIdFromSignature(messageHash [32]byte, signature []byte) (types.ID, error) {
	return _Accounts.Contract.GetAccountIdFromSignature(&_Accounts.CallOpts, messageHash, signature)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (_Accounts *AccountsCallerSession) GetAccountIdFromSignature(messageHash [32]byte, signature []byte) (types.ID, error) {
	return _Accounts.Contract.GetAccountIdFromSignature(&_Accounts.CallOpts, messageHash, signature)
}

// IdentityHashToAccount is a free data retrieval call binding the contract method 0x17aba2d3.
//
// Solidity: function identityHashToAccount(bytes32 ) constant returns(bytes8)
func (_Accounts *AccountsCaller) IdentityHashToAccount(opts *bind.CallOpts, arg0 [32]byte) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := &[]interface{}{ret0}
	err := _Accounts.contract.Call(opts, out, "identityHashToAccount", arg0)
	return *ret0, err
}

// IdentityHashToAccount is a free data retrieval call binding the contract method 0x17aba2d3.
//
// Solidity: function identityHashToAccount(bytes32 ) constant returns(bytes8)
func (_Accounts *AccountsSession) IdentityHashToAccount(arg0 [32]byte) (types.ID, error) {
	return _Accounts.Contract.IdentityHashToAccount(&_Accounts.CallOpts, arg0)
}

// IdentityHashToAccount is a free data retrieval call binding the contract method 0x17aba2d3.
//
// Solidity: function identityHashToAccount(bytes32 ) constant returns(bytes8)
func (_Accounts *AccountsCallerSession) IdentityHashToAccount(arg0 [32]byte) (types.ID, error) {
	return _Accounts.Contract.IdentityHashToAccount(&_Accounts.CallOpts, arg0)
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address sender, bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCaller) IsControllerOf(opts *bind.CallOpts, sender common.Address, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _Accounts.contract.Call(opts, out, "isControllerOf", sender, accountId)
	return *ret0, err
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address sender, bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsSession) IsControllerOf(sender common.Address, accountId types.ID) (bool, error) {
	return _Accounts.Contract.IsControllerOf(&_Accounts.CallOpts, sender, accountId)
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address sender, bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsControllerOf(sender common.Address, accountId types.ID) (bool, error) {
	return _Accounts.Contract.IsControllerOf(&_Accounts.CallOpts, sender, accountId)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCaller) IsTemporary(opts *bind.CallOpts, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _Accounts.contract.Call(opts, out, "isTemporary", accountId)
	return *ret0, err
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsSession) IsTemporary(accountId types.ID) (bool, error) {
	return _Accounts.Contract.IsTemporary(&_Accounts.CallOpts, accountId)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsTemporary(accountId types.ID) (bool, error) {
	return _Accounts.Contract.IsTemporary(&_Accounts.CallOpts, accountId)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (_Accounts *AccountsCaller) NumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := &[]interface{}{ret0}
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

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (_Accounts *AccountsTransactor) Create(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _Accounts.contract.Transact(opts, "create")
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (_Accounts *AccountsSession) Create() (*ethTypes.Transaction, error) {
	return _Accounts.Contract.Create(&_Accounts.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (_Accounts *AccountsTransactorSession) Create() (*ethTypes.Transaction, error) {
	return _Accounts.Contract.Create(&_Accounts.TransactOpts)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (_Accounts *AccountsTransactor) CreateTemporary(opts *bind.TransactOpts, identityHash [32]byte) (*ethTypes.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createTemporary", identityHash)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (_Accounts *AccountsSession) CreateTemporary(identityHash [32]byte) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.CreateTemporary(&_Accounts.TransactOpts, identityHash)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (_Accounts *AccountsTransactorSession) CreateTemporary(identityHash [32]byte) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.CreateTemporary(&_Accounts.TransactOpts, identityHash)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (_Accounts *AccountsTransactor) SetController(opts *bind.TransactOpts, controller common.Address) (*ethTypes.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setController", controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (_Accounts *AccountsSession) SetController(controller common.Address) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.SetController(&_Accounts.TransactOpts, controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (_Accounts *AccountsTransactorSession) SetController(controller common.Address) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.SetController(&_Accounts.TransactOpts, controller)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (_Accounts *AccountsTransactor) UnlockTemporary(opts *bind.TransactOpts, identityPreimage [32]byte, newOwner common.Address, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Accounts.contract.Transact(opts, "unlockTemporary", identityPreimage, newOwner, passwordSignature)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (_Accounts *AccountsSession) UnlockTemporary(identityPreimage [32]byte, newOwner common.Address, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.UnlockTemporary(&_Accounts.TransactOpts, identityPreimage, newOwner, passwordSignature)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (_Accounts *AccountsTransactorSession) UnlockTemporary(identityPreimage [32]byte, newOwner common.Address, passwordSignature []byte) (*ethTypes.Transaction, error) {
	return _Accounts.Contract.UnlockTemporary(&_Accounts.TransactOpts, identityPreimage, newOwner, passwordSignature)
}

// AccountsSignUpIterator is returned from FilterSignUp and is used to iterate over the raw logs and unpacked data for SignUp events raised by the Accounts contract.
type AccountsSignUpIterator struct {
	Event *AccountsSignUp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
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
	AccountId types.ID
	Raw       ethTypes.Log // Blockchain specific contextual infos
}

// FilterSignUp is a free log retrieval operation binding the contract event 0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4.
//
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
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
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
func (_Accounts *AccountsFilterer) ParseSignUpFromReceipt(receipt *ethTypes.Receipt) (*AccountsSignUp, error) {
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
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
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

	logs chan ethTypes.Log     // Log channel receiving the found contract events
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
	AccountId    types.ID
	Raw          ethTypes.Log // Blockchain specific contextual infos
}

// FilterTemporaryCreated is a free log retrieval operation binding the contract event 0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e.
//
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
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
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
func (_Accounts *AccountsFilterer) ParseTemporaryCreatedFromReceipt(receipt *ethTypes.Receipt) (*AccountsTemporaryCreated, error) {
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
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
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

	logs chan ethTypes.Log     // Log channel receiving the found contract events
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
	AccountId    types.ID
	NewOwner     common.Address
	Raw          ethTypes.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0.
//
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (_Accounts *AccountsFilterer) FilterUnlocked(opts *bind.FilterOpts, identityHash [][32]byte, accountId []types.ID) (*AccountsUnlockedIterator, error) {

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
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (_Accounts *AccountsFilterer) ParseUnlockedFromReceipt(receipt *ethTypes.Receipt) (*AccountsUnlocked, error) {
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
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (_Accounts *AccountsFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *AccountsUnlocked, identityHash [][32]byte, accountId []types.ID) (event.Subscription, error) {

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
