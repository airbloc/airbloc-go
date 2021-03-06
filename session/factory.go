package session

import (
	"context"

	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/logger"
	"github.com/pkg/errors"
)

type Factory interface {
	NewSession(ctx context.Context, opts ...FactoryOption) (Session, account.Account, error)
}

type sessionFactory struct {
	client      *blockchain.Client
	feePayer    account.FeePayer
	deployments bind.Deployments
}

func NewFactory(ctx context.Context, opt BaseOption) (Factory, error) {
	var client *blockchain.Client
	if opt.KAS != nil {
		// connect to Klaytn API Service (KAS)
		cli, err := blockchain.NewClientWithKAS(ctx, *opt.KAS)
		if err != nil {
			return nil, errors.Wrapf(err, "connect KAS")
		}
		client = cli
	} else {
		// connect to the Klaytn node directly
		cli, err := blockchain.NewClient(ctx, opt.BlockchainEndpoint)
		if err != nil {
			return nil, errors.Wrapf(
				err, "initialize klaytn client. url=%s",
				opt.BlockchainEndpoint,
			)
		}
		client = cli
	}

	var token *string = nil
	if opt.FeePayerToken != "" {
		token = &opt.FeePayerToken
	}

	feePayer, err := account.NewFeePayer(ctx, nil, opt.FeePayerEndpoint, token)
	if err != nil {
		client.Close()
		return nil, errors.Wrapf(
			err, "initialize fee payer client. url=%s",
			opt.FeePayerEndpoint,
		)
	}

	feePayerAddr, err := feePayer.Address(ctx)
	if err != nil {
		client.Close()
		return nil, errors.Wrapf(err, "fetch fee payer address")
	}

	log := logger.New("session-factory")
	log.Info("Using feePayer {}", feePayerAddr.Hex())

	deployments, err := bind.GetDeploymentsFrom(opt.DeploymentPath)
	if err != nil {
		client.Close()
		return nil, errors.Wrapf(
			err, "fetch deployments from %s",
			opt.DeploymentPath,
		)
	}
	if deployments != nil {
		log.Info("Using deployment at {}", opt.DeploymentPath)
	}

	return &sessionFactory{
		client:      client,
		feePayer:    feePayer,
		deployments: deployments,
	}, nil
}

func MustNewFactory(ctx context.Context, opt BaseOption) Factory {
	if sf, err := NewFactory(ctx, opt); err != nil {
		panic(err)
	} else {
		return sf
	}
}

func (sf sessionFactory) NewSession(ctx context.Context, opts ...FactoryOption) (
	Session,
	account.Account,
	error,
) {
	factoryOption := sessionFactoryOption{}
	for _, opt := range opts {
		opt(&factoryOption)
	}
	if err := factoryOption.Validate(); err != nil {
		return Session{},
			account.Account{},
			errors.Wrap(err, "validate factory option")
	}

	var (
		acc = account.NewKeyedAccount(factoryOption.key)
		err error
	)
	if factoryOption.useFeePayer {
		acc, err = account.NewKeyedAccountWithFeePayer(
			ctx,
			factoryOption.key,
			sf.feePayer,
		)
		if err != nil {
			return Session{},
				account.Account{},
				errors.Wrap(err, "new keyed account with fee payer")
		}
	}

	deployments := sf.deployments
	if factoryOption.deployments != nil {
		deployments = factoryOption.deployments
	}
	sess, err := NewSession(Config{
		Account:     acc,
		Client:      sf.client,
		Deployments: deployments,
	})
	return sess, acc, err
}

func (sf sessionFactory) Close() {
	sf.client.Close()
}
