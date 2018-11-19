package account

import (
	"context"

	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type Manager struct {
	client   *blockchain.Client
	contract *adapter.Accounts
}

func NewManager(client *blockchain.Client) *Manager {
	return &Manager{
		client:   client,
		contract: client.Contracts.Accounts,
	}
}

func (manager *Manager) Create(ctx context.Context) (ablCommon.ID, error) {
	tx, err := manager.contract.Create(manager.client.Account())
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event, err := manager.contract.ParseSignUpFromReceipt(receipt)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}
	return ablCommon.ID(event.AccountId), err
}

func (manager *Manager) CreateTemporary(
	ctx context.Context,
	proxy ethCommon.Address,
) error {
	tx, err := manager.contract.CreateTemporary(manager.client.Account(), proxy)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	// TODO do something
	return nil
}

func (manager *Manager) CreateUsingProxy(
	ctx context.Context,
	owner, proxy, proof ethCommon.Address,
) (ablCommon.ID, error) {
	tx, err := manager.contract.CreateUsingProxy(manager.client.Account(), owner, proxy, proof)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event, err := manager.contract.ParseSignUpFromReceipt(receipt)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}
	return ablCommon.ID(event.AccountId), err
}
