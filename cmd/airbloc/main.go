package main

import (
	"fmt"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse/service"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"strings"

	logger2 "github.com/airbloc/airbloc-go/logger"
	"github.com/airbloc/airbloc-go/node/userdelegateapi"
	"github.com/airbloc/airbloc-go/userdelegate"
	"github.com/airbloc/logger"
	log2 "log"

	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/node/serverapi"
	"github.com/jinzhu/configor"
)

var (
	log    = logger.New(name)
	config = node.NewConfig()

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

		verbose   bool
		logLevel  string
		logFilter string
	}

	// list of CLI commands and flags
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start Airbloc API server.",
		Long:  "Start Airbloc REST/gRPC API server.",
		Run:   start("api,warehouse", "apps,server.accounts,collections,data,dauth,exchange,schemas,warehouse"),
	}

	userDelegateCmd = &cobra.Command{
		Use:   "userdelegate",
		Short: "Start Airbloc user delegate daemon.",
		Long:  "Start user delegate daemon, watching and supervising user's data event.",
		Run:   start("userdelegate,warehouse", ""),
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

	// list of available APIs and services
	AvailableAPIs = map[string]node.Constructor{
		"apps":        serverapi.NewAppsAPI,
		"collections": serverapi.NewCollectionsAPI,
		"data":        serverapi.NewDataAPI,
		"dauth":       serverapi.NewDAuthAPI,
		"exchange":    serverapi.NewExchangeAPI,
		"schemas":     serverapi.NewSchemaAPI,
		"warehouse":   serverapi.NewWarehouseAPI,

		"server.accounts":       serverapi.NewAccountsAPI,
		"userdelegate.accounts": userdelegateapi.NewAccountAPI,
	}
	AvailableServices = map[string]node.ServiceConstructor{
		"api":          node.NewAPIService,
		"warehouse":    warehouseservice.New,
		"userdelegate": userdelegate.NewService,
	}
)

func init() {
	cobra.OnInitialize(loadConfig)
	rflags := rootCmd.PersistentFlags()

	rflags.StringVarP(&rootFlags.dataDir, "datadir", "d", "~/.airbloc", "Data directory")
	rflags.StringVarP(&rootFlags.configPath, "config", "c", "$DATADIR/config.yml", "Config file")
	rflags.StringVarP(&rootFlags.keyPath, "keystore", "k", "", "Keystore file for node (default is $DATADIR/private.key)")
	rflags.StringVar(&rootFlags.private, "private", "", "Raw 32-byte private key with 0x prefix (Not Recommended)")

	rflags.StringVar(&config.Blockchain.Endpoint, "ethereum", config.Blockchain.Endpoint, "Ethereum RPC endpoint")
	rflags.StringVar(&config.MetaDB.MongoDBEndpoint, "metadb", config.MetaDB.MongoDBEndpoint, "Metadatabase endpoint")
	rflags.StringSliceVar(&config.P2P.BootNodes, "bootnodes", config.P2P.BootNodes, "Bootstrap Node multiaddr for P2P")

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
		userDelegateCmd,
		versionCmd,
	)
}

func loadConfig() {
	args := strings.Join(os.Args, " ")
	if strings.HasSuffix(args, "help") || strings.HasSuffix(args, "version") {
		return
	}

	// get data directory
	dataDir, err := homedir.Expand(rootFlags.dataDir)
	if err != nil {
		log.Error("Error: failed to resolve data directory {}", err, rootFlags.dataDir)
	}
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Error("Error: failed to create data directory {}", err, rootFlags.dataDir)
	}

	configPath := rootFlags.configPath
	configPath = strings.Replace(configPath, "$DATADIR", dataDir, 1)
	if err := configor.Load(config, configPath); err != nil {
		log.Error("Error: failed to load config from {}", err, configPath)
		os.Exit(1)
	}

	// override key path
	if rootFlags.keyPath != "" {
		config.PrivateKeyPath = rootFlags.keyPath
	} else {
		keyPath := strings.Replace(rootFlags.keyPath, "$DATADIR", dataDir, 1)
		config.PrivateKeyPath = keyPath
	}

	// setup loggers
	if rootFlags.verbose {
		rootFlags.logLevel = "*"
	}
	logger2.Setup(os.Stdout, rootFlags.logLevel, rootFlags.logFilter)
	logger.SetLogger(logger.NewStandardOutput(os.Stdout, rootFlags.logLevel, rootFlags.logFilter))
	log2.SetOutput(os.Stderr)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error("Error: %v", err)
		os.Exit(1)
	}
}

func start(serviceNames string, apiNames string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		nodeKey := loadNodeKey()
		backend, err := node.NewAirblocBackend(nodeKey, config)
		if err != nil {
			log.Error("Error: init", err)
			os.Exit(1)
		}
		defer backend.Stop()

		// attach services using config
		serviceNames := strings.Split(serviceNames, ",")
		registerServices(backend, serviceNames)

		if len(apiNames) != 0 {
			// attach APIs using config
			apiNames := strings.Split(apiNames, ",")
			registerApis(backend, apiNames)
		}

		if err := backend.Start(); err != nil {
			log.Error("Error: failed to start airbloc", err)
			os.Exit(1)
		}
	}
}

func loadNodeKey() *key.Key {
	if rootFlags.private != "" {
		// load from command-line argument
		if len(rootFlags.private) != 66 || !strings.HasPrefix(rootFlags.private, "0x") {
			log.Error("Error: Invalid private key.")
			os.Exit(1)
		}
		rawKey := []byte(strings.TrimPrefix(rootFlags.private, "0x"))
		k, err := crypto.ToECDSA(rawKey)
		if err != nil {
			log.Error("Error: Invalid ECDSA key", err)
			os.Exit(1)
		}
		return key.FromECDSA(k)
	} else {
		k, err := key.Load(config.PrivateKeyPath)
		if err != nil {
			log.Error("Error: failed to load private key from the given path", err)
			os.Exit(1)
		}
		return k
	}
}

func registerServices(backend node.Backend, serviceNames []string) {
	for _, name := range serviceNames {
		serviceConstructor, exists := AvailableServices[name]
		if !exists {
			log.Error("Error: service {} does not exist.", name)
			os.Exit(1)
		}

		service, err := serviceConstructor(backend)
		if err != nil {
			log.Error("Error: failed to create service {}", err, name)
			os.Exit(1)
		}
		backend.AttachService(name, service)
	}
}

func registerApis(backend node.Backend, apiNames []string) {
	apiService, ok := backend.GetService("api").(*node.APIService)
	if !ok {
		log.Error("Error: API service is not registered.")
		os.Exit(1)
	}

	for _, name := range apiNames {
		apiConstructor, exists := AvailableAPIs[name]
		if !exists {
			log.Error("Error: API %s does not exist.", name)
			os.Exit(1)
		}

		api, err := apiConstructor(backend)
		if err != nil {
			log.Error("Error: failed to create API {}", err, name)
			os.Exit(1)
		}
		api.AttachToAPI(apiService)
	}
}
