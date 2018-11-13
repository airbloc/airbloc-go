package collections

import (
	"github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type Collection struct {
	AppId    ethCommon.Hash
	SchemaId common.ID
	Policy   *IncentivePolicy
}

type IncentivePolicy struct {
	DataProducer  float64
	DataProcessor float64
	DataRelayer   float64
	DataSource    float64
}
