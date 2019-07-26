package api_test

import (
	"testing"

	"github.com/airbloc/airbloc-go/provider/api"
	"github.com/airbloc/airbloc-go/shared/service"
)

func TestNewAppRegistryAPI(t *testing.T) {
	var backend service.Backend
	api.NewAppRegistryAPI(backend)
}
