package api

import (
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/pkg/errors"
)

// Airbloc implements Airbloc node service.
// it composes all service used by Airbloc.
type AirblocBackend struct {
	Kms           *key.Manager
	Ethclient     *blockchain.Client
	MetaDatabase  metadb.Database
	LocalDatabase localdb.Database
	Config        *Config

	Services map[string]Service
}

func NewAirblocBackend(config *Config) (*AirblocBackend, error) {
	nodeKey, err := key.Load(config.PrivateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load private key from the given path")
	}

	metaDatabase, err := metadb.NewBigchainDB(
		config.MetaDB.BigchainDBEndpoint,
		config.MetaDB.MongoDBEndpoint,
		nodeKey.DeriveEd25519KeyPair(),
		config.MetaDB.Version)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize metadatabase")
	}

	localDatabase, err := localdb.NewBadgerDatabase(config.LocalDB.Path, 1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize local database")
	}

	kms := key.NewManager(nodeKey, localDatabase)

	ethclient, err := blockchain.NewClient(config.Blockchain.Endpoint, nodeKey.PrivateKey, config.Blockchain.Option)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize Ethereum client")
	}

	return &AirblocBackend{
		Kms:           kms,
		Ethclient:     ethclient,
		MetaDatabase:  metaDatabase,
		LocalDatabase: localDatabase,
		Config:        config,
	}, nil
}

func (airbloc *AirblocBackend) Attach(name string, service Service) {
	airbloc.Services[name] = service
}

func (airbloc *AirblocBackend) Start() error {
	for name, service := range airbloc.Services {
		if err := service.Start(); err != nil {
			return errors.Wrapf(err, "failed to start %s service", name)
		}
	}
}

func (airbloc *AirblocBackend) Stop() {
	for _, service := range airbloc.Services {
		service.Stop()
	}
}
