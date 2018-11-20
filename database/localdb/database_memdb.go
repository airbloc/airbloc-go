package localdb

type MemDB struct {
	db map[string][]byte
	v  int
}

func NewMemDB() Database {
	return &MemDB{
		db: make(map[string][]byte),
	}
}

func (db *MemDB) Path() string {
	return ""
}

func (db *MemDB) Version() int {
	return db.v
}

func (db *MemDB) Put(key []byte, value []byte) error {
	db.db[string(key)] = value
	return nil
}

func (db *MemDB) Has(key []byte) (bool, error) {
	_, ok := db.db[string(key)]
	return ok, nil
}

func (db *MemDB) Get(key []byte) ([]byte, error) {
	return db.db[string(key)], nil
}

func (db *MemDB) Delete(key []byte) error {
	delete(db.db, string(key))
	return nil
}

func (db *MemDB) Close() error {
	return nil
}

func (db *MemDB) NewBatch() Batch {
	return nil
}

func (db *MemDB) getDB() interface{} {
	return db.db
}
