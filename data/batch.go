package data

import (
	"strings"

	"github.com/airbloc/airbloc-go/common"
)

type OwnerIDs map[common.ID][]common.ID
type Bundles map[common.ID]OwnerIDs

// Batch contains and points multiple data ID.
// it manages data IDs using trie-like structure.
type Batch struct {
	ID    string
	Count int
	set   Bundles
}

func newBatch(id string) *Batch {
	return &Batch{
		ID:    id,
		Count: 0,
		set:   make(Bundles),
	}
}

// Append adds a data ID to the batch.
func (batch *Batch) Append(dataID common.DataID) {
	// padding is same in data id now.
	ownerIDs := batch.set[dataID.BundleID]
	if ownerIDs == nil {
		ownerIDs = make(OwnerIDs)
	}

	rowIDs := ownerIDs[dataID.OwnerID]
	if rowIDs == nil {
		rowIDs = []common.ID{}
	}
	rowIDs = append(rowIDs, dataID.RowID)

	batch.set[dataID.BundleID][dataID.OwnerID] = rowIDs
	batch.Count += 1
}

// Iterator returns an iterator channel that can be used to
// traverse data IDs in for-range loop.
func (batch *Batch) Iterator() chan common.DataID {
	ch := make(chan common.DataID)
	go func() {
		for bundleID, ownerIDs := range batch.set {
			for ownerID, rowIDs := range ownerIDs {
				for _, rowID := range rowIDs {
					ch <- common.DataID{
						BundleID: bundleID,
						OwnerID:  ownerID,
						RowID:    rowID,
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

// Marshall encodes a batch to the bytes.
func (batch *Batch) Marshal() []byte {
	var csv strings.Builder
	for dataID := range batch.Iterator() {
		csv.WriteString(dataID.String())
		csv.WriteString(",")
	}

	csvStr := csv.String()
	if len(csvStr) > 0 {
		// exclude last comma
		csvStr = csvStr[:len(csvStr)-1]
	}
	return []byte(csvStr)
}

// UnmarshalBatch decodes a batch from the bytes.
func UnmarshalBatch(batchID string, rawBatch []byte) (*Batch, error) {
	dataIDs := strings.Split(string(rawBatch), ",")
	batch := newBatch(batchID)
	for _, rawDataID := range dataIDs {
		dataID, err := common.NewDataID(rawDataID)
		if err != nil {
			return nil, err
		}
		batch.Append(*dataID)
	}
	return batch, nil
}
