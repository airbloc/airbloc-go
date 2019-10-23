package blockchain

import (
	"context"
	"math/big"

	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
)

type ContractConstructor func(common.Address, common.Hash, *big.Int, abi.ABI, bind.ContractBackend) interface{}

type TxClient interface {
	bind.ContractBackend
	Account(context.Context, ...*TransactOpts) *TransactOpts
	GetContract(interface{}) interface{}
	SignTransaction(context.Context, bind.SignerFn, *types.Transaction) (*types.Transaction, error)
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	WaitDeployed(context.Context, *types.Transaction) (*types.Receipt, error)
}

type EventIterator interface {
	Next() bool
	Event() interface{}
	Error() error
	Close() error
}
