package collections

import (
	"github.com/ethereum/go-ethereum/common"
)

type Collection struct {
	AppId    common.Hash
	SchemaId common.Hash
	Policy   *IncentivePolicy
}

type IncentivePolicy struct {
	DataProducer  float64
	DataProcessor float64
	DataRelayer   float64
	DataSource    float64
}
