package common

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type Data struct {
	OwnerAnID ID     `json:"ownerAnId"`
	RowID     ID     `json:"rowId"`
	Payload   string `json:"payload"`
}

type EncryptedData struct {
	OwnerAnID ID     `json:"ownerAnId"`
	RowID     ID     `json:"rowId"`
	Capsule   []byte `json:"capsule"`
	Payload   []byte `json:"payload"`
}

type DataID struct {
	Empty    ID `json:"empty"`
	BundleID ID `json:"bundleId"`
	OwnerID  ID `json:"ownerId"`
	RowID    ID `json:"rowId"`
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

func NewDataID(dataID string) (*DataID, error) {
	log.Println(dataID)
	empty, err := HexToID(convert(dataID, 0))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse {empty} from the given data ID.")
	}

	bundleID, err := HexToID(convert(dataID, 1))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse bundle ID from the given data ID.")
	}

	ownerID, err := HexToID(convert(dataID, 2))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse owner ID from the given data ID")
	}

	rowID, err := HexToID(convert(dataID, 3))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse row ID from the given data ID")
	}

	return &DataID{
		Empty:    empty,
		BundleID: bundleID,
		OwnerID:  ownerID,
		RowID:    rowID,
	}, nil
}

func (id *DataID) String() string {
	return fmt.Sprintf(
		"%s%s%s%s",
		id.Empty.Hex(),
		id.BundleID.Hex(),
		id.OwnerID.Hex(),
		id.RowID.Hex(),
	)
}
