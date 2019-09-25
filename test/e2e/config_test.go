package e2e

import (
	"crypto/ecdsa"
	"log"
	"net/http"

	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/jinzhu/configor"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/stretchr/testify/require"
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
	serverConfig := new(service.Config)
	if err := configor.Load(serverConfig, ConfigFilePath); err != nil {
		log.Fatalln("Unable to load service config from", ConfigFilePath)
	}

	privateKey, err := crypto.LoadECDSA("../../" + serverConfig.PrivateKeyPath)
	require.NoError(t, err)

	// load contract deployments
	deployments := make(map[string]common.Address)
	res, err := http.Get(serverConfig.Blockchain.DeploymentPath)
	require.NoError(t, err)
	defer res.Body.Close()
	require.NoError(t, json.NewDecoder(res.Body).Decode(&deployments))

	userDelegateConfig := new(service.Config)
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
