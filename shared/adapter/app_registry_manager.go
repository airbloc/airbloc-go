package adapter

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/logger"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

type appRegistryManager struct {
	IAppRegistryContract
	log *logger.Logger
}

// NewAppRegistryManager makes new *NewAppRegistryManager struct
func NewAppRegistryManager(client blockchain.TxClient) IAppRegistryManager {
	return &appRegistryManager{
		IAppRegistryContract: client.GetContract(&AppRegistryContract{}).(*AppRegistryContract),
		log:                  logger.New("app-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (manager *appRegistryManager) Register(ctx context.Context, opts *blockchain.TransactOpts, appName string) error {
	receipt, err := manager.IAppRegistryContract.Register(ctx, opts, appName)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IAppRegistryContract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App registered.", logger.Attrs{"name": evt[0].AppName})
	return nil
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (manager *appRegistryManager) Unregister(ctx context.Context, opts *blockchain.TransactOpts, appName string) error {
	receipt, err := manager.IAppRegistryContract.Unregister(ctx, opts, appName)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IAppRegistryContract.ParseUnregistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App unregistered.", logger.Attrs{"name": evt[0].AppName})
	return nil

}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (manager *appRegistryManager) TransferAppOwner(ctx context.Context, opts *blockchain.TransactOpts, appName string, newOwner common.Address) error {
	receipt, err := manager.IAppRegistryContract.TransferAppOwner(ctx, opts, appName, newOwner)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IAppRegistryContract.ParseAppOwnerTransferredFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("App owner transfered.", logger.Attrs{
		"prev-owner": evt[0].OldOwner.Hex(),
		"new-owner":  evt[0].NewOwner.Hex(),
	})
	return nil
}
