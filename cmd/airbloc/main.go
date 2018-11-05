package main

import (
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/exchange"
	"github.com/airbloc/airbloc-go/schemas"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/ethereum/go-ethereum/log"
	"strings"
)

var (
	AvailableAPIs = map[string]api.Constructor{
		"account":     account.NewAPI,
		"collections": collections.NewAPI,
		"data":        data.NewAPI,
		"exchange":    exchange.NewAPI,
		"schemas":     schemas.NewAPI,
		"warehouse":   warehouse.NewAPI,
	}
	AvailableServices = map[string]api.ServiceConstructor{
		"api": api.NewAPIService,
	}
)

func main() {
	config := api.DefaultConfig()

	backend, err := api.NewAirblocBackend(config)
	if err != nil {
		log.Error("Failed to initialize Airbloc backend.", "error", err)
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
