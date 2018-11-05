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

type Service struct {
	client      *blockchain.Client
	contract    *adapter.Exchange
	contractABI abi.ABI
}

func NewService(
	client *blockchain.Client,
	addr ethCommon.Address,
) (*Service, error) {
	exchange, err := adapter.NewExchange(addr, client)
	if err != nil {
		return nil, err
	}

	rawABI := strings.NewReader(adapter.ExchangeABI)
	contractABI, err := abi.JSON(rawABI)
	if err != nil {
		return nil, err
	}

	return &Service{
		client:      client,
		contract:    exchange,
		contractABI: contractABI,
	}, nil
}

// exchange.sol
func (s *Service) Order(ctx context.Context, offeror, offeree, contract ethCommon.Address) (ablCommon.ID, error) {
	tx, err := s.contract.Order(s.client.Account(), offeror, offeree, contract)
	if err != nil {
		return ablCommon.ID{}, err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return ablCommon.ID{}, err
	}

	event := adapter.ExchangeOfferPresented{}
	if err = s.contractABI.Unpack(
		&event,
		"OfferPresented",
		receipt.Logs[0].Data,
	); err != nil {
		return ablCommon.ID{}, err
	}

	return ablCommon.ID(event.OfferId), err
}

func (s *Service) Settle(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := s.contract.Settle(s.client.Account(), offerId)
	if err != nil {
		return err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	event := adapter.ExchangeOfferSettled{}
	if err = s.contractABI.Unpack(
		&event,
		"OfferSettled",
		receipt.Logs[0].Data,
	); err != nil {
		return err
	}

	return err
}

func (s *Service) Reject(ctx context.Context, offerId ablCommon.ID) error {
	tx, err := s.contract.Reject(s.client.Account(), offerId)
	if err != nil {
		return err
	}

	receipt, err := s.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	event := adapter.ExchangeOfferSettled{}
	if err = s.contractABI.Unpack(
		&event,
		"OfferRejected",
		receipt.Logs[0].Data,
	); err != nil {
		return err
	}

	return err
}

// other contract
func (s *Service) OpenOrder(funcInfo abi.ABI, ctx context.Context) {
}

func (s *Service) CloseOrder(ctx context.Context) {

}
