package airbloc

import (
	"context"
	"fmt"

	ablbind "github.com/airbloc/airbloc-go/bind"

	"github.com/klaytn/klaytn/blockchain/types"
	klayClient "github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

type sessionData struct{ *Client }
type Session struct {
	sessionData

	account     Account
	deployments ablbind.Deployments
}

func NewSession(cfg Config) (Session, error) {
	sess := Session{
		sessionData: sessionData{cfg.Client},
		account:     cfg.Account,
		deployments: cfg.Deployments,
	}
	return sess, nil
}

func (sess Session) Client() *klayClient.Client {
	return sess.sessionData.clientData.Client
}

func (sess Session) Deployment(contract string) (ablbind.Deployment, bool) {
	return sess.deployments.Get(contract)
}

func (sess Session) Transactor(ctx context.Context, opts ...*ablbind.TransactOpts) *ablbind.TransactOpts {
	return ablbind.MergeTxOpts(ctx, sess.account.txOpts(), opts...)
}

func (sess Session) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	if sess.account.isReadOnly() {
		return nil, errors.New("session is on read-only mode")
	}

	txType := tx.Type()

	if txType == types.TxTypeValueTransfer ||
		txType == types.TxTypeSmartContractExecution {
		hash, err := sess.SendRawTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
		return sess.WaitMinedWithHash(ctx, hash)
	}

	if !sess.account.isDelegated() {
		return nil, errors.New("session is non-delegate mode")
	}

	if txType == types.TxTypeFeeDelegatedValueTransfer ||
		txType == types.TxTypeFeeDelegatedSmartContractExecution {
		hash, err := sess.account.SendTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
		return sess.WaitMinedWithHash(ctx, hash)
	}
	return nil, errors.New("invalid transaction type")
}

func (sess Session) MakeTransaction(opts *ablbind.TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	txType, txValues, err := opts.MakeTransactionData(sess, contract, input)
	if err != nil {
		return nil, err
	}

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
