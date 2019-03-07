package collections

import (
	"context"
	"github.com/airbloc/airbloc-go/shared/types"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
)

type Collections struct {
	client   blockchain.TxClient
	contract *adapter.CollectionRegistry
}

func New(
	client blockchain.TxClient,
) *Collections {
	contract := client.GetContract(&adapter.CollectionRegistry{})
	return &Collections{
		client:   client,
		contract: contract.(*adapter.CollectionRegistry),
	}
}

func (s *Collections) Register(ctx context.Context, collection *Collection) (types.ID, error) {
	// damn EVM
	dataProducerRatio := big.NewFloat(collection.IncentivePolicy.DataProvider)
	dataProducerRatio.Mul(dataProducerRatio, big.NewFloat(100*params.Ether))
	solidityDataProducerRatio := new(big.Int)
	dataProducerRatio.Int(solidityDataProducerRatio)

	tx, err := s.contract.Register(
		s.client.Account(),
		collection.AppId,
		collection.Schema.Id,
		solidityDataProducerRatio,
	)
	if err != nil {
		return types.ID{}, err
	}
	receipt, err := s.client.WaitMined(context.Background(), tx)
	if err != nil {
		return types.ID{}, err
	}
	event, err := s.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return types.ID{}, err
	}
	return types.ID(event.CollectionId), nil
}

func (s *Collections) Unregister(ctx context.Context, collectionId types.ID) error {
	tx, err := s.contract.Unregister(s.client.Account(), collectionId)
	if err != nil {
		return err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	// do something with event
	_, err = s.contract.ParseUnregistrationFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to receive Unregistration event")
	}
	return nil
}

func (s *Collections) Get(id types.ID) (*Collection, error) {
	result, err := s.contract.Get(nil, id)
	if err != nil {
		return nil, err
	}

	// here's little trick converting e.g.) 35 ETH to 0.35 (35%)
	dataProviderRatioPercentage := big.NewInt(0)
	dataProviderRatioPercentage.Div(result.IncentiveRatioSelf, big.NewInt(params.Ether))
	dataProviderRatio := float64(dataProviderRatioPercentage.Int64() / 100)

	collection := NewCollection(
		result.AppId,
		result.SchemaId,
		IncentivePolicy{
			DataProvider: dataProviderRatio,
			DataOwner:    1 - dataProviderRatio,
		},
	)
	collection.Id = id
	return collection, nil
}

func (s *Collections) ListID(ctx context.Context, appId types.ID) ([]types.ID, error) {
	filterOpts := &bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	}
	events, err := s.contract.FilterRegistration(filterOpts, nil, types.IDFilter(appId))
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan Registrations in CollectionRegistry")
	}
	defer events.Close()

	var collections []types.ID
	for events.Next() {
		collectionId := types.ID(events.Event.CollectionId)
		collections = append(collections, collectionId)
	}
	if events.Error() != nil {
		return nil, errors.Wrap(events.Error(), "failed to iterate over Registrations in CollectionRegistry")
	}
	return collections, nil
}

func (s *Collections) List(ctx context.Context, appId types.ID) ([]*Collection, error) {
	collectionIds, err := s.ListID(ctx, appId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get collection IDs")
	}

	var collections []*Collection
	for _, collectionId := range collectionIds {
		collection, err := s.Get(collectionId)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get collection %s", collectionId.Hex())
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (s *Collections) Exists(id types.ID) (bool, error) {
	return s.contract.Exists(nil, id)
}

func (s *Collections) IsCollectionAllowed(id, userId types.ID) (bool, error) {
	return s.contract.IsCollectionAllowed(nil, id, userId)
}

func (s *Collections) GetContract() *adapter.CollectionRegistry {
	return s.contract
}
