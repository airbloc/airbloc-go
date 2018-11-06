package api

import (
	"time"

	"github.com/airbloc/airbloc-go/blockchain"
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

	Warehouse struct {
		DefaultStorage string

		Http struct {
			Timeout         time.Duration
			MaxConnsPerHost int
		}

		LocalStorage struct {
			SavePath string
			Endpoint string
		}

		S3 struct {
			Region     string
			AccessKey  string
			Bucket     string
			PathPrefix string
		}
	}
}

func DefaultConfig() (config *Config) {
	config = new(Config)
	config.PrivateKeyPath = "private.key"
	config.Port = 9124

	config.LocalDB.Path = "local.db"
	config.LocalDB.Version = 1

	config.MetaDB.BigchainDBEndpoint = "http://localhost:9984/api/v1"
	config.MetaDB.MongoDBEndpoint = "mongo://localhost:27017"
	config.MetaDB.Version = 1

	config.Blockchain.Endpoint = "http://localhost:8545"
	config.Blockchain.Option.Confirmation = 1

	config.Warehouse.DefaultStorage = "local"
	config.Warehouse.Http.Timeout = 30 * time.Second
	config.Warehouse.Http.MaxConnsPerHost = 10

	config.Warehouse.LocalStorage.SavePath = "local/warehouse/"
	config.Warehouse.LocalStorage.Endpoint = "http://localhost:9125/"
	return
}
