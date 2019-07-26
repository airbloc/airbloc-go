package api_test

import (
	"testing"

	"github.com/airbloc/airbloc-go/provider/api"
	"github.com/airbloc/airbloc-go/shared/service"
)

func TestDataTypeRegistryAPI(t *testing.T) {
	var backend service.Backend
	api.NewDataTypeRegistryAPI(backend)
}
