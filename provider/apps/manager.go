package apps

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/ethereum/go-ethereum/common"
)

// Manager is contract wrapper struct
type Manager struct {
	client   blockchain.TxClient
	contract *adapter.AppRegistry
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) adapter.AppRegistryManager {
	contract := client.GetContract(&adapter.AppRegistry{})
	types.DataId{}
	return &Manager{
		client:   client,
		contract: contract.(*adapter.AppRegistry),
	}
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (manager *Manager) Register(ctx context.Context, appName string) error {
	tx, err := manager.contract.Register(manager.client.Account(), appName)
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
// Solidity: function unregister(string appName) returns()
func (manager *Manager) Unregister(ctx context.Context, appName string) error {
	tx, err := manager.contract.Unregister(manager.client.Account(), appName)
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
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (manager *Manager) Get(appName string) (types.App, error) {
	return manager.contract.Get(nil, appName)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (manager *Manager) TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) error {
	tx, err := manager.contract.TransferAppOwner(manager.client.Account(), appName, newOwner)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseAppOwnerTransferredFromReceipt(receipt)
	return err
}
