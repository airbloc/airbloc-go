package managers

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	types "github.com/airbloc/airbloc-go/shared/adapter/types"
	wrappers "github.com/airbloc/airbloc-go/shared/adapter/wrappers"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source accounts.go -destination ./mocks/mock_accounts.go -package mocks IAccountsManager

type IAccountsManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	wrappers.IAccountsCalls

	// Transact methods
	Create(
		ctx context.Context,
		opts *ablbind.TransactOpts,
	) (
		types.ID,
		error,
	)
	CreateTemporary(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		identityHash common.Hash,
	) (
		types.ID,
		error,
	)
	SetController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		controller common.Address,
	) error

	UnlockTemporary(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		identityPreimage common.Hash,
		newOwner common.Address,
		passwordSignature []byte,
	) error

	// Event methods
	wrappers.IAccountsFilterer
	wrappers.IAccountsWatcher
}

// accountsManager is contract wrapper struct
type accountsManager struct {
	wrappers.IAccountsContract
	client ablbind.ContractBackend
	log    *logger.Logger
}

// NewAccountsManager makes new *accountsManager struct
func NewAccountsManager(client ablbind.ContractBackend, contract interface{}) interface{} {
	return &accountsManager{
		IAccountsContract: contract.(*wrappers.AccountsContract),
		client:            client,
		log:               logger.New("accounts"),
	}
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (manager *accountsManager) Create(ctx context.Context, opts *ablbind.TransactOpts) (types.ID, error) {
	receipt, err := manager.IAccountsContract.Create(ctx, opts)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IAccountsContract.ParseSignUpFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Account created.", logger.Attrs{
		"account_id": evt[0].AccountId.Hex(),
		"owner":      evt[0].Owner.Hex(),
	})
	return evt[0].AccountId, nil
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns()
func (manager *accountsManager) CreateTemporary(ctx context.Context, opts *ablbind.TransactOpts, identityHash common.Hash) (types.ID, error) {
	receipt, err := manager.IAccountsContract.CreateTemporary(ctx, opts, identityHash)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IAccountsContract.ParseTemporaryCreatedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary account created.", logger.Attrs{
		"account_id": evt[0].AccountId.Hex(),
		"proxy":      evt[0].Proxy.Hex(),
	})
	return evt[0].AccountId, nil
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (manager *accountsManager) UnlockTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityPreimage common.Hash,
	newOwner common.Address,
	passwordSignature []byte,
) error {
	receipt, err := manager.IAccountsContract.UnlockTemporary(ctx, opts, identityPreimage, newOwner, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IAccountsContract.ParseUnlockedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary account unlocked.", logger.Attrs{
		"account_id": evt[0].AccountId.Hex(),
		"new_owner":  evt[0].NewOwner.Hex(),
	})
	return nil
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (manager *accountsManager) SetController(ctx context.Context, opts *ablbind.TransactOpts, controller common.Address) error {
	_, err := manager.IAccountsContract.SetController(ctx, opts, controller)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	manager.log.Info("Controller changed.", logger.Attrs{"controller": controller.Hex()})
	return nil
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (manager *accountsManager) GetAccountId(owner common.Address) (types.ID, error) {
	if owner == (common.Address{}) {
		return manager.IAccountsContract.GetAccountId(manager.client.Transactor(context.Background()).From)
	} else {
		return manager.IAccountsContract.GetAccountId(owner)
	}
}