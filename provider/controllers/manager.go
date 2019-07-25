package controllers

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/ethereum/go-ethereum/common"
)

// Manager is contract wrapper struct
type Manager struct {
	client   blockchain.TxClient
	contract *adapter.ControllerRegistry
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) adapter.ControllerRegistryManager {
	contract := client.GetContract(&adapter.ControllerRegistry{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.ControllerRegistry),
	}
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (manager *Manager) Register(ctx context.Context, controllerAddr common.Address) error {
	tx, err := manager.contract.Register(manager.client.Account(), controllerAddr)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseRegistrationFromReceipt(receipt)
	return err
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (manager *Manager) Get(controllerAddr common.Address) (types.DataController, error) {
	return manager.contract.Get(nil, controllerAddr)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (manager *Manager) Exists(controller common.Address) (bool, error) {
	return manager.contract.Exists(nil, controller)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (manager *Manager) IsOwner() (bool, error) {
	return manager.contract.IsOwner(nil)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (manager *Manager) Owner() (common.Address, error) {
	return manager.contract.Owner(nil)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (manager *Manager) RenounceOwnership(ctx context.Context) error {
	tx, err := manager.contract.RenounceOwnership(manager.client.Account())
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseOwnershipTransferredFromReceipt(receipt)
	return err
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (manager *Manager) TransferOwnership(ctx context.Context, newOwner common.Address) error {
	tx, err := manager.contract.TransferOwnership(manager.client.Account(), newOwner)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseOwnershipTransferredFromReceipt(receipt)
	return err
}
