package cryptoconditions

import "github.com/pkg/errors"

// Fulfillment defines the fulfillment interface.
type Fulfillment interface {
	// ConditionType returns the type of condition this fulfillment fulfills.
	ConditionType() ConditionType

	// Condition generates the condition that this fulfillment fulfills.
	Condition() *Condition

	// Cost calculates the cost metric of this fulfillment.
	Cost() int

	// Encode encodes the fulfillment into binary format.
	Encode() ([]byte, error)

	// Validate checks whether this fulfillment correctly validates the given
	// condition using the specified message.
	// It returns nil if it does, an error indicating the problem otherwise.
	Validate(*Condition, []byte) error

	// fingerprint calculates the fingerprint of the condition this fulfillment
	// fulfills.
	fingerprint() []byte

	// fingerprintContents returns data that is hashed when calculating the
	// fingerprint.
	fingerprintContents() []byte
}

// compoundConditionFulfillment is an interface that fulfillments for compound
// conditions have to implement to be able to indicate the condition types of
// their sub-fulfillments.
type compoundConditionFulfillment interface {
	// subConditionTypes returns the set with all the different types
	// amongst sub-conditions of this fulfillment.
	subConditionTypes() ConditionTypeSet
}

// fulfillmentDoesNotMatchConditionError is the error we throw when trying to
// validate a condition with a wrong
// fulfillment.
var fulfillmentDoesNotMatchConditionError = errors.New(
	"The fulfillment does not match the given condition")
