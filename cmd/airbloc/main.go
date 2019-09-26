package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/airbloc/airbloc-go/consumer"
	"github.com/airbloc/airbloc-go/controller"
	"github.com/airbloc/airbloc-go/provider"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/airbloc/logger"
	"github.com/klaytn/klaytn/crypto"
	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	ablLog = logger.New(name)
	config = service.NewConfig()

	rootCmd = &cobra.Command{
		Use:     name,
		Short:   descShort,
		Long:    descLong,
		Version: Version,
	}

	// top-level flags, independent from node.Config
	rootFlags struct {
		configPath string
		dataDir    string

		keyPath string
		private string

		blockchainEndpoint string
		deploymentPath     string
		mongoEndpoint      string
		bootnodes          []string

		verbose   bool
		logLevel  string
		logFilter string
	}

	// list of CLI commands and flags
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start Airbloc API server.",
		Long:  "Start Airbloc REST/gRPC API server.",
		Run:   start("warehouse,provider"),
		//Run:   start("provider,consumer,warehouse"),
	}

	dataControllerCmd = &cobra.Command{
		Use:   "controller",
		Short: "Start Airbloc data controller daemon.",
		Long:  "Start data controller daemon, watching and supervising user's data event.",
		Run:   start("warehouse,controller"),
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display a version.",
		Run: func(*cobra.Command, []string) {
			fmt.Println("Client:")
			fmt.Println("  Version:", Version)
			fmt.Println("  Git Commit:", GitCommit)
			fmt.Println("  Git Branch:", GitBranch)
			fmt.Println("  Build Date:", BuildDate)
		},
	}

	// list of available Services
	AvailableServices = map[string]service.Constructor{
		"provider":   provider.NewService,
		"consumer":   consumer.NewService,
		"controller": controller.NewService,
		"warehouse":  warehouse.NewService,
	}
)

func init() {
	log.SetFlags(log.Lshortfile)
	cobra.OnInitialize(loadConfig)
	rflags := rootCmd.PersistentFlags()

	rflags.StringVarP(&rootFlags.dataDir, "datadir", "d", "~/.airbloc", "Data directory")
	rflags.StringVarP(&rootFlags.configPath, "config", "c", "", "Config file")
	rflags.StringVarP(&rootFlags.keyPath, "keystore", "k", "", "Keystore file for node (default is $DATADIR/private.key)")
	rflags.StringVar(&rootFlags.private, "private", "", "Raw 32-byte private key with 0x prefix (Not Recommended)")

	rflags.StringVar(&rootFlags.blockchainEndpoint, "klaytn", "", "Klaytn RPC endpoint")
	rflags.StringVar(&rootFlags.deploymentPath, "deployment", "", "Path or URL of deployment.json")
	rflags.StringVar(&rootFlags.mongoEndpoint, "metadb", "", "Metadatabase endpoint")
	rflags.StringSliceVar(&rootFlags.bootnodes, "bootnodes", nil, "Bootstrap Node multiaddr for P2P")

	rflags.BoolVarP(&rootFlags.verbose, "verbose", "v", true, "Verbose output")
	rflags.StringVar(&rootFlags.logFilter, "logfilter", "*", "Log only from specific packages (e.g. warehouse,users)")

	// server options
	f := serverCmd.Flags()
	f.IntVarP(&config.Port, "port", "p", config.Port, "Port of gRPC Server API endpoint.")
	f.StringVar(&config.Warehouse.DefaultStorage, "warehouse.storage", config.Warehouse.DefaultStorage,
		"Type of warehouse storage. [local|s3|gcs|azure]")

	rootCmd.AddCommand(
		initCmd,
		serverCmd,
		dataControllerCmd,
		versionCmd,
	)
}

