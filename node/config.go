package node

import (
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	PrivateKeyPath string `default:"private.key" yaml:"privateKeyPath"`
	Port           int    `default:"9124" yaml:"port"`

	P2P struct {
		ListenAddr string   `default:"/ip4/0.0.0.0/tcp/2470" yaml:"listenAddr"`
		BootNodes  []string `yaml:"bootNodes"`
	} `yaml:"p2p"`

	LocalDB struct {
		Path    string `default:"local/"`
		Version int    `default:"1"`
	} `yaml:"localDb"`

	MetaDB struct {
		BigchainDBEndpoint string `default:"http://localhost:9984" yaml:"bigchainDbEndpoint"`
		ProxyEndpoint      string `default:"http://localhost:9983" yaml:"proxyEndpoint"`
		MongoDBEndpoint    string `default:"mongodb://localhost:27017" yaml:"mongoDbEndpoint"`
		Version            int    `default:"1"`
	} `yaml:"metaDb"`

	Blockchain struct {
		Endpoint string `default:"http://localhost:8545"`
		Options  struct {
			MinConfirmations int `default:"1" yaml:"minConfirmations"`
		}
		DeploymentPath string `default:"deployment.local.json" yaml:"deploymentPath"`
	}

	Warehouse warehouse.Config

	UserDelegate struct {
		AccountIds []string `yaml:"accountIds"`
	} `yaml:"userDelegate"`
}

// NewConfig returns node configurations with default value.
func NewConfig() *Config {
	config := new(Config)
	defaults.SetDefaults(config)
	return config
}
