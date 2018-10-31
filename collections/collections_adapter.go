package collections

import (
	"context"

	"strings"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// TODO: localdb integration
type Adapter struct {
	db          *localdb.Model
	client      *ethclient.Client
	account     *bind.TransactOpts
	contract    *adapter.CollectionRegistry
	contractABI abi.ABI
}

func NewAdapter(
	db localdb.Database,
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
		db:          localdb.NewModel(db, "collection"),
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
	if receipt.Status == types.ReceiptStatusFailed {
		return common.Hash{}, errors.New("tx reverted")
	}

	event, err := adt.ParseRegisteredEvent(receipt.Logs[0].Data)
	if err != nil {
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
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("tx reverted")
	}

	// do something with event
	_, err = adt.ParseUnregsiteredEvent(receipt.Logs[0].Data)
	return err
}

func (adt *Adapter) Check(id common.Hash) (bool, error) {
	return adt.contract.Check(nil, id)
}

func (adt *Adapter) CheckAllowed(id, uid common.Hash) (bool, error) {
	return adt.contract.CheckAllowed(nil, id, uid)
}

func (adt *Adapter) Allow(ctx context.Context, account *bind.TransactOpts, id, uid common.Hash) error {
	tx, err := adt.contract.Allow(account, id, uid)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, adt.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("tx reverted")
	}

	_, err = adt.ParseAllowedEvent(receipt.Logs[0].Data)
	return err
}

func (adt *Adapter) Deny(ctx context.Context, account *bind.TransactOpts, id, uid common.Hash) error {
	tx, err := adt.contract.Deny(account, id, uid)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, adt.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("tx reverted")
	}

	_, err = adt.ParseDenideEvent(receipt.Logs[0].Data)
	return err
}
