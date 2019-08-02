package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/pkg/errors"
)

// Manager is contract wrapper struct
type consentsManager struct {
	contract IConsentsContract
	log      *logger.Logger
}

// Address is getter method of Consents.address
func (manager *consentsManager) Address() common.Address {
	return manager.contract.Address()
}

// TxHash is getter method of Consents.txHash
func (manager *consentsManager) TxHash() common.Hash {
	return manager.contract.TxHash()
}

// CreatedAt is getter method of Consents.createdAt
func (manager *consentsManager) CreatedAt() *big.Int {
	return manager.contract.CreatedAt()
}

// NewConsentsManager makes new *consentsManager struct
func NewConsentsManager(client blockchain.TxClient) IConsentsManager {
	return &consentsManager{
		contract: NewConsentsContract(client),
		log:      logger.New("consents"),
	}
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (manager *consentsManager) Consent(ctx context.Context, appName string, action uint8, dataType string, allowed bool) error {
	receipt, err := manager.contract.Consent(ctx, appName, action, dataType, allowed)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented.", logger.Attrs{
		"app-name":   evt.AppName,
		"data-type":  evt.DataType,
		"account-id": evt.UserId.Hex(),
		"allowed":    evt.Allowed,
	})
	return err
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf92519d8.
//
// Solidity: function consentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed) returns()
func (manager *consentsManager) ConsentByController(
	ctx context.Context,
	userId types.ID,
	appName string,
	action uint8,
	dataType string,
	allowed bool,
) error {
	receipt, err := manager.contract.ConsentByController(ctx, userId, appName, action, dataType, allowed)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented by controller.", logger.Attrs{
		"app-name":   evt.AppName,
		"data-type":  evt.DataType,
		"account-id": evt.UserId.Hex(),
		"allowed":    evt.Allowed,
	})
	return err
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xedf2ef20.
//
// Solidity: function modifyConsentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed, bytes passwordSignature) returns()
func (manager *consentsManager) ModifyConsentByController(
	ctx context.Context,
	userId types.ID,
	appName string,
	action uint8,
	dataType string,
	allowed bool,
	passwordSignature []byte,
) error {
	receipt, err := manager.contract.ModifyConsentByController(ctx, userId, appName, action, dataType, allowed, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consent modified by controller.", logger.Attrs{
		"app-name":   evt.AppName,
		"data-type":  evt.DataType,
		"account-id": evt.UserId.Hex(),
		"allowed":    evt.Allowed,
	})
	return err
}

// IsAllowed is a free data retrieval call binding the contract method 0xa1d2bbf5.
//
// Solidity: function isAllowed(uint8 action, bytes8 userId, string appName, string dataType) constant returns(bool)
func (manager *consentsManager) IsAllowed(
	userId types.ID,
	appName string,
	action uint8,
	dataType string,
) (bool, error) {
	return manager.contract.IsAllowed(userId, appName, action, dataType)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x118642e1.
//
// Solidity: function isAllowedAt(uint8 action, bytes8 userId, string appName, string dataType, uint256 blockNumber) constant returns(bool)
func (manager *consentsManager) IsAllowedAt(
	userId types.ID,
	appName string,
	action uint8,
	dataType string,
	blockNumber *big.Int,
) (bool, error) {
	return manager.contract.IsAllowedAt(userId, appName, action, dataType, blockNumber)
}

// FilterConsented is a free log retrieval operation binding the contract event 0xd0bd2a4b9fcbb6eee35bf0e8d542816e1d5244740220e033fff96b0abd805fac.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes32 indexed app, string appName, string dataType, bool allowed)
func (manager consentsManager) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, app []common.Address) (*ConsentsConsentedIterator, error) {
	return manager.contract.FilterConsented(opts, action, userId, app)
}

// WatchConsented is a free log subscription operation binding the contract event 0xd0bd2a4b9fcbb6eee35bf0e8d542816e1d5244740220e033fff96b0abd805fac.
//
// Solidity: event Consented(uint8 indexed action, bytes8 indexed userId, bytes32 indexed app, string appName, string dataType, bool allowed)
func (manager consentsManager) WatchConsented(opts *bind.WatchOpts, sink chan<- *ConsentsConsented, action []uint8, userId []types.ID, app []common.Address) (event.Subscription, error) {
	return manager.contract.WatchConsented(opts, sink, action, userId, app)
}
