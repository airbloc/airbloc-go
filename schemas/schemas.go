package schemas

import (
	"context"
	"encoding/json"

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

func (s *Schemas) Register(ctx context.Context, name string, schema map[string]interface{}) (common.ID, error) {
	rawSchema, err := json.Marshal(schema)
	if err != nil {
		return common.ID{}, errors.Wrap(err, "given schema is not a valid JSON schema")
	}

	if nameExists, err := s.NameExists(name); err != nil {
		return common.ID{}, err
	} else if nameExists {
		return common.ID{}, ErrNameExists
	}

	// register schema to the blockchain and get ID
	dtx, err := s.contract.Register(s.client.Account(), name)
	if err != nil {
		return common.ID{}, err
	}

	receipt, err := s.client.WaitMined(ctx, dtx)
	if err != nil {
		return common.ID{}, errors.Wrap(err, "failed to wait for tx to be mined")
	}

	event, err := s.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return common.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	schemaId := common.ID(event.Id)
	s.log.Info("Registered new schema %s with", name, logger.Attrs{"id": schemaId.Hex()})

	// create metadata
	metadata := map[string]interface{}{
		"name":   name,
		"id":     schemaId.Hex(),
		"schema": string(rawSchema),
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

func (s *Schemas) Unregister(ctx context.Context, id common.ID) error {
	tx, err := s.contract.Unregister(s.client.Account(), id)
	if err != nil {
		return err
	}

	if _, err := s.client.WaitMined(ctx, tx); err != nil {
		return err
	}

	query := bson.NewDocument(bson.EC.String("data.id", id.Hex()))
	metadata, err := s.db.RetrieveAsset(query)
	if err != nil {
		return errors.Wrap(err, "failed to find the asset on metadb")
	}
	return s.db.Burn(metadata.Lookup("id").StringValue())
}
