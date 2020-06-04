package bind

import (
	"context"
	"math/big"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type ContractTransactor interface {
	// PendingCodeAt returns the code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	// PendingNonceAt retrieves the current pending nonce associated with an account.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// EstimateGas tries to estimate the gas needed to execute a specific
	// transaction based on the current pending state of the backend blockchain.
	// There is no guarantee that this is the true gas limit requirement as other
	// transactions may be added or removed by miners, but it should provide a basis
	// for setting a reasonable default.
	EstimateGas(ctx context.Context, call klaytn.CallMsg) (gas uint64, err error)
	// SendTransaction injects the transaction into the pending pool for execution.
	SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)
	// ChainID can return the chain ID of the chain.
	ChainID(ctx context.Context) (*big.Int, error)
}

// ContractBackend is an interface that used by airbloc contracts
type ContractBackend interface {
	bind.ContractCaller
	bind.ContractFilterer
	ContractTransactor
	Client() *client.Client
	Deployment(string) (Deployment, bool)
	Transactor(context.Context, ...*TransactOpts) *TransactOpts
	MakeTransaction(*TransactOpts, *common.Address, []byte) (*types.Transaction, error)
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	WaitMinedWithHash(context.Context, common.Hash) (*types.Receipt, error)
	WaitDeployed(context.Context, *types.Transaction) (*types.Receipt, error)
}

// EventIterator is an interface for all return value of contract's filterer methods
type EventIterator interface {
	Next() bool
	Event() interface{}
	Error() error
	Close() error
}
