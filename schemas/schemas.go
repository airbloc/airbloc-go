package schemas

import (
	"context"
	"github.com/azer/logger"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/pkg/errors"
)

var (
	ErrNameExists = errors.New("Schema name already exists")
)

type Schemas struct {
	db       *metadb.Model
	client   blockchain.TxClient
	contract *adapter.SchemaRegistry
	log      *logger.Logger
}

func New(db metadb.Database, client blockchain.TxClient) *Schemas {
	contract := client.GetContract(&adapter.SchemaRegistry{})
	return &Schemas{
		db:       metadb.NewModel(db, "schema"),
		client:   client,
		contract: contract.(*adapter.SchemaRegistry),
		log:      logger.New("schemas"),
	}
}

func (s *Schemas) Register(schema *Schema) (common.ID, error) {
	if nameExists, err := s.NameExists(schema.Name); err != nil {
		return common.ID{}, err
	} else if nameExists {
		return common.ID{}, ErrNameExists
	}

	// register schema to the blockchain and get ID
	dtx, err := s.contract.Register(s.client.Account(), schema.Name)
	if err != nil {
		return common.ID{}, err
	}

	receipt, err := s.client.WaitMined(context.Background(), dtx)
	if err != nil {
		return common.ID{}, errors.Wrap(err, "failed to wait for tx to be mined")
	}

	event, err := s.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return common.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	schemaId := common.ID(event.Id)
	s.log.Info("Registered new schema %s with", schema.Name, logger.Attrs{"id": schemaId.Hex()})

	// create metadata
	metadata := map[string]interface{}{
		"name":   schema.Name,
		"id":     schemaId.Hex(),
		"schema": schema.Schema,
	}

	if _, err := s.db.Create(metadata, nil); err != nil {
		return schemaId, errors.Wrap(err, "failed to save metadata")
	}
	return schemaId, nil
}

func (s *Schemas) NameExists(name string) (bool, error) {
	hashedName := crypto.Keccak256Hash([]byte(name))
	return s.contract.NameExists(nil, hashedName)
}

// Get retrieves a schema from metadatabase.
func (s *Schemas) Get(id common.ID) (*Schema, error) {
	result, err := s.db.RetrieveAsset(bson.M{"id": id.Hex()})
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve schema from MetaDB")
	}
	schema, err := NewSchema(result["name"].(string), result["schema"].(string))
	if err != nil {
		return nil, err
	}
	schema.Id = id
	return schema, nil
}

func (s *Schemas) Unregister(id common.ID) error {
	tx, err := s.contract.Unregister(s.client.Account(), id)
	if err != nil {
		return err
	}
	if _, err := s.client.WaitMined(context.Background(), tx); err != nil {
		return err
	}
	result, err := s.db.RetrieveAsset(bson.M{"id": id.Hex()})
	if err != nil {
		return errors.Wrap(err, "failed to find the asset on metadb")
	}
	return s.db.Burn(result["_id"].(string))
}
