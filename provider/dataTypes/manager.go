package dataTypes

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
	contract *adapter.DataTypeRegistry
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) adapter.DataTypeRegistryManager {
	contract := client.GetContract(&adapter.DataTypeRegistry{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.DataTypeRegistry),
	}
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (manager *Manager) Register(ctx context.Context, name string, schemaHash common.Hash) error {
	tx, err := manager.contract.Register(manager.client.Account(), name, schemaHash)
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

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (manager *Manager) Unregister(ctx context.Context, name string) error {
	tx, err := manager.contract.Unregister(manager.client.Account(), name)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseUnregistrationFromReceipt(receipt)
	return err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (manager *Manager) Get(ctx context.Context, name string) (types.DataType, error) {
	return manager.contract.Get(nil, name)
}
