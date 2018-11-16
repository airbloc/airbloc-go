package api

import (
	"time"
)

type Config struct {
	PrivateKeyPath string `default:"private.key" yaml:"privateKeyPath"`
	Port           int    `default:"9124" yaml:"port"`

	LocalDB struct {
		Path    string `default:"local/"`
		Version int    `default:"1"`
	} `yaml:"localDb"`

	MetaDB struct {
		BigchainDBEndpoint string `default:"http://localhost:9984" yaml:"bigchainDbEndpoint"`
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

	Warehouse struct {
		DefaultStorage string `default:"local" yaml:"defaultStorage"`

		Http struct {
			Timeout         time.Duration `default:"30s"`
			MaxConnsPerHost int           `default:"5" yaml:"maxConnsPerHost"`
		}

		LocalStorage struct {
			SavePath string `default:"local/warehouse"`
			Endpoint string `default:"http://localhost:80"`
		}

		S3 struct {
			Region     string `yaml:"region"`
			AccessKey  string `yaml:"accessKey"`
			Bucket     string `yaml:"bucket"`
			PathPrefix string `yaml:"pathPrefix"`
		}
	}
}
