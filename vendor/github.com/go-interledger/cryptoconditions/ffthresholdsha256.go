package cryptoconditions

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

//TODO Currently not working due to missing ASN.1 features from our dependency
// When resolved, remove all the mentions in the documentation.

// FfThresholdSha256 implements the THRESHOLD-SHA-256 fulfillment.
// This type of fulfillment is currently not supported.
type FfThresholdSha256 struct {
	Threshold uint16 `asn1:"-"`

	SubFulfillments []Fulfillment `asn1:"tag:0,explicit,set,choice:fulfillment"`
	SubConditions   []*Condition  `asn1:"tag:1,explicit,set,choice:condition"`
}

//TODO ADD NORMALIZE METHOD that makes sure the FF is of minimal size by replacing (threshold - nbFulfillments) fulfillments
// with their conditions, choosing those fulfillments that have the biggest (fulfillmentSize - conditionSize).

// NewThresholdSha256 creates a new THRESHOLD-SHA-256 fulfillment.
// This type of fulfillment is currently not supported.
func NewThresholdSha256(threshold uint16, subFulfillments []Fulfillment, subConditions []*Condition) *FfThresholdSha256 {
	return &FfThresholdSha256{
		Threshold:       threshold,
		SubFulfillments: subFulfillments,
		SubConditions:   subConditions,
	}
}

func (f FfThresholdSha256) ConditionType() ConditionType {
	return CTThresholdSha256
}

func (f FfThresholdSha256) Cost() int {
	// The cost is the sum of the F.threshold largest cost values of all
	// sub-conditions, added to 1024 times the total number of sub-conditions.
	conditionCosts := make([]int,
		len(f.SubFulfillments)+len(f.SubConditions))
	index := 0
	for _, fulfillment := range f.SubFulfillments {
		conditionCosts[index] = fulfillment.Cost()
		index++
	}
	for _, condition := range f.SubConditions {
		conditionCosts[index] = condition.Cost()
		index++
	}
	sort.Ints(conditionCosts)
	// We need the sum of the [threshold] highest costs.
	tHighest := conditionCosts[len(conditionCosts)-int(f.Threshold):]
	sum := 0
	for _, cost := range tHighest {
		sum += cost
	}
	return sum + 1024*len(conditionCosts)
}

func (f FfThresholdSha256) fingerprintContents() []byte {
	subConditions := make([]*Condition, len(f.SubFulfillments))
	for i, sff := range f.SubFulfillments {
		subConditions[i] = sff.Condition()
	}
	content := struct {
		Threshold     uint16       `asn1:"tag:0"`
		SubConditions []*Condition `asn1:"tag:1,explicit,set,choice:condition"`
	}{
		Threshold:     f.Threshold,
		SubConditions: subConditions,
	}

	encoded, err := ASN1Context.Encode(content)
	if err != nil {
		//TODO
		panic(err)
	}

	return encoded
}

func (f FfThresholdSha256) fingerprint() []byte {
	hash := sha256.Sum256(f.fingerprintContents())
	return hash[:]
}

func (f FfThresholdSha256) subConditionTypes() ConditionTypeSet {
	var set ConditionTypeSet
	for _, sff := range f.SubFulfillments {
		set.addRelevant(sff)
	}
	for _, sc := range f.SubConditions {
		set.addRelevant(sc)
	}
	// As per RFC:
	// This is the set of types and subtypes of all sub-crypto-conditions,
	// recursively excluding the type of the root crypto-condition.
	set.remove(f.ConditionType())
	return set
}

func (f FfThresholdSha256) Condition() *Condition {
	return NewCompoundCondition(f.ConditionType(), f.fingerprint(), f.Cost(), f.subConditionTypes())
}

func (f FfThresholdSha256) Encode() ([]byte, error) {
	return encodeFulfillment(f)
}

func (f FfThresholdSha256) Validate(condition *Condition, message []byte) error {
	if !matches(f, condition) {
		return fulfillmentDoesNotMatchConditionError
	}

	th := int(f.Threshold)
	if th == 0 {
		return nil
	}

	// Check if we have enough fulfillments.
	if len(f.SubFulfillments) < th {
		return fmt.Errorf("Not enough fulfillments: %v of %v", len(f.SubFulfillments), th)
	}

	// Try to verify the fulfillments one by one.
	for _, ff := range f.SubFulfillments {
		if ff.Validate(nil, message) == nil {
			th--
			if th == 0 {
				break
			}
		}
	}

	if th != 0 {
		return fmt.Errorf("Could only verify %v of %v fulfillments", int(f.Threshold)-th, th)
	}
	return nil
}
