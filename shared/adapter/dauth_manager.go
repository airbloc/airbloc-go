package adapter

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
)

type DAuthManager struct {
	consents IConsentsManager
}

func NewDAuthManager(client blockchain.TxClient) *DAuthManager {
	return &DAuthManager{
		consents: NewConsentsManager(client),
	}
}

func appNameToAddr(appName string) common.Address {
	return common.BytesToAddress(crypto.Keccak256([]byte(appName)))
}

func (manager *DAuthManager) Allow(
	ctx context.Context,
	action uint8,
	appName, dataType string,
) error {
	consentData := types.ConsentData{
		Action:   action,
		DataType: dataType,
		Allow:    1,
	}
	return manager.consents.Consent(ctx, appName, consentData)
}

func (manager *DAuthManager) AllowByController(
	ctx context.Context,
	action uint8,
	userId types.ID,
	appName, dataType string,
	passwordSignature []byte,
) error {
	iter, err := manager.consents.FilterConsented(&bind.FilterOpts{
		Context: ctx,
		Start:   manager.consents.CreatedAt().Uint64(),
	}, []uint8{action}, []types.ID{userId}, []common.Address{appNameToAddr(appName)})
	if err != nil {
		return err
	}

	consentData := types.ConsentData{
		Action:   action,
		DataType: dataType,
		Allow:    1,
	}

	if iter.Next() {
		return manager.consents.ConsentByController(ctx, userId, appName, consentData)
	}
	return manager.consents.ModifyConsentByController(ctx, userId, appName, consentData, passwordSignature)
}

func (manager *DAuthManager) Deny(
	ctx context.Context,
	action uint8,
	appName, dataType string,
) error {
	consentData := types.ConsentData{
		Action:   action,
		DataType: dataType,
		Allow:    0,
	}
	return manager.consents.Consent(ctx, appName, consentData)
}

func (manager *DAuthManager) DenyByController(ctx context.Context,
	action uint8,
	userId types.ID,
	appName, dataType string,
	passwordSignature []byte,
) error {
	iter, err := manager.consents.FilterConsented(&bind.FilterOpts{
		Context: ctx,
		Start:   manager.consents.CreatedAt().Uint64(),
	}, []uint8{action}, []types.ID{userId}, []common.Address{appNameToAddr(appName)})
	if err != nil {
		return err
	}

	consentData := types.ConsentData{
		Action:   action,
		DataType: dataType,
		Allow:    0,
	}

	if iter.Next() {
		return manager.consents.ConsentByController(ctx, userId, appName, consentData)
	}
	return manager.consents.ModifyConsentByController(ctx, userId, appName, consentData, passwordSignature)
}
