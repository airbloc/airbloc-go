package blockchain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/logger"

	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	klayClient "github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

type Client struct {
	*klayClient.Client

	transactor *adapter.TransactOpts
	readOnly   bool

	feePayer           common.Address
	feePayerUrl        *url.URL
	feePayerTransactor *adapter.TransactOpts
	delegated          bool

	log *logger.Logger
}

func NewClient(ctx context.Context, opts Options) (*Client, error) {
	log := logger.New("klaytn")

	if _, err := url.Parse(opts.Endpoint); err != nil {
		return nil, errors.Errorf("invalid URL: %s", opts.Endpoint)
	}

	// try to connect to Ethereum
	c, err := klayClient.DialContext(ctx, opts.Endpoint)
	if err != nil {
		return nil, err
	}
	cid, err := c.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	log.Info("Using {} network", getChainName(cid))

	client := &Client{
		Client: c,
		log:    log,
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

func (c Client) Transactor(ctx context.Context, opts ...*adapter.TransactOpts) *adapter.TransactOpts {
	return adapter.MergeTxOpts(ctx, c.transactor, opts...)
}

func (c *Client) SetTransactor(key *ecdsa.PrivateKey) {
	c.transactor = adapter.NewKeyedTransactor(key)
	c.readOnly = key == nil
}

func (c *Client) FeePayer() common.Address {
	return c.feePayer
}

func (c *Client) SetFeePayerWithKey(key *ecdsa.PrivateKey) {
	c.feePayerTransactor = adapter.NewKeyedFeePayerTransactor(key)
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

func (c *Client) sendTxToBlockchain(ctx context.Context, tx *types.Transaction) error {
	return c.Client.SendTransaction(ctx, tx)
}

func (c *Client) sendTxToDelegate(ctx context.Context, tx *types.Transaction) error {
	// check fee payer
	feePayer, _ := tx.FeePayer()
	if feePayer != c.feePayer {
		return errors.New("fee payer mismatching")
	}

	// sign and send to blockchain
	if c.feePayerTransactor != nil {
		chainID, err := c.ChainID(ctx)
		if err != nil {
			return errors.Wrap(err, "fetching chain id")
		}

		signer := types.NewEIP155Signer(chainID)
		signedTx, err := c.feePayerTransactor.Signer(signer, c.feePayer, tx)
		if err != nil {
			return errors.Wrap(err, "signing tx")
		}

		return c.Client.SendTransaction(ctx, signedTx)
	}

	// request to fee payer worker
	if c.feePayerUrl != nil {
		rawTxData, err := tx.MarshalJSON()
		if err != nil {
			return errors.Wrap(err, "marshaling tx")
		}

		resp, err := http.Post(c.feePayerUrl.RequestURI(), "application/json", bytes.NewReader(rawTxData))
		if err != nil {
			return errors.Wrap(err, "sending tx to delegate")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated { // transaction created
			d, _ := ioutil.ReadAll(resp.Body)
			c.log.Error("Failed to request transaction", logger.Attrs{
				"status-code": resp.StatusCode,
				"message":     string(d),
			})
			return errors.New("request failed")
		}
	}

	return nil
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	txType := tx.Type()
	switch txType {
	case types.TxTypeValueTransfer:
		fallthrough
	case types.TxTypeSmartContractExecution:
		return c.sendTxToBlockchain(ctx, tx)
	case types.TxTypeFeeDelegatedValueTransfer:
		fallthrough
	case types.TxTypeFeeDelegatedSmartContractDeploy:
		return c.sendTxToDelegate(ctx, tx)
	default:
		return errors.New("unsupported transaction type")
	}
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
