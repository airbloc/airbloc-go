// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source simple_token_wrapper.go -destination ./mocks/mock_simple_token.go -package mocks ISimpleTokenManager,ISimpleTokenContract
type ISimpleTokenManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	ISimpleTokenCalls

	// Transact methods
	AddMinter(ctx context.Context, account common.Address) error
	Approve(ctx context.Context, spender common.Address, value *big.Int) (bool, error)
	DecreaseAllowance(ctx context.Context, spender common.Address, subtractedValue *big.Int) (bool, error)
	IncreaseAllowance(ctx context.Context, spender common.Address, addedValue *big.Int) (bool, error)
	Mint(ctx context.Context, account common.Address, amount *big.Int) (bool, error)
	RenounceMinter(ctx context.Context) error
	Transfer(ctx context.Context, recipient common.Address, amount *big.Int) (bool, error)
	TransferFrom(ctx context.Context, sender common.Address, recipient common.Address, amount *big.Int) (bool, error)

	// Event methods
	ISimpleTokenFilterer
	ISimpleTokenWatcher
}

type ISimpleTokenCalls interface {
	Allowance(owner common.Address, spender common.Address) (*big.Int, error)
	BalanceOf(account common.Address) (*big.Int, error)
	IsMinter(account common.Address) (bool, error)
	TotalSupply() (*big.Int, error)
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

type ISimpleTokenEvents interface {
	ISimpleTokenFilterer
	ISimpleTokenParser
	ISimpleTokenWatcher
}

type ISimpleTokenFilterer interface {
	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SimpleTokenApprovalIterator, error)
	FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*SimpleTokenMinterAddedIterator, error)
	FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*SimpleTokenMinterRemovedIterator, error)
	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleTokenTransferIterator, error)
}

type ISimpleTokenParser interface {
	ParseApprovalFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenApproval, error)
	ParseMinterAddedFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenMinterAdded, error)
	ParseMinterRemovedFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenMinterRemoved, error)
	ParseTransferFromReceipt(receipt *klayTypes.Receipt) (*SimpleTokenTransfer, error)
}

type ISimpleTokenWatcher interface {
	WatchApproval(opts *bind.WatchOpts, sink chan<- *SimpleTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)
	WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *SimpleTokenMinterAdded, account []common.Address) (event.Subscription, error)
	WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *SimpleTokenMinterRemoved, account []common.Address) (event.Subscription, error)
	WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error)
}

type ISimpleTokenContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	ISimpleTokenCalls
	ISimpleTokenTransacts
	ISimpleTokenEvents
}

// Manager is contract wrapper struct
type SimpleTokenContract struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int
	client    blockchain.TxClient

	SimpleTokenCaller
	SimpleTokenFilterer
	SimpleTokenTransactor
}

// Address is getter method of Accounts.address
func (c *SimpleTokenContract) Address() common.Address {
	return c.address
}

// TxHash is getter method of Accounts.txHash
func (c *SimpleTokenContract) TxHash() common.Hash {
	return c.txHash
}

// CreatedAt is getter method of Accounts.createdAt
func (c *SimpleTokenContract) CreatedAt() *big.Int {
	return c.createdAt
}

func newSimpleTokenContract(address common.Address, txHash common.Hash, createdAt *big.Int, parsedABI abi.ABI, backend bind.ContractBackend) interface{} {
	contract := bind.NewBoundContract(address, parsedABI, backend, backend, backend)

	return &SimpleTokenContract{
		address:   address,
		txHash:    txHash,
		createdAt: createdAt,
		client:    backend.(blockchain.TxClient),

		SimpleTokenCaller:     SimpleTokenCaller{contract: contract},
		SimpleTokenTransactor: SimpleTokenTransactor{contract: contract},
		SimpleTokenFilterer:   SimpleTokenFilterer{contract: contract},
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("SimpleToken", newSimpleTokenContract)
	blockchain.RegisterSelector("0x983b2d56", "addMinter(address)")
	blockchain.RegisterSelector("0x095ea7b3", "approve(address,uint256)")
	blockchain.RegisterSelector("0xa457c2d7", "decreaseAllowance(address,uint256)")
	blockchain.RegisterSelector("0x39509351", "increaseAllowance(address,uint256)")
	blockchain.RegisterSelector("0x40c10f19", "mint(address,uint256)")
	blockchain.RegisterSelector("0x98650275", "renounceMinter()")
	blockchain.RegisterSelector("0xa9059cbb", "transfer(address,uint256)")
	blockchain.RegisterSelector("0x23b872dd", "transferFrom(address,address,uint256)")
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (c *SimpleTokenContract) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return c.SimpleTokenCaller.Allowance(nil, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (c *SimpleTokenContract) BalanceOf(account common.Address) (*big.Int, error) {
	return c.SimpleTokenCaller.BalanceOf(nil, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (c *SimpleTokenContract) IsMinter(account common.Address) (bool, error) {
	return c.SimpleTokenCaller.IsMinter(nil, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (c *SimpleTokenContract) TotalSupply() (*big.Int, error) {
	return c.SimpleTokenCaller.TotalSupply(nil)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (c *SimpleTokenContract) AddMinter(ctx context.Context, account common.Address) (*klayTypes.Receipt, error) {
	tx, err := c.SimpleTokenTransactor.AddMinter(c.client.Account(), account)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (c *SimpleTokenContract) Approve(ctx context.Context, spender common.Address, value *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.SimpleTokenTransactor.Approve(c.client.Account(), spender, value)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (c *SimpleTokenContract) DecreaseAllowance(ctx context.Context, spender common.Address, subtractedValue *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.SimpleTokenTransactor.DecreaseAllowance(c.client.Account(), spender, subtractedValue)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (c *SimpleTokenContract) IncreaseAllowance(ctx context.Context, spender common.Address, addedValue *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.SimpleTokenTransactor.IncreaseAllowance(c.client.Account(), spender, addedValue)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (c *SimpleTokenContract) Mint(ctx context.Context, account common.Address, amount *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.SimpleTokenTransactor.Mint(c.client.Account(), account, amount)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (c *SimpleTokenContract) RenounceMinter(ctx context.Context) (*klayTypes.Receipt, error) {
	tx, err := c.SimpleTokenTransactor.RenounceMinter(c.client.Account())
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (c *SimpleTokenContract) Transfer(ctx context.Context, recipient common.Address, amount *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.SimpleTokenTransactor.Transfer(c.client.Account(), recipient, amount)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (c *SimpleTokenContract) TransferFrom(ctx context.Context, sender common.Address, recipient common.Address, amount *big.Int) (*klayTypes.Receipt, error) {
	tx, err := c.SimpleTokenTransactor.TransferFrom(c.client.Account(), sender, recipient, amount)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
