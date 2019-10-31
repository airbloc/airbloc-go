package managers

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	wrappers "github.com/airbloc/airbloc-go/shared/adapter/wrappers"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source data_type_registry.go -destination ./mocks/mock_data_type_registry.go -package mocks IDataTypeRegistryManager

type IDataTypeRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	wrappers.IDataTypeRegistryCalls

	// Transact methods
	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		name string,
		schemaHash common.Hash,
	) error

	Unregister(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		name string,
	) error

	// Event methods
	wrappers.IDataTypeRegistryFilterer
	wrappers.IDataTypeRegistryWatcher
}

// dataTypeRegistryManager is contract wrapper struct
type dataTypeRegistryManager struct {
	wrappers.IDataTypeRegistryContract
	client ablbind.ContractBackend
	log    *logger.Logger
}

// NewDataTypeRegistryManager makes new *dataTypeRegistryManager struct
func NewDataTypeRegistryManager(client ablbind.ContractBackend, contract interface{}) interface{} {
	return &dataTypeRegistryManager{
		IDataTypeRegistryContract: contract.(*wrappers.DataTypeRegistryContract),
		client: client,
		log:    logger.New("dataTypeRegistry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (manager *dataTypeRegistryManager) Register(ctx context.Context, opts *ablbind.TransactOpts, name string, schemaHash common.Hash) error {
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
func (manager *dataTypeRegistryManager) Unregister(ctx context.Context, opts *ablbind.TransactOpts, name string) error {
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
