package client

import (
	"net/http"

	"net/url"

	"fmt"

	"github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/pkg/errors"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	baseHeader http.Header
}

// load config example in unchainio/pkg/xconfig
func New(config ClientConfig) (*Client, error) {

	url, err := url.Parse(config.Url)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse url from the config")
	}

	client := &Client{
		baseURL:    url,
		httpClient: http.DefaultClient,
		baseHeader: http.Header{},
	}

	// Set API Key & Secret header
	for k, v := range config.Headers {
		client.baseHeader.Set(k, v)
	}

	return client, nil
}

func (c *Client) GetBlock(blockHeight string) (transaction.Block, error) {
	var block transaction.Block
	path := fmt.Sprintf("blocks/%s", blockHeight)
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return block, errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &block)
	if err != nil {
		return block, errors.Wrap(err, "Request unsuccessful")
	}

	return block, nil
}

func (c *Client) GetTransaction(transactionID string) (transaction.Transaction, error) {
	var txn transaction.Transaction

	path := fmt.Sprintf("transactions/%s", transactionID)
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return txn, errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &txn)
	if err != nil {
		return txn, errors.Wrap(err, "Request unsuccessful")
	}

	return txn, nil
}

func (c *Client) ListBlocks(transactionID string) ([]transaction.Block, error) {
	var blocks []transaction.Block

	path := fmt.Sprintf("blocks?transaction_id=%s", transactionID)
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return blocks, errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &blocks)
	if err != nil {
		return blocks, errors.Wrap(err, "Request unsuccessful")
	}

	return blocks, nil
}

func (c *Client) ListOutputs(pubKey string, spent bool) ([]transaction.OutputLocation, error) {
	var outputs []transaction.OutputLocation

	var path string
	if spent {
		path = fmt.Sprintf("outputs?public_key=%s?spent=true", pubKey)
	} else {
		path = fmt.Sprintf("outputs?public_key=%s", pubKey)
	}
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return outputs, errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &outputs)
	if err != nil {
		return outputs, errors.Wrap(err, "Request unsuccessful")
	}

	return outputs, nil
}

func (c *Client) ListTransactions(assetID, operation string) ([]transaction.Transaction, error) {
	var txns []transaction.Transaction

	path := fmt.Sprintf("transactions?asset_id=%s?%s", assetID, operation)

	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return txns, errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &txns)
	if err != nil {
		return txns, errors.Wrap(err, "Request unsuccessful")
	}

	return txns, nil
}

type TxrOptions struct {
	path string
}

/*
	Default transaction mode is async: client immediately receives response from server
	transaction validity is not checked
*/
func (c *Client) PostTransaction(txn *transaction.Transaction) error {

	req, err := c.newRequest("POST", "transactions", txn)
	if err != nil {
		return errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &txn)
	if err != nil {
		return errors.Wrap(err, "Request unsuccessful")
	}

	// FIXME make sure that returned txn decodes output public keys as []byte instead of []string
	return nil
}

/*
	Client receives a response when the transaction is valid
*/
func (c *Client) PostTransactionSync(txn *transaction.Transaction) error {

	req, err := c.newRequest("POST", "transactions?mode=sync", txn)
	if err != nil {
		return errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &txn)
	if err != nil {
		return errors.Wrap(err, "Request unsuccessful")
	}

	// FIXME make sure that returned txn decodes output public keys as []byte instead of []string
	return nil
}

/* Client receives a response when transaction is committed to block */
func (c *Client) PostTransactionCommit(txn *transaction.Transaction) error {

	req, err := c.newRequest("POST", "transactions?mode=commit", txn)
	if err != nil {
		return errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &txn)
	if err != nil {
		return errors.Wrap(err, "Request unsuccessful")
	}

	// FIXME make sure that returned txn decodes output public keys as []byte instead of []string
	return nil
}

// TODO add search string to request
func (c *Client) SearchAsset(search string, limit int) ([]transaction.Asset, error) {
	var assets []transaction.Asset

	req, err := c.newRequest("GET", "assets", nil)
	if err != nil {
		return assets, errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &assets)
	if err != nil {
		return assets, errors.Wrap(err, "Request unsuccessful")
	}

	return assets, nil
}

// TODO add search string to request
func (c *Client) SearchMetadata(search string, limit int) ([]transaction.Metadata, error) {
	var metadatas []transaction.Metadata

	req, err := c.newRequest("GET", "metadata", nil)
	if err != nil {
		return metadatas, errors.Wrap(err, "Could not create http request")
	}

	err = c.do(req, &metadatas)
	if err != nil {
		return metadatas, errors.Wrap(err, "Request unsuccessful")
	}

	return metadatas, nil
}
