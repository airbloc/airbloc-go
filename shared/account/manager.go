package account

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Manager is contract wrapper struct
type Manager struct {
	contract adapter.IAccountsContract
	log      *logger.Logger
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) adapter.IAccountsManager {
	return &Manager{
		contract: adapter.NewAccountsContract(client),
		log:      logger.New("accounts"),
	}
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (manager *Manager) Create(ctx context.Context) (types.ID, error) {
	receipt, err := manager.contract.Create(ctx)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseSignUpFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Account created.", logger.Attrs{
		"account_id": event.AccountId.Hex(),
		"owner":      event.Owner.Hex(),
	})
	return event.AccountId, nil
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns()
func (manager *Manager) CreateTemporary(ctx context.Context, identityHash ethCommon.Hash) (types.ID, error) {
	receipt, err := manager.contract.CreateTemporary(ctx, identityHash)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseTemporaryCreatedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary account created.", logger.Attrs{
		"account_id": event.AccountId.Hex(),
		"proxy":      event.Proxy.Hex(),
	})
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
	receipt, err := manager.contract.UnlockTemporary(ctx, identityPreimage, newOwner, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseUnlockedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary account unlocked.", logger.Attrs{
		"account_id": event.AccountId.Hex(),
		"new_owner":  event.NewOwner.Hex(),
	})
	return err
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (manager *Manager) SetController(ctx context.Context, controller ethCommon.Address) error {
	receipt, err := manager.contract.SetController(ctx, controller)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	event, err := manager.contract.ParseControllerChangedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Controller changed.", logger.Attrs{
		"account_id":     event.AccountId.Hex(),
		"new_controller": event.NewController.Hex(),
	})
	return nil
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (manager *Manager) GetAccount(accountId types.ID) (types.Account, error) {
	return manager.contract.GetAccount(accountId)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (manager *Manager) GetAccountId(owner ethCommon.Address) (types.ID, error) {
	return manager.contract.GetAccountId(owner)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (manager *Manager) GetAccountIdFromSignature(messageHash ethCommon.Hash, signature []byte) (types.ID, error) {
	return manager.contract.GetAccountIdFromSignature(messageHash, signature)
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts(bytes8 ) constant returns(address owner, uint8 status, address controller, address passwordProof)
func (manager *Manager) Accounts(accountId types.ID) (types.Account, error) {
	return manager.contract.Accounts(accountId)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 accountId) constant returns(bool)
func (manager *Manager) Exists(accountId types.ID) (bool, error) {
	return manager.contract.Exists(accountId)
}

// IdentityHashToAccount is a free data retrieval call binding the contract method 0x17aba2d3.
//
// Solidity: function identityHashToAccount(bytes32 ) constant returns(bytes8)
func (manager *Manager) IdentityHashToAccount(identityHash ethCommon.Hash) (types.ID, error) {
	return manager.contract.IdentityHashToAccount(identityHash)
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address sender, bytes8 accountId) constant returns(bool)
func (manager *Manager) IsControllerOf(sender ethCommon.Address, accountId types.ID) (bool, error) {
	return manager.contract.IsControllerOf(sender, accountId)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 accountId) constant returns(bool)
func (manager *Manager) IsTemporary(accountId types.ID) (bool, error) {
	return manager.contract.IsTemporary(accountId)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (manager *Manager) NumberOfAccounts() (*big.Int, error) {
	return manager.contract.NumberOfAccounts()
}
