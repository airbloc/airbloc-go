package metadb

import (
	"context"
	"github.com/bigchaindb/go-bigchaindb-driver/pkg/client"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"golang.org/x/crypto/ed25519"
	"net/http"
)

type bigchainDB struct {
	bdb    *client.Client
	client *http.Client
	mdb    *mongo.Database
	key    *txn.KeyPair
	v      int
}

// http://localhost:9984 or external gateway
func NewBigchainDB(bdbUrl, mdbUrl string, key *txn.KeyPair, version int) (Database, error) {
	log.Debug("BigchainDB initiated", "endpoint", bdbUrl)
	config := client.ClientConfig{Url: bdbUrl}
	bdbClient, err := client.New(config)
	if err != nil {
		return nil, err
	}

	block, err := bdbClient.GetBlock("0")
	if err != nil {
		panic(err)
	}
	log.Debug("Connection Test", "block", block)

	mdbClient, err := mongo.NewClient(mdbUrl)
	if err != nil {
		return nil, err
	}

	return &bigchainDB{
		bdb:    bdbClient,
		mdb:    mdbClient.Database(BigchainDBName),
		client: &http.Client{},
		key:    key,
		v:      version,
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

	inTxn, err := db.prepareTx(tx)
	if err != nil {
		return
	}

	err = db.sendIntermediateTx(inTxn)
	return
}

func (db *bigchainDB) RetrieveOne(
	ctx context.Context,
	query *bson.Document,
	opts ...findopt.One,
) (*bson.Document, error) {
	metaDB := db.mdb.Collection(BigchainAssetCollection)
	res := metaDB.FindOne(ctx, query, opts...)

	doc := bson.NewDocument()
	err := res.Decode(&doc)
	if err != nil {
		return nil, err
	}

	// TODO: filter burned assets
	return doc, nil
}

func (db *bigchainDB) RetrieveMany(
	ctx context.Context,
	query *bson.Document,
	opts ...findopt.Find,
) (*bson.Document, error) {
	metaDB := db.mdb.Collection(BigchainAssetCollection)

	cursor, err := metaDB.Find(ctx, query, opts...)
	if err != nil {
		return nil, err
	}

	doc := bson.NewDocument()
	cursor.Decode(&doc)

	// TODO: filter burned assets
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

	inTxn, err := db.prepareTx(tx)
	if err != nil {
		return
	}

	err = db.sendIntermediateTx(inTxn)
	return
}

func (db *bigchainDB) Burn(txid string, mode Mode) error {
	return db.Append(txid, BigchainBurnAddress, txn.Metadata{"status": "BURNED"}, mode)
}

func (db *bigchainDB) Close() (err error) {
	err = db.mdb.Client().Disconnect(context.Background())
	return
}
