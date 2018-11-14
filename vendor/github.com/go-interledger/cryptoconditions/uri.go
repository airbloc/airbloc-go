package cryptoconditions

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/kalaspuffar/base64url"
	"github.com/pkg/errors"
)

// generateURI generates a URI for the given condition.
func generateURI(condition *Condition) string {
	params := make(url.Values)
	params.Set("cost", fmt.Sprintf("%d", condition.Cost()))
	params.Set("fpt", strings.ToLower(condition.Type().String()))

	if condition.Type().IsCompound() {
		subtypesString := ""
		for _, st := range condition.SubTypes().AllTypes() {
			subtypesString += strings.ToLower(st.String()) + ","
		}
		if subtypesString != "" {
			params.Set("subtypes", strings.TrimSuffix(subtypesString, ","))
		}
	}

	encodedFingerprint := base64url.Encode(condition.Fingerprint())
	uri := url.URL{
		Scheme:   "ni",
		Path:     "/sha-256;" + encodedFingerprint,
		RawQuery: params.Encode(),
	}

	return uri.String()
}

// ParseURI parses a URI into a Condition.
func ParseURI(uri string) (*Condition, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse URI")
	}
	params := u.Query()

	// Find the condition type.
	conditionTypeString := strings.ToUpper(params.Get("fpt"))
	conditionType, found := conditionTypeDictionary[conditionTypeString]
	if !found {
		return nil, errors.Errorf(
			"unknown condition type: %s", params.Get("fpt"))
	}

	// Parse the fingerprint.
	pathParts := strings.SplitN(u.Path, ";", 2)
	if len(pathParts) != 2 {
		return nil, errors.New("incorrectly formatted URI, no semicolon found")
	}
	fingerprint, err := base64url.Decode(pathParts[1])
	if err != nil {
		return nil, errors.Wrap(err,
			"failed to decode base64url encoded fingerprint")
	}

	// Parse cost.
	parsedInt, err := strconv.ParseInt(params.Get("cost"), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err,
			"failed to parse cost value %s", params.Get("cost"))
	}
	cost := int(parsedInt)

	// Parse subtypes.
	subtypeSet := &ConditionTypeSet{}
	if params.Get("subtypes") != "" {
		subtypeStrings := strings.Split(params.Get("subtypes"), ",")
		for _, subtypeString := range subtypeStrings {
			subType, found := conditionTypeDictionary[strings.ToUpper(subtypeString)]
			if !found {
				return nil, errors.Errorf(
					"unknown condition type in subconditions: %s",
					subtypeString)
			}
			subtypeSet.add(subType)
		}
	}

	return &Condition{
		conditionType: conditionType,
		fingerprint:   fingerprint,
		cost:          cost,
		subTypes:      *subtypeSet,
	}, nil
}
