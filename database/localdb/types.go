package localdb

import (
	"github.com/dgraph-io/badger"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ErrWrongType = errors.New("wrong database instance")
)

func GetBadgerDB(db Database) (*badger.DB, error) {
	v, ok := db.getDB().(*badger.DB)
	if !ok {
		return nil, ErrWrongType
	}
	return v, nil
}

func GetLevelDB(db Database) (*leveldb.DB, error) {
	v, ok := db.getDB().(*leveldb.DB)
	if !ok {
		return nil, ErrWrongType
	}
	return v, nil
}
