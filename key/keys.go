package key

import (
	"crypto/ecdsa"

	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	libp2pCrypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

// Key is an ECDSA keypair with SECP256K1 curve
// which is compatible with Ethereum (ECDSA), ECIES and Ed25519 (BigchainDB, by deriving a key).
type Key struct {
	*ecdsa.PrivateKey
	Ed25519Private  ed25519.PrivateKey
	ECIESPrivate    *ecies.PrivateKey
	EthereumAddress common.Address
}

func FromECDSA(key *ecdsa.PrivateKey) *Key {
	return &Key{
		PrivateKey:      key,
		Ed25519Private:  ed25519.NewKeyFromSeed(key.D.Bytes()),
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

// rawEd25519PrivKey returns raw 64-byte Ed25519 private key
func (key *Key) rawEd25519PrivKey() (privKey []byte) {
	privKey = make([]byte, ed25519.PrivateKeySize)
	copy(privKey, key.Ed25519Private)
	return
}

// rawEd25519PublicKey returns raw 32-byte Ed25519 public key
func (key *Key) rawEd25519Public() (publicKey []byte) {
	publicKey = make([]byte, ed25519.PublicKeySize)
	copy(publicKey, key.Ed25519Private[32:])
	return
}

// DeriveBigchainDBKeyPair returns an Ed25519 keypair
// that can be used for signing BigchainDB transactions on Airbloc.
func (key *Key) DeriveBigchainDBKeyPair() *txn.KeyPair {
	return &txn.KeyPair{
		PrivateKey: key.Ed25519Private,
		PublicKey:  key.rawEd25519Public(),
	}
}

// DeriveLibp2pKeyPair returns an Ed25519 keypair for Libp2p identity
func (key *Key) DeriveLibp2pKeyPair() (libp2pCrypto.PrivKey, error) {
	return libp2pCrypto.UnmarshalEd25519PrivateKey(
		append(
			key.rawEd25519PrivKey(),
			key.rawEd25519Public()...,
		),
	)
}
