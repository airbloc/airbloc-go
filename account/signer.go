package account

import (
	"crypto/ecdsa"

	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/crypto"
)

type MessageSigner func(message []byte) (signature []byte, err error)

func NewWalletMessageSigner(account accounts.Account, wallet accounts.Wallet, passphrase *string) MessageSigner {
	return func(message []byte) (signature []byte, err error) {
		if passphrase != nil {
			signature, err = wallet.SignHash(account, message)
		} else {
			signature, err = wallet.SignHashWithPassphrase(account, *passphrase, message)
		}
		return
	}
}

func NewKeyedMessageSigner(key *ecdsa.PrivateKey) MessageSigner {
	return func(message []byte) (signature []byte, err error) {
		signature, err = crypto.Sign(message, key)
		return
	}
}
