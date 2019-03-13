package api

import "github.com/airbloc/airbloc-go/shared/service"

type API interface {
	AttachToAPI(api *Service)
}

type Constructor func(backend service.Backend) (API, error)
