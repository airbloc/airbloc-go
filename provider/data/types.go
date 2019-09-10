package data

import (
	"github.com/airbloc/airbloc-go/shared/types"
	"go.mongodb.org/mongo-driver/bson"
)

type data struct {
	CollectionId types.ID
	UserId       types.ID
	IngestedAt   types.Time
	Payload      string
}

type bundle struct {
	Id         string   `json:"bundleId" mapstructure:"bundleId"`
	Uri        string   `json:"uri" mapstructure:"uri"`
	Provider   string   `json:"provider" mapstructure:"provider"`
	Collection string   `json:"collection" mapstructure:"collection"`
	IngestedAt int64    `json:"ingestedAt" mapstructure:"ingestedAt"`
	DataIds    []string `json:"-" mapstructure:"-"`
	RawDataIds []bson.D `json:"dataIds" mapstructure:"dataIds"`
}
