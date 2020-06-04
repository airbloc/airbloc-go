package localdb

import (
	"os"

	"github.com/dgraph-io/badger"
	"github.com/pkg/errors"
)

type badgerDB struct {
	fn string // filename for reporting
	db *badger.DB
	v  int
}

func NewBadgerDatabase(path string, version int) (Database, error) {
	if fi, err := os.Stat(path); err == nil {
		if !fi.IsDir() {
			return nil, errors.Errorf("database: open %s: not a directory", path)
		}
	} else if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return nil, err
		}
	}

	opts := badger.DefaultOptions(path)
	opts.WithNumMemtables(5)
	opts.WithSyncWrites(false)
	opts.WithNumCompactors(3)
	opts.WithReadOnly(false)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &badgerDB{
		fn: path,
		db: db,
		v:  version,
	}, nil
}

func (db *badgerDB) Path() string {
	return db.fn
}

func (db *badgerDB) Version() int {
	return db.v
}

func (db *badgerDB) Put(key []byte, value []byte) error {
	txn := db.db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set(key, value)
	if err != nil {
		return err
	}
	return txn.Commit()
}

func (db *badgerDB) Has(key []byte) (bool, error) {
	txn := db.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(key)
	// badger.ErrKeyNotFound
	if err == badger.ErrKeyNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	value, err := item.ValueCopy(nil)
	return value != nil, err
}

func (db *badgerDB) Get(key []byte) ([]byte, error) {
	txn := db.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(key)
	if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}

func (db *badgerDB) Delete(key []byte) error {
	txn := db.db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Delete(key)
	if err != nil {
		return err
	}
	return txn.Commit()
}

func (db *badgerDB) NewIterator() *badger.Iterator {
	txn := db.db.NewTransaction(false)
	return txn.NewIterator(badger.DefaultIteratorOptions)
}

func (db *badgerDB) Close() error {
	err := db.db.Close()
	if err != nil {
		return errors.Wrap(err, "database: close database")
	}
	return nil
}

func (db *badgerDB) NewBatch() Batch {

	txn := db.db.NewTransaction(true)

	return &badgerBatch{db: db.db, txn: txn}
}

func (db *badgerDB) getDB() interface{} {
	return db.db
}

type badgerBatch struct {
	db   *badger.DB
	txn  *badger.Txn
	size int
}

func (b *badgerBatch) Put(key, value []byte) error {

	err := b.txn.Set(key, value)
	b.size += len(value)

	return err
}

func (b *badgerBatch) Write() error {
	return b.txn.Commit()
}

func (b *badgerBatch) ValueSize() int {
	return b.size
}

func (b *badgerBatch) Reset() {
	b.txn.Discard()
	b.size = 0
}
