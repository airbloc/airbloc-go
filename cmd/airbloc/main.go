package main

import (
	"fmt"
	"github.com/airbloc/airbloc-go/userdelegate"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strings"

	"github.com/airbloc/airbloc-go/node/userdelegateapi"

	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/node/serverapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jinzhu/configor"
)

var (
	commands = []cli.Command{
		{
			Name:   "userdelegate",
			Usage:  "Launch a user delegate daemon",
			Action: start("userdelegate", ""),
		},
	}
	flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
			Value: "config.yml",
		},
	}
	AvailableAPIs = map[string]node.Constructor{
		"apps":        serverapi.NewAppsAPI,
		"collections": serverapi.NewCollectionsAPI,
		"data":        serverapi.NewDataAPI,
		"exchange":    serverapi.NewExchangeAPI,
		"schemas":     serverapi.NewSchemaAPI,
		"warehouse":   serverapi.NewWarehouseAPI,
		"account":     userdelegateapi.NewAccountAPI, // TODO: it's not supposed to be in here
	}
	AvailableServices = map[string]node.ServiceConstructor{
		"api":          node.NewAPIService,
		"userdelegate": userdelegate.NewService,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "airbloc"
	app.Description = "A node of Airbloc Protocol, which is decentralized data exchange protocol."
	app.Commands = commands
	app.Flags = flags
	app.Action = start("api", "account,collections,data,exchange,schemas,warehouse")

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func start(serviceNames string, apiNames string) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		config := new(node.Config)
		if err := configor.Load(config, "config.yml"); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to load configurations.")
			return err
		}

		// setup logger
		glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
		glogger.Verbosity(log.Lvl(log.LvlTrace))
		log.Root().SetHandler(glogger)

		backend, err := node.NewAirblocBackend(config)
		if err != nil {
			log.Error("Failed to initialize Airbloc backend.")
			return err
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
			log.Error("Service does not exist", "name", name)
			panic(name)
		}

		service, err := serviceConstructor(backend)
		if err != nil {
			log.Error("Failed to create service", "name", name, "error", err)
			panic(name)
		}
		backend.AttachService(name, service)
	}
}

func registerApis(backend node.Backend, apiNames []string) {
	apiService, ok := backend.GetService("api").(*node.APIService)
	if !ok {
		log.Error("API service is not registered")
		panic(nil)
	}

	for _, name := range apiNames {
		apiConstructor, exists := AvailableAPIs[name]
		if !exists {
			log.Error("Service does not exist", "name", name)
			panic(name)
		}

		api, err := apiConstructor(backend)
		if err != nil {
			log.Error("Failed to create service", "name", name, "error", err)
			panic(name)
		}
		api.AttachToAPI(apiService)
	}
}
