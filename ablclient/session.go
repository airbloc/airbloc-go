package ablclient

import (
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
)

type Session struct {
	session
}

type session struct {
	AccountId     types.ID
	WalletAddress common.Address
	Key           *key.Key
}

func NewSession(accountId types.ID, walletAddress common.Address, password string) *Session {
	identityHash := crypto.Keccak256Hash(walletAddress.Bytes())
	priv := key.DeriveFromPassword(identityHash, password)
	return &Session{
		session{
			AccountId:     accountId,
			WalletAddress: walletAddress,
			Key:           priv,
		},
	}
}

func (session *session) Sign(hash common.Hash) ([]byte, error) {
	return crypto.Sign(hash[:], session.Key.PrivateKey)
}
