package key

import (
	"crypto/ecdsa"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
	"math/big"

	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

var (
	passwordSalt = []byte("AirblocPassword")
	one = big.NewInt(1)
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

// DeriveFromPassword uses PBKDF2 with
func DeriveFromPassword(identity common.Hash, password string) (*Key) {
	passwordHash := sha3.Sum256([]byte(password))

	// make PBKDF2 input material
	// with sha3_256(identity hash) + sha3_256(password_hash)
	keybase := make([]byte, 64)
	copy(keybase[:32], identity[:])
	copy(keybase[32:], passwordHash[:])

	curveParams := crypto.S256().Params()
	material := pbkdf2.Key(keybase, passwordSalt, 1024, curveParams.BitSize/8+8, sha3.New256)

	// make sure that k <= N
	k := new(big.Int).SetBytes(material)
	n := new(big.Int).Sub(curveParams.N, one)
	k.Mod(k, n)
	k.Add(k, one)

	private := new(ecdsa.PrivateKey)
	private.PublicKey.Curve = crypto.S256()
	private.D = k
	private.PublicKey.X, private.PublicKey.Y = crypto.S256().ScalarBaseMult(k.Bytes())
	return FromECDSA(private)
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
