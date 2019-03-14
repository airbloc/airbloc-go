package apps

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var (
	ErrPrefix          = "apps : "
	ErrCannotCreateTx  = ErrPrefix + "failed to create tx"
	ErrCannotCall      = ErrPrefix + "call failed"
	ErrContext         = ErrPrefix + "context error"
	ErrTxReverted      = ErrPrefix + "tx reverted"
	ErrEventParseError = ErrPrefix + "failed to parse event from receipt"
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

func (apps *Manager) TransferOwnership(ctx context.Context, appId types.ID, newOwner ethCommon.Address) (bool, error) {
	tx, err := apps.contract.TransferAppOwner(apps.client.Account(), appId, newOwner)
	if err != nil {
		return false, errors.Wrap(err, ErrCannotCreateTx)
	}

	_, err = apps.client.WaitMined(ctx, tx)
	if ctx.Err() != nil {
		return false, errors.Wrap(err, ErrContext)
	}

	if err != nil {
		return false, errors.Wrap(err, ErrTxReverted)
	}
	return true, nil
}

func (apps *Manager) IsOwner(ctx context.Context, appId types.ID, owner ethCommon.Address) (bool, error) {
	result, err := apps.contract.IsOwner(nil, appId, owner)
	if err != nil {
		return false, errors.Wrap(err, ErrCannotCall)
	}
	return result, nil
}

func (apps *Manager) Register(ctx context.Context, name string) (types.ID, error) {
	tx, err := apps.contract.Register(apps.client.Account(), name)
	if err != nil {
		return types.ID{}, errors.Wrap(err, ErrCannotCreateTx)
	}

	receipt, err := apps.client.WaitMined(ctx, tx)
	if ctx.Err() != nil {
		return types.ID{}, errors.Wrap(err, ErrContext)
	}

	if err != nil {
		return types.ID{}, errors.Wrap(err, ErrTxReverted)
	}

	event, err := apps.contract.ParseRegisteredFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, ErrEventParseError)
	}
	return event.AppId, nil
}

func (apps *Manager) Unregister(ctx context.Context, appId types.ID) (bool, error) {
	tx, err := apps.contract.Unregister(apps.client.Account(), appId)
	if err != nil {
		return false, errors.Wrap(err, ErrCannotCreateTx)
	}

	_, err = apps.client.WaitMined(ctx, tx)
	if ctx.Err() != nil {
		return false, errors.Wrap(err, ErrContext)
	}

	if err != nil {
		return false, errors.Wrap(err, ErrTxReverted)
	}
	return true, nil
}
