package collections

import (
	"github.com/airbloc/airbloc-go/shared/schemas"
	"github.com/airbloc/airbloc-go/shared/types"
)

type Collection struct {
	Id types.ID

	// owner of the collection
	AppId           types.ID
	IncentivePolicy IncentivePolicy

	// format of the data the app want to collect
	Schema    schemas.Schema
	CreatedAt types.Time
}

// IncentivePolicy is a policy about sharing revenue of the data
// between each stakeholders at the data delivery pipeline.
type IncentivePolicy struct {
	DataProvider  float64
	DataProcessor float64
	DataRelayer   float64
	DataOwner     float64
}

func NewCollection(appId types.ID, schemaId types.ID, policy IncentivePolicy) *Collection {
	collection := &Collection{
		AppId:           appId,
		IncentivePolicy: policy,
	}
	collection.Schema.Id = schemaId
	return collection
}
