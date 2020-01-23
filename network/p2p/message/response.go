package message

import (
	"github.com/airbloc/airbloc-go/account"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/crypto"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Response interface {
	Message
	Signature() hexutil.Bytes
	SetSignature(sign hexutil.Bytes)
}

func SignResponseMessage(resp Response, acc account.Account) error {
	if resp == nil {
		return errors.New("nil response")
	}
	if resp.ID() == (uuid.UUID{}) {
		return errors.New("empty uuid")
	}
	var id [32]byte
	copy(id[:], resp.ID().Bytes())
	sign, err := acc.SignMessage(id[:])
	if err != nil {
		return errors.Wrap(err, "failed to sign response message with account")
	}
	resp.SetSignature(sign)
	return nil
}

func VerifyResponseMessage(resp Response, sender common.Address) (bool, error) {
	if resp == nil {
		return false, errors.New("nil response")
	}
	if resp.ID() == (uuid.UUID{}) {
		return false, errors.New("empty uuid")
	}
	var id [32]byte
	copy(id[:], resp.ID().Bytes())
	pubkeyBytes, err := crypto.Ecrecover(id[:], resp.Signature())
	if err != nil {
		return false, errors.Wrap(err, "failed to recover signer's public key from signature")
	}

	pubkey, err := crypto.UnmarshalPubkey(pubkeyBytes)
	if err != nil {
		return false, errors.Wrap(err, "failed to unmarshal public key")
	}

	return crypto.PubkeyToAddress(*pubkey) == sender, nil
}
