package account

import (
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Session struct {
	AccountId     types.ID
	WalletAddress ethCommon.Address
	Key           *key.Key
}

func NewSession(accountId types.ID, walletAddress ethCommon.Address, password string) *Session {
	identityHash := crypto.Keccak256Hash(walletAddress.Bytes())
	priv := key.DeriveFromPassword(identityHash, password)
	return &Session{
		AccountId:     accountId,
		WalletAddress: walletAddress,
		Key:           priv,
	}
}

func (session *Session) Sign(hash ethCommon.Hash) ([]byte, error) {
	return crypto.Sign(hash[:], session.Key.PrivateKey)
}
