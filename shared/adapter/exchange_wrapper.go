// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"

	blockchain "github.com/airbloc/airbloc-go/shared/blockchain"
	types "github.com/airbloc/airbloc-go/shared/types"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source exchange_wrapper.go -destination ./mocks/mock_exchange.go -package mocks IExchangeManager,IExchangeContract
type IExchangeManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IExchangeCalls

	// Transact methods
	AddDataIds(
		ctx context.Context,
		offerId types.ID,
		dataIds []types.DataId,
	) error

	Cancel(
		ctx context.Context,
		offerId types.ID,
	) error

	Order(
		ctx context.Context,
		offerId types.ID,
	) error

	Prepare(
		ctx context.Context,
		provider string,
		consumer common.Address,
		escrow common.Address,
		escrowSign [4]byte,
		escrowArgs []byte,
		dataIds []types.DataId,
	) (
		types.ID,
		error,
	)
	Reject(
		ctx context.Context,
		offerId types.ID,
	) error

	Settle(
		ctx context.Context,
		offerId types.ID,
	) error

	// Event methods
	IExchangeFilterer
	IExchangeWatcher
}

type IExchangeCalls interface {
	GetOffer(
		offerId types.ID,
	) (
		types.Offer,
		error,
	)
	GetOfferMembers(
		offerId types.ID,
	) (
		common.Address,
		common.Address,
		error,
	)
	OfferExists(
		offerId types.ID,
	) (
		bool,
		error,
	)
}

type IExchangeTransacts interface {
	AddDataIds(
		ctx context.Context,
		offerId types.ID,
		dataIds []types.DataId,
	) (*chainTypes.Receipt, error)
	Cancel(
		ctx context.Context,
		offerId types.ID,
	) (*chainTypes.Receipt, error)
	Order(
		ctx context.Context,
		offerId types.ID,
	) (*chainTypes.Receipt, error)
	Prepare(
		ctx context.Context,
		provider string,
		consumer common.Address,
		escrow common.Address,
		escrowSign [4]byte,
		escrowArgs []byte,
		dataIds []types.DataId,
	) (*chainTypes.Receipt, error)
	Reject(
		ctx context.Context,
		offerId types.ID,
	) (*chainTypes.Receipt, error)
	Settle(
		ctx context.Context,
		offerId types.ID,
	) (*chainTypes.Receipt, error)
}

type IExchangeEvents interface {
	IExchangeFilterer
	IExchangeParser
	IExchangeWatcher
}

type IExchangeFilterer interface {
	FilterEscrowExecutionFailed(
		opts *bind.FilterOpts,

	) (*ExchangeEscrowExecutionFailedIterator, error)
	FilterOfferCanceled(
		opts *bind.FilterOpts,
		offerId []types.ID,

	) (*ExchangeOfferCanceledIterator, error)
	FilterOfferPrepared(
		opts *bind.FilterOpts,
		offerId []types.ID,

	) (*ExchangeOfferPreparedIterator, error)
	FilterOfferPresented(
		opts *bind.FilterOpts,
		offerId []types.ID,

	) (*ExchangeOfferPresentedIterator, error)
	FilterOfferReceipt(
		opts *bind.FilterOpts,
		offerId []types.ID,

		consumer []common.Address,

	) (*ExchangeOfferReceiptIterator, error)
	FilterOfferRejected(
		opts *bind.FilterOpts,
		offerId []types.ID,
		consumer []common.Address,
	) (*ExchangeOfferRejectedIterator, error)
	FilterOfferSettled(
		opts *bind.FilterOpts,
		offerId []types.ID,
		consumer []common.Address,
	) (*ExchangeOfferSettledIterator, error)
}

type IExchangeParser interface {
	ParseEscrowExecutionFailed(log chainTypes.Log) (*ExchangeEscrowExecutionFailed, error)
	ParseEscrowExecutionFailedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeEscrowExecutionFailed, error)
	ParseOfferCanceled(log chainTypes.Log) (*ExchangeOfferCanceled, error)
	ParseOfferCanceledFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferCanceled, error)
	ParseOfferPrepared(log chainTypes.Log) (*ExchangeOfferPrepared, error)
	ParseOfferPreparedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferPrepared, error)
	ParseOfferPresented(log chainTypes.Log) (*ExchangeOfferPresented, error)
	ParseOfferPresentedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferPresented, error)
	ParseOfferReceipt(log chainTypes.Log) (*ExchangeOfferReceipt, error)
	ParseOfferReceiptFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferReceipt, error)
	ParseOfferRejected(log chainTypes.Log) (*ExchangeOfferRejected, error)
	ParseOfferRejectedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferRejected, error)
	ParseOfferSettled(log chainTypes.Log) (*ExchangeOfferSettled, error)
	ParseOfferSettledFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferSettled, error)
}

