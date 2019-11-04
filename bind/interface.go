package bind

import (
	"context"
	"math/big"

	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
)

// ContractBase is an interface that applies all airbloc contracts
type ContractBase interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int
	GetSelectors() map[string]string
}

// ContractBackend is an interface that used by airbloc contracts
type ContractBackend interface {
	bind.ContractBackend
	Transactor(context.Context, ...*TransactOpts) *TransactOpts
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	WaitDeployed(context.Context, *types.Transaction) (*types.Receipt, error)
	GetDeployment(string) (Deployment, bool)
}

// EventIterator is an interface for all return value of contract's filterer methods
type EventIterator interface {
	Next() bool
	Event() interface{}
	Error() error
	Close() error
}
