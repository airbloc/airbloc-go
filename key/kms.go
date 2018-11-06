package key

import (
	"crypto/rand"

	"github.com/airbloc/airbloc-go/database/localdb"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

// Manager is local key-manager service (KMS) containing the account's private key
// and re-encryption keys that owned by the account.
type Manager struct {
	OwnerKey      *Key
	localDatabase localdb.Database
}

func NewManager(ownerKey *Key, localDatabase localdb.Database) *Manager {
	return &Manager{
		OwnerKey:      ownerKey,
		localDatabase: localDatabase,
	}
}

func (kms *Manager) Encrypt(payload string) ([]byte, error) {
	publicKey := kms.OwnerKey.ECIESPrivate.PublicKey
	return ecies.Encrypt(rand.Reader, &publicKey, []byte(payload), nil, nil)
}

func (kms *Manager) Decrypt(encryptedPayload []byte) (string, error) {
	privateKey := kms.OwnerKey.ECIESPrivate
	payload, err := privateKey.Decrypt(encryptedPayload, nil, nil)
	return string(payload), err
}

func (kms *Manager) SignEthTx(tx *types.Transaction) (*types.Transaction, error) {
	return types.SignTx(tx, types.EIP155Signer{}, kms.OwnerKey.PrivateKey)
}

func (kms *Manager) SignBDBTx(tx *txn.Transaction) error {
	return tx.Sign([]*txn.KeyPair{kms.OwnerKey.DeriveBigchainDBKeyPair()})
}
