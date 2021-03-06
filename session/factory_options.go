package session

import (
	"crypto/ecdsa"
	"errors"

	"github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/pkg/kas"
)

type BaseOption struct {
	KAS                *kas.Config
	BlockchainEndpoint string `split_words:"true" default:"https://api.baobab.klaytn.net:8651"`
	DeploymentPath     string `split_words:"true"`
	FeePayerEndpoint   string `split_words:"true" default:"http://localhost:3470/api"`
	FeePayerToken      string `split_words:"true"`
}

type sessionFactoryOption struct {
	// required
	key *ecdsa.PrivateKey

	// optional
	client      *blockchain.Client
	useFeePayer bool
	deployments bind.Deployments
}

func (opt sessionFactoryOption) Validate() error {
	if opt.key == nil {
		return errors.New("ecdsa privateKey required")
	}
	return nil
}

type FactoryOption func(*sessionFactoryOption)

func WithClient(client *blockchain.Client) FactoryOption {
	return func(opt *sessionFactoryOption) { opt.client = client }
}

func WithKey(key *ecdsa.PrivateKey) FactoryOption {
	return func(opt *sessionFactoryOption) { opt.key = key }
}

func WithUsingFeePayer() FactoryOption {
	return func(opt *sessionFactoryOption) { opt.useFeePayer = true }
}

func WithDeployments(deployments bind.Deployments) FactoryOption {
	return func(opt *sessionFactoryOption) { opt.deployments = deployments }
}
