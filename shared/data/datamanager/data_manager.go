package datamanager

import (
	"context"
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/collections"
	"github.com/airbloc/airbloc-go/shared/data"
	"github.com/airbloc/airbloc-go/shared/database/localdb"
	"github.com/airbloc/airbloc-go/shared/database/metadb"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/p2p"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/airbloc-go/shared/warehouse"
	"github.com/mitchellh/mapstructure"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/pkg/errors"
)

type Manager struct {
	kms         key.Manager
	client      blockchain.TxClient
	metadb      *metadb.Model
	p2p         p2p.Server
	warehouse   *warehouse.DataWarehouse
	registry    *adapter.DataRegistry
	collections *collections.Collections
	batches     *data.BatchManager
}

func NewManager(
	kms key.Manager,
	p2p p2p.Server,
	metaDB metadb.Database,
	localDB localdb.Database,
	client blockchain.TxClient,
	warehouse *warehouse.DataWarehouse,
) *Manager {
	batches := data.NewBatchManager(localDB)
	contract := client.GetContract(&adapter.DataRegistry{})
	return &Manager{
		kms:         kms,
		client:      client,
		p2p:         p2p,
		warehouse:   warehouse,
		registry:    contract.(*adapter.DataRegistry),
		collections: collections.New(client),
		batches:     batches,
		metadb:      metadb.NewModel(metaDB, "bundles"),
	}
}

func (manager *Manager) Batches() *data.BatchManager {
	return manager.batches
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

type getDataResult struct {
	CollectionId types.ID
	UserId       types.ID
	IngestedAt   types.Time
	Payload      string
}

func (manager *Manager) Get(rawDataId string) (*getDataResult, error) {
	dataId, err := types.NewDataId(rawDataId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse data ID %s", rawDataId)
	}

	bundle, err := manager.warehouse.Get(dataId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", rawDataId)
	}

	// prevent runtime error
	if uint32(len(bundle.Data[dataId.UserId])) <= dataId.RowId.Uint32() {
		return nil, errors.New("data does not exists")
	}

	encryptedData := bundle.Data[dataId.UserId][dataId.RowId.Uint32()]
	d, err := manager.kms.DecryptData(encryptedData)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt data %s", dataId.String())
	}

	return &getDataResult{
		CollectionId: bundle.Collection,
		UserId:       d.UserId,
		IngestedAt:   bundle.IngestedAt,
		Payload:      d.Payload,
	}, nil
}

func (manager *Manager) GetBatch(batch *data.Batch) ([]*getDataResult, error) {
	result := make([]*getDataResult, batch.Count)
	bundles := make(map[types.ID]*data.Bundle, batch.Count)

	index := 0
	for dataId := range batch.Iterator() {
		if _, alreadyFetched := bundles[dataId.BundleId]; !alreadyFetched {
			b, err := manager.warehouse.Get(&dataId)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", dataId.String())
			}
			bundles[dataId.BundleId] = b
		}

		bundle := bundles[dataId.BundleId]

		// prevent runtime error
		if uint32(len(bundle.Data[dataId.UserId])) < dataId.RowId.Uint32() {
			return nil, errors.New("data does not exists")
		}

		encryptedData := bundle.Data[dataId.UserId][dataId.RowId.Uint32()]
		d, err := manager.kms.DecryptData(encryptedData)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decrypt data %s", dataId.String())
		}

		result[index] = &getDataResult{
			CollectionId: bundle.Collection,
			UserId:       d.UserId,
			IngestedAt:   bundle.IngestedAt,
			Payload:      d.Payload,
		}

		// prevent runtime error
		if batch.Count == index {
			break
		}
		index++
	}
	return result, nil
}

type bundleInfo struct {
	Id         string        `json:"bundleId" mapstructure:"bundleId"`
	Uri        string        `json:"uri" mapstructure:"uri"`
	Provider   string        `json:"provider" mapstructure:"provider"`
	Collection string        `json:"collection" mapstructure:"collection"`
	IngestedAt int64         `json:"ingestedAt" mapstructure:"ingestedAt"`
	DataIds    []string      `json:"-" mapstructure:"-"`
	RawDataIds []primitive.D `json:"dataIds" mapstructure:"dataIds"`
}

func (manager *Manager) GetBundleInfo(ctx context.Context, id types.ID) (*bundleInfo, error) {
	rawBundle, err := manager.metadb.RetrieveAsset(bson.M{"bundleId": id.Hex()})
	if err != nil {
		return nil, errors.Wrap(err, "retrieving bundle data")
	}

	// debug
	//d, _ := json.MarshalIndent(rawBundle, "", "    ")
	//log.Println(string(d))

	bundleInfo := new(bundleInfo)
	bundleInfo.Id = id.Hex()
	if err := mapstructure.Decode(rawBundle, bundleInfo); err != nil {
		return nil, errors.Wrap(err, "decoding document")
	}

	bundleInfo.DataIds = make([]string, len(bundleInfo.RawDataIds))
	for index, id := range bundleInfo.RawDataIds {
		rawDataId := new(types.RawDataId)
		if err := mapstructure.Decode(id.Map(), rawDataId); err != nil {
			return nil, errors.Wrap(err, "decoding rawDataId")
		}

		dataId, err := rawDataId.Convert()
		if err != nil {
			return nil, errors.Wrap(err, "converting dataId")
		}
		bundleInfo.DataIds[index] = dataId.Hex()
	}
	bundleInfo.RawDataIds = nil

	return bundleInfo, nil
}
