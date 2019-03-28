package schemas

import (
	"context"
	"github.com/airbloc/logger"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/database/metadb"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrNameExists = errors.New("Schema name already exists")
)

type Schemas struct {
	db       metadb.Database
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

func (s *Schemas) Register(ctx context.Context, schema *Schema) (types.ID, error) {
	if nameExists, err := s.NameExists(schema.Name); err != nil {
		return types.ID{}, err
	} else if nameExists {
		return types.ID{}, ErrNameExists
	}

	// register schema to the blockchain and get ID
	dtx, err := s.contract.Register(s.client.Account(), schema.Name)
	if err != nil {
		return types.ID{}, err
	}

	receipt, err := s.client.WaitMined(context.Background(), dtx)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to wait for tx to be mined")
	}

	event, err := s.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	schemaId := types.ID(event.Id)
	s.log.Info("Registered new schema {} with", schema.Name, logger.Attrs{"id": schemaId.Hex()})

	// create metadata
	metadata := bson.M{
		"name":   schema.Name,
		"id":     schemaId.Hex(),
		"schema": schema.Schema,
	}

	if err := s.db.Insert(ctx, []interface{}{metadata}, nil); err != nil {
		return schemaId, errors.Wrap(err, "failed to save metadata")
	}
	return schemaId, nil
}

func (s *Schemas) NameExists(name string) (bool, error) {
	hashedName := crypto.Keccak256Hash([]byte(name))
	return s.contract.NameExists(nil, hashedName)
}

// Get retrieves a schema from metadatabase.
func (s *Schemas) Get(ctx context.Context, id types.ID) (*Schema, error) {
	rawRes, err := s.db.Find(ctx, bson.M{"id": id.Hex()}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve schema from MetaDB")
	}

	if len(rawRes) == 0 {
		return nil, errors.New("empty result")
	}

	name, exists := rawRes[0]["name"]
	if !exists {
		return nil, errors.New("name field does not exists")
	}

	schema, exists := rawRes[0]["schema"]
	if !exists {
		return nil, errors.New("schema field does not exists")
	}

	scheme, err := NewSchema(name.(string), schema.(string))
	if err != nil {
		return nil, err
	}
	scheme.Id = id
	return scheme, nil
}

func (s *Schemas) Unregister(ctx context.Context, id types.ID) error {
	tx, err := s.contract.Unregister(s.client.Account(), id)
	if err != nil {
		return err
	}
	if _, err := s.client.WaitMined(context.Background(), tx); err != nil {
		return err
	}
	result, err := s.db.Find(ctx, bson.M{"id": id.Hex()}, nil)
	if err != nil {
		return errors.Wrap(err, "failed to find the asset on metadb")
	}
	return s.db.Delete(ctx, bson.M{"_id": result[0]["_id"]}, nil)
}
