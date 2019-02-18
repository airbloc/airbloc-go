package data

import (
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