func loadConfig() {
	args := strings.Join(os.Args, " ")
	if strings.HasSuffix(args, "help") || strings.HasSuffix(args, "version") {
		return
	}

	// get data directory
	dataDir, err := home.Expand(rootFlags.dataDir)
	if err != nil {
		ablLog.Error("Error: failed to resolve data directory {}", err, rootFlags.dataDir)
	}
	if err = os.MkdirAll(dataDir, os.ModePerm); err != nil {
		ablLog.Error("Error: failed to create data directory {}", err, rootFlags.dataDir)
	}

	// override config
	if rootFlags.configPath != "" {
		configPath := rootFlags.configPath
		configPath = strings.Replace(configPath, "$DATADIR", dataDir, 1)
		configData, err := ioutil.ReadFile(configPath)
		if err != nil {
			ablLog.Error("Error: failed to read config file {}", err, configPath)
			os.Exit(1)
		}
		if err = yaml.Unmarshal(configData, config); err != nil {
			ablLog.Error("Error: failed to unmarshal config data {}", err, configPath)
			os.Exit(1)
		}
	}

	// setup loggers
	if rootFlags.verbose {
		rootFlags.logLevel = "*"
	}
	if rootFlags.logLevel != "*" {
		config.LogLevel = rootFlags.logLevel
	}
	if rootFlags.logFilter != "*" {
		config.LogFilter = rootFlags.logFilter
	}

	// setup keypath
	if rootFlags.keyPath != "" {
		config.KeyPath = rootFlags.keyPath
	} else {
		config.KeyPath = strings.Replace(config.KeyPath, "$DATADIR", dataDir, 1)
	}

	if rootFlags.private != "" {
		config.Key = rootFlags.private
	}
	if rootFlags.blockchainEndpoint != "" {
		config.Blockchain.Endpoint = rootFlags.blockchainEndpoint
	}
	if rootFlags.deploymentPath != "" {
		config.Blockchain.DeploymentPath = rootFlags.deploymentPath
	}
	if rootFlags.mongoEndpoint != "" {
		config.MetaDB.MongoDBEndpoint = rootFlags.mongoEndpoint
	}
	if rootFlags.bootnodes != nil {
		config.P2P.BootNodes = rootFlags.bootnodes
	}

	logger.SetLogger(logger.NewStandardOutput(os.Stdout, config.LogLevel, config.LogFilter))
	log.SetOutput(os.Stderr)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		ablLog.Error("Error", err)
		os.Exit(1)
	}
}

func start(serviceNames string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		nodeKey := loadNodeKey()
		backend, err := service.NewAirblocBackend(nodeKey, config)
		if err != nil {
			ablLog.Error("Error: init error", err)
			os.Exit(1)
		}
		defer backend.Stop()

		// attach services using config
		serviceNames := strings.Split(serviceNames, ",")
		registerServices(backend, serviceNames)

		if err := backend.Start(); err != nil {
			ablLog.Error("Error: failed to start airbloc", err)
			os.Exit(1)
		}
	}
}

func loadNodeKey() *key.Key {
	if config.Key != "" {
		if len(rootFlags.private) != 66 || !strings.HasPrefix(rootFlags.private, "0x") {
			ablLog.Error("Error: Invalid private key.")
			os.Exit(1)
		}
		rawKey, err := hex.DecodeString(strings.TrimPrefix(rootFlags.private, "0x"))
		if err != nil {
			ablLog.Error("Error: Failed to decode hex", err)
			os.Exit(1)
		}
		k, err := crypto.ToECDSA(rawKey)
		if err != nil {
			ablLog.Error("Error: Invalid ECDSA key", err)
			os.Exit(1)
		}
		return key.FromECDSA(k)
	} else {
		k, err := key.Load(config.KeyPath)
		if err != nil {
			ablLog.Error("Error: failed to load private key from the given path", err)
			os.Exit(1)
		}
		return k
	}
}

func registerServices(backend service.Backend, serviceNames []string) {
	for _, serviceName := range serviceNames {
		serviceConstructor, exists := AvailableServices[serviceName]
		if !exists {
			ablLog.Error("Error: service {} does not exist.", serviceName)
			os.Exit(1)
		}

		svc, err := serviceConstructor(backend)
		if err != nil {
			ablLog.Error("Error: failed to create service {}", err, serviceName)
			os.Exit(1)
		}
		backend.AttachService(serviceName, svc)
	}
}
