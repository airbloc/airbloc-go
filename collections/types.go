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
	DataProducer  float32
	DataProcessor float32
	DataRelayer   float32
	DataSource    float32
}
