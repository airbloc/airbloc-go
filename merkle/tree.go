package merkle

import (
	"math/big"

	"github.com/dgraph-io/badger"
)

type tree []*level

func (t tree) flush(
	db *badger.DB, h int,
	bn *big.Int, dt []byte,
) error {
	return db.Update(func(txn *badger.Txn) error {
		for i := 0; i <= h; i++ { // to the top level
			for _, n := range *t[i] {
				if err := txn.Set(n.Key(i, bn), dt); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
