package account

import (
	"strings"

	"github.com/pkg/errors"

	"context"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type Manager struct {
	client      *blockchain.Client
	contract    *adapter.Accounts
	contractABI abi.ABI
}

func NewManager(client *blockchain.Client, addr ethCommon.Address) (*Manager, error) {
	accounts, err := adapter.NewAccounts(addr, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind to Accounts")
	}

	rawABI := strings.NewReader(adapter.AccountsABI)
	contractABI, err := abi.JSON(rawABI)
	if err != nil {
		return nil, err
	}

	return &Manager{
		client:      client,
		contract:    accounts,
		contractABI: contractABI,
	}, nil
}

func (manager *Manager) Create(ctx context.Context) (ablCommon.ID, error) {
	tx, err := manager.contract.Create(manager.client.Account())
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event := adapter.AccountsSignUp{}
	if err := manager.contractABI.Unpack(
		&event,
		"SignUp",
		receipt.Logs[0].Data,
	); err != nil {
		return ablCommon.ID{}, err
	}

	return ablCommon.ID(event.AccountId), err
}

func (manager *Manager) CreateTemporary(
	ctx context.Context,
	proxy ethCommon.Address,
) error {
	tx, err := manager.contract.CreateTemporary(manager.client.Account(), proxy)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	// TODO do something
	return nil
}

func (manager *Manager) CreateUsingProxy(
	ctx context.Context,
	owner, proxy, proof ethCommon.Address,
) (ablCommon.ID, error) {
	tx, err := manager.contract.CreateUsingProxy(manager.client.Account(), owner, proxy, proof)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event := adapter.AccountsSignUp{}
	if err := manager.contractABI.Unpack(
		&event,
		"SignUp",
		receipt.Logs[0].Data,
	); err != nil {
		return ablCommon.ID{}, err
	}

	return ablCommon.ID(event.AccountId), err
}
