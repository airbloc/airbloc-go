package blockchain

import (
	"context"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/pkg/errors"
)

type Client struct {
	*client.Client
	ctx        context.Context
	cfg        ClientOpt
	key        *key.Key
	transactor *TransactOpts
	contracts  *contractManager
	logger     *logger.Logger
}

func NewClient(key *key.Key, rawurl string, cfg ClientOpt) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log := logger.New("klaytn")

	log.Debug(rawurl)
	// URL validation
	l, err := url.Parse(rawurl)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid URL: %s", rawurl)
	}
	if l.Scheme != "ws" {
		log.Error("Warning: You're using {} endpoint for Ethereum. Using WebSocket is recommended.",
			strings.ToUpper(l.Scheme))
	}

	// try to connect to Ethereum
	klayClient, err := client.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	cid, err := klayClient.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	log.Info("Using {} network", getChainName(cid))

	klayent := &Client{
		Client: klayClient,
		ctx:    context.TODO(),
		cfg:    cfg,
		key:    key,
		logger: log,
	}

	cm := NewContractManager(klayent)
	if err := cm.Load(cfg.DeploymentPath); err != nil {
		return nil, err
	}
	klayent.contracts = cm
	klayent.SetAccount(key)
	return klayent, nil
}

func (c Client) Account(ctx context.Context, opts ...*TransactOpts) *TransactOpts {
	mergedOpts := &TransactOpts{
		From:     c.transactor.From,
		FeePayer: c.transactor.FeePayer,
		Signer:   c.transactor.Signer,
		Nonce:    c.transactor.Nonce,
		Value:    c.transactor.Value,
		GasPrice: c.transactor.GasPrice,
		GasLimit: c.transactor.GasLimit,
		TxType:   c.transactor.TxType,
	}
	mergedOpts.Context = ctx
	for _, opt := range opts {
		if opt.From != (common.Address{}) {
			mergedOpts.From = opt.From
		}
		if opt.FeePayer != (common.Address{}) {
			mergedOpts.FeePayer = opt.FeePayer
		}
		if opt.Signer != nil {
			mergedOpts.Signer = opt.Signer
		}
		if opt.Nonce != nil {
			mergedOpts.Nonce = opt.Nonce
		}
		if opt.Value != nil {
			mergedOpts.Value = opt.Value
		}
		if opt.GasPrice != nil {
			mergedOpts.GasPrice = opt.GasPrice
		}
		if opt.GasLimit != 0 {
			mergedOpts.GasLimit = opt.GasLimit
		}
		if opt.TxType != 0 {
			mergedOpts.TxType = opt.TxType
		}
	}
	return mergedOpts
}

func (c *Client) SetAccount(key *key.Key) {
	c.transactor = NewKeyedTransactor(key.PrivateKey)
	c.transactor.TxType = types.TxTypeValueTransfer
}

func (c *Client) GetContract(contractType interface{}) interface{} {
	contract := c.contracts.GetContract(contractType)
	if contract == nil {
		panic("Contract not registered: " + reflect.ValueOf(contractType).Type().Name())
	}
	return contract
}

func (c *Client) SignTransaction(
	ctx context.Context,
	signerFn bind.SignerFn,
	tx *types.Transaction,
) (*types.Transaction, error) {
	chainID, err := c.ChainID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	signer := types.NewEIP155Signer(chainID)

	from, err := tx.From()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sender from transaction")
	}

	signedTx, err := signerFn(signer, from, tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign transaction with sender")
	}

	txType := tx.Type()
	if txType == types.TxTypeFeeDelegatedValueTransfer ||
		txType == types.TxTypeFeeDelegatedSmartContractExecution {
		feePayer, err := tx.FeePayer()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get fee payer from transaction")
		}
		if from == c.transactor.From {
			return nil, errors.Wrap(err, "sender and fee payer should not be same")
		}
		if feePayer != crypto.PubkeyToAddress(c.key.PublicKey) {
			return nil, errors.Wrap(err, "fee payer is not the same with request")
		}

		if err = signedTx.SignFeePayer(signer, c.key.PrivateKey); err != nil {
			return nil, errors.Wrap(err, "failed to sign transaction with fee payer")
		}
	}

	return signedTx, nil
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

// WaitMined waits until transcaction created.
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	methodName, details := GetTransactionDetails(c.contracts, tx)
	timer := c.logger.Timer()

	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		timer.End("Transaction to {} failed", methodName, details)
		return nil, ErrTxFailed
	}
	timer.End("Transacted {}", methodName, details)
	// err = c.waitConfirmation(ctx)
	return receipt, err
}

// WaitDeployed waits until contract created.
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
	// err = c.waitConfirmation(ctx)
	return receipt, err
}
