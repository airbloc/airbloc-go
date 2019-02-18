package common

import (
	"encoding/hex"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/pkg/errors"
)

type Data struct {
	UserId      ID     `json:"userId"`
	RowId       RowId  `json:"rowId"`
	CollectedAt Time   `json:"collecedAt"`
	Payload     string `json:"payload"`
}

type EncryptedData struct {
	UserId      ID     `json:"userId"`
	RowId       RowId  `json:"rowId"`
	Capsule     []byte `json:"capsule"`
	CollectedAt Time   `json:"collectedAt"`
	Payload     []byte `json:"payload"`
}

type DataId struct {
	BundleId ID    `json:"bundleId"`
	UserId   ID    `json:"userId"`
	RowId    RowId `json:"rowId"`
}

func convert(dataID string, index int) string {
	var str string
	if index == 0 {
		str = dataID[:IDStrLength]
	} else {
		str = dataID[IDStrLength*index : IDStrLength*(index+1)]
	}
	return str
}

func NewDataId(dataID string) (id *DataId, err error) {
	id = new(DataId)

	id.BundleId, err = HexToID(convert(dataID, 0))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse bundle ID from the given data ID.")
	}

	id.UserId, err = HexToID(convert(dataID, 1))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse owner ID from the given data ID")
	}

	rowId, err := hex.DecodeString(dataID[IDStrLength*2:])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse row ID from the given data ID")
	}
	copy(id.RowId[:], rowId)

	return
}

func (id *DataId) Hex() string {
	return id.String()
}

func (id *DataId) String() string {
	return fmt.Sprintf(
		"%s%s%s",
		id.BundleId.Hex(),
		id.UserId.Hex(),
		hex.EncodeToString(id.RowId[:]),
	)
}

type RawDataId struct {
	BundleId    string             `json:"bundleId" mapstructure:"bundleId"`
	UserId      string             `json:"userId" mapstructure:"userId"`
	RowId       string             `json:"rowId" mapstructure:"rowId"`
	CollectedAt primitive.DateTime `json:"collectedAt" mapstructure:"collectedAt"`
	IngestedAt  primitive.DateTime `json:"ingestedAt" mapstructure:"ingestedAt"`
}

func (id *RawDataId) Convert() (*DataId, error) {
	return NewDataId(id.BundleId + id.UserId + id.RowId)
}
