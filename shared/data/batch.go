package data

import (
	"github.com/airbloc/airbloc-go/shared/types"
	"strings"
)

type UserIds map[types.ID][][4]byte
type Bundles map[types.ID]UserIds

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
func (batch *Batch) Append(dataID types.DataId) {
	// padding is same in data id now.
	userIds := batch.set[dataID.BundleId]
	if userIds == nil {
		userIds = make(UserIds)
	}

	rowIDs := userIds[dataID.UserId]
	if rowIDs == nil {
		rowIDs = [][4]byte{}
	}
	rowIDs = append(rowIDs, dataID.RowId)

	batch.set[dataID.BundleId][dataID.UserId] = rowIDs
	batch.Count += 1
}

// Iterator returns an iterator channel that can be used to
// traverse data IDs in for-range loop.
func (batch *Batch) Iterator() chan types.DataId {
	ch := make(chan types.DataId)
	go func() {
		for bundleID, userIds := range batch.set {
			for userId, rowIds := range userIds {
				for _, rowId := range rowIds {
					ch <- types.DataId{
						BundleId: bundleID,
						UserId:   userId,
						RowId:    rowId,
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
	for _, rawDataId := range dataIDs {
		dataID, err := types.NewDataId(rawDataId)
		if err != nil {
			return nil, err
		}
		batch.Append(*dataID)
	}
	return batch, nil
}
