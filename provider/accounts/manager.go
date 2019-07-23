package accounts

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type Manager struct {
	client   blockchain.TxClient
	contract *adapter.Accounts
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.Accounts{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.Accounts),
	}
}

func (manager *Manager) Create(ctx context.Context) (types.ID, error) {
	tx, err := manager.contract.Create(manager.client.Account())
	if err != nil {
		return types.ID{}, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return types.ID{}, err
	}

	event, err := manager.contract.ParseSignUpFromReceipt(receipt)
	if err != nil {
		return types.ID{}, err
	}

	return event.AccountId, nil
}

func (manager *Manager) CreateTemporary(ctx context.Context, identityHash ethCommon.Hash) (types.ID, error) {
	tx, err := manager.contract.CreateTemporary(manager.client.Account(), identityHash)
	if err != nil {
		return types.ID{}, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return types.ID{}, err
	}

	event, err := manager.contract.ParseTemporaryCreatedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, err
	}

	return event.AccountId, nil
}

func (manager *Manager) UnlockTemporary(
	ctx context.Context,
	identityPreimage ethCommon.Hash,
	newOwner ethCommon.Address,
	passwordSignature []byte,
) error {
	tx, err := manager.contract.UnlockTemporary(manager.client.Account(), identityPreimage, newOwner, passwordSignature)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = manager.contract.ParseUnlockedFromReceipt(receipt)
	return err
}

func (manager *Manager) SetController(ctx context.Context, controller ethCommon.Address) error {
	tx, err := manager.contract.SetController(manager.client.Account(), controller)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	return err
}

func (manager *Manager) GetAccount(accountId types.ID) (types.Account, error) {
	return manager.contract.GetAccount(nil, accountId)
}

func (manager *Manager) GetAccountId(owner ethCommon.Address) (types.ID, error) {
	return manager.contract.GetAccountId(nil, owner)
}

func (manager *Manager) GetAccountIdFromSignature(messageHash ethCommon.Hash, signature []byte) (types.ID, error) {
	return manager.contract.GetAccountIdFromSignature(nil, messageHash, signature)
}

func (manager *Manager) IsTemporary(accountId types.ID) (bool, error) {
	return manager.contract.IsTemporary(nil, accountId)
}

func (manager *Manager) IsControllerOf(controller ethCommon.Address, accountId types.ID) (bool, error) {
	return manager.contract.IsControllerOf(nil, controller, accountId)
}

func (manager *Manager) Exists(accountId types.ID) (bool, error) {
	return manager.contract.Exists(nil, accountId)
}
