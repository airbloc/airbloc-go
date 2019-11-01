package bind

import (
	"crypto/ecdsa"
	"errors"

	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
)

func TxRawKeySigner(key *ecdsa.PrivateKey) bind.SignerFn {
	return func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if addr != crypto.PubkeyToAddress(key.PublicKey) {
			return nil, errors.New("unauthorized sign request")
		}

		h := signer.Hash(tx)
		sig, err := types.NewTxSignatureWithValues(signer, h, key)
		if err != nil {
			return nil, err
		}

		tx.SetSignature(types.TxSignatures{sig})
		return tx, nil
	}
}

func DelegateTxRawKeySigner(key *ecdsa.PrivateKey) bind.SignerFn {
	return func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if addr != crypto.PubkeyToAddress(key.PublicKey) {
			return nil, errors.New("unauthorized sign request")
		}

		h, err := signer.HashFeePayer(tx)
		if err != nil {
			return nil, err
		}

		sig, err := types.NewTxSignatureWithValues(signer, h, key)
		if err != nil {
			return nil, err
		}

		if err := tx.SetFeePayerSignatures(types.TxSignatures{sig}); err != nil {
			return nil, err
		}
		return tx, nil
	}
}

func TxWalletSigner(account accounts.Account, wallet accounts.Wallet, passphrase *string) bind.SignerFn {
	return func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		hash := signer.Hash(tx)

		var (
			rawsig []byte
			err    error
		)
		if passphrase != nil {
			rawsig, err = wallet.SignHash(account, hash[:])
		} else {
			rawsig, err = wallet.SignHashWithPassphrase(account, *passphrase, hash[:])
		}
		if err != nil {
			return nil, err
		}

		sig := &types.TxSignature{}
		sig.R, sig.S, sig.V, err = signer.SignatureValues(rawsig)
		if err != nil {
			return nil, err
		}

		tx.SetSignature(types.TxSignatures{sig})
		return tx, nil
	}
}

func DelegateTxWalletSigner(account accounts.Account, wallet accounts.Wallet, passphrase *string) bind.SignerFn {
	return func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		hash, err := signer.HashFeePayer(tx)
		if err != nil {
			return nil, err
		}

		var rawsig []byte
		if passphrase != nil {
			rawsig, err = wallet.SignHash(account, hash[:])
		} else {
			rawsig, err = wallet.SignHashWithPassphrase(account, *passphrase, hash[:])
		}
		if err != nil {
			return nil, err
		}

		sig := &types.TxSignature{}
		sig.R, sig.S, sig.V, err = signer.SignatureValues(rawsig)
		if err != nil {
			return nil, err
		}

		if err := tx.SetFeePayerSignatures(types.TxSignatures{sig}); err != nil {
			return nil, err
		}
		return tx, nil
	}
}
