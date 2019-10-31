package contracts

import (
	"errors"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	types "github.com/airbloc/airbloc-go/shared/adapter/types"
	platform "github.com/klaytn/klaytn"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// AccountsABI is the input ABI used to generate the binding from.
const (
	AccountsAddress   = "0xdFe31bb8C440F4b0C6319Fe95378E1418A240e2F"
	AccountsTxHash    = "0xe391330246621ef3702bce179b1865268940dc56dbb1429fd1f5567dc3044dce"
	AccountsCreatedAt = "0x000000000000000000000000000000000000000000000000000000000063b9ea"
	AccountsABI       = "[{\"constant\":true,\"inputs\":[],\"name\":\"numberOfAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"signature\":\"0x0f03e4c3\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"identityHashToAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x17aba2d3\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"accounts\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"0xf4a3fad5\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"controllerReg\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"SignUp\",\"signature\":\"0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"TemporaryCreated\",\"signature\":\"0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"accountId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Unlocked\",\"signature\":\"0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"create\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0xefc81a8c\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"createTemporary\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x56003f0f\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identityPreimage\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"unlockTemporary\",\"outputs\":[],\"payable\":false,\"signature\":\"0x2299219d\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"signature\":\"0x92eefe9b\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"getAccount\",\"outputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0xf9292ddb\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"getAccountByIdentityHash\",\"outputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0xc75aeb7e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getAccountId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0xe0b490f7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"getAccountIdByIdentityHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x793d5046\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getAccountIdFromSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x23d0601d\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"isTemporary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x6b886888\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"isControllerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xa83038e7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x97e4fea7\",\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewAccountsCaller(contract *ablbind.BoundContract) AccountsCaller {
	return AccountsCaller{contract: contract}
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewAccountsTransactor(contract *ablbind.BoundContract) AccountsTransactor {
	return AccountsTransactor{contract: contract}
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFilterer struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

func NewAccountsFilterer(contract *ablbind.BoundContract) AccountsFilterer {
	return AccountsFilterer{contract: contract}
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCaller) Exists(opts *bind.CallOpts, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "exists", accountId)
	return *ret0, err
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns(types.Account)
func (_Accounts *AccountsCaller) GetAccount(opts *bind.CallOpts, accountId types.ID) (types.Account, error) {
	var (
		ret0 = new(types.Account)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccount", accountId)
	return *ret0, err
}

// GetAccountByIdentityHash is a free data retrieval call binding the contract method 0xc75aeb7e.
//
// Solidity: function getAccountByIdentityHash(bytes32 identityHash) constant returns(types.Account)
func (_Accounts *AccountsCaller) GetAccountByIdentityHash(opts *bind.CallOpts, identityHash common.Hash) (types.Account, error) {
	var (
		ret0 = new(types.Account)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccountByIdentityHash", identityHash)
	return *ret0, err
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountId(opts *bind.CallOpts, sender common.Address) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccountId", sender)
	return *ret0, err
}

// GetAccountIdByIdentityHash is a free data retrieval call binding the contract method 0x793d5046.
//
// Solidity: function getAccountIdByIdentityHash(bytes32 identityHash) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountIdByIdentityHash(opts *bind.CallOpts, identityHash common.Hash) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccountIdByIdentityHash", identityHash)
	return *ret0, err
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountIdFromSignature(opts *bind.CallOpts, messageHash common.Hash, signature []byte) (types.ID, error) {
	var (
		ret0 = new(types.ID)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccountIdFromSignature", messageHash, signature)
	return *ret0, err
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address sender, bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCaller) IsControllerOf(opts *bind.CallOpts, sender common.Address, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isControllerOf", sender, accountId)
	return *ret0, err
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 accountId) constant returns(bool)
func (_Accounts *AccountsCaller) IsTemporary(opts *bind.CallOpts, accountId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isTemporary", accountId)
	return *ret0, err
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (_Accounts *AccountsTransactor) Create(opts *ablbind.TransactOpts) (*chainTypes.Transaction, error) {
	return _Accounts.contract.Transact(opts, "create")
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (_Accounts *AccountsTransactor) CreateTemporary(opts *ablbind.TransactOpts, identityHash common.Hash) (*chainTypes.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createTemporary", identityHash)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (_Accounts *AccountsTransactor) SetController(opts *ablbind.TransactOpts, controller common.Address) (*chainTypes.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setController", controller)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (_Accounts *AccountsTransactor) UnlockTemporary(opts *ablbind.TransactOpts, identityPreimage common.Hash, newOwner common.Address, passwordSignature []byte) (*chainTypes.Transaction, error) {
	return _Accounts.contract.Transact(opts, "unlockTemporary", identityPreimage, newOwner, passwordSignature)
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
func (_Accounts *AccountsFilterer) FilterSignUp(opts *bind.FilterOpts, owner []common.Address) (ablbind.EventIterator, error) {

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
func (_Accounts *AccountsFilterer) WatchSignUp(opts *bind.WatchOpts, sink chan<- *AccountsSignUp, owner []common.Address) (event.Subscription, error) {

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
func (_Accounts *AccountsFilterer) ParseSignUp(log chainTypes.Log) (*AccountsSignUp, error) {
	evt := new(AccountsSignUp)
	if err := _Accounts.contract.UnpackLog(evt, "SignUp", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterSignUp parses the event from given transaction receipt.
//
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
func (_Accounts *AccountsFilterer) ParseSignUpFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsSignUp, error) {
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
func (_Accounts *AccountsFilterer) FilterTemporaryCreated(opts *bind.FilterOpts, proxy []common.Address, identityHash []common.Hash) (ablbind.EventIterator, error) {

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
func (_Accounts *AccountsFilterer) WatchTemporaryCreated(opts *bind.WatchOpts, sink chan<- *AccountsTemporaryCreated, proxy []common.Address, identityHash []common.Hash) (event.Subscription, error) {

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
func (_Accounts *AccountsFilterer) ParseTemporaryCreated(log chainTypes.Log) (*AccountsTemporaryCreated, error) {
	evt := new(AccountsTemporaryCreated)
	if err := _Accounts.contract.UnpackLog(evt, "TemporaryCreated", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterTemporaryCreated parses the event from given transaction receipt.
//
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
func (_Accounts *AccountsFilterer) ParseTemporaryCreatedFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsTemporaryCreated, error) {
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
func (_Accounts *AccountsFilterer) FilterUnlocked(opts *bind.FilterOpts, identityHash []common.Hash, accountId []types.ID) (ablbind.EventIterator, error) {

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
func (_Accounts *AccountsFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *AccountsUnlocked, identityHash []common.Hash, accountId []types.ID) (event.Subscription, error) {

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
func (_Accounts *AccountsFilterer) ParseUnlocked(log chainTypes.Log) (*AccountsUnlocked, error) {
	evt := new(AccountsUnlocked)
	if err := _Accounts.contract.UnpackLog(evt, "Unlocked", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// FilterUnlocked parses the event from given transaction receipt.
//
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (_Accounts *AccountsFilterer) ParseUnlockedFromReceipt(receipt *chainTypes.Receipt) ([]*AccountsUnlocked, error) {
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
