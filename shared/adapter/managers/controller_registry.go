package managers

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	wrappers "github.com/airbloc/airbloc-go/shared/adapter/wrappers"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source controller_registry.go -destination ./mocks/mock_controller_registry.go -package mocks IControllerRegistryManager

type IControllerRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	wrappers.IControllerRegistryCalls

	// Transact methods
	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		controllerAddr common.Address,
	) error

	// Event methods
	wrappers.IControllerRegistryFilterer
	wrappers.IControllerRegistryWatcher
}

// controllerRegistryManager is contract wrapper struct
type controllerRegistryManager struct {
	wrappers.IControllerRegistryContract
	client ablbind.ContractBackend
	log    *logger.Logger
}

// NewControllerRegistryManager makes new *controllerRegistryManager struct
func NewControllerRegistryManager(client ablbind.ContractBackend, contract interface{}) interface{} {
	return &controllerRegistryManager{
		IControllerRegistryContract: contract.(*wrappers.ControllerRegistryContract),
		client: client,
		log:    logger.New("controllerRegistry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (manager *controllerRegistryManager) Register(ctx context.Context, opts *ablbind.TransactOpts, controllerAddr common.Address) error {
	receipt, err := manager.IControllerRegistryContract.Register(ctx, opts, controllerAddr)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IControllerRegistryContract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data controller registered.", logger.Attrs{"controller": evt[0].Controller.Hex()})
	return nil
}
