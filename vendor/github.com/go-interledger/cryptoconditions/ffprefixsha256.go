package cryptoconditions

import (
	"bytes"
	"crypto/sha256"

	"github.com/pkg/errors"
)

// FfPrefixSha256 implements the PREFIX-SHA-256 fulfillment.
type FfPrefixSha256 struct {
	Prefix           []byte `asn1:"tag:0"`
	MaxMessageLength uint32 `asn1:"tag:1"`

	// Only have either a sub-fulfillment or a sub-condition.
	SubFulfillment Fulfillment `asn1:"tag:2,explicit,choice:fulfillment"`
	subCondition   *Condition  `asn1:"-"`
}

// NewPrefixSha256 creates a new PREFIX-SHA-256 fulfillment.
func NewPrefixSha256(prefix []byte, maxMessageLength uint32, subFf Fulfillment) *FfPrefixSha256 {
	return &FfPrefixSha256{
		Prefix:           prefix,
		MaxMessageLength: maxMessageLength,
		SubFulfillment:   subFf,
	}
}

// PrefixSha256Unfulfilled creates an unfulfilled PREFIX-SHA-256 fulfillment.
func NewPrefixSha256Unfulfilled(prefix []byte, maxMessageLength uint32, subCondition *Condition) *FfPrefixSha256 {
	return &FfPrefixSha256{
		Prefix:           prefix,
		MaxMessageLength: maxMessageLength,
		subCondition:     subCondition,
	}
}

func (f FfPrefixSha256) ConditionType() ConditionType {
	return CTPrefixSha256
}

// SubCondition returns the sub-condition of this fulfillment.
func (f FfPrefixSha256) SubCondition() *Condition {
	if f.IsFulfilled() {
		return f.SubFulfillment.Condition()
	} else {
		return f.subCondition
	}
}

// IsFulfilled returns true if this fulfillment is fulfilled,
// i.e. when it contains a sub-fulfillment.
// If false, it only contains a sub-condition.
func (f FfPrefixSha256) IsFulfilled() bool {
	return f.SubFulfillment != nil
}

func (f FfPrefixSha256) Cost() int {
	return len(f.Prefix) +
		int(f.MaxMessageLength) +
		f.SubCondition().Cost() +
		1024
}

func (f FfPrefixSha256) fingerprintContents() []byte {
	content := struct {
		Prefix           []byte      `asn1:"tag:0"`
		MaxMessageLength uint32      `asn1:"tag:1"`
		SubCondition     interface{} `asn1:"tag:2,explicit,choice:condition"`
	}{
		Prefix:           f.Prefix,
		MaxMessageLength: f.MaxMessageLength,
		SubCondition:     castToEncodableCondition(f.SubCondition()),
	}

	encoded, err := ASN1Context.Encode(content)
	if err != nil {
		panic(err) //TODO check when this can happen
	}

	return encoded
}

func (f FfPrefixSha256) fingerprint() []byte {
	hash := sha256.Sum256(f.fingerprintContents())
	return hash[:]
}

func (f FfPrefixSha256) subConditionTypes() ConditionTypeSet {
	var set ConditionTypeSet
	if f.IsFulfilled() {
		set.addRelevant(f.SubFulfillment)
	} else {
		set.addRelevant(f.subCondition)
	}
	// As per RFC:
	// This is the set of types and subtypes of all sub-crypto-conditions,
	// recursively excluding the type of the root crypto-condition.
	set.remove(f.ConditionType())
	return set
}

func (f FfPrefixSha256) Condition() *Condition {
	return NewCompoundCondition(f.ConditionType(), f.fingerprint(), f.Cost(), f.subConditionTypes())
}

func (f FfPrefixSha256) Encode() ([]byte, error) {
	return encodeFulfillment(f)
}

func (f FfPrefixSha256) Validate(condition *Condition, message []byte) error {
	if !matches(f, condition) {
		return fulfillmentDoesNotMatchConditionError
	}

	if !f.IsFulfilled() {
		return errors.New("cannot validate unfulfilled fulfillment.")
	}

	if len(message) > int(f.MaxMessageLength) {
		return errors.Errorf(
			"message length of %d exceeds limit of %d",
			len(message), f.MaxMessageLength)
	}

	buffer := new(bytes.Buffer)
	buffer.Write(f.Prefix)
	buffer.Write(message)
	newMessage := buffer.Bytes()

	return errors.Wrapf(f.SubFulfillment.Validate(nil, newMessage),
		"failed to validate sub-fulfillment with message %x", newMessage)
}
