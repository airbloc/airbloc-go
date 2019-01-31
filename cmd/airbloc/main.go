package main

import (
	"github.com/airbloc/airbloc-go/warehouse"
	"log"
	"os"
	"strings"

	logger2 "github.com/airbloc/airbloc-go/logger"
	"github.com/airbloc/airbloc-go/userdelegate"
	"github.com/azer/logger"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"

	"github.com/airbloc/airbloc-go/node/userdelegateapi"

	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/node/serverapi"
	"github.com/jinzhu/configor"
)

var (
	mainLogger = logger.New("airbloc")

	// list of CLI commands and flags
	commands = []cli.Command{
		{
			Name:    "userdelegate",
			Aliases: []string{"ud"},
			Usage:   "Launch a user delegate daemon",
			Action:  start("userdelegate,warehouse", ""),
			Flags:   flags,
		},
	}
	flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
			Value: "config.yml",
		},
		cli.StringFlag{
			Name:  "loglevel",
			Usage: "Log output verbosity [MUTE|INFO|TIMER|*]",
			Value: "*",
		},
		cli.StringFlag{
			Name:  "logfilter",
			Usage: "Filter logs from specific packages (e.g. warehouse,users)",
			Value: "*",
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
		"warehouse":    warehouse.NewService,
		"userdelegate": userdelegate.NewService,
	}
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	app := cli.NewApp()
	app.Name = "airbloc"
	app.Description = "A node of Airbloc Protocol, which is decentralized data exchange protocol."
	app.Commands = commands
	app.Flags = flags
	app.Action = start("api,warehouse", "apps,server.accounts,collections,data,dauth,exchange,schemas,warehouse")

	err := app.Run(os.Args)
	if err != nil {
		mainLogger.Error("Error: %+v", err)
		os.Exit(1)
	}
}

func start(serviceNames string, apiNames string) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		logger2.Setup(os.Stdout, ctx.String("loglevel"), ctx.String("logfilter"))

		config := new(node.Config)
		if err := configor.Load(config, ctx.String("config")); err != nil {
			return errors.Wrapf(err, "failed to load config from %s", ctx.String("config"))
		}

		backend, err := node.NewAirblocBackend(config)
		if err != nil {
			return errors.Wrap(err, "failed to initialize backend")
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
		return backend.Start()
	}
}

func registerServices(backend node.Backend, serviceNames []string) {
	for _, name := range serviceNames {
		serviceConstructor, exists := AvailableServices[name]
		if !exists {
			mainLogger.Error("Error: service %s does not exist.", name)
			os.Exit(1)
		}

		service, err := serviceConstructor(backend)
		if err != nil {
			mainLogger.Error("Error: failed to create service %s: %+v", name, err)
			os.Exit(1)
		}
		backend.AttachService(name, service)
	}
}

func registerApis(backend node.Backend, apiNames []string) {
	apiService, ok := backend.GetService("api").(*node.APIService)
	if !ok {
		mainLogger.Error("Error: API service is not registered.")
		os.Exit(1)
	}

	for _, name := range apiNames {
		apiConstructor, exists := AvailableAPIs[name]
		if !exists {
			mainLogger.Error("Error: API %s does not exist.", name)
			os.Exit(1)
		}

		api, err := apiConstructor(backend)
		if err != nil {
			mainLogger.Error("Error: failed to create API %s: %+v", name, err)
			os.Exit(1)
		}
		api.AttachToAPI(apiService)
	}
}
