package service

import (
	"os"
	"testing"
	"time"

	"github.com/kelseyhightower/envconfig"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

const (
	TestConfigYaml = `
key: "test"

p2p:
  bootNodes: ["a", "b", "c", "d", "e"]

blockchain:
  endpoint: "test_endpoint"

warehouse:
  http:
    timeout: "24h"
  debug:
    disableUserAuthValidation: true
`
	TestConfigPath = "../../test/test_config.yml"
)

func TestConfig(t *testing.T) {
	Convey("Testing service configuration", t, func() {
		actualConfig := Config{}
		expectedConfig := Config{
			Env: "production",

			Key:     "",
			KeyPath: "private.key",

			Host: "localhost",
			Port: 2471,

			LogLevel:  "*",
			LogFilter: "*",
			P2P: struct {
				ListenAddr string   `default:"/ip4/0.0.0.0/tcp/2472" yaml:"listenAddr" split_words:"true"`
				BootNodes  []string `yaml:"bootNodes" split_words:"true"`
			}{
				ListenAddr: "/ip4/0.0.0.0/tcp/2472",
				BootNodes:  nil,
			},
			LocalDB: struct {
				Path    string `default:"local/"`
				Version int    `default:"1"`
			}{
				Path:    "local/",
				Version: 1,
			},
			MetaDB: struct {
				MongoDBEndpoint string `default:"mongodb://localhost:27017" yaml:"mongoDbEndpoint" split_words:"true"`
				Version         int    `default:"1"`
			}{
				MongoDBEndpoint: "mongodb://localhost:27017",
				Version:         1,
			},
			ResourceDB: struct {
				Endpoint string `default:"localhost:9090"`
			}{
				Endpoint: "localhost:9090",
			},
			Blockchain: struct {
				Endpoint string `default:"http://localhost:8545" yaml:"endpoint"`
				Options  struct {
					MinConfirmations int `default:"1" yaml:"minConfirmations" split_words:"true"`
				}
				DeploymentPath string `default:"deployment.local.json" yaml:"deploymentPath" split_words:"true"`
			}{
				Endpoint: "http://localhost:8545",
				Options: struct {
					MinConfirmations int `default:"1" yaml:"minConfirmations" split_words:"true"`
				}{
					MinConfirmations: 1,
				},
				DeploymentPath: "deployment.local.json",
			},
			Warehouse: struct {
				DefaultStorage string `default:"local" yaml:"defaultStorage" split_words:"true"`
				Http           struct {
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
			}{
				DefaultStorage: "local",
				Http: struct {
					Timeout         time.Duration `default:"30s"`
					MaxConnsPerHost int           `default:"5" yaml:"maxConnsPerHost" split_words:"true"`
				}{
					Timeout:         30 * time.Second,
					MaxConnsPerHost: 5,
				},
				LocalStorage: struct {
					SavePath string `default:"local/warehouse" yaml:"savepath" split_words:"true"`
					Endpoint string `default:"http://localhost:80" yaml:"endpoint"`
				}{
					SavePath: "local/warehouse",
					Endpoint: "http://localhost:80",
				},
				S3: struct {
					Region     string `default:"ap-northeast-1" yaml:"region"`
					Bucket     string `yaml:"bucket"`
					PathPrefix string `yaml:"prefix" split_words:"true"`
				}{
					Region: "ap-northeast-1",
				},
				Debug: struct {
					DisableUserAuthValidation bool `default:"false" yaml:"disableUserAuthValidation" split_words:"true"`
					DisableSchemaValidation   bool `default:"false" yaml:"disableSchemaValidation" split_words:"true"`
				}{
					DisableSchemaValidation:   false,
					DisableUserAuthValidation: false,
				},
			},
			Controller: struct {
				AccountIds []string `yaml:"accountIds" split_words:"true"`
			}{},
		}
		os.Clearenv()

		Convey("Should set to default correctly.", func() {
			envconfig.MustProcess("TEST", &actualConfig)
			So(actualConfig, ShouldResemble, expectedConfig)
		})

		Convey("Should get environment correctly.", func() {
			So(os.Setenv("TEST_KEY", "test"), ShouldBeNil)
			So(os.Setenv("TEST_P2P_BOOT_NODES", "a,b,c,d,e"), ShouldBeNil)
			So(os.Setenv("TEST_BLOCKCHAIN_ENDPOINT", "test_endpoint"), ShouldBeNil)
			So(os.Setenv("TEST_WAREHOUSE_HTTP_TIMEOUT", "24h"), ShouldBeNil)
			So(os.Setenv("TEST_WAREHOUSE_DEBUG_DISABLE_USER_AUTH_VALIDATION", "true"), ShouldBeNil)
			envconfig.MustProcess("TEST", &actualConfig)

			expectedConfig.Key = "test"
			expectedConfig.P2P.BootNodes = []string{"a", "b", "c", "d", "e"}
			expectedConfig.Blockchain.Endpoint = "test_endpoint"
			expectedConfig.Warehouse.Http.Timeout = 24 * time.Hour
			expectedConfig.Warehouse.Debug.DisableUserAuthValidation = true
			So(actualConfig, ShouldResemble, expectedConfig)
		})

		Convey("Should get config from external config file.", func() {
			envconfig.MustProcess("TEST", &actualConfig)

			So(yaml.Unmarshal([]byte(TestConfigYaml), &actualConfig), ShouldBeNil)

			expectedConfig.Key = "test"
			expectedConfig.P2P.BootNodes = []string{"a", "b", "c", "d", "e"}
			expectedConfig.Blockchain.Endpoint = "test_endpoint"
			expectedConfig.Warehouse.Http.Timeout = 24 * time.Hour
			expectedConfig.Warehouse.Debug.DisableUserAuthValidation = true
			So(actualConfig, ShouldResemble, expectedConfig)
		})
	})
}
