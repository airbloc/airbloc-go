package bind

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
)

// TransactOpts is the collection of authorization data required to create a
// valid Klaytn transaction.
type TransactOpts struct {
	From     common.Address
	FeePayer common.Address
	Signer   bind.SignerFn
	Value    *big.Int
	GasPrice *big.Int
	GasLimit uint64
	TxType   types.TxType
	Context  context.Context
}

func (opts *TransactOpts) MakeTransactionData(client ContractBackend, contract *common.Address, input []byte) (types.TxType, map[types.TxValueKeyType]interface{}, error) {
	var err error

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
	nonce, err := client.PendingNonceAt(opts.Context, opts.From)
	if err != nil {
		return types.TxType(0), nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
	}

	// gas price
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = client.SuggestGasPrice(opts.Context)
		if err != nil {
			return types.TxType(0), nil, fmt.Errorf("failed to suggest gas price: %v", err)
		}
	}

	// gas limit
	if contract != nil {
		if code, err := client.PendingCodeAt(opts.Context, *contract); err != nil {
			return types.TxType(0), nil, err
		} else if len(code) == 0 {
			return types.TxType(0), nil, bind.ErrNoCode
		}
	}

	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		gasLimit, err = client.EstimateGas(opts.Context, klaytn.CallMsg{
			From:  opts.From,
			To:    contract,
			Value: value,
			Data:  input,
		})
		if err != nil {
			return types.TxType(0), nil, fmt.Errorf("failed to estimate gas needed: %v", err)
		}
		gasLimit, _ = new(big.Float).Mul(big.NewFloat(1.5), new(big.Float).SetUint64(gasLimit)).Uint64()
	}

	if contract == nil {
		return types.TxType(0), nil, bind.ErrNoCode
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

	return opts.TxType, values, nil
}

func (opts *TransactOpts) MakeTransaction(client ContractBackend, contract *common.Address, input []byte) (*types.Transaction, error) {
	txType, txValues, err := opts.MakeTransactionData(client, contract, input)
	if err != nil {
		return nil, err
	}

	rawTx, err := types.NewTransactionWithMap(txType, txValues)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	return rawTx, nil
}

func NewKeyedTransactor(key *ecdsa.PrivateKey) *TransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	return &TransactOpts{
		From:   keyAddr,
		Signer: TxRawKeySigner(key),
	}
}

func NewKeyedFeePayerTransactor(key *ecdsa.PrivateKey) *TransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	return &TransactOpts{
		From:   keyAddr,
		Signer: DelegateTxRawKeySigner(key),
	}
}

func NewWalletTransactor(account accounts.Account, wallet accounts.Wallet, passphrase *string) *TransactOpts {
	return &TransactOpts{
		From:   account.Address,
		Signer: TxWalletSigner(account, wallet, passphrase),
	}
}

func NewWalletFeePayerTransactor(account accounts.Account, wallet accounts.Wallet, passphrase *string) *TransactOpts {
	return &TransactOpts{
		From:   account.Address,
		Signer: DelegateTxWalletSigner(account, wallet, passphrase),
	}
}

func MergeTxOpts(ctx context.Context, origin *TransactOpts, opts ...*TransactOpts) *TransactOpts {
	mergedOpts := &TransactOpts{}
	if origin != nil {
		mergedOpts = &TransactOpts{
			From:     origin.From,
			FeePayer: origin.FeePayer,
			Signer:   origin.Signer,
			Value:    origin.Value,
			GasPrice: origin.GasPrice,
			TxType:   origin.TxType,
		}
	}
	mergedOpts.Context = ctx
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if opt.From != (common.Address{}) {
			mergedOpts.From = opt.From
		}
		if opt.FeePayer != (common.Address{}) {
			mergedOpts.FeePayer = opt.FeePayer
		}
		if opt.Signer != nil {
			mergedOpts.Signer = opt.Signer
		}
		if opt.Value != nil {
			mergedOpts.Value = opt.Value
		}
		if opt.GasPrice != nil {
			mergedOpts.GasPrice = opt.GasPrice
		}
		if opt.TxType != 0 {
			mergedOpts.TxType = opt.TxType
		}
	}
	return mergedOpts
}
