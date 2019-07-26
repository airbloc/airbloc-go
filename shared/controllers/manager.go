package controllers

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Manager is contract wrapper struct
type manager struct {
	contract adapter.IControllerRegistryContract
	log      *logger.Logger
}

// NewManager makes new *manager struct
func NewManager(client blockchain.TxClient) adapter.IControllerRegistryManager {
	return &manager{
		contract: adapter.NewControllerRegistryContract(client),
		log:      logger.New("controller-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (manager *manager) Register(ctx context.Context, controllerAddr common.Address) error {
	receipt, err := manager.contract.Register(ctx, controllerAddr)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data controller registered.", logger.Attrs{"controller": event.Controller.Hex()})
	return err
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (manager *manager) Get(controllerAddr common.Address) (types.DataController, error) {
	return manager.contract.Get(controllerAddr)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (manager *manager) Exists(controller common.Address) (bool, error) {
	return manager.contract.Exists(controller)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (manager *manager) IsOwner() (bool, error) {
	return manager.contract.IsOwner()
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (manager *manager) Owner() (common.Address, error) {
	return manager.contract.Owner()
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (manager *manager) RenounceOwnership(ctx context.Context) error {
	receipt, err := manager.contract.RenounceOwnership(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseOwnershipTransferredFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Ownership renounced.", logger.Attrs{
		"prev-owner": event.PreviousOwner.Hex(),
		"new-owner":  event.NewOwner.Hex(),
	})
	return err
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (manager *manager) TransferOwnership(ctx context.Context, newOwner common.Address) error {
	receipt, err := manager.contract.TransferOwnership(ctx, newOwner)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseOwnershipTransferredFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Ownership transferred.", logger.Attrs{
		"prev-owner": event.PreviousOwner.Hex(),
		"new-owner":  event.NewOwner.Hex(),
	})
	return err
}
