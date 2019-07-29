package exchange

import (
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// Manager is contract wrapper struct
type manager struct {
	contract adapter.IExchangeContract
	log      *logger.Logger
}

// NewManager makes new *manager struct
func NewManager(client blockchain.TxClient) adapter.IExchangeManager {
	return &manager{
		contract: adapter.NewExchangeContract(client),
		log:      logger.New("exchange"),
	}
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (manager *manager) Prepare(
	ctx context.Context,
	provider string,
	consumer ethCommon.Address,
	escrow ethCommon.Address,
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
	receipt, err := manager.contract.Prepare(ctx, provider, consumer, escrow, escrowSign, escrowArgs, ids)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseOfferPreparedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer prepared.", logger.Attrs{
		"offer-id":          event.OfferId.Hex(),
		"provider-app-name": event.ProviderAppName,
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

			err = manager.AddDataIds(ctx, event.OfferId, dataIds[start:end])
			if err != nil {
				break
			}
		}
	}
	return event.OfferId, err
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (manager *manager) AddDataIds(ctx context.Context, offerId types.ID, dataIds []types.DataId) error {
	_, err := manager.contract.AddDataIds(ctx, offerId, dataIds)
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
func (manager *manager) Order(ctx context.Context, offerId types.ID) error {
	receipt, err := manager.contract.Order(ctx, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseOfferPresentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer presented.", logger.Attrs{"offer-id": event.OfferId.Hex()})
	return err
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (manager *manager) Cancel(ctx context.Context, offerId types.ID) error {
	receipt, err := manager.contract.Cancel(ctx, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseOfferCanceledFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer cancelled.", logger.Attrs{"offer-id": event.OfferId.Hex()})
	return err
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (manager *manager) Settle(ctx context.Context, offerId types.ID) error {
	receipt, err := manager.contract.Settle(ctx, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseOfferSettledFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer settled", logger.Attrs{"offer-id": event.OfferId.Hex()})
	return nil
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (manager *manager) Reject(ctx context.Context, offerId types.ID) error {
	receipt, err := manager.contract.Reject(ctx, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseOfferRejectedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer rejected", logger.Attrs{"offer-id": event.OfferId.Hex()})
	return nil
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (manager *manager) GetOffer(offerId types.ID) (types.Offer, error) {
	return manager.contract.GetOffer(offerId)
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (manager *manager) GetOfferMembers(offerId types.ID) (ethCommon.Address, ethCommon.Address, error) {
	return manager.contract.GetOfferMembers(offerId)
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (manager *manager) OfferExists(offerId types.ID) (bool, error) {
	return manager.contract.OfferExists(offerId)
}

// FilterEscrowExecutionFailed is a free log retrieval operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (manager manager) FilterEscrowExecutionFailed(opts *bind.FilterOpts) (*adapter.ExchangeEscrowExecutionFailedIterator, error) {
	return manager.contract.FilterEscrowExecutionFailed(opts)
}

// WatchEscrowExecutionFailed is a free log subscription operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (manager manager) WatchEscrowExecutionFailed(opts *bind.WatchOpts, sink chan<- *adapter.ExchangeEscrowExecutionFailed) (event.Subscription, error) {
	return manager.contract.WatchEscrowExecutionFailed(opts, sink)
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (manager manager) FilterOfferCanceled(opts *bind.FilterOpts, offerId []types.ID) (*adapter.ExchangeOfferCanceledIterator, error) {
	return manager.contract.FilterOfferCanceled(opts, offerId)
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (manager manager) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *adapter.ExchangeOfferCanceled, offerId []types.ID) (event.Subscription, error) {
	return manager.contract.WatchOfferCanceled(opts, sink, offerId)
}

// FilterOfferPrepared is a free log retrieval operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (manager manager) FilterOfferPrepared(opts *bind.FilterOpts, offerId []types.ID) (*adapter.ExchangeOfferPreparedIterator, error) {
	return manager.contract.FilterOfferPrepared(opts, offerId)
}

// WatchOfferPrepared is a free log subscription operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (manager manager) WatchOfferPrepared(opts *bind.WatchOpts, sink chan<- *adapter.ExchangeOfferPrepared, offerId []types.ID) (event.Subscription, error) {
	return manager.contract.WatchOfferPrepared(opts, sink, offerId)
}

// FilterOfferPresented is a free log retrieval operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (manager manager) FilterOfferPresented(opts *bind.FilterOpts, offerId []types.ID) (*adapter.ExchangeOfferPresentedIterator, error) {
	return manager.contract.FilterOfferPresented(opts, offerId)
}

// WatchOfferPresented is a free log subscription operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (manager manager) WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *adapter.ExchangeOfferPresented, offerId []types.ID) (event.Subscription, error) {
	return manager.contract.WatchOfferPresented(opts, sink, offerId)
}

// FilterOfferReceipt is a free log retrieval operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (manager manager) FilterOfferReceipt(opts *bind.FilterOpts, offerId []types.ID, consumer []ethCommon.Address) (*adapter.ExchangeOfferReceiptIterator, error) {
	return manager.contract.FilterOfferReceipt(opts, offerId, consumer)
}

// WatchOfferReceipt is a free log subscription operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (manager manager) WatchOfferReceipt(opts *bind.WatchOpts, sink chan<- *adapter.ExchangeOfferReceipt, offerId []types.ID, consumer []ethCommon.Address) (event.Subscription, error) {
	return manager.contract.WatchOfferReceipt(opts, sink, offerId, consumer)
}

// FilterOfferRejected is a free log retrieval operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (manager manager) FilterOfferRejected(opts *bind.FilterOpts, offerId []types.ID, consumer []ethCommon.Address) (*adapter.ExchangeOfferRejectedIterator, error) {
	return manager.contract.FilterOfferRejected(opts, offerId, consumer)
}

// WatchOfferRejected is a free log subscription operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (manager manager) WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *adapter.ExchangeOfferRejected, offerId []types.ID, consumer []ethCommon.Address) (event.Subscription, error) {
	return manager.contract.WatchOfferRejected(opts, sink, offerId, consumer)
}

// FilterOfferSettled is a free log retrieval operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (manager manager) FilterOfferSettled(opts *bind.FilterOpts, offerId []types.ID, consumer []ethCommon.Address) (*adapter.ExchangeOfferSettledIterator, error) {
	return manager.contract.FilterOfferSettled(opts, offerId, consumer)
}

// WatchOfferSettled is a free log subscription operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (manager manager) WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *adapter.ExchangeOfferSettled, offerId []types.ID, consumer []ethCommon.Address) (event.Subscription, error) {
	return manager.contract.WatchOfferSettled(opts, sink, offerId, consumer)
}
