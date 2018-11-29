package metadb

import (
	"context"
	"fmt"
	"strings"

	"github.com/azer/logger"
	"github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
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
	assetData := make(map[string]interface{})
	assetData["type"] = model.Name
	assetData["data"] = immutableData
	asset := transaction.Asset{Data: assetData}

	tx, err := model.database.Create(asset, mutableData, BigchainTxModeDefault)
	if err != nil {
		return tx, err
	}
	truncatedId := fmt.Sprintf("%s…%s", (*tx.ID)[:6], (*tx.ID)[58:])
	model.log.Info("Metadata created with", logger.Attrs{"id": truncatedId})
	return tx, nil
}

func (model *Model) RetrieveAsset(query *bson.Document) (*bson.Document, error) {
	panic("implement me")
}

func (model *Model) RetrieveMany(context.Context, *bson.Document, ...findopt.Find) (*bson.Document, error) {
	panic("implement me")
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
