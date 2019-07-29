package apps

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

// FilterAppOwnerTransferred is a free log retrieval operation binding the contract event 0x9323f5fe9b72ac1fe704b80d8d53e6538f63d0d041068ef274c35b41d1cbc1de.
//
// Solidity: event AppOwnerTransferred(bytes32 indexed hashedAppName, string appName, address indexed oldOwner, address newOwner)
func (manager manager) FilterAppOwnerTransferred(opts *bind.FilterOpts, hashedAppName []common.Hash, oldOwner []common.Address) (*adapter.AppRegistryAppOwnerTransferredIterator, error) {
	return manager.contract.FilterAppOwnerTransferred(opts, hashedAppName, oldOwner)
}

// WatchAppOwnerTransferred is a free log subscription operation binding the contract event 0x9323f5fe9b72ac1fe704b80d8d53e6538f63d0d041068ef274c35b41d1cbc1de.
//
// Solidity: event AppOwnerTransferred(bytes32 indexed hashedAppName, string appName, address indexed oldOwner, address newOwner)
func (manager manager) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryAppOwnerTransferred, hashedAppName []common.Hash, oldOwner []common.Address) (event.Subscription, error) {
	return manager.contract.WatchAppOwnerTransferred(opts, sink, hashedAppName, oldOwner)
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xe7e1383b88439b9522e6630da35051999780d58947518c9a3d1620d19b1bc886.
//
// Solidity: event Registration(bytes32 indexed hashedAppName, string appName)
func (manager manager) FilterRegistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*adapter.AppRegistryRegistrationIterator, error) {
	return manager.contract.FilterRegistration(opts, hashedAppName)
}

// WatchRegistration is a free log subscription operation binding the contract event 0xe7e1383b88439b9522e6630da35051999780d58947518c9a3d1620d19b1bc886.
//
// Solidity: event Registration(bytes32 indexed hashedAppName, string appName)
func (manager manager) WatchRegistration(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryRegistration, hashedAppName []common.Hash) (event.Subscription, error) {
	return manager.contract.WatchRegistration(opts, sink, hashedAppName)
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0xe2d00f1029a39aa8484ea81b8c2794180aa48eb0a1f28c721acdc94789f1d638.
//
// Solidity: event Unregistration(bytes32 indexed hashedAppName, string appName)
func (manager manager) FilterUnregistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*adapter.AppRegistryUnregistrationIterator, error) {
	return manager.contract.FilterUnregistration(opts, hashedAppName)
}

// WatchUnregistration is a free log subscription operation binding the contract event 0xe2d00f1029a39aa8484ea81b8c2794180aa48eb0a1f28c721acdc94789f1d638.
//
// Solidity: event Unregistration(bytes32 indexed hashedAppName, string appName)
func (manager manager) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryUnregistration, hashedAppName []common.Hash) (event.Subscription, error) {
	return manager.contract.WatchUnregistration(opts, sink, hashedAppName)
}
