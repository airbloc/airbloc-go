package types

import (
	"encoding/hex"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pkg/errors"
)

type DataId [20]byte

func (id DataId) BundleId() (bundleId ID) {
	copy(bundleId[:], id[:IDLength])
	return
}

func (id DataId) UserId() (userId ID) {
	copy(userId[:], id[IDLength:IDLength*2])
	return
}

func (id DataId) RowId() (rowId RowId) {
	copy(rowId[:], id[IDLength*2:])
	return
}

func NewDataId(dataID string) (id DataId, err error) {
	convert := func(dataID string, index int) string {
		var str string
		if index == 0 {
			str = dataID[:IDStrLength]
		} else {
			str = dataID[IDStrLength*index : IDStrLength*(index+1)]
		}
		return str
	}

	bundleId, err := HexToID(convert(dataID, 0))
	if err != nil {
		err = errors.Wrap(err, "failed to parse bundle ID from the given data ID.")
		return
	}
	copy(id[:IDLength], bundleId[:])

	userId, err := HexToID(convert(dataID, 1))
	if err != nil {
		err = errors.Wrap(err, "failed to parse owner ID from the given data ID")
		return
	}
	copy(id[IDStrLength:IDStrLength*2], userId[:])

	rowId, err := hex.DecodeString(dataID[IDStrLength*2:])
	if err != nil {
		err = errors.Wrap(err, "failed to parse row ID from the given data ID")
		return
	}
	copy(id[IDStrLength*2:], rowId[:])

	return
}

func (id DataId) Hex() string {
	return id.String()
}

func (id DataId) String() string {
	return fmt.Sprintf(
		"%s%s%s",
		id.BundleId().Hex(),
		id.UserId().Hex(),
		id.RowId().Hex(),
	)
}

type RawDataId struct {
	BundleId    string             `json:"bundleId" mapstructure:"bundleId"`
	UserId      string             `json:"userId" mapstructure:"userId"`
	RowId       string             `json:"rowId" mapstructure:"rowId"`
	CollectedAt primitive.DateTime `json:"collectedAt" mapstructure:"collectedAt"`
}

func (id *RawDataId) Convert() (DataId, error) {
	return NewDataId(id.BundleId + id.UserId + id.RowId)
}
