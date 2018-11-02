package localdb

type Model struct {
	db     Database
	prefix []byte
}

func NewModel(db Database, prefix string) *Model {
	return &Model{
		db:     db,
		prefix: []byte(prefix),
	}
}

func (m *Model) Put(key string, value []byte) error {
	return m.db.Put(append(m.prefix, []byte(key)...), value)
}

func (m *Model) Get(key string) ([]byte, error) {
	return m.db.Get(append(m.prefix, []byte(key)...))
}

func (m *Model) Has(key string) (bool, error) {
	return m.db.Has(append(m.prefix, []byte(key)...))
}

func (m *Model) Delete(key string) error {
	return m.db.Delete(append(m.prefix, []byte(key)...))
}

func (m *Model) NewBatch() *ModelBatch {
	return &ModelBatch{
		batch:  m.db.NewBatch(),
		prefix: m.prefix,
	}
}

type ModelBatch struct {
	batch  Batch
	prefix []byte
}

func (b *ModelBatch) Put(key string, value []byte) error {
	return b.batch.Put(append(b.prefix, []byte(key)...), value)
}

func (b *ModelBatch) Write() error {
	return b.batch.Write()
}

func (b *ModelBatch) ValueSize() int {
	return b.batch.ValueSize()
}

func (b *ModelBatch) Reset() {
	b.batch.Reset()
}
