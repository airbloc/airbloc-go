package api

type API interface {
	AttachToAPI(api *APIService)
}

type Constructor func(airbloc *AirblocBackend) (API, error)

type Service interface {
	Start() error
	Stop()
}

type ServiceConstructor func(airbloc *AirblocBackend) (Service, error)
