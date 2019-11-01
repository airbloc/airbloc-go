package blockchain

import (
	"fmt"
	"reflect"

	"github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/bind/managers"
	"github.com/airbloc/airbloc-go/bind/wrappers"
	"github.com/airbloc/logger"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/pkg/errors"
)

type constructorData struct {
	manager  func(bind.ContractBackend, interface{}) interface{}
	contract func(bind.Deployment, bind.ContractBackend) interface{}
}

var (
	constructors = map[string]constructorData{
		"Accounts": {
			contract: wrappers.NewAccountsContract,
			manager:  managers.NewAccountsManager,
		},
		"AppRegistry": {
			contract: wrappers.NewAppRegistryContract,
			manager:  managers.NewAppRegistryManager,
		},
		"Consents": {
			contract: wrappers.NewConsentsContract,
			manager:  managers.NewConsentsManager,
		},
		"ControllerRegistry": {
			contract: wrappers.NewControllerRegistryContract,
			manager:  managers.NewControllerRegistryManager,
		},
		"DataTypeRegistry": {
			contract: wrappers.NewDataTypeRegistryContract,
			manager:  managers.NewDataTypeRegistryManager,
		},
		"Exchange": {
			contract: wrappers.NewExchangeContract,
			manager:  managers.NewExchangeManager,
		},
	}
)

type ContractManager struct {
	Accounts           managers.IAccountsManager
	AppRegistry        managers.IAppRegistryManager
	Consents           managers.IConsentsManager
	ControllerRegistry managers.IControllerRegistryManager
	DataTypeRegistry   managers.IDataTypeRegistryManager
	Exchange           managers.IExchangeManager

	addrToName      map[common.Address]string
	addrToSelectors map[common.Address]map[[4]byte]string
}

func NewContractManager(client bind.ContractBackend, deployments bind.Deployments) (ContractManager, error) {
	m := ContractManager{
		addrToName:      make(map[common.Address]string),
		addrToSelectors: make(map[common.Address]map[[4]byte]string),
	}
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(&m)

	for i := 0; i < t.NumField(); i++ {
		field := t.FieldByIndex([]int{i})

		name := field.Name

		constructor, exist := constructors[name]
		if !exist {
			return ContractManager{}, errors.Errorf("constructor for %+v does not exist", name)
		}

		deployment := bind.Deployment{}
		if deployments != nil {
			if d, ok := deployments.Get(name); ok {
				deployment = d
			}
		}

		contract := constructor.contract(deployment, client)
		if contractBase, ok := contract.(bind.ContractBase); ok {
			manager := constructor.manager(client, contract)
			if reflect.TypeOf(manager).Implements(field.Type) {
				v.Elem().FieldByName(name).Set(reflect.ValueOf(manager))
				m.addrToName[contractBase.Address()] = name
				m.addrToSelectors[contractBase.Address()] = contractBase.GetSignatures()
			}
		}
	}

	return m, nil
}

func (cm ContractManager) getSignatureFromTxData(addr common.Address, txdata []byte) (string, error) {
	if txdata == nil {
		return "", errors.New("nil txdata")
	}
	if len(txdata) < 4 {
		return "", errors.New("txdata too short")
	}

	var sign [4]byte
	copy(sign[:], txdata[:4])

	if signHex, registered := cm.addrToSelectors[addr][sign]; registered {
		return signHex, nil
	} else {
		return hexutil.Encode(sign[:]), nil
	}
}

func (cm ContractManager) GetTransactionDetails(tx *types.Transaction) (string, logger.Attrs, error) {
	addr := tx.To()
	if addr == nil {
		return "", nil, errors.Wrap(errors.New("nil address"), "getting destination address from tx")
	}

	var methodName string
	if name, registered := cm.addrToName[*addr]; registered {
		signature, err := cm.getSignatureFromTxData(*addr, tx.Data())
		if err != nil {
			return "", nil, errors.Wrap(err, "getting signature from txdata")
		}
		methodName = fmt.Sprintf("%s.%s", name, signature)
	} else {
		methodName = addr.Hex()
	}

	attrs := logger.Attrs{
		"txid": tx.Hash().Hex(),
		"gas":  tx.Gas(),
	}
	if tx.Value() != nil {
		attrs["value"] = fmt.Sprintf("%s klay", WeiToKlay(tx.Value()).Text('e', 2))
	}
	return methodName, attrs, nil
}
