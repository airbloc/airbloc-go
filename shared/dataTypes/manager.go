package dataTypes

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/ethereum/go-ethereum/event"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Manager is contract wrapper struct
type manager struct {
	contract adapter.IDataTypeRegistryContract
	log      *logger.Logger
}

// NewManager makes new *manager struct
func NewManager(client blockchain.TxClient) adapter.IDataTypeRegistryManager {
	return &manager{
		contract: adapter.NewDataTypeRegistryContract(client),
		log:      logger.New("data-type-registry"),
	}
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (manager *manager) Register(ctx context.Context, name string, schemaHash common.Hash) error {
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
func (manager *manager) Unregister(ctx context.Context, name string) error {
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
func (manager *manager) Get(name string) (types.DataType, error) {
	return manager.contract.Get(name)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (manager *manager) Exists(name string) (bool, error) {
	return manager.contract.Exists(name)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (manager *manager) IsOwner(name string, owner common.Address) (bool, error) {
	return manager.contract.IsOwner(name, owner)
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (manager manager) FilterRegistration(opts *bind.FilterOpts) (*adapter.DataTypeRegistryRegistrationIterator, error) {
	return manager.contract.FilterRegistration(opts)
}

// WatchRegistration is a free log subscription operation binding the contract event 0xd510136a132b28d5bccd27cc4dd52d556d9982ab168ba54b1e775d4d0f1ca948.
//
// Solidity: event Registration(string name)
func (manager manager) WatchRegistration(opts *bind.WatchOpts, sink chan<- *adapter.DataTypeRegistryRegistration) (event.Subscription, error) {
	return manager.contract.WatchRegistration(opts, sink)
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (manager manager) FilterUnregistration(opts *bind.FilterOpts) (*adapter.DataTypeRegistryUnregistrationIterator, error) {
	return manager.contract.FilterUnregistration(opts)
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x2c7e9e18beb0594fa2ccaf8412bbe719d47f3c1efb1349e2ba03c1a3e4f64c83.
//
// Solidity: event Unregistration(string name)
func (manager manager) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *adapter.DataTypeRegistryUnregistration) (event.Subscription, error) {
	return manager.contract.WatchUnregistration(opts, sink)
}
