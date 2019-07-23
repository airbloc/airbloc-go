package exchange

import (
	"math/big"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
)

type Manager struct {
	client   blockchain.TxClient
	contract *adapter.Exchange
}

func NewManager(client blockchain.TxClient) *Manager {
	contract := client.GetContract(&adapter.Exchange{})
	return &Manager{
		client:   client,
		contract: contract.(*adapter.Exchange),
	}
}

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

func (manager *Manager) AddDataIds(ctx context.Context, offerId types.ID, dataIds [][20]byte) error {
	tx, err := manager.contract.AddDataIds(manager.client.Account(), offerId, dataIds)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	return err
}

func (manager *Manager) Order(ctx context.Context, offerId types.ID) error {
	tx, err := manager.contract.Order(manager.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	return err
}

func (manager *Manager) Cancel(ctx context.Context, offerId types.ID) error {
	tx, err := manager.contract.Cancel(manager.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	return err
}

func (manager *Manager) GetOffer(offerId types.ID) (struct {
	Provider string
	Consumer common.Address
	DataIds  [][20]byte
	At       *big.Int
	Until    *big.Int
	Escrow   struct {
		Addr common.Address
		Sign [4]byte
		Args []byte
	}
	Status uint8
}, error) {
	return manager.contract.GetOffer(nil, offerId)
}
