package account

import (
	"testing"

	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/types"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSession_Sign(t *testing.T) {
	accountId, _ := types.HexToID("deadbeefdeadbeef")
	walletAddress := common2.HexToAddress("0xdeadbeef397a8ea5464f8cc753645d42e14b79ea")
	password := "foobar1234"
	session := NewSession(accountId, walletAddress, password)

	msg := crypto.Keccak256Hash([]byte("foo"))
	sig, err := session.Sign(msg)
	assert.NoError(t, err, "Failed to sign")

	// test recovering public key
	identityHash := crypto.Keccak256Hash(walletAddress.Bytes())
	priv := key.DeriveFromPassword(identityHash, password)

	recoveredPub, err := crypto.Ecrecover(msg[:], sig)
	assert.NoError(t, err)
	assert.Equal(t, recoveredPub, crypto.FromECDSAPub(&priv.PublicKey))
}
