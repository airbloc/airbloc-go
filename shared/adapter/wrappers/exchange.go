package wrappers

import (
	"context"
	"math/big"
	"strings"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	contracts "github.com/airbloc/airbloc-go/shared/adapter/contracts"
	types "github.com/airbloc/airbloc-go/shared/adapter/types"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source exchange.go -destination ./mocks/mock_exchange.go -package mocks IExchangeContract

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
		opts *ablbind.TransactOpts,
		offerId types.ID,
		dataIds []types.DataId,
	) (*chainTypes.Receipt, error)
	Cancel(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		offerId types.ID,
	) (*chainTypes.Receipt, error)
	Order(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		offerId types.ID,
	) (*chainTypes.Receipt, error)
	Prepare(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		provider string,
		consumer common.Address,
		escrow common.Address,
		escrowSign [4]byte,
		escrowArgs []byte,
		dataIds []types.DataId,
	) (*chainTypes.Receipt, error)
	Reject(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		offerId types.ID,
	) (*chainTypes.Receipt, error)
	Settle(
		ctx context.Context,
		opts *ablbind.TransactOpts,
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

	) (ablbind.EventIterator, error)
	FilterOfferCanceled(
		opts *bind.FilterOpts,
		offerId []types.ID,

	) (ablbind.EventIterator, error)
	FilterOfferPrepared(
		opts *bind.FilterOpts,
		offerId []types.ID,

	) (ablbind.EventIterator, error)
	FilterOfferPresented(
		opts *bind.FilterOpts,
		offerId []types.ID,

	) (ablbind.EventIterator, error)
	FilterOfferReceipt(
		opts *bind.FilterOpts,
		offerId []types.ID,

		consumer []common.Address,

	) (ablbind.EventIterator, error)
	FilterOfferRejected(
		opts *bind.FilterOpts,
		offerId []types.ID,
		consumer []common.Address,
	) (ablbind.EventIterator, error)
	FilterOfferSettled(
		opts *bind.FilterOpts,
		offerId []types.ID,
		consumer []common.Address,
	) (ablbind.EventIterator, error)
}

type IExchangeParser interface {
	ParseEscrowExecutionFailed(log chainTypes.Log) (*contracts.ExchangeEscrowExecutionFailed, error)
	ParseEscrowExecutionFailedFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ExchangeEscrowExecutionFailed, error)
	ParseOfferCanceled(log chainTypes.Log) (*contracts.ExchangeOfferCanceled, error)
	ParseOfferCanceledFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ExchangeOfferCanceled, error)
	ParseOfferPrepared(log chainTypes.Log) (*contracts.ExchangeOfferPrepared, error)
	ParseOfferPreparedFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ExchangeOfferPrepared, error)
	ParseOfferPresented(log chainTypes.Log) (*contracts.ExchangeOfferPresented, error)
	ParseOfferPresentedFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ExchangeOfferPresented, error)
	ParseOfferReceipt(log chainTypes.Log) (*contracts.ExchangeOfferReceipt, error)
	ParseOfferReceiptFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ExchangeOfferReceipt, error)
	ParseOfferRejected(log chainTypes.Log) (*contracts.ExchangeOfferRejected, error)
	ParseOfferRejectedFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ExchangeOfferRejected, error)
	ParseOfferSettled(log chainTypes.Log) (*contracts.ExchangeOfferSettled, error)
	ParseOfferSettledFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ExchangeOfferSettled, error)
}

type IExchangeWatcher interface {
	WatchEscrowExecutionFailed(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ExchangeEscrowExecutionFailed,

	) (event.Subscription, error)
	WatchOfferCanceled(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ExchangeOfferCanceled,
		offerId []types.ID,

	) (event.Subscription, error)
	WatchOfferPrepared(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ExchangeOfferPrepared,
		offerId []types.ID,

	) (event.Subscription, error)
	WatchOfferPresented(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ExchangeOfferPresented,
		offerId []types.ID,

	) (event.Subscription, error)
	WatchOfferReceipt(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ExchangeOfferReceipt,
		offerId []types.ID,

		consumer []common.Address,

	) (event.Subscription, error)
	WatchOfferRejected(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ExchangeOfferRejected,
		offerId []types.ID,
		consumer []common.Address,
	) (event.Subscription, error)
	WatchOfferSettled(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ExchangeOfferSettled,
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
	ablbind.Deployment
	client ablbind.ContractBackend

	contracts.ExchangeCaller
	contracts.ExchangeFilterer
	contracts.ExchangeTransactor
}

func NewExchangeContract(deployment ablbind.Deployment, backend ablbind.ContractBackend) interface{} {
	if deployment.Address() == (common.Address{}) {
		evmABI, err := abi.JSON(strings.NewReader(contracts.ExchangeABI))
		if err != nil {
			panic(err)
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(contracts.ExchangeAddress),
			common.HexToHash(contracts.ExchangeTxHash),
			new(big.Int).SetBytes(common.HexToHash(contracts.ExchangeCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, backend)

	contract := &ExchangeContract{
		Deployment: deployment,
		client:     backend,

		ExchangeCaller:     contracts.NewExchangeCaller(base),
		ExchangeTransactor: contracts.NewExchangeTransactor(base),
		ExchangeFilterer:   contracts.NewExchangeFilterer(base),
	}

	return contract
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
	opts *ablbind.TransactOpts,
	offerId types.ID,
	dataIds []types.DataId,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.AddDataIds(c.client.Transactor(ctx, opts), offerId, dataIds)
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
	opts *ablbind.TransactOpts,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Cancel(c.client.Transactor(ctx, opts), offerId)
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
	opts *ablbind.TransactOpts,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Order(c.client.Transactor(ctx, opts), offerId)
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
	opts *ablbind.TransactOpts,
	provider string,
	consumer common.Address,
	escrow common.Address,
	escrowSign [4]byte,
	escrowArgs []byte,
	dataIds []types.DataId,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Prepare(c.client.Transactor(ctx, opts), provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
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
	opts *ablbind.TransactOpts,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Reject(c.client.Transactor(ctx, opts), offerId)
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
	opts *ablbind.TransactOpts,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := c.ExchangeTransactor.Settle(c.client.Transactor(ctx, opts), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
