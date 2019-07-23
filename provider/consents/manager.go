package consents

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
)

type Manager struct {
	client   blockchain.TxClient
	contract *adapter.Consents
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.Consents{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.Consents),
	}
}

func (manager *Manager) Consent(ctx context.Context, action uint8, appName, dataType string, allowed bool) error {
	tx, err := manager.contract.Consent(
		manager.client.Account(),
		action, appName, dataType, allowed,
	)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseConsentedFromReceipt(receipt)
	return err
}

func (manager *Manager) ConsentByController(
	ctx context.Context,
	action uint8,
	userId types.ID,
	appName, dataType string,
	allowed bool,
) error {
	tx, err := manager.contract.ConsentByController(
		manager.client.Account(),
		action, userId, appName, dataType, allowed,
	)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseConsentedFromReceipt(receipt)
	return err
}

func (manager *Manager) ModifyConsentByController(
	ctx context.Context,
	action uint8,
	userId types.ID,
	appName, dataType string,
	allowed bool,
	passwordSignature []byte,
) error {
	tx, err := manager.contract.ModifyConsentByController(
		manager.client.Account(),
		action, userId, appName, dataType, allowed, passwordSignature,
	)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseConsentedFromReceipt(receipt)
	return err
}

func (manager *Manager) IsAllowed(
	ctx context.Context,
	action uint8,
	userId types.ID,
	appName, dataType string,
) (bool, error) {
	return manager.contract.IsAllowed(nil, action, userId, appName, dataType)
}

func (manager *Manager) IsAllowedAt(
	ctx context.Context,
	action uint8,
	userId types.ID,
	appName, dataType string,
	blockNumber *big.Int,
) (bool, error) {
	return manager.contract.IsAllowedAt(nil, action, userId, appName, dataType, blockNumber)
}
