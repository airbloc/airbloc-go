package blockchain

import (
	"reflect"

	"github.com/pkg/errors"
)

var (
	ErrContractNotFound = errors.New("contract not found")
)

type ContractManager struct {
	storage map[reflect.Type]interface{}
}

func NewContractManager() *ContractManager {
	return &ContractManager{
		storage: make(map[reflect.Type]interface{}),
	}
}

func (cm *ContractManager) Get(c interface{}) (interface{}, error) {
	c, ok := cm.storage[reflect.ValueOf(c).Type()]
	if !ok {
		return nil, ErrContractNotFound
	}
	return c, nil
}

func (cm *ContractManager) Set(c interface{}) {
	cm.storage[reflect.ValueOf(c).Type()] = c
}
