package provider

import (
	pAPI "github.com/airbloc/airbloc-go/provider/api"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
)

var apis = []api.Constructor{
	pAPI.NewAccountsAPI,
	pAPI.NewAppRegistryAPI,
	pAPI.NewControllerRegistryAPI,
	pAPI.NewDataAPI,
	pAPI.NewDAuthAPI,
	pAPI.NewDataTypeRegistryAPI,
	pAPI.NewExchangeAPI,
	pAPI.NewWarehouseAPI,
	pAPI.NewUserAPI,
}

type Service struct {
	*api.Service
}

func NewService(backend service.Backend) (service.Service, error) {
	svc, err := api.NewService(backend)
	if err != nil {
		return nil, err
	}

	apiSvc := svc.(*api.Service)

	for _, apiConstructor := range apis {
		apiInstance, err := apiConstructor(backend)
		if err != nil {
			return nil, err
		}

		apiInstance.AttachToAPI(apiSvc)
	}

	return &Service{Service: apiSvc}, nil
}

func (service Service) Start() error {
	return service.Service.Start()
}

func (service Service) Stop() {
	service.Service.Stop()
}
