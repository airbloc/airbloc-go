package apps

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
	contract adapter.IAppRegistryContract
	log      *logger.Logger
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) adapter.IAppRegistryManager {
	return &manager{
		contract: adapter.NewAppRegistryContract(client),
		log:      logger.New("app-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (manager *manager) Register(ctx context.Context, appName string) error {
	receipt, err := manager.contract.Register(ctx, appName)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App registered.", logger.Attrs{"name": event.AppName})
	return err
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (manager *manager) Unregister(ctx context.Context, appName string) error {
	receipt, err := manager.contract.Unregister(ctx, appName)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseUnregistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App unregistered.", logger.Attrs{"name": event.AppName})
	return err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (manager *manager) Get(appName string) (types.App, error) {
	return manager.contract.Get(appName)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (manager *manager) TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) error {
	receipt, err := manager.contract.TransferAppOwner(ctx, appName, newOwner)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseAppOwnerTransferredFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App owner transfered.", logger.Attrs{
		"prev-owner": event.OldOwner.Hex(),
		"new-owner":  event.NewOwner.Hex(),
	})
	return err
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (manager *manager) Exists(appName string) (bool, error) {
	return manager.contract.Exists(appName)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (manager *manager) IsOwner(appName string, owner common.Address) (bool, error) {
	return manager.contract.IsOwner(appName, owner)
}
