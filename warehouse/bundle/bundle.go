package bundle

import (
	"time"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/data"
	"github.com/mailru/easyjson"
)

type Bundle struct {
	Id         common.ID `json:"id"`
	Uri        string    `json:"uri"`
	Provider   common.ID `json:"provider"`
	Collection common.ID `json:"collection"`
	DataCount  int       `json:"dataCount"`
	IngestedAt time.Time `json:"ingestedAt"`

	Data []*data.EncryptedData `json:"data"`
}

func Unmarshal(bundleData []byte) (*Bundle, error) {
	var bundle Bundle
	err := easyjson.Unmarshal(bundleData, &bundle)
	return &bundle, err
}

func (bundle *Bundle) Marshal() (bundleData []byte, err error) {
	bundleData, err = easyjson.Marshal(bundle)
	return
}
