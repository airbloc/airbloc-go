package bind

import (
	"context"

	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
)

type ContractBackend interface {
	bind.ContractBackend
	Transactor(context.Context, ...*TransactOpts) *TransactOpts
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	WaitDeployed(context.Context, *types.Transaction) (*types.Receipt, error)
}

type EventIterator interface {
	Next() bool
	Event() interface{}
	Error() error
	Close() error
}
