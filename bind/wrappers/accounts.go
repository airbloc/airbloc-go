package wrappers

import (
	"context"
	"math/big"
	"strings"

	ablbind "github.com/airbloc/airbloc-go/bind"
	contracts "github.com/airbloc/airbloc-go/bind/contracts"
	types "github.com/airbloc/airbloc-go/bind/types"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source accounts.go -destination ./mocks/mock_accounts.go -package mocks IAccountsContract

type IAccountsCalls interface {
	Exists(
		accountId types.ID,
	) (
		bool,
		error,
	)
	GetAccount(
		accountId types.ID,
	) (
		types.Account,
		error,
	)
	GetAccountByIdentityHash(
		identityHash common.Hash,
	) (
		types.Account,
		error,
	)
	GetAccountId(
		sender common.Address,
	) (
		types.ID,
		error,
	)
	GetAccountIdByIdentityHash(
		identityHash common.Hash,
	) (
		types.ID,
		error,
	)
	GetAccountIdFromSignature(
		messageHash common.Hash,
		signature []byte,
	) (
		types.ID,
		error,
	)
	IsControllerOf(
		sender common.Address,
		accountId types.ID,
	) (
		bool,
		error,
	)
	IsTemporary(
		accountId types.ID,
	) (
		bool,
		error,
	)
}

type IAccountsTransacts interface {
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

type IAccountsEvents interface {
	IAccountsFilterer
	IAccountsParser
	IAccountsWatcher
}

type IAccountsFilterer interface {
	FilterSignUp(
		opts *bind.FilterOpts,
		owner []common.Address,

	) (ablbind.EventIterator, error)
	FilterTemporaryCreated(
		opts *bind.FilterOpts,
		proxy []common.Address,
		identityHash []common.Hash,

	) (ablbind.EventIterator, error)
	FilterUnlocked(
		opts *bind.FilterOpts,
		identityHash []common.Hash,
		accountId []types.ID,

	) (ablbind.EventIterator, error)
}

type IAccountsParser interface {
	ParseSignUp(log chainTypes.Log) (*contracts.AccountsSignUp, error)
	ParseSignUpFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.AccountsSignUp, error)
	ParseTemporaryCreated(log chainTypes.Log) (*contracts.AccountsTemporaryCreated, error)
	ParseTemporaryCreatedFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.AccountsTemporaryCreated, error)
	ParseUnlocked(log chainTypes.Log) (*contracts.AccountsUnlocked, error)
	ParseUnlockedFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.AccountsUnlocked, error)
}

type IAccountsWatcher interface {
	WatchSignUp(
		opts *bind.WatchOpts,
		sink chan<- *contracts.AccountsSignUp,
		owner []common.Address,

	) (event.Subscription, error)
	WatchTemporaryCreated(
		opts *bind.WatchOpts,
		sink chan<- *contracts.AccountsTemporaryCreated,
		proxy []common.Address,
		identityHash []common.Hash,

	) (event.Subscription, error)
	WatchUnlocked(
		opts *bind.WatchOpts,
		sink chan<- *contracts.AccountsUnlocked,
		identityHash []common.Hash,
		accountId []types.ID,

	) (event.Subscription, error)
}

type IAccountsContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IAccountsCalls
	IAccountsTransacts
	IAccountsEvents
}

// Manager is contract wrapper struct
type AccountsContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	contracts.AccountsCaller
	contracts.AccountsFilterer
	contracts.AccountsTransactor
}

func NewAccountsContract(deployment ablbind.Deployment, backend ablbind.ContractBackend) interface{} {
	if deployment.Address() == (common.Address{}) {
		evmABI, err := abi.JSON(strings.NewReader(contracts.AccountsABI))
		if err != nil {
			panic(err)
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(contracts.AccountsAddress),
			common.HexToHash(contracts.AccountsTxHash),
			new(big.Int).SetBytes(common.HexToHash(contracts.AccountsCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, backend)

	contract := &AccountsContract{
		Deployment: deployment,
		client:     backend,

		AccountsCaller:     contracts.NewAccountsCaller(base),
		AccountsTransactor: contracts.NewAccountsTransactor(base),
		AccountsFilterer:   contracts.NewAccountsFilterer(base),
	}

	return contract
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 accountId) constant returns(bool)
func (c *AccountsContract) Exists(
	accountId types.ID,
) (

	bool,
	error,
) {
	return c.AccountsCaller.Exists(nil, accountId)
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (c *AccountsContract) GetAccount(
	accountId types.ID,
) (

	types.Account,
	error,
) {
	return c.AccountsCaller.GetAccount(nil, accountId)
}

// GetAccountByIdentityHash is a free data retrieval call binding the contract method 0xc75aeb7e.
//
// Solidity: function getAccountByIdentityHash(bytes32 identityHash) constant returns((address,uint8,address,address))
func (c *AccountsContract) GetAccountByIdentityHash(
	identityHash common.Hash,
) (

	types.Account,
	error,
) {
	return c.AccountsCaller.GetAccountByIdentityHash(nil, identityHash)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (c *AccountsContract) GetAccountId(
	sender common.Address,
) (

	types.ID,
	error,
) {
	return c.AccountsCaller.GetAccountId(nil, sender)
}

// GetAccountIdByIdentityHash is a free data retrieval call binding the contract method 0x793d5046.
//
// Solidity: function getAccountIdByIdentityHash(bytes32 identityHash) constant returns(bytes8)
func (c *AccountsContract) GetAccountIdByIdentityHash(
	identityHash common.Hash,
) (

	types.ID,
	error,
) {
	return c.AccountsCaller.GetAccountIdByIdentityHash(nil, identityHash)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (c *AccountsContract) GetAccountIdFromSignature(
	messageHash common.Hash,
	signature []byte,
) (

	types.ID,
	error,
) {
	return c.AccountsCaller.GetAccountIdFromSignature(nil, messageHash, signature)
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address sender, bytes8 accountId) constant returns(bool)
func (c *AccountsContract) IsControllerOf(
	sender common.Address,
	accountId types.ID,
) (

	bool,
	error,
) {
	return c.AccountsCaller.IsControllerOf(nil, sender, accountId)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 accountId) constant returns(bool)
func (c *AccountsContract) IsTemporary(
	accountId types.ID,
) (

	bool,
	error,
) {
	return c.AccountsCaller.IsTemporary(nil, accountId)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (c *AccountsContract) Create(
	ctx context.Context,
	opts *ablbind.TransactOpts,
) (*chainTypes.Receipt, error) {
	tx, err := c.AccountsTransactor.Create(c.client.Transactor(ctx, opts))
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (c *AccountsContract) CreateTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityHash common.Hash,
) (*chainTypes.Receipt, error) {
	tx, err := c.AccountsTransactor.CreateTemporary(c.client.Transactor(ctx, opts), identityHash)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (c *AccountsContract) SetController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	controller common.Address,
) (*chainTypes.Receipt, error) {
	tx, err := c.AccountsTransactor.SetController(c.client.Transactor(ctx, opts), controller)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (c *AccountsContract) UnlockTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityPreimage common.Hash,
	newOwner common.Address,
	passwordSignature []byte,
) (*chainTypes.Receipt, error) {
	tx, err := c.AccountsTransactor.UnlockTemporary(c.client.Transactor(ctx, opts), identityPreimage, newOwner, passwordSignature)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
