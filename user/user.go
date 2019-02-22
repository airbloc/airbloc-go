package user

import (
	"context"
	"encoding/json"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/mitchellh/mapstructure"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
	"log"
)

type Manager struct {
	kms         key.Manager
	metadb      metadb.Database
	warehouse   *warehouse.DataWarehouse
	collections *collections.Collections
}

func NewManager(
	kms key.Manager,
	metaDB metadb.Database,
	client blockchain.TxClient,
	warehouse *warehouse.DataWarehouse,
) *Manager {
	return &Manager{
		kms:         kms,
		warehouse:   warehouse,
		metadb:      metaDB,
		collections: collections.New(client),
	}
}

type userData struct {
	CollectionId string `json:"_id" mapstructure:"-"`
	Data         []struct {
		CollectedAt int64  `json:"collectedAt"`
		IngestedAt  int64  `json:"ingestedAt"`
		Payload     string `json:"payload"`
	} `json:"data" mapstructure:"-"`
}

func (manager *Manager) GetData(ctx context.Context, id common.ID, from int64) ([]*userData, error) {
	var cond bson.D
	if from == 0 {
		cond = bson.D{{
			"$eq", bson.A{
				"$$collections.ingestedAt",
				bson.D{{"$max", "$collections.ingestedAt"}},
			},
		}}
	} else {
		cond = bson.D{{
			"$gte", bson.A{"$$collections.ingestedAt", from},
		}}
	}

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"data.data.dataIds.userId", id.Hex()}}}},
		{{"$replaceRoot", bson.D{{"newRoot", "$data.data"}}}},
		{{"$project", bson.D{
			{"ingestedAt", 1},
			{"collection", 1},
			{"dataIds", bson.D{{
				"$filter", bson.D{
					{"input", "$dataIds"},
					{"as", "dataId"},
					{"cond", bson.D{{
						"$eq", bson.A{"$$dataId.userId", id.Hex()},
					}}},
				},
			}}},
		}}},
		{{"$sort", bson.D{{"ingestedAt", -1}}}},
		{{"$group", bson.D{
			{"_id", "$collection"},
			{"collections", bson.D{{
				"$push", "$$ROOT",
			}}},
		}}},
		{{"$project", bson.D{{
			"collections", bson.D{{
				"$filter", bson.D{
					{"input", "$collections"},
					{"as", "collections"},
					{"cond", cond},
				},
			}},
		}}}},
	}

	cur, err := manager.metadb.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "aggregating data pipeline")
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		elem := &bson.D{}
		if err := cur.Decode(elem); err != nil {
			return nil, errors.Wrap(err, "retrieving document")
		}

		d, _ := json.MarshalIndent(elem, "", "    ")
		log.Println(string(d))
	}

	return nil, nil
}

type userDataInfo struct {
	CollectionId string `json:"collection" mapstructure:"collection"`
	IngestedAt   int64  `json:"ingestedAt"`
	DataIds      []struct {
		common.DataId
		CollectedAt int64 `json:"collectedAt"`
	} `json:"dataIds" mapstructure:"-"`
	RawDataIds []bson.D `json:"rawDataIds" mapstructure:"dataIds"`
}

func (manager *Manager) GetDataIds(ctx context.Context, id common.ID) ([]userDataInfo, error) {
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"data.data.dataIds.userId", id.Hex()}}}},
		{{"$replaceRoot", bson.D{{"newRoot", "$data.data"}}}},
		{{"$project", bson.D{
			{"ingestedAt", 1},
			{"collection", 1},
			{"dataIds", bson.D{{
				"$filter", bson.D{
					{"input", "$dataIds"},
					{"as", "dataId"},
					{"cond", bson.D{{
						"$eq", bson.A{"$$dataId.userId", id.Hex()},
					}}},
				},
			}}},
		}}},
	}

	cur, err := manager.metadb.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "aggregating data pipeline")
	}
	defer cur.Close(ctx)

	var infoes []userDataInfo
	for cur.Next(ctx) {
		elem := &bson.D{}
		if err := cur.Decode(elem); err != nil {
			return nil, errors.Wrap(err, "retrieving document")
		}

		var collection userDataInfo
		if err := mapstructure.Decode(elem.Map(), &collection); err != nil {
			return nil, errors.Wrap(err, "decoding document")
		}

		// dataIds
		collection.DataIds = make([]struct {
			common.DataId
			CollectedAt int64 `json:"collectedAt"`
		}, len(collection.RawDataIds))

		for i, idPack := range collection.RawDataIds {
			var rawDataId common.RawDataId
			if err := mapstructure.Decode(idPack.Map(), &rawDataId); err != nil {
				return nil, errors.Wrap(err, "decoding dataId")
			}

			dataId, err := rawDataId.Convert()
			if err != nil {
				return nil, errors.Wrap(err, "converting dataId")
			}
			collection.DataIds[i].DataId = *dataId
			collection.DataIds[i].CollectedAt = int64(rawDataId.CollectedAt)
		}
		collection.RawDataIds = nil
		infoes = append(infoes, collection)
	}

	return infoes, nil
}
