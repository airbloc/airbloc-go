package exchange

import (
	"context"
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
)

type Manager struct {
	client   blockchain.TxClient
	contract *adapter.Exchange
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.Exchange{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.Exchange),
	}
}

func (manager *Manager) Settle(ctx context.Context, offerId types.ID) error {
	tx, err := manager.contract.Settle(manager.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}

func (manager *Manager) Reject(ctx context.Context, offerId types.ID) error {
	tx, err := manager.contract.Reject(manager.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}
