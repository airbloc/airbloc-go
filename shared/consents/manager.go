package consents

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/pkg/errors"
)

// Manager is contract wrapper struct
type manager struct {
	contract adapter.IConsentsContract
	log      *logger.Logger
}

// NewManager makes new *manager struct
func NewManager(client blockchain.TxClient) adapter.IConsentsManager {
	return &manager{
		contract: adapter.NewConsentsContract(client),
		log:      logger.New("consents"),
	}
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (manager *manager) Consent(ctx context.Context, action uint8, appName, dataType string, allowed bool) error {
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
func (manager *manager) ConsentByController(
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
func (manager *manager) ModifyConsentByController(
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
func (manager *manager) IsAllowed(
	action uint8,
	userId types.ID,
	appName, dataType string,
) (bool, error) {
	return manager.contract.IsAllowed(action, userId, appName, dataType)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x118642e1.
//
// Solidity: function isAllowedAt(uint8 action, bytes8 userId, string appName, string dataType, uint256 blockNumber) constant returns(bool)
func (manager *manager) IsAllowedAt(
	action uint8,
	userId types.ID,
	appName, dataType string,
	blockNumber *big.Int,
) (bool, error) {
	return manager.contract.IsAllowedAt(action, userId, appName, dataType, blockNumber)
}

// FilterConsented is a free log retrieval operation binding the contract event 0xd0bd2a4b9fcbb6eee35bf0e8d542816e1d5244740220e033fff96b0abd805fac.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes32 indexed app, string appName, string dataType, bool allowed)
func (manager manager) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, app []common.Hash) (*adapter.ConsentsConsentedIterator, error) {
	return manager.contract.FilterConsented(opts, action, userId, app)
}

// WatchConsented is a free log subscription operation binding the contract event 0xd0bd2a4b9fcbb6eee35bf0e8d542816e1d5244740220e033fff96b0abd805fac.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes32 indexed app, string appName, string dataType, bool allowed)
func (manager manager) WatchConsented(opts *bind.WatchOpts, sink chan<- *adapter.ConsentsConsented, action []uint8, userId []types.ID, app []common.Hash) (event.Subscription, error) {
	return manager.contract.WatchConsented(opts, sink, action, userId, app)
}
