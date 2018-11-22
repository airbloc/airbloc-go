package api

import (
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
)

type Backend interface {
	Kms() key.Manager
	Client() *blockchain.Client
	MetaDatabase() metadb.Database
	LocalDatabase() localdb.Database
	Config() *Config

	Service
	GetService(string) Service
	AttachService(string, Service)
	DetachService(string)
}

type API interface {
	AttachToAPI(api *APIService)
}

type Constructor func(airbloc Backend) (API, error)

type Service interface {
	Start() error
	Stop()
}

type ServiceConstructor func(airbloc Backend) (Service, error)
