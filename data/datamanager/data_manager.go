package datamanager

import (
	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/pkg/errors"
)

type Manager struct {
	kms       key.Manager
	client    blockchain.TxClient
	metadb    *metadb.Database
	p2p       p2p.Server
	warehouse *warehouse.DataWarehouse
	registry  *adapter.DataRegistry
	batches   *data.BatchManager
}

func NewManager(
	kms key.Manager,
	p2p p2p.Server,
	localDB localdb.Database,
	client blockchain.TxClient,
) *Manager {
	batches := data.NewBatchManager(localDB)
	contract := client.GetContract(&adapter.DataRegistry{})
	return &Manager{
		kms:      kms,
		client:   client,
		p2p:      p2p,
		registry: contract.(*adapter.DataRegistry),
		batches:  batches,
	}
}

func (manager *Manager) Batches() *data.BatchManager {
	return manager.batches
}

func (manager *Manager) decryptData(bundle *data.Bundle, dataID *ablCommon.DataID) (*ablCommon.Data, error) {
	// TODO: Needs paging method
	var encryptedData *ablCommon.EncryptedData
	for _, d := range bundle.Data {
		if d.OwnerAnID == dataID.OwnerID && d.RowID == dataID.RowID {
			encryptedData = d
		}
	}
	if encryptedData == nil {
		return nil, errors.New("cannot find any data matches given data id")
	}
	return manager.kms.DecryptExternalData(encryptedData)
}

func (manager *Manager) Get(dataId string) (*ablCommon.Data, error) {
	id, err := ablCommon.NewDataID(dataId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse data ID %s", dataId)
	}

	bundle, err := manager.warehouse.Get(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId)
	}

	d, err := manager.decryptData(bundle, id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt data %s", id.String())
	}
	return d, nil
}

func (manager *Manager) GetBatch(batch *data.Batch) ([]*ablCommon.Data, error) {
	bundles := make(map[ablCommon.ID]*data.Bundle)
	dataList := make([]*ablCommon.Data, batch.Count)

	for dataId := range batch.Iterator() {
		if _, alreadyFetched := bundles[dataId.BundleID]; !alreadyFetched {
			b, err := manager.warehouse.Get(&dataId)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId.String())
			}
			bundles[dataId.BundleID] = b
		}

		// try to decrypt data using own private key / re-encryption key
		d, err := manager.decryptData(bundles[dataId.BundleID], &dataId)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decrypt data %s", dataId.String())
		}
		dataList = append(dataList, d)
	}
	return dataList, nil
}
