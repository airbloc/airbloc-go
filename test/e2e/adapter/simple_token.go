// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testadapter

import (
	"context"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/types"
	ethereum "github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
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
	_ = klayTypes.BloomLookup
	_ = event.NewSubscription
)

// SimpleTokenABI is the input ABI used to generate the binding from.
const SimpleTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x095ea7b3\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"signature\":\"0x18160ddd\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x23b872dd\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x39509351\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x40c10f19\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"signature\":\"0x70a08231\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"signature\":\"0x983b2d56\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"signature\":\"0x98650275\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xa457c2d7\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xa9059cbb\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xaa271e1a\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"signature\":\"0xdd62ed3e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"signature\":\"0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"signature\":\"0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"signature\":\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"signature\":\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\"type\":\"event\"}]"

// SimpleToken is an auto generated Go binding around an Ethereum contract.
type SimpleToken struct {
	address               common.Address
	txHash                common.Hash
	createdAt             *big.Int
	SimpleTokenCaller     // Read-only binding to the contract
	SimpleTokenTransactor // Write-only binding to the contract
	SimpleTokenFilterer   // Log filterer for contract events
}

// Address is getter method of SimpleToken.address
func (_SimpleToken *SimpleToken) Address() common.Address {
	return _SimpleToken.address
}

// TxHash is getter method of SimpleToken.txHash
func (_SimpleToken *SimpleToken) TxHash() common.Hash {
	return _SimpleToken.txHash
}

// CreatedAt is getter method of SimpleToken.createdAt
func (_SimpleToken *SimpleToken) CreatedAt() *big.Int {
	return _SimpleToken.createdAt
}

// SimpleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleTokenSession struct {
	Contract     *SimpleToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleTokenRaw struct {
	Contract *SimpleToken // Generic contract binding to access the raw methods on
}

// NewSimpleToken creates a new instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleToken(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (*SimpleToken, error) {
	contract, err := bindSimpleToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleToken{
		address:               address,
		txHash:                txHash,
		createdAt:             createdAt,
		SimpleTokenCaller:     SimpleTokenCaller{contract: contract},
		SimpleTokenTransactor: SimpleTokenTransactor{contract: contract},
		SimpleTokenFilterer:   SimpleTokenFilterer{contract: contract},
	}, nil
}

// bindSimpleToken binds a generic wrapper to an already deployed contract.
func bindSimpleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleToken *SimpleTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleToken.Contract.SimpleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleToken *SimpleTokenRaw) Transfer(opts *bind.TransactOpts) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.SimpleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleToken *SimpleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.SimpleTokenTransactor.contract.Transact(opts, method, params...)
}

// SimpleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleTokenCallerSession struct {
	Contract *SimpleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SimpleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleTokenCallerRaw struct {
	Contract *SimpleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NewSimpleTokenCaller creates a new read-only instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenCaller(address common.Address, caller bind.ContractCaller) (*SimpleTokenCaller, error) {
	contract, err := bindSimpleToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenCaller{contract: contract}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleToken *SimpleTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleToken.Contract.contract.Call(opts, result, method, params...)
}

// SimpleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleTokenTransactorSession struct {
	Contract     *SimpleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SimpleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleTokenTransactorRaw struct {
	Contract *SimpleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleTokenTransactor creates a new write-only instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTokenTransactor, error) {
	contract, err := bindSimpleToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleToken *SimpleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleToken *SimpleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.contract.Transact(opts, method, params...)
}

// SimpleTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewSimpleTokenFilterer creates a new log filterer instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleTokenFilterer, error) {
	contract, err := bindSimpleToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenFilterer{contract: contract}, nil
}

//go:generate mockgen -source simple_token.go -destination ./mocks/mock_simple_token.go -package mocks ISimpleTokenManager,ISimpleTokenContract
type ISimpleTokenManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	Allowance(owner common.Address, spender common.Address) (*big.Int, error)
	BalanceOf(account common.Address) (*big.Int, error)
	IsMinter(account common.Address) (bool, error)
	TotalSupply() (*big.Int, error)

	// Transact methods
	AddMinter(ctx context.Context, account common.Address) error
	Approve(ctx context.Context, spender common.Address, value *big.Int) (bool, error)
	DecreaseAllowance(ctx context.Context, spender common.Address, subtractedValue *big.Int) (bool, error)
	IncreaseAllowance(ctx context.Context, spender common.Address, addedValue *big.Int) (bool, error)
	Mint(ctx context.Context, account common.Address, amount *big.Int) (bool, error)
	RenounceMinter(ctx context.Context) error
	Transfer(ctx context.Context, recipient common.Address, amount *big.Int) (bool, error)
	TransferFrom(ctx context.Context, sender common.Address, recipient common.Address, amount *big.Int) (bool, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SimpleTokenApprovalIterator, error)
	WatchApproval(opts *bind.WatchOpts, sink chan<- *SimpleTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*SimpleTokenMinterAddedIterator, error)
	WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *SimpleTokenMinterAdded, account []common.Address) (event.Subscription, error)

	FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*SimpleTokenMinterRemovedIterator, error)
	WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *SimpleTokenMinterRemoved, account []common.Address) (event.Subscription, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleTokenTransferIterator, error)
	WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error)
}

type ISimpleTokenContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int
	Filterer() SimpleTokenFilterer

	ISimpleTokenCalls
	ISimpleTokenTransacts
	ISimpleTokenEvents
}

// SimpleTokenContract is contract wrapper struct
type SimpleTokenContract struct {
	client   blockchain.TxClient
	contract *SimpleToken
	SimpleTokenFilterer
}

// Address is getter method of SimpleToken.address
func (c *SimpleTokenContract) Address() common.Address {
	return c.contract.Address()
}

// TxHash is getter method of SimpleToken.txHash
func (c *SimpleTokenContract) TxHash() common.Hash {
	return c.contract.TxHash()
}

// CreatedAt is getter method of SimpleToken.createdAt
func (c *SimpleTokenContract) CreatedAt() *big.Int {
	return c.contract.CreatedAt()
}

// Filterer is getter method of SimpleToken.SimpleTokenFilterer
func (c *SimpleTokenContract) Filterer() SimpleTokenFilterer {
	return c.SimpleTokenFilterer
}

// NewSimpleTokenContract makes new *SimpleTokenContract struct
func NewSimpleTokenContract(client blockchain.TxClient) ISimpleTokenContract {
	contract := client.GetContract(&SimpleToken{}).(*SimpleToken)
	return &SimpleTokenContract{
		client:              client,
		contract:            contract,
		SimpleTokenFilterer: contract.SimpleTokenFilterer,
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("SimpleToken", (&SimpleToken{}).new)
	blockchain.RegisterSelector("0x983b2d56", "addMinter(address)")
	blockchain.RegisterSelector("0x095ea7b3", "approve(address,uint256)")
	blockchain.RegisterSelector("0xa457c2d7", "decreaseAllowance(address,uint256)")
	blockchain.RegisterSelector("0x39509351", "increaseAllowance(address,uint256)")
	blockchain.RegisterSelector("0x40c10f19", "mint(address,uint256)")
	blockchain.RegisterSelector("0x98650275", "renounceMinter()")
	blockchain.RegisterSelector("0xa9059cbb", "transfer(address,uint256)")
	blockchain.RegisterSelector("0x23b872dd", "transferFrom(address,address,uint256)")
}

func (_SimpleToken *SimpleToken) new(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (interface{}, error) {
	return NewSimpleToken(address, txHash, createdAt, backend)
}

type ISimpleTokenCalls interface {
	Allowance(owner common.Address, spender common.Address) (*big.Int, error)
	BalanceOf(account common.Address) (*big.Int, error)
	IsMinter(account common.Address) (bool, error)
	TotalSupply() (*big.Int, error)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (c *SimpleTokenContract) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return c.contract.Allowance(nil, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := &[]interface{}{ret0}
	err := _SimpleToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.Allowance(&_SimpleToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.Allowance(&_SimpleToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (c *SimpleTokenContract) BalanceOf(account common.Address) (*big.Int, error) {
	return c.contract.BalanceOf(nil, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := &[]interface{}{ret0}
	err := _SimpleToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.BalanceOf(&_SimpleToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.BalanceOf(&_SimpleToken.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (c *SimpleTokenContract) IsMinter(account common.Address) (bool, error) {
	return c.contract.IsMinter(nil, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_SimpleToken *SimpleTokenCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _SimpleToken.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_SimpleToken *SimpleTokenSession) IsMinter(account common.Address) (bool, error) {
	return _SimpleToken.Contract.IsMinter(&_SimpleToken.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_SimpleToken *SimpleTokenCallerSession) IsMinter(account common.Address) (bool, error) {
	return _SimpleToken.Contract.IsMinter(&_SimpleToken.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (c *SimpleTokenContract) TotalSupply() (*big.Int, error) {
	return c.contract.TotalSupply(nil)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := &[]interface{}{ret0}
	err := _SimpleToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) TotalSupply() (*big.Int, error) {
	return _SimpleToken.Contract.TotalSupply(&_SimpleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SimpleToken.Contract.TotalSupply(&_SimpleToken.CallOpts)
}

type ISimpleTokenTransacts interface {
	AddMinter(ctx context.Context, account common.Address) (*klayTypes.Receipt, error)
	Approve(ctx context.Context, spender common.Address, value *big.Int) (*klayTypes.Receipt, error)
	DecreaseAllowance(ctx context.Context, spender common.Address, subtractedValue *big.Int) (*klayTypes.Receipt, error)
	IncreaseAllowance(ctx context.Context, spender common.Address, addedValue *big.Int) (*klayTypes.Receipt, error)
	Mint(ctx context.Context, account common.Address, amount *big.Int) (*klayTypes.Receipt, error)
	RenounceMinter(ctx context.Context) (*klayTypes.Receipt, error)
	Transfer(ctx context.Context, recipient common.Address, amount *big.Int) (*klayTypes.Receipt, error)
	TransferFrom(ctx context.Context, sender common.Address, recipient common.Address, amount *big.Int) (*klayTypes.Receipt, error)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (c *SimpleTokenContract) AddMinter(ctx context.Context, account common.Address) (*klayTypes.Receipt, error) {
	tx, err := c.contract.AddMinter(c.client.Account(), account)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_SimpleToken *SimpleTokenTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*klayTypes.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_SimpleToken *SimpleTokenSession) AddMinter(account common.Address) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.AddMinter(&_SimpleToken.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_SimpleToken *SimpleTokenTransactorSession) AddMinter(account common.Address) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.AddMinter(&_SimpleToken.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (c *SimpleTokenContract) Approve(ctx context.Context, spender common.Address, value *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.contract.Approve(c.client.Account(), spender, value)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SimpleToken *SimpleTokenSession) Approve(spender common.Address, value *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.Approve(&_SimpleToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.Approve(&_SimpleToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (c *SimpleTokenContract) DecreaseAllowance(ctx context.Context, spender common.Address, subtractedValue *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.contract.DecreaseAllowance(c.client.Account(), spender, subtractedValue)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SimpleToken *SimpleTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.DecreaseAllowance(&_SimpleToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.DecreaseAllowance(&_SimpleToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (c *SimpleTokenContract) IncreaseAllowance(ctx context.Context, spender common.Address, addedValue *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.contract.IncreaseAllowance(c.client.Account(), spender, addedValue)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SimpleToken *SimpleTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.IncreaseAllowance(&_SimpleToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.IncreaseAllowance(&_SimpleToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (c *SimpleTokenContract) Mint(ctx context.Context, account common.Address, amount *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.contract.Mint(c.client.Account(), account, amount)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenSession) Mint(account common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.Mint(&_SimpleToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.Mint(&_SimpleToken.TransactOpts, account, amount)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (c *SimpleTokenContract) RenounceMinter(ctx context.Context) (*klayTypes.Receipt, error) {
	tx, err := c.contract.RenounceMinter(c.client.Account())
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SimpleToken *SimpleTokenTransactor) RenounceMinter(opts *bind.TransactOpts) (*klayTypes.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SimpleToken *SimpleTokenSession) RenounceMinter() (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.RenounceMinter(&_SimpleToken.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SimpleToken *SimpleTokenTransactorSession) RenounceMinter() (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.RenounceMinter(&_SimpleToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (c *SimpleTokenContract) Transfer(ctx context.Context, recipient common.Address, amount *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.contract.Transfer(c.client.Account(), recipient, amount)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenSession) Transfer(recipient common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.Transfer(&_SimpleToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.Transfer(&_SimpleToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (c *SimpleTokenContract) TransferFrom(ctx context.Context, sender common.Address, recipient common.Address, amount *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.contract.TransferFrom(c.client.Account(), sender, recipient, amount)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.TransferFrom(&_SimpleToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*klayTypes.Transaction, error) {
	return _SimpleToken.Contract.TransferFrom(&_SimpleToken.TransactOpts, sender, recipient, amount)
}

type ISimpleTokenEvents interface {
	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SimpleTokenApprovalIterator, error)
	ParseApprovalFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenApproval, error)
	WatchApproval(opts *bind.WatchOpts, sink chan<- *SimpleTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*SimpleTokenMinterAddedIterator, error)
	ParseMinterAddedFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenMinterAdded, error)
	WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *SimpleTokenMinterAdded, account []common.Address) (event.Subscription, error)

	FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*SimpleTokenMinterRemovedIterator, error)
	ParseMinterRemovedFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenMinterRemoved, error)
	WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *SimpleTokenMinterRemoved, account []common.Address) (event.Subscription, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleTokenTransferIterator, error)
	ParseTransferFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenTransfer, error)
	WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error)
}

// SimpleTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SimpleToken contract.
type SimpleTokenApprovalIterator struct {
	Event *SimpleTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan klayTypes.Log    // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTokenApproval)
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
		it.Event = new(SimpleTokenApproval)
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
func (it *SimpleTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTokenApproval represents a Approval event raised by the SimpleToken contract.
type SimpleTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     klayTypes.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SimpleToken *SimpleTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SimpleTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SimpleToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenApprovalIterator{contract: _SimpleToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// FilterApproval parses the event from given transaction receipt.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (manager *SimpleTokenContract) ParseApprovalFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenApproval, error) {
	return manager.contract.ParseApprovalFromReceipt(receipt)
}

// FilterApproval parses the event from given transaction receipt.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SimpleToken *SimpleTokenFilterer) ParseApprovalFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenApproval, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925") {
			event := new(SimpleTokenApproval)
			if err := _SimpleToken.contract.UnpackLog(event, "Approval", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Approval event not found")
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SimpleToken *SimpleTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SimpleTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SimpleToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTokenApproval)
				if err := _SimpleToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SimpleTokenMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the SimpleToken contract.
type SimpleTokenMinterAddedIterator struct {
	Event *SimpleTokenMinterAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan klayTypes.Log    // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleTokenMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTokenMinterAdded)
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
		it.Event = new(SimpleTokenMinterAdded)
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
func (it *SimpleTokenMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTokenMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTokenMinterAdded represents a MinterAdded event raised by the SimpleToken contract.
type SimpleTokenMinterAdded struct {
	Account common.Address
	Raw     klayTypes.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_SimpleToken *SimpleTokenFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*SimpleTokenMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SimpleToken.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenMinterAddedIterator{contract: _SimpleToken.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// FilterMinterAdded parses the event from given transaction receipt.
//
// Solidity: event MinterAdded(address indexed account)
func (manager *SimpleTokenContract) ParseMinterAddedFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenMinterAdded, error) {
	return manager.contract.ParseMinterAddedFromReceipt(receipt)
}

// FilterMinterAdded parses the event from given transaction receipt.
//
// Solidity: event MinterAdded(address indexed account)
func (_SimpleToken *SimpleTokenFilterer) ParseMinterAddedFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenMinterAdded, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6") {
			event := new(SimpleTokenMinterAdded)
			if err := _SimpleToken.contract.UnpackLog(event, "MinterAdded", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("MinterAdded event not found")
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_SimpleToken *SimpleTokenFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *SimpleTokenMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SimpleToken.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTokenMinterAdded)
				if err := _SimpleToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// SimpleTokenMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the SimpleToken contract.
type SimpleTokenMinterRemovedIterator struct {
	Event *SimpleTokenMinterRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan klayTypes.Log    // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleTokenMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTokenMinterRemoved)
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
		it.Event = new(SimpleTokenMinterRemoved)
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
func (it *SimpleTokenMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTokenMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTokenMinterRemoved represents a MinterRemoved event raised by the SimpleToken contract.
type SimpleTokenMinterRemoved struct {
	Account common.Address
	Raw     klayTypes.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_SimpleToken *SimpleTokenFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*SimpleTokenMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SimpleToken.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenMinterRemovedIterator{contract: _SimpleToken.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// FilterMinterRemoved parses the event from given transaction receipt.
//
// Solidity: event MinterRemoved(address indexed account)
func (manager *SimpleTokenContract) ParseMinterRemovedFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenMinterRemoved, error) {
	return manager.contract.ParseMinterRemovedFromReceipt(receipt)
}

// FilterMinterRemoved parses the event from given transaction receipt.
//
// Solidity: event MinterRemoved(address indexed account)
func (_SimpleToken *SimpleTokenFilterer) ParseMinterRemovedFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenMinterRemoved, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692") {
			event := new(SimpleTokenMinterRemoved)
			if err := _SimpleToken.contract.UnpackLog(event, "MinterRemoved", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("MinterRemoved event not found")
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_SimpleToken *SimpleTokenFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *SimpleTokenMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SimpleToken.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTokenMinterRemoved)
				if err := _SimpleToken.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// SimpleTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SimpleToken contract.
type SimpleTokenTransferIterator struct {
	Event *SimpleTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan klayTypes.Log    // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTokenTransfer)
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
		it.Event = new(SimpleTokenTransfer)
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
func (it *SimpleTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTokenTransfer represents a Transfer event raised by the SimpleToken contract.
type SimpleTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   klayTypes.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SimpleToken *SimpleTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenTransferIterator{contract: _SimpleToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// FilterTransfer parses the event from given transaction receipt.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (manager *SimpleTokenContract) ParseTransferFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenTransfer, error) {
	return manager.contract.ParseTransferFromReceipt(receipt)
}

// FilterTransfer parses the event from given transaction receipt.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SimpleToken *SimpleTokenFilterer) ParseTransferFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenTransfer, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") {
			event := new(SimpleTokenTransfer)
			if err := _SimpleToken.contract.UnpackLog(event, "Transfer", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Transfer event not found")
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SimpleToken *SimpleTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTokenTransfer)
				if err := _SimpleToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
