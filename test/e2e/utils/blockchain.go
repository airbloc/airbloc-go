package e2eutils

import (
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/key"
)

const (
	deploymentPath = "http://localhost:8500"
	blockchainPath = "https://api.baobab.klaytn.net:8651"
	testerKeyPath  = "../tester.key"
)

func ConnectBlockchain() (*blockchain.Client, error) {
	clientKey, err := key.Load(testerKeyPath)
	if err != nil {
		return nil, err
	}

	client, err := blockchain.NewClient(clientKey, blockchainPath, blockchain.ClientOpt{
		Confirmation:   5,
		DeploymentPath: deploymentPath,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
