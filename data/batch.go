package data

import (
	"math/big"
	"strings"

	"github.com/airbloc/airbloc-go/common"
)

// Batch contains and points multiple data ID.
// it manages data IDs using trie-like structure.
type Batch struct {
	ID    string
	Count int
	set   map[common.ID]map[*big.Int][]common.ID
}

func newBatch(id string) *Batch {
	return &Batch{
		ID:    id,
		Count: 0,
		set:   make(map[common.ID]map[*big.Int][]common.ID),
	}
}

// Append adds a data ID to the batch.
func (batch *Batch) Append(dataID common.DataID) {
	bundles := batch.set[dataID.CollectionID]
	if bundles == nil {
		bundles = make(map[*big.Int][]common.ID)
	}

	ownerIDs := bundles[dataID.BundleIndex]
	if ownerIDs == nil {
		ownerIDs = []common.ID{}
	}
	ownerIDs = append(ownerIDs, dataID.OwnerID)

	batch.set[dataID.CollectionID][dataID.BundleIndex] = ownerIDs
	batch.Count += 1
}

// Iterator returns an iterator channel that can be used to
// traverse data IDs in for-range loop.
func (batch *Batch) Iterator() chan common.DataID {
	ch := make(chan common.DataID)
	go func() {
		for collectionID, bundles := range batch.set {
			for bundleIndex, ownerIDs := range bundles {
				for _, ownerID := range ownerIDs {
					ch <- common.DataID{
						CollectionID: collectionID,
						BundleIndex:  bundleIndex,
						OwnerID:      ownerID,
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
