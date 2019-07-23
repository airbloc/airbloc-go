package types

import "github.com/ethereum/go-ethereum/common"

type DataType struct {
	Name       string
	Owner      common.Address
	SchemaHash [32]byte
}
