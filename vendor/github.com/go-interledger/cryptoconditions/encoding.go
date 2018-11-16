package cryptoconditions

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/stevenroose/asn1"
)

//TODO IMPLICIT TAGGING
// Because our ASN.1 package does not support implicit tagging, we manually
// tag all values in all structs.  Normally, a flag IMPLICIT on the choices for
// conditions and fulfillments should make it possible to remove that manual
// tag numbers.

// ASN1Context defines the ASN.1 context that is used to encode and decode
// objects.  It explicitly requires encoding and decoding to happen in strict
// DER format and it also defines the CHOICE mapping for
// fulfillments (`fulfillment`) and conditions (`condition`).
var ASN1Context *asn1.Context

type encodablePreimageSha256 struct {
	Fingerprint []byte `asn1:"tag:0"`
	Cost        int    `asn1:"tag:1"`
}

type encodablePrefixSha256 struct {
	Fingerprint []byte         `asn1:"tag:0"`
	Cost        int            `asn1:"tag:1"`
	SubTypes    asn1.BitString `asn1:"tag:2"`
}

type encodableThresholdSha256 struct {
	Fingerprint []byte         `asn1:"tag:0"`
	Cost        int            `asn1:"tag:1"`
	SubTypes    asn1.BitString `asn1:"tag:2"`
}

type encodableRsaSha256 struct {
	Fingerprint []byte `asn1:"tag:0"`
	Cost        int    `asn1:"tag:1"`
}

type encodableEd25519Sha256 struct {
	Fingerprint []byte `asn1:"tag:0"`
	Cost        int    `asn1:"tag:1"`
}

// castToEncodableCondition translates the condition to an encodable struct.
func castToEncodableCondition(condition *Condition) interface{} {
	switch condition.Type() {
	case CTPreimageSha256:
		return encodablePreimageSha256{
			Fingerprint: condition.Fingerprint(),
			Cost:        condition.Cost(),
		}

	case CTPrefixSha256:
		return encodablePrefixSha256{
			Fingerprint: condition.Fingerprint(),
			Cost:        condition.Cost(),
			SubTypes:    asn1.BitString(condition.SubTypes()),
		}

	case CTThresholdSha256:
		return encodableThresholdSha256{
			Fingerprint: condition.Fingerprint(),
			Cost:        condition.Cost(),
			SubTypes:    asn1.BitString(condition.SubTypes()),
		}

	case CTRsaSha256:
		return encodableRsaSha256{
			Fingerprint: condition.Fingerprint(),
			Cost:        condition.Cost(),
		}

	case CTEd25519Sha256:
		return encodableEd25519Sha256{
			Fingerprint: condition.Fingerprint(),
			Cost:        condition.Cost(),
		}
	}
	return nil
}

// encodeCondition encodes the given condition to it's DER encoding.
func encodeCondition(condition *Condition) ([]byte, error) {
	var encoded = castToEncodableCondition(condition)

	//TODO determine when an error is possible
	encoding, err := ASN1Context.EncodeWithOptions(encoded, "choice:condition")
	if err != nil {
		return nil, errors.Wrap(err, "ASN.1 encoding failed")
	}
	return encoding, nil
}

// DecodeCondition decodes the DER encoding of a condition.
func DecodeCondition(encodedCondition []byte) (*Condition, error) {
	var obj interface{}
	rest, err := ASN1Context.DecodeWithOptions(
		encodedCondition, &obj, "choice:condition")
	if err != nil {
		return nil, errors.Wrap(err, "ASN.1 decoding failed")
	}
	if len(rest) != 0 {
		return nil, errors.Errorf(
			"Encoding was not minimal. Excess bytes: %x", rest)
	}

	var cond *Condition
	switch obj.(type) {
	case encodablePreimageSha256:
		c := obj.(encodablePreimageSha256)
		cond = NewSimpleCondition(CTPreimageSha256, c.Fingerprint, c.Cost)
	case encodablePrefixSha256:
		c := obj.(encodablePrefixSha256)
		cond = NewCompoundCondition(CTPrefixSha256, c.Fingerprint, c.Cost, ConditionTypeSet(c.SubTypes))
	case encodableThresholdSha256:
		c := obj.(encodableThresholdSha256)
		cond = NewCompoundCondition(CTThresholdSha256, c.Fingerprint, c.Cost, ConditionTypeSet(c.SubTypes))
	case encodableRsaSha256:
		c := obj.(encodableRsaSha256)
		cond = NewSimpleCondition(CTRsaSha256, c.Fingerprint, c.Cost)
	case encodableEd25519Sha256:
		c := obj.(encodableEd25519Sha256)
		cond = NewSimpleCondition(CTEd25519Sha256, c.Fingerprint, c.Cost)

	default:
		return nil, errors.New("encoding was not a condition")
	}

	return cond, nil
}

