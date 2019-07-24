package exchange

import (
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
)

// Manager is contract wrapper struct
type Manager struct {
	client   blockchain.TxClient
	contract *adapter.Exchange
}

// NewManager makes new *Manager struct
func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.Exchange{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.Exchange),
	}
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (manager *Manager) Prepare(
	ctx context.Context,
	provider string,
	consumer ethCommon.Address,
	escrow ethCommon.Address,
	escrowSign [4]byte,
	escrowArgs []byte,
	dataIds ...[20]byte,
) (types.ID, error) {
	var err error
	var ids [][20]byte
	// if length of dataIds exceeds 20,
	// this won't put dataIds when Prepare() calls. and makes array ids keeps nil state
	if len(dataIds) < 20 {
		ids = dataIds
	}
	tx, err := manager.contract.Prepare(
		manager.client.Account(),
		provider, consumer,
		escrow, escrowSign, escrowArgs,
		ids,
	)
	if err != nil {
		return types.ID{}, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return types.ID{}, err
	}

	event, err := manager.contract.ParseOfferPreparedFromReceipt(receipt)
	if err != nil {
		return types.ID{}, err
	}

	offerId := event.OfferId

	// then, splits ids into chunks which maximum length is 20.
	// and adds in offer struct one by one.
	if ids == nil {
		l := len(dataIds)
		for i := 0; i < l; i += 20 {
			start := i
			end := i + 20
			if end > l {
				end = l
			}

			err := manager.AddDataIds(ctx, offerId, dataIds[start:end])
			if err != nil {
				return offerId, err
			}
		}
	}
	return offerId, err
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (manager *Manager) AddDataIds(ctx context.Context, offerId types.ID, dataIds [][20]byte) error {
	tx, err := manager.contract.AddDataIds(manager.client.Account(), offerId, dataIds)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	return err
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (manager *Manager) Order(ctx context.Context, offerId types.ID) error {
	tx, err := manager.contract.Order(manager.client.Account(), offerId)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	_, err = manager.contract.ParseOfferPresentedFromReceipt(receipt)
	return err
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (manager *Manager) Cancel(ctx context.Context, offerId types.ID) error {
	tx, err := manager.contract.Cancel(manager.client.Account(), offerId)
	if err != nil {
		return err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	_, err = manager.contract.ParseOfferCanceledFromReceipt(receipt)
	return err
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (manager *Manager) GetOffer(offerId types.ID) (types.Offer, error) {
	return manager.contract.GetOffer(nil, offerId)
}
