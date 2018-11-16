package account

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
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
