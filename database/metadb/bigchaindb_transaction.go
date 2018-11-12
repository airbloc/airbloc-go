package metadb

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-interledger/cryptoconditions"
	"github.com/kalaspuffar/base64url"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	Ed25519Magic = []byte{0x30, 0x22, 0x80, 0x20}
	Ed25519Cost  = 131072
)

type IntermediateTxn struct {
	Transaction txn.Transaction `json:"transaction"`
	Signatures  struct {
		PublicKey string `json:"publicKey"`
		Signature string `json:"signature"`
	} `json:"signatures"`
}

func (db *bigchainDB) prepareTx(tx *txn.Transaction) (*IntermediateTxn, error) {
	signature, err := db.preFulfill(tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign transaction")
	}

	// prepare an intermediate transaction
	inTxn := IntermediateTxn{Transaction: *tx}
	inTxn.Signatures.Signature = base64.StdEncoding.EncodeToString(signature)
	inTxn.Signatures.PublicKey = base64.StdEncoding.EncodeToString(db.key.PublicKey)
	return &inTxn, nil
}

// Sign signs transaction according to BEP-13,
// but DOES NOT fulfill the inputs.
func (db *bigchainDB) preFulfill(t *txn.Transaction) ([]byte, error) {
	// Set transaction ID to ctnull value
	t.ID = nil

	signedTx := *t
	input := signedTx.Inputs[0]
	var serializedTxn strings.Builder
	s, err := t.String()
	if err != nil {
		return nil, err
	}
	serializedTxn.WriteString(s)

	// If fulfills is not empty add to make unique serialization Txn
	if input.Fulfills != nil {
		serializedTxn.WriteString(input.Fulfills.TransactionID)
		serializedTxn.WriteString(string(input.Fulfills.OutputIndex))
	}

	log.Debug("Signing Transaction", "payload", serializedTxn.String())
	bytesToSign := sha3.Sum256([]byte(serializedTxn.String()))

	// rand reader is ignored within Sign method; crypto.Hash(0) is sanity check to
	// make sure bytes_to_sign is not hashed already - ed25519.PrivateKey cannot sign hashed msg
	return db.key.PrivateKey.Sign(rand.Reader, bytesToSign[:], crypto.Hash(0))
}

func (db *bigchainDB) sendIntermediateTx(inTxn *IntermediateTxn) error {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(inTxn)

	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", "http://localhost:8984/transactions", buffer)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := db.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	if response.StatusCode != 200 {
		log.Error("BigchainDB request failed!", "status", response.StatusCode, "response", string(body))
		return errors.Errorf("server returned %s", response.Status)
	}
	return err
}

func (db *bigchainDB) sendTx(tx *txn.Transaction, mode Mode) error {
	switch mode {
	case BigchainTxModeDefault:
		return db.bdb.PostTransaction(tx)
	case BigchainTxModeCommit:
		return db.bdb.PostTransactionCommit(tx)
	case BigchainTxModeSync:
		return db.bdb.PostTransactionSync(tx)
	default:
		return errors.New("invalid mode (0 = default, 1 = commit, 2 = sync)")
	}
}

func (db *bigchainDB) newCreateTransaction(asset txn.Asset, metadata txn.Metadata) (*txn.Transaction, error) {
	// generate condition
	out, err := newOutput(db.key.PublicKey, BigchainDBAmount)
	if err != nil {
		return nil, err
	}

	// create transaction
	return txn.NewCreateTransaction(
		asset,
		metadata,
		[]txn.Output{out},
		[]ed25519.PublicKey{db.key.PublicKey},
	)
}

func (db *bigchainDB) newTransferTransaction(txid string, to ed25519.PublicKey, metadata txn.Metadata) (*txn.Transaction, error) {
	tx, err := db.bdb.GetTransaction(txid)
	if err != nil {
		return nil, err
	}

	cond := newEd25519Condition(to)
	out, err := txn.NewOutput(*cond, BigchainDBAmount)
	if err != nil {
		return nil, err
	}

	return txn.NewTransferTransaction(
		[]txn.Transaction{tx},
		[]txn.Output{out},
		metadata,
	)
}

// newOutput creates simple Ed25519-SHA256 condition for given public key.
func newOutput(pubKey ed25519.PublicKey, amount string) (txn.Output, error) {
	if amount == "" {
		amount = "1"
	}

	condition := newEd25519Condition(pubKey)
	return txn.Output{
		Condition: txn.Condition{
			Uri: generateURI(condition),
			Details: txn.ConditionDetail{
				PublicKey: base58.Encode(pubKey),
				Type:      strings.ToLower(condition.Type().String()),
			},
		},
		Amount:     amount,
		PublicKeys: []string{base58.Encode(pubKey)},
	}, nil
}

func generateURI(condition *cryptoconditions.Condition) string {
	return fmt.Sprintf("ni:///sha-256;%s?fpt=%s&cost=%d",
		base64url.Encode(condition.Fingerprint()),
		strings.ToLower(condition.Type().String()),
		condition.Cost())
}

func newEd25519Condition(pubKey ed25519.PublicKey) *cryptoconditions.Condition {
	var ed25519Fingerprint [36]byte
	copy(ed25519Fingerprint[:4], Ed25519Magic)
	copy(ed25519Fingerprint[4:], pubKey)

	fingerprint := sha256.Sum256(ed25519Fingerprint[:])
	return cryptoconditions.NewSimpleCondition(cryptoconditions.CTEd25519Sha256, fingerprint[:], Ed25519Cost)
}
