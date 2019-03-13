package dauth

import (
	"context"
	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/pkg/errors"
)

var (
	ErrCollectionNotFound = errors.New("collection not found.")
)

// Manager is user-side DAuth (Data Collection Authentification) manager.
// NOTE that this is for server-side management (user delegate, account proxy);
// it is not supposed to be called by client directly.
type Manager struct {
	ethclient          blockchain.TxClient
	collectionRegistry *adapter.CollectionRegistry
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.CollectionRegistry{})
	return &Manager{
		ethclient:          client,
		collectionRegistry: contract.(*adapter.CollectionRegistry),
	}
}

// Allow allows data provider to collect certain kinds (Collection) of user's data.
func (manager *Manager) Allow(collectionId types.ID, passwordSig []byte) error {
	if exists, err := manager.Exists(collectionId); err != nil {
		return errors.Wrap(err, "failed to check collection existence")
	} else if !exists {
		return ErrCollectionNotFound
	}

	tx, err := manager.collectionRegistry.AllowByPassword(manager.ethclient.Account(), collectionId, passwordSig)
	if err != nil {
		return errors.Wrap(err, "failed to transact AllowByPassword")
	}
	if _, err := manager.ethclient.WaitMined(context.Background(), tx); err != nil {
		return errors.Wrap(err, "transaction execution failed")
	}
	return nil
}

// AllowByDelegate
func (manager *Manager) AllowByDelegate(collectionId types.ID, accountId types.ID) error {
	if exists, err := manager.Exists(collectionId); err != nil {
		return errors.Wrap(err, "failed to check collection existence")
	} else if !exists {
		return ErrCollectionNotFound
	}

	tx, err := manager.collectionRegistry.AllowByDelegate(manager.ethclient.Account(), collectionId, accountId)
	if err != nil {
		return errors.Wrap(err, "failed to transact AllowByDelegate")
	}
	if _, err := manager.ethclient.WaitMined(context.Background(), tx); err != nil {
		return errors.Wrap(err, "transaction execution failed")
	}
	return nil
}

// Allow allows data provider to collect certain kinds (Collection) of user's data.
func (manager *Manager) Deny(collectionId types.ID, passwordSig []byte) error {
	if exists, err := manager.Exists(collectionId); err != nil {
		return errors.Wrap(err, "failed to check collection existence")
	} else if !exists {
		return ErrCollectionNotFound
	}

	tx, err := manager.collectionRegistry.DenyByPassword(manager.ethclient.Account(), collectionId, passwordSig)
	if err != nil {
		return errors.Wrap(err, "failed to transact AllowByPassword")
	}
	if _, err := manager.ethclient.WaitMined(context.Background(), tx); err != nil {
		return errors.Wrap(err, "transaction execution failed")
	}
	return nil
}

// AllowByDelegate
func (manager *Manager) DenyByDelegate(collectionId types.ID, accountId types.ID) error {
	if exists, err := manager.Exists(collectionId); err != nil {
		return errors.Wrap(err, "failed to check collection existence")
	} else if !exists {
		return ErrCollectionNotFound
	}

	tx, err := manager.collectionRegistry.DenyByDelegate(manager.ethclient.Account(), collectionId, accountId)
	if err != nil {
		return errors.Wrap(err, "failed to transact DenyByDelegate")
	}
	if _, err := manager.ethclient.WaitMined(context.Background(), tx); err != nil {
		return errors.Wrap(err, "transaction execution failed")
	}
	return nil
}

// IsCollectionAllowed returns true if the given user allowed data collection
// of the given collection (data type) through DAuth.
func (manager *Manager) IsCollectionAllowed(collectionId types.ID, accountId types.ID) (bool, error) {
	return manager.collectionRegistry.IsCollectionAllowed(nil, collectionId, accountId)
}

// Exists checks that given collection is exists.
func (manager *Manager) Exists(collectionId types.ID) (exists bool, err error) {
	return manager.collectionRegistry.Exists(nil, collectionId)
}
