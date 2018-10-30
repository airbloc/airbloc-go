package collections

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Collection struct {
	AppId    common.Hash
	SchemaId common.Hash
	Policy   *IncentivePolicy
}

type IncentivePolicy struct {
	DataProducer  *big.Int
	DataProcessor *big.Int
	DataRelayer   *big.Int
	DataSource    *big.Int
}
