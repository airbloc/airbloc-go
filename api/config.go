package api

import (
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/warehouse"
)

type Config struct {
	PrivateKeyPath string
	Port           int

	LocalDB struct {
		Path    string
		Version int
	}

	MetaDB struct {
		BigchainDBEndpoint string
		MongoDBEndpoint    string
		Version            int
	}

	Blockchain struct {
		Endpoint string
		Option   blockchain.ClientOpt
	}

	Warehouse warehouse.Config
}

func DefaultConfig() (config *Config) {
	config.PrivateKeyPath = "private.key"
	config.Port = 9124

	config.LocalDB.Path = "local.db"
	config.LocalDB.Version = 1

	config.MetaDB.BigchainDBEndpoint = "http://localhost:9984/api/v1"
	config.MetaDB.MongoDBEndpoint = "mongo://localhost:27017"
	config.MetaDB.Version = 1

	config.Blockchain.Endpoint = "http://localhost:8545"
	config.Blockchain.Option.Confirmation = 1

	config.Warehouse = warehouse.DefaultConfig()
	return
}
