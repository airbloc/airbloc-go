package data

import (
	"errors"
	"github.com/airbloc/airbloc-go/merkle"
	"github.com/json-iterator/go"
	"golang.org/x/crypto/sha3"

	"github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Bundle struct {
	Id         string      `json:"-"`
	Uri        string      `json:"-"`
	Provider   common.ID   `json:"provider"`
	Collection common.ID   `json:"collection"`
	DataCount  int         `json:"dataCount"`
	IngestedAt common.Time `json:"ingestedAt"`

	// mapping(userId => []data)
	Data map[common.ID][]*common.EncryptedData `json:"data"`
	tree *merkle.MainTree                      `json:"-"`
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
			bundle.Data[userId][i].RowId = common.UintToRowId(uint32(i))
		}
	}
	return nil
}

type marshalBundle struct {
	Provider   common.ID     `json:"provider"`
	Collection common.ID     `json:"collection"`
	DataCount  int           `json:"dataCount"`
	IngestedAt common.Time   `json:"ingestedAt"`
	Data       []marshalData `json:"data"`
}

type marshalData struct {
	UserId common.ID               `json:"userId"`
	Data   []*common.EncryptedData `json:"data"`
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

	bundle.Data = make(map[common.ID][]*common.EncryptedData, len(data.Data))
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
