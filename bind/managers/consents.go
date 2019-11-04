package managers

import (
	"context"
	"math/big"

	ablbind "github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/bind/contracts"
	types "github.com/airbloc/airbloc-go/bind/types"
	logger "github.com/airbloc/logger"
	common "github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

//go:generate mockgen -source consents.go -destination ./mocks/mock_consents.go -package mocks IConsentsManager

type ConsentsManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	contracts.ConsentsCaller

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

	contracts.ConsentsEventFilterer
	contracts.ConsentsEventWatcher
}

// consentsManager is contract wrapper struct
type consentsManager struct {
	*contracts.ConsentsContract
	client ablbind.ContractBackend
	log    *logger.Logger
}

// NewConsentsManager makes new *consentsManager struct
func NewConsentsManager(backend ablbind.ContractBackend) (ConsentsManager, error) {
	contract, err := contracts.NewConsentsContract(backend)
	if err != nil {
		return nil, err
	}

	return &consentsManager{
		ConsentsContract: contract,
		client:           backend,
		log:              logger.New("consents"),
	}, nil
}

// Consent is a paid mutator transaction binding the contract method 0xcd4dc804.
//
// Solidity: function consent(string appName, (uint8,string,bool) consentData) returns()
func (manager *consentsManager) Consent(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
	consentData types.ConsentData,
) error {
	receipt, err := manager.ConsentsContract.Consent(ctx, opts, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ConsentsContract.ParseConsentedFromReceipt(receipt)
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

// ConsentByController is a paid mutator transaction binding the contract method 0xf573f89a.
//
// Solidity: function consentByController(bytes8 userId, string appName, (uint8,string,bool) consentData) returns()
func (manager *consentsManager) ConsentByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
) error {
	receipt, err := manager.ConsentsContract.ConsentByController(ctx, opts, userId, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ConsentsContract.ParseConsentedFromReceipt(receipt)
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

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, (uint8,string,bool)[] consentData) returns()
func (manager *consentsManager) ConsentMany(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
	consentData []types.ConsentData,
) error {
	receipt, err := manager.ConsentsContract.ConsentMany(ctx, opts, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ConsentsContract.ParseConsentedFromReceipt(receipt)
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
	receipt, err := manager.ConsentsContract.ConsentManyByController(ctx, opts, userId, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ConsentsContract.ParseConsentedFromReceipt(receipt)
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

// ModifyConsentByController is a paid mutator transaction binding the contract method 0x0bfec389.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, (uint8,string,bool) consentData, bytes passwordSignature) returns()
func (manager *consentsManager) ModifyConsentByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
	passwordSignature []byte,
) error {
	receipt, err := manager.ConsentsContract.ModifyConsentByController(ctx, opts, userId, appName, consentData, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ConsentsContract.ParseConsentedFromReceipt(receipt)
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
	receipt, err := manager.ConsentsContract.ModifyConsentManyByController(ctx, opts, userId, appName, consentData, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ConsentsContract.ParseConsentedFromReceipt(receipt)
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
