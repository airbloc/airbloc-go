package collections

import (
	"context"
	"github.com/ethereum/go-ethereum/params"
	"math/big"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: localdb integration
type Collections struct {
	db       *localdb.Model
	client   *blockchain.Client
	contract *adapter.CollectionRegistry
}

func New(
	db localdb.Database,
	client *blockchain.Client,
) (*Collections, error) {

	return &Collections{
		db:       localdb.NewModel(db, "collection"),
		client:   client,
		contract: client.Contracts.CollectionRegistry,
	}, nil
}

func (s *Collections) Get(id common.Hash) (*Collection, error) {
	appId, schemaId, err := s.contract.Get(nil, id)
	if err != nil {
		return nil, err
	}

	return &Collection{
		AppId:    appId,
		SchemaId: schemaId,
	}, nil
}

func (s *Collections) Register(ctx context.Context, collection *Collection) (common.Hash, error) {
	// damn EVM
	dataProducerRatio := big.NewFloat(collection.Policy.DataProducer)
	dataProducerRatio.Mul(dataProducerRatio, big.NewFloat(params.Ether))
	solidityDataProducerRatio := new(big.Int)
	dataProducerRatio.Int(solidityDataProducerRatio)

	tx, err := s.contract.Register(
		s.client.Account(),
		[8]byte{'T', 'O', 'D', 'O'}, // TODO: AppID
		collection.SchemaId,
		solidityDataProducerRatio,
	)

	if err != nil {
		return common.Hash{}, err
	}

	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return common.Hash{}, err
	}

	var event adapter.CollectionRegistryRegistered
	if err := s.client.GetEventFromReceipt("CollectionRegistry", "Registered", &event, receipt); err != nil {
		return common.Hash{}, err
	}
	return event.ColectionId, nil
}

func (s *Collections) Unregister(ctx context.Context, collectionId common.Hash) error {
	tx, err := s.contract.Unregister(s.client.Account(), collectionId)
	if err != nil {
		return err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	// do something with event
	event := adapter.CollectionRegistryUnregistered{}
	return s.client.GetEventFromReceipt("CollectionRegistry", "Unregistered", &event, receipt)
}

func (s *Collections) Check(id common.Hash) (bool, error) {
	return s.contract.Check(nil, id)
}

func (s *Collections) CheckAllowed(id, uid common.Hash) (bool, error) {
	return s.contract.CheckAllowed(nil, id, uid)
}

func (s *Collections) Allow(ctx context.Context, account *bind.TransactOpts, id, uid common.Hash) error {
	tx, err := s.contract.Allow(account, id, uid)
	if err != nil {
		return err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	event := adapter.CollectionRegistryAllowed{}
	return s.client.GetEventFromReceipt("CollectionRegistry", "Allowed", &event, receipt)
}

func (s *Collections) Deny(ctx context.Context, account *bind.TransactOpts, id, uid common.Hash) error {
	tx, err := s.contract.Deny(account, id, uid)
	if err != nil {
		return err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	event := adapter.CollectionRegistryDenied{}
	return s.client.GetEventFromReceipt("CollectionRegistry", "Denied", &event, receipt)
}
