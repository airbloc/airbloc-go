package exchange

import (
	"strings"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

type Service struct {
	db          *localdb.Model
	client      *ethclient.Client
	account     *bind.TransactOpts
	contract    *adapter.Exchange
	contractABI abi.ABI
}

func NewService(
	db localdb.Database,
	client *ethclient.Client,
	account *bind.TransactOpts,
	addr common.Address,
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
		db:          localdb.NewModel(db, "collection"),
		client:      client,
		account:     account,
		contract:    exchange,
		contractABI: contractABI,
	}, nil
}

// exchange.sol
func (s *Service) Order(ctx context.Context) {

}

func (s *Service) Settle(ctx context.Context) {

}

func (s *Service) Reject(ctx context.Context) {

}

// other contract
func (s *Service) OpenOrder(ctx context.Context) {

}

func (s *Service) CloseOrder(ctx context.Context) {

}
