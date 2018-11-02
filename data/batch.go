package data

import (
	"github.com/airbloc/airbloc-go/common"
	"strings"
)

// Batch contains and points multiple data ID.
// it manages data IDs using trie-like structure.
type Batch struct {
	Id    string
	Count int
	set   map[common.ID][]int
}

func newBatch(id string) *Batch {
	return &Batch{
		Id:    id,
		Count: 0,
		set:   map[common.ID][]int{},
	}
}

// Append adds a data ID to the batch.
func (batch *Batch) Append(dataId ID) {
	indices := batch.set[dataId.BundleID]
	if indices == nil {
		indices = []int{}
	}
	indices = append(indices, dataId.Index)

	batch.set[dataId.BundleID] = indices
	batch.Count += 1
}

// Iterator returns an iterator channel that can be used to
// traverse data IDs in for-range loop.
func (batch *Batch) Iterator() chan ID {
	ch := make(chan ID)
	go func() {
		for bundleId, indices := range batch.set {
			for _, index := range indices {
				ch <- ID{bundleId, index}
			}
		}
		close(ch)
	}()
	return ch
}

// Marshall encodes a batch to the bytes.
func (batch *Batch) Marshal() []byte {
	var csv strings.Builder
	for dataId := range batch.Iterator() {
		csv.WriteString(dataId.String())
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
func UnmarshalBatch(batchId string, rawBatch []byte) (*Batch, error) {
	dataIds := strings.Split(string(rawBatch), ",")
	batch := newBatch(batchId)
	for _, rawDataId := range dataIds {
		dataId, err := NewDataID(rawDataId)
		if err != nil {
			return nil, err
		}
		batch.Append(*dataId)
	}
	return batch, nil
}
