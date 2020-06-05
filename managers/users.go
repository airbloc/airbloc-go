package managers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/pkg/errors"

	ablbind "github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/contracts"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source users.go -destination ./mocks/mock_users.go -package mocks IUsersManager

type UsersManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	contracts.UsersCaller

	Create(
		ctx context.Context,
		opts *ablbind.TransactOpts,
	) (
		[8]byte,
		error,
	)
	CreateTemporary(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		identityHash common.Hash,
	) (
		[8]byte,
		error,
	)
	SetController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		newController common.Address,
	) error

	UnlockTemporary(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		identityPreimage common.Hash,
		newOwner common.Address,
	) error

	contracts.UsersEventFilterer
	contracts.UsersEventWatcher
}

// usersManager is contract wrapper struct
type usersManager struct {
	*contracts.UsersContract
	client ablbind.ContractBackend
	log    logger.Logger
}

// NewUsersManager makes new *usersManager struct
func NewUsersManager(backend ablbind.ContractBackend) (UsersManager, error) {
	contract, err := contracts.NewUsersContract(backend)
	if err != nil {
		return nil, err
	}

	return &usersManager{
		UsersContract: contract,
		client:        backend,
		log:           logger.New("users"),
	}, nil
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(bytes8)
func (manager *usersManager) Create(
	ctx context.Context,
	opts *ablbind.TransactOpts,
) (
	[8]byte,
	error,
) {
	receipt, err := manager.UsersContract.Create(ctx, opts)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.UsersContract.ParseSignedUpFromReceipt(receipt)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("User created.", logger.Attrs{"user-id": fmt.Sprintf("%x", evt[0].UserId)})
	return evt[0].UserId, nil
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns(bytes8)
func (manager *usersManager) CreateTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityHash common.Hash,
) (
	[8]byte,
	error,
) {
	receipt, err := manager.UsersContract.CreateTemporary(ctx, opts, identityHash)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.UsersContract.ParseTemporaryCreatedFromReceipt(receipt)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary user created.", logger.Attrs{
		"user-id":       fmt.Sprintf("%x", evt[0].UserId),
		"identity-hash": evt[0].IdentityHash.Hex(),
	})
	return evt[0].UserId, nil
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (manager *usersManager) SetController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	newController common.Address,
) error {
	receipt, err := manager.UsersContract.SetController(ctx, opts, newController)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.UsersContract.ParseControllerChangedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Controller changed.", logger.Attrs{
		"user-id":        fmt.Sprintf("%x", evt[0].UserId),
		"old-controller": evt[0].OldController.Hex(),
		"new-controller": evt[0].NewController.Hex(),
	})
	return nil
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x564929bf.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner) returns()
func (manager *usersManager) UnlockTemporary(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	identityPreimage common.Hash,
	newOwner common.Address,
) error {
	receipt, err := manager.UsersContract.UnlockTemporary(ctx, opts, identityPreimage, newOwner)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.UsersContract.ParseTemporaryUnlockedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary user unlocked.", logger.Attrs{
		"user-id":   fmt.Sprintf("%x", evt[0].UserId),
		"new-owner": evt[0].NewOwner.Hex(),
	})
	return nil
}
