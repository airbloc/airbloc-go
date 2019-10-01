package adapter

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/logger"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

type dataTypeRegistryManager struct {
	IDataTypeRegistryContract
	log *logger.Logger
}

// NewDataTypeRegistryManager makes new *dataTypeRegistryManager struct
func NewDataTypeRegistryManager(client blockchain.TxClient) IDataTypeRegistryManager {
	return &dataTypeRegistryManager{
		IDataTypeRegistryContract: client.GetContract(&DataTypeRegistryContract{}).(*DataTypeRegistryContract),
		log: logger.New("data-type-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (manager *dataTypeRegistryManager) Register(ctx context.Context, opts *blockchain.TransactOpts, name string, schemaHash common.Hash) error {
	receipt, err := manager.IDataTypeRegistryContract.Register(ctx, opts, name, schemaHash)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IDataTypeRegistryContract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data type registered.", logger.Attrs{"name": evt[0].Name})
	return nil
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (manager *dataTypeRegistryManager) Unregister(ctx context.Context, opts *blockchain.TransactOpts, name string) error {
	receipt, err := manager.IDataTypeRegistryContract.Unregister(ctx, opts, name)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IDataTypeRegistryContract.ParseUnregistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data type unregistered.", logger.Attrs{"name": evt[0].Name})
	return nil
}
