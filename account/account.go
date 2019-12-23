package account

import (
	"context"
	"crypto/ecdsa"

	ablbind "github.com/airbloc/airbloc-go/bind"

	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/pkg/errors"
)

type Account struct {
	pubKey   *ecdsa.PublicKey
	signer   MessageSigner
	account  *ablbind.TransactOpts
	feePayer *FeePayer
}

func (acc Account) Address() common.Address {
	return acc.account.From
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

func (acc Account) PublicKey() ecdsa.PublicKey {
	return *acc.pubKey
}

func (acc Account) SignMessage(message []byte) ([]byte, error) {
	return acc.signer(message)
}

func (acc Account) SendTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	return acc.feePayer.Transact(ctx, tx)
}

func NewKeyedAccount(key *ecdsa.PrivateKey) Account {
	return Account{
		pubKey:  &key.PublicKey,
		signer:  NewKeyedMessageSigner(key),
		account: ablbind.NewKeyedTransactor(key),
	}
}

func NewWalletAccount(account accounts.Account, wallet accounts.Wallet) (Account, error) {
	signer := NewWalletMessageSigner(account, wallet, nil)

	// vulfpeck - 1612
	hash := []byte{1, 6, 1, 2, '*'}

	signature, err := signer(hash)
	if err != nil {
		return Account{}, errors.Wrap(errors.Wrap(err, "failed to get public key"), "sign message")
	}

	pubKeyBytes, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		return Account{}, errors.Wrap(errors.Wrap(err, "failed to get public key"), "recover pubkey from signature")
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return Account{}, errors.Wrap(errors.Wrap(err, "failed to get public key"), "unmarshal pubkey")
	}

	return Account{
		pubKey:  pubKey,
		signer:  NewWalletMessageSigner(account, wallet, nil),
		account: ablbind.NewWalletTransactor(account, wallet, nil),
	}, nil
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
	acc, err := NewWalletAccount(account, wallet)
	if err != nil {
		return Account{}, err
	}
	return newAccountWithFeePayer(acc, feePayerUrl)
}
