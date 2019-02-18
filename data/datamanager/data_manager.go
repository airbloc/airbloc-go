package datamanager

import (
	"context"
	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/mitchellh/mapstructure"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

type Manager struct {
	kms         key.Manager
	client      blockchain.TxClient
	metadb      metadb.Database
	p2p         p2p.Server
	warehouse   *warehouse.DataWarehouse
	registry    *adapter.DataRegistry
	collections *collections.Collections
	batches     *data.BatchManager
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
		kms:         kms,
		client:      client,
		p2p:         p2p,
		warehouse:   warehouse,
		registry:    contract.(*adapter.DataRegistry),
		collections: collections.New(client),
		batches:     batches,
	}
}

func (manager *Manager) Batches() *data.BatchManager {
	return manager.batches
}

func (manager *Manager) encrypt(data *common.Data) (*common.EncryptedData, error) {
	encryptedPayload, err := manager.kms.Encrypt(data.Payload)
	if err != nil {
		return nil, err
	}
	return &common.EncryptedData{
		UserId:  data.UserId,
		Payload: encryptedPayload,
	}, nil
}

type getDataResult struct {
	CollectionId common.ID
	UserId       common.ID
	IngestedAt   common.Time
	Payload      string
}

func (manager *Manager) Get(rawDataId string) (*getDataResult, error) {
	dataId, err := common.NewDataId(rawDataId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse data ID %s", rawDataId)
	}

	bundle, err := manager.warehouse.Get(dataId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve bundle of data %s", rawDataId)
	}

	// prevent runtime error
	if uint32(len(bundle.Data[dataId.UserId])) < dataId.RowId.Uint32() {
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
	bundles := make(map[common.ID]*data.Bundle, batch.Count)

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

func (manager *Manager) GetBundleInfo(ctx context.Context, id common.ID) (*bundleInfo, error) {
	rawBundle, err := manager.metadb.RetrieveOne(ctx, bson.M{"bundleId": id.Hex()})
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
		rawDataId := new(common.RawDataId)
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

type userInfo struct {
	AppId        string `json:"appId" mapstructure:"-"`
	SchemaId     string `json:"schemaId" mapstructure:"-"`
	CollectionId string `json:"_id" mapstructure:"_id"`
	DataIds      []struct {
		Id         string `json:"id"`
		IngestedAt int64  `json:"ingestedAt"`
	} `json:"dataIds" mapstructure:"-"`
	RawDataIds [][]primitive.D `json:"-" mapstructure:"dataIds"`
}

func (manager *Manager) GetUserInfo(ctx context.Context, id common.ID) ([]*userInfo, error) {
	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"data.data.dataIds.userId", id.Hex()}}}},
		bson.D{{"$project", bson.D{
			{"data.data.ingestedAt", 1},
			{"data.data.collection", 1},
			{"data.data.dataIds", bson.D{{
				"$filter", bson.D{
					{"input", "$data.data.dataIds"},
					{"as", "dataId"},
					{"cond", bson.D{{
						"$eq", bson.A{"$$dataId.userId", id.Hex()},
					}}},
				},
			}}},
		}}},
		bson.D{{"$addFields", bson.D{{
			"data.data.dataIds", bson.D{{
				"ingestedAt", "$data.data.ingestedAt",
			}},
		}}}},
		bson.D{{"$group", bson.D{
			{"_id", "$data.data.collection"},
			{"dataIds", bson.D{{
				"$addToSet", "$data.data.dataIds",
			}}},
		}}},
	}

	cur, err := manager.metadb.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "aggregating data pipeline")
	}
	defer cur.Close(ctx)

	var infoes []*userInfo
	for cur.Next(ctx) {
		elem := &bson.D{}
		if err := cur.Decode(elem); err != nil {
			return nil, errors.Wrap(err, "retrieving document")
		}

		// debug
		//d, _ := json.MarshalIndent(elem.Map(), "", "    ")
		//log.Println(string(d))

		collection := new(userInfo)
		if err := mapstructure.Decode(elem.Map(), &collection); err != nil {
			return nil, errors.Wrap(err, "decoding document")
		}

		// appId, schemaId, etc...
		collectionId, err := common.HexToID(collection.CollectionId)
		if err != nil {
			return nil, errors.Wrap(err, "converting collectionId")
		}

		collectionInfo, err := manager.collections.Get(collectionId)
		collection.AppId = collectionInfo.AppId.Hex()
		collection.SchemaId = collectionInfo.Schema.Id.Hex()

		// dataIds
		index := 0
		for _, idPack := range collection.RawDataIds {
			collection.DataIds = append(collection.DataIds, make([]struct {
				Id         string `json:"id"`
				IngestedAt int64  `json:"ingestedAt"`
			}, len(idPack))...)
			for _, id := range idPack {
				rawDataId := new(common.RawDataId)
				if err := mapstructure.Decode(id.Map(), rawDataId); err != nil {
					return nil, errors.Wrap(err, "decoding rawDataId")
				}

				dataId, err := rawDataId.Convert()
				if err != nil {
					return nil, errors.Wrap(err, "converting dataId")
				}
				collection.DataIds[index].Id = dataId.Hex()
				collection.DataIds[index].IngestedAt = int64(rawDataId.IngestedAt)
				index++
			}
		}
		collection.RawDataIds = nil
		infoes = append(infoes, collection)
	}
	return infoes, nil
}
