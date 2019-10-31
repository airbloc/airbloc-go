package blockchain

import (
	"reflect"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/adapter/managers"
	"github.com/airbloc/airbloc-go/shared/adapter/wrappers"

	"github.com/pkg/errors"
)

type ManagerConstructor func(adapter.ContractBackend, interface{}) interface{}
type ContractConstructor func(adapter.Deployment, adapter.ContractBackend) interface{}

var (
	contractConstructors = map[string]ContractConstructor{
		"Accounts":           wrappers.NewAccountsContract,
		"AppRegistry":        wrappers.NewAppRegistryContract,
		"Consents":           wrappers.NewConsentsContract,
		"ControllerRegistry": wrappers.NewControllerRegistryContract,
		"DataTypeRegistry":   wrappers.NewDataTypeRegistryContract,
		"Exchange":           wrappers.NewExchangeContract,
	}
	managerConstructors = map[string]ManagerConstructor{
		"Accounts":           managers.NewAccountsManager,
		"AppRegistry":        managers.NewAppRegistryManager,
		"Consents":           managers.NewConsentsManager,
		"ControllerRegistry": managers.NewControllerRegistryManager,
		"DataTypeRegistry":   managers.NewDataTypeRegistryManager,
		"Exchange":           managers.NewExchangeManager,
	}
)

type ContractManager struct {
	Accounts           managers.IAccountsManager
	AppRegistry        managers.IAppRegistryManager
	Consents           managers.IConsentsManager
	ControllerRegistry managers.IControllerRegistryManager
	DataTypeRegistry   managers.IDataTypeRegistryManager
	Exchange           managers.IExchangeManager
}

func NewContractManager(client adapter.ContractBackend, deployments adapter.Deployments) (ContractManager, error) {
	m := ContractManager{}
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(&m)

	for i := 0; i < t.NumField(); i++ {
		field := t.FieldByIndex([]int{i})

		name := field.Name

		contractConstructor, ok := contractConstructors[name]
		if !ok {
			return ContractManager{}, errors.Errorf("constructor for %+v does not exist", name)
		}

		managerConstructor, ok := managerConstructors[name]
		if !ok {
			return ContractManager{}, errors.Errorf("constructor for %+v does not exist", name)
		}

		deployment := adapter.Deployment{}
		if deployments != nil {
			if d, ok := deployments.Get(name); ok {
				deployment = d
			}
		}

		instance := managerConstructor(client, contractConstructor(deployment, client))
		if reflect.TypeOf(instance).Implements(field.Type) {
			v.Elem().FieldByName(name).Set(reflect.ValueOf(instance))
		}
	}

	return m, nil
}
