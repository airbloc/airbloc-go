package types

import "github.com/ethereum/go-ethereum/common"

type DataType struct {
	Name       string         "json:\"name\""
	Owner      common.Address "json:\"owner\""
	SchemaHash common.Hash    "json:\"schemaHash\""
}
