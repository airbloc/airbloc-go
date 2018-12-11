package schemas

import "github.com/airbloc/airbloc-go/common"

// Schema is a data format definition used on Airbloc.
// Airbloc shares all schema through the apps, using common schema registry,
// and ensures that data ingested by apps should follow the format.
type Schema struct {
	Id     common.ID
	Name   string
	Schema map[string]interface{}
}
