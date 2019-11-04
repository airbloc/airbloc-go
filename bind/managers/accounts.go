package managers

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	ablbind "github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/bind/contracts"
	types "github.com/airbloc/airbloc-go/bind/types"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source accounts.go -destination ./mocks/mock_accounts.go -package mocks IAccountsManager

type AccountsManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	contracts.AccountsCaller

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

	contracts.AccountsEventFilterer
	contracts.AccountsEventWatcher
}

// accountsManager is contract wrapper struct
type accountsManager struct {
	*contracts.AccountsContract
	client ablbind.ContractBackend
	log    *logger.Logger
}

// NewAccountsManager makes new *accountsManager struct
func NewAccountsManager(backend ablbind.ContractBackend) (AccountsManager, error) {
	contract, err := contracts.NewAccountsContract(backend)
	if err != nil {
		return nil, err
	}

	return &accountsManager{
		AccountsContract: contract,
		client:           backend,
		log:              logger.New("accounts"),
	}, nil
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (manager *accountsManager) Create(
	ctx context.Context,
	opts *ablbind.TransactOpts,
) (
	types.ID,
	error,
) {
	receipt, err := manager.AccountsContract.Create(ctx, opts)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.AccountsContract.ParseSignUpFromReceipt(receipt)
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
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (manager *accountsManager) CreateTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityHash common.Hash,
) (
	types.ID,
	error,
) {
	receipt, err := manager.AccountsContract.CreateTemporary(ctx, opts, identityHash)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.AccountsContract.ParseTemporaryCreatedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary account created.", logger.Attrs{
		"account_id": evt[0].AccountId.Hex(),
		"proxy":      evt[0].Proxy.Hex(),
	})
	return evt[0].AccountId, nil
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (manager *accountsManager) SetController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	controller common.Address,
) error {
	_, err := manager.AccountsContract.SetController(ctx, opts, controller)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	manager.log.Info("Controller changed.", logger.Attrs{"controller": controller.Hex()})
	return nil
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
	receipt, err := manager.AccountsContract.UnlockTemporary(ctx, opts, identityPreimage, newOwner, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.AccountsContract.ParseUnlockedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary account unlocked.", logger.Attrs{
		"account_id": evt[0].AccountId.Hex(),
		"new_owner":  evt[0].NewOwner.Hex(),
	})
	return nil
}
