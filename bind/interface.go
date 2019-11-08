package bind

import (
	"context"

	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
)

// ContractBackend is an interface that used by airbloc contracts
type ContractBackend interface {
	bind.ContractBackend
	Deployment(string) (Deployment, bool)
	Transactor(context.Context, ...*TransactOpts) *TransactOpts
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	WaitDeployed(context.Context, *types.Transaction) (*types.Receipt, error)
}

// EventIterator is an interface for all return value of contract's filterer methods
type EventIterator interface {
	Next() bool
	Event() interface{}
	Error() error
	Close() error
}
