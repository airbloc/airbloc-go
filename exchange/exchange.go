package exchange

import (
	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Exchange struct {
	client   *blockchain.Client
	contract *adapter.Exchange
}

func New(client *blockchain.Client) (*Exchange, error) {
	return &Exchange{
		client:   client,
		contract: client.Contracts.Exchange,
	}, nil
}

// exchange.sol
func (exchange *Exchange) Order(ctx context.Context, offeror, offeree, contract ethCommon.Address) (ablCommon.ID, error) {
	tx, err := exchange.contract.Order(exchange.client.Account(), offeror, offeree, contract)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event, err := exchange.contract.ParseOfferPresentedFromReceipt(receipt)
	if err != nil {
		return ablCommon.ID{}, errors.Wrap(err, "failed to parse event from the receipt")
	}

	return ablCommon.ID(event.OfferId), err
}

func (exchange *Exchange) Settle(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := exchange.contract.Settle(exchange.client.Account(), offerId)
	if err != nil {
		return err
	}

	receipt, err := exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = exchange.contract.ParseOfferSettledFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse event from the receipt")
	}
	return err
}

func (exchange *Exchange) Reject(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := exchange.contract.Reject(exchange.client.Account(), offerId)
	if err != nil {
		return err
	}

	receipt, err := exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	_, err = exchange.contract.ParseOfferSettledFromReceipt(receipt)
	if err != nil {
		return errors.Wrap(err, "failed to parse event from the receipt")
	}
	return err
}

// other contract
func (exchange *Exchange) OpenOrder(funcInfo abi.ABI, ctx context.Context) {
}

func (exchange *Exchange) CloseOrder(ctx context.Context) {

}
