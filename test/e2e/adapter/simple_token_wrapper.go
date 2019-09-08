// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"math/big"

	blockchain "github.com/airbloc/airbloc-go/shared/blockchain"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source simple_token_wrapper.go -destination ./mocks/mock_simple_token.go -package mocks ISimpleTokenManager,ISimpleTokenContract
type ISimpleTokenManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	ISimpleTokenCalls

	// Transact methods

	// Event methods
	ISimpleTokenFilterer
	ISimpleTokenWatcher
}

type ISimpleTokenCalls interface {
}

type ISimpleTokenTransacts interface {
}

type ISimpleTokenEvents interface {
	ISimpleTokenFilterer
	ISimpleTokenParser
	ISimpleTokenWatcher
}

type ISimpleTokenFilterer interface {
	FilterApproval(
		opts *bind.FilterOpts,
		owner []common.Address,
		spender []common.Address,

	) (*SimpleTokenApprovalIterator, error)
	FilterMinterAdded(
		opts *bind.FilterOpts,
		account []common.Address,
	) (*SimpleTokenMinterAddedIterator, error)
	FilterMinterRemoved(
		opts *bind.FilterOpts,
		account []common.Address,
	) (*SimpleTokenMinterRemovedIterator, error)
	FilterTransfer(
		opts *bind.FilterOpts,
		from []common.Address,
		to []common.Address,

	) (*SimpleTokenTransferIterator, error)
}

type ISimpleTokenParser interface {
	ParseApproval(log chainTypes.Log) (*SimpleTokenApproval, error)
	ParseApprovalFromReceipt(receipt *chainTypes.Receipt) ([]*SimpleTokenApproval, error)
	ParseMinterAdded(log chainTypes.Log) (*SimpleTokenMinterAdded, error)
	ParseMinterAddedFromReceipt(receipt *chainTypes.Receipt) ([]*SimpleTokenMinterAdded, error)
	ParseMinterRemoved(log chainTypes.Log) (*SimpleTokenMinterRemoved, error)
	ParseMinterRemovedFromReceipt(receipt *chainTypes.Receipt) ([]*SimpleTokenMinterRemoved, error)
	ParseTransfer(log chainTypes.Log) (*SimpleTokenTransfer, error)
	ParseTransferFromReceipt(receipt *chainTypes.Receipt) ([]*SimpleTokenTransfer, error)
}

type ISimpleTokenWatcher interface {
	WatchApproval(
		opts *bind.WatchOpts,
		sink chan<- *SimpleTokenApproval,
		owner []common.Address,
		spender []common.Address,

	) (event.Subscription, error)
	WatchMinterAdded(
		opts *bind.WatchOpts,
		sink chan<- *SimpleTokenMinterAdded,
		account []common.Address,
	) (event.Subscription, error)
	WatchMinterRemoved(
		opts *bind.WatchOpts,
		sink chan<- *SimpleTokenMinterRemoved,
		account []common.Address,
	) (event.Subscription, error)
	WatchTransfer(
		opts *bind.WatchOpts,
		sink chan<- *SimpleTokenTransfer,
		from []common.Address,
		to []common.Address,

	) (event.Subscription, error)
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
}