type IExchangeWatcher interface {
	WatchEscrowExecutionFailed(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeEscrowExecutionFailed,

	) (event.Subscription, error)
	WatchOfferCanceled(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferCanceled,
		offerId []types.ID,

	) (event.Subscription, error)
	WatchOfferPrepared(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferPrepared,
		offerId []types.ID,

	) (event.Subscription, error)
	WatchOfferPresented(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferPresented,
		offerId []types.ID,

	) (event.Subscription, error)
	WatchOfferReceipt(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferReceipt,
		offerId []types.ID,

		consumer []common.Address,

	) (event.Subscription, error)
	WatchOfferRejected(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferRejected,
		offerId []types.ID,
		consumer []common.Address,
	) (event.Subscription, error)
	WatchOfferSettled(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferSettled,
		offerId []types.ID,
		consumer []common.Address,
	) (event.Subscription, error)
}

type IExchangeContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IExchangeCalls
	IExchangeTransacts
	IExchangeEvents
}

// Manager is contract wrapper struct
type ExchangeContract struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int
	client    blockchain.TxClient

	ExchangeCaller
	ExchangeFilterer
	ExchangeTransactor
}

// Address is getter method of Accounts.address
func (c *ExchangeContract) Address() common.Address {
	return c.address
}

// TxHash is getter method of Accounts.txHash
func (c *ExchangeContract) TxHash() common.Hash {
	return c.txHash
}

// CreatedAt is getter method of Accounts.createdAt
func (c *ExchangeContract) CreatedAt() *big.Int {
	return c.createdAt
}

func newExchangeContract(address common.Address, txHash common.Hash, createdAt *big.Int, parsedABI abi.ABI, backend bind.ContractBackend) interface{} {
	contract := bind.NewBoundContract(address, parsedABI, backend, backend, backend)

	return &ExchangeContract{
		address:   address,
		txHash:    txHash,
		createdAt: createdAt,
		client:    backend.(blockchain.TxClient),

		ExchangeCaller:     ExchangeCaller{contract: contract},
		ExchangeTransactor: ExchangeTransactor{contract: contract},
		ExchangeFilterer:   ExchangeFilterer{contract: contract},
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("Exchange", newExchangeContract)
	blockchain.RegisterSelector("0x367a9005", "addDataIds(bytes8,bytes20[])")
	blockchain.RegisterSelector("0xb2d9ba39", "cancel(bytes8)")
	blockchain.RegisterSelector("0x0cf833fb", "order(bytes8)")
	blockchain.RegisterSelector("0x77e61c33", "prepare(string,address,address,bytes4,bytes,bytes20[])")
	blockchain.RegisterSelector("0x6622e153", "reject(bytes8)")
	blockchain.RegisterSelector("0xa60d9b5f", "settle(bytes8)")
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (c *ExchangeContract) GetOffer(
	offerId types.ID,
) (

	types.Offer,
	error,
) {
	return c.ExchangeCaller.GetOffer(nil, offerId)
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (c *ExchangeContract) GetOfferMembers(
	offerId types.ID,
) (

	common.Address,
	common.Address,
	error,
) {
	return c.ExchangeCaller.GetOfferMembers(nil, offerId)
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (c *ExchangeContract) OfferExists(
	offerId types.ID,
) (

	bool,
	error,
) {
	return c.ExchangeCaller.OfferExists(nil, offerId)
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (c *ExchangeContract) AddDataIds(
	ctx context.Context,
	offerId types.ID,
	dataIds []types.DataId,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.AddDataIds(c.client.Account(), offerId, dataIds)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (c *ExchangeContract) Cancel(
	ctx context.Context,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Cancel(c.client.Account(), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (c *ExchangeContract) Order(
	ctx context.Context,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Order(c.client.Account(), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (c *ExchangeContract) Prepare(
	ctx context.Context,
	provider string,
	consumer common.Address,
	escrow common.Address,
	escrowSign [4]byte,
	escrowArgs []byte,
	dataIds []types.DataId,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Prepare(c.client.Account(), provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (c *ExchangeContract) Reject(
	ctx context.Context,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Reject(c.client.Account(), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (c *ExchangeContract) Settle(
	ctx context.Context,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Settle(c.client.Account(), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
