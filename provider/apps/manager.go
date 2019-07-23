package apps

import (
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/ethereum/go-ethereum/common"
)

type Manager struct {
	client   blockchain.TxClient
	contract *adapter.AppRegistry
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.AppRegistry{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.AppRegistry),
	}
}

func (manager *Manager) Register(appName string) error {
	tx, err := manager.contract.Register(manager.client.Account(), appName)
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

func (manager *Manager) Unregister(appName string) error {
	tx, err := manager.contract.Unregister(manager.client.Account(), appName)
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

func (manager *Manager) Get(appName string) (types.App, error) {
	return manager.contract.Get(nil, appName)
}

func (manager *Manager) TransferAppOwner(appName string, newOwner common.Address) error {
	tx, err := manager.contract.TransferAppOwner(manager.client.Account(), appName, newOwner)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseAppOwnerTransferredFromReceipt(receipt)
	return err

}
