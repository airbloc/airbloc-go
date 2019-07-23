package types

import "github.com/ethereum/go-ethereum/common"

type App struct {
	Name       string
	Owner      common.Address
	HashedName [32]byte
}
