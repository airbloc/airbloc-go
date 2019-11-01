package airbloc

import (
	"context"
	"crypto/ecdsa"

	"github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/logger"
)

type airbloc struct {
	ctx    context.Context
	cancel context.CancelFunc
	client *blockchain.Client
	logger *logger.Logger

	blockchain.ContractManager
}

func NewAirbloc(ctx context.Context, endpoint string) (*airbloc, error) {
	return NewAirblocWithOption(ctx, Options{
		Blockchain: blockchain.Options{Endpoint: endpoint},
	})
}

func NewAirblocWithOption(ctx context.Context, opt Options) (*airbloc, error) {
	ablctx, cancel := context.WithCancel(ctx)
	client, err := blockchain.NewClient(ablctx, opt.Blockchain)
	if err != nil {
		return nil, err
	}

	deployments, err := bind.GetDeploymentsFrom(opt.DeploymentPath)
	if err != nil {
		return nil, err
	}

	manager, err := blockchain.NewContractManager(client, deployments)
	if err != nil {
		client.Close()
		return nil, err
	}

	return &airbloc{
		ctx:             ablctx,
		cancel:          cancel,
		client:          client,
		ContractManager: manager,
	}, nil
}

func (abl *airbloc) Client() *blockchain.Client {
	return abl.client
}

func (abl *airbloc) GetTransactor(ctx context.Context, txOpts ...*bind.TransactOpts) *bind.TransactOpts {
	return abl.client.Transactor(ctx, txOpts...)
}

func (abl *airbloc) ChangeTransactor(key *ecdsa.PrivateKey) {
	abl.client.SetTransactor(key)
}

func (abl *airbloc) Close() {
	abl.cancel()
	abl.client.Close()
}
