package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/pkg/errors"
)

type TransactOpts struct {
	From     common.Address
	FeePayer common.Address
	Signer   bind.SignerFn
	Nonce    *big.Int
	Value    *big.Int
	GasPrice *big.Int
	GasLimit uint64
	TxType   types.TxType
	Context  context.Context
}

func NewKeyedTransactor(key *ecdsa.PrivateKey) *TransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	return &TransactOpts{
		From: keyAddr,
		Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if addr != keyAddr {
				return nil, errors.New("unauthorized sign request")
			}
			return types.SignTx(tx, signer, key)
		},
	}
}

type BoundContract struct {
	address common.Address
	abi     abi.ABI
	client  TxClient
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
		address: address,
		abi:     abi,
		client:  transactor.(TxClient),
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

	delegated := opts.FeePayer != (common.Address{})
	contractExec := input != nil

	switch {
	case contractExec && delegated:
		opts.TxType = types.TxTypeFeeDelegatedSmartContractExecution
	case !contractExec && delegated:
		opts.TxType = types.TxTypeFeeDelegatedValueTransfer
	case contractExec && !delegated:
		opts.TxType = types.TxTypeSmartContractExecution
	case !contractExec && !delegated:
		opts.TxType = types.TxTypeValueTransfer
	}

	// value
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	// nonce
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = c.client.PendingNonceAt(opts.Context, opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// gas price
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = c.client.SuggestGasPrice(opts.Context)
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas price: %v", err)
		}
	}

	// gas limit
	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		if contract != nil {
			if code, err := c.client.PendingCodeAt(opts.Context, c.address); err != nil {
				return nil, err
			} else if len(code) == 0 {
				return nil, bind.ErrNoCode
			}
		}

		from := opts.From
		if delegated {
			from = opts.FeePayer
		}

		gasLimit, err = c.client.EstimateGas(opts.Context, klaytn.CallMsg{
			From:  from,
			To:    contract,
			Value: value,
			Data:  input,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
		}
	}
	gasLimit, _ = new(big.Float).Mul(big.NewFloat(1.5), new(big.Float).SetUint64(gasLimit)).Uint64()

	if contract == nil {
		return nil, bind.ErrNoCode
	}

	values := map[types.TxValueKeyType]interface{}{
		types.TxValueKeyNonce:    nonce,
		types.TxValueKeyFrom:     opts.From,
		types.TxValueKeyTo:       *contract,
		types.TxValueKeyAmount:   value,
		types.TxValueKeyGasLimit: gasLimit,
		types.TxValueKeyGasPrice: gasPrice,
	}

	if delegated {
		values[types.TxValueKeyFeePayer] = opts.FeePayer
	}

	if contractExec {
		values[types.TxValueKeyData] = input
	}

	rawTx, err := types.NewTransactionWithMap(opts.TxType, values)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}

	signedTx, err := c.client.SignTransaction(opts.Context, opts.Signer, rawTx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign transaction")
	}

	if err = c.client.SendTransaction(opts.Context, signedTx); err != nil {
		return nil, err
	}
	return signedTx, nil
}
