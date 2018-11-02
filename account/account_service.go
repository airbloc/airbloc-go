package account

import (
	"strings"

	"context"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type Service struct {
	client      *blockchain.Client
	contract    *adapter.Accounts
	contractABI abi.ABI
}

func NewService(
	client *blockchain.Client,
	addr ethCommon.Address,
) (*Service, error) {
	accounts, err := adapter.NewAccounts(addr, client)
	if err != nil {
		return nil, err
	}

	rawABI := strings.NewReader(adapter.AccountsABI)
	contractABI, err := abi.JSON(rawABI)
	if err != nil {
		return nil, err
	}

	return &Service{
		client:      client,
		contract:    accounts,
		contractABI: contractABI,
	}, nil
}

func (s *Service) Create(ctx context.Context) (ablCommon.ID, error) {
	tx, err := s.contract.Create(s.client.Account.Opts)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event := adapter.AccountsSignUp{}
	if err := s.contractABI.Unpack(
		&event,
		"SignUp",
		receipt.Logs[0].Data,
	); err != nil {
		return ablCommon.ID{}, err
	}

	return ablCommon.ID(event.AccountId), err
}

func (s *Service) CreateTemporary(
	ctx context.Context,
	proxy ethCommon.Address,
) error {
	tx, err := s.contract.CreateTemporary(s.client.Account.Opts, proxy)
	if err != nil {
		return err
	}

	_, err = s.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	// TODO do something
	return nil
}

func (s *Service) CreateUsingProxy(
	ctx context.Context,
	owner, proxy, proof ethCommon.Address,
) (ablCommon.ID, error) {
	tx, err := s.contract.CreateUsingProxy(s.client.Account.Opts, owner, proxy, proof)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event := adapter.AccountsSignUp{}
	if err := s.contractABI.Unpack(
		&event,
		"SignUp",
		receipt.Logs[0].Data,
	); err != nil {
		return ablCommon.ID{}, err
	}

	return ablCommon.ID(event.AccountId), err
}
