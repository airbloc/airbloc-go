package adapter

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/logger"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

type controllerRegistryManager struct {
	IControllerRegistryContract
	log *logger.Logger
}

// NewControllerRegistryManager makes new *controllerRegistryManager struct
func NewControllerRegistryManager(client blockchain.TxClient) IControllerRegistryManager {
	return &controllerRegistryManager{
		IControllerRegistryContract: client.GetContract(&ControllerRegistryContract{}).(*ControllerRegistryContract),
		log:                         logger.New("controller-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (manager *controllerRegistryManager) Register(ctx context.Context, controllerAddr common.Address) error {
	receipt, err := manager.IControllerRegistryContract.Register(ctx, controllerAddr)
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
