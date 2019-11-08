package airbloc

import (
	"context"

	ablbind "github.com/airbloc/airbloc-go/bind"

	"github.com/klaytn/klaytn/blockchain/types"
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
	return ablbind.MergeTxOpts(ctx, sess.account.account, opts...)
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
