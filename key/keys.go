package key

import (
	"crypto/ecdsa"

	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

// Key is an ECDSA keypair with SECP256K1 curve
// which is compatible with Ethereum (ECDSA), ECIES and Ed25519 (BigchainDB, by deriving a key).
type Key struct {
	*ecdsa.PrivateKey
	ECIESPrivate    *ecies.PrivateKey
	EthereumAddress common.Address
}

func FromECDSA(key *ecdsa.PrivateKey) *Key {
	return &Key{
		PrivateKey:      key,
		ECIESPrivate:    ecies.ImportECDSA(key),
		EthereumAddress: crypto.PubkeyToAddress(key.PublicKey),
	}
}

// Generate creates new random ECDSA Key.
func Generate() (*Key, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate ECDSA key.")
	}
	return FromECDSA(privateKey), nil
}

// Load loads a private key from the given file.
// the file must contain a 32-byte hex-encoded ECDSA private key with SECP256k1 curve.
func Load(path string) (*Key, error) {
	privateKey, err := crypto.LoadECDSA(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load ECDSA key.")
	}
	return FromECDSA(privateKey), nil
}

// Save saves a private key to the given file.
func (key *Key) Save(path string) error {
	if err := crypto.SaveECDSA(path, key.PrivateKey); err != nil {
		return errors.Wrap(err, "failed to save ECDSA key to given path.")
	}
	return nil
}

// DeriveEd25519KeyPair returns an Ed25519 keypair
// that can be used for signing BigchainDB transactions on Airbloc.
func (key *Key) DeriveEd25519KeyPair() *txn.KeyPair {
	privateKey := ed25519.NewKeyFromSeed(key.D.Bytes())

	publicKey := make([]byte, ed25519.PublicKeySize)
	copy(publicKey, privateKey[32:])

	return &txn.KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}
