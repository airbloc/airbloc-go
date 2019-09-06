package provider

import (
	pAPI "github.com/airbloc/airbloc-go/provider/api"
	serviceLib "github.com/airbloc/airbloc-go/shared/service"
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
	pAPI.NewConsentsAPI,
}

type service struct {
	*api.Service
}

func NewService(backend serviceLib.Backend) (serviceLib.Service, error) {
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

	return &service{Service: apiSvc}, nil
}

func (service service) Start() error {
	return service.Service.Start()
}

func (service service) Stop() {
	service.Service.Stop()
}
