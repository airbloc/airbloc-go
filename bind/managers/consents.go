package managers

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	types "github.com/airbloc/airbloc-go/shared/adapter/types"
	wrappers "github.com/airbloc/airbloc-go/shared/adapter/wrappers"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source consents.go -destination ./mocks/mock_consents.go -package mocks IConsentsManager

type IConsentsManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	wrappers.IConsentsCalls

	// Transact methods
	Consent(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
		consentData types.ConsentData,
	) error

	ConsentByController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData types.ConsentData,
	) error

	ConsentMany(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
		consentData []types.ConsentData,
	) error

	ConsentManyByController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData []types.ConsentData,
	) error

	ModifyConsentByController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData types.ConsentData,
		passwordSignature []byte,
	) error

	ModifyConsentManyByController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData []types.ConsentData,
		passwordSignature []byte,
	) error

	// Event methods
	wrappers.IConsentsFilterer
	wrappers.IConsentsWatcher
}

// consentsManager is contract wrapper struct
type consentsManager struct {
	wrappers.IConsentsContract
	client ablbind.ContractBackend
	log    *logger.Logger
}

// NewConsentsManager makes new *consentsManager struct
func NewConsentsManager(client ablbind.ContractBackend, contract interface{}) interface{} {
	return &consentsManager{
		IConsentsContract: contract.(*wrappers.ConsentsContract),
		client:            client,
		log:               logger.New("consents"),
	}
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (manager *consentsManager) Consent(ctx context.Context, opts *ablbind.TransactOpts, appName string, consentData types.ConsentData) error {
	receipt, err := manager.IConsentsContract.Consent(ctx, opts, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to one data.", logger.Attrs{
		"app-name":   evt[0].AppName,
		"data-type":  evt[0].DataType,
		"account-id": evt[0].UserId.Hex(),
		"allowed":    evt[0].Allowed,
	})
	return nil
}

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, (uint8,string,bool)[] consentData) returns()
func (manager *consentsManager) ConsentMany(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
	consentData []types.ConsentData,
) error {
	receipt, err := manager.IConsentsContract.ConsentMany(ctx, opts, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to many data.", logger.Attrs{
		"app-name":   evt[0].AppName,
		"account-id": evt[0].UserId.Hex(),
		"data-count": len(evt),
	})
	return nil
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf92519d8.
//
// Solidity: function consentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed) returns()
func (manager *consentsManager) ConsentByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
) error {
	receipt, err := manager.IConsentsContract.ConsentByController(ctx, opts, userId, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to one data by controller.", logger.Attrs{
		"app-name":   evt[0].AppName,
		"data-type":  evt[0].DataType,
		"account-id": evt[0].UserId.Hex(),
		"allowed":    evt[0].Allowed,
	})
	return nil
}

// ConsentManyByController is a paid mutator transaction binding the contract method 0xae6d5034.
//
// Solidity: function consentManyByController(bytes8 userId, string appName, (uint8,string,bool)[] consentData) returns()
func (manager *consentsManager) ConsentManyByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData []types.ConsentData,
) error {
	receipt, err := manager.IConsentsContract.ConsentManyByController(ctx, opts, userId, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to many data by controller.", logger.Attrs{
		"app-name":   evt[0].AppName,
		"account-id": evt[0].UserId.Hex(),
		"data-count": len(evt),
	})
	return nil
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xedf2ef20.
//
// Solidity: function modifyConsentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed, bytes passwordSignature) returns()
func (manager *consentsManager) ModifyConsentByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
	passwordSignature []byte,
) error {
	receipt, err := manager.IConsentsContract.ModifyConsentByController(ctx, opts, userId, appName, consentData, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consent modified by controller.", logger.Attrs{
		"app-name":   evt[0].AppName,
		"data-type":  evt[0].DataType,
		"account-id": evt[0].UserId.Hex(),
		"allowed":    evt[0].Allowed,
	})
	return nil
}

// ModifyConsentManyByController is a paid mutator transaction binding the contract method 0xe031b1cf.
//
// Solidity: function modifyConsentManyByController(bytes8 userId, string appName, (uint8,string,bool)[] consentData, bytes passwordSignature) returns()
func (manager *consentsManager) ModifyConsentManyByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData []types.ConsentData,
	passwordSignature []byte,
) error {
	receipt, err := manager.IConsentsContract.ModifyConsentManyByController(ctx, opts, userId, appName, consentData, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.IConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consent modified by controller.", logger.Attrs{
		"app-name":   evt[0].AppName,
		"data-type":  evt[0].DataType,
		"account-id": evt[0].UserId.Hex(),
		"data-count": len(evt),
	})
	return nil
}
