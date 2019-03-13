package exchange

import (
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
	offeree, escrowAddr ethCommon.Address,
	escrowSign [4]byte, escrowArgs []byte,
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
		offeree, escrowAddr,
		escrowSign, escrowArgs,
		ids,
	)
	if err != nil {
		return types.ID{}, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return types.ID{}, err
	}
	offerId := types.BytesToID(receipt.Logs[0].Topics[1].Bytes())

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
	if err != nil {
		return err
	}
	return nil
}

// manager.sol
func (manager *Manager) Order(ctx context.Context, offerId types.ID) error {
	tx, err := manager.contract.Order(manager.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}

func (manager *Manager) Settle(ctx context.Context, offerId types.ID) (*adapter.ExchangeReceipt, error) {
	manager.client.Account().GasLimit = 6000000
	tx, err := manager.contract.Settle(manager.client.Account(), offerId)
	if err != nil {
		return nil, err
	}

	receipt, err := manager.client.WaitMined(ctx, tx)
	if err != nil {
		return nil, err
	}

	evt, err := manager.contract.ParseReceiptFromReceipt(receipt)
	if err != nil {
		return nil, err
	}
	// FIXME: right zero padding
	evt.OfferId = offerId

	return evt, nil
}

func (manager *Manager) Reject(ctx context.Context, offerId types.ID) error {
	tx, err := manager.contract.Reject(manager.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = manager.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}

func (manager *Manager) GetOfferCompact(offerId types.ID) (*OfferCompact, error) {
	from, to, escrow, err := manager.contract.GetOfferCompact(nil, offerId)
	if err != nil {
		return nil, err
	}

	return &OfferCompact{
		From:   from,
		To:     to,
		Escrow: escrow,
	}, nil
}

func (manager *Manager) GetOffer(offerId types.ID) (*Offer, error) {
	from, to, dataIds, addr, sign, args, status, err := manager.contract.GetOffer(nil, offerId)
	if err != nil {
		return nil, err
	}

	escrow := &Escrow{
		Addr: addr,
		Sign: sign,
		Args: args,
	}

	offer := &Offer{
		From:    from,
		To:      to,
		DataIds: dataIds,
		Escrow:  escrow,
		Status:  status,
	}
	return offer, nil
}

func (manager *Manager) GetReceiptsByOfferor(offeror ethCommon.Address) ([][8]byte, error) {
	return manager.contract.GetReceiptsByOfferor(nil, offeror)
}

func (manager *Manager) GetReceiptsByOfferee(offeree ethCommon.Address) ([][8]byte, error) {
	return manager.contract.GetReceiptsByOfferee(nil, offeree)
}

func (manager *Manager) GetReceiptsByEscrow(escrow ethCommon.Address) ([][8]byte, error) {
	return manager.contract.GetReceiptsByEscrow(nil, escrow)
}
