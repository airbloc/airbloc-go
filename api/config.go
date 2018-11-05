package api

import "github.com/airbloc/airbloc-go/blockchain"

type Config struct {
	PrivateKeyPath string
	Port           int

	Service string
	API     string

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
}
