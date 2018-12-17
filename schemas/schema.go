package schemas

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/xeipuuv/gojsonschema"
)

// Schema is a data format definition used on Airbloc.
// Airbloc shares all schema through the apps, using common schema registry,
// and ensures that data ingested by apps should follow the format.
type Schema struct {
	Id     common.ID
	Name   string
	Schema string

	// for JSON schema validation
	jsonSchema *gojsonschema.Schema
}

// NewSchema creates a schema from given ID and string.
// an error will be returned when given schema does not met JSON Schema format.
// for details, please see https://json-schema.org/
func NewSchema(name string, schema string) (*Schema, error) {
	jsonSchema, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(schema))
	if err != nil {
		return nil, err
	}
	return &Schema{
		Name:       name,
		Schema:     schema,
		jsonSchema: jsonSchema,
	}, nil
}

// IsValidFormat returns true if the format of the data matches with
// the schema specified in the data collection.
func (schema *Schema) IsValidFormat(data common.Data) (bool, error) {
	payload := gojsonschema.NewStringLoader(data.Payload)
	results, err := schema.jsonSchema.Validate(payload)
	if err != nil {
		return false, err
	}
	return results.Valid(), nil
}
