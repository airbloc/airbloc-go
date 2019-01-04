package metadb

import (
	"context"
	"fmt"
	"strings"

	"github.com/azer/logger"
	"github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/mongodb/mongo-go-driver/bson"
	"golang.org/x/crypto/ed25519"
)

type Model struct {
	database Database
	Name     string
	log      *logger.Logger
}

// NewModel creates a collection for given metadatabase, with the name of the data type.
func NewModel(database Database, name string) *Model {
	loggerName := fmt.Sprintf("metadb (%s)", strings.ToLower(name))
	return &Model{
		Name:     name,
		database: database,
		log:      logger.New(loggerName),
	}
}

func (model *Model) Create(immutableData map[string]interface{}, mutableData map[string]interface{}) (*transaction.Transaction, error) {
	asset := transaction.Asset{Data: map[string]interface{}{
		"type": model.Name,
		"data": immutableData,
	}}
	tx, err := model.database.Create(asset, mutableData, BigchainTxModeDefault)
	if err != nil {
		return tx, err
	}
	truncatedId := fmt.Sprintf("%s…%s", (*tx.ID)[:6], (*tx.ID)[58:])
	model.log.Info("Metadata created with", logger.Attrs{"id": truncatedId})
	return tx, nil
}

func unwrap(rawAsset interface{}) bson.M {
	// Unwraps {
	//   "_id": MongoObjectId,
	//   "id": BigchainDBTxID,
	//   "data": {
	//     "type": ModelType,
	//     "data": ModelData
	//   }
	// }
	// into {"_id": BigchainDBTxID, ...ModelData}
	asset := rawAsset.(bson.M)
	assetData := asset["data"].(bson.D).Map()
	data := assetData["data"].(bson.D).Map()

	data["_id"] = asset["_id"]
	return data
}

func wrap(query bson.M) bson.M {
	// since data is included in "data.data" object, we need to wrap query
	wrappedQuery := bson.M{}
	for fieldName, value := range query {
		wrappedQuery["data.data."+fieldName] = value
	}
	return wrappedQuery
}

func (model *Model) RetrieveAsset(query bson.M) (bson.M, error) {
	asset, err := model.database.RetrieveOne(context.Background(), wrap(query))
	if err != nil {
		return nil, err
	}
	return unwrap(asset), err
}

func (model *Model) RetrieveMany(ctx context.Context, query bson.M) ([]bson.M, error) {
	assets, err := model.database.RetrieveMany(ctx, wrap(query))
	if err != nil {
		return nil, err
	}

	// unwrap results
	var results []bson.M
	for _, asset := range assets {
		results = append(results, unwrap(asset))
	}
	return results, nil
}

func (model *Model) Append(string, ed25519.PublicKey, transaction.Metadata, Mode) error {
	panic("implement me")
}

func (model *Model) Burn(assetId string) error {
	return model.database.Burn(assetId, BigchainTxModeDefault)
}

func (model *Model) Close() error {
	return model.database.Close()
}
