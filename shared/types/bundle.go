package types

import (
	"github.com/airbloc/airbloc-go/shared/merkle"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

type Bundle struct {
	Id         string `json:"-"`
	Uri        string `json:"-"`
	Provider   ID     `json:"provider"`
	Collection ID     `json:"collection"`
	DataCount  int    `json:"dataCount"`
	IngestedAt Time   `json:"ingestedAt"`

	// mapping(userId => []data)
	Data map[ID][]*EncryptedData `json:"data"`
	tree *merkle.MainTree        `json:"-"`
}

func (bundle *Bundle) Hash() (ethCommon.Hash, error) {
	bundleData, err := json.Marshal(bundle)
	if err != nil {
		return ethCommon.Hash{}, err
	}
	return sha3.Sum256(bundleData), nil
}

func (bundle *Bundle) SetupRowId() error {
	if bundle.Data == nil {
		return errors.New("nil data")
	}

	for userId, rowData := range bundle.Data {
		for i := range rowData {
			bundle.Data[userId][i].RowId = UintToRowId(uint32(i))
		}
	}
	return nil
}

type marshalBundle struct {
	Provider   ID            `json:"provider"`
	Collection ID            `json:"collection"`
	DataCount  int           `json:"dataCount"`
	IngestedAt Time          `json:"ingestedAt"`
	Data       []marshalData `json:"data"`
}

type marshalData struct {
	UserId ID               `json:"userId"`
	Data   []*EncryptedData `json:"data"`
}

func (bundle *Bundle) UnmarshalJSON(d []byte) error {
	data := new(marshalBundle)
	err := json.Unmarshal(d, &data)
	if err != nil {
		return err
	}

	bundle.Provider = data.Provider
	bundle.Collection = data.Collection
	bundle.DataCount = data.DataCount
	bundle.IngestedAt = data.IngestedAt

	bundle.Data = make(map[ID][]*EncryptedData, len(data.Data))
	for _, encryptedData := range data.Data {
		bundle.Data[encryptedData.UserId] = encryptedData.Data
	}

	return nil
}

func (bundle *Bundle) MarshalJSON() ([]byte, error) {
	data := make([]marshalData, len(bundle.Data))

	i := 0
	for userId, encryptedData := range bundle.Data {
		data[i], i = marshalData{
			UserId: userId,
			Data:   encryptedData,
		}, i+1
	}

	return json.Marshal(&marshalBundle{
		Provider:   bundle.Provider,
		Collection: bundle.Collection,
		DataCount:  bundle.DataCount,
		IngestedAt: bundle.IngestedAt,
		Data:       data,
	})
}

func (bundle *Bundle) generateSMT() (err error) {
	leaves := make(map[[8]byte][][4]byte, len(bundle.Data))
	for userId, rowData := range bundle.Data {
		leaves[userId] = make([][4]byte, len(rowData))
		for i, data := range rowData {
			leaves[userId][i] = data.RowId
		}
	}

	bundle.tree, err = merkle.NewMainTree(leaves)
	return
}

// SetupUserProof creates a root of 64-depth SMT (Sparse Merkle Tree),
// which can be used as an accumulator of User IDs for the bundle.
func (bundle *Bundle) SetupUserProof() (root ethCommon.Hash, _ error) {
	if bundle.tree == nil {
		if err := bundle.generateSMT(); err != nil {
			return root, errors.Wrap(err, "setup user proof")
		}
	}
	root = bundle.tree.Root()
	return
}

func (bundle *Bundle) GenerateProof(rowId RowId, userId ID) ([]byte, error) {
	if bundle.tree == nil {
		if err := bundle.generateSMT(); err != nil {
			return nil, errors.Wrap(err, "setup user proof")
		}
	}
	return bundle.tree.GenerateProof(rowId, userId)
}
