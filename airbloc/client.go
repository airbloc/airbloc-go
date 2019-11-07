package airbloc

import (
	"context"
	"math/big"
	"net/url"

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

type Client struct {
	*klayClient.Client
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
		Client: client,
		log:    log,
	}, nil
}

// WaitMined waits until transcaction created.
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	methodName, details := "TODO: mocked method name", "TODO: mocked details"
	timer := c.log.Timer()

	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		timer.End("Transaction to {} failed", methodName, details)
		return nil, errors.New("tx failed")
	}
	timer.End("Transacted {}", methodName, details)
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
