package managers

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	types "github.com/airbloc/airbloc-go/shared/adapter/types"
	wrappers "github.com/airbloc/airbloc-go/shared/adapter/wrappers"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source exchange.go -destination ./mocks/mock_exchange.go -package mocks IExchangeManager

type IExchangeManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	wrappers.IExchangeCalls

	// Transact methods
	AddDataIds(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		offerId types.ID,
		dataIds []types.DataId,
	) error

	Cancel(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		offerId types.ID,
	) error

	Order(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		offerId types.ID,
	) error

	Prepare(
		ctx context.Context,
		opts *ablbind.TransactOpts,
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
		opts *ablbind.TransactOpts,
		offerId types.ID,
	) error

	Settle(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		offerId types.ID,
	) error

	// Event methods
	wrappers.IExchangeFilterer
	wrappers.IExchangeWatcher
}

// exchangeManager is contract wrapper struct
type exchangeManager struct {
	wrappers.IExchangeContract
	client ablbind.ContractBackend
	log    *logger.Logger
}

// NewExchangeManager makes new *exchangeManager struct
func NewExchangeManager(client ablbind.ContractBackend, contract interface{}) interface{} {
	return &exchangeManager{
		IExchangeContract: contract.(*wrappers.ExchangeContract),
		client:            client,
		log:               logger.New("exchange"),
	}
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (manager *exchangeManager) Prepare(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	provider string,
	consumer common.Address,
	escrow common.Address,
	escrowSign [4]byte,
	escrowArgs []byte,
	dataIds []types.DataId,
) (types.ID, error) {
	var err error
	var ids []types.DataId
	// if length of dataIds exceeds 20,
	// this won't put dataIds when Prepare() calls. and makes array ids keeps nil state
	if len(dataIds) < 20 {
		ids = dataIds
	}
	receipt, err := manager.IExchangeContract.Prepare(ctx, opts, provider, consumer, escrow, escrowSign, escrowArgs, ids)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IExchangeContract.ParseOfferPreparedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer prepared.", logger.Attrs{
		"offer-id":          evt[0].OfferId.Hex(),
		"provider-app-name": evt[0].ProviderAppName,
	})

	// then, splits ids into chunks which maximum length is 20.
	// and adds in offer struct one by one.
	if ids == nil {
		l := len(dataIds)
		for i := 0; i < l; i += 20 {
			start := i
			end := i + 20
			if end > l {
				end = l
			}

			err = manager.AddDataIds(ctx, opts, evt[0].OfferId, dataIds[start:end])
			if err != nil {
				break
			}
		}
	}
	return evt[0].OfferId, nil
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (manager *exchangeManager) AddDataIds(ctx context.Context, opts *ablbind.TransactOpts, offerId types.ID, dataIds []types.DataId) error {
	_, err := manager.IExchangeContract.AddDataIds(ctx, opts, offerId, dataIds)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	manager.log.Info("Offer updated.", logger.Attrs{
		"offer-id":    offerId.Hex(),
		"data-length": len(dataIds),
	})
	return nil
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (manager *exchangeManager) Order(ctx context.Context, opts *ablbind.TransactOpts, offerId types.ID) error {
	receipt, err := manager.IExchangeContract.Order(ctx, opts, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IExchangeContract.ParseOfferPresentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer presented.", logger.Attrs{"offer-id": evt[0].OfferId.Hex()})
	return nil
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (manager *exchangeManager) Cancel(ctx context.Context, opts *ablbind.TransactOpts, offerId types.ID) error {
	receipt, err := manager.IExchangeContract.Cancel(ctx, opts, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IExchangeContract.ParseOfferCanceledFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer cancelled.", logger.Attrs{"offer-id": evt[0].OfferId.Hex()})
	return nil
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (manager *exchangeManager) Settle(ctx context.Context, opts *ablbind.TransactOpts, offerId types.ID) error {
	receipt, err := manager.IExchangeContract.Settle(ctx, opts, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IExchangeContract.ParseOfferSettledFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer settled", logger.Attrs{"offer-id": evt[0].OfferId.Hex()})
	return nil
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (manager *exchangeManager) Reject(ctx context.Context, opts *ablbind.TransactOpts, offerId types.ID) error {
	receipt, err := manager.IExchangeContract.Reject(ctx, opts, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IExchangeContract.ParseOfferRejectedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer rejected", logger.Attrs{"offer-id": evt[0].OfferId.Hex()})
	return nil
}
