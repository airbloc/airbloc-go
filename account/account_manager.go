package account

import (
	"context"

	"github.com/azer/logger"
	"github.com/ethereum/go-ethereum/crypto"
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
	client   blockchain.TxClient
	contract *adapter.Accounts
	log      *logger.Logger
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.Accounts{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.Accounts),
		log:      logger.New("accounts"),
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
	manager.log.Info("New temporary account created.", logger.Attrs{"id": accountId.Hex()})

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
		Delegate:      account.Delegate,
		PasswordProof: account.PasswordProof,
	}, nil
}

func (manager *Manager) GetByIdentity(identity string) (*Account, error) {
	identityHash := manager.HashIdentity(identity)
	accountId, err := manager.contract.IdentityHashToAccount(nil, identityHash)
	if err != nil {
		return nil, errors.Wrap(err, "call IdentityHashToAccount failed")
	}
	return manager.Get(accountId)
}

// HashIdentity hashes Personal-Identifiable ID (PIID) (e.g. Email).
func (manager *Manager) HashIdentity(identity string) (identityHash ethCommon.Hash) {
	// make sure that double hashing is required for commitment scheme without revaling real identity.
	// TODO: use salted hash to prevent rainbow table
	identityHash = crypto.Keccak256Hash([]byte(identity))
	identityHash = crypto.Keccak256Hash(identityHash[:])
	return
}

func (manager *Manager) TestPassword(messageHash ethCommon.Hash, signature []byte) (bool, error) {
	accountId, err := manager.contract.GetAccountIdFromSignature(nil, messageHash, signature)
	if err != nil {
		return false, errors.Wrap(err, "call Accounts.getAccountIdFromSignature reverted")
	}
	manager.log.Info("Successfully tested password.", logger.Attrs{"account": accountId})
	return true, nil
}

func (manager *Manager) IsDelegateOf(delegateAddr ethCommon.Address, accountId ablCommon.ID) (bool, error) {
	return manager.contract.IsDelegateOf(nil, delegateAddr, accountId)
}

func (manager *Manager) GetContract() *adapter.Accounts {
	return manager.contract
}
