package airbloc

import (
	"context"
	"math/big"
	"net/url"
	"time"

	ablbind "github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/logger"

	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	klayClient "github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

// getChainName returns chain name by chain ID (network ID), according to EIP-155.
func getChainName(cid *big.Int) string {
	switch cid.String() {
	case "1":
		return "Ethereum Main"
	case "3":
		return "Ethereum Ropsten Test"
	case "4":
		return "Ethereum Rinkeby Test"
	case "1000":
		return "Klaytn Aspen Test"
	case "1001":
		return "Klaytn Baobab Test"
	case "8217":
		return "Klaytn Main"
	}
	return "EVM Private"
}

type clientData struct {
	*klayClient.Client
}

type Client struct {
	clientData
	log *logger.Logger
}

func NewClient(ctx context.Context, endpoint string) (*Client, error) {
	log := logger.New("klaytn")

	if _, err := url.Parse(endpoint); err != nil {
		return nil, errors.Errorf("invalid URL: %s", endpoint)
	}

	client, err := klayClient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	cid, err := client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	log.Info("Using {} network", getChainName(cid))

	return &Client{
		clientData: clientData{client},
		log:        log,
	}, nil
}

func (c *Client) Client() *klayClient.Client {
	return c.clientData.Client
}

func (c *Client) Deployment(string) (ablbind.Deployment, bool) {
	return ablbind.Deployment{}, false // use default information (mainnet)
}

func (c *Client) Transactor(ctx context.Context, opts ...*ablbind.TransactOpts) *ablbind.TransactOpts {
	return ablbind.MergeTxOpts(ctx, nil, opts...)
}

func (c *Client) waitMined(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, err := c.TransactionReceipt(ctx, hash)
		if receipt != nil {
			return receipt, nil
		}
		if err != nil {
			c.log.Debug("Receipt retrieval failed", "err", err)
		} else {
			c.log.Debug("Transaction not yet mined")
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

// WaitMined waits until transcaction created.
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return c.WaitMinedWithHash(ctx, tx.Hash())
}

func (c *Client) WaitMinedWithHash(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	receipt, err := c.waitMined(ctx, hash)
	if err != nil {
		return nil, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return nil, errors.New("tx failed")
	}
	return receipt, err
}

// WaitDeployed waits until contract created.
func (c *Client) WaitDeployed(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	txType := tx.Type()

	if txType != types.TxTypeSmartContractDeploy &&
		txType != types.TxTypeFeeDelegatedSmartContractDeploy &&
		txType != types.TxTypeFeeDelegatedSmartContractDeployWithRatio {
		return nil, errors.New("tx is not contract creation")
	}

	receipt, err := c.WaitMined(ctx, tx)
	if err != nil {
		return nil, err
	}
	if receipt.ContractAddress == (common.Address{}) {
		return nil, errors.New("zero address")
	}

	code, err := c.CodeAt(ctx, receipt.ContractAddress, nil)
	if err == nil && len(code) == 0 {
		err = bind.ErrNoCodeAfterDeploy
	}
	return receipt, err
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	hash, err := c.SendRawTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	return c.WaitMinedWithHash(ctx, hash)
}

func (c *Client) MakeTransaction(opts *ablbind.TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	return opts.MakeTransaction(c, contract, input)
}
