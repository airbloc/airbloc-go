package metadb

import (
	"context"
	"github.com/azer/logger"
	"github.com/pkg/errors"
	"net/http"

	"github.com/bigchaindb/go-bigchaindb-driver/pkg/client"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"golang.org/x/crypto/ed25519"
)

type bigchainDB struct {
	bdb      *client.Client
	client   *http.Client
	proxyUrl string
	mdb      *mongo.Database
	key      *txn.KeyPair
	v        int
	log      *logger.Logger
}

// http://localhost:9984 or external gateway
func NewBigchainDB(bdbUrl, mdbUrl, proxyUrl string, key *txn.KeyPair, version int) (Database, error) {
	config := client.ClientConfig{Url: bdbUrl}
	bdbClient, err := client.New(config)
	if err != nil {
		return nil, err
	}
	mdbClient, err := mongo.NewClient(mdbUrl)
	if err != nil {
		return nil, err
	}
	db := &bigchainDB{
		bdb:      bdbClient,
		mdb:      mdbClient.Database(BigchainDBName),
		proxyUrl: proxyUrl,
		client:   &http.Client{},
		key:      key,
		v:        version,
		log:      logger.New("bigchaindb"),
	}
	if err := db.dial(); err != nil {
		return db, errors.Wrap(err, "unable to connect to BigchainDB")
	}
	return db, nil
}

func (db *bigchainDB) dial() error {
	block, err := db.bdb.GetBlock("0")
	if err != nil {
		return err
	}
	db.log.Info("Connected to BigchainDB. Received", logger.Attrs{
		"block": block,
	})
	return nil
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

	results, err := db.sendIntermediateTx(inTxn)
	if !results.Exists("id") {
		return nil, errors.New("server returned no transaction ID")
	}
	txId := string(results.GetStringBytes("id"))
	tx.ID = &txId
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

	_, err = db.sendIntermediateTx(inTxn)
	return
}

func (db *bigchainDB) Burn(txid string, mode Mode) error {
	return db.Append(txid, BigchainBurnAddress, txn.Metadata{"status": "BURNED"}, mode)
}

func (db *bigchainDB) Close() (err error) {
	err = db.mdb.Client().Disconnect(context.Background())
	return
}
