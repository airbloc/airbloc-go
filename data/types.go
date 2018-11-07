package data

import (
	"fmt"
	"strconv"

	"github.com/airbloc/airbloc-go/common"
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

type ID struct {
	BundleID common.ID
	Index    int
}

func NewDataID(dataId string) (*ID, error) {
	bundleId, err := common.IDFromString(dataId[:common.IDLength])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse data ID from the given data ID.")
	}

	index, err := strconv.Atoi(dataId[common.IDLength+1:])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse data index from the given data ID.")
	}

	return &ID{
		BundleID: bundleId,
		Index:    index,
	}, nil
}

func (id *ID) String() string {
	return fmt.Sprintf("%s/%d", id.BundleID.String(), id.Index)
}
