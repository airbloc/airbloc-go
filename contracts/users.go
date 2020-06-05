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

// UsersABI is the input ABI used to generate the binding from.
const (
	UsersAddress   = "0x5DA52F7f59bC0c0Bf706c6DCBbc65D9750CE1551"
	UsersTxHash    = "0x8b46629739031ef2ddb33edf1eec96d0abd158eeeca9c7d1079110ad4cebc5bb"
	UsersCreatedAt = "0x0000000000000000000000000000000000000000000000000000000000d6d8a8"
	UsersABI       = "[{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"roleName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"}],\"name\":\"revokeAction\",\"outputs\":[],\"payable\":false,\"signature\":\"0x10a789f8\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ACTION_CONSENT_CREATE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"signature\":\"0x5e99d46f\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"createRole\",\"outputs\":[],\"payable\":false,\"signature\":\"0x617551d4\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfUsers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"signature\":\"0x75ce46a7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_TEMP_DATA_CONTROLLER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"signature\":\"0x7db664b8\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DATA_CONTROLLER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"signature\":\"0x8007a4b6\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"action\",\"type\":\"string\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x96c57715\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ACTION_CONSENT_MODIFY\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"signature\":\"0x9c1808aa\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"signature\":\"0xac1465c6\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"subject\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"unbindRole\",\"outputs\":[],\"payable\":false,\"signature\":\"0xb31ec2af\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"identityHashToUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0xb85860ae\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"subject\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"bindRole\",\"outputs\":[],\"payable\":false,\"signature\":\"0xc8526eed\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ACTION_USER_TRANSFER_OWNERSHIP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"signature\":\"0xd0140505\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"name\":\"roleName\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"}],\"name\":\"grantAction\",\"outputs\":[],\"payable\":false,\"signature\":\"0xe602362d\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"signature\":\"0xeadfe35f\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"controllerReg\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"SignedUp\",\"signature\":\"0xb60195123201ae92c75c0daf6703a414df492031738ac8d8f5a09c81277c0a87\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"oldController\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"ControllerChanged\",\"signature\":\"0xf2fe27b4381d3faea12791cd499b21e10627d1cd2f09101c2c33547ecf711fa1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"feePayer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"TemporaryCreated\",\"signature\":\"0x67048177171beb7aeaaa2c3ac9ab1bc44f168a64b33ef5d502d607a39bb43480\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"TemporaryUnlocked\",\"signature\":\"0xbd9af3a42cb6d433090684c1dd3af90f4ef6f0ae87fb219c7b38d388385c9be7\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleCreation\",\"signature\":\"0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleRemoval\",\"signature\":\"0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleBound\",\"signature\":\"0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleUnbound\",\"signature\":\"0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"actionName\",\"type\":\"string\"}],\"name\":\"ActionGranted\",\"signature\":\"0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"resourceId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"actionName\",\"type\":\"string\"}],\"name\":\"ActionRevoked\",\"signature\":\"0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"create\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0xefc81a8c\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"createTemporary\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x56003f0f\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identityPreimage\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"unlockTemporary\",\"outputs\":[],\"payable\":false,\"signature\":\"0x564929bf\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"signature\":\"0x92eefe9b\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x47ba65d2\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"getByIdentityHash\",\"outputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"signature\":\"0x628ecbda\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x65f68c89\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityHash\",\"type\":\"bytes32\"}],\"name\":\"getIdByIdentityHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"signature\":\"0x93f9c90e\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"isTemporary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x6b886888\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"isControllerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0xa83038e7\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"signature\":\"0x97e4fea7\",\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

// UsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsersCaller interface {
	Exists(
		ctx context.Context,
		userId [8]byte,
	) (
		bool,
		error,
	)
	Get(
		ctx context.Context,
		userId [8]byte,
	) (
		types.User,
		error,
	)
	GetByIdentityHash(
		ctx context.Context,
		identityHash common.Hash,
	) (
		types.User,
		error,
	)
	GetId(
		ctx context.Context,
		owner common.Address,
	) (
		[8]byte,
		error,
	)
	GetIdByIdentityHash(
		ctx context.Context,
		identityHash common.Hash,
	) (
		[8]byte,
		error,
	)
	IsControllerOf(
		ctx context.Context,
		controller common.Address,
		userId [8]byte,
	) (
		bool,
		error,
	)
	IsTemporary(
		ctx context.Context,
		userId [8]byte,
	) (
		bool,
		error,
	)
}

