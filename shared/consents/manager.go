package consents

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
)

// Manager is contract wrapper struct
type Manager struct {
	client   blockchain.TxClient
	contract *adapter.Consents
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) adapter.ConsentsManager {
	contract := client.GetContract(&adapter.Consents{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.Consents),
	}
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
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

// ConsentByController is a paid mutator transaction binding the contract method 0xf92519d8.
//
// Solidity: function consentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed) returns()
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

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xedf2ef20.
//
// Solidity: function modifyConsentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed, bytes passwordSignature) returns()
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

// IsAllowed is a free data retrieval call binding the contract method 0xa1d2bbf5.
//
// Solidity: function isAllowed(uint8 action, bytes8 userId, string appName, string dataType) constant returns(bool)
func (manager *Manager) IsAllowed(
	action uint8,
	userId types.ID,
	appName, dataType string,
) (bool, error) {
	return manager.contract.IsAllowed(nil, action, userId, appName, dataType)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x118642e1.
//
// Solidity: function isAllowedAt(uint8 action, bytes8 userId, string appName, string dataType, uint256 blockNumber) constant returns(bool)
func (manager *Manager) IsAllowedAt(
	action uint8,
	userId types.ID,
	appName, dataType string,
	blockNumber *big.Int,
) (bool, error) {
	return manager.contract.IsAllowedAt(nil, action, userId, appName, dataType, blockNumber)
}
