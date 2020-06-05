package managers

import (
	"context"
	"fmt"
	"math/big"

	ablbind "github.com/airbloc/airbloc-go/bind"
	types "github.com/airbloc/airbloc-go/bind/types"
	"github.com/airbloc/airbloc-go/contracts"
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
		userId [8]byte,
		appName string,
		consentData types.ConsentData,
	) error

	ConsentMany(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId [8]byte,
		appName string,
		consentData []types.ConsentData,
	) error

	contracts.ConsentsEventFilterer
	contracts.ConsentsEventWatcher
}

// consentsManager is contract wrapper struct
type consentsManager struct {
	*contracts.ConsentsContract
	client ablbind.ContractBackend
	log    logger.Logger
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
	userId [8]byte,
	appName string,
	consentData types.ConsentData,
) error {
	receipt, err := manager.ConsentsContract.Consent(ctx, opts, userId, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to one data.", logger.Attrs{
		"app-name":  evt[0].AppName,
		"data-type": evt[0].DataType,
		"user-id":   fmt.Sprintf("%x", evt[0].UserId),
		"allowed":   evt[0].Allowed,
	})
	return nil
}

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, (uint8,string,bool)[] consentData) returns()
func (manager *consentsManager) ConsentMany(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId [8]byte,
	appName string,
	consentData []types.ConsentData,
) error {
	receipt, err := manager.ConsentsContract.ConsentMany(ctx, opts, userId, appName, consentData)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.ConsentsContract.ParseConsentedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Consented to many data.", logger.Attrs{
		"app-name":   evt[0].AppName,
		"user-id":    fmt.Sprintf("%x", evt[0].UserId),
		"data-count": len(evt),
	})
	return nil
}
