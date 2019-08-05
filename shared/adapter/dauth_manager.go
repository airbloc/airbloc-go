package adapter

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	action types.ConsentActionTypes,
	appName, dataType string,
) error {
	return manager.consents.Consent(ctx, appName, uint8(action), dataType, true)
}

func (manager *DAuthManager) AllowByController(
	ctx context.Context,
	action types.ConsentActionTypes,
	userId types.ID,
	appName, dataType string,
	passwordSignature []byte,
) error {
	iter, err := manager.consents.FilterConsented(&bind.FilterOpts{
		Context: ctx,
		Start:   manager.consents.CreatedAt().Uint64(),
	}, []uint8{uint8(action)}, []types.ID{userId}, []common.Address{appNameToAddr(appName)})
	if err != nil {
		return err
	}

	if iter.Next() {
		return manager.consents.ConsentByController(ctx, userId, appName, uint8(action), dataType, true)
	}
	return manager.consents.ModifyConsentByController(ctx, userId, appName, uint8(action), dataType, true, passwordSignature)
}

func (manager *DAuthManager) Deny(
	ctx context.Context,
	action types.ConsentActionTypes,
	appName, dataType string,
) error {
	return manager.consents.Consent(ctx, appName, uint8(action), dataType, false)
}

func (manager *DAuthManager) DenyByController(ctx context.Context,
	action types.ConsentActionTypes,
	userId types.ID,
	appName, dataType string,
	passwordSignature []byte,
) error {
	iter, err := manager.consents.FilterConsented(&bind.FilterOpts{
		Context: ctx,
		Start:   manager.consents.CreatedAt().Uint64(),
	}, []uint8{uint8(action)}, []types.ID{userId}, []common.Address{appNameToAddr(appName)})
	if err != nil {
		return err
	}

	if iter.Next() {
		return manager.consents.ConsentByController(ctx, userId, appName, uint8(action), dataType, false)
	}
	return manager.consents.ModifyConsentByController(ctx, userId, appName, uint8(action), dataType, false, passwordSignature)
}
