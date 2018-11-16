package cryptoconditions

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"math/big"

	"github.com/pkg/errors"
)

const (
	// The RSA parameters we have to use.
	ffRsaSha256MinimumModulusLength = 128
	ffRsaSha256MaximumModulusLength = 512
	ffRsaSha256PublicExponent       = 65537
)

var ffRsaSha256PssOpts = &rsa.PSSOptions{
	SaltLength: 32,
	Hash:       crypto.SHA256,
}

// NewFfRsaSha256 implements the RSA-SHA-256 fulfillment.
type FfRsaSha256 struct {
	Modulus   []byte `asn1:"tag:0"`
	Signature []byte `asn1:"tag:1"`
}

// RsaSha256 creates a new RSA-SHA-256 fulfillment.
func NewRsaSha256(modulus []byte, signature []byte) (*FfRsaSha256, error) {
	if len(modulus) < ffRsaSha256MinimumModulusLength {
		return nil, errors.New("modulus is too small.")
	}
	if len(modulus) > ffRsaSha256MaximumModulusLength {
		return nil, errors.New("modulus is too large.")
	}

	return &FfRsaSha256{
		Modulus:   modulus,
		Signature: signature,
	}, nil
}

// PublicKey returns the RSA public key.
func (f FfRsaSha256) PublicKey() *rsa.PublicKey {
	return &rsa.PublicKey{
		N: new(big.Int).SetBytes(f.Modulus),
		E: ffRsaSha256PublicExponent,
	}
}

func (f FfRsaSha256) ConditionType() ConditionType {
	return CTRsaSha256
}

func (f FfRsaSha256) Cost() int {
	return len(f.Modulus) * len(f.Modulus)
}

func (f FfRsaSha256) fingerprintContents() []byte {
	content := struct {
		Modulus []byte `asn1:"tag:0"`
	}{
		Modulus: f.Modulus,
	}

	encoded, err := ASN1Context.Encode(content)
	if err != nil {
		panic(err) //TODO check when this can happen
	}

	return encoded
}

func (f FfRsaSha256) fingerprint() []byte {
	hash := sha256.Sum256(f.fingerprintContents())
	return hash[:]
}

func (f FfRsaSha256) Condition() *Condition {
	return NewSimpleCondition(f.ConditionType(), f.fingerprint(), f.Cost())
}

func (f FfRsaSha256) Encode() ([]byte, error) {
	return encodeFulfillment(f)
}

func (f FfRsaSha256) Validate(condition *Condition, message []byte) error {
	if !matches(f, condition) {
		return fulfillmentDoesNotMatchConditionError
	}

	hashed := sha256.Sum256(message)
	err := rsa.VerifyPSS(
		f.PublicKey(), crypto.SHA256, hashed[:], f.Signature, ffRsaSha256PssOpts)

	return errors.Wrapf(err,
		"failed to verify RSA signature of binary message \"%x\" (hex)", message)
}
