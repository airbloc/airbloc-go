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

// ConsentsABI is the input ABI used to generate the binding from.
const (
	ConsentsAddress   = "0x1bd64C662f408437ABC0012F345e686506b01b11"
	ConsentsTxHash    = "0x9c6012dc1f1bb52875dfe4224eb0fec99678871438440c767ab5b99b622092ff"
	ConsentsCreatedAt = "0x0000000000000000000000000000000000000000000000000000000000d6d8ad"
	ConsentsABI       = "[{\"constant\":true,\"inputs\":[],\"name\":\"CONSENT_DATA_MAX_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"signature\":\"0x6d332de7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"userReg\",\"type\":\"address\"},{\"name\":\"appReg\",\"type\":\"address\"},{\"name\":\"controllerReg\",\"type\":\"address\"},{\"name\":\"dataTypeReg\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"appId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"dataType\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"Consented\",\"signature\":\"0x30eae66e1c1ac9fa8cbc19b0340774f536960bfdd4c8bdf58f89b0b7a50a47c6\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"components\":[{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"consentData\",\"type\":\"tuple\"}],\"name\":\"consent\",\"outputs\":[],\"payable\":false,\"signature\":\"0xf5146122\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"components\":[{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"consentData\",\"type\":\"tuple[]\"}],\"name\":\"consentMany\",\"outputs\":[],\"payable\":false,\"signature\":\"0xcb401b1b\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x50615985\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"},{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"uint8\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"isAllowedAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x7cdda67c\",\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

// ConsentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsentsCaller interface {
	IsAllowed(
		ctx context.Context,
		userId types.ID,
		appName string,
		action uint8,
		dataType string,
	) (
		bool,
		error,
	)
	IsAllowedAt(
		ctx context.Context,
		userId types.ID,
		appName string,
		action uint8,
		dataType string,
		blockNumber *big.Int,
	) (
		bool,
		error,
	)
}

type consentsCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// IsAllowed is a free data retrieval call binding the contract method 0x50615985.
//
// Solidity: function isAllowed(bytes8 userId, string appName, uint8 action, string dataType) constant returns(bool)
func (_Consents *consentsCaller) IsAllowed(ctx context.Context, userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Consents.contract.Call(&bind.CallOpts{Context: ctx}, out, "isAllowed", userId, appName, action, dataType)
	return *ret0, err
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x7cdda67c.
//
// Solidity: function isAllowedAt(bytes8 userId, string appName, uint8 action, string dataType, uint256 blockNumber) constant returns(bool)
func (_Consents *consentsCaller) IsAllowedAt(ctx context.Context, userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Consents.contract.Call(&bind.CallOpts{Context: ctx}, out, "isAllowedAt", userId, appName, action, dataType, blockNumber)
	return *ret0, err
}

// ConsentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsentsTransactor interface {
	Consent(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData types.ConsentData,
	) (*chainTypes.Receipt, error)
	ConsentMany(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData []types.ConsentData,
	) (*chainTypes.Receipt, error)
}

type consentsTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
	backend  ablbind.ContractBackend
}

// Consent is a paid mutator transaction binding the contract method 0xf5146122.
//
// Solidity: function consent(bytes8 userId, string appName, (uint8,string,bool) consentData) returns()
func (_Consents *consentsTransactor) Consent(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _Consents.contract.Transact(opts, "consent", userId, appName, consentData)
}

// ConsentMany is a paid mutator transaction binding the contract method 0xcb401b1b.
//
// Solidity: function consentMany(bytes8 userId, string appName, (uint8,string,bool)[] consentData) returns()
func (_Consents *consentsTransactor) ConsentMany(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData []types.ConsentData,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _Consents.contract.Transact(opts, "consentMany", userId, appName, consentData)
}

type ConsentsEvents interface {
	ConsentsEventFilterer
	ConsentsEventParser
	ConsentsEventWatcher
}

// ConsentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsentsEventFilterer interface {
	// Filterer
	FilterConsented(
		opts *bind.FilterOpts,
		action []uint8, userId []types.ID, appId []types.ID,
	) (ablbind.EventIterator, error)
}

type ConsentsEventParser interface {
	// Parser
	ParseConsented(log chainTypes.Log) (*ConsentsConsented, error)
	ParseConsentedFromReceipt(receipt *chainTypes.Receipt) ([]*ConsentsConsented, error)
}

