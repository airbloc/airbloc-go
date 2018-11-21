package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractConstructor interface {
	New(common.Address, bind.ContractBackend) (interface{}, error)
}

type TxClient interface {
	bind.ContractBackend
	Account() *bind.TransactOpts
	GetContract(interface{}) (interface{}, error)
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	WaitDeployed(context.Context, *types.Transaction) (*types.Receipt, error)
}
