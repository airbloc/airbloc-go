package key

import (
	"github.com/airbloc/airbloc-go/pkg/types"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
)

type Manager interface {
	NodeKey() *Key
	EncryptData(*types.Data) (*types.EncryptedData, error)
	DecryptData(*types.EncryptedData) (*types.Data, error)
	Encrypt(string) ([]byte, error)
	Decrypt([]byte) (string, error)
	SignEthTx(*klayTypes.Transaction) (*klayTypes.Transaction, error)
}
