package blockchain

import (
	"io"
	"os"
	"reflect"

	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var (
	ErrContractNotFound = errors.New("contract not found")
)

var ContractList map[string]ContractConstructor

type ContractManager struct {
	client  TxClient
	storage map[reflect.Type]interface{}
}

func NewContractManager(client TxClient) *ContractManager {
	return &ContractManager{
		client:  client,
		storage: make(map[reflect.Type]interface{}),
	}
}

func (cm *ContractManager) Load(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "contract manager : failed to open file")
	}
	defer f.Close()
	return cm.load(f)
}

func (cm *ContractManager) load(reader io.Reader) error {
	decoder := json.NewDecoder(reader)

	contracts := make(map[string]common.Address)
	if err := decoder.Decode(&contracts); err != nil {
		return errors.Wrap(err, "contract maanger : failed to decode json")
	}

	for name, addr := range contracts {
		contract, err := ContractList[name](addr, cm.client)
		if err != nil {
			return errors.Wrap(err, "contract manager : failed to get contract")
		}
		cm.SetContract(contract)
	}
	return nil
}

func (cm *ContractManager) GetContract(c interface{}) (interface{}, error) {
	c, ok := cm.storage[reflect.ValueOf(c).Type()]
	if !ok {
		return nil, ErrContractNotFound
	}
	return c, nil
}

func (cm *ContractManager) SetContract(c interface{}) {
	cm.storage[reflect.ValueOf(c).Type()] = c
}
