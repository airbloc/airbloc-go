package cryptoconditions

import "crypto/sha256"

// FfPreimageSha256 implements the PREIMAGE-SHA-256 fulfillment.
type FfPreimageSha256 struct {
	Preimage []byte `asn1:"tag:0"`
}

// NewPreimageSha256 creates a new PREIMAGE-SHA-256 fulfillment.
func NewPreimageSha256(preimage []byte) *FfPreimageSha256 {
	return &FfPreimageSha256{
		Preimage: preimage,
	}
}

func (f FfPreimageSha256) ConditionType() ConditionType {
	return CTPreimageSha256
}

func (f FfPreimageSha256) Cost() int {
	return len(f.Preimage)
}

func (f FfPreimageSha256) fingerprintContents() []byte {
	return f.Preimage
}

func (f FfPreimageSha256) fingerprint() []byte {
	hash := sha256.Sum256(f.fingerprintContents())
	return hash[:]
}

func (f FfPreimageSha256) Condition() *Condition {
	return NewSimpleCondition(f.ConditionType(), f.fingerprint(), f.Cost())
}

func (f FfPreimageSha256) Encode() ([]byte, error) {
	return encodeFulfillment(f)
}

func (f FfPreimageSha256) Validate(condition *Condition, message []byte) error {
	if !matches(f, condition) {
		return fulfillmentDoesNotMatchConditionError
	}

	// For a preimage fulfillment, no additional check is required.
	return nil
}
