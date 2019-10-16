package service

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env string `default:"production"`

	// private keys
	Key     string `yaml:"key"`
	KeyPath string `default:"private.key" yaml:"keyPath" split_words:"true"`

	// api/rpc listen address
	Host string `default:"0.0.0.0" yaml:"host"`
	Port int    `default:"2471" yaml:"port"`

	// logger settings
	LogLevel  string `default:"*" yaml:"logLevel" split_words:"true"`
	LogFilter string `default:"*" yaml:"logFilter" split_words:"true"`

	P2P struct {
		ListenAddr string   `default:"/ip4/0.0.0.0/tcp/2472" yaml:"listenAddr" split_words:"true"`
		BootNodes  []string `yaml:"bootNodes" split_words:"true"`
	} `yaml:"p2p"`

	LocalDB struct {
		Path    string `default:"local/"`
		Version int    `default:"1"`
	} `yaml:"localDB" split_words:"true"`

	MetaDB struct {
		MongoDBEndpoint string `default:"mongodb://localhost:27017" yaml:"mongoDbEndpoint" split_words:"true"`
		Version         int    `default:"1"`
	} `yaml:"metaDB" split_words:"true"`

	ResourceDB struct {
		Endpoint string `default:"localhost:9090"`
	} `yaml:"resourceDB" split_words:"true"`

	Blockchain struct {
		Endpoint string `default:"http://localhost:8545" yaml:"endpoint"`
		Options  struct {
			MinConfirmations int `default:"1" yaml:"minConfirmations" split_words:"true"`
		}
		DeploymentPath string `default:"deployment.local.json" yaml:"deploymentPath" split_words:"true"`
	} `yaml:"blockchain"`

	Warehouse struct {
		DefaultStorage string `default:"local" yaml:"defaultStorage" split_words:"true"`

		Http struct {
			Timeout         time.Duration `default:"30s"`
			MaxConnsPerHost int           `default:"5" yaml:"maxConnsPerHost" split_words:"true"`
		} `yaml:"http"`

		LocalStorage struct {
			SavePath string `default:"local/warehouse" yaml:"savepath" split_words:"true"`
			Endpoint string `default:"http://localhost:80" yaml:"endpoint"`
		} `yaml:"localStorage" split_words:"true"`

		S3 struct {
			Region     string `default:"ap-northeast-1" yaml:"region"`
			Bucket     string `yaml:"bucket"`
			PathPrefix string `yaml:"prefix" split_words:"true"`
		}

		Debug struct {
			DisableUserAuthValidation bool `default:"false" yaml:"disableUserAuthValidation" split_words:"true"`
			DisableSchemaValidation   bool `default:"false" yaml:"disableSchemaValidation" split_words:"true"`
		}
	} `yaml:"warehouse"`

	Controller struct {
		AccountIds []string `yaml:"accountIds" split_words:"true"`
	} `yaml:"controller" split_words:"true"`
}

// NewConfig returns node configurations with default value.
func NewConfig() *Config {
	config := new(Config)
	envconfig.MustProcess("airbloc", config)

	switch config.Env {
	case "production":
		// set production defaults
	case "development":
		// set development defaults
	case "testing":
		// set testing defaults
	}

	return config
}
