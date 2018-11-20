package datamanager

import (
	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/pkg/errors"
)

type Manager struct {
	kms       *key.Manager
	client    blockchain.TxClient
	metadb    *metadb.Database
	warehouse *warehouse.DataWarehouse
	registry  *adapter.DataRegistry
	batches   *data.BatchManager
}

func NewManager(
	kms *key.Manager,
	localDB localdb.Database,
	client blockchain.TxClient,
	contract *adapter.DataRegistry,
) (*Manager, error) {
	batches := data.NewBatchManager(localDB)

	return &Manager{
		kms:      kms,
		client:   client,
		registry: contract,
		batches:  batches,
	}, nil
}

func (manager *Manager) Batches() *data.BatchManager {
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
	d, err := manager.kms.DecryptExternalData(encryptedData)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt data %s", dataId)
	}
	return d, nil
}

func (manager *Manager) GetBatch(batch *data.Batch) ([]*ablCommon.Data, error) {
	bundles := make(map[ablCommon.ID]*data.Bundle)
	dataList := make([]*ablCommon.Data, batch.Count)

	for dataId := range batch.Iterator() {
		if _, alreadyFetched := bundles[dataId.BundleID]; !alreadyFetched {
			b, err := manager.warehouse.Get(dataId.BundleID.String())
			if err != nil {
				return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId.String())
			}
			bundles[dataId.BundleID] = b
		}
		encryptedData := bundles[dataId.BundleID].Data[dataId.Index]

		// try to decrypt data using own private key / re-encryption key
		d, err := manager.kms.DecryptExternalData(encryptedData)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decrypt data %s", dataId)
		}
		dataList = append(dataList, d)
	}
	return dataList, nil
}
