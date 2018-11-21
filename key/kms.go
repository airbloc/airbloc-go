package key

import (
	"crypto/rand"

	"github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/airbloc/airbloc-go/database/localdb"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

// Manager is local key-manager service (KMS) containing the account's private key
// and re-encryption keys that owned by the account.
type KeyManager struct {
	OwnerKey      *Key
	localDatabase localdb.Database
}

func NewKeyManager(ownerKey *Key, localDatabase localdb.Database) Manager {
	log.Debug("Private key loaded", "address", ownerKey.EthereumAddress.Hex())
	return &KeyManager{
		OwnerKey:      ownerKey,
		localDatabase: localDatabase,
	}
}

// DecryptExternalData looks for re-encryption key, then tries to decrypt.
func (kms *KeyManager) DecryptExternalData(data *common.EncryptedData) (*common.Data, error) {
	// TODO: Implement key.Manager.DecryptExternalData
	return nil, nil
}

func (kms *KeyManager) Encrypt(payload string) ([]byte, error) {
	publicKey := kms.OwnerKey.ECIESPrivate.PublicKey
	return ecies.Encrypt(rand.Reader, &publicKey, []byte(payload), nil, nil)
}

func (kms *KeyManager) Decrypt(encryptedPayload []byte) (string, error) {
	privateKey := kms.OwnerKey.ECIESPrivate
	payload, err := privateKey.Decrypt(encryptedPayload, nil, nil)
	return string(payload), err
}

func (kms *KeyManager) SignEthTx(tx *types.Transaction) (*types.Transaction, error) {
	return types.SignTx(tx, types.EIP155Signer{}, kms.OwnerKey.PrivateKey)
}

func (kms *KeyManager) SignBDBTx(tx *txn.Transaction) error {
	return tx.Sign([]*txn.KeyPair{kms.OwnerKey.DeriveBigchainDBKeyPair()})
}
