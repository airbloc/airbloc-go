package bundle

import (
	"time"

	"github.com/airbloc/airbloc-go/common"
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

func Unmarshal(bundleData []byte) (*Bundle, error) {
	var bundle Bundle
	err := easyjson.Unmarshal(bundleData, &bundle)
	return &bundle, err
}

func (bundle *Bundle) Marshal() (bundleData []byte, err error) {
	bundleData, err = easyjson.Marshal(bundle)
	return
}
