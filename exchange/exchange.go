package exchange

import (
	"strings"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
)

type Exchange struct {
	client      *blockchain.Client
	contract    *adapter.Exchange
	contractABI abi.ABI
}

func New(
	client *blockchain.Client,
	addr ethCommon.Address,
) (*Exchange, error) {
	exchange, err := adapter.NewExchange(addr, client)
	if err != nil {
		return nil, err
	}

	rawABI := strings.NewReader(adapter.ExchangeABI)
	contractABI, err := abi.JSON(rawABI)
	if err != nil {
		return nil, err
	}

	return &Exchange{
		client:      client,
		contract:    exchange,
		contractABI: contractABI,
	}, nil
}

// exchange.sol
func (exchange *Exchange) Order(ctx context.Context, offeror, offeree, contract ethCommon.Address) (ablCommon.ID, error) {
	tx, err := exchange.contract.Order(s.client.Account(), offeror, offeree, contract)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event := adapter.ExchangeOfferPresented{}
	if err = exchange.contractABI.Unpack(
		&event,
		"OfferPresented",
		receipt.Logs[0].Data,
	); err != nil {
		return ablCommon.ID{}, err
	}

	return ablCommon.ID(event.OfferId), err
}

func (exchange *Exchange) Settle(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := exchange.contract.Settle(s.client.Account(), offerId)
	if err != nil {
		return err
	}

	receipt, err := exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	event := adapter.ExchangeOfferSettled{}
	if err = exchange.contractABI.Unpack(
		&event,
		"OfferSettled",
		receipt.Logs[0].Data,
	); err != nil {
		return err
	}

	return err
}

func (exchange *Exchange) Reject(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := exchange.contract.Reject(s.client.Account(), offerId)
	if err != nil {
		return err
	}

	receipt, err := exchange.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	event := adapter.ExchangeOfferSettled{}
	if err = exchange.contractABI.Unpack(
		&event,
		"OfferRejected",
		receipt.Logs[0].Data,
	); err != nil {
		return err
	}

	return err
}

// other contract
func (exchange *Exchange) OpenOrder(funcInfo abi.ABI, ctx context.Context) {
}

func (exchange *Exchange) CloseOrder(ctx context.Context) {

}
