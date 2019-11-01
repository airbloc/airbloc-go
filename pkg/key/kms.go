package key

import (
	"crypto/rand"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/logger"

	"github.com/airbloc/airbloc-go/shared/database/localdb"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/crypto/ecies"
)

// Manager is local key-manager service (KMS) containing the account's private key
// and re-encryption keys that owned by the account.
type manager struct {
	ownerKey      *Key
	localDatabase localdb.Database
	log           *logger.Logger
}

func NewKeyManager(ownerKey *Key, localDatabase localdb.Database) Manager {
	kms := &manager{
		ownerKey:      ownerKey,
		localDatabase: localDatabase,
		log:           logger.New("kms"),
	}
	kms.log.Info("Node key loaded.", logger.Attrs{"address": ownerKey.EthereumAddress.Hex()})
	return kms
}

// NodeKey returns the owner key.
func (kms *manager) NodeKey() *Key {
	return kms.ownerKey
}

func (kms *manager) EncryptData(data *types.Data) (*types.EncryptedData, error) {
	encryptedData, err := kms.Encrypt(data.Payload)
	if err != nil {
		return nil, err
	}
	return &types.EncryptedData{
		Payload:     encryptedData,
		UserId:      data.UserId,
		RowId:       data.RowId,
		CollectedAt: data.CollectedAt,
	}, nil
}

func (kms *manager) DecryptData(encryptedData *types.EncryptedData) (*types.Data, error) {
	data, err := kms.Decrypt(encryptedData.Payload)
	if err != nil {
		return nil, err
	}
	return &types.Data{
		Payload: data,
		UserId:  encryptedData.UserId,
		RowId:   encryptedData.RowId,
	}, nil
}

func (kms *manager) Encrypt(payload string) ([]byte, error) {
	publicKey := kms.ownerKey.ECIESPrivate.PublicKey
	return ecies.Encrypt(rand.Reader, &publicKey, []byte(payload), nil, nil)
}

func (kms *manager) Decrypt(encryptedPayload []byte) (string, error) {
	privateKey := kms.ownerKey.ECIESPrivate
	payload, err := privateKey.Decrypt(encryptedPayload, nil, nil)
	return string(payload), err
}

func (kms *manager) SignEthTx(tx *klayTypes.Transaction) (*klayTypes.Transaction, error) {
	return klayTypes.SignTx(tx, klayTypes.EIP155Signer{}, kms.ownerKey.PrivateKey)
}
