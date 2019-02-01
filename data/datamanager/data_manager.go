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
	warehouse *warehouse.DataWarehouse,
) *Manager {
	batches := data.NewBatchManager(localDB)
	contract := client.GetContract(&adapter.DataRegistry{})
	return &Manager{
		kms:       kms,
		client:    client,
		p2p:       p2p,
		warehouse: warehouse,
		registry:  contract.(*adapter.DataRegistry),
		batches:   batches,
	}
}

func (manager *Manager) Batches() *data.BatchManager {
	return manager.batches
}

func (manager *Manager) encrypt(data *ablCommon.Data) (*ablCommon.EncryptedData, error) {
	encryptedPayload, err := manager.kms.Encrypt(data.Payload)
	if err != nil {
		return nil, err
	}
	return &ablCommon.EncryptedData{
		UserId:  data.UserId,
		Payload: encryptedPayload,
	}, nil
}

func (manager *Manager) decrypt(bundle *data.Bundle, dataID *ablCommon.DataId) (*ablCommon.Data, error) {
	// TODO: Needs paging method
	var encryptedData *ablCommon.EncryptedData
	for _, d := range bundle.Data {
		if d.UserId == dataID.UserId && d.RowId == dataID.RowId {
			encryptedData = d
		}
	}
	if encryptedData == nil {
		return nil, errors.New("cannot find any data matches given data id")
	}
	return manager.kms.DecryptData(encryptedData)
}

func (manager *Manager) Get(dataId string) (*data.Bundle, *ablCommon.Data, error) {
	id, err := ablCommon.NewDataId(dataId)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to parse data ID %s", dataId)
	}

	bundle, err := manager.warehouse.Get(id)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId)
	}

	d, err := manager.decrypt(bundle, id)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to decrypt data %s", id.String())
	}
	return bundle, d, nil
}

func (manager *Manager) GetBatch(batch *data.Batch) (map[*data.Bundle][]*ablCommon.Data, error) {
	result := make(map[*data.Bundle][]*ablCommon.Data)
	bundles := make(map[ablCommon.ID]*data.Bundle)

	for dataId := range batch.Iterator() {
		if _, alreadyFetched := bundles[dataId.BundleId]; !alreadyFetched {
			b, err := manager.warehouse.Get(&dataId)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId.String())
			}
			bundles[dataId.BundleId] = b
		}

		// try to decrypt data using own private key / re-encryption key
		d, err := manager.decrypt(bundles[dataId.BundleId], &dataId)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decrypt data %s", dataId.String())
		}
		result[bundles[dataId.BundleId]] = append(result[bundles[dataId.BundleId]], d)
	}
	return result, nil
}
