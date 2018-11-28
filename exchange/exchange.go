package exchange

import (
	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Exchange struct {
	client   blockchain.TxClient
	contract *adapter.Exchange
}

func New(client blockchain.TxClient) (*Exchange, error) {
	raw, err := client.GetContract(&adapter.Exchange{})
	if err != nil {
		return nil, err
	}

	contract, ok := raw.(*adapter.Exchange)
	if !ok {
		return nil, blockchain.ErrContractNotFound
	}
	return &Exchange{
		client:   client,
		contract: contract,
	}, nil
}

func (exchange *Exchange) Prepare(
	ctx context.Context,
	offeror, offeree, escrow ethCommon.Address,
	sign [4]byte, args []byte, dataIds [][16]byte,
) (ablCommon.ID, error) {
	var err error
	var ids [][16]byte
	if len(dataIds) < 20 {
		ids = dataIds
	}
	tx, err := exchange.contract.Prepare(
		exchange.client.Account(),
		offeror, offeree, escrow,
		sign, args, ids,
	)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event, err := exchange.contract.ParseOfferPreparedFromReceipt(receipt)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to parse event from the receipt")
	}
	offerId := ablCommon.ID(event.OfferId)

	if ids == nil {
		l := len(dataIds)
		for i := 0; i < l; i += 20 {
			start := i
			end := i + 20
			if end > l {
				end = l
			}

			err := exchange.AddDataIds(ctx, offerId, dataIds[start:end])
			if err != nil {
				return offerId, err
			}
		}
	}
	return offerId, err
}

func (exchange *Exchange) AddDataIds(ctx context.Context, offerId ablCommon.ID, dataIds [][16]byte) error {
	tx, err := exchange.contract.AddDataIds(exchange.client.Account(), offerId, dataIds)
	if err != nil {
		return err
	}

	_, err = exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}

// exchange.sol
func (exchange *Exchange) Order(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := exchange.contract.Order(exchange.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}

func (exchange *Exchange) Settle(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := exchange.contract.Settle(exchange.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}

func (exchange *Exchange) Reject(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := exchange.contract.Reject(exchange.client.Account(), offerId)
	if err != nil {
		return err
	}

	_, err = exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}

func (exchange *Exchange) CloseOrder(ctx context.Context, offerId ablCommon.ID, abi abi.ABI, args ...interface{}) error {
	_, _, escrow, _, err := exchange.contract.GetOfferCompact(nil, offerId)
	if err != nil {
		return err
	}

	contract := bind.NewBoundContract(
		escrow, abi,
		exchange.client,
		exchange.client,
		exchange.client,
	)

	tx, err := contract.Transact(exchange.client.Account(), "close", args...)
	if err != nil {
		return err
	}

	_, err = exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}
