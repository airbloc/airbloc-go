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
type appRegistryManager struct {
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

// NewAppRegistryManager makes new *Manager struct
func NewAppRegistryManager(client blockchain.TxClient) IAppRegistryManager {
	return &appRegistryManager{
		contract: NewAppRegistryContract(client),
		log:      logger.New("app-registry"),
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

// FilterAppOwnerTransferred is a free log retrieval operation binding the contract event 0x9323f5fe9b72ac1fe704b80d8d53e6538f63d0d041068ef274c35b41d1cbc1de.
//
// Solidity: event AppOwnerTransferred(bytes32 indexed hashedAppName, string appName, address indexed oldOwner, address newOwner)
func (manager appRegistryManager) FilterAppOwnerTransferred(opts *bind.FilterOpts, addr []common.Address, oldOwner []common.Address) (*AppRegistryAppOwnerTransferredIterator, error) {
	return manager.contract.FilterAppOwnerTransferred(opts, addr, oldOwner)
}

// WatchAppOwnerTransferred is a free log subscription operation binding the contract event 0x9323f5fe9b72ac1fe704b80d8d53e6538f63d0d041068ef274c35b41d1cbc1de.
//
// Solidity: event AppOwnerTransferred(bytes32 indexed hashedAppName, string appName, address indexed oldOwner, address newOwner)
func (manager appRegistryManager) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AppRegistryAppOwnerTransferred, addr []common.Address, oldOwner []common.Address) (event.Subscription, error) {
	return manager.contract.WatchAppOwnerTransferred(opts, sink, addr, oldOwner)
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xe7e1383b88439b9522e6630da35051999780d58947518c9a3d1620d19b1bc886.
//
// Solidity: event Registration(bytes32 indexed hashedAppName, string appName)
func (manager appRegistryManager) FilterRegistration(opts *bind.FilterOpts, addr []common.Address) (*AppRegistryRegistrationIterator, error) {
	return manager.contract.FilterRegistration(opts, addr)
}

// WatchRegistration is a free log subscription operation binding the contract event 0xe7e1383b88439b9522e6630da35051999780d58947518c9a3d1620d19b1bc886.
//
// Solidity: event Registration(bytes32 indexed hashedAppName, string appName)
func (manager appRegistryManager) WatchRegistration(opts *bind.WatchOpts, sink chan<- *AppRegistryRegistration, addr []common.Address) (event.Subscription, error) {
	return manager.contract.WatchRegistration(opts, sink, addr)
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0xe2d00f1029a39aa8484ea81b8c2794180aa48eb0a1f28c721acdc94789f1d638.
//
// Solidity: event Unregistration(bytes32 indexed hashedAppName, string appName)
func (manager appRegistryManager) FilterUnregistration(opts *bind.FilterOpts, addr []common.Address) (*AppRegistryUnregistrationIterator, error) {
	return manager.contract.FilterUnregistration(opts, addr)
}

// WatchUnregistration is a free log subscription operation binding the contract event 0xe2d00f1029a39aa8484ea81b8c2794180aa48eb0a1f28c721acdc94789f1d638.
//
// Solidity: event Unregistration(bytes32 indexed hashedAppName, string appName)
func (manager appRegistryManager) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *AppRegistryUnregistration, addr []common.Address) (event.Subscription, error) {
	return manager.contract.WatchUnregistration(opts, sink, addr)
}
