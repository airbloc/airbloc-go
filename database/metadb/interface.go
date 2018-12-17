package metadb

import (
	"context"

	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/mongodb/mongo-go-driver/bson"
	"golang.org/x/crypto/ed25519"
)

type Mode int

const (
	BigchainTxModeDefault = Mode(iota)
	BigchainTxModeCommit
	BigchainTxModeSync

	BigchainDBAmount = "1"

	// attach bigchainDB use mongoDB
	BigchainDBName          = "bigchain"
	BigchainTxCollection    = "transactions"
	BigchainMetaCollection  = "metadata"
	BigchainAssetCollection = "assets"
)

var (
	BigchainBurnAddress = ed25519.PublicKey("BurnBurnBurnBurnBurnBurnBurnBurnBurnBurnBurn")
)

type Database interface {
	Create(txn.Asset, txn.Metadata, Mode) (*txn.Transaction, error)
	RetrieveOne(context.Context, bson.M) (bson.M, error)
	RetrieveMany(context.Context, bson.M) ([]bson.M, error)
	Append(string, ed25519.PublicKey, txn.Metadata, Mode) error
	Burn(string, Mode) error
	Close() error
}
