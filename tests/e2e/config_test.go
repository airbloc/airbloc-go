package e2e

import (
	"crypto/ecdsa"
	"github.com/airbloc/airbloc-go/shared/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/require"
	"log"
	"os"
)

const ConfigFilePath = "../../config.yml"
const UserDelegateConfigPath = "../../config-userdelegate.yml"

type testConfig struct {
	TransactorPrivateKey *ecdsa.PrivateKey
	UserDelegateAddress  common.Address
	EthereumEndpoint     string
	DeployedContracts    map[string]common.Address
}

func (t *T) loadConfig() {
	serverConfig := new(node.Config)
	if err := configor.Load(serverConfig, ConfigFilePath); err != nil {
		log.Fatalln("Unable to load node config from", ConfigFilePath)
	}

	privateKey, err := crypto.LoadECDSA("../../" + serverConfig.PrivateKeyPath)
	require.NoError(t, err)

	// load contract deployments
	deployments := make(map[string]common.Address)
	file, err := os.OpenFile("../../"+serverConfig.Blockchain.DeploymentPath, os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(file).Decode(&deployments))

	userDelegateConfig := new(node.Config)
	if err := configor.Load(userDelegateConfig, UserDelegateConfigPath); err != nil {
		log.Fatalln("Unable to load user delegate config from", UserDelegateConfigPath)
	}

	udKey, err := crypto.LoadECDSA("../../" + userDelegateConfig.PrivateKeyPath)
	require.NoError(t, err)

	t.config = &testConfig{
		TransactorPrivateKey: privateKey,
		UserDelegateAddress:  crypto.PubkeyToAddress(udKey.PublicKey),
		DeployedContracts:    deployments,
		EthereumEndpoint:     serverConfig.Blockchain.Endpoint,
	}
}
