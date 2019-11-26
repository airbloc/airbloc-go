package session

import (
	"context"
	"fmt"

	"github.com/airbloc/airbloc-go/account"
	ablbind "github.com/airbloc/airbloc-go/bind"

	"github.com/klaytn/klaytn/blockchain/types"
	klayClient "github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

type Session struct {
	ablbind.ContractBackend

	account     account.Account
	deployments ablbind.Deployments
}

func NewSession(cfg Config) (Session, error) {
	sess := Session{
		ContractBackend: cfg.Client,
		account:         cfg.Account,
		deployments:     cfg.Deployments,
	}
	return sess, nil
}

func (sess Session) Client() *klayClient.Client {
	return sess.ContractBackend.Client()
}

func (sess Session) Deployment(contract string) (ablbind.Deployment, bool) {
	return sess.deployments.Get(contract)
}

func (sess Session) Transactor(ctx context.Context, opts ...*ablbind.TransactOpts) *ablbind.TransactOpts {
	return ablbind.MergeTxOpts(ctx, sess.account.TxOpts(), opts...)
}

func (sess Session) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	hash, err := sess.sendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	return sess.WaitMinedWithHash(ctx, hash)
}

func (sess Session) sendTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	if sess.account.IsReadOnly() {
		return common.Hash{}, errors.New("session is on read-only mode")
	}

	txType := tx.Type()

	if txType == types.TxTypeValueTransfer ||
		txType == types.TxTypeSmartContractExecution {
		return sess.Client().SendRawTransaction(ctx, tx)
	}

	if !sess.account.IsDelegated() {
		return common.Hash{}, errors.New("session is non-delegate mode")
	}

	if txType == types.TxTypeFeeDelegatedValueTransfer ||
		txType == types.TxTypeFeeDelegatedSmartContractExecution {
		return sess.account.SendTransaction(ctx, tx)
	}
	return common.Hash{}, errors.New("invalid transaction type")
}

func (sess Session) MakeTransaction(opts *ablbind.TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	txType, txValues, err := opts.MakeTransactionData(sess, contract, input)
	if err != nil {
		return nil, err
	}

	if sess.account.IsDelegated() {
		switch txType {
		case types.TxTypeValueTransfer:
			txType = types.TxTypeFeeDelegatedValueTransfer
		case types.TxTypeSmartContractExecution:
			txType = types.TxTypeFeeDelegatedSmartContractExecution
		}
		txValues[types.TxValueKeyFeePayer] = common.Address{}
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
