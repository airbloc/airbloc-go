package key

import (
	"github.com/airbloc/airbloc-go/common"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/core/types"
)

type Manager interface {
	NodeKey() *Key
	EncryptData(*common.Data) (*common.EncryptedData, error)
	DecryptData(*common.EncryptedData) (*common.Data, error)
	Encrypt(string) ([]byte, error)
	Decrypt([]byte) (string, error)
	SignEthTx(*types.Transaction) (*types.Transaction, error)
	SignBDBTx(*txn.Transaction) error
}
