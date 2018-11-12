package main

import (
	"fmt"
	accountAPI "github.com/airbloc/airbloc-go/account/api"
	"github.com/airbloc/airbloc-go/api"
	collectionsAPI "github.com/airbloc/airbloc-go/collections/api"
	dataAPI "github.com/airbloc/airbloc-go/data/api"
	exchangeAPI "github.com/airbloc/airbloc-go/exchange/api"
	schemasAPI "github.com/airbloc/airbloc-go/schemas/api"
	warehouseAPI "github.com/airbloc/airbloc-go/warehouse/api"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jinzhu/configor"
	"os"
	"strings"
)

var (
	AvailableAPIs = map[string]api.Constructor{
		"account":     accountAPI.New,
		"collections": collectionsAPI.New,
		"data":        dataAPI.New,
		"exchange":    exchangeAPI.New,
		"schemas":     schemasAPI.New,
		"warehouse":   warehouseAPI.New,
	}
	AvailableServices = map[string]api.ServiceConstructor{
		"api": api.NewAPIService,
	}
)

func main() {
	config := new(api.Config)
	if err := configor.Load(config, "config.yml"); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load configurations.")
		panic(err)
	}

	// setup logger
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	glogger.Verbosity(log.Lvl(4))
	log.Root().SetHandler(glogger)

	backend, err := api.NewAirblocBackend(config)
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

func registerServices(backend *api.AirblocBackend, serviceNames []string) {
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

func registerApis(backend *api.AirblocBackend, apiNames []string) {
	apiService, ok := backend.Services["api"].(*api.APIService)
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
