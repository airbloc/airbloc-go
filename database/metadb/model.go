package metadb

import (
	"context"

	"github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"golang.org/x/crypto/ed25519"
)

type Model struct {
	database Database
	Name     string
}

func (model *Model) Create(immutableData map[string]interface{}, mutableData map[string]interface{}) (*transaction.Transaction, error) {
	assetData := make(map[string]interface{})
	assetData["type"] = model.Name
	assetData["data"] = immutableData
	asset := transaction.Asset{Data: assetData}

	return model.database.Create(asset, mutableData, BigchainTxModeDefault)
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

func NewModel(database Database, name string) *Model {
	return &Model{database, name}
}
