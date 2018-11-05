package data

import (
	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type Manager struct {
	kms      *key.Manager
	client   *blockchain.Client
	registry *adapter.DataRegistry
	batches  *BatchManager
}

func NewManager(
	kms *key.Manager,
	client *blockchain.Client,
	localDB localdb.Database,
	address common.Address,
) (*Manager, error) {

	registry, err := adapter.NewDataRegistry(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind to DataRegistry")
	}

	batches := NewBatchManager(localDB)

	return &Manager{
		kms:      kms,
		client:   client,
		registry: registry,
		batches:  batches,
	}, nil
}

func (manager *Manager) Batches() *BatchManager {
	return manager.batches
}
