package dauth

import (
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
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

// IsCollectionAllowed returns true if the given user allowed data collection
// of the given collection (data type) through DAuth.
func (manager *Manager) IsCollectionAllowed(collectionId types.ID, accountId types.ID) (bool, error) {
	return manager.collectionRegistry.IsCollectionAllowed(nil, collectionId, accountId)
}
