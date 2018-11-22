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
	kms           key.Manager
	ethclient     *blockchain.Client
	metaDatabase  metadb.Database
	localDatabase localdb.Database
	config        *Config
	services      map[string]Service
}

func NewAirblocBackend(config *Config) (Backend, error) {
	nodeKey, err := key.Load(config.PrivateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load private key from the given path")
	}

	metaDatabase, err := metadb.NewBigchainDB(
		config.MetaDB.BigchainDBEndpoint,
		config.MetaDB.MongoDBEndpoint,
		nodeKey.DeriveBigchainDBKeyPair(),
		config.MetaDB.Version)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize metadatabase")
	}

	localDatabase, err := localdb.NewBadgerDatabase(config.LocalDB.Path, 1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize local database")
	}

	kms := key.NewKeyManager(nodeKey, localDatabase)

	// setup ethereum client
	clientOpt := blockchain.ClientOpt{
		Confirmation: config.Blockchain.Options.MinConfirmations,
	}
	ethclient, err := blockchain.NewClient(nodeKey, config.Blockchain.Endpoint, clientOpt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize Ethereum client")
	}
	ethclient.SetAccount(nodeKey)

	return &AirblocBackend{
		kms:           kms,
		ethclient:     ethclient,
		metaDatabase:  metaDatabase,
		localDatabase: localDatabase,
		config:        config,
		services:      make(map[string]Service),
	}, nil
}

func (airbloc *AirblocBackend) Kms() key.Manager {
	return airbloc.kms
}

func (airbloc *AirblocBackend) Client() *blockchain.Client {
	return airbloc.ethclient
}

func (airbloc *AirblocBackend) MetaDatabase() metadb.Database {
	return airbloc.metaDatabase
}

func (airbloc *AirblocBackend) LocalDatabase() localdb.Database {
	return airbloc.localDatabase
}

func (airbloc *AirblocBackend) Config() *Config {
	return airbloc.config
}

func (airbloc *AirblocBackend) Start() error {
	for name, service := range airbloc.services {
		if err := service.Start(); err != nil {
			return errors.Wrapf(err, "failed to start %s service", name)
		}
	}
	return nil
}

func (airbloc *AirblocBackend) Stop() {
	for _, service := range airbloc.services {
		service.Stop()
	}
	airbloc.ethclient.Close()
	airbloc.localDatabase.Close()
	airbloc.metaDatabase.Close()
}

func (airbloc *AirblocBackend) GetService(name string) Service {
	return airbloc.services[name]
}

func (airbloc *AirblocBackend) AttachService(name string, service Service) {
	airbloc.services[name] = service
}

func (airbloc *AirblocBackend) DetachService(name string) {
	delete(airbloc.services, name)
}
