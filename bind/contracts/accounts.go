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

// AccountsABI is the input ABI used to generate the binding from.
const (
	AccountsAddress   = "0x4E48E6356EB97BF9869EbCd7568780E3b9f1DB57"
	AccountsTxHash    = "0xd6b0456617b9c3c994ccddc9dc6d8ab71dfdd033978b51bc1f63c35cc6afa0df"
	AccountsCreatedAt = "0x0000000000000000000000000000000000000000000000000000000000823881"
	AccountsABI       = "[{\"constant\":true,\"inputs\":[],\"name\":\"numberOfAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"signature\":\"0x0f03e4c3\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"identityHashToAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x17aba2d3\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"accounts\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"0xf4a3fad5\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"controllerReg\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"SignUp\",\"signature\":\"0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"TemporaryCreated\",\"signature\":\"0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"accountId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Unlocked\",\"signature\":\"0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"create\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0xefc81a8c\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"createTemporary\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x56003f0f\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identityPreimage\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"unlockTemporary\",\"outputs\":[],\"payable\":false,\"signature\":\"0x2299219d\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"signature\":\"0x92eefe9b\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"getAccount\",\"outputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0xf9292ddb\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"getAccountByIdentityHash\",\"outputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0xc75aeb7e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getAccountId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0xe0b490f7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"getAccountIdByIdentityHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x793d5046\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getAccountIdFromSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x23d0601d\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"isTemporary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x6b886888\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"isControllerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xa83038e7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x97e4fea7\",\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller interface {
	Exists(
		ctx context.Context,
		accountId types.ID,
	) (
		bool,
		error,
	)
	GetAccount(
		ctx context.Context,
		accountId types.ID,
	) (
		types.Account,
		error,
	)
	GetAccountByIdentityHash(
		ctx context.Context,
		identityHash common.Hash,
	) (
		types.Account,
		error,
	)
	GetAccountId(
		ctx context.Context,
		sender common.Address,
	) (
		types.ID,
		error,
	)
	GetAccountIdByIdentityHash(
		ctx context.Context,
		identityHash common.Hash,
	) (
		types.ID,
		error,
	)
	GetAccountIdFromSignature(
		ctx context.Context,
		messageHash common.Hash,
		signature []byte,
	) (
		types.ID,
		error,
	)
	IsControllerOf(
		ctx context.Context,
		sender common.Address,
		accountId types.ID,
	) (
		bool,
		error,
	)
	IsTemporary(
		ctx context.Context,
		accountId types.ID,
	) (
		bool,
		error,
	)
}

type accountsCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 accountId) constant returns(bool)
func (_Accounts *accountsCaller) Exists(ctx context.Context, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Accounts.contract.Call(&bind.CallOpts{Context: ctx}, out, "exists", accountId)
	return *ret0, err
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns(types.Account)
func (_Accounts *accountsCaller) GetAccount(ctx context.Context, accountId types.ID) (types.Account, error) {
	var (
		ret0 = new(types.Account)
	)
	out := ret0

	err := _Accounts.contract.Call(&bind.CallOpts{Context: ctx}, out, "getAccount", accountId)
	return *ret0, err
}

// GetAccountByIdentityHash is a free data retrieval call binding the contract method 0xc75aeb7e.
//
// Solidity: function getAccountByIdentityHash(bytes32 identityHash) constant returns(types.Account)
func (_Accounts *accountsCaller) GetAccountByIdentityHash(ctx context.Context, identityHash common.Hash) (types.Account, error) {
	var (
		ret0 = new(types.Account)
	)
	out := ret0

	err := _Accounts.contract.Call(&bind.CallOpts{Context: ctx}, out, "getAccountByIdentityHash", identityHash)
	return *ret0, err
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (_Accounts *accountsCaller) GetAccountId(ctx context.Context, sender common.Address) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := ret0

	err := _Accounts.contract.Call(&bind.CallOpts{Context: ctx}, out, "getAccountId", sender)
	return *ret0, err
}

// GetAccountIdByIdentityHash is a free data retrieval call binding the contract method 0x793d5046.
//
// Solidity: function getAccountIdByIdentityHash(bytes32 identityHash) constant returns(bytes8)
func (_Accounts *accountsCaller) GetAccountIdByIdentityHash(ctx context.Context, identityHash common.Hash) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := ret0

	err := _Accounts.contract.Call(&bind.CallOpts{Context: ctx}, out, "getAccountIdByIdentityHash", identityHash)
	return *ret0, err
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (_Accounts *accountsCaller) GetAccountIdFromSignature(ctx context.Context, messageHash common.Hash, signature []byte) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := ret0

	err := _Accounts.contract.Call(&bind.CallOpts{Context: ctx}, out, "getAccountIdFromSignature", messageHash, signature)
	return *ret0, err
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address sender, bytes8 accountId) constant returns(bool)
func (_Accounts *accountsCaller) IsControllerOf(ctx context.Context, sender common.Address, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Accounts.contract.Call(&bind.CallOpts{Context: ctx}, out, "isControllerOf", sender, accountId)
	return *ret0, err
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 accountId) constant returns(bool)
func (_Accounts *accountsCaller) IsTemporary(ctx context.Context, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Accounts.contract.Call(&bind.CallOpts{Context: ctx}, out, "isTemporary", accountId)
	return *ret0, err
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor interface {
	Create(
		ctx context.Context,
		opts *ablbind.TransactOpts,
	) (*chainTypes.Receipt, error)
	CreateTemporary(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		identityHash common.Hash,
	) (*chainTypes.Receipt, error)
	SetController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		controller common.Address,
	) (*chainTypes.Receipt, error)
	UnlockTemporary(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		identityPreimage common.Hash,
		newOwner common.Address,
		passwordSignature []byte,
	) (*chainTypes.Receipt, error)
}

type accountsTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
	backend  ablbind.ContractBackend
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (_Accounts *accountsTransactor) Create(
	ctx context.Context,
	opts *ablbind.TransactOpts,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	tx, err := _Accounts.contract.Transact(opts, "create")
	if err != nil {
		return nil, err
	}
	return _Accounts.backend.WaitMined(ctx, tx)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (_Accounts *accountsTransactor) CreateTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityHash common.Hash,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	tx, err := _Accounts.contract.Transact(opts, "createTemporary", identityHash)
	if err != nil {
		return nil, err
	}
	return _Accounts.backend.WaitMined(ctx, tx)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (_Accounts *accountsTransactor) SetController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	controller common.Address,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	tx, err := _Accounts.contract.Transact(opts, "setController", controller)
	if err != nil {
		return nil, err
	}
	return _Accounts.backend.WaitMined(ctx, tx)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (_Accounts *accountsTransactor) UnlockTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityPreimage common.Hash,
	newOwner common.Address,
	passwordSignature []byte,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	tx, err := _Accounts.contract.Transact(opts, "unlockTemporary", identityPreimage, newOwner, passwordSignature)
	if err != nil {
		return nil, err
	}
	return _Accounts.backend.WaitMined(ctx, tx)
}

type AccountsEvents interface {
	AccountsEventFilterer
	AccountsEventParser
	AccountsEventWatcher
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsEventFilterer interface {
	// Filterer
	FilterSignUp(
		opts *bind.FilterOpts,
		owner []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterTemporaryCreated(
		opts *bind.FilterOpts,
		proxy []common.Address, identityHash []common.Hash,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterUnlocked(
		opts *bind.FilterOpts,
		identityHash []common.Hash, accountId []types.ID,
	) (ablbind.EventIterator, error)
}

type AccountsEventParser interface {
	// Parser
	ParseSignUp(log chainTypes.Log) (*AccountsSignUp, error)
	ParseSignUpFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsSignUp, error)

	// Parser
	ParseTemporaryCreated(log chainTypes.Log) (*AccountsTemporaryCreated, error)
	ParseTemporaryCreatedFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsTemporaryCreated, error)

	// Parser
	ParseUnlocked(log chainTypes.Log) (*AccountsUnlocked, error)
	ParseUnlockedFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsUnlocked, error)
}

