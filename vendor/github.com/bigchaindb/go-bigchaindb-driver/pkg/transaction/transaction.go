package transaction

import (
	"encoding/json"

	"encoding/hex"

	"bytes"

	"strings"

	"net/url"

	"github.com/go-interledger/cryptoconditions"
	"github.com/kalaspuffar/base64url"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/sha3"
)

func New(
	operation string,
	asset Asset,
	metadata Metadata,
	inputs []Input,
	outputs []Output,
) (*Transaction, error) {

	if !(operation == "CREATE" || operation == "TRANSFER") {
		return &Transaction{}, errors.New("Not a valid operation - expecting 'CREATE' or 'TRANSFER'")
	}

	return &Transaction{
		Asset:     asset,
		ID:        nil,
		Inputs:    inputs,
		Metadata:  metadata,
		Operation: operation,
		Outputs:   outputs,
		Version:   "2.0",
	}, nil
}

func NewCreateTransaction(
	asset Asset,
	metadata Metadata,
	outputs []Output,
	issuers []ed25519.PublicKey,
) (*Transaction, error) {
	// Create list of unfulfilled unfulfilledInputs
	unfulfilledInputs := createInputs([]ed25519.PublicKey{issuers[0]})
	return New("CREATE", asset, metadata, unfulfilledInputs, outputs)
}

func createInputs(publicKeys []ed25519.PublicKey) []Input {
	var inputs []Input
	for _, pubKey := range publicKeys {
		input := Input{
			// New input is unfulfilled - fulfillment string is ctnull
			Fulfillment:  nil,
			Fulfills:     nil,
			OwnersBefore: []string{base58.Encode(pubKey)},
		}
		inputs = append(inputs, input)
	}
	return inputs
}

func createInputsFromUnspentTransactions(unspentTransactions []Transaction) ([]Input, error) {
	var inputs []Input

	var unspentOutputs []Output
	for _, ut := range unspentTransactions {
		unspentOutputs = append(unspentOutputs, ut.Outputs...)
	}

	for _, uo := range unspentOutputs {
		input := Input{
			OwnersBefore: uo.PublicKeys,
		}
		inputs = append(inputs, input)

	}
	return inputs, nil
}

// TODO clarify starting point transfer txn: Outputlocations or unspent transactions?
func NewTransferTransaction(
	unspentTransactions []Transaction,
	outputs []Output,
	metadata Metadata,
) (*Transaction, error) {
	inputs, err := createInputsFromUnspentTransactions(unspentTransactions)
	if err != nil {
		return nil, err
	}

	var asset Asset
	// FIXME make sure all unspent txns point to same asset
	if unspentTransactions[0].Operation == "CREATE" {
		asset.ID = unspentTransactions[0].ID
	} else {
		asset = unspentTransactions[0].Asset
	}

	return New("TRANSFER", asset, metadata, inputs, outputs)
}

func NewOutput(condition cryptoconditions.Condition, amount string) (Output, error) {
	if amount == "" {
		amount = "1"
	}

	if condition.Type() == cryptoconditions.CTThresholdSha256 {
		return Output{}, errors.New("No support for treshold-sha-256 yet")
	}

	return Output{
		Condition: Condition{
			Uri: generateURI(condition.Type().String(), base64url.Encode(condition.Fingerprint())),
			Details: ConditionDetail{
				PublicKey: base58.Encode(condition.Fingerprint()),
				Type:      strings.ToLower(condition.Type().String()),
			},
		},
		Amount:     amount,
		PublicKeys: []string{base58.Encode(condition.Fingerprint())},
	}, nil
}

// generateURI generates a URI for the given condition.
// TODO - BigchainDB regexp to validate URI does not match with go-interledger cryptoconditions
// because it forces order in cost and fpt query params
// ^ni:///sha-256;([a-zA-Z0-9_-]{0,86})[?](fpt=(ed25519|threshold)-sha-256(&)?|cost=[0-9]+(&)?|subtypes=ed25519-sha-256(&)?){2,3}$
// FIXME Example uri with base64url encoding that cant be parsed:
// ni:///sha-256;0FOQ3QM4qNHvr-Yf0F37WNU8TUeqE-AykSBRhsnigNk?cost=131072&fpt=ed25519-sha-256
func generateURI(cType, encodedFingerprint string) string {
	params := make(url.Values)
	// FIXME hardcoded costs
	params.Set("cost", "131072")
	params.Set("fpt", strings.ToLower(cType))

	uri := url.URL{
		Scheme:   "ni",
		Path:     "/sha-256;" + encodedFingerprint,
		RawQuery: params.Encode(),
	}

	return uri.String()
}

/*
	The ID of a transaction is the SHA3-256 hash of the transaction.
*/
func (t *Transaction) createID() (string, error) {

	// Strip ID of txn
	tn := &Transaction{
		ID:        nil,
		Version:   t.Version,
		Inputs:    t.Inputs,
		Outputs:   t.Outputs,
		Operation: t.Operation,
		Asset:     t.Asset,
		Metadata:  t.Metadata,
	}
	// Serialize transaction - encoding/json follows RFC7159 and BDB marshalling
	dbytes, err := tn.JSON()
	if err != nil {
		return "", err
	}

	// Return hash of serialized txn object
	h := sha3.Sum256(dbytes)
	return hex.EncodeToString(h[:]), nil
}

func (t *Transaction) String() (string, error) {

	dbytes, err := t.JSON()
	if err != nil {
		return "", err
	}
	dstr := string(dbytes[:])
	return dstr, nil
}

func (t *Transaction) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(t)
	if err != nil {
		return nil, err
	}
	dbytes := buffer.Bytes()

	buffer1 := &bytes.Buffer{}
	err = json.Compact(buffer1, dbytes)
	if err != nil {
		return nil, err
	}
	dbytes = buffer1.Bytes()

	return dbytes, nil
}
