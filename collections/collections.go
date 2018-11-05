package collections

import (
	"context"

	"strings"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

// TODO: localdb integration
type Service struct {
	db          *localdb.Model
	client      *blockchain.Client
	account     *bind.TransactOpts
	contract    *adapter.CollectionRegistry
	contractABI abi.ABI
}

func NewService(
	db localdb.Database,
	client *blockchain.Client,
	account *bind.TransactOpts,
	addr common.Address,
) (*Service, error) {
	collection, err := adapter.NewCollectionRegistry(addr, client)
	if err != nil {
		return nil, err
	}

	rawABI := strings.NewReader(adapter.CollectionRegistryABI)
	contractABI, err := abi.JSON(rawABI)
	if err != nil {
		return nil, err
	}

	return &Service{
		db:          localdb.NewModel(db, "collection"),
		client:      client,
		account:     account,
		contract:    collection,
		contractABI: contractABI,
	}, nil
}

func (s *Service) Get(id common.Hash) (*Collection, error) {
	appId, schemaId, err := s.contract.Get(nil, id)
	if err != nil {
		return nil, err
	}

	return &Collection{
		AppId:    appId,
		SchemaId: schemaId,
	}, nil
}

func (s *Service) Register(ctx context.Context, collection *Collection) (common.Hash, error) {
	tx, err := s.contract.Register(
		s.account,
		collection.AppId,
		collection.SchemaId,
		collection.Policy.DataProducer,
	)
	if err != nil {
		return common.Hash{}, err
	}

	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return common.Hash{}, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return common.Hash{}, errors.New("tx reverted")
	}

	event, err := s.ParseRegisteredEvent(receipt.Logs[0].Data)
	if err != nil {
		return common.Hash{}, err
	}

	return event.ColectionId, nil
}

func (s *Service) Unregister(ctx context.Context, collectionId common.Hash) error {
	tx, err := s.contract.Unregister(s.account, collectionId)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("tx reverted")
	}

	// do something with event
	_, err = s.ParseUnregsiteredEvent(receipt.Logs[0].Data)
	return err
}

func (s *Service) Check(id common.Hash) (bool, error) {
	return s.contract.Check(nil, id)
}

func (s *Service) CheckAllowed(id, uid common.Hash) (bool, error) {
	return s.contract.CheckAllowed(nil, id, uid)
}

func (s *Service) Allow(ctx context.Context, account *bind.TransactOpts, id, uid common.Hash) error {
	tx, err := s.contract.Allow(account, id, uid)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("tx reverted")
	}

	_, err = s.ParseAllowedEvent(receipt.Logs[0].Data)
	return err
}

func (s *Service) Deny(ctx context.Context, account *bind.TransactOpts, id, uid common.Hash) error {
	tx, err := s.contract.Deny(account, id, uid)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("tx reverted")
	}

	_, err = s.ParseDenideEvent(receipt.Logs[0].Data)
	return err
}
