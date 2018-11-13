package schemas

import (
	"context"
	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: metadb integration
type Schemas struct {
	db          metadb.Database
	client      *blockchain.Client
	contract    *adapter.SchemaRegistry
	contractABI abi.ABI
}

func New(db metadb.Database, client *blockchain.Client) *Schemas {
	return &Schemas{
		db:       db,
		client:   client,
		contract: client.Contracts.SchemaRegistry,
	}
}

func (s *Schemas) Register(ctx context.Context, name string, data map[string]interface{}) (common.Hash, error) {
	dtx, err := s.contract.Register(s.client.Account())
	if err != nil {
		return common.Hash{}, err
	}

	receipt, err := bind.WaitMined(ctx, s.client, dtx)
	if err != nil {
		return common.Hash{}, err
	}

	event := adapter.SchemaRegistryRegistered{}
	if err := s.client.GetEventFromReceipt("SchemaRegistry", "Registered", &event, receipt); err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to parse a event from receipt")
	}
	// TODO: add metadata
	return common.Hash(event.Id), nil
}

func (s *Schemas) Unregister(ctx context.Context, id common.Hash) error {
	dtx, err := s.contract.Unregister(s.client.Account(), id)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, s.client, dtx)
	if err != nil {
		return err
	}

	event := adapter.SchemaRegistryUnregistered{}
	if err := s.contractABI.Unpack(
		&event,
		"Unregistered",
		receipt.Logs[0].Data,
	); err != nil {
		return err
	}

	// TODO: retrieve by id and burn it
	return nil
}
