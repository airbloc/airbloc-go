package airbloc

import (
	"context"
	"crypto/ecdsa"
	"net/http"

	ablbind "github.com/airbloc/airbloc-go/bind"

	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/blockchain/types"
)

type Account struct {
	account  *ablbind.TransactOpts
	feePayer *feePayerClient
}

func (acc Account) isReadOnly() bool {
	return acc.account == nil
}

func (acc Account) isDelegated() bool {
	return acc.feePayer != nil
}

func (acc *Account) SetAccount(account *ablbind.TransactOpts) {
	acc.account = account
}

func (acc *Account) SetFeePayer(ctx context.Context, feePayerUrl string) error {
	return acc.feePayer.SetFeePayer(ctx, feePayerUrl)
}

func (acc Account) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return acc.feePayer.Transact(ctx, tx)
}

func NewKeyedAccount(key *ecdsa.PrivateKey) Account {
	return Account{account: ablbind.NewKeyedTransactor(key)}
}

func NewWalletAccount(account accounts.Account, wallet accounts.Wallet) Account {
	return Account{account: ablbind.NewWalletTransactor(account, wallet, nil)}
}

func newAccountWithFeePayer(ctx context.Context, acc Account, rawFeePayerUrl string) (Account, error) {
	client := &feePayerClient{client: http.DefaultClient}
	if err := client.SetFeePayer(ctx, rawFeePayerUrl); err != nil {
		return Account{}, err
	}
	acc.feePayer = client
	return acc, nil
}

func NewKeyedAccountWithFeePayer(ctx context.Context, key *ecdsa.PrivateKey, feePayerUrl string) (Account, error) {
	return newAccountWithFeePayer(ctx, NewKeyedAccount(key), feePayerUrl)
}

func NewWalletAccountWithFeePayer(ctx context.Context, account accounts.Account, wallet accounts.Wallet, feePayerUrl string) (Account, error) {
	return newAccountWithFeePayer(ctx, NewWalletAccount(account, wallet), feePayerUrl)
}
