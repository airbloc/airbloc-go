package bind

import (
	"encoding/json"

	"github.com/airbloc/logger"

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
	log     logger.Logger
	*bind.BoundContract
}

func NewBoundContract(
	address common.Address,
	abi abi.ABI,
	name string,
	backend ContractBackend,
) *BoundContract {
	return &BoundContract{
		address:       address,
		abi:           abi,
		client:        backend,
		log:           logger.New(name),
		BoundContract: bind.NewBoundContract(address, abi, backend, backend.Client(), backend),
	}
}

func (c *BoundContract) logTx(method string, params ...interface{}) {
	abiMethod := c.abi.Methods[method]

	attr := make(map[string]interface{})
	for index := range abiMethod.Inputs {
		attr[abiMethod.Inputs[index].Name] = params[index]
	}
	d, err := json.Marshal(attr)
	if err != nil {
		c.log.Error("failed to log tx. err={}", err)
		return
	}

	c.log.Info(abiMethod.Sig()+" input={}", string(d))
}

func (c *BoundContract) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	c.logTx(method, params...)
	return c.BoundContract.Call(opts, result, method, params...)
}

func (c *BoundContract) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Receipt, error) {
	c.logTx(method, params...)
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	return c.transact(opts, &c.address, input)
}

func (c *BoundContract) Transfer(opts *TransactOpts) (*types.Receipt, error) {
	return c.transact(opts, &c.address, nil)
}

func (c *BoundContract) transact(opts *TransactOpts, contract *common.Address, input []byte) (*types.Receipt, error) {
	if opts == nil {
		opts = &TransactOpts{}
	}

	opts = c.client.Transactor(opts.Context, opts)

	unsignedTx, err := c.client.MakeTransaction(opts, contract, input)
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

	return c.client.SendTransaction(opts.Context, signedTx)
}
