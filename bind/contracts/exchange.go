package contracts

import (
	"context"
	"errors"
	"math/big"
	"strings"

	ablbind "github.com/airbloc/airbloc-go/bind"
	types "github.com/airbloc/airbloc-go/bind/types"
	platform "github.com/klaytn/klaytn"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// ExchangeABI is the input ABI used to generate the binding from.
const (
	ExchangeAddress   = "0x5A7E1B30627D09901981168EA001A1d9c80F8fcA"
	ExchangeTxHash    = "0x5adc6c14bb6dcafb3e4095db425cc8ab22a3f237132f7ea5cac50717b2daab89"
	ExchangeCreatedAt = "0x000000000000000000000000000000000000000000000000000000000063b9f6"
	ExchangeABI       = "[{\"inputs\":[{\"name\":\"appReg\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"providerAppName\",\"type\":\"string\"}],\"name\":\"OfferPrepared\",\"signature\":\"0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"providerAppName\",\"type\":\"string\"}],\"name\":\"OfferPresented\",\"signature\":\"0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"providerAppName\",\"type\":\"string\"}],\"name\":\"OfferCanceled\",\"signature\":\"0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"OfferSettled\",\"signature\":\"0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"providerAppName\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"OfferReceipt\",\"signature\":\"0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"OfferRejected\",\"signature\":\"0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"EscrowExecutionFailed\",\"signature\":\"0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\"},{\"name\":\"escrowSign\",\"type\":\"bytes4\"},{\"name\":\"escrowArgs\",\"type\":\"bytes\"},{\"name\":\"dataIds\",\"type\":\"bytes20[]\"}],\"name\":\"prepare\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x77e61c33\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"},{\"name\":\"dataIds\",\"type\":\"bytes20[]\"}],\"name\":\"addDataIds\",\"outputs\":[],\"payable\":false,\"signature\":\"0x367a9005\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"signature\":\"0x0cf833fb\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"signature\":\"0xb2d9ba39\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"signature\":\"0xa60d9b5f\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"signature\":\"0x6622e153\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"offerExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xc4a03da9\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"getOffer\",\"outputs\":[{\"components\":[{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"dataIds\",\"type\":\"bytes20[]\"},{\"name\":\"at\",\"type\":\"uint256\"},{\"name\":\"until\",\"type\":\"uint256\"},{\"components\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"sign\",\"type\":\"bytes4\"},{\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"escrow\",\"type\":\"tuple\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x107f04b4\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"getOfferMembers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"0x72dfa465\",\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller interface {
	GetOffer(
		ctx context.Context,
		offerId types.ID,
	) (
		types.Offer,
		error,
	)
	GetOfferMembers(
		ctx context.Context,
		offerId types.ID,
	) (
		common.Address,
		common.Address,
		error,
	)
	OfferExists(
		ctx context.Context,
		offerId types.ID,
	) (
		bool,
		error,
	)
}

type exchangeCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns(types.Offer)
func (_Exchange *exchangeCaller) GetOffer(ctx context.Context, offerId types.ID) (types.Offer, error) {
	var (
		ret0 = new(types.Offer)
	)
	out := ret0

	err := _Exchange.contract.Call(&bind.CallOpts{Context: ctx}, out, "getOffer", offerId)
	return *ret0, err
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (_Exchange *exchangeCaller) GetOfferMembers(ctx context.Context, offerId types.ID) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}

	err := _Exchange.contract.Call(&bind.CallOpts{Context: ctx}, out, "getOfferMembers", offerId)
	return *ret0, *ret1, err
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (_Exchange *exchangeCaller) OfferExists(ctx context.Context, offerId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Exchange.contract.Call(&bind.CallOpts{Context: ctx}, out, "offerExists", offerId)
	return *ret0, err
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor interface {
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

type exchangeTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
	backend  ablbind.ContractBackend
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (_Exchange *exchangeTransactor) AddDataIds(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	offerId types.ID,
	dataIds []types.DataId,
) (*chainTypes.Receipt, error) {
	tx, err := _Exchange.contract.Transact(_Exchange.backend.Transactor(ctx, opts), "addDataIds", offerId, dataIds)
	if err != nil {
		return nil, err
	}
	return _Exchange.backend.WaitMined(ctx, tx)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (_Exchange *exchangeTransactor) Cancel(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := _Exchange.contract.Transact(_Exchange.backend.Transactor(ctx, opts), "cancel", offerId)
	if err != nil {
		return nil, err
	}
	return _Exchange.backend.WaitMined(ctx, tx)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (_Exchange *exchangeTransactor) Order(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := _Exchange.contract.Transact(_Exchange.backend.Transactor(ctx, opts), "order", offerId)
	if err != nil {
		return nil, err
	}
	return _Exchange.backend.WaitMined(ctx, tx)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (_Exchange *exchangeTransactor) Prepare(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	provider string,
	consumer common.Address,
	escrow common.Address,
	escrowSign [4]byte,
	escrowArgs []byte,
	dataIds []types.DataId,
) (*chainTypes.Receipt, error) {
	tx, err := _Exchange.contract.Transact(_Exchange.backend.Transactor(ctx, opts), "prepare", provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
	if err != nil {
		return nil, err
	}
	return _Exchange.backend.WaitMined(ctx, tx)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (_Exchange *exchangeTransactor) Reject(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := _Exchange.contract.Transact(_Exchange.backend.Transactor(ctx, opts), "reject", offerId)
	if err != nil {
		return nil, err
	}
	return _Exchange.backend.WaitMined(ctx, tx)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (_Exchange *exchangeTransactor) Settle(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	offerId types.ID,
) (*chainTypes.Receipt, error) {
	tx, err := _Exchange.contract.Transact(_Exchange.backend.Transactor(ctx, opts), "settle", offerId)
	if err != nil {
		return nil, err
	}
	return _Exchange.backend.WaitMined(ctx, tx)
}

type ExchangeEvents interface {
	ExchangeEventFilterer
	ExchangeEventParser
	ExchangeEventWatcher
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeEventFilterer interface {
	// Filterer
	FilterEscrowExecutionFailed(
		opts *bind.FilterOpts,

	) (ablbind.EventIterator, error)

	// Filterer
	FilterOfferCanceled(
		opts *bind.FilterOpts,
		offerId []types.ID,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterOfferPrepared(
		opts *bind.FilterOpts,
		offerId []types.ID,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterOfferPresented(
		opts *bind.FilterOpts,
		offerId []types.ID,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterOfferReceipt(
		opts *bind.FilterOpts,
		offerId []types.ID, consumer []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterOfferRejected(
		opts *bind.FilterOpts,
		offerId []types.ID, consumer []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterOfferSettled(
		opts *bind.FilterOpts,
		offerId []types.ID, consumer []common.Address,
	) (ablbind.EventIterator, error)
}

type ExchangeEventParser interface {
	// Parser
	ParseEscrowExecutionFailed(log chainTypes.Log) (*ExchangeEscrowExecutionFailed, error)
	ParseEscrowExecutionFailedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeEscrowExecutionFailed, error)

	// Parser
	ParseOfferCanceled(log chainTypes.Log) (*ExchangeOfferCanceled, error)
	ParseOfferCanceledFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferCanceled, error)

	// Parser
	ParseOfferPrepared(log chainTypes.Log) (*ExchangeOfferPrepared, error)
	ParseOfferPreparedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferPrepared, error)

	// Parser
	ParseOfferPresented(log chainTypes.Log) (*ExchangeOfferPresented, error)
	ParseOfferPresentedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferPresented, error)

	// Parser
	ParseOfferReceipt(log chainTypes.Log) (*ExchangeOfferReceipt, error)
	ParseOfferReceiptFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferReceipt, error)

	// Parser
	ParseOfferRejected(log chainTypes.Log) (*ExchangeOfferRejected, error)
	ParseOfferRejectedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferRejected, error)

	// Parser
	ParseOfferSettled(log chainTypes.Log) (*ExchangeOfferSettled, error)
	ParseOfferSettledFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferSettled, error)
}

type ExchangeEventWatcher interface {
	// Watcher
	WatchEscrowExecutionFailed(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeEscrowExecutionFailed,

	) (event.Subscription, error)

	// Watcher
	WatchOfferCanceled(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferCanceled,
		offerId []types.ID,
	) (event.Subscription, error)

	// Watcher
	WatchOfferPrepared(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferPrepared,
		offerId []types.ID,
	) (event.Subscription, error)

	// Watcher
	WatchOfferPresented(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferPresented,
		offerId []types.ID,
	) (event.Subscription, error)

	// Watcher
	WatchOfferReceipt(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferReceipt,
		offerId []types.ID, consumer []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchOfferRejected(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferRejected,
		offerId []types.ID, consumer []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchOfferSettled(
		opts *bind.WatchOpts,
		sink chan<- *ExchangeOfferSettled,
		offerId []types.ID, consumer []common.Address,
	) (event.Subscription, error)
}

type exchangeEvents struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeEscrowExecutionFailedIterator is returned from FilterEscrowExecutionFailed and is used to iterate over the raw logs and unpacked data for EscrowExecutionFailed events raised by the Exchange contract.
type ExchangeEscrowExecutionFailedIterator struct {
	Evt *ExchangeEscrowExecutionFailed // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeEscrowExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ExchangeEscrowExecutionFailed)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(ExchangeEscrowExecutionFailed)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeEscrowExecutionFailedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeEscrowExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEscrowExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEscrowExecutionFailed represents a EscrowExecutionFailed event raised by the Exchange contract.
type ExchangeEscrowExecutionFailed struct {
	Reason []byte
	Raw    chainTypes.Log // Blockchain specific contextual infos
}

// FilterEscrowExecutionFailed is a free log retrieval operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *exchangeEvents) FilterEscrowExecutionFailed(opts *bind.FilterOpts) (ablbind.EventIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EscrowExecutionFailed")
	if err != nil {
		return nil, err
	}
	return &ExchangeEscrowExecutionFailedIterator{contract: _Exchange.contract, event: "EscrowExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchEscrowExecutionFailed is a free log subscription operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *exchangeEvents) WatchEscrowExecutionFailed(opts *bind.WatchOpts, sink chan<- *ExchangeEscrowExecutionFailed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EscrowExecutionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(ExchangeEscrowExecutionFailed)
				if err := _Exchange.contract.UnpackLog(evt, "EscrowExecutionFailed", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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

// ParseEscrowExecutionFailed is a log parse operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *exchangeEvents) ParseEscrowExecutionFailed(log chainTypes.Log) (*ExchangeEscrowExecutionFailed, error) {
	evt := new(ExchangeEscrowExecutionFailed)
	if err := _Exchange.contract.UnpackLog(evt, "EscrowExecutionFailed", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterEscrowExecutionFailed parses the event from given transaction receipt.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *exchangeEvents) ParseEscrowExecutionFailedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeEscrowExecutionFailed, error) {
	var evts []*ExchangeEscrowExecutionFailed
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7") {
			evt, err := _Exchange.ParseEscrowExecutionFailed(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("EscrowExecutionFailed event not found")
	}
	return evts, nil
}

// ExchangeOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the Exchange contract.
type ExchangeOfferCanceledIterator struct {
	Evt *ExchangeOfferCanceled // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ExchangeOfferCanceled)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(ExchangeOfferCanceled)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferCanceledIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferCanceled represents a OfferCanceled event raised by the Exchange contract.
type ExchangeOfferCanceled struct {
	OfferId         types.ID
	ProviderAppName string
	Raw             chainTypes.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) FilterOfferCanceled(opts *bind.FilterOpts, offerId []types.ID) (ablbind.EventIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferCanceled", offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferCanceledIterator{contract: _Exchange.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferCanceled, offerId []types.ID) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferCanceled", offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(ExchangeOfferCanceled)
				if err := _Exchange.contract.UnpackLog(evt, "OfferCanceled", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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

// ParseOfferCanceled is a log parse operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) ParseOfferCanceled(log chainTypes.Log) (*ExchangeOfferCanceled, error) {
	evt := new(ExchangeOfferCanceled)
	if err := _Exchange.contract.UnpackLog(evt, "OfferCanceled", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterOfferCanceled parses the event from given transaction receipt.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) ParseOfferCanceledFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferCanceled, error) {
	var evts []*ExchangeOfferCanceled
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a") {
			evt, err := _Exchange.ParseOfferCanceled(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("OfferCanceled event not found")
	}
	return evts, nil
}

// ExchangeOfferPreparedIterator is returned from FilterOfferPrepared and is used to iterate over the raw logs and unpacked data for OfferPrepared events raised by the Exchange contract.
type ExchangeOfferPreparedIterator struct {
	Evt *ExchangeOfferPrepared // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ExchangeOfferPrepared)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(ExchangeOfferPrepared)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferPreparedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferPrepared represents a OfferPrepared event raised by the Exchange contract.
type ExchangeOfferPrepared struct {
	OfferId         types.ID
	ProviderAppName string
	Raw             chainTypes.Log // Blockchain specific contextual infos
}

// FilterOfferPrepared is a free log retrieval operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) FilterOfferPrepared(opts *bind.FilterOpts, offerId []types.ID) (ablbind.EventIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferPrepared", offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferPreparedIterator{contract: _Exchange.contract, event: "OfferPrepared", logs: logs, sub: sub}, nil
}

// WatchOfferPrepared is a free log subscription operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) WatchOfferPrepared(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPrepared, offerId []types.ID) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferPrepared", offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(ExchangeOfferPrepared)
				if err := _Exchange.contract.UnpackLog(evt, "OfferPrepared", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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

// ParseOfferPrepared is a log parse operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) ParseOfferPrepared(log chainTypes.Log) (*ExchangeOfferPrepared, error) {
	evt := new(ExchangeOfferPrepared)
	if err := _Exchange.contract.UnpackLog(evt, "OfferPrepared", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterOfferPrepared parses the event from given transaction receipt.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) ParseOfferPreparedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferPrepared, error) {
	var evts []*ExchangeOfferPrepared
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5") {
			evt, err := _Exchange.ParseOfferPrepared(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("OfferPrepared event not found")
	}
	return evts, nil
}

// ExchangeOfferPresentedIterator is returned from FilterOfferPresented and is used to iterate over the raw logs and unpacked data for OfferPresented events raised by the Exchange contract.
type ExchangeOfferPresentedIterator struct {
	Evt *ExchangeOfferPresented // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferPresentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ExchangeOfferPresented)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(ExchangeOfferPresented)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferPresentedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferPresentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferPresentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferPresented represents a OfferPresented event raised by the Exchange contract.
type ExchangeOfferPresented struct {
	OfferId         types.ID
	ProviderAppName string
	Raw             chainTypes.Log // Blockchain specific contextual infos
}

// FilterOfferPresented is a free log retrieval operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) FilterOfferPresented(opts *bind.FilterOpts, offerId []types.ID) (ablbind.EventIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferPresented", offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferPresentedIterator{contract: _Exchange.contract, event: "OfferPresented", logs: logs, sub: sub}, nil
}

// WatchOfferPresented is a free log subscription operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPresented, offerId []types.ID) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferPresented", offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(ExchangeOfferPresented)
				if err := _Exchange.contract.UnpackLog(evt, "OfferPresented", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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

// ParseOfferPresented is a log parse operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) ParseOfferPresented(log chainTypes.Log) (*ExchangeOfferPresented, error) {
	evt := new(ExchangeOfferPresented)
	if err := _Exchange.contract.UnpackLog(evt, "OfferPresented", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterOfferPresented parses the event from given transaction receipt.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *exchangeEvents) ParseOfferPresentedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferPresented, error) {
	var evts []*ExchangeOfferPresented
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa") {
			evt, err := _Exchange.ParseOfferPresented(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("OfferPresented event not found")
	}
	return evts, nil
}

// ExchangeOfferReceiptIterator is returned from FilterOfferReceipt and is used to iterate over the raw logs and unpacked data for OfferReceipt events raised by the Exchange contract.
type ExchangeOfferReceiptIterator struct {
	Evt *ExchangeOfferReceipt // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ExchangeOfferReceipt)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(ExchangeOfferReceipt)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferReceiptIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferReceipt represents a OfferReceipt event raised by the Exchange contract.
type ExchangeOfferReceipt struct {
	OfferId         types.ID
	ProviderAppName string
	Consumer        common.Address
	Result          []byte
	Raw             chainTypes.Log // Blockchain specific contextual infos
}

// FilterOfferReceipt is a free log retrieval operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *exchangeEvents) FilterOfferReceipt(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (ablbind.EventIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferReceipt", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferReceiptIterator{contract: _Exchange.contract, event: "OfferReceipt", logs: logs, sub: sub}, nil
}

// WatchOfferReceipt is a free log subscription operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *exchangeEvents) WatchOfferReceipt(opts *bind.WatchOpts, sink chan<- *ExchangeOfferReceipt, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferReceipt", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(ExchangeOfferReceipt)
				if err := _Exchange.contract.UnpackLog(evt, "OfferReceipt", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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

// ParseOfferReceipt is a log parse operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *exchangeEvents) ParseOfferReceipt(log chainTypes.Log) (*ExchangeOfferReceipt, error) {
	evt := new(ExchangeOfferReceipt)
	if err := _Exchange.contract.UnpackLog(evt, "OfferReceipt", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterOfferReceipt parses the event from given transaction receipt.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *exchangeEvents) ParseOfferReceiptFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferReceipt, error) {
	var evts []*ExchangeOfferReceipt
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa") {
			evt, err := _Exchange.ParseOfferReceipt(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("OfferReceipt event not found")
	}
	return evts, nil
}

// ExchangeOfferRejectedIterator is returned from FilterOfferRejected and is used to iterate over the raw logs and unpacked data for OfferRejected events raised by the Exchange contract.
type ExchangeOfferRejectedIterator struct {
	Evt *ExchangeOfferRejected // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ExchangeOfferRejected)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(ExchangeOfferRejected)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferRejectedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferRejected represents a OfferRejected event raised by the Exchange contract.
type ExchangeOfferRejected struct {
	OfferId  types.ID
	Consumer common.Address
	Raw      chainTypes.Log // Blockchain specific contextual infos
}

// FilterOfferRejected is a free log retrieval operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *exchangeEvents) FilterOfferRejected(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (ablbind.EventIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferRejected", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferRejectedIterator{contract: _Exchange.contract, event: "OfferRejected", logs: logs, sub: sub}, nil
}

// WatchOfferRejected is a free log subscription operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *exchangeEvents) WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *ExchangeOfferRejected, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferRejected", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(ExchangeOfferRejected)
				if err := _Exchange.contract.UnpackLog(evt, "OfferRejected", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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

// ParseOfferRejected is a log parse operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *exchangeEvents) ParseOfferRejected(log chainTypes.Log) (*ExchangeOfferRejected, error) {
	evt := new(ExchangeOfferRejected)
	if err := _Exchange.contract.UnpackLog(evt, "OfferRejected", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterOfferRejected parses the event from given transaction receipt.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *exchangeEvents) ParseOfferRejectedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferRejected, error) {
	var evts []*ExchangeOfferRejected
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843") {
			evt, err := _Exchange.ParseOfferRejected(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("OfferRejected event not found")
	}
	return evts, nil
}

// ExchangeOfferSettledIterator is returned from FilterOfferSettled and is used to iterate over the raw logs and unpacked data for OfferSettled events raised by the Exchange contract.
type ExchangeOfferSettledIterator struct {
	Evt *ExchangeOfferSettled // Event containing the contract specifics and raw log

	contract *ablbind.BoundContract // Generic contract to use for unpacking event data
	event    string                 // Event name to use for unpacking event data

	logs chan chainTypes.Log   // Log channel receiving the found contract events
	sub  platform.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ExchangeOfferSettled)
			if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Evt.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Evt = new(ExchangeOfferSettled)
		if err := it.contract.UnpackLog(it.Evt, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Evt.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferSettledIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferSettled represents a OfferSettled event raised by the Exchange contract.
type ExchangeOfferSettled struct {
	OfferId  types.ID
	Consumer common.Address
	Raw      chainTypes.Log // Blockchain specific contextual infos
}

// FilterOfferSettled is a free log retrieval operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *exchangeEvents) FilterOfferSettled(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (ablbind.EventIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferSettled", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferSettledIterator{contract: _Exchange.contract, event: "OfferSettled", logs: logs, sub: sub}, nil
}

// WatchOfferSettled is a free log subscription operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *exchangeEvents) WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferSettled, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferSettled", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(ExchangeOfferSettled)
				if err := _Exchange.contract.UnpackLog(evt, "OfferSettled", log); err != nil {
					return err
				}
				evt.Raw = log

				select {
				case sink <- evt:
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

// ParseOfferSettled is a log parse operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *exchangeEvents) ParseOfferSettled(log chainTypes.Log) (*ExchangeOfferSettled, error) {
	evt := new(ExchangeOfferSettled)
	if err := _Exchange.contract.UnpackLog(evt, "OfferSettled", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterOfferSettled parses the event from given transaction receipt.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *exchangeEvents) ParseOfferSettledFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferSettled, error) {
	var evts []*ExchangeOfferSettled
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e") {
			evt, err := _Exchange.ParseOfferSettled(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("OfferSettled event not found")
	}
	return evts, nil
}

// Manager is contract wrapper struct
type ExchangeContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	ExchangeCaller
	ExchangeTransactor
	ExchangeEvents
}

func NewExchangeContract(backend ablbind.ContractBackend) (*ExchangeContract, error) {
	deployment, exist := backend.Deployment("Exchange")
	if !exist {
		evmABI, err := abi.JSON(strings.NewReader(ExchangeABI))
		if err != nil {
			return nil, err
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(ExchangeAddress),
			common.HexToHash(ExchangeTxHash),
			new(big.Int).SetBytes(common.HexToHash(ExchangeCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, "Exchange", backend)

	contract := &ExchangeContract{
		Deployment: deployment,
		client:     backend,

		ExchangeCaller:     &exchangeCaller{base},
		ExchangeTransactor: &exchangeTransactor{base, backend},
		ExchangeEvents:     &exchangeEvents{base},
	}

	return contract, nil
}
