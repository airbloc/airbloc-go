package cryptoconditions

import "bytes"

// max returns the highest of both integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the lowest of both integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// fulfills determines if the fulfillment is able to fulfill the condition.
// The trivial way of doing this is to compare the ff.Condition() with the
// condition. However, this requires the generation of the condition while we
// can determine more efficiently whether or not they are going to match.
func matches(ff Fulfillment, cond *Condition) bool {
	if cond == nil {
		return true
	}

	if ff.ConditionType() != cond.Type() {
		return false
	}

	if ff.Cost() > cond.Cost() {
		return false
	}

	if !bytes.Equal(ff.fingerprint(), cond.Fingerprint()) {
		return false
	}

	//TODO subtype check? or just condition check after all?
	return ff.Condition().Equals(cond)
}
