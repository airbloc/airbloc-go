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
	"github.com/pkg/errors"
)

type Client struct {
	*client.Client
	ctx        context.Context
	cfg        ClientOpt
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
		TransactOpts: c.transactor.TransactOpts,
		FeePayer:     c.transactor.FeePayer,
		TxType:       c.transactor.TxType,
	}
	mergedOpts.Context = ctx
	for _, opt := range opts {
		if opt.TransactOpts != nil {
			mergedOpts.TransactOpts = opt.TransactOpts
		}
		if opt.FeePayer != (common.Address{}) {
			mergedOpts.FeePayer = opt.FeePayer
		}
		if opt.TxType != 0 {
			mergedOpts.TxType = opt.TxType
		}
	}
	return mergedOpts
}

func (c *Client) SetAccount(key *key.Key) {
	c.transactor = &TransactOpts{
		TransactOpts: bind.NewKeyedTransactor(key.PrivateKey),
		TxType:       types.TxTypeValueTransfer,
	}
}

func (c *Client) GetContract(contractType interface{}) interface{} {
	contract := c.contracts.GetContract(contractType)
	if contract == nil {
		panic("Contract not registered: " + reflect.ValueOf(contractType).Type().Name())
	}
	return contract
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	// TODO: task-waiting action required
	return c.Client.SendTransaction(ctx, tx)
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
