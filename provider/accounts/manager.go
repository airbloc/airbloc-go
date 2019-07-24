package accounts

import (
	"context"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

// Manager is contract wrapper struct
type Manager struct {
	client   blockchain.TxClient
	contract *adapter.Accounts
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.Accounts{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.Accounts),
	}
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
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

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns()
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

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
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

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (manager *Manager) SetController(ctx context.Context, controller ethCommon.Address) error {
	tx, err := manager.contract.SetController(manager.client.Account(), controller)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	return err
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (manager *Manager) GetAccount(accountId types.ID) (types.Account, error) {
	return manager.contract.GetAccount(nil, accountId)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (manager *Manager) GetAccountId(owner ethCommon.Address) (types.ID, error) {
	return manager.contract.GetAccountId(nil, owner)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (manager *Manager) GetAccountIdFromSignature(messageHash ethCommon.Hash, signature []byte) (types.ID, error) {
	return manager.contract.GetAccountIdFromSignature(nil, messageHash, signature)
}
