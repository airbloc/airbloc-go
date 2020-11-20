package managers

import (
	"context"
	"math/big"

	ablbind "github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/contracts"
	"github.com/airbloc/logger"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

//go:generate mockgen -source controller_registry.go -destination ./mocks/mock_controller_registry.go -package mocks IControllerRegistryManager

type ControllerRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	contracts.ControllerRegistryCaller

	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		controllerAddr common.Address,
	) error

	contracts.ControllerRegistryEventFilterer
	contracts.ControllerRegistryEventWatcher
}

// controllerRegistryManager is contract wrapper struct
type controllerRegistryManager struct {
	*contracts.ControllerRegistryContract
	client ablbind.ContractBackend
	log    logger.Logger
}

// NewControllerRegistryManager makes new *controllerRegistryManager struct
func NewControllerRegistryManager(backend ablbind.ContractBackend) (ControllerRegistryManager, error) {
	contract, err := contracts.NewControllerRegistryContract(backend)
	if err != nil {
		return nil, err
	}

	return &controllerRegistryManager{
		ControllerRegistryContract: contract,
		client:                     backend,
		log:                        logger.New("controller_registry"),
	}, nil
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (manager *controllerRegistryManager) Register(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	controllerAddr common.Address,
) error {
	receipt, err := manager.ControllerRegistryContract.Register(ctx, opts, controllerAddr)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ControllerRegistryContract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data controller registered.", logger.Attrs{"controller": evt[0].Controller.Hex()})
	return nil
}
