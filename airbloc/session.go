package airbloc

import (
	"context"
	"fmt"

	ablbind "github.com/airbloc/airbloc-go/bind"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

type Session struct {
	*Client

	account     Account
	deployments ablbind.Deployments
}

func NewSession(cfg Config) (Session, error) {
	session := Session{
		Client:      cfg.Client,
		account:     cfg.Account,
		deployments: cfg.Deployments,
	}
	return session, nil
}

func (sess Session) Deployment(contract string) (ablbind.Deployment, bool) {
	return sess.deployments.Get(contract)
}

func (sess Session) Transactor(ctx context.Context, opts ...*ablbind.TransactOpts) *ablbind.TransactOpts {
	return ablbind.MergeTxOpts(ctx, sess.account.txOpts(), opts...)
}

func (sess Session) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if sess.account.isReadOnly() {
		return errors.New("session is on read-only mode")
	}

	txType := tx.Type()

	if txType == types.TxTypeValueTransfer ||
		txType == types.TxTypeSmartContractExecution {
		return sess.Client.SendTransaction(ctx, tx)
	}

	if !sess.account.isDelegated() {
		return errors.New("session is non-delegate mode")
	}

	if txType == types.TxTypeFeeDelegatedValueTransfer ||
		txType == types.TxTypeFeeDelegatedSmartContractExecution {
		return sess.account.SendTransaction(ctx, tx)
	}
	return errors.New("invalid transaction type")
}

func (sess Session) MakeTransaction(opts *ablbind.TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	pTxType, txValues, err := opts.MakeTransactionData(sess, contract, input)
	if err != nil {
		return nil, err
	}
	if pTxType == nil {
		return nil, errors.New("unsupported tx type")
	}

	txType := *pTxType

	if sess.account.isDelegated() {
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
