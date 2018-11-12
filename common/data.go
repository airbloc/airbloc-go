package common

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type Data struct {
	Payload   string `json:"payload"`
	OwnerAnid string `json:"ownerAnid"`
}

type EncryptedData struct {
	OwnerAnid string `json:"ownerAnid"`
	Payload   []byte `json:"payload"`
	Capsule   []byte `json:"capsule"`
}

type DataID struct {
	BundleID ID
	Index    int
}

func NewDataID(dataId string) (*DataID, error) {
	bundleId, err := IDFromString(dataId[:IDLength])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse data ID from the given data ID.")
	}

	index, err := strconv.Atoi(dataId[IDLength+1:])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse data index from the given data ID.")
	}

	return &DataID{
		BundleID: bundleId,
		Index:    index,
	}, nil
}

func (id *DataID) String() string {
	return fmt.Sprintf("%s/%d", id.BundleID.String(), id.Index)
}
