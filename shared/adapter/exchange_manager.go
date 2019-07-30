package adapter

import (
	"math/big"

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
type exchangeManager struct {
	contract IExchangeContract
	log      *logger.Logger
}

// Address is getter method of Exchange.address
func (manager *exchangeManager) Address() ethCommon.Address {
	return manager.contract.Address()
}

// TxHash is getter method of Exchange.txHash
func (manager *exchangeManager) TxHash() ethCommon.Hash {
	return manager.contract.TxHash()
}

// CreatedAt is getter method of Exchange.createdAt
func (manager *exchangeManager) CreatedAt() *big.Int {
	return manager.contract.CreatedAt()
}

// NewExchangeManager makes new *accountsManager struct
func NewExchangeManager(client blockchain.TxClient) IExchangeManager {
	return &exchangeManager{
		contract: NewExchangeContract(client),
		log:      logger.New("exchange"),
	}
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (manager *exchangeManager) Prepare(
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

	evt, err := manager.contract.ParseOfferPreparedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer prepared.", logger.Attrs{
		"offer-id":          evt.OfferId.Hex(),
		"provider-app-name": evt.ProviderAppName,
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

			err = manager.AddDataIds(ctx, evt.OfferId, dataIds[start:end])
			if err != nil {
				break
			}
		}
	}
	return evt.OfferId, err
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (manager *exchangeManager) AddDataIds(ctx context.Context, offerId types.ID, dataIds []types.DataId) error {
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
func (manager *exchangeManager) Order(ctx context.Context, offerId types.ID) error {
	receipt, err := manager.contract.Order(ctx, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseOfferPresentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer presented.", logger.Attrs{"offer-id": evt.OfferId.Hex()})
	return err
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (manager *exchangeManager) Cancel(ctx context.Context, offerId types.ID) error {
	receipt, err := manager.contract.Cancel(ctx, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseOfferCanceledFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer cancelled.", logger.Attrs{"offer-id": evt.OfferId.Hex()})
	return err
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (manager *exchangeManager) Settle(ctx context.Context, offerId types.ID) error {
	receipt, err := manager.contract.Settle(ctx, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseOfferSettledFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer settled", logger.Attrs{"offer-id": evt.OfferId.Hex()})
	return nil
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (manager *exchangeManager) Reject(ctx context.Context, offerId types.ID) error {
	receipt, err := manager.contract.Reject(ctx, offerId)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseOfferRejectedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Offer rejected", logger.Attrs{"offer-id": evt.OfferId.Hex()})
	return nil
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (manager *exchangeManager) GetOffer(offerId types.ID) (types.Offer, error) {
	return manager.contract.GetOffer(offerId)
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (manager *exchangeManager) GetOfferMembers(offerId types.ID) (ethCommon.Address, ethCommon.Address, error) {
	return manager.contract.GetOfferMembers(offerId)
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (manager *exchangeManager) OfferExists(offerId types.ID) (bool, error) {
	return manager.contract.OfferExists(offerId)
}

// FilterEscrowExecutionFailed is a free log retrieval operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (manager exchangeManager) FilterEscrowExecutionFailed(opts *bind.FilterOpts) (*ExchangeEscrowExecutionFailedIterator, error) {
	return manager.contract.FilterEscrowExecutionFailed(opts)
}

// WatchEscrowExecutionFailed is a free log subscription operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (manager exchangeManager) WatchEscrowExecutionFailed(opts *bind.WatchOpts, sink chan<- *ExchangeEscrowExecutionFailed) (event.Subscription, error) {
	return manager.contract.WatchEscrowExecutionFailed(opts, sink)
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (manager exchangeManager) FilterOfferCanceled(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferCanceledIterator, error) {
	return manager.contract.FilterOfferCanceled(opts, offerId)
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (manager exchangeManager) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferCanceled, offerId []types.ID) (event.Subscription, error) {
	return manager.contract.WatchOfferCanceled(opts, sink, offerId)
}

// FilterOfferPrepared is a free log retrieval operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (manager exchangeManager) FilterOfferPrepared(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPreparedIterator, error) {
	return manager.contract.FilterOfferPrepared(opts, offerId)
}

// WatchOfferPrepared is a free log subscription operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (manager exchangeManager) WatchOfferPrepared(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPrepared, offerId []types.ID) (event.Subscription, error) {
	return manager.contract.WatchOfferPrepared(opts, sink, offerId)
}

// FilterOfferPresented is a free log retrieval operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (manager exchangeManager) FilterOfferPresented(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPresentedIterator, error) {
	return manager.contract.FilterOfferPresented(opts, offerId)
}

// WatchOfferPresented is a free log subscription operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (manager exchangeManager) WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPresented, offerId []types.ID) (event.Subscription, error) {
	return manager.contract.WatchOfferPresented(opts, sink, offerId)
}

// FilterOfferReceipt is a free log retrieval operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (manager exchangeManager) FilterOfferReceipt(opts *bind.FilterOpts, offerId []types.ID, consumer []ethCommon.Address) (*ExchangeOfferReceiptIterator, error) {
	return manager.contract.FilterOfferReceipt(opts, offerId, consumer)
}

// WatchOfferReceipt is a free log subscription operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (manager exchangeManager) WatchOfferReceipt(opts *bind.WatchOpts, sink chan<- *ExchangeOfferReceipt, offerId []types.ID, consumer []ethCommon.Address) (event.Subscription, error) {
	return manager.contract.WatchOfferReceipt(opts, sink, offerId, consumer)
}

// FilterOfferRejected is a free log retrieval operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (manager exchangeManager) FilterOfferRejected(opts *bind.FilterOpts, offerId []types.ID, consumer []ethCommon.Address) (*ExchangeOfferRejectedIterator, error) {
	return manager.contract.FilterOfferRejected(opts, offerId, consumer)
}

// WatchOfferRejected is a free log subscription operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (manager exchangeManager) WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *ExchangeOfferRejected, offerId []types.ID, consumer []ethCommon.Address) (event.Subscription, error) {
	return manager.contract.WatchOfferRejected(opts, sink, offerId, consumer)
}

// FilterOfferSettled is a free log retrieval operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (manager exchangeManager) FilterOfferSettled(opts *bind.FilterOpts, offerId []types.ID, consumer []ethCommon.Address) (*ExchangeOfferSettledIterator, error) {
	return manager.contract.FilterOfferSettled(opts, offerId, consumer)
}

// WatchOfferSettled is a free log subscription operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (manager exchangeManager) WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferSettled, offerId []types.ID, consumer []ethCommon.Address) (event.Subscription, error) {
	return manager.contract.WatchOfferSettled(opts, sink, offerId, consumer)
}
