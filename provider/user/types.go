package user

import (
	"github.com/airbloc/airbloc-go/shared/types"
	"go.mongodb.org/mongo-driver/bson"
)

type userDataPayload struct {
	types.DataId
	CollectedAt int64  `json:"collectedAt"`
	IngestedAt  int64  `json:"ingestedAt"`
	Payload     string `json:"payload"`
}

type userData struct {
	CollectionId string            `json:"collectionId"`
	Data         []userDataPayload `json:"data" mapstructure:"-"`
}

type userDataInfo struct {
	CollectionId string `json:"collection" mapstructure:"collection"`
	IngestedAt   int64  `json:"ingestedAt"`
	DataIds      []struct {
		types.DataId
		CollectedAt int64 `json:"collectedAt"`
	} `json:"dataIds" mapstructure:"-"`
	RawDataIds []bson.D `json:"rawDataIds" mapstructure:"dataIds"`
}
