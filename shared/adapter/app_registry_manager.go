package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type appRegistryManager struct {
	AppRegistryFilterer
	contract IAppRegistryContract
	log      *logger.Logger
}

// Address is getter method of AppRegistry.address
func (manager *appRegistryManager) Address() common.Address {
	return manager.contract.Address()
}

// TxHash is getter method of AppRegistry.txHash
func (manager *appRegistryManager) TxHash() common.Hash {
	return manager.contract.TxHash()
}

// CreatedAt is getter method of AppRegistry.createdAt
func (manager *appRegistryManager) CreatedAt() *big.Int {
	return manager.contract.CreatedAt()
}

// NewAppRegistryManager makes new *NewAppRegistryManager struct
func NewAppRegistryManager(client blockchain.TxClient) IAppRegistryManager {
	contract := NewAppRegistryContract(client)
	return &appRegistryManager{
		AppRegistryFilterer: contract.Filterer(),
		contract:            contract,
		log:                 logger.New("app-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (manager *appRegistryManager) Register(ctx context.Context, appName string) error {
	receipt, err := manager.contract.Register(ctx, appName)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App registered.", logger.Attrs{"name": evt.AppName})
	return err
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (manager *appRegistryManager) Unregister(ctx context.Context, appName string) error {
	receipt, err := manager.contract.Unregister(ctx, appName)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseUnregistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App unregistered.", logger.Attrs{"name": evt.AppName})
	return err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (manager *appRegistryManager) Get(appName string) (types.App, error) {
	return manager.contract.Get(appName)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (manager *appRegistryManager) TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) error {
	receipt, err := manager.contract.TransferAppOwner(ctx, appName, newOwner)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseAppOwnerTransferredFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App owner transfered.", logger.Attrs{
		"prev-owner": evt.OldOwner.Hex(),
		"new-owner":  evt.NewOwner.Hex(),
	})
	return err
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (manager *appRegistryManager) Exists(appName string) (bool, error) {
	return manager.contract.Exists(appName)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (manager *appRegistryManager) IsOwner(appName string, owner common.Address) (bool, error) {
	return manager.contract.IsOwner(appName, owner)
}
