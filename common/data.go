package common

import (
	"fmt"
	"math/big"

	"github.com/pkg/errors"
)

type Data struct {
	Payload   string `json:"payload"`
	OwnerAnid ID     `json:"ownerAnid"`
}

type EncryptedData struct {
	OwnerAnid ID     `json:"ownerAnid"`
	Payload   []byte `json:"payload"`
	Capsule   []byte `json:"capsule"`
}

type DataID struct {
	CollectionID ID
	BundleIndex  *big.Int
	OwnerID      ID
}

func NewDataID(dataId string) (*DataID, error) {
	collectionID, err := HexToID(dataId[:IDLength])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse collection ID from the given data ID.")
	}

	bundleIndex, ok := new(big.Int).SetString(dataId[IDLength+1:IDLength*2], 16)
	if !ok {
		return nil, errors.New("cannot parse hex string to bundleIndex")
	}

	ownerID, err := HexToID(dataId[IDLength*2+1 : IDLength*3])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse owner ID from the given data ID")
	}

	return &DataID{
		CollectionID: collectionID,
		BundleIndex:  bundleIndex,
		OwnerID:      ownerID,
	}, nil
}

func (id *DataID) String() string {
	return fmt.Sprintf(
		"%s%s%s",
		id.CollectionID.Hex(),
		id.BundleIndex.Text(16),
		id.OwnerID.Hex(),
	)
}
