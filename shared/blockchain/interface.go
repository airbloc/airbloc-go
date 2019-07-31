package blockchain

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractConstructor func(common.Address, common.Hash, *big.Int, bind.ContractBackend) (interface{}, error)

type TxClient interface {
	bind.ContractBackend
	Account() *bind.TransactOpts
	GetContract(interface{}) interface{}
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	WaitDeployed(context.Context, *types.Transaction) (*types.Receipt, error)
}
