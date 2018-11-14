package collections

import (
	"github.com/airbloc/airbloc-go/common"
)

type Collection struct {
	AppId    common.ID
	SchemaId common.ID
	Policy   *IncentivePolicy
}

type IncentivePolicy struct {
	DataProvider  float64
	DataProcessor float64
	DataRelayer   float64
	DataOwner     float64
}
