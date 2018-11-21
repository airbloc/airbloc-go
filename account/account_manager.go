package account

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

var (
	ErrNoAccount = errors.New("Account not found.")
)

type Manager struct {
	client   *blockchain.Client
	contract *adapter.Accounts
}

func NewManager(client *blockchain.Client) *Manager {
	return &Manager{
		client:   client,
		contract: client.Contracts.Accounts,
	}
}

func (manager *Manager) CreateTemporary(identityHash ethCommon.Hash) (ablCommon.ID, error) {
	tx, err := manager.contract.CreateTemporary(manager.client.Account(), identityHash)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := manager.client.WaitMined(context.Background(), tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event, err := manager.contract.ParseTemporaryCreatedFromReceipt(receipt)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}
	accountId := ablCommon.ID(event.AccountId)
	log.Debug("New temporary account created", "accountId", accountId.String())

	return accountId, nil
}

func (manager *Manager) CreateUsingProxy(owner ethCommon.Address, passwordSignature []byte) (ablCommon.ID, error) {
	tx, err := manager.contract.CreateUsingProxy(manager.client.Account(), owner, passwordSignature)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := manager.client.WaitMined(context.Background(), tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event, err := manager.contract.ParseSignUpFromReceipt(receipt)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}
	return ablCommon.ID(event.AccountId), err
}

func (manager *Manager) Get(accountId ablCommon.ID) (*Account, error) {
	account, err := manager.contract.Accounts(nil, accountId)
	if err != nil {
		return nil, errors.Wrap(err, "call Accounts failed")
	}

	if Status(account.Status) == StatusNone {
		return nil, ErrNoAccount
	}

	return &Account{
		ID:            accountId,
		Owner:         account.Owner,
		Status:        Status(account.Status),
		Proxy:         account.Proxy,
		PasswordProof: account.PasswordProof,
	}, nil
}

func (manager *Manager) GetByIdentity(identity string) (*Account, error) {
	// make sure that double hashing is required
	identityHash := crypto.Keccak256Hash([]byte(identity))
	identityHash = crypto.Keccak256Hash(identityHash[:])

	accountId, err := manager.contract.IdentityHashToAccount(nil, identityHash)
	if err != nil {
		return nil, errors.Wrap(err, "call IdentityHashToAccount failed")
	}
	return manager.Get(accountId)
}

func (manager *Manager) TestPassword(messageHash ethCommon.Hash, signature []byte) (bool, error) {
	accountId, err := manager.contract.GetAccountIdFromSignature(nil, messageHash, signature)
	if err != nil {
		return false, errors.Wrap(err, "call Accounts.getAccountIdFromSignature reverted")
	}
	log.Trace("Successfully tested password", "accountId", accountId)
	return true, nil
}
