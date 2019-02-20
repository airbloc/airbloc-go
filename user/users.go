package user

import (
	"context"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/mitchellh/mapstructure"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
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

func (manager *Manager) GetData(ctx context.Context, id common.ID) ([]*userData, error) {
	manager.GetDataIds(ctx, id)
	return nil, nil
}

type userDataInfo struct {
	CollectionId string `json:"_id" mapstructure:"_id"`
	DataIds      []struct {
		Id          string `json:"id"`
		IngestedAt  int64  `json:"ingestedAt"`
		CollectedAt int64  `json:"collectedAt"`
	} `json:"dataIds" mapstructure:"-"`
	RawDataIds [][]primitive.D `json:"-" mapstructure:"dataIds"`
}

func (manager *Manager) GetDataIds(ctx context.Context, id common.ID) ([]*userDataInfo, error) {
	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"data.data.dataIds.userId", id.Hex()}}}},
		bson.D{{"$project", bson.D{
			{"data.data.ingestedAt", 1},
			{"data.data.collectedAt", 1},
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

	var infoes []*userDataInfo
	for cur.Next(ctx) {
		elem := &bson.D{}
		if err := cur.Decode(elem); err != nil {
			return nil, errors.Wrap(err, "retrieving document")
		}

		// debug
		//d, _ := json.MarshalIndent(elem.Map(), "", "    ")
		//log.Println(string(d))

		collection := new(userDataInfo)
		if err := mapstructure.Decode(elem.Map(), &collection); err != nil {
			return nil, errors.Wrap(err, "decoding document")
		}

		// dataIds
		index := 0
		for _, idPack := range collection.RawDataIds {
			collection.DataIds = append(collection.DataIds, make([]struct {
				Id          string `json:"id"`
				IngestedAt  int64  `json:"ingestedAt"`
				CollectedAt int64  `json:"collectedAt"`
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
