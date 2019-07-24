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
func NewManager(client blockchain.TxClient) *Manager {
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
func (manager *Manager) Get(ctx context.Context, controllerAddr common.Address) (types.DataController, error) {
	return manager.contract.Get(nil, controllerAddr)
}