type ConsentsEventWatcher interface {
	// Watcher
	WatchConsented(
		opts *bind.WatchOpts,
		sink chan<- *ConsentsConsented,
		action []uint8, userId []types.ID, appId []types.ID,
	) (event.Subscription, error)
}

type consentsEvents struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsentsConsentedIterator is returned from FilterConsented and is used to iterate over the raw logs and unpacked data for Consented events raised by the Consents contract.
type ConsentsConsentedIterator struct {
	Evt *ConsentsConsented // Event containing the contract specifics and raw log

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
func (it *ConsentsConsentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(ConsentsConsented)
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
		it.Evt = new(ConsentsConsented)
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
func (it *ConsentsConsentedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConsentsConsentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsentsConsentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsentsConsented represents a Consented event raised by the Consents contract.
type ConsentsConsented struct {
	Action   uint8
	UserId   types.ID
	AppId    types.ID
	AppName  string
	DataType string
	Allowed  bool
	Raw      chainTypes.Log // Blockchain specific contextual infos
}

// FilterConsented is a free log retrieval operation binding the contract event 0x30eae66e1c1ac9fa8cbc19b0340774f536960bfdd4c8bdf58f89b0b7a50a47c6.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes8 indexed appId, string appName, string dataType, bool allowed)
func (_Consents *consentsEvents) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appId []types.ID) (ablbind.EventIterator, error) {

	var actionRule []interface{}
	for _, actionItem := range action {
		actionRule = append(actionRule, actionItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _Consents.contract.FilterLogs(opts, "Consented", actionRule, userIdRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return &ConsentsConsentedIterator{contract: _Consents.contract, event: "Consented", logs: logs, sub: sub}, nil
}

// WatchConsented is a free log subscription operation binding the contract event 0x30eae66e1c1ac9fa8cbc19b0340774f536960bfdd4c8bdf58f89b0b7a50a47c6.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes8 indexed appId, string appName, string dataType, bool allowed)
func (_Consents *consentsEvents) WatchConsented(opts *bind.WatchOpts, sink chan<- *ConsentsConsented, action []uint8, userId []types.ID, appId []types.ID) (event.Subscription, error) {

	var actionRule []interface{}
	for _, actionItem := range action {
		actionRule = append(actionRule, actionItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _Consents.contract.WatchLogs(opts, "Consented", actionRule, userIdRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(ConsentsConsented)
				if err := _Consents.contract.UnpackLog(evt, "Consented", log); err != nil {
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

// ParseConsented is a log parse operation binding the contract event 0x30eae66e1c1ac9fa8cbc19b0340774f536960bfdd4c8bdf58f89b0b7a50a47c6.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes8 indexed appId, string appName, string dataType, bool allowed)
func (_Consents *consentsEvents) ParseConsented(log chainTypes.Log) (*ConsentsConsented, error) {
	evt := new(ConsentsConsented)
	if err := _Consents.contract.UnpackLog(evt, "Consented", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseConsentedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes8 indexed appId, string appName, string dataType, bool allowed)
func (_Consents *consentsEvents) ParseConsentedFromReceipt(receipt *chainTypes.Receipt) ([]*ConsentsConsented, error) {
	var evts []*ConsentsConsented
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x30eae66e1c1ac9fa8cbc19b0340774f536960bfdd4c8bdf58f89b0b7a50a47c6") {
			evt, err := _Consents.ParseConsented(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("Consented event not found")
	}
	return evts, nil
}

// Manager is contract wrapper struct
type ConsentsContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	ConsentsCaller
	ConsentsTransactor
	ConsentsEvents
}

func NewConsentsContract(backend ablbind.ContractBackend) (*ConsentsContract, error) {
	deployment, exist := backend.Deployment("Consents")
	if !exist {
		evmABI, err := abi.JSON(strings.NewReader(ConsentsABI))
		if err != nil {
			return nil, err
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(ConsentsAddress),
			common.HexToHash(ConsentsTxHash),
			new(big.Int).SetBytes(common.HexToHash(ConsentsCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, "Consents", backend)

	contract := &ConsentsContract{
		Deployment: deployment,
		client:     backend,

		ConsentsCaller:     &consentsCaller{base},
		ConsentsTransactor: &consentsTransactor{base, backend},
		ConsentsEvents:     &consentsEvents{base},
	}

	return contract, nil
}
