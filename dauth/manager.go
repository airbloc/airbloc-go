package dauth

import (
	"context"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/common"
	"github.com/pkg/errors"
)

var (
	ErrCollectionNotFound = errors.New("collection not found.")
)

// Manager is user-side DAuth (Data Collection Authentification) manager.
// NOTE that this is for server-side management (user delegate, account proxy);
// it is not supposed to be called by client directly.
type Manager struct {
	ethclient          *blockchain.Client
	collectionRegistry *adapter.CollectionRegistry
}

func NewManager(client *blockchain.Client) *Manager {
	contract := client.GetContract(&adapter.CollectionRegistry{})
	return &Manager{
		ethclient:          client,
		collectionRegistry: contract.(*adapter.CollectionRegistry),
	}
}

// Allow allows data provider to collect certain kinds (Collection) of user's data.
func (manager *Manager) Allow(collectionId common.ID, passwordSig []byte) error {
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

// Allow allows data provider to collect certain kinds (Collection) of user's data.
func (manager *Manager) Deny(collectionId common.ID, passwordSig []byte) error {
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

// Exists checks that given collection is exists.
func (manager *Manager) Exists(collectionId common.ID) (exists bool, err error) {
	return manager.collectionRegistry.Exists(nil, collectionId)
}
