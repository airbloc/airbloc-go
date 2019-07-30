package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/ethereum/go-ethereum/event"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Manager is contract wrapper struct
type dataTypeRegistryManager struct {
	contract IDataTypeRegistryContract
	log      *logger.Logger
}

// Address is getter method of DataTypeRegistry.address
func (manager *dataTypeRegistryManager) Address() common.Address {
	return manager.contract.Address()
}

// TxHash is getter method of DataTypeRegistry.txHash
func (manager *dataTypeRegistryManager) TxHash() common.Hash {
	return manager.contract.TxHash()
}

// CreatedAt is getter method of DataTypeRegistry.createdAt
func (manager *dataTypeRegistryManager) CreatedAt() *big.Int {
	return manager.contract.CreatedAt()
}

// NewDataTypeRegistryManager makes new *dataTypeRegistryManager struct
func NewDataTypeRegistryManager(client blockchain.TxClient) IDataTypeRegistryManager {
	return &dataTypeRegistryManager{
		contract: NewDataTypeRegistryContract(client),
		log:      logger.New("data-type-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (manager *dataTypeRegistryManager) Register(ctx context.Context, name string, schemaHash common.Hash) error {
	receipt, err := manager.contract.Register(ctx, name, schemaHash)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data type registered.", logger.Attrs{"name": evt.Name})
	return err
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (manager *dataTypeRegistryManager) Unregister(ctx context.Context, name string) error {
	receipt, err := manager.contract.Unregister(ctx, name)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseUnregistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Data type unregistered.", logger.Attrs{"name": evt.Name})
	return err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (manager *dataTypeRegistryManager) Get(name string) (types.DataType, error) {
	return manager.contract.Get(name)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (manager *dataTypeRegistryManager) Exists(name string) (bool, error) {
	return manager.contract.Exists(name)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (manager *dataTypeRegistryManager) IsOwner(name string, owner common.Address) (bool, error) {
	return manager.contract.IsOwner(name, owner)
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (manager dataTypeRegistryManager) FilterRegistration(opts *bind.FilterOpts) (*DataTypeRegistryRegistrationIterator, error) {
	return manager.contract.FilterRegistration(opts)
}

// WatchRegistration is a free log subscription operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (manager dataTypeRegistryManager) WatchRegistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryRegistration) (event.Subscription, error) {
	return manager.contract.WatchRegistration(opts, sink)
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (manager dataTypeRegistryManager) FilterUnregistration(opts *bind.FilterOpts) (*DataTypeRegistryUnregistrationIterator, error) {
	return manager.contract.FilterUnregistration(opts)
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (manager dataTypeRegistryManager) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *DataTypeRegistryUnregistration) (event.Subscription, error) {
	return manager.contract.WatchUnregistration(opts, sink)
}
