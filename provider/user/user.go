package user

import (
	"context"
	"fmt"
	"github.com/airbloc/airbloc-go/provider/collections"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/database/metadb"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/mitchellh/mapstructure"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

type Manager struct {
	kms         key.Manager
	metadb      metadb.Database
	warehouse   *warehouse.Manager
	collections *collections.Manager
}

func NewManager(
	kms key.Manager,
	metaDB metadb.Database,
	client blockchain.TxClient,
	warehouse *warehouse.Manager,
) *Manager {
	return &Manager{
		kms:         kms,
		warehouse:   warehouse,
		metadb:      metaDB,
		collections: collections.NewManager(client),
	}
}

type userData struct {
	CollectionId string            `json:"collectionId"`
	Data         []userDataPayload `json:"data" mapstructure:"-"`
}

type userDataPayload struct {
	types.DataId
	CollectedAt int64  `json:"collectedAt"`
	IngestedAt  int64  `json:"ingestedAt"`
	Payload     string `json:"payload"`
}

func (manager *Manager) GetData(ctx context.Context, id types.ID, from int64) ([]userData, error) {
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

	cond = bson.D{{
		"$project", bson.D{{
			"collections", bson.D{{
				"$filter", bson.D{
					{"input", "$collections"},
					{"as", "collections"},
					{"cond", cond},
				},
			}},
		}},
	}}

	infoes, err := manager.getDataIds(ctx, id, cond)
	if err != nil {
		return nil, errors.Wrap(err, "get data ids")
	}

	usersData := make([]userData, len(infoes))
	for i, info := range infoes {
		usersData[i] = userData{
			CollectionId: info.CollectionId,
			Data:         make([]userDataPayload, len(info.DataIds)),
		}

		for j, dataId := range info.DataIds {
			bundle, err := manager.warehouse.Get(&dataId.DataId)
			if err != nil {
				return nil, errors.Wrap(err, "fetching bundle")
			}

			if _, ok := bundle.Data[dataId.UserId]; !ok {
				return nil, errors.Errorf("cannot find user %s on given bundle", dataId.UserId)
			}

			if uint32(len(bundle.Data[dataId.UserId])) <= dataId.RowId.Uint32() {
				return nil, fmt.Errorf("cannot find row data at given userId")
			}

			encryptedData := bundle.Data[dataId.UserId][dataId.RowId.Uint32()]
			if encryptedData.RowId.Uint32() != dataId.RowId.Uint32() {
				return nil, fmt.Errorf(
					"rowId mismatching : expected %s, actual %s",
					dataId.RowId.Hex(), encryptedData.RowId.Hex(),
				)
			}

			payload, err := manager.kms.DecryptData(encryptedData)
			if err != nil {
				return nil, errors.Wrap(err, "decrypting data")
			}

			usersData[i].Data[j] = userDataPayload{
				DataId:      dataId.DataId,
				CollectedAt: dataId.CollectedAt,
				IngestedAt:  info.IngestedAt,
				Payload:     payload.Payload,
			}
		}
	}

	return usersData, nil
}

type userDataInfo struct {
	CollectionId string `json:"collection" mapstructure:"collection"`
	IngestedAt   int64  `json:"ingestedAt"`
	DataIds      []struct {
		types.DataId
		CollectedAt int64 `json:"collectedAt"`
	} `json:"dataIds" mapstructure:"-"`
	RawDataIds []bson.D `json:"rawDataIds" mapstructure:"dataIds"`
}

type userDataInfoQueryResponse struct {
	CollectionId string   `json:"collectionId" mapstructure:"_id"`
	Collections  []bson.D `json:"collections" mapstructure:"collections"`
}

func (manager *Manager) getDataIds(ctx context.Context, id types.ID, cond ...bson.D) ([]userDataInfo, error) {
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
		{{"$group", bson.D{
			{"_id", "$collection"},
			{"collections", bson.D{{
				"$push", "$$ROOT",
			}}},
		}}},
	}
	pipeline = append(pipeline, cond...)

	cur, err := manager.metadb.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "aggregating data pipeline")
	}
	defer cur.Close(ctx)

	var infoes []userDataInfo
	for cur.Next(ctx) {
		var doc bson.D
		if err := cur.Decode(&doc); err != nil {
			return nil, errors.Wrap(err, "retrieving document")
		}

		var resp userDataInfoQueryResponse
		if err := mapstructure.Decode(doc.Map(), &resp); err != nil {
			return nil, errors.Wrap(err, "decoding document")
		}

		for _, rawCollection := range resp.Collections {
			var collection userDataInfo
			if err := mapstructure.Decode(rawCollection.Map(), &collection); err != nil {
				return nil, errors.Wrap(err, "decoding document")
			}

			// dataIds
			collection.DataIds = make([]struct {
				types.DataId
				CollectedAt int64 `json:"collectedAt"`
			}, len(collection.RawDataIds))

			for i, idPack := range collection.RawDataIds {
				var rawDataId types.RawDataId
				if err := mapstructure.Decode(idPack.Map(), &rawDataId); err != nil {
					return nil, errors.Wrap(err, "decoding dataId")
				}

				dataId, err := rawDataId.Convert()
				if err != nil {
					return nil, errors.Wrap(err, "failed to unmarshal data ID")
				}
				collection.DataIds[i].DataId = *dataId
				collection.DataIds[i].CollectedAt = int64(rawDataId.CollectedAt)
			}
			collection.RawDataIds = nil
			infoes = append(infoes, collection)
		}
	}

	return infoes, nil
}

// returns all of user's dataIds
func (manager *Manager) GetDataIds(ctx context.Context, id types.ID) ([]userDataInfo, error) {
	return manager.getDataIds(ctx, id)
}
