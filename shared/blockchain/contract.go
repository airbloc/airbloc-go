package blockchain

import (
	"github.com/json-iterator/go"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	// ContractList is filled by automatically generated contract binds in package `adapter`.
	ContractList = make(map[string]ContractConstructor)
)

type ContractManager struct {
	client     TxClient
	addrToName map[common.Address]string
	storage    map[reflect.Type]interface{}
}

func NewContractManager(client TxClient) *ContractManager {
	return &ContractManager{
		client:     client,
		storage:    make(map[reflect.Type]interface{}),
		addrToName: make(map[common.Address]string),
	}
}

func (cm *ContractManager) Load(path string) error {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := http.Get(path)
		if err != nil {
			return errors.Wrap(err, "failed to load deployment from url")
		}
		defer resp.Body.Close()
		return cm.load(resp.Body)
	}
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to load deployment from file")
	}
	defer f.Close()
	return cm.load(f)
}

func (cm *ContractManager) load(reader io.Reader) error {
	decoder := json.NewDecoder(reader)

	contracts := make(map[string]common.Address)
	if err := decoder.Decode(&contracts); err != nil {
		return errors.Wrap(err, "contract maanger: failed to decode json")
	}

	for name, addr := range contracts {
		contract, err := ContractList[name](addr, cm.client)
		if err != nil {
			return errors.Wrap(err, "contract manager: failed to get contract")
		}
		cm.addrToName[addr] = name
		cm.SetContract(contract)
	}
	return nil
}

func (cm *ContractManager) GetContract(c interface{}) interface{} {
	return cm.storage[reflect.ValueOf(c).Type()]
}

func (cm *ContractManager) SetContract(c interface{}) {
	cm.storage[reflect.ValueOf(c).Type()] = c
}
