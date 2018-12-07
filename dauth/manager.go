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
func (manager *Manager) Allow(ctx context.Context, collectionId common.ID, passwordSig []byte) error {
	if exists, err := manager.Exists(collectionId); err != nil {
		return errors.Wrap(err, "failed to check collection existence")
	} else if !exists {
		return ErrCollectionNotFound
	}

	tx, err := manager.collectionRegistry.AllowByPassword(manager.ethclient.Account(), collectionId, passwordSig)
	if err != nil {
		return errors.Wrap(err, "failed to transact AllowByPassword")
	}
	if _, err := manager.ethclient.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "transaction execution failed")
	}
	return nil
}

// AllowByDelegate
func (manager *Manager) AllowByDelegate(ctx context.Context, collectionId common.ID, accountId common.ID) error {
	if exists, err := manager.Exists(collectionId); err != nil {
		return errors.Wrap(err, "failed to check collection existence")
	} else if !exists {
		return ErrCollectionNotFound
	}

	tx, err := manager.collectionRegistry.AllowByDelegate(manager.ethclient.Account(), collectionId, accountId)
	if err != nil {
		return errors.Wrap(err, "failed to transact AllowByDelegate")
	}
	if _, err := manager.ethclient.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "transaction execution failed")
	}
	return nil
}

// Allow allows data provider to collect certain kinds (Collection) of user's data.
func (manager *Manager) Deny(ctx context.Context, collectionId common.ID, passwordSig []byte) error {
	if exists, err := manager.Exists(collectionId); err != nil {
		return errors.Wrap(err, "failed to check collection existence")
	} else if !exists {
		return ErrCollectionNotFound
	}

	tx, err := manager.collectionRegistry.AllowByPassword(manager.ethclient.Account(), collectionId, passwordSig)
	if err != nil {
		return errors.Wrap(err, "failed to transact AllowByPassword")
	}
	if _, err := manager.ethclient.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "transaction execution failed")
	}
	return nil
}

// AllowByDelegate
func (manager *Manager) DenyByDelegate(ctx context.Context, collectionId common.ID, accountId common.ID) error {
	if exists, err := manager.Exists(collectionId); err != nil {
		return errors.Wrap(err, "failed to check collection existence")
	} else if !exists {
		return ErrCollectionNotFound
	}

	tx, err := manager.collectionRegistry.DenyByDelegate(manager.ethclient.Account(), collectionId, accountId)
	if err != nil {
		return errors.Wrap(err, "failed to transact DenyByDelegate")
	}
	if _, err := manager.ethclient.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "transaction execution failed")
	}
	return nil
}

// Exists checks that given collection is exists.
func (manager *Manager) Exists(collectionId common.ID) (exists bool, err error) {
	return manager.collectionRegistry.Exists(nil, collectionId)
}
