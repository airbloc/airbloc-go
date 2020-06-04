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

// AppRegistryABI is the input ABI used to generate the binding from.
const (
	AppRegistryAddress   = "0x5A9E5AD83d10e25a9764E76d1e191E9cc2491CD2"
	AppRegistryTxHash    = "0x312b15fea0e2ee850406f74b9d644e5ecd9d79329e0bcda855dad83e1330d2d1"
	AppRegistryCreatedAt = "0x0000000000000000000000000000000000000000000000000000000000d6d8a1"
	AppRegistryABI       = "[{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"roleName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"}],\"name\":\"revokeAction\",\"outputs\":[],\"payable\":false,\"signature\":\"0x10a789f8\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"createRole\",\"outputs\":[],\"payable\":false,\"signature\":\"0x617551d4\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"action\",\"type\":\"string\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x96c57715\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"signature\":\"0xac1465c6\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"subject\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"unbindRole\",\"outputs\":[],\"payable\":false,\"signature\":\"0xb31ec2af\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"subject\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"bindRole\",\"outputs\":[],\"payable\":false,\"signature\":\"0xc8526eed\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"roleName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"}],\"name\":\"grantAction\",\"outputs\":[],\"payable\":false,\"signature\":\"0xe602362d\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"Registration\",\"signature\":\"0xf5ef63b2a629619cbea83fa4979a52fa156e638837dd66bc4683dd7fbb8f4709\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"Unregistration\",\"signature\":\"0xfe4087afbde6adacfa58b7ae569b654c6445bdf8781edab97394ea74d2a5f85e\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"appName\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"AppOwnerTransferred\",\"signature\":\"0xd862515a389f54c6adf72b970712257c420c06e58830f73050bc0b9aca86f96d\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleCreation\",\"signature\":\"0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleRemoval\",\"signature\":\"0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleBound\",\"signature\":\"0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleUnbound\",\"signature\":\"0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"actionName\",\"type\":\"string\"}],\"name\":\"ActionGranted\",\"signature\":\"0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"actionName\",\"type\":\"string\"}],\"name\":\"ActionRevoked\",\"signature\":\"0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0xf2c298be\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"signature\":\"0x6598a1ae\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x693ec85e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"getId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0xbee51f3b\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x261a323e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xbde1eee7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appName\",\"type\":\"string\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAppOwner\",\"outputs\":[],\"payable\":false,\"signature\":\"0x1a9dff9f\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
)

// AppRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppRegistryCaller interface {
	Exists(
		ctx context.Context,
		appName string,
	) (
		bool,
		error,
	)
	Get(
		ctx context.Context,
		appName string,
	) (
		types.App,
		error,
	)
	GetId(
		ctx context.Context,
		appName string,
	) (
		types.ID,
		error,
	)
	IsOwner(
		ctx context.Context,
		appName string,
		owner common.Address,
	) (
		bool,
		error,
	)
}

type appRegistryCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (_AppRegistry *appRegistryCaller) Exists(ctx context.Context, appName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _AppRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "exists", appName)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns(types.App)
func (_AppRegistry *appRegistryCaller) Get(ctx context.Context, appName string) (types.App, error) {
	var (
		ret0 = new(types.App)
	)
	out := ret0

	err := _AppRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "get", appName)
	return *ret0, err
}

// GetId is a free data retrieval call binding the contract method 0xbee51f3b.
//
// Solidity: function getId(string appName) constant returns(bytes8)
func (_AppRegistry *appRegistryCaller) GetId(ctx context.Context, appName string) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := ret0

	err := _AppRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "getId", appName)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (_AppRegistry *appRegistryCaller) IsOwner(ctx context.Context, appName string, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _AppRegistry.contract.Call(&bind.CallOpts{Context: ctx}, out, "isOwner", appName, owner)
	return *ret0, err
}

// AppRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppRegistryTransactor interface {
	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
	) (*chainTypes.Receipt, error)
	TransferAppOwner(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
		newOwner common.Address,
	) (*chainTypes.Receipt, error)
	Unregister(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
	) (*chainTypes.Receipt, error)
}

type appRegistryTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
	backend  ablbind.ContractBackend
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns(bytes8)
func (_AppRegistry *appRegistryTransactor) Register(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _AppRegistry.contract.Transact(opts, "register", appName)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (_AppRegistry *appRegistryTransactor) TransferAppOwner(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
	newOwner common.Address,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _AppRegistry.contract.Transact(opts, "transferAppOwner", appName, newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (_AppRegistry *appRegistryTransactor) Unregister(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _AppRegistry.contract.Transact(opts, "unregister", appName)
}

type AppRegistryEvents interface {
	AppRegistryEventFilterer
	AppRegistryEventParser
	AppRegistryEventWatcher
}

// AppRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppRegistryEventFilterer interface {
	// Filterer
	FilterActionGranted(
		opts *bind.FilterOpts,
		resourceId []types.ID,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterActionRevoked(
		opts *bind.FilterOpts,
		resourceId []types.ID,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterAppOwnerTransferred(
		opts *bind.FilterOpts,
		appId []types.ID, oldOwner []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRegistration(
		opts *bind.FilterOpts,
		appId []types.ID,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRoleBound(
		opts *bind.FilterOpts,
		resourceId []types.ID, subject []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRoleCreation(
		opts *bind.FilterOpts,
		resourceId []types.ID,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRoleRemoval(
		opts *bind.FilterOpts,
		resourceId []types.ID,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRoleUnbound(
		opts *bind.FilterOpts,
		resourceId []types.ID, subject []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterUnregistration(
		opts *bind.FilterOpts,
		appId []types.ID,
	) (ablbind.EventIterator, error)
}

type AppRegistryEventParser interface {
	// Parser
	ParseActionGranted(log chainTypes.Log) (*AppRegistryActionGranted, error)
	ParseActionGrantedFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryActionGranted, error)

	// Parser
	ParseActionRevoked(log chainTypes.Log) (*AppRegistryActionRevoked, error)
	ParseActionRevokedFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryActionRevoked, error)

	// Parser
	ParseAppOwnerTransferred(log chainTypes.Log) (*AppRegistryAppOwnerTransferred, error)
	ParseAppOwnerTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryAppOwnerTransferred, error)

	// Parser
	ParseRegistration(log chainTypes.Log) (*AppRegistryRegistration, error)
	ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRegistration, error)

	// Parser
	ParseRoleBound(log chainTypes.Log) (*AppRegistryRoleBound, error)
	ParseRoleBoundFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRoleBound, error)

	// Parser
	ParseRoleCreation(log chainTypes.Log) (*AppRegistryRoleCreation, error)
	ParseRoleCreationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRoleCreation, error)

	// Parser
	ParseRoleRemoval(log chainTypes.Log) (*AppRegistryRoleRemoval, error)
	ParseRoleRemovalFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRoleRemoval, error)

	// Parser
	ParseRoleUnbound(log chainTypes.Log) (*AppRegistryRoleUnbound, error)
	ParseRoleUnboundFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRoleUnbound, error)

	// Parser
	ParseUnregistration(log chainTypes.Log) (*AppRegistryUnregistration, error)
	ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryUnregistration, error)
}

type AppRegistryEventWatcher interface {
	// Watcher
	WatchActionGranted(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryActionGranted,
		resourceId []types.ID,
	) (event.Subscription, error)

	// Watcher
	WatchActionRevoked(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryActionRevoked,
		resourceId []types.ID,
	) (event.Subscription, error)

	// Watcher
	WatchAppOwnerTransferred(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryAppOwnerTransferred,
		appId []types.ID, oldOwner []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchRegistration(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryRegistration,
		appId []types.ID,
	) (event.Subscription, error)

	// Watcher
	WatchRoleBound(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryRoleBound,
		resourceId []types.ID, subject []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchRoleCreation(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryRoleCreation,
		resourceId []types.ID,
	) (event.Subscription, error)

	// Watcher
	WatchRoleRemoval(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryRoleRemoval,
		resourceId []types.ID,
	) (event.Subscription, error)

	// Watcher
	WatchRoleUnbound(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryRoleUnbound,
		resourceId []types.ID, subject []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchUnregistration(
		opts *bind.WatchOpts,
		sink chan<- *AppRegistryUnregistration,
		appId []types.ID,
	) (event.Subscription, error)
}

type appRegistryEvents struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryActionGrantedIterator is returned from FilterActionGranted and is used to iterate over the raw logs and unpacked data for ActionGranted events raised by the AppRegistry contract.
type AppRegistryActionGrantedIterator struct {
	Evt *AppRegistryActionGranted // Event containing the contract specifics and raw log

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
func (it *AppRegistryActionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryActionGranted)
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
		it.Evt = new(AppRegistryActionGranted)
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
func (it *AppRegistryActionGrantedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryActionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryActionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryActionGranted represents a ActionGranted event raised by the AppRegistry contract.
type AppRegistryActionGranted struct {
	ResourceId types.ID
	RoleName   string
	ActionName string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterActionGranted is a free log retrieval operation binding the contract event 0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9.
//
// Solidity: event ActionGranted(bytes8 indexed resourceId, string roleName, string actionName)
func (_AppRegistry *appRegistryEvents) FilterActionGranted(opts *bind.FilterOpts, resourceId []types.ID) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "ActionGranted", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryActionGrantedIterator{contract: _AppRegistry.contract, event: "ActionGranted", logs: logs, sub: sub}, nil
}

// WatchActionGranted is a free log subscription operation binding the contract event 0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9.
//
// Solidity: event ActionGranted(bytes8 indexed resourceId, string roleName, string actionName)
func (_AppRegistry *appRegistryEvents) WatchActionGranted(opts *bind.WatchOpts, sink chan<- *AppRegistryActionGranted, resourceId []types.ID) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "ActionGranted", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryActionGranted)
				if err := _AppRegistry.contract.UnpackLog(evt, "ActionGranted", log); err != nil {
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

// ParseActionGranted is a log parse operation binding the contract event 0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9.
//
// Solidity: event ActionGranted(bytes8 indexed resourceId, string roleName, string actionName)
func (_AppRegistry *appRegistryEvents) ParseActionGranted(log chainTypes.Log) (*AppRegistryActionGranted, error) {
	evt := new(AppRegistryActionGranted)
	if err := _AppRegistry.contract.UnpackLog(evt, "ActionGranted", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseActionGrantedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event ActionGranted(bytes8 indexed resourceId, string roleName, string actionName)
func (_AppRegistry *appRegistryEvents) ParseActionGrantedFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryActionGranted, error) {
	var evts []*AppRegistryActionGranted
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9") {
			evt, err := _AppRegistry.ParseActionGranted(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("ActionGranted event not found")
	}
	return evts, nil
}

// AppRegistryActionRevokedIterator is returned from FilterActionRevoked and is used to iterate over the raw logs and unpacked data for ActionRevoked events raised by the AppRegistry contract.
type AppRegistryActionRevokedIterator struct {
	Evt *AppRegistryActionRevoked // Event containing the contract specifics and raw log

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
func (it *AppRegistryActionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryActionRevoked)
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
		it.Evt = new(AppRegistryActionRevoked)
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
func (it *AppRegistryActionRevokedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryActionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryActionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryActionRevoked represents a ActionRevoked event raised by the AppRegistry contract.
type AppRegistryActionRevoked struct {
	ResourceId types.ID
	RoleName   string
	ActionName string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterActionRevoked is a free log retrieval operation binding the contract event 0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430.
//
// Solidity: event ActionRevoked(bytes8 indexed resourceId, string roleName, string actionName)
func (_AppRegistry *appRegistryEvents) FilterActionRevoked(opts *bind.FilterOpts, resourceId []types.ID) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "ActionRevoked", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryActionRevokedIterator{contract: _AppRegistry.contract, event: "ActionRevoked", logs: logs, sub: sub}, nil
}

// WatchActionRevoked is a free log subscription operation binding the contract event 0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430.
//
// Solidity: event ActionRevoked(bytes8 indexed resourceId, string roleName, string actionName)
func (_AppRegistry *appRegistryEvents) WatchActionRevoked(opts *bind.WatchOpts, sink chan<- *AppRegistryActionRevoked, resourceId []types.ID) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "ActionRevoked", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryActionRevoked)
				if err := _AppRegistry.contract.UnpackLog(evt, "ActionRevoked", log); err != nil {
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

// ParseActionRevoked is a log parse operation binding the contract event 0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430.
//
// Solidity: event ActionRevoked(bytes8 indexed resourceId, string roleName, string actionName)
func (_AppRegistry *appRegistryEvents) ParseActionRevoked(log chainTypes.Log) (*AppRegistryActionRevoked, error) {
	evt := new(AppRegistryActionRevoked)
	if err := _AppRegistry.contract.UnpackLog(evt, "ActionRevoked", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseActionRevokedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event ActionRevoked(bytes8 indexed resourceId, string roleName, string actionName)
func (_AppRegistry *appRegistryEvents) ParseActionRevokedFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryActionRevoked, error) {
	var evts []*AppRegistryActionRevoked
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430") {
			evt, err := _AppRegistry.ParseActionRevoked(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("ActionRevoked event not found")
	}
	return evts, nil
}

// AppRegistryAppOwnerTransferredIterator is returned from FilterAppOwnerTransferred and is used to iterate over the raw logs and unpacked data for AppOwnerTransferred events raised by the AppRegistry contract.
type AppRegistryAppOwnerTransferredIterator struct {
	Evt *AppRegistryAppOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *AppRegistryAppOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryAppOwnerTransferred)
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
		it.Evt = new(AppRegistryAppOwnerTransferred)
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
func (it *AppRegistryAppOwnerTransferredIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryAppOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryAppOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryAppOwnerTransferred represents a AppOwnerTransferred event raised by the AppRegistry contract.
type AppRegistryAppOwnerTransferred struct {
	AppId    types.ID
	AppName  string
	OldOwner common.Address
	NewOwner common.Address
	Raw      chainTypes.Log // Blockchain specific contextual infos
}

// FilterAppOwnerTransferred is a free log retrieval operation binding the contract event 0xd862515a389f54c6adf72b970712257c420c06e58830f73050bc0b9aca86f96d.
//
// Solidity: event AppOwnerTransferred(bytes8 indexed appId, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *appRegistryEvents) FilterAppOwnerTransferred(opts *bind.FilterOpts, appId []types.ID, oldOwner []common.Address) (ablbind.EventIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "AppOwnerTransferred", appIdRule, oldOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryAppOwnerTransferredIterator{contract: _AppRegistry.contract, event: "AppOwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchAppOwnerTransferred is a free log subscription operation binding the contract event 0xd862515a389f54c6adf72b970712257c420c06e58830f73050bc0b9aca86f96d.
//
// Solidity: event AppOwnerTransferred(bytes8 indexed appId, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *appRegistryEvents) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AppRegistryAppOwnerTransferred, appId []types.ID, oldOwner []common.Address) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "AppOwnerTransferred", appIdRule, oldOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryAppOwnerTransferred)
				if err := _AppRegistry.contract.UnpackLog(evt, "AppOwnerTransferred", log); err != nil {
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

// ParseAppOwnerTransferred is a log parse operation binding the contract event 0xd862515a389f54c6adf72b970712257c420c06e58830f73050bc0b9aca86f96d.
//
// Solidity: event AppOwnerTransferred(bytes8 indexed appId, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *appRegistryEvents) ParseAppOwnerTransferred(log chainTypes.Log) (*AppRegistryAppOwnerTransferred, error) {
	evt := new(AppRegistryAppOwnerTransferred)
	if err := _AppRegistry.contract.UnpackLog(evt, "AppOwnerTransferred", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseAppOwnerTransferredFromReceipt parses the event from given transaction receipt.
//
// Solidity: event AppOwnerTransferred(bytes8 indexed appId, string appName, address indexed oldOwner, address newOwner)
func (_AppRegistry *appRegistryEvents) ParseAppOwnerTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryAppOwnerTransferred, error) {
	var evts []*AppRegistryAppOwnerTransferred
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xd862515a389f54c6adf72b970712257c420c06e58830f73050bc0b9aca86f96d") {
			evt, err := _AppRegistry.ParseAppOwnerTransferred(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("AppOwnerTransferred event not found")
	}
	return evts, nil
}

// AppRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the AppRegistry contract.
type AppRegistryRegistrationIterator struct {
	Evt *AppRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *AppRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryRegistration)
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
		it.Evt = new(AppRegistryRegistration)
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
func (it *AppRegistryRegistrationIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRegistration represents a Registration event raised by the AppRegistry contract.
type AppRegistryRegistration struct {
	AppId   types.ID
	AppName string
	Raw     chainTypes.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xf5ef63b2a629619cbea83fa4979a52fa156e638837dd66bc4683dd7fbb8f4709.
//
// Solidity: event Registration(bytes8 indexed appId, string appName)
func (_AppRegistry *appRegistryEvents) FilterRegistration(opts *bind.FilterOpts, appId []types.ID) (ablbind.EventIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Registration", appIdRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRegistrationIterator{contract: _AppRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// WatchRegistration is a free log subscription operation binding the contract event 0xf5ef63b2a629619cbea83fa4979a52fa156e638837dd66bc4683dd7fbb8f4709.
//
// Solidity: event Registration(bytes8 indexed appId, string appName)
func (_AppRegistry *appRegistryEvents) WatchRegistration(opts *bind.WatchOpts, sink chan<- *AppRegistryRegistration, appId []types.ID) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Registration", appIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryRegistration)
				if err := _AppRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
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

// ParseRegistration is a log parse operation binding the contract event 0xf5ef63b2a629619cbea83fa4979a52fa156e638837dd66bc4683dd7fbb8f4709.
//
// Solidity: event Registration(bytes8 indexed appId, string appName)
func (_AppRegistry *appRegistryEvents) ParseRegistration(log chainTypes.Log) (*AppRegistryRegistration, error) {
	evt := new(AppRegistryRegistration)
	if err := _AppRegistry.contract.UnpackLog(evt, "Registration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRegistrationFromReceipt parses the event from given transaction receipt.
//
// Solidity: event Registration(bytes8 indexed appId, string appName)
func (_AppRegistry *appRegistryEvents) ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRegistration, error) {
	var evts []*AppRegistryRegistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf5ef63b2a629619cbea83fa4979a52fa156e638837dd66bc4683dd7fbb8f4709") {
			evt, err := _AppRegistry.ParseRegistration(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("Registration event not found")
	}
	return evts, nil
}

// AppRegistryRoleBoundIterator is returned from FilterRoleBound and is used to iterate over the raw logs and unpacked data for RoleBound events raised by the AppRegistry contract.
type AppRegistryRoleBoundIterator struct {
	Evt *AppRegistryRoleBound // Event containing the contract specifics and raw log

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
func (it *AppRegistryRoleBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryRoleBound)
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
		it.Evt = new(AppRegistryRoleBound)
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
func (it *AppRegistryRoleBoundIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryRoleBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRoleBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRoleBound represents a RoleBound event raised by the AppRegistry contract.
type AppRegistryRoleBound struct {
	ResourceId types.ID
	Subject    common.Address
	RoleName   string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRoleBound is a free log retrieval operation binding the contract event 0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372.
//
// Solidity: event RoleBound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_AppRegistry *appRegistryEvents) FilterRoleBound(opts *bind.FilterOpts, resourceId []types.ID, subject []common.Address) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "RoleBound", resourceIdRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRoleBoundIterator{contract: _AppRegistry.contract, event: "RoleBound", logs: logs, sub: sub}, nil
}

// WatchRoleBound is a free log subscription operation binding the contract event 0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372.
//
// Solidity: event RoleBound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_AppRegistry *appRegistryEvents) WatchRoleBound(opts *bind.WatchOpts, sink chan<- *AppRegistryRoleBound, resourceId []types.ID, subject []common.Address) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "RoleBound", resourceIdRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryRoleBound)
				if err := _AppRegistry.contract.UnpackLog(evt, "RoleBound", log); err != nil {
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

// ParseRoleBound is a log parse operation binding the contract event 0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372.
//
// Solidity: event RoleBound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_AppRegistry *appRegistryEvents) ParseRoleBound(log chainTypes.Log) (*AppRegistryRoleBound, error) {
	evt := new(AppRegistryRoleBound)
	if err := _AppRegistry.contract.UnpackLog(evt, "RoleBound", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRoleBoundFromReceipt parses the event from given transaction receipt.
//
// Solidity: event RoleBound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_AppRegistry *appRegistryEvents) ParseRoleBoundFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRoleBound, error) {
	var evts []*AppRegistryRoleBound
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372") {
			evt, err := _AppRegistry.ParseRoleBound(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("RoleBound event not found")
	}
	return evts, nil
}

// AppRegistryRoleCreationIterator is returned from FilterRoleCreation and is used to iterate over the raw logs and unpacked data for RoleCreation events raised by the AppRegistry contract.
type AppRegistryRoleCreationIterator struct {
	Evt *AppRegistryRoleCreation // Event containing the contract specifics and raw log

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
func (it *AppRegistryRoleCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryRoleCreation)
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
		it.Evt = new(AppRegistryRoleCreation)
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
func (it *AppRegistryRoleCreationIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryRoleCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRoleCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRoleCreation represents a RoleCreation event raised by the AppRegistry contract.
type AppRegistryRoleCreation struct {
	ResourceId types.ID
	RoleName   string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRoleCreation is a free log retrieval operation binding the contract event 0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a.
//
// Solidity: event RoleCreation(bytes8 indexed resourceId, string roleName)
func (_AppRegistry *appRegistryEvents) FilterRoleCreation(opts *bind.FilterOpts, resourceId []types.ID) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "RoleCreation", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRoleCreationIterator{contract: _AppRegistry.contract, event: "RoleCreation", logs: logs, sub: sub}, nil
}

// WatchRoleCreation is a free log subscription operation binding the contract event 0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a.
//
// Solidity: event RoleCreation(bytes8 indexed resourceId, string roleName)
func (_AppRegistry *appRegistryEvents) WatchRoleCreation(opts *bind.WatchOpts, sink chan<- *AppRegistryRoleCreation, resourceId []types.ID) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "RoleCreation", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryRoleCreation)
				if err := _AppRegistry.contract.UnpackLog(evt, "RoleCreation", log); err != nil {
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

// ParseRoleCreation is a log parse operation binding the contract event 0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a.
//
// Solidity: event RoleCreation(bytes8 indexed resourceId, string roleName)
func (_AppRegistry *appRegistryEvents) ParseRoleCreation(log chainTypes.Log) (*AppRegistryRoleCreation, error) {
	evt := new(AppRegistryRoleCreation)
	if err := _AppRegistry.contract.UnpackLog(evt, "RoleCreation", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRoleCreationFromReceipt parses the event from given transaction receipt.
//
// Solidity: event RoleCreation(bytes8 indexed resourceId, string roleName)
func (_AppRegistry *appRegistryEvents) ParseRoleCreationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRoleCreation, error) {
	var evts []*AppRegistryRoleCreation
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a") {
			evt, err := _AppRegistry.ParseRoleCreation(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("RoleCreation event not found")
	}
	return evts, nil
}

// AppRegistryRoleRemovalIterator is returned from FilterRoleRemoval and is used to iterate over the raw logs and unpacked data for RoleRemoval events raised by the AppRegistry contract.
type AppRegistryRoleRemovalIterator struct {
	Evt *AppRegistryRoleRemoval // Event containing the contract specifics and raw log

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
func (it *AppRegistryRoleRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryRoleRemoval)
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
		it.Evt = new(AppRegistryRoleRemoval)
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
func (it *AppRegistryRoleRemovalIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryRoleRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRoleRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRoleRemoval represents a RoleRemoval event raised by the AppRegistry contract.
type AppRegistryRoleRemoval struct {
	ResourceId types.ID
	RoleName   string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRoleRemoval is a free log retrieval operation binding the contract event 0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08.
//
// Solidity: event RoleRemoval(bytes8 indexed resourceId, string roleName)
func (_AppRegistry *appRegistryEvents) FilterRoleRemoval(opts *bind.FilterOpts, resourceId []types.ID) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "RoleRemoval", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRoleRemovalIterator{contract: _AppRegistry.contract, event: "RoleRemoval", logs: logs, sub: sub}, nil
}

// WatchRoleRemoval is a free log subscription operation binding the contract event 0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08.
//
// Solidity: event RoleRemoval(bytes8 indexed resourceId, string roleName)
func (_AppRegistry *appRegistryEvents) WatchRoleRemoval(opts *bind.WatchOpts, sink chan<- *AppRegistryRoleRemoval, resourceId []types.ID) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "RoleRemoval", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryRoleRemoval)
				if err := _AppRegistry.contract.UnpackLog(evt, "RoleRemoval", log); err != nil {
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

// ParseRoleRemoval is a log parse operation binding the contract event 0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08.
//
// Solidity: event RoleRemoval(bytes8 indexed resourceId, string roleName)
func (_AppRegistry *appRegistryEvents) ParseRoleRemoval(log chainTypes.Log) (*AppRegistryRoleRemoval, error) {
	evt := new(AppRegistryRoleRemoval)
	if err := _AppRegistry.contract.UnpackLog(evt, "RoleRemoval", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRoleRemovalFromReceipt parses the event from given transaction receipt.
//
// Solidity: event RoleRemoval(bytes8 indexed resourceId, string roleName)
func (_AppRegistry *appRegistryEvents) ParseRoleRemovalFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRoleRemoval, error) {
	var evts []*AppRegistryRoleRemoval
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08") {
			evt, err := _AppRegistry.ParseRoleRemoval(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("RoleRemoval event not found")
	}
	return evts, nil
}

// AppRegistryRoleUnboundIterator is returned from FilterRoleUnbound and is used to iterate over the raw logs and unpacked data for RoleUnbound events raised by the AppRegistry contract.
type AppRegistryRoleUnboundIterator struct {
	Evt *AppRegistryRoleUnbound // Event containing the contract specifics and raw log

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
func (it *AppRegistryRoleUnboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryRoleUnbound)
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
		it.Evt = new(AppRegistryRoleUnbound)
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
func (it *AppRegistryRoleUnboundIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryRoleUnboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRoleUnboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRoleUnbound represents a RoleUnbound event raised by the AppRegistry contract.
type AppRegistryRoleUnbound struct {
	ResourceId types.ID
	Subject    common.Address
	RoleName   string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRoleUnbound is a free log retrieval operation binding the contract event 0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3.
//
// Solidity: event RoleUnbound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_AppRegistry *appRegistryEvents) FilterRoleUnbound(opts *bind.FilterOpts, resourceId []types.ID, subject []common.Address) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "RoleUnbound", resourceIdRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRoleUnboundIterator{contract: _AppRegistry.contract, event: "RoleUnbound", logs: logs, sub: sub}, nil
}

// WatchRoleUnbound is a free log subscription operation binding the contract event 0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3.
//
// Solidity: event RoleUnbound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_AppRegistry *appRegistryEvents) WatchRoleUnbound(opts *bind.WatchOpts, sink chan<- *AppRegistryRoleUnbound, resourceId []types.ID, subject []common.Address) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "RoleUnbound", resourceIdRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryRoleUnbound)
				if err := _AppRegistry.contract.UnpackLog(evt, "RoleUnbound", log); err != nil {
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

// ParseRoleUnbound is a log parse operation binding the contract event 0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3.
//
// Solidity: event RoleUnbound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_AppRegistry *appRegistryEvents) ParseRoleUnbound(log chainTypes.Log) (*AppRegistryRoleUnbound, error) {
	evt := new(AppRegistryRoleUnbound)
	if err := _AppRegistry.contract.UnpackLog(evt, "RoleUnbound", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRoleUnboundFromReceipt parses the event from given transaction receipt.
//
// Solidity: event RoleUnbound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_AppRegistry *appRegistryEvents) ParseRoleUnboundFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryRoleUnbound, error) {
	var evts []*AppRegistryRoleUnbound
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3") {
			evt, err := _AppRegistry.ParseRoleUnbound(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("RoleUnbound event not found")
	}
	return evts, nil
}

// AppRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the AppRegistry contract.
type AppRegistryUnregistrationIterator struct {
	Evt *AppRegistryUnregistration // Event containing the contract specifics and raw log

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
func (it *AppRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AppRegistryUnregistration)
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
		it.Evt = new(AppRegistryUnregistration)
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
func (it *AppRegistryUnregistrationIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryUnregistration represents a Unregistration event raised by the AppRegistry contract.
type AppRegistryUnregistration struct {
	AppId   types.ID
	AppName string
	Raw     chainTypes.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0xfe4087afbde6adacfa58b7ae569b654c6445bdf8781edab97394ea74d2a5f85e.
//
// Solidity: event Unregistration(bytes8 indexed appId, string appName)
func (_AppRegistry *appRegistryEvents) FilterUnregistration(opts *bind.FilterOpts, appId []types.ID) (ablbind.EventIterator, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Unregistration", appIdRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryUnregistrationIterator{contract: _AppRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// WatchUnregistration is a free log subscription operation binding the contract event 0xfe4087afbde6adacfa58b7ae569b654c6445bdf8781edab97394ea74d2a5f85e.
//
// Solidity: event Unregistration(bytes8 indexed appId, string appName)
func (_AppRegistry *appRegistryEvents) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *AppRegistryUnregistration, appId []types.ID) (event.Subscription, error) {

	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Unregistration", appIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AppRegistryUnregistration)
				if err := _AppRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
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

// ParseUnregistration is a log parse operation binding the contract event 0xfe4087afbde6adacfa58b7ae569b654c6445bdf8781edab97394ea74d2a5f85e.
//
// Solidity: event Unregistration(bytes8 indexed appId, string appName)
func (_AppRegistry *appRegistryEvents) ParseUnregistration(log chainTypes.Log) (*AppRegistryUnregistration, error) {
	evt := new(AppRegistryUnregistration)
	if err := _AppRegistry.contract.UnpackLog(evt, "Unregistration", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseUnregistrationFromReceipt parses the event from given transaction receipt.
//
// Solidity: event Unregistration(bytes8 indexed appId, string appName)
func (_AppRegistry *appRegistryEvents) ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*AppRegistryUnregistration, error) {
	var evts []*AppRegistryUnregistration
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xfe4087afbde6adacfa58b7ae569b654c6445bdf8781edab97394ea74d2a5f85e") {
			evt, err := _AppRegistry.ParseUnregistration(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("Unregistration event not found")
	}
	return evts, nil
}

// Manager is contract wrapper struct
type AppRegistryContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	AppRegistryCaller
	AppRegistryTransactor
	AppRegistryEvents
}

func NewAppRegistryContract(backend ablbind.ContractBackend) (*AppRegistryContract, error) {
	deployment, exist := backend.Deployment("AppRegistry")
	if !exist {
		evmABI, err := abi.JSON(strings.NewReader(AppRegistryABI))
		if err != nil {
			return nil, err
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(AppRegistryAddress),
			common.HexToHash(AppRegistryTxHash),
			new(big.Int).SetBytes(common.HexToHash(AppRegistryCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, "AppRegistry", backend)

	contract := &AppRegistryContract{
		Deployment: deployment,
		client:     backend,

		AppRegistryCaller:     &appRegistryCaller{base},
		AppRegistryTransactor: &appRegistryTransactor{base, backend},
		AppRegistryEvents:     &appRegistryEvents{base},
	}

	return contract, nil
}
