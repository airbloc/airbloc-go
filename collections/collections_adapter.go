package collections

import (
	"context"

	"strings"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Adapter struct {
	client      *ethclient.Client
	account     *bind.TransactOpts
	contract    *adapter.CollectionRegistry
	contractABI abi.ABI
}

func NewAdapter(
	client *ethclient.Client,
	account *bind.TransactOpts,
	addr common.Address,
) (*Adapter, error) {
	collection, err := adapter.NewCollectionRegistry(addr, client)
	if err != nil {
		return nil, err
	}

	rawABI := strings.NewReader(adapter.CollectionRegistryABI)
	contractABI, err := abi.JSON(rawABI)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		client:      client,
		account:     account,
		contract:    collection,
		contractABI: contractABI,
	}, nil
}

func (adt *Adapter) Get(id common.Hash) (*Collection, error) {
	appId, schemaId, err := adt.contract.Get(nil, id)
	if err != nil {
		return nil, err
	}

	return &Collection{
		AppId:    appId,
		SchemaId: schemaId,
	}, nil
}

func (adt *Adapter) Register(ctx context.Context, collection *Collection) (common.Hash, error) {
	tx, err := adt.contract.Register(
		adt.account,
		collection.AppId,
		collection.SchemaId,
		collection.Policy.DataProducer,
	)
	if err != nil {
		return common.Hash{}, err
	}

	receipt, err := bind.WaitMined(ctx, adt.client, tx)
	if err != nil {
		return common.Hash{}, err
	}

	event := adapter.CollectionRegistryCollectionRegistered{}
	if err := adt.contractABI.Unpack(
		&event,
		"CollectionRegistered",
		receipt.Logs[0].Data,
	); err != nil {
		return common.Hash{}, err
	}

	return event.ColectionId, nil
}

func (adt *Adapter) Unregister(ctx context.Context, collectionId common.Hash) error {
	tx, err := adt.contract.Unregister(adt.account, collectionId)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, adt.client, tx)
	if err != nil {
		return err
	}

	event := adapter.CollectionRegistryCollectionUnregistered{}
	if err := adt.contractABI.Unpack(
		&event,
		"CollectionUnregistered",
		receipt.Logs[0].Data,
	); err != nil {
		return err
	}

	return nil
}
