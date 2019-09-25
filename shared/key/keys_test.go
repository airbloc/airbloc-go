package key

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/klaytn/klaytn/crypto"
	"github.com/stretchr/testify/assert"
)

func TestDeriveFromPassword(t *testing.T) {
	identity := crypto.Keccak256Hash([]byte("Identity!"))
	key := DeriveFromPassword(identity, "alslalsl")

	fmt.Printf("Key: %s\n", hex.EncodeToString(key.PrivateKey.D.Bytes()))
	fmt.Printf("PubKey: %s\n", hex.EncodeToString(crypto.CompressPubkey(&key.PublicKey)))

	// sign test
	hash := crypto.Keccak256Hash([]byte("Message"))

	sig, err := crypto.Sign(hash[:], key.PrivateKey)
	assert.NoError(t, err, "Failed to sign using the derived key.")

	recoveredPub, err := crypto.SigToPub(hash[:], sig)
	assert.NoError(t, err, "Failed to recover from pubkey.")

	fmt.Printf("Recovered Pub: %s\n", hex.EncodeToString(crypto.CompressPubkey(recoveredPub)))

	signatureIsCorrect := crypto.VerifySignature(crypto.CompressPubkey(&key.PublicKey), hash[:], sig[:64])
	assert.True(t, signatureIsCorrect, "Signature is incorrect!")
}

func TestGenerate(t *testing.T) {
	_, err := Generate()
	require.NoError(t, err)
}
