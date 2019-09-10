package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// ErrNoAccount is error implementation of "no account"
type ErrNoAccount struct{}

func (err ErrNoAccount) Error() string {
	return "no account"
}

// accountsManager is contract wrapper struct
type accountsManager struct {
	AccountsFilterer
	contract IAccountsContract
	client   blockchain.TxClient
	log      *logger.Logger
}

// Address is getter method of Accounts.address
func (manager *accountsManager) Address() ethCommon.Address {
	return manager.contract.Address()
}

// TxHash is getter method of Accounts.txHash
func (manager *accountsManager) TxHash() ethCommon.Hash {
	return manager.contract.TxHash()
}

// CreatedAt is getter method of Accounts.createdAt
func (manager *accountsManager) CreatedAt() *big.Int {
	return manager.contract.CreatedAt()
}

// NewAccountsManager makes new *accountsManager struct
func NewAccountsManager(client blockchain.TxClient) IAccountsManager {
	contract := NewAccountsContract(client)
	return &accountsManager{
		AccountsFilterer: contract.Filterer(),
		contract:         contract,
		client:           client,
		log:              logger.New("accounts"),
	}
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (manager *accountsManager) Create(ctx context.Context) (types.ID, error) {
	receipt, err := manager.contract.Create(ctx)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseSignUpFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Account created.", logger.Attrs{
		"account_id": evt.AccountId.Hex(),
		"owner":      evt.Owner.Hex(),
	})
	return evt.AccountId, nil
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns()
func (manager *accountsManager) CreateTemporary(ctx context.Context, identityHash ethCommon.Hash) (types.ID, error) {
	receipt, err := manager.contract.CreateTemporary(ctx, identityHash)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseTemporaryCreatedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary account created.", logger.Attrs{
		"account_id": evt.AccountId.Hex(),
		"proxy":      evt.Proxy.Hex(),
	})
	return evt.AccountId, nil
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (manager *accountsManager) UnlockTemporary(
	ctx context.Context,
	identityPreimage ethCommon.Hash,
	newOwner ethCommon.Address,
	passwordSignature []byte,
) error {
	receipt, err := manager.contract.UnlockTemporary(ctx, identityPreimage, newOwner, passwordSignature)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	evt, err := manager.contract.ParseUnlockedFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse a event from the receipt")
	}

	manager.log.Info("Temporary account unlocked.", logger.Attrs{
		"account_id": evt.AccountId.Hex(),
		"new_owner":  evt.NewOwner.Hex(),
	})
	return err
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (manager *accountsManager) SetController(ctx context.Context, controller ethCommon.Address) error {
	_, err := manager.contract.SetController(ctx, controller)
	if err != nil {
		return errors.Wrap(err, "failed to transact")
	}

	manager.log.Info("Controller changed.", logger.Attrs{"controller": controller.Hex()})
	return nil
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (manager *accountsManager) GetAccount(accountId types.ID) (types.Account, error) {
	return manager.contract.GetAccount(accountId)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
func (manager *accountsManager) GetAccountId(owner ethCommon.Address) (types.ID, error) {
	if owner == (ethCommon.Address{}) {
		return manager.contract.GetAccountId(manager.client.Account().From)
	} else {
		return manager.contract.GetAccountId(owner)
	}
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (manager *accountsManager) GetAccountIdFromSignature(messageHash ethCommon.Hash, signature []byte) (types.ID, error) {
	return manager.contract.GetAccountIdFromSignature(messageHash, signature)
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts(bytes8 ) constant returns(address owner, uint8 status, address controller, address passwordProof)
func (manager *accountsManager) Accounts(accountId types.ID) (types.Account, error) {
	return manager.contract.Accounts(accountId)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(bytes8 accountId) constant returns(bool)
func (manager *accountsManager) Exists(accountId types.ID) (bool, error) {
	return manager.contract.Exists(accountId)
}

// IdentityHashToAccount is a free data retrieval call binding the contract method 0x17aba2d3.
//
// Solidity: function identityHashToAccount(bytes32 ) constant returns(bytes8)
func (manager *accountsManager) IdentityHashToAccount(identityHash ethCommon.Hash) (types.ID, error) {
	return manager.contract.IdentityHashToAccount(identityHash)
}

// IsControllerOf is a free data retrieval call binding the contract method 0xa83038e7.
//
// Solidity: function isControllerOf(address sender, bytes8 accountId) constant returns(bool)
func (manager *accountsManager) IsControllerOf(sender ethCommon.Address, accountId types.ID) (bool, error) {
	return manager.contract.IsControllerOf(sender, accountId)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(bytes8 accountId) constant returns(bool)
func (manager *accountsManager) IsTemporary(accountId types.ID) (bool, error) {
	return manager.contract.IsTemporary(accountId)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (manager *accountsManager) NumberOfAccounts() (*big.Int, error) {
	return manager.contract.NumberOfAccounts()
}
