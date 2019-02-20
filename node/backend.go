package node

import (
	"os"
	"os/signal"
	"runtime"

	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
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
	p2pServer     p2p.Server
	services      map[string]Service
}

func NewAirblocBackend(nodeKey *key.Key, config *Config) (Backend, error) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	metaDatabase, err := metadb.NewBigchainDB(
		config.MetaDB.BigchainDBEndpoint,
		config.MetaDB.MongoDBEndpoint,
		config.MetaDB.ProxyEndpoint,
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

	// setup P2P
	// bootnode information should be given from config.
	var bootInfos []peerstore.PeerInfo
	for _, addr := range config.P2P.BootNodes {
		m, err := multiaddr.NewMultiaddr(addr)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid libp2p multiaddr: %s", addr)
		}
		bootInfo, err := peerstore.InfoFromP2pAddr(m)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid P2P peer address: %s", addr)
		}
		bootInfos = append(bootInfos, *bootInfo)
	}

	addr, err := multiaddr.NewMultiaddr(config.P2P.ListenAddr)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid listen address: %s", config.P2P.ListenAddr)
	}
	p2pServer, err := p2p.NewAirblocServer(nodeKey, addr, bootInfos)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize P2P server")
	}

	// setup ethereum client
	clientOpt := blockchain.ClientOpt{
		Confirmation:   config.Blockchain.Options.MinConfirmations,
		DeploymentPath: config.Blockchain.DeploymentPath,
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
		p2pServer:     p2pServer,
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

func (airbloc *AirblocBackend) P2P() p2p.Server {
	return airbloc.p2pServer
}

func (airbloc *AirblocBackend) Config() *Config {
	return airbloc.config
}

func (airbloc *AirblocBackend) Start() error {
	if err := airbloc.P2P().Start(); err != nil {
		return errors.Wrapf(err, "failed to start P2P service")
	}
	if err := p2p.StartNameServer(airbloc.P2P()); err != nil {
		return errors.Wrapf(err, "failed to start P2P address lookup service")
	}

	for name, service := range airbloc.services {
		if err := service.Start(); err != nil {
			return errors.Wrapf(err, "failed to start %s service", name)
		}
	}

	// wait for interrupt
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	select {
	case <-interruptCh:
		break
	}
	return nil
}

func (airbloc *AirblocBackend) Stop() {
	for _, service := range airbloc.services {
		service.Stop()
	}
	airbloc.p2pServer.Stop()
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
