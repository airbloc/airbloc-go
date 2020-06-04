package key

import (
	"testing"

	"github.com/klaytn/klaytn/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeriveFromPassword(t *testing.T) {
	identity := crypto.Keccak256Hash([]byte("Identity!"))
	key := DeriveFromPassword(identity, "alslalsl")

	// sign test
	hash := crypto.Keccak256Hash([]byte("Message"))

	sig, err := crypto.Sign(hash[:], key.PrivateKey)
	assert.NoError(t, err, "Failed to sign using the derived key.")

	recoveredPub, err := crypto.SigToPub(hash[:], sig)
	assert.NoError(t, err, "Failed to recover from pubkey.")
	assert.Equal(t, key.PublicKey, *recoveredPub)

	signatureIsCorrect := crypto.VerifySignature(crypto.CompressPubkey(&key.PublicKey), hash[:], sig[:64])
	assert.True(t, signatureIsCorrect, "Signature is incorrect!")
}

func TestGenerate(t *testing.T) {
	_, err := Generate()
	require.NoError(t, err)
}
