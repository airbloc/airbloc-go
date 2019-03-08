package service

type Service interface {
	Start() error
	Stop()
}

type Constructor func(backend Backend) (Service, error)
