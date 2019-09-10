package data

import (
	"context"

	"github.com/airbloc/airbloc-go/provider/data/batch"
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/database/localdb"
	"github.com/airbloc/airbloc-go/shared/database/metadb"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/p2p"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type Manager struct {
	kms       key.Manager
	client    blockchain.TxClient
	metadb    metadb.Database
	p2p       p2p.Server
	warehouse *warehouse.Manager
	//registry  *adapter.DataRegistry
	dataTypes adapter.IDataTypeRegistryContract
	batches   *batch.BatchManager
}

func NewManager(
	kms key.Manager,
	p2p p2p.Server,
	metaDB metadb.Database,
	localDB localdb.Database,
	client blockchain.TxClient,
	warehouse *warehouse.Manager,
) *Manager {
	batches := batch.NewBatchManager(localDB)
	//contract := client.GetContract(&adapter.DataRegistry{})
	return &Manager{
		kms:       kms,
		client:    client,
		p2p:       p2p,
		warehouse: warehouse,
		//registry:  contract.(*adapter.DataRegistry),
		dataTypes: adapter.NewDataTypeRegistryContract(client),
		batches:   batches,
		metadb:    metadb.NewModel(metaDB, "bundles"),
	}
}

func (manager *Manager) Batches() *batch.BatchManager {
	return manager.batches
}

func (manager *Manager) decrypt(bundleData *warehouse.Bundle, dataId types.DataId) (data, error) {
	// prevent runtime error
	if uint32(len(bundleData.Data[dataId.UserId()])) < dataId.RowId().Uint32() {
		return data{}, errors.New("data does not exists")
	}

	encryptedData := bundleData.Data[dataId.UserId()][dataId.RowId().Uint32()]
	d, err := manager.kms.DecryptData(encryptedData)
	if err != nil {
		return data{}, errors.Wrapf(err, "failed to decrypt data %s", dataId.String())
	}

	return data{
		CollectionId: bundleData.Collection,
		UserId:       d.UserId,
		IngestedAt:   bundleData.IngestedAt,
		Payload:      d.Payload,
	}, nil
}

func (manager *Manager) encrypt(data *types.Data) (*types.EncryptedData, error) {
	encryptedPayload, err := manager.kms.Encrypt(data.Payload)
	if err != nil {
		return nil, err
	}
	return &types.EncryptedData{
		UserId:  data.UserId,
		Payload: encryptedPayload,
	}, nil
}

func (manager *Manager) Get(dataId types.DataId) (data, error) {
	bundleData, err := manager.warehouse.Get(dataId)
	if err != nil {
		return data{}, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId)
	}
	return manager.decrypt(bundleData, dataId)
}

func (manager *Manager) GetBatch(batchInfo *batch.Batch) ([]data, error) {
	var err error

	batchData := make([]data, batchInfo.Count)
	bundles := make(map[types.ID]*warehouse.Bundle, batchInfo.Count)

	index := 0
	for dataId := range batchInfo.Iterator() {
		bundleId := dataId.BundleId()

		if _, alreadyFetched := bundles[bundleId]; !alreadyFetched {
			bundles[bundleId], err = manager.warehouse.Get(dataId)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId.String())
			}
		}

		batchData[index], err = manager.decrypt(bundles[bundleId], dataId)
		if err != nil {
			return nil, err
		}

		// prevent runtime error
		if batchInfo.Count == index {
			break
		}
		index++
	}
	return batchData, nil
}

func (manager *Manager) GetBundle(ctx context.Context, bundleId types.ID) (*bundle, error) {
	rawBundle, err := manager.metadb.Find(ctx, bson.M{"bundleId": bundleId.Hex()}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve bundle data")
	}

	bundleData := &bundle{Id: bundleId.Hex()}
	if err = mapstructure.Decode(rawBundle, bundleData); err != nil {
		return nil, errors.Wrap(err, "failed to decode document")
	}

	bundleData.DataIds = make([]string, len(bundleData.RawDataIds))
	for index, id := range bundleData.RawDataIds {
		dataId, _, err := types.RawIdToDataId(id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get dataId")
		}
		bundleData.DataIds[index] = dataId.String()
	}
	bundleData.RawDataIds = nil

	return bundleData, nil
}
