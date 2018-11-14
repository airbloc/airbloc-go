package cryptoconditions

import (
	"crypto/sha256"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

// ffEd25519Sha256Cost is the fixed cost value for ED25519-SHA-256 fulfillments.
const ffEd25519Sha256Cost = 131072

// FfEd25519Sha256 implements the ED25519-SHA-256 fulfillment.
type FfEd25519Sha256 struct {
	PublicKey []byte `asn1:"tag:0"`
	Signature []byte `asn1:"tag:1"`
}

// NewEd25519Sha256 creates a new ED25519-SHA-256 fulfillment.
func NewEd25519Sha256(pubkey []byte, signature []byte) (*FfEd25519Sha256, error) {
	if len(pubkey) != ed25519.PublicKeySize {
		return nil, errors.Errorf(
			"wrong pubkey size (%d)", len(pubkey))
	}
	if len(signature) != ed25519.SignatureSize && len(signature) != 0 {
		return nil, errors.Errorf(
			"wrong signature size (%d)", len(signature))
	}
	return &FfEd25519Sha256{
		PublicKey: pubkey,
		Signature: signature,
	}, nil
}

// Ed25519PublicKey returns the Ed25519 public key.
func (f FfEd25519Sha256) Ed25519PublicKey() ed25519.PublicKey {
	return ed25519.PublicKey(f.PublicKey)
}

func (f FfEd25519Sha256) ConditionType() ConditionType {
	return CTEd25519Sha256
}

func (f FfEd25519Sha256) Cost() int {
	return ffEd25519Sha256Cost
}

func (f FfEd25519Sha256) fingerprintContents() []byte {
	content := struct {
		PubKey []byte `asn1:"tag:0"`
	}{
		PubKey: f.PublicKey,
	}

	encoded, err := ASN1Context.Encode(content)
	if err != nil {
		panic(err) //TODO check when this can happen
	}

	return encoded
}

func (f FfEd25519Sha256) fingerprint() []byte {
	hash := sha256.Sum256(f.fingerprintContents())
	return hash[:]
}

func (f FfEd25519Sha256) Condition() *Condition {
	return NewSimpleCondition(f.ConditionType(), f.fingerprint(), f.Cost())
}

func (f FfEd25519Sha256) Encode() ([]byte, error) {
	return encodeFulfillment(f)
}

func (f FfEd25519Sha256) Validate(condition *Condition, message []byte) error {
	if !matches(f, condition) {
		return fulfillmentDoesNotMatchConditionError
	}

	if ed25519.Verify(f.Ed25519PublicKey(), message, f.Signature) {
		return nil
	} else {
		return fmt.Errorf("Unable to Validate Ed25519Sha256 fulfillment: "+
			"signature verification failed for message %x", message)
	}
}
