package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	*ethclient.Client
	ctx        context.Context
	cfg        ClientOpt
	transactor *bind.TransactOpts
}

func NewClient(transactor *bind.TransactOpts, url string, cfg ClientOpt) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:     client,
		ctx:        context.Background(),
		cfg:        cfg,
		transactor: transactor,
	}, err
}

func (c Client) Account() *bind.TransactOpts {
	return c.transactor
}

func (c *Client) SetAccount(account *bind.TransactOpts) {
	c.transactor = account
}

func (c *Client) waitConfirmation(ctx context.Context) error {
	ch := make(chan *types.Header)
	sub, err := c.SubscribeNewHead(c.ctx, ch)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for count := c.cfg.Confirmation; count > 0; {
		select {
		case <-ch:
			count--
		case <-ctx.Done():
			return context.DeadlineExceeded
		}
	}
	return err
}

// Wait Mined
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return nil, ErrTxFailed
	}
	err = c.waitConfirmation(ctx)
	return receipt, err
}

// Wait Deployed
func (c *Client) WaitDeployed(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	if tx.To() != nil {
		return nil, ErrTxNoContract
	}

	receipt, err := c.WaitMined(ctx, tx)
	if err != nil {
		return nil, err
	}
	if receipt.ContractAddress == (common.Address{}) {
		return nil, ErrZeroAddress
	}

	code, err := c.CodeAt(ctx, receipt.ContractAddress, nil)
	if err == nil && len(code) == 0 {
		err = bind.ErrNoCodeAfterDeploy
	}
	err = c.waitConfirmation(ctx)
	return receipt, err
}
