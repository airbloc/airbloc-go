package managers

import (
	"context"
	"math/big"

	ablbind "github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/bind/contracts"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

//go:generate mockgen -source data_type_registry.go -destination ./mocks/mock_data_type_registry.go -package mocks IDataTypeRegistryManager

type DataTypeRegistryManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	contracts.DataTypeRegistryCaller

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

	contracts.DataTypeRegistryEventFilterer
	contracts.DataTypeRegistryEventWatcher
}

// dataTypeRegistryManager is contract wrapper struct
type dataTypeRegistryManager struct {
	*contracts.DataTypeRegistryContract
	client ablbind.ContractBackend
	log    *logger.Logger
}

// NewDataTypeRegistryManager makes new *dataTypeRegistryManager struct
func NewDataTypeRegistryManager(backend ablbind.ContractBackend) (DataTypeRegistryManager, error) {
	contract, err := contracts.NewDataTypeRegistryContract(backend)
	if err != nil {
		return nil, err
	}

	return &dataTypeRegistryManager{
		DataTypeRegistryContract: contract,
		client: backend,
		log:    logger.New("data_type_registry"),
	}, nil
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (manager *dataTypeRegistryManager) Register(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	name string,
	schemaHash common.Hash,
) error {
	receipt, err := manager.DataTypeRegistryContract.Register(ctx, opts, name, schemaHash)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.DataTypeRegistryContract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data type registered.", logger.Attrs{"name": evt[0].Name})
	return nil
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (manager *dataTypeRegistryManager) Unregister(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	name string,
) error {
	receipt, err := manager.DataTypeRegistryContract.Unregister(ctx, opts, name)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.DataTypeRegistryContract.ParseUnregistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data type unregistered.", logger.Attrs{"name": evt[0].Name})
	return nil
}