// encodeFulfillment encodes the given fulfillment to it's DER encoding.
func encodeFulfillment(fulfillment Fulfillment) ([]byte, error) {
	//TODO determine when an error is possible
	encoded, err := ASN1Context.EncodeWithOptions(
		fulfillment, "choice:fulfillment")
	if err != nil {
		return nil, errors.Wrap(err, "ASN.1 encoding failed")
	}
	return encoded, nil
}

// DecodeFulfillment decodes the DER encoding of a fulfillment.
func DecodeFulfillment(encodedFulfillment []byte) (Fulfillment, error) {
	var obj interface{}
	rest, err := ASN1Context.DecodeWithOptions(
		encodedFulfillment, &obj, "choice:fulfillment")
	if err != nil {
		return nil, errors.Wrap(err, "ASN.1 decoding failed")
	}
	if len(rest) != 0 {
		return nil, errors.Errorf(
			"Encoding was not minimal. Excess bytes: %x", rest)
	}

	// Do some reflection magic to derive a pointer to the struct in obj.
	ptr := reflect.Indirect(reflect.New(reflect.TypeOf(obj)))
	ptr.Set(reflect.ValueOf(obj))
	obj = ptr.Addr().Interface()

	// Check whether the object we got is in fact a Fulfillment.
	fulfillment, ok := obj.(Fulfillment)
	if !ok {
		return nil, errors.New("Encoded object was not a fulfillment")
	}
	return fulfillment, nil
}

// buildAsn1Context builds the context for ASN.1 encoding and decoding.
// It forces the use of DER and specifies the tags for the CHOICES used for
// conditions and fulfillments.
func buildASN1Context() *asn1.Context {
	ctx := asn1.NewContext()
	ctx.SetDer(true, true)

	// Define the Condition CHOICE element.
	conditionChoices := []asn1.Choice{
		{
			Options: fmt.Sprintf("tag:%d", CTPreimageSha256),
			Type:    reflect.TypeOf(encodablePreimageSha256{}),
		},
		{
			Options: fmt.Sprintf("tag:%d", CTPrefixSha256),
			Type:    reflect.TypeOf(encodablePrefixSha256{}),
		},
		{
			Options: fmt.Sprintf("tag:%d", CTThresholdSha256),
			Type:    reflect.TypeOf(encodableThresholdSha256{}),
		},
		{
			Options: fmt.Sprintf("tag:%d", CTRsaSha256),
			Type:    reflect.TypeOf(encodableRsaSha256{}),
		},
		{
			Options: fmt.Sprintf("tag:%d", CTEd25519Sha256),
			Type:    reflect.TypeOf(encodableEd25519Sha256{}),
		},
	}
	if err := ctx.AddChoice("condition", conditionChoices); err != nil {
		panic(err)
	}

	// Define the Fulfillment CHOICE element.
	fulfillmentChoices := []asn1.Choice{
		{
			Options: fmt.Sprintf("tag:%d", CTPreimageSha256),
			Type:    reflect.TypeOf(FfPreimageSha256{}),
		},
		{
			Options: fmt.Sprintf("tag:%d", CTPrefixSha256),
			Type:    reflect.TypeOf(FfPrefixSha256{}),
		},
		{
			Options: fmt.Sprintf("tag:%d", CTThresholdSha256),
			Type:    reflect.TypeOf(FfThresholdSha256{}),
		},
		{
			Options: fmt.Sprintf("tag:%d", CTRsaSha256),
			Type:    reflect.TypeOf(FfRsaSha256{}),
		},
		{
			Options: fmt.Sprintf("tag:%d", CTEd25519Sha256),
			Type:    reflect.TypeOf(FfEd25519Sha256{}),
		},
	}
	if err := ctx.AddChoice("fulfillment", fulfillmentChoices); err != nil {
		panic(err)
	}

	return ctx
}

func init() {
	ASN1Context = buildASN1Context()
}
