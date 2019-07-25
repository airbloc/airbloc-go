package consents

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/pkg/errors"
)

// Manager is contract wrapper struct
type Manager struct {
	contract adapter.IConsentsContract
	log      *logger.Logger
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) adapter.IConsentsManager {
	return &Manager{
		contract: adapter.NewConsentsContract(client),
		log:      logger.New("consents"),
	}
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (manager *Manager) Consent(ctx context.Context, action uint8, appName, dataType string, allowed bool) error {
	receipt, err := manager.contract.Consent(ctx, action, appName, dataType, allowed)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented.", logger.Attrs{
		"app-name":   event.AppName,
		"data-type":  event.DataType,
		"account-id": event.UserId.Hex(),
		"allowed":    event.Allowed,
	})
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
	receipt, err := manager.contract.ConsentByController(ctx, action, userId, appName, dataType, allowed)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented by controller.", logger.Attrs{
		"app-name":   event.AppName,
		"data-type":  event.DataType,
		"account-id": event.UserId.Hex(),
		"allowed":    event.Allowed,
	})
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
	receipt, err := manager.contract.ModifyConsentByController(ctx, action, userId, appName, dataType, allowed, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consent modified by controller.", logger.Attrs{
		"app-name":   event.AppName,
		"data-type":  event.DataType,
		"account-id": event.UserId.Hex(),
		"allowed":    event.Allowed,
	})
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
	return manager.contract.IsAllowed(action, userId, appName, dataType)
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
	return manager.contract.IsAllowedAt(action, userId, appName, dataType, blockNumber)
}
