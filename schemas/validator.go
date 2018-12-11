package schemas

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/azer/logger"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

// Validator ensures data providers to only register data to Airbloc
// from users who have allowed collection of the data through DAuth.
type Validator struct {
	jsonSchema *gojsonschema.Schema
	log        *logger.Logger
}

// NewValidator creates validator instance.
func NewValidator(schema *Schema) (*Validator, error) {
	jsonSchema, err := gojsonschema.NewSchema(gojsonschema.NewGoLoader(schema.Schema))
	if err != nil {
		return nil, errors.Wrap(err, "invalid JSONSchema")
	}
	return &Validator{
		jsonSchema: jsonSchema,
		log:        logger.New("schema-validator"),
	}, nil
}

// IsValidFormat returns true if the format of the data matches with
// the schema specified in the data collection.
func (validator *Validator) IsValidFormat(data common.Data) bool {
	payload := gojsonschema.NewStringLoader(data.Payload)
	results, err := validator.jsonSchema.Validate(payload)
	if err != nil {
		validator.log.Error("error: %s", err.Error(), logger.Attrs{"user": data.OwnerAnid.Hex()})
		return false
	}
	return results.Valid()
}
