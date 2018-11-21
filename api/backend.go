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
	Kms           key.Manager
	Ethclient     blockchain.TxClient
	MetaDatabase  metadb.Database
	LocalDatabase localdb.Database
	Config        *Config
	services      map[string]Service
}

func NewAirblocBackend(config *Config) (*AirblocBackend, error) {
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
		Kms:           kms,
		Ethclient:     ethclient,
		MetaDatabase:  metaDatabase,
		LocalDatabase: localDatabase,
		Config:        config,
		services:      make(map[string]Service),
	}, nil
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
	airbloc.Ethclient.Close()
	airbloc.LocalDatabase.Close()
	airbloc.MetaDatabase.Close()
}

func (airbloc *AirblocBackend) GetService(name string) Service {
	return airbloc.services[name]
}

func (airbloc *AirblocBackend) AttachService(name string, service Service) {
	airbloc.Services[name] = service
}

func (airbloc *AirblocBackend) DettachService(name string) {
	delete(airbloc.Services, name)
}
