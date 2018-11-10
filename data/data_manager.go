package data

import (
	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type Manager struct {
	kms       *key.Manager
	client    *blockchain.Client
	metadb    *metadb.Database
	warehouse *warehouse.DataWarehouse
	registry  *adapter.DataRegistry
	batches   *BatchManager
}

func NewManager(
	kms *key.Manager,
	client *blockchain.Client,
	localDB localdb.Database,
	address ethCommon.Address,
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

func (manager *Manager) Get(dataId string) (*ablCommon.Data, error) {
	id, err := ablCommon.NewDataID(dataId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse data ID %s", dataId)
	}

	bundle, err := manager.warehouse.Get(id.BundleID.String())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId)
	}
	encryptedData := bundle.Data[id.Index]

	// try to decrypt data using own private key / re-encryption key
	data, err := manager.kms.DecryptExternalData(encryptedData)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt data %s", dataId)
	}
	return data, nil
}
