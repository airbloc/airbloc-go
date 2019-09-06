// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
	"math/big"

	types "github.com/airbloc/airbloc-go/shared/types"
	platform "github.com/klaytn/klaytn"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"inputs\":[{\"name\":\"appReg\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"providerAppName\",\"type\":\"string\"}],\"name\":\"OfferPrepared\",\"signature\":\"0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"providerAppName\",\"type\":\"string\"}],\"name\":\"OfferPresented\",\"signature\":\"0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"providerAppName\",\"type\":\"string\"}],\"name\":\"OfferCanceled\",\"signature\":\"0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"OfferSettled\",\"signature\":\"0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"providerAppName\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"OfferReceipt\",\"signature\":\"0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"offerId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"OfferRejected\",\"signature\":\"0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"EscrowExecutionFailed\",\"signature\":\"0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\"},{\"name\":\"escrowSign\",\"type\":\"bytes4\"},{\"name\":\"escrowArgs\",\"type\":\"bytes\"},{\"name\":\"dataIds\",\"type\":\"bytes20[]\"}],\"name\":\"prepare\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x77e61c33\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"},{\"name\":\"dataIds\",\"type\":\"bytes20[]\"}],\"name\":\"addDataIds\",\"outputs\":[],\"payable\":false,\"signature\":\"0x367a9005\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"signature\":\"0x0cf833fb\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"signature\":\"0xb2d9ba39\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"signature\":\"0xa60d9b5f\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"signature\":\"0x6622e153\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"offerExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xc4a03da9\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"getOffer\",\"outputs\":[{\"components\":[{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"dataIds\",\"type\":\"bytes20[]\"},{\"name\":\"at\",\"type\":\"uint256\"},{\"name\":\"until\",\"type\":\"uint256\"},{\"components\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"sign\",\"type\":\"bytes4\"},{\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"escrow\",\"type\":\"tuple\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x107f04b4\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"getOfferMembers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"0x72dfa465\",\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int

	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// Address is getter method of Exchange.address
func (_Exchange *Exchange) Address() common.Address {
	return _Exchange.address
}

// TxHash is getter method of Exchange.txHash
func (_Exchange *Exchange) TxHash() common.Hash {
	return _Exchange.txHash
}

// CreatedAt is getter method of Exchange.createdAt
func (_Exchange *Exchange) CreatedAt() *big.Int {
	return _Exchange.createdAt
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns(types.Offer)
func (_Exchange *ExchangeCaller) GetOffer(opts *bind.CallOpts, offerId types.ID) (types.Offer, error) {
	var (
		ret0 = new(types.Offer)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getOffer", offerId)
	return *ret0, err
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns(types.Offer)
func (_Exchange *ExchangeSession) GetOffer(offerId types.ID) (types.Offer, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, offerId)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns(types.Offer)
func (_Exchange *ExchangeCallerSession) GetOffer(offerId types.ID) (types.Offer, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, offerId)
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (_Exchange *ExchangeCaller) GetOfferMembers(opts *bind.CallOpts, offerId types.ID) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Exchange.contract.Call(opts, out, "getOfferMembers", offerId)
	return *ret0, *ret1, err
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (_Exchange *ExchangeSession) GetOfferMembers(offerId types.ID) (common.Address, common.Address, error) {
	return _Exchange.Contract.GetOfferMembers(&_Exchange.CallOpts, offerId)
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (_Exchange *ExchangeCallerSession) GetOfferMembers(offerId types.ID) (common.Address, common.Address, error) {
	return _Exchange.Contract.GetOfferMembers(&_Exchange.CallOpts, offerId)
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (_Exchange *ExchangeCaller) OfferExists(opts *bind.CallOpts, offerId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "offerExists", offerId)
	return *ret0, err
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (_Exchange *ExchangeSession) OfferExists(offerId types.ID) (bool, error) {
	return _Exchange.Contract.OfferExists(&_Exchange.CallOpts, offerId)
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (_Exchange *ExchangeCallerSession) OfferExists(offerId types.ID) (bool, error) {
	return _Exchange.Contract.OfferExists(&_Exchange.CallOpts, offerId)
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (_Exchange *ExchangeTransactor) AddDataIds(opts *bind.TransactOpts, offerId types.ID, dataIds []types.DataId) (*chainTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addDataIds", offerId, dataIds)
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (_Exchange *ExchangeSession) AddDataIds(offerId types.ID, dataIds []types.DataId) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.AddDataIds(&_Exchange.TransactOpts, offerId, dataIds)
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (_Exchange *ExchangeTransactorSession) AddDataIds(offerId types.ID, dataIds []types.DataId) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.AddDataIds(&_Exchange.TransactOpts, offerId, dataIds)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactor) Cancel(opts *bind.TransactOpts, offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancel", offerId)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (_Exchange *ExchangeSession) Cancel(offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, offerId)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactorSession) Cancel(offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactor) Order(opts *bind.TransactOpts, offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "order", offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (_Exchange *ExchangeSession) Order(offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactorSession) Order(offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, offerId)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (_Exchange *ExchangeTransactor) Prepare(opts *bind.TransactOpts, provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*chainTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "prepare", provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (_Exchange *ExchangeSession) Prepare(provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Prepare(&_Exchange.TransactOpts, provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (_Exchange *ExchangeTransactorSession) Prepare(provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Prepare(&_Exchange.TransactOpts, provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactor) Reject(opts *bind.TransactOpts, offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "reject", offerId)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (_Exchange *ExchangeSession) Reject(offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Reject(&_Exchange.TransactOpts, offerId)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactorSession) Reject(offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Reject(&_Exchange.TransactOpts, offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactor) Settle(opts *bind.TransactOpts, offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "settle", offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (_Exchange *ExchangeSession) Settle(offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Settle(&_Exchange.TransactOpts, offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactorSession) Settle(offerId types.ID) (*chainTypes.Transaction, error) {
	return _Exchange.Contract.Settle(&_Exchange.TransactOpts, offerId)
}

// ExchangeEscrowExecutionFailedIterator is returned from FilterEscrowExecutionFailed and is used to iterate over the raw logs and unpacked data for EscrowExecutionFailed events raised by the Exchange contract.
type ExchangeEscrowExecutionFailedIterator struct {
	Event *ExchangeEscrowExecutionFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

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
			it.Event = new(ExchangeEscrowExecutionFailed)
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
		it.Event = new(ExchangeEscrowExecutionFailed)
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
func (_Exchange *ExchangeFilterer) FilterEscrowExecutionFailed(opts *bind.FilterOpts) (*ExchangeEscrowExecutionFailedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EscrowExecutionFailed")
	if err != nil {
		return nil, err
	}
	return &ExchangeEscrowExecutionFailedIterator{contract: _Exchange.contract, event: "EscrowExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchEscrowExecutionFailed is a free log subscription operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *ExchangeFilterer) WatchEscrowExecutionFailed(opts *bind.WatchOpts, sink chan<- *ExchangeEscrowExecutionFailed) (event.Subscription, error) {

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
				event := new(ExchangeEscrowExecutionFailed)
				if err := _Exchange.contract.UnpackLog(event, "EscrowExecutionFailed", log); err != nil {
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

// ParseEscrowExecutionFailed is a log parse operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *ExchangeFilterer) ParseEscrowExecutionFailed(log chainTypes.Log) (*ExchangeEscrowExecutionFailed, error) {
	event := new(ExchangeEscrowExecutionFailed)
	if err := _Exchange.contract.UnpackLog(event, "EscrowExecutionFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterEscrowExecutionFailed parses the event from given transaction receipt.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *ExchangeFilterer) ParseEscrowExecutionFailedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeEscrowExecutionFailed, error) {
	var events []*ExchangeEscrowExecutionFailed
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7") {
			event, err := _Exchange.ParseEscrowExecutionFailed(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("EscrowExecutionFailed event not found")
	}
	return events, nil
}

// ExchangeOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the Exchange contract.
type ExchangeOfferCanceledIterator struct {
	Event *ExchangeOfferCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

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
			it.Event = new(ExchangeOfferCanceled)
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
		it.Event = new(ExchangeOfferCanceled)
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
func (_Exchange *ExchangeFilterer) FilterOfferCanceled(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferCanceledIterator, error) {

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
func (_Exchange *ExchangeFilterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferCanceled, offerId []types.ID) (event.Subscription, error) {

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
				event := new(ExchangeOfferCanceled)
				if err := _Exchange.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
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

// ParseOfferCanceled is a log parse operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferCanceled(log chainTypes.Log) (*ExchangeOfferCanceled, error) {
	event := new(ExchangeOfferCanceled)
	if err := _Exchange.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterOfferCanceled parses the event from given transaction receipt.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferCanceledFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferCanceled, error) {
	var events []*ExchangeOfferCanceled
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a") {
			event, err := _Exchange.ParseOfferCanceled(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("OfferCanceled event not found")
	}
	return events, nil
}

// ExchangeOfferPreparedIterator is returned from FilterOfferPrepared and is used to iterate over the raw logs and unpacked data for OfferPrepared events raised by the Exchange contract.
type ExchangeOfferPreparedIterator struct {
	Event *ExchangeOfferPrepared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

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
			it.Event = new(ExchangeOfferPrepared)
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
		it.Event = new(ExchangeOfferPrepared)
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
func (_Exchange *ExchangeFilterer) FilterOfferPrepared(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPreparedIterator, error) {

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
func (_Exchange *ExchangeFilterer) WatchOfferPrepared(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPrepared, offerId []types.ID) (event.Subscription, error) {

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
				event := new(ExchangeOfferPrepared)
				if err := _Exchange.contract.UnpackLog(event, "OfferPrepared", log); err != nil {
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

// ParseOfferPrepared is a log parse operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferPrepared(log chainTypes.Log) (*ExchangeOfferPrepared, error) {
	event := new(ExchangeOfferPrepared)
	if err := _Exchange.contract.UnpackLog(event, "OfferPrepared", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterOfferPrepared parses the event from given transaction receipt.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferPreparedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferPrepared, error) {
	var events []*ExchangeOfferPrepared
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5") {
			event, err := _Exchange.ParseOfferPrepared(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("OfferPrepared event not found")
	}
	return events, nil
}

// ExchangeOfferPresentedIterator is returned from FilterOfferPresented and is used to iterate over the raw logs and unpacked data for OfferPresented events raised by the Exchange contract.
type ExchangeOfferPresentedIterator struct {
	Event *ExchangeOfferPresented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

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
			it.Event = new(ExchangeOfferPresented)
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
		it.Event = new(ExchangeOfferPresented)
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
func (_Exchange *ExchangeFilterer) FilterOfferPresented(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPresentedIterator, error) {

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
func (_Exchange *ExchangeFilterer) WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPresented, offerId []types.ID) (event.Subscription, error) {

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
				event := new(ExchangeOfferPresented)
				if err := _Exchange.contract.UnpackLog(event, "OfferPresented", log); err != nil {
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

// ParseOfferPresented is a log parse operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferPresented(log chainTypes.Log) (*ExchangeOfferPresented, error) {
	event := new(ExchangeOfferPresented)
	if err := _Exchange.contract.UnpackLog(event, "OfferPresented", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterOfferPresented parses the event from given transaction receipt.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferPresentedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferPresented, error) {
	var events []*ExchangeOfferPresented
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa") {
			event, err := _Exchange.ParseOfferPresented(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("OfferPresented event not found")
	}
	return events, nil
}

// ExchangeOfferReceiptIterator is returned from FilterOfferReceipt and is used to iterate over the raw logs and unpacked data for OfferReceipt events raised by the Exchange contract.
type ExchangeOfferReceiptIterator struct {
	Event *ExchangeOfferReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

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
			it.Event = new(ExchangeOfferReceipt)
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
		it.Event = new(ExchangeOfferReceipt)
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
func (_Exchange *ExchangeFilterer) FilterOfferReceipt(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferReceiptIterator, error) {

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
func (_Exchange *ExchangeFilterer) WatchOfferReceipt(opts *bind.WatchOpts, sink chan<- *ExchangeOfferReceipt, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

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
				event := new(ExchangeOfferReceipt)
				if err := _Exchange.contract.UnpackLog(event, "OfferReceipt", log); err != nil {
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

// ParseOfferReceipt is a log parse operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *ExchangeFilterer) ParseOfferReceipt(log chainTypes.Log) (*ExchangeOfferReceipt, error) {
	event := new(ExchangeOfferReceipt)
	if err := _Exchange.contract.UnpackLog(event, "OfferReceipt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterOfferReceipt parses the event from given transaction receipt.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *ExchangeFilterer) ParseOfferReceiptFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferReceipt, error) {
	var events []*ExchangeOfferReceipt
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa") {
			event, err := _Exchange.ParseOfferReceipt(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("OfferReceipt event not found")
	}
	return events, nil
}

// ExchangeOfferRejectedIterator is returned from FilterOfferRejected and is used to iterate over the raw logs and unpacked data for OfferRejected events raised by the Exchange contract.
type ExchangeOfferRejectedIterator struct {
	Event *ExchangeOfferRejected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

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
			it.Event = new(ExchangeOfferRejected)
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
		it.Event = new(ExchangeOfferRejected)
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
func (_Exchange *ExchangeFilterer) FilterOfferRejected(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferRejectedIterator, error) {

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
func (_Exchange *ExchangeFilterer) WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *ExchangeOfferRejected, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

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
				event := new(ExchangeOfferRejected)
				if err := _Exchange.contract.UnpackLog(event, "OfferRejected", log); err != nil {
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

// ParseOfferRejected is a log parse operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) ParseOfferRejected(log chainTypes.Log) (*ExchangeOfferRejected, error) {
	event := new(ExchangeOfferRejected)
	if err := _Exchange.contract.UnpackLog(event, "OfferRejected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterOfferRejected parses the event from given transaction receipt.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) ParseOfferRejectedFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferRejected, error) {
	var events []*ExchangeOfferRejected
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843") {
			event, err := _Exchange.ParseOfferRejected(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("OfferRejected event not found")
	}
	return events, nil
}

// ExchangeOfferSettledIterator is returned from FilterOfferSettled and is used to iterate over the raw logs and unpacked data for OfferSettled events raised by the Exchange contract.
type ExchangeOfferSettledIterator struct {
	Event *ExchangeOfferSettled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

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
			it.Event = new(ExchangeOfferSettled)
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
		it.Event = new(ExchangeOfferSettled)
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
func (_Exchange *ExchangeFilterer) FilterOfferSettled(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferSettledIterator, error) {

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
func (_Exchange *ExchangeFilterer) WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferSettled, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

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
				event := new(ExchangeOfferSettled)
				if err := _Exchange.contract.UnpackLog(event, "OfferSettled", log); err != nil {
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

// ParseOfferSettled is a log parse operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) ParseOfferSettled(log chainTypes.Log) (*ExchangeOfferSettled, error) {
	event := new(ExchangeOfferSettled)
	if err := _Exchange.contract.UnpackLog(event, "OfferSettled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FilterOfferSettled parses the event from given transaction receipt.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) ParseOfferSettledFromReceipt(receipt *chainTypes.Receipt) ([]*ExchangeOfferSettled, error) {
	var events []*ExchangeOfferSettled
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e") {
			event, err := _Exchange.ParseOfferSettled(*log)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("OfferSettled event not found")
	}
	return events, nil
}
