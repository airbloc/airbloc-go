package blockchain

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"net/http"
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

type Client struct {
	*klayClient.Client
	deployments ablbind.Deployments

	transactor *ablbind.TransactOpts
	readOnly   bool

	feePayer           common.Address
	feePayerUrl        *url.URL
	feePayerTransactor *ablbind.TransactOpts
	delegated          bool

	log *logger.Logger
}

func NewClient(ctx context.Context, opts Options) (*Client, error) {
	log := logger.New("klaytn")

	deployments, err := ablbind.GetDeploymentsFrom(opts.DeploymentPath)
	if err != nil {
		return nil, err
	}

	if _, err = url.Parse(opts.Endpoint); err != nil {
		return nil, errors.Errorf("invalid URL: %s", opts.Endpoint)
	}

	// try to connect to Klaytn
	blockchainClient, err := klayClient.DialContext(ctx, opts.Endpoint)
	if err != nil {
		return nil, err
	}
	cid, err := blockchainClient.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	log.Info("Using {} network", getChainName(cid))

	client := &Client{
		Client:      blockchainClient,
		deployments: deployments,
		log:         log,
	}

	if opts.Key != nil {
		client.SetTransactor(opts.Key)
	}

	if opts.FeePayerEndpoint != "" {
		if err := client.SetFeePayerWithUrl(opts.FeePayerEndpoint); err != nil {
			return nil, errors.Wrap(err, "initializing fee payer with endpoint")
		}
	}

	if opts.FeePayerKey != nil {
		client.SetFeePayerWithKey(opts.FeePayerKey)
	}

	return client, nil
}

func (c Client) Transactor(ctx context.Context, opts ...*ablbind.TransactOpts) *ablbind.TransactOpts {
	return ablbind.MergeTxOpts(ctx, c.transactor, opts...)
}

func (c *Client) SetTransactor(key *ecdsa.PrivateKey) {
	c.transactor = ablbind.NewKeyedTransactor(key)
	c.readOnly = key == nil
}

func (c *Client) FeePayer() common.Address {
	return c.feePayer
}

func (c *Client) SetFeePayerWithKey(key *ecdsa.PrivateKey) {
	c.feePayerTransactor = ablbind.NewKeyedFeePayerTransactor(key)
	c.feePayer = c.feePayerTransactor.From
	c.delegated = c.feePayer != (common.Address{})
}

func (c *Client) SetFeePayerWithUrl(endpoint string) (err error) {
	c.feePayerUrl, err = url.Parse(endpoint)
	if err != nil {
		return errors.Errorf("invalid URL: %s", endpoint)
	}

	client := &http.Client{Timeout: time.Minute}
	resp, err := client.Get(endpoint + "/address")
	if err != nil {
		return errors.Wrap(err, "requesting fee payer address")
	}
	defer resp.Body.Close()

	var addr struct {
		Address common.Address `json:"address"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&addr); err != nil {
		return errors.Wrap(err, "decoding json response body")
	}

	c.feePayer = addr.Address
	c.delegated = c.feePayer != (common.Address{})
	return nil
}

func (c *Client) GetDeployment(name string) (ablbind.Deployment, bool) {
	return c.deployments.Get(name)
}

// WaitMined waits until transcaction created.
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	//methodName, details := GetTransactionDetails(c.contracts, tx)
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
