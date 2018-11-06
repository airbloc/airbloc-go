package api

import (
	"os"
	"time"

	"github.com/airbloc/airbloc-go/blockchain"
	"gopkg.in/yaml.v2"
)

type Config struct {
	PrivateKeyPath string `yaml:"private_key_path"`
	Port           int    `yaml:"port"`

	LocalDB struct {
		Path    string `yaml:"path"`
		Version int    `yaml:"version"`
	} `yaml:"local_db"`

	MetaDB struct {
		BigchainDBEndpoint string `yaml:"bigchain_db_endpoint"`
		MongoDBEndpoint    string `yaml:"mongo_db_endpoint"`
		Version            int    `yaml:"version"`
	} `yaml:"meta_db"`

	Blockchain struct {
		Endpoint string               `yaml:"endpoint"`
		Option   blockchain.ClientOpt `yaml:"option"`
	} `yaml:"blockchain"`

	Warehouse struct {
		DefaultStorage string `yaml:"default_storage"`

		Http struct {
			Timeout         time.Duration `yaml:"timeout"`
			MaxConnsPerHost int           `yaml:"max_conns_per_host"`
		} `yaml:"http"`

		LocalStorage struct {
			SavePath string `yaml:"save_path"`
			Endpoint string `yaml:"endpoint"`
		} `yaml:"local_storage"`

		S3 struct {
			Region     string `yaml:"region"`
			AccessKey  string `yaml:"access_key"`
			Bucket     string `yaml:"bucket"`
			PathPrefix string `yaml:"path_prefix"`
		} `yaml:"s3"`
	} `yaml:"warehouse"`
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

func ImportConfig(filepath string) (*Config, error) {
	config := new(Config)
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func (cfg *Config) Export(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	return encoder.Encode(cfg)
}
