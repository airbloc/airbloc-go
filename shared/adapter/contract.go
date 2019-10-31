package adapter

import (
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

type BoundContract struct {
	address common.Address
	abi     abi.ABI
	client  ContractBackend
	*bind.BoundContract
}

func NewBoundContract(
	address common.Address,
	abi abi.ABI,
	backend ContractBackend,
) *BoundContract {
	return &BoundContract{
		address:       address,
		abi:           abi,
		client:        backend,
		BoundContract: bind.NewBoundContract(address, abi, backend, backend, backend),
	}
}

func (c *BoundContract) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	return c.transact(opts, &c.address, input)
}

func (c *BoundContract) Transfer(opts *TransactOpts) (*types.Transaction, error) {
	return c.transact(opts, &c.address, nil)
}

func (c *BoundContract) transact(opts *TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	if opts == nil {
		opts = &TransactOpts{}
	}

	unsignedTx, err := opts.MakeTransaction(c, contract, input)
	if err != nil {
		return nil, err
	}

	chainID, err := c.client.ChainID(opts.Context)
	if err != nil {
		return nil, errors.Wrap(err, "get chain id")
	}

	signedTx, err := opts.Signer(types.NewEIP155Signer(chainID), opts.From, unsignedTx)
	if err != nil {
		return nil, errors.Wrap(err, "signing transaction")
	}

	if err = c.client.SendTransaction(opts.Context, signedTx); err != nil {
		return nil, err
	}
	return signedTx, nil
}
