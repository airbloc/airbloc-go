package blockchain

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/airbloc/logger"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/pkg/errors"
)

func (c *Client) sendTxToBlockchain(ctx context.Context, tx *types.Transaction) error {
	return c.Client.SendTransaction(ctx, tx)
}

func (c *Client) sendTxToDelegateWithKey(ctx context.Context, tx *types.Transaction) error {
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

func (c *Client) sendTxToDelegateWithURL(ctx context.Context, tx *types.Transaction) error {
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
	return nil
}

func (c *Client) sendTxToDelegate(ctx context.Context, tx *types.Transaction) error {
	if !c.delegated {
		return errors.New("client is not delegate mode. please set fee payer")
	}

	// check fee payer
	feePayer, _ := tx.FeePayer()
	if feePayer != c.feePayer {
		return errors.New("fee payer mismatching")
	}

	switch {
	case c.feePayerTransactor != nil:
		return c.sendTxToDelegateWithKey(ctx, tx)
	case c.feePayerUrl != nil:
		return c.sendTxToDelegateWithURL(ctx, tx)
	default:
		return nil
	}
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
