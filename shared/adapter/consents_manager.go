package adapter

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/pkg/errors"
)

type consentsManager struct {
	IConsentsContract
	log *logger.Logger
}

// NewConsentsManager makes new *consentsManager struct
func NewConsentsManager(client blockchain.TxClient) IConsentsManager {
	return &consentsManager{
		IConsentsContract: client.GetContract(&ConsentsContract{}).(*ConsentsContract),
		log:               logger.New("consents"),
	}
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (manager *consentsManager) Consent(ctx context.Context, appName string, consentData types.ConsentData) error {
	receipt, err := manager.IConsentsContract.Consent(ctx, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to one data.", logger.Attrs{
		"app-name":   evt.AppName,
		"data-type":  evt.DataType,
		"account-id": evt.UserId.Hex(),
		"allowed":    evt.Allowed,
	})
	return nil
}

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, (uint8,string,bool)[] consentData) returns()
func (manager *consentsManager) ConsentMany(
	ctx context.Context,
	appName string,
	consentData []types.ConsentData,
) error {
	receipt, err := manager.IConsentsContract.ConsentMany(ctx, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to many data.", logger.Attrs{
		"app-name":   evt.AppName,
		"account-id": evt.UserId.Hex(),
		"data-count": len(receipt.Logs),
	})
	return nil
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf92519d8.
//
// Solidity: function consentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed) returns()
func (manager *consentsManager) ConsentByController(
	ctx context.Context,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
) error {
	receipt, err := manager.IConsentsContract.ConsentByController(ctx, userId, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to one data by controller.", logger.Attrs{
		"app-name":   evt.AppName,
		"data-type":  evt.DataType,
		"account-id": evt.UserId.Hex(),
		"allowed":    evt.Allowed,
	})
	return nil
}

// ConsentManyByController is a paid mutator transaction binding the contract method 0xae6d5034.
//
// Solidity: function consentManyByController(bytes8 userId, string appName, (uint8,string,bool)[] consentData) returns()
func (manager *consentsManager) ConsentManyByController(
	ctx context.Context,
	userId types.ID,
	appName string,
	consentData []types.ConsentData,
) error {
	receipt, err := manager.IConsentsContract.ConsentManyByController(ctx, userId, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to many data by controller.", logger.Attrs{
		"app-name":   evt.AppName,
		"account-id": evt.UserId.Hex(),
		"data-count": len(receipt.Logs),
	})
	return nil
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xedf2ef20.
//
// Solidity: function modifyConsentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed, bytes passwordSignature) returns()
func (manager *consentsManager) ModifyConsentByController(
	ctx context.Context,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
	passwordSignature []byte,
) error {
	receipt, err := manager.IConsentsContract.ModifyConsentByController(ctx, userId, appName, consentData, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consent modified by controller.", logger.Attrs{
		"app-name":   evt.AppName,
		"data-type":  evt.DataType,
		"account-id": evt.UserId.Hex(),
		"allowed":    evt.Allowed,
	})
	return nil
}

// ModifyConsentManyByController is a paid mutator transaction binding the contract method 0xe031b1cf.
//
// Solidity: function modifyConsentManyByController(bytes8 userId, string appName, (uint8,string,bool)[] consentData, bytes passwordSignature) returns()
func (manager *consentsManager) ModifyConsentManyByController(
	ctx context.Context,
	userId types.ID,
	appName string,
	consentData []types.ConsentData,
	passwordSignature []byte,
) error {
	receipt, err := manager.IConsentsContract.ModifyConsentManyByController(ctx, userId, appName, consentData, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consent modified by controller.", logger.Attrs{
		"app-name":   evt.AppName,
		"data-type":  evt.DataType,
		"account-id": evt.UserId.Hex(),
		"data-count": len(receipt.Logs),
	})
	return nil
}
