package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
)

// ErrNoAccount is error implementation of "no account"
type ErrNoAccount struct{}

func (err ErrNoAccount) Error() string {
	return "no account"
}

// accountsManager is contract wrapper struct
type accountsManager struct {
	contract IAccountsContract
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
	return &accountsManager{
		contract: NewAccountsContract(client),
		log:      logger.New("accounts"),
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
	return manager.contract.GetAccountId(owner)
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

// FilterSignUp is a free log retrieval operation binding the contract event 0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4.
//
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
func (manager *accountsManager) FilterSignUp(opts *bind.FilterOpts, owner []ethCommon.Address) (*AccountsSignUpIterator, error) {
	return manager.contract.FilterSignUp(opts, owner)
}

// WatchSignUp is a free log subscription operation binding the contract event 0xb98ae0923087f0b489e49e611630c8accd44d415c9fcbd5d59c6511877754ec4.
//
// Solidity: event SignUp(address indexed owner, bytes8 accountId)
func (manager *accountsManager) WatchSignUp(opts *bind.WatchOpts, sink chan<- *AccountsSignUp, owner []ethCommon.Address) (event.Subscription, error) {
	return manager.contract.WatchSignUp(opts, sink, owner)
}

// FilterTemporaryCreated is a free log retrieval operation binding the contract event 0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e.
//
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
func (manager *accountsManager) FilterTemporaryCreated(opts *bind.FilterOpts, proxy []ethCommon.Address, identityHash []ethCommon.Hash) (*AccountsTemporaryCreatedIterator, error) {
	return manager.contract.FilterTemporaryCreated(opts, proxy, identityHash)
}

// WatchTemporaryCreated is a free log subscription operation binding the contract event 0x7f475d23ee7af49ec9e9b689d8eddd76ab367e3d326ba1658216174b5adbf52e.
//
// Solidity: event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId)
func (manager *accountsManager) WatchTemporaryCreated(opts *bind.WatchOpts, sink chan<- *AccountsTemporaryCreated, proxy []ethCommon.Address, identityHash []ethCommon.Hash) (event.Subscription, error) {
	return manager.contract.WatchTemporaryCreated(opts, sink, proxy, identityHash)
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0.
//
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (manager *accountsManager) FilterUnlocked(opts *bind.FilterOpts, identityHash []ethCommon.Hash, accountId []types.ID) (*AccountsUnlockedIterator, error) {
	return manager.contract.FilterUnlocked(opts, identityHash, accountId)
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x97e37defaf20fab5209164d8e3b54fdb1bd84d7ec6def1886c587be543d41bc0.
//
// Solidity: event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner)
func (manager *accountsManager) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *AccountsUnlocked, identityHash []ethCommon.Hash, accountId []types.ID) (event.Subscription, error) {
	return manager.contract.WatchUnlocked(opts, sink, identityHash, accountId)
}
