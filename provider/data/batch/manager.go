package batch

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/airbloc/airbloc-go/shared/database/localdb"
	"github.com/pkg/errors"
)

// BatchManager creates a batch (pointing multiple Data ID)
// and manages batches by storing them on the local database.
type BatchManager struct {
	localStorage *localdb.Model
	batches      map[string]*Batch
}

// NewBatchManager creates a instance of BatchManager.
func NewBatchManager(localDatabase localdb.Database) *BatchManager {
	return &BatchManager{
		localStorage: localdb.NewModel(localDatabase, "batch"),
		batches:      map[string]*Batch{},
	}
}

func createBatchId() (string, error) {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (operator *BatchManager) Create() (*Batch, string, error) {
	batchId, err := createBatchId()
	if err != nil {
		return nil, batchId, err
	}

	batch := newBatch(batchId)
	operator.batches[batchId] = batch
	return batch, batchId, nil
}

func (operator *BatchManager) Get(batchId string) (*Batch, error) {
	if operator.batches[batchId] != nil {
		return operator.batches[batchId], nil
	}

	// try in local storage
	existsInLocal, err := operator.localStorage.Has(batchId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query batchId to local database")
	}
	if !existsInLocal {
		return nil, errors.Errorf("batchId %s does not exist.", batchId)
	}

	batchData, err := operator.localStorage.Get(batchId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load batchId to local database")
	}

	return UnmarshalBatch(batchId, batchData)
}

func (operator *BatchManager) Save(batchId string) error {
	batch := operator.batches[batchId]
	if batch != nil {
		return errors.Errorf("batchId %s not found.", batchId)
	}

	batchData := batch.Marshal()
	return operator.localStorage.Put(batchId, batchData)
}

func (operator *BatchManager) Delete(batchId string) error {
	if operator.batches[batchId] != nil {
		return errors.Errorf("batchId %s not found.", batchId)
	}
	delete(operator.batches, batchId)
	return operator.localStorage.Delete(batchId)
}
