package blockchain

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/params"
)

var (
	supportTxTypes = map[types.TxType]interface{}{
		types.TxTypeValueTransfer:                      "",
		types.TxTypeFeeDelegatedValueTransfer:          "",
		types.TxTypeSmartContractExecution:             "",
		types.TxTypeFeeDelegatedSmartContractExecution: "",
	}
)

type TransactOpts struct {
	*bind.TransactOpts
	FeePayer common.Address
	TxType   types.TxType
}

type BoundContract struct {
	address    common.Address
	abi        abi.ABI
	transactor bind.ContractTransactor
	*bind.BoundContract
}

func NewBoundContract(
	address common.Address,
	abi abi.ABI,
	caller bind.ContractCaller,
	transactor bind.ContractTransactor,
	filterer bind.ContractFilterer,
) *BoundContract {
	return &BoundContract{
		address:    address,
		abi:        abi,
		transactor: transactor,
		BoundContract: bind.NewBoundContract(
			address,
			abi,
			caller,
			transactor,
			filterer,
		),
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
	var err error
	if opts == nil {
		return nil, errors.New("nil transcatOpts")
	}

	// value
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	// nonce
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = c.transactor.PendingNonceAt(opts.Context, opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// gas price
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = c.transactor.SuggestGasPrice(opts.Context)
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas price: %v", err)
		}
	}

	// gas limit
	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		if contract != nil {
			if code, err := c.transactor.PendingCodeAt(opts.Context, c.address); err != nil {
				return nil, err
			} else if len(code) == 0 {
				return nil, bind.ErrNoCode
			}
		}

		msg := klaytn.CallMsg{From: opts.From, To: contract, Value: value, Data: input}
		gasLimit, err = c.transactor.EstimateGas(opts.Context, msg)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
		}
	}

	values := map[types.TxValueKeyType]interface{}{
		types.TxValueKeyTo:     contract,
		types.TxValueKeyAmount: value,
		types.TxValueKeyNonce:  nonce,
		types.TxValueKeyGasPrice: new(big.Int).Add(
			gasPrice, new(big.Int).Mul(
				big.NewInt(5),
				big.NewInt(params.Ston),
			),
		),
		types.TxValueKeyGasLimit: gasLimit,
	}

	if _, ok := supportTxTypes[opts.TxType]; !ok {
		return nil, errors.New("unsupported transaction type")
	}

	if opts.TxType == types.TxTypeFeeDelegatedValueTransfer ||
		opts.TxType == types.TxTypeFeeDelegatedSmartContractExecution {
		values[types.TxValueKeyFeePayer] = opts.FeePayer
	}

	if opts.TxType == types.TxTypeSmartContractExecution ||
		opts.TxType == types.TxTypeFeeDelegatedSmartContractExecution {
		values[types.TxValueKeyData] = input
	}

	rawTx, err := types.NewTransactionWithMap(opts.TxType, values)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}

	chainID, err := c.transactor.ChainID(opts.Context)
	if err != nil {
		return nil, err
	}

	signer := types.NewEIP155Signer(chainID)
	signedTx, err := opts.Signer(signer, opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	if err := c.transactor.SendTransaction(opts.Context, signedTx); err != nil {
		return nil, err
	}
	return signedTx, nil
}
