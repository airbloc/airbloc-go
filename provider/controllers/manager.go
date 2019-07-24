package controllers

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/ethereum/go-ethereum/common"
)

type Manager struct {
	client   blockchain.TxClient
	contract *adapter.ControllerRegistry
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.ControllerRegistry{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.ControllerRegistry),
	}
}

func (manager *Manager) Register(ctx context.Context, controllerAddr common.Address) error {
	tx, err := manager.contract.Register(manager.client.Account(), controllerAddr)
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

func (manager *Manager) Get(ctx context.Context, controllerAddr common.Address) (types.DataController, error) {
	return manager.contract.Get(nil, controllerAddr)
}
