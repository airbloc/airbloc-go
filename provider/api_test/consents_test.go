package api_test

import (
	"testing"

	"github.com/airbloc/airbloc-go/provider/api"
	"github.com/airbloc/airbloc-go/shared/service"
)

func TestConsentsAPI(t *testing.T) {
	var backend service.Backend
	api.NewConsentsAPI(backend)
}
