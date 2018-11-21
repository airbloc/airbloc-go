package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/node/serverapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jinzhu/configor"
)

var (
	AvailableAPIs = map[string]node.Constructor{
		"collections": serverapi.NewCollectionsAPI,
		"data":        serverapi.NewDataAPI,
		"exchange":    serverapi.NewExchangeAPI,
		"schemas":     serverapi.NewSchemaAPI,
		"warehouse":   serverapi.NewWarehouseAPI,
	}
	AvailableServices = map[string]node.ServiceConstructor{
		"api": node.NewAPIService,
	}
)

func main() {
	config := new(node.Config)
	if err := configor.Load(config, "config.yml"); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load configurations.")
		panic(err)
	}

	// setup logger
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	glogger.Verbosity(log.Lvl(log.LvlTrace))
	log.Root().SetHandler(glogger)

	backend, err := node.NewAirblocBackend(config)
	if err != nil {
		log.Error("Failed to initialize Airbloc backend.")
		log.Error(err.Error())
		return
	}
	defer backend.Stop()

	// attach services using config
	serviceNames := strings.Split("api", ",")
	registerServices(backend, serviceNames)

	// attach APIs using config
	apiNames := strings.Split("account,collections,data,exchange,schemas,warehouse", ",")
	registerApis(backend, apiNames)

	backend.Start()
}

func registerServices(backend *node.AirblocBackend, serviceNames []string) {
	for _, name := range serviceNames {
		serviceConstructor, exists := AvailableServices[name]
		if !exists {
			log.Error("Service does not exist", "name", name)
			panic(nil)
		}

		service, err := serviceConstructor(backend)
		if err != nil {
			log.Error("Failed to create service", "name", name, "error", err)
			panic(err)
		}
		backend.Attach(name, service)
	}
}

func registerApis(backend *node.AirblocBackend, apiNames []string) {
	apiService, ok := backend.Services["api"].(*node.APIService)
	if !ok {
		log.Error("API service is not registered")
		panic(nil)
	}

	for _, name := range apiNames {
		apiConstructor, exists := AvailableAPIs[name]
		if !exists {
			log.Error("Service does not exist", "name", name)
			panic(nil)
		}

		api, err := apiConstructor(backend)
		if err != nil {
			log.Error("Failed to create service", "name", name, "error", err)
			panic(err)
		}
		api.AttachToAPI(apiService)
	}
}
