package node

import (
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/database/localdb"
	"github.com/airbloc/airbloc-go/shared/database/metadb"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/p2p"
)

type Backend interface {
	Kms() key.Manager
	Client() *blockchain.Client
	MetaDatabase() metadb.Database
	LocalDatabase() localdb.Database
	Config() *Config
	P2P() p2p.Server

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
