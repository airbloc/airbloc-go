package types

import "github.com/ethereum/go-ethereum/common"

type App struct {
	Name       string         "json:\"name\""
	Owner      common.Address "json:\"owner\""
	HashedName common.Hash    "json:\"hashedName\""
}
