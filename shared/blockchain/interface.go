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
	Account() *bind.TransactOpts
	GetContract(interface{}) interface{}
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	WaitDeployed(context.Context, *types.Transaction) (*types.Receipt, error)
}
