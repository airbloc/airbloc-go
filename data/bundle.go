package data

import (
	"time"

	"golang.org/x/crypto/sha3"

	"github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson"
)

type Bundle struct {
	Id         common.ID `json:"-"`
	Uri        string    `json:"-"`
	Provider   common.ID `json:"provider"`
	Collection common.ID `json:"collection"`
	DataCount  int       `json:"dataCount"`
	IngestedAt time.Time `json:"ingestedAt"`

	Data []*common.EncryptedData `json:"data"`
}

func UnmarshalBundle(bundleData []byte) (*Bundle, error) {
	var bundle Bundle
	err := easyjson.Unmarshal(bundleData, &bundle)
	return &bundle, err
}

func (bundle *Bundle) Marshal() (bundleData []byte, err error) {
	bundleData, err = easyjson.Marshal(bundle)
	return
}

func (bundle *Bundle) Hash() (ethCommon.Hash, error) {
	bundleData, err := bundle.Marshal()
	if err != nil {
		return ethCommon.Hash{}, err
	}
	return sha3.Sum256(bundleData), nil
}
