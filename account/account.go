package account

import (
	"context"
	"crypto/ecdsa"

	ablbind "github.com/airbloc/airbloc-go/bind"

	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
)

type Account struct {
	account  *ablbind.TransactOpts
	feePayer *FeePayer
}

func (acc Account) IsReadOnly() bool {
	return acc.account == nil
}

func (acc Account) IsDelegated() bool {
	return acc.feePayer != nil
}

func (acc Account) TxOpts() *ablbind.TransactOpts {
	return acc.account
}

func (acc Account) SendTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	return acc.feePayer.Transact(ctx, tx)
}

func NewKeyedAccount(key *ecdsa.PrivateKey) Account {
	return Account{account: ablbind.NewKeyedTransactor(key)}
}

func NewWalletAccount(account accounts.Account, wallet accounts.Wallet) Account {
	return Account{account: ablbind.NewWalletTransactor(account, wallet, nil)}
}

func newAccountWithFeePayer(acc Account, rawFeePayerUrl string) (Account, error) {
	feePayer, err := NewFeePayer(nil, rawFeePayerUrl)
	if err != nil {
		return Account{}, err
	}

	acc.feePayer = feePayer
	return acc, nil
}

func NewKeyedAccountWithFeePayer(key *ecdsa.PrivateKey, feePayerUrl string) (Account, error) {
	return newAccountWithFeePayer(NewKeyedAccount(key), feePayerUrl)
}

func NewWalletAccountWithFeePayer(account accounts.Account, wallet accounts.Wallet, feePayerUrl string) (Account, error) {
	return newAccountWithFeePayer(NewWalletAccount(account, wallet), feePayerUrl)
}