type usersCaller struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 userId) constant returns(bool)
func (_Users *usersCaller) Exists(ctx context.Context, userId [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Users.contract.Call(&bind.CallOpts{Context: ctx}, out, "exists", userId)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x47ba65d2.
//
// Solidity: function get(bytes8 userId) constant returns(types.User)
func (_Users *usersCaller) Get(ctx context.Context, userId [8]byte) (types.User, error) {
	var (
		ret0 = new(types.User)
	)
	out := ret0

	err := _Users.contract.Call(&bind.CallOpts{Context: ctx}, out, "get", userId)
	return *ret0, err
}

// GetByIdentityHash is a free data retrieval call binding the contract method 0x628ecbda.
//
// Solidity: function getByIdentityHash(bytes32 identityHash) constant returns(types.User)
func (_Users *usersCaller) GetByIdentityHash(ctx context.Context, identityHash common.Hash) (types.User, error) {
	var (
		ret0 = new(types.User)
	)
	out := ret0

	err := _Users.contract.Call(&bind.CallOpts{Context: ctx}, out, "getByIdentityHash", identityHash)
	return *ret0, err
}

// GetId is a free data retrieval call binding the contract method 0x65f68c89.
//
// Solidity: function getId(address owner) constant returns(bytes8)
func (_Users *usersCaller) GetId(ctx context.Context, owner common.Address) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0

	err := _Users.contract.Call(&bind.CallOpts{Context: ctx}, out, "getId", owner)
	return *ret0, err
}

// GetIdByIdentityHash is a free data retrieval call binding the contract method 0x93f9c90e.
//
// Solidity: function getIdByIdentityHash(bytes32 identityHash) constant returns(bytes8)
func (_Users *usersCaller) GetIdByIdentityHash(ctx context.Context, identityHash common.Hash) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0

	err := _Users.contract.Call(&bind.CallOpts{Context: ctx}, out, "getIdByIdentityHash", identityHash)
	return *ret0, err
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address controller, bytes8 userId) constant returns(bool)
func (_Users *usersCaller) IsControllerOf(ctx context.Context, controller common.Address, userId [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Users.contract.Call(&bind.CallOpts{Context: ctx}, out, "isControllerOf", controller, userId)
	return *ret0, err
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 userId) constant returns(bool)
func (_Users *usersCaller) IsTemporary(ctx context.Context, userId [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0

	err := _Users.contract.Call(&bind.CallOpts{Context: ctx}, out, "isTemporary", userId)
	return *ret0, err
}

// UsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsersTransactor interface {
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
		newController common.Address,
	) (*chainTypes.Receipt, error)
	UnlockTemporary(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		identityPreimage common.Hash,
		newOwner common.Address,
	) (*chainTypes.Receipt, error)
}

type usersTransactor struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
	backend  ablbind.ContractBackend
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (_Users *usersTransactor) Create(
	ctx context.Context,
	opts *ablbind.TransactOpts,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _Users.contract.Transact(opts, "create")
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (_Users *usersTransactor) CreateTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityHash common.Hash,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _Users.contract.Transact(opts, "createTemporary", identityHash)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_Users *usersTransactor) SetController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	newController common.Address,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _Users.contract.Transact(opts, "setController", newController)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x564929bf.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner) returns()
func (_Users *usersTransactor) UnlockTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityPreimage common.Hash,
	newOwner common.Address,
) (*chainTypes.Receipt, error) {
	if opts == nil {
		opts = &ablbind.TransactOpts{}
	}
	opts.Context = ctx

	return _Users.contract.Transact(opts, "unlockTemporary", identityPreimage, newOwner)
}

type UsersEvents interface {
	UsersEventFilterer
	UsersEventParser
	UsersEventWatcher
}

// UsersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsersEventFilterer interface {
	// Filterer
	FilterActionGranted(
		opts *bind.FilterOpts,
		resourceId [][8]byte,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterActionRevoked(
		opts *bind.FilterOpts,
		resourceId [][8]byte,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterControllerChanged(
		opts *bind.FilterOpts,
		userId [][8]byte,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRoleBound(
		opts *bind.FilterOpts,
		resourceId [][8]byte, subject []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRoleCreation(
		opts *bind.FilterOpts,
		resourceId [][8]byte,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRoleRemoval(
		opts *bind.FilterOpts,
		resourceId [][8]byte,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterRoleUnbound(
		opts *bind.FilterOpts,
		resourceId [][8]byte, subject []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterSignedUp(
		opts *bind.FilterOpts,
		owner []common.Address,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterTemporaryCreated(
		opts *bind.FilterOpts,
		proxy []common.Address, feePayer []common.Address, identityHash []common.Hash,
	) (ablbind.EventIterator, error)

	// Filterer
	FilterTemporaryUnlocked(
		opts *bind.FilterOpts,
		identityHash []common.Hash, userId [][8]byte,
	) (ablbind.EventIterator, error)
}

type UsersEventParser interface {
	// Parser
	ParseActionGranted(log chainTypes.Log) (*UsersActionGranted, error)
	ParseActionGrantedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersActionGranted, error)

	// Parser
	ParseActionRevoked(log chainTypes.Log) (*UsersActionRevoked, error)
	ParseActionRevokedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersActionRevoked, error)

	// Parser
	ParseControllerChanged(log chainTypes.Log) (*UsersControllerChanged, error)
	ParseControllerChangedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersControllerChanged, error)

	// Parser
	ParseRoleBound(log chainTypes.Log) (*UsersRoleBound, error)
	ParseRoleBoundFromReceipt(receipt *chainTypes.Receipt) ([]*UsersRoleBound, error)

	// Parser
	ParseRoleCreation(log chainTypes.Log) (*UsersRoleCreation, error)
	ParseRoleCreationFromReceipt(receipt *chainTypes.Receipt) ([]*UsersRoleCreation, error)

	// Parser
	ParseRoleRemoval(log chainTypes.Log) (*UsersRoleRemoval, error)
	ParseRoleRemovalFromReceipt(receipt *chainTypes.Receipt) ([]*UsersRoleRemoval, error)

	// Parser
	ParseRoleUnbound(log chainTypes.Log) (*UsersRoleUnbound, error)
	ParseRoleUnboundFromReceipt(receipt *chainTypes.Receipt) ([]*UsersRoleUnbound, error)

	// Parser
	ParseSignedUp(log chainTypes.Log) (*UsersSignedUp, error)
	ParseSignedUpFromReceipt(receipt *chainTypes.Receipt) ([]*UsersSignedUp, error)

	// Parser
	ParseTemporaryCreated(log chainTypes.Log) (*UsersTemporaryCreated, error)
	ParseTemporaryCreatedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersTemporaryCreated, error)

	// Parser
	ParseTemporaryUnlocked(log chainTypes.Log) (*UsersTemporaryUnlocked, error)
	ParseTemporaryUnlockedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersTemporaryUnlocked, error)
}

type UsersEventWatcher interface {
	// Watcher
	WatchActionGranted(
		opts *bind.WatchOpts,
		sink chan<- *UsersActionGranted,
		resourceId [][8]byte,
	) (event.Subscription, error)

	// Watcher
	WatchActionRevoked(
		opts *bind.WatchOpts,
		sink chan<- *UsersActionRevoked,
		resourceId [][8]byte,
	) (event.Subscription, error)

	// Watcher
	WatchControllerChanged(
		opts *bind.WatchOpts,
		sink chan<- *UsersControllerChanged,
		userId [][8]byte,
	) (event.Subscription, error)

	// Watcher
	WatchRoleBound(
		opts *bind.WatchOpts,
		sink chan<- *UsersRoleBound,
		resourceId [][8]byte, subject []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchRoleCreation(
		opts *bind.WatchOpts,
		sink chan<- *UsersRoleCreation,
		resourceId [][8]byte,
	) (event.Subscription, error)

	// Watcher
	WatchRoleRemoval(
		opts *bind.WatchOpts,
		sink chan<- *UsersRoleRemoval,
		resourceId [][8]byte,
	) (event.Subscription, error)

	// Watcher
	WatchRoleUnbound(
		opts *bind.WatchOpts,
		sink chan<- *UsersRoleUnbound,
		resourceId [][8]byte, subject []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchSignedUp(
		opts *bind.WatchOpts,
		sink chan<- *UsersSignedUp,
		owner []common.Address,
	) (event.Subscription, error)

	// Watcher
	WatchTemporaryCreated(
		opts *bind.WatchOpts,
		sink chan<- *UsersTemporaryCreated,
		proxy []common.Address, feePayer []common.Address, identityHash []common.Hash,
	) (event.Subscription, error)

	// Watcher
	WatchTemporaryUnlocked(
		opts *bind.WatchOpts,
		sink chan<- *UsersTemporaryUnlocked,
		identityHash []common.Hash, userId [][8]byte,
	) (event.Subscription, error)
}

type usersEvents struct {
	contract *ablbind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersActionGrantedIterator is returned from FilterActionGranted and is used to iterate over the raw logs and unpacked data for ActionGranted events raised by the Users contract.
type UsersActionGrantedIterator struct {
	Evt *UsersActionGranted // Event containing the contract specifics and raw log

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
func (it *UsersActionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersActionGranted)
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
		it.Evt = new(UsersActionGranted)
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
func (it *UsersActionGrantedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersActionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersActionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersActionGranted represents a ActionGranted event raised by the Users contract.
type UsersActionGranted struct {
	ResourceId [8]byte
	RoleName   string
	ActionName string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterActionGranted is a free log retrieval operation binding the contract event 0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9.
//
// Solidity: event ActionGranted(bytes8 indexed resourceId, string roleName, string actionName)
func (_Users *usersEvents) FilterActionGranted(opts *bind.FilterOpts, resourceId [][8]byte) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "ActionGranted", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return &UsersActionGrantedIterator{contract: _Users.contract, event: "ActionGranted", logs: logs, sub: sub}, nil
}

// WatchActionGranted is a free log subscription operation binding the contract event 0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9.
//
// Solidity: event ActionGranted(bytes8 indexed resourceId, string roleName, string actionName)
func (_Users *usersEvents) WatchActionGranted(opts *bind.WatchOpts, sink chan<- *UsersActionGranted, resourceId [][8]byte) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "ActionGranted", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersActionGranted)
				if err := _Users.contract.UnpackLog(evt, "ActionGranted", log); err != nil {
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
func (_Users *usersEvents) ParseActionGranted(log chainTypes.Log) (*UsersActionGranted, error) {
	evt := new(UsersActionGranted)
	if err := _Users.contract.UnpackLog(evt, "ActionGranted", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseActionGrantedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event ActionGranted(bytes8 indexed resourceId, string roleName, string actionName)
func (_Users *usersEvents) ParseActionGrantedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersActionGranted, error) {
	var evts []*UsersActionGranted
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xe03e108285a5925f590b2999b168b964f1fe8f10e0081058fe0c01c38b7e3bd9") {
			evt, err := _Users.ParseActionGranted(*log)
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

// UsersActionRevokedIterator is returned from FilterActionRevoked and is used to iterate over the raw logs and unpacked data for ActionRevoked events raised by the Users contract.
type UsersActionRevokedIterator struct {
	Evt *UsersActionRevoked // Event containing the contract specifics and raw log

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
func (it *UsersActionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersActionRevoked)
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
		it.Evt = new(UsersActionRevoked)
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
func (it *UsersActionRevokedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersActionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersActionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersActionRevoked represents a ActionRevoked event raised by the Users contract.
type UsersActionRevoked struct {
	ResourceId [8]byte
	RoleName   string
	ActionName string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterActionRevoked is a free log retrieval operation binding the contract event 0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430.
//
// Solidity: event ActionRevoked(bytes8 indexed resourceId, string roleName, string actionName)
func (_Users *usersEvents) FilterActionRevoked(opts *bind.FilterOpts, resourceId [][8]byte) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "ActionRevoked", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return &UsersActionRevokedIterator{contract: _Users.contract, event: "ActionRevoked", logs: logs, sub: sub}, nil
}

// WatchActionRevoked is a free log subscription operation binding the contract event 0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430.
//
// Solidity: event ActionRevoked(bytes8 indexed resourceId, string roleName, string actionName)
func (_Users *usersEvents) WatchActionRevoked(opts *bind.WatchOpts, sink chan<- *UsersActionRevoked, resourceId [][8]byte) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "ActionRevoked", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersActionRevoked)
				if err := _Users.contract.UnpackLog(evt, "ActionRevoked", log); err != nil {
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
func (_Users *usersEvents) ParseActionRevoked(log chainTypes.Log) (*UsersActionRevoked, error) {
	evt := new(UsersActionRevoked)
	if err := _Users.contract.UnpackLog(evt, "ActionRevoked", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseActionRevokedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event ActionRevoked(bytes8 indexed resourceId, string roleName, string actionName)
func (_Users *usersEvents) ParseActionRevokedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersActionRevoked, error) {
	var evts []*UsersActionRevoked
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x853537b9aa72341c9043ddf56d5bc062180a8f122964f92d899c2734aad5f430") {
			evt, err := _Users.ParseActionRevoked(*log)
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

// UsersControllerChangedIterator is returned from FilterControllerChanged and is used to iterate over the raw logs and unpacked data for ControllerChanged events raised by the Users contract.
type UsersControllerChangedIterator struct {
	Evt *UsersControllerChanged // Event containing the contract specifics and raw log

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
func (it *UsersControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersControllerChanged)
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
		it.Evt = new(UsersControllerChanged)
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
func (it *UsersControllerChangedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersControllerChanged represents a ControllerChanged event raised by the Users contract.
type UsersControllerChanged struct {
	UserId        [8]byte
	OldController common.Address
	NewController common.Address
	Raw           chainTypes.Log // Blockchain specific contextual infos
}

// FilterControllerChanged is a free log retrieval operation binding the contract event 0xf2fe27b4381d3faea12791cd499b21e10627d1cd2f09101c2c33547ecf711fa1.
//
// Solidity: event ControllerChanged(bytes8 indexed userId, address oldController, address newController)
func (_Users *usersEvents) FilterControllerChanged(opts *bind.FilterOpts, userId [][8]byte) (ablbind.EventIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "ControllerChanged", userIdRule)
	if err != nil {
		return nil, err
	}
	return &UsersControllerChangedIterator{contract: _Users.contract, event: "ControllerChanged", logs: logs, sub: sub}, nil
}

// WatchControllerChanged is a free log subscription operation binding the contract event 0xf2fe27b4381d3faea12791cd499b21e10627d1cd2f09101c2c33547ecf711fa1.
//
// Solidity: event ControllerChanged(bytes8 indexed userId, address oldController, address newController)
func (_Users *usersEvents) WatchControllerChanged(opts *bind.WatchOpts, sink chan<- *UsersControllerChanged, userId [][8]byte) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "ControllerChanged", userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersControllerChanged)
				if err := _Users.contract.UnpackLog(evt, "ControllerChanged", log); err != nil {
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

// ParseControllerChanged is a log parse operation binding the contract event 0xf2fe27b4381d3faea12791cd499b21e10627d1cd2f09101c2c33547ecf711fa1.
//
// Solidity: event ControllerChanged(bytes8 indexed userId, address oldController, address newController)
func (_Users *usersEvents) ParseControllerChanged(log chainTypes.Log) (*UsersControllerChanged, error) {
	evt := new(UsersControllerChanged)
	if err := _Users.contract.UnpackLog(evt, "ControllerChanged", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseControllerChangedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event ControllerChanged(bytes8 indexed userId, address oldController, address newController)
func (_Users *usersEvents) ParseControllerChangedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersControllerChanged, error) {
	var evts []*UsersControllerChanged
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf2fe27b4381d3faea12791cd499b21e10627d1cd2f09101c2c33547ecf711fa1") {
			evt, err := _Users.ParseControllerChanged(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("ControllerChanged event not found")
	}
	return evts, nil
}

// UsersRoleBoundIterator is returned from FilterRoleBound and is used to iterate over the raw logs and unpacked data for RoleBound events raised by the Users contract.
type UsersRoleBoundIterator struct {
	Evt *UsersRoleBound // Event containing the contract specifics and raw log

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
func (it *UsersRoleBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersRoleBound)
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
		it.Evt = new(UsersRoleBound)
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
func (it *UsersRoleBoundIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersRoleBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersRoleBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersRoleBound represents a RoleBound event raised by the Users contract.
type UsersRoleBound struct {
	ResourceId [8]byte
	Subject    common.Address
	RoleName   string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRoleBound is a free log retrieval operation binding the contract event 0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372.
//
// Solidity: event RoleBound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_Users *usersEvents) FilterRoleBound(opts *bind.FilterOpts, resourceId [][8]byte, subject []common.Address) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "RoleBound", resourceIdRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return &UsersRoleBoundIterator{contract: _Users.contract, event: "RoleBound", logs: logs, sub: sub}, nil
}

// WatchRoleBound is a free log subscription operation binding the contract event 0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372.
//
// Solidity: event RoleBound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_Users *usersEvents) WatchRoleBound(opts *bind.WatchOpts, sink chan<- *UsersRoleBound, resourceId [][8]byte, subject []common.Address) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "RoleBound", resourceIdRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersRoleBound)
				if err := _Users.contract.UnpackLog(evt, "RoleBound", log); err != nil {
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
func (_Users *usersEvents) ParseRoleBound(log chainTypes.Log) (*UsersRoleBound, error) {
	evt := new(UsersRoleBound)
	if err := _Users.contract.UnpackLog(evt, "RoleBound", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRoleBoundFromReceipt parses the event from given transaction receipt.
//
// Solidity: event RoleBound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_Users *usersEvents) ParseRoleBoundFromReceipt(receipt *chainTypes.Receipt) ([]*UsersRoleBound, error) {
	var evts []*UsersRoleBound
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf1292e948ea47db3c62a8d29d6c7d2272d236bc10fa66fddb7b2e1570b1b1372") {
			evt, err := _Users.ParseRoleBound(*log)
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

// UsersRoleCreationIterator is returned from FilterRoleCreation and is used to iterate over the raw logs and unpacked data for RoleCreation events raised by the Users contract.
type UsersRoleCreationIterator struct {
	Evt *UsersRoleCreation // Event containing the contract specifics and raw log

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
func (it *UsersRoleCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersRoleCreation)
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
		it.Evt = new(UsersRoleCreation)
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
func (it *UsersRoleCreationIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersRoleCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersRoleCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersRoleCreation represents a RoleCreation event raised by the Users contract.
type UsersRoleCreation struct {
	ResourceId [8]byte
	RoleName   string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRoleCreation is a free log retrieval operation binding the contract event 0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a.
//
// Solidity: event RoleCreation(bytes8 indexed resourceId, string roleName)
func (_Users *usersEvents) FilterRoleCreation(opts *bind.FilterOpts, resourceId [][8]byte) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "RoleCreation", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return &UsersRoleCreationIterator{contract: _Users.contract, event: "RoleCreation", logs: logs, sub: sub}, nil
}

// WatchRoleCreation is a free log subscription operation binding the contract event 0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a.
//
// Solidity: event RoleCreation(bytes8 indexed resourceId, string roleName)
func (_Users *usersEvents) WatchRoleCreation(opts *bind.WatchOpts, sink chan<- *UsersRoleCreation, resourceId [][8]byte) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "RoleCreation", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersRoleCreation)
				if err := _Users.contract.UnpackLog(evt, "RoleCreation", log); err != nil {
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
func (_Users *usersEvents) ParseRoleCreation(log chainTypes.Log) (*UsersRoleCreation, error) {
	evt := new(UsersRoleCreation)
	if err := _Users.contract.UnpackLog(evt, "RoleCreation", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRoleCreationFromReceipt parses the event from given transaction receipt.
//
// Solidity: event RoleCreation(bytes8 indexed resourceId, string roleName)
func (_Users *usersEvents) ParseRoleCreationFromReceipt(receipt *chainTypes.Receipt) ([]*UsersRoleCreation, error) {
	var evts []*UsersRoleCreation
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xc2b6ee9f52b3de3bb61425ddae45ae6d999023b01a3950ee19667b235ed8b23a") {
			evt, err := _Users.ParseRoleCreation(*log)
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

// UsersRoleRemovalIterator is returned from FilterRoleRemoval and is used to iterate over the raw logs and unpacked data for RoleRemoval events raised by the Users contract.
type UsersRoleRemovalIterator struct {
	Evt *UsersRoleRemoval // Event containing the contract specifics and raw log

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
func (it *UsersRoleRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersRoleRemoval)
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
		it.Evt = new(UsersRoleRemoval)
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
func (it *UsersRoleRemovalIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersRoleRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersRoleRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersRoleRemoval represents a RoleRemoval event raised by the Users contract.
type UsersRoleRemoval struct {
	ResourceId [8]byte
	RoleName   string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRoleRemoval is a free log retrieval operation binding the contract event 0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08.
//
// Solidity: event RoleRemoval(bytes8 indexed resourceId, string roleName)
func (_Users *usersEvents) FilterRoleRemoval(opts *bind.FilterOpts, resourceId [][8]byte) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "RoleRemoval", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return &UsersRoleRemovalIterator{contract: _Users.contract, event: "RoleRemoval", logs: logs, sub: sub}, nil
}

// WatchRoleRemoval is a free log subscription operation binding the contract event 0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08.
//
// Solidity: event RoleRemoval(bytes8 indexed resourceId, string roleName)
func (_Users *usersEvents) WatchRoleRemoval(opts *bind.WatchOpts, sink chan<- *UsersRoleRemoval, resourceId [][8]byte) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "RoleRemoval", resourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersRoleRemoval)
				if err := _Users.contract.UnpackLog(evt, "RoleRemoval", log); err != nil {
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
func (_Users *usersEvents) ParseRoleRemoval(log chainTypes.Log) (*UsersRoleRemoval, error) {
	evt := new(UsersRoleRemoval)
	if err := _Users.contract.UnpackLog(evt, "RoleRemoval", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRoleRemovalFromReceipt parses the event from given transaction receipt.
//
// Solidity: event RoleRemoval(bytes8 indexed resourceId, string roleName)
func (_Users *usersEvents) ParseRoleRemovalFromReceipt(receipt *chainTypes.Receipt) ([]*UsersRoleRemoval, error) {
	var evts []*UsersRoleRemoval
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf62cd9d9e6db5c8c5690d7f9adf2224040edc6e63b93e00db8ad092ab6923c08") {
			evt, err := _Users.ParseRoleRemoval(*log)
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

// UsersRoleUnboundIterator is returned from FilterRoleUnbound and is used to iterate over the raw logs and unpacked data for RoleUnbound events raised by the Users contract.
type UsersRoleUnboundIterator struct {
	Evt *UsersRoleUnbound // Event containing the contract specifics and raw log

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
func (it *UsersRoleUnboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersRoleUnbound)
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
		it.Evt = new(UsersRoleUnbound)
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
func (it *UsersRoleUnboundIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersRoleUnboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersRoleUnboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersRoleUnbound represents a RoleUnbound event raised by the Users contract.
type UsersRoleUnbound struct {
	ResourceId [8]byte
	Subject    common.Address
	RoleName   string
	Raw        chainTypes.Log // Blockchain specific contextual infos
}

// FilterRoleUnbound is a free log retrieval operation binding the contract event 0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3.
//
// Solidity: event RoleUnbound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_Users *usersEvents) FilterRoleUnbound(opts *bind.FilterOpts, resourceId [][8]byte, subject []common.Address) (ablbind.EventIterator, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "RoleUnbound", resourceIdRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return &UsersRoleUnboundIterator{contract: _Users.contract, event: "RoleUnbound", logs: logs, sub: sub}, nil
}

// WatchRoleUnbound is a free log subscription operation binding the contract event 0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3.
//
// Solidity: event RoleUnbound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_Users *usersEvents) WatchRoleUnbound(opts *bind.WatchOpts, sink chan<- *UsersRoleUnbound, resourceId [][8]byte, subject []common.Address) (event.Subscription, error) {

	var resourceIdRule []interface{}
	for _, resourceIdItem := range resourceId {
		resourceIdRule = append(resourceIdRule, resourceIdItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "RoleUnbound", resourceIdRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersRoleUnbound)
				if err := _Users.contract.UnpackLog(evt, "RoleUnbound", log); err != nil {
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
func (_Users *usersEvents) ParseRoleUnbound(log chainTypes.Log) (*UsersRoleUnbound, error) {
	evt := new(UsersRoleUnbound)
	if err := _Users.contract.UnpackLog(evt, "RoleUnbound", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseRoleUnboundFromReceipt parses the event from given transaction receipt.
//
// Solidity: event RoleUnbound(bytes8 indexed resourceId, address indexed subject, string roleName)
func (_Users *usersEvents) ParseRoleUnboundFromReceipt(receipt *chainTypes.Receipt) ([]*UsersRoleUnbound, error) {
	var evts []*UsersRoleUnbound
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xe3ee4a62256915250287945244f05e6f185e5b0be4c5f0e85f29e3044f120ce3") {
			evt, err := _Users.ParseRoleUnbound(*log)
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

// UsersSignedUpIterator is returned from FilterSignedUp and is used to iterate over the raw logs and unpacked data for SignedUp events raised by the Users contract.
type UsersSignedUpIterator struct {
	Evt *UsersSignedUp // Event containing the contract specifics and raw log

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
func (it *UsersSignedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersSignedUp)
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
		it.Evt = new(UsersSignedUp)
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
func (it *UsersSignedUpIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersSignedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersSignedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersSignedUp represents a SignedUp event raised by the Users contract.
type UsersSignedUp struct {
	Owner  common.Address
	UserId [8]byte
	Raw    chainTypes.Log // Blockchain specific contextual infos
}

// FilterSignedUp is a free log retrieval operation binding the contract event 0xb60195123201ae92c75c0daf6703a414df492031738ac8d8f5a09c81277c0a87.
//
// Solidity: event SignedUp(address indexed owner, bytes8 userId)
func (_Users *usersEvents) FilterSignedUp(opts *bind.FilterOpts, owner []common.Address) (ablbind.EventIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "SignedUp", ownerRule)
	if err != nil {
		return nil, err
	}
	return &UsersSignedUpIterator{contract: _Users.contract, event: "SignedUp", logs: logs, sub: sub}, nil
}

// WatchSignedUp is a free log subscription operation binding the contract event 0xb60195123201ae92c75c0daf6703a414df492031738ac8d8f5a09c81277c0a87.
//
// Solidity: event SignedUp(address indexed owner, bytes8 userId)
func (_Users *usersEvents) WatchSignedUp(opts *bind.WatchOpts, sink chan<- *UsersSignedUp, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "SignedUp", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersSignedUp)
				if err := _Users.contract.UnpackLog(evt, "SignedUp", log); err != nil {
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

// ParseSignedUp is a log parse operation binding the contract event 0xb60195123201ae92c75c0daf6703a414df492031738ac8d8f5a09c81277c0a87.
//
// Solidity: event SignedUp(address indexed owner, bytes8 userId)
func (_Users *usersEvents) ParseSignedUp(log chainTypes.Log) (*UsersSignedUp, error) {
	evt := new(UsersSignedUp)
	if err := _Users.contract.UnpackLog(evt, "SignedUp", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseSignedUpFromReceipt parses the event from given transaction receipt.
//
// Solidity: event SignedUp(address indexed owner, bytes8 userId)
func (_Users *usersEvents) ParseSignedUpFromReceipt(receipt *chainTypes.Receipt) ([]*UsersSignedUp, error) {
	var evts []*UsersSignedUp
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xb60195123201ae92c75c0daf6703a414df492031738ac8d8f5a09c81277c0a87") {
			evt, err := _Users.ParseSignedUp(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("SignedUp event not found")
	}
	return evts, nil
}

// UsersTemporaryCreatedIterator is returned from FilterTemporaryCreated and is used to iterate over the raw logs and unpacked data for TemporaryCreated events raised by the Users contract.
type UsersTemporaryCreatedIterator struct {
	Evt *UsersTemporaryCreated // Event containing the contract specifics and raw log

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
func (it *UsersTemporaryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersTemporaryCreated)
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
		it.Evt = new(UsersTemporaryCreated)
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
func (it *UsersTemporaryCreatedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersTemporaryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersTemporaryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersTemporaryCreated represents a TemporaryCreated event raised by the Users contract.
type UsersTemporaryCreated struct {
	Proxy        common.Address
	FeePayer     common.Address
	IdentityHash common.Hash
	UserId       [8]byte
	Raw          chainTypes.Log // Blockchain specific contextual infos
}

// FilterTemporaryCreated is a free log retrieval operation binding the contract event 0x67048177171beb7aeaaa2c3ac9ab1bc44f168a64b33ef5d502d607a39bb43480.
//
// Solidity: event TemporaryCreated(address indexed proxy, address indexed feePayer, bytes32 indexed identityHash, bytes8 userId)
func (_Users *usersEvents) FilterTemporaryCreated(opts *bind.FilterOpts, proxy []common.Address, feePayer []common.Address, identityHash []common.Hash) (ablbind.EventIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var feePayerRule []interface{}
	for _, feePayerItem := range feePayer {
		feePayerRule = append(feePayerRule, feePayerItem)
	}
	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "TemporaryCreated", proxyRule, feePayerRule, identityHashRule)
	if err != nil {
		return nil, err
	}
	return &UsersTemporaryCreatedIterator{contract: _Users.contract, event: "TemporaryCreated", logs: logs, sub: sub}, nil
}

// WatchTemporaryCreated is a free log subscription operation binding the contract event 0x67048177171beb7aeaaa2c3ac9ab1bc44f168a64b33ef5d502d607a39bb43480.
//
// Solidity: event TemporaryCreated(address indexed proxy, address indexed feePayer, bytes32 indexed identityHash, bytes8 userId)
func (_Users *usersEvents) WatchTemporaryCreated(opts *bind.WatchOpts, sink chan<- *UsersTemporaryCreated, proxy []common.Address, feePayer []common.Address, identityHash []common.Hash) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var feePayerRule []interface{}
	for _, feePayerItem := range feePayer {
		feePayerRule = append(feePayerRule, feePayerItem)
	}
	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "TemporaryCreated", proxyRule, feePayerRule, identityHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersTemporaryCreated)
				if err := _Users.contract.UnpackLog(evt, "TemporaryCreated", log); err != nil {
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

// ParseTemporaryCreated is a log parse operation binding the contract event 0x67048177171beb7aeaaa2c3ac9ab1bc44f168a64b33ef5d502d607a39bb43480.
//
// Solidity: event TemporaryCreated(address indexed proxy, address indexed feePayer, bytes32 indexed identityHash, bytes8 userId)
func (_Users *usersEvents) ParseTemporaryCreated(log chainTypes.Log) (*UsersTemporaryCreated, error) {
	evt := new(UsersTemporaryCreated)
	if err := _Users.contract.UnpackLog(evt, "TemporaryCreated", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseTemporaryCreatedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event TemporaryCreated(address indexed proxy, address indexed feePayer, bytes32 indexed identityHash, bytes8 userId)
func (_Users *usersEvents) ParseTemporaryCreatedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersTemporaryCreated, error) {
	var evts []*UsersTemporaryCreated
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x67048177171beb7aeaaa2c3ac9ab1bc44f168a64b33ef5d502d607a39bb43480") {
			evt, err := _Users.ParseTemporaryCreated(*log)
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

// UsersTemporaryUnlockedIterator is returned from FilterTemporaryUnlocked and is used to iterate over the raw logs and unpacked data for TemporaryUnlocked events raised by the Users contract.
type UsersTemporaryUnlockedIterator struct {
	Evt *UsersTemporaryUnlocked // Event containing the contract specifics and raw log

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
func (it *UsersTemporaryUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Evt = new(UsersTemporaryUnlocked)
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
		it.Evt = new(UsersTemporaryUnlocked)
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
func (it *UsersTemporaryUnlockedIterator) Event() interface{} {
	return it.Evt
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersTemporaryUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersTemporaryUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersTemporaryUnlocked represents a TemporaryUnlocked event raised by the Users contract.
type UsersTemporaryUnlocked struct {
	IdentityHash common.Hash
	UserId       [8]byte
	NewOwner     common.Address
	Raw          chainTypes.Log // Blockchain specific contextual infos
}

// FilterTemporaryUnlocked is a free log retrieval operation binding the contract event 0xbd9af3a42cb6d433090684c1dd3af90f4ef6f0ae87fb219c7b38d388385c9be7.
//
// Solidity: event TemporaryUnlocked(bytes32 indexed identityHash, bytes8 indexed userId, address newOwner)
func (_Users *usersEvents) FilterTemporaryUnlocked(opts *bind.FilterOpts, identityHash []common.Hash, userId [][8]byte) (ablbind.EventIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "TemporaryUnlocked", identityHashRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return &UsersTemporaryUnlockedIterator{contract: _Users.contract, event: "TemporaryUnlocked", logs: logs, sub: sub}, nil
}

// WatchTemporaryUnlocked is a free log subscription operation binding the contract event 0xbd9af3a42cb6d433090684c1dd3af90f4ef6f0ae87fb219c7b38d388385c9be7.
//
// Solidity: event TemporaryUnlocked(bytes32 indexed identityHash, bytes8 indexed userId, address newOwner)
func (_Users *usersEvents) WatchTemporaryUnlocked(opts *bind.WatchOpts, sink chan<- *UsersTemporaryUnlocked, identityHash []common.Hash, userId [][8]byte) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "TemporaryUnlocked", identityHashRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				evt := new(UsersTemporaryUnlocked)
				if err := _Users.contract.UnpackLog(evt, "TemporaryUnlocked", log); err != nil {
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

// ParseTemporaryUnlocked is a log parse operation binding the contract event 0xbd9af3a42cb6d433090684c1dd3af90f4ef6f0ae87fb219c7b38d388385c9be7.
//
// Solidity: event TemporaryUnlocked(bytes32 indexed identityHash, bytes8 indexed userId, address newOwner)
func (_Users *usersEvents) ParseTemporaryUnlocked(log chainTypes.Log) (*UsersTemporaryUnlocked, error) {
	evt := new(UsersTemporaryUnlocked)
	if err := _Users.contract.UnpackLog(evt, "TemporaryUnlocked", log); err != nil {
		return nil, err
	}
	return evt, nil
}

// ParseTemporaryUnlockedFromReceipt parses the event from given transaction receipt.
//
// Solidity: event TemporaryUnlocked(bytes32 indexed identityHash, bytes8 indexed userId, address newOwner)
func (_Users *usersEvents) ParseTemporaryUnlockedFromReceipt(receipt *chainTypes.Receipt) ([]*UsersTemporaryUnlocked, error) {
	var evts []*UsersTemporaryUnlocked
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xbd9af3a42cb6d433090684c1dd3af90f4ef6f0ae87fb219c7b38d388385c9be7") {
			evt, err := _Users.ParseTemporaryUnlocked(*log)
			if err != nil {
				return nil, err
			}
			evts = append(evts, evt)
		}
	}

	if len(evts) == 0 {
		return nil, errors.New("TemporaryUnlocked event not found")
	}
	return evts, nil
}

// Manager is contract wrapper struct
type UsersContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	UsersCaller
	UsersTransactor
	UsersEvents
}

func NewUsersContract(backend ablbind.ContractBackend) (*UsersContract, error) {
	deployment, exist := backend.Deployment("Users")
	if !exist {
		evmABI, err := abi.JSON(strings.NewReader(UsersABI))
		if err != nil {
			return nil, err
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(UsersAddress),
			common.HexToHash(UsersTxHash),
			new(big.Int).SetBytes(common.HexToHash(UsersCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, "Users", backend)

	contract := &UsersContract{
		Deployment: deployment,
		client:     backend,

		UsersCaller:     &usersCaller{base},
		UsersTransactor: &usersTransactor{base, backend},
		UsersEvents:     &usersEvents{base},
	}

	return contract, nil
}
