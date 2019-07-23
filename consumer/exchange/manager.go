package exchange

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/ethereum/go-ethereum/common"
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

func (manager *Manager) GetOffer(offerId types.ID) (struct {
	Provider string
	Consumer common.Address
	DataIds  [][20]byte
	At       *big.Int
	Until    *big.Int
	Escrow   struct {
		Addr common.Address
		Sign [4]byte
		Args []byte
	}
	Status uint8
}, error) {
	return manager.contract.GetOffer(nil, offerId)
}
