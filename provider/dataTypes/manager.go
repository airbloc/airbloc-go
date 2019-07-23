package dataTypes

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/ethereum/go-ethereum/common"
)

type Manager struct {
	client   blockchain.TxClient
	contract *adapter.DataTypeRegistry
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.DataTypeRegistry{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.DataTypeRegistry),
	}
}

func (manager *Manager) Register(ctx context.Context, name string, schemaHash common.Hash) error {
	tx, err := manager.contract.Register(manager.client.Account(), name, schemaHash)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseRegistrationFromReceipt(receipt)
	return err
}

func (manager *Manager) Unregister(ctx context.Context, name string) error {
	tx, err := manager.contract.Unregister(manager.client.Account(), name)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseUnregistrationFromReceipt(receipt)
	return err
}

func (manager *Manager) Get(ctx context.Context, name string) (types.DataType, error) {
	return manager.contract.Get(nil, name)
}
