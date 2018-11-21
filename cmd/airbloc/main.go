package main

import (
	"fmt"
	"os"
	"strings"

	accountAPI "github.com/airbloc/airbloc-go/account/api"
	"github.com/airbloc/airbloc-go/api"
	collectionsAPI "github.com/airbloc/airbloc-go/collections/api"
	dataManageAPI "github.com/airbloc/airbloc-go/data/datamanager/api"
	exchangeAPI "github.com/airbloc/airbloc-go/exchange/api"
	schemasAPI "github.com/airbloc/airbloc-go/schemas/api"
	warehouseAPI "github.com/airbloc/airbloc-go/warehouse/api"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jinzhu/configor"
)

var (
	AvailableAPIs = map[string]api.Constructor{
		"account":     accountAPI.New,
		"collections": collectionsAPI.New,
		"data":        dataManageAPI.New,
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
	glogger.Verbosity(log.Lvl(log.LvlTrace))
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

func registerServices(backend api.Backend, serviceNames []string) {
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
		backend.AttachService(name, service)
	}
}

func registerApis(backend api.Backend, apiNames []string) {
	apiService, ok := backend.GetService("api").(*api.APIService)
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
