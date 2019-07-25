package dataTypes

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
type Manager struct {
	contract adapter.IDataTypeRegistryContract
	log      *logger.Logger
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) adapter.IDataTypeRegistryManager {
	return &Manager{
		contract: adapter.NewDataTypeRegistryContract(client),
		log:      logger.New("data-type-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (manager *Manager) Register(ctx context.Context, name string, schemaHash common.Hash) error {
	receipt, err := manager.contract.Register(ctx, name, schemaHash)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data type registered.", logger.Attrs{"name": event.Name})
	return err
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (manager *Manager) Unregister(ctx context.Context, name string) error {
	receipt, err := manager.contract.Unregister(ctx, name)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseUnregistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data type unregistered.", logger.Attrs{"name": event.Name})
	return err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (manager *Manager) Get(name string) (types.DataType, error) {
	return manager.contract.Get(name)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (manager *Manager) Exists(name string) (bool, error) {
	return manager.contract.Exists(name)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (manager *Manager) IsOwner(name string, owner common.Address) (bool, error) {
	return manager.contract.IsOwner(name, owner)
}
