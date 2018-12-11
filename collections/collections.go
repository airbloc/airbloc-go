package collections

import (
	"context"
	"github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"math/big"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
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

func (s *Collections) Register(ctx context.Context, collection *Collection) (common.ID, error) {
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
		return common.ID{}, err
	}
	receipt, err := s.client.WaitMined(context.Background(), tx)
	if err != nil {
		return common.ID{}, err
	}
	event, err := s.contract.ParseRegistrationFromReceipt(receipt)
	if err != nil {
		return common.ID{}, err
	}
	return common.ID(event.CollectionId), nil
}

func (s *Collections) Unregister(ctx context.Context, collectionId common.ID) error {
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

func (s *Collections) Get(id common.ID) (*Collection, error) {
	result, err := s.contract.Get(nil, id)
	if err != nil {
		return nil, err
	}

	// here's little trick converting e.g.) 35 ETH to 0.35 (35%)
	dataProviderRatioPercentage := big.NewInt(0)
	dataProviderRatioPercentage.Div(result.IncentiveRatioSelf, big.NewInt(params.Ether))
	dataProviderRatio := float64(dataProviderRatioPercentage.Int64() / 100)

	return NewCollection(
		result.AppId,
		result.SchemaId,
		IncentivePolicy{
			DataProvider: dataProviderRatio,
			DataOwner:    1 - dataProviderRatio,
		},
	), nil
}

func (s *Collections) Exists(id common.ID) (bool, error) {
	return s.contract.Exists(nil, id)
}

func (s *Collections) IsCollectionAllowed(id, userId common.ID) (bool, error) {
	return s.contract.IsCollectionAllowed(nil, id, userId)
}
