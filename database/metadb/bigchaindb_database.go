package metadb

import (
	"context"

	"github.com/bigchaindb/go-bigchaindb-driver/pkg/client"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"golang.org/x/crypto/ed25519"
)

type bigchainDB struct {
	bdb *client.Client
	mdb *mongo.Database
	key *txn.KeyPair
	v   int
}

// http://localhost:9984 or external gateway
func NewBigchainDB(bdbUrl, mdbUrl string, key *txn.KeyPair, version int) (Database, error) {
	config := client.ClientConfig{Url: bdbUrl}
	bdbClient, err := client.New(config)
	if err != nil {
		return nil, err
	}

	mdbClient, err := mongo.NewClient(mdbUrl)
	if err != nil {
		return nil, err
	}

	return &bigchainDB{
		bdb: bdbClient,
		mdb: mdbClient.Database(BigchainDBName),
		key: key,
		v:   version,
	}, nil
}

func (db *bigchainDB) Create(
	asset txn.Asset,
	metadata txn.Metadata,
	mode Mode,
) (tx *txn.Transaction, err error) {
	tx, err = db.newCreateTransaction(asset, metadata)
	if err != nil {
		return
	}

	if err = db.signTx(tx); err != nil {
		return
	}

	if err = db.sendTx(tx, mode); err != nil {
		return
	}

	return
}

func (db *bigchainDB) RetrieveOne(
	ctx context.Context,
	query *bson.Document,
	opts ...findopt.One,
) (*bson.Document, error) {
	metaDB := db.mdb.Collection(BigchainMetaCollection)
	res := metaDB.FindOne(ctx, query, opts...)

	doc := bson.NewDocument()
	err := res.Decode(&doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (db *bigchainDB) RetrieveMany(
	ctx context.Context,
	query *bson.Document,
	opts ...findopt.Find,
) (*bson.Document, error) {
	metaDB := db.mdb.Collection(BigchainMetaCollection)

	cursor, err := metaDB.Find(ctx, query, opts...)
	if err != nil {
		return nil, err
	}

	doc := bson.NewDocument()
	cursor.Decode(&doc)
	return doc, nil
}

func (db *bigchainDB) Append(
	txid string,
	to ed25519.PublicKey,
	metadata txn.Metadata,
	mode Mode,
) (err error) {
	var tx *txn.Transaction
	tx, err = db.newTransferTransaction(txid, to, metadata)
	if err != nil {
		return
	}

	if err = db.signTx(tx); err != nil {
		return
	}

	if err = db.sendTx(tx, mode); err != nil {
		return
	}

	return
}

func (db *bigchainDB) Burn(txid string, mode Mode) error {
	return db.Append(txid, BigchainBurnAddress, txn.Metadata{"status": "BURNED"}, mode)
}

func (db *bigchainDB) Close() (err error) {
	err = db.mdb.Client().Disconnect(context.Background())
	return
}
