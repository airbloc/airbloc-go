package metadb

import (
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

func (db *bigchainDB) signTx(tx *txn.Transaction) error {
	return tx.Sign([]*txn.KeyPair{db.key})
}

func (db *bigchainDB) sendTx(tx *txn.Transaction, mode Mode) error {
	switch mode {
	case BigchainTxModeDefault:
		return db.bdb.PostTransaction(tx)
	case BigchainTxModeCommit:
		return db.bdb.PostTransactionCommit(tx)
	case BigchainTxModeSync:
		return db.bdb.PostTransactionSync(tx)
	default:
		return errors.New("invalid mode (0 = default, 1 = commit, 2 = sync)")
	}
}

func (db *bigchainDB) newCreateTransaction(asset txn.Asset, metadata txn.Metadata) (*txn.Transaction, error) {
	// generate condition
	cond := txn.NewEd25519Condition(db.key.PublicKey)
	out, err := txn.NewOutput(*cond, BigchainDBAmount)
	if err != nil {
		return nil, err
	}

	// create transaction
	return txn.NewCreateTransaction(
		asset,
		metadata,
		[]txn.Output{out},
		[]ed25519.PublicKey{db.key.PublicKey},
	)
}

func (db *bigchainDB) newTransferTransaction(txid string, to ed25519.PublicKey, metadata txn.Metadata) (*txn.Transaction, error) {
	tx, err := db.bdb.GetTransaction(txid)
	if err != nil {
		return nil, err
	}

	cond := txn.NewEd25519Condition(to)
	out, err := txn.NewOutput(*cond, BigchainDBAmount)
	if err != nil {
		return nil, err
	}

	return txn.NewTransferTransaction(
		[]txn.Transaction{tx},
		[]txn.Output{out},
		metadata,
	)
}
