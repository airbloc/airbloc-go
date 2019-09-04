package service

import (
	"time"

	defaults "github.com/mcuadros/go-defaults"
)

type Config struct {
	PrivateKeyPath string `default:"private.key" yaml:"privateKeyPath"`
	Host           string `default:"localhost" yaml:"host"`
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
		MongoDBEndpoint string `default:"mongodb://localhost:27017" yaml:"mongoDbEndpoint"`
		Version         int    `default:"1"`
	} `yaml:"metaDb"`

	ResourceDB struct {
		Endpoint string `default:"localhost:9090"`
	}

	Blockchain struct {
		Endpoint string `default:"http://localhost:8545" yaml:"endpoint"`
		Options  struct {
			MinConfirmations int `default:"1" yaml:"minConfirmations"`
		}
		DeploymentPath string `default:"deployment.local.json" yaml:"deploymentPath"`
	} `yaml:"blockchain"`

	Warehouse struct {
		DefaultStorage string `default:"local" yaml:"defaultStorage"`

		Http struct {
			Timeout         time.Duration `default:"30s"`
			MaxConnsPerHost int           `default:"5" yaml:"maxConnsPerHost"`
		}

		LocalStorage struct {
			SavePath string `default:"local/warehouse" yaml:"savepath"`
			Endpoint string `default:"http://localhost:80" yaml:"endpoint"`
		} `yaml:"localStorage"`

		S3 struct {
			Region     string `default:"ap-northeast-1" yaml:"region"`
			AccessKey  string `yaml:"accessKey"`
			SecretKey  string `yaml:"secretKey"`
			Token      string `default:"" yaml:"token"`
			Bucket     string `yaml:"bucket"`
			PathPrefix string `yaml:"prefix"`
		}

		Debug struct {
			DisableUserAuthValidation bool `default:"false" yaml:"disableUserAuthValidation"`
			DisableSchemaValidation   bool `default:"false" yaml:"disableSchemaValidation"`
		}
	} `yaml:"warehouse"`

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