type AccountsEventWatcher interface {
	// Watcher
	WatchSignUp(
		opts *bind.WatchOpts,
		sink chan<- *AccountsSignUp,
		owner []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchTemporaryCreated(
		opts *bind.WatchOpts,
		sink chan<- *AccountsTemporaryCreated,
		proxy []common.Address, identityHash []common.Hash,
	) (event.Subscription, error)

	// Watcher
	WatchUnlocked(
		opts *bind.WatchOpts,
		sink chan<- *AccountsUnlocked,
		identityHash []common.Hash, accountId []types.ID,
	) (event.Subscription, error)
}

type accountsEvents struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsSignUpIterator is returned from FilterSignUp and is used to iterate over the raw logs and unpacked data for SignUp events raised by the Accounts contract.
type AccountsSignUpIterator struct {
	Evt *AccountsSignUp // Event containing the contract specifics and raw log

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
func (it *AccountsSignUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AccountsSignUp)
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
		it.Evt = new(AccountsSignUp)
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
func (it *AccountsSignUpIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsSignUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsSignUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsSignUp represents a SignUp event raised by the Accounts contract.
type AccountsSignUp struct {
	Owner     common.Address
	AccountId types.ID
	Raw       chainTypes.Log // Blockchain specific contextual infos
}

// FilterSignUp is a free log retrieval operation binding the contract event 0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4.
//
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
func (_Accounts *accountsEvents) FilterSignUp(opts *bind.FilterOpts, owner []common.Address) (ablbind.EventIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "SignUp", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsSignUpIterator{contract: _Accounts.contract, event: "SignUp", logs: logs, sub: sub}, nil
}

// WatchSignUp is a free log subscription operation binding the contract event 0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4.
//
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
func (_Accounts *accountsEvents) WatchSignUp(opts *bind.WatchOpts, sink chan<- *AccountsSignUp, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "SignUp", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AccountsSignUp)
				if err := _Accounts.contract.UnpackLog(evt, "SignUp", log); err != nil {
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

// ParseSignUp is a log parse operation binding the contract event 0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4.
//
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
func (_Accounts *accountsEvents) ParseSignUp(log chainTypes.Log) (*AccountsSignUp, error) {
	evt := new(AccountsSignUp)
	if err := _Accounts.contract.UnpackLog(evt, "SignUp", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseSignUpFromReceipt parses the event from given transaction receipt.
//
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
func (_Accounts *accountsEvents) ParseSignUpFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsSignUp, error) {
	var evts []*AccountsSignUp
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4") {
			evt, err := _Accounts.ParseSignUp(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("SignUp event not found")
	}
	return evts, nil
}

// AccountsTemporaryCreatedIterator is returned from FilterTemporaryCreated and is used to iterate over the raw logs and unpacked data for TemporaryCreated events raised by the Accounts contract.
type AccountsTemporaryCreatedIterator struct {
	Evt *AccountsTemporaryCreated // Event containing the contract specifics and raw log

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
func (it *AccountsTemporaryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AccountsTemporaryCreated)
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
		it.Evt = new(AccountsTemporaryCreated)
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
func (it *AccountsTemporaryCreatedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsTemporaryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsTemporaryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsTemporaryCreated represents a TemporaryCreated event raised by the Accounts contract.
type AccountsTemporaryCreated struct {
	Proxy        common.Address
	IdentityHash common.Hash
	AccountId    types.ID
	Raw          chainTypes.Log // Blockchain specific contextual infos
}

// FilterTemporaryCreated is a free log retrieval operation binding the contract event 0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e.
//
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
func (_Accounts *accountsEvents) FilterTemporaryCreated(opts *bind.FilterOpts, proxy []common.Address, identityHash []common.Hash) (ablbind.EventIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "TemporaryCreated", proxyRule, identityHashRule)
	if err != nil {
		return nil, err
	}
	return &AccountsTemporaryCreatedIterator{contract: _Accounts.contract, event: "TemporaryCreated", logs: logs, sub: sub}, nil
}

// WatchTemporaryCreated is a free log subscription operation binding the contract event 0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e.
//
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
func (_Accounts *accountsEvents) WatchTemporaryCreated(opts *bind.WatchOpts, sink chan<- *AccountsTemporaryCreated, proxy []common.Address, identityHash []common.Hash) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "TemporaryCreated", proxyRule, identityHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AccountsTemporaryCreated)
				if err := _Accounts.contract.UnpackLog(evt, "TemporaryCreated", log); err != nil {
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

// ParseTemporaryCreated is a log parse operation binding the contract event 0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e.
//
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
func (_Accounts *accountsEvents) ParseTemporaryCreated(log chainTypes.Log) (*AccountsTemporaryCreated, error) {
	evt := new(AccountsTemporaryCreated)
	if err := _Accounts.contract.UnpackLog(evt, "TemporaryCreated", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseTemporaryCreatedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
func (_Accounts *accountsEvents) ParseTemporaryCreatedFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsTemporaryCreated, error) {
	var evts []*AccountsTemporaryCreated
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e") {
			evt, err := _Accounts.ParseTemporaryCreated(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("TemporaryCreated event not found")
	}
	return evts, nil
}

// AccountsUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the Accounts contract.
type AccountsUnlockedIterator struct {
	Evt *AccountsUnlocked // Event containing the contract specifics and raw log

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
func (it *AccountsUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(AccountsUnlocked)
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
		it.Evt = new(AccountsUnlocked)
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
func (it *AccountsUnlockedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsUnlocked represents a Unlocked event raised by the Accounts contract.
type AccountsUnlocked struct {
	IdentityHash common.Hash
	AccountId    types.ID
	NewOwner     common.Address
	Raw          chainTypes.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0.
//
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (_Accounts *accountsEvents) FilterUnlocked(opts *bind.FilterOpts, identityHash []common.Hash, accountId []types.ID) (ablbind.EventIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Unlocked", identityHashRule, accountIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountsUnlockedIterator{contract: _Accounts.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0.
//
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (_Accounts *accountsEvents) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *AccountsUnlocked, identityHash []common.Hash, accountId []types.ID) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Unlocked", identityHashRule, accountIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(AccountsUnlocked)
				if err := _Accounts.contract.UnpackLog(evt, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0.
//
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (_Accounts *accountsEvents) ParseUnlocked(log chainTypes.Log) (*AccountsUnlocked, error) {
	evt := new(AccountsUnlocked)
	if err := _Accounts.contract.UnpackLog(evt, "Unlocked", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseUnlockedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (_Accounts *accountsEvents) ParseUnlockedFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsUnlocked, error) {
	var evts []*AccountsUnlocked
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0") {
			evt, err := _Accounts.ParseUnlocked(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("Unlocked event not found")
	}
	return evts, nil
}

// Manager is contract wrapper struct
type AccountsContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	AccountsCaller
	AccountsTransactor
	AccountsEvents
}

func NewAccountsContract(backend ablbind.ContractBackend) (*AccountsContract, error) {
	deployment, exist := backend.Deployment("Accounts")
	if !exist {
		evmABI, err := abi.JSON(strings.NewReader(AccountsABI))
		if err != nil {
			return nil, err
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(AccountsAddress),
			common.HexToHash(AccountsTxHash),
			new(big.Int).SetBytes(common.HexToHash(AccountsCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, "Accounts", backend)

	contract := &AccountsContract{
		Deployment: deployment,
		client:     backend,

		AccountsCaller:     &accountsCaller{base},
		AccountsTransactor: &accountsTransactor{base, backend},
		AccountsEvents:     &accountsEvents{base},
	}

	return contract, nil
}
