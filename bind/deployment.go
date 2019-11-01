package bind

import (
	"encoding/json"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/common"
)

type deploymentData struct {
	Address   common.Address `json:"address"`
	TxHash    common.Hash    `json:"tx_hash"`
	CreatedAt *big.Int       `json:"created_at"`
	ParsedABI abi.ABI        `json:"abi"`
}

type Deployment struct{ deploymentData }

func (d Deployment) Address() common.Address {
	return d.deploymentData.Address
}

func (d Deployment) TxHash() common.Hash {
	return d.deploymentData.TxHash
}

func (d Deployment) CreatedAt() *big.Int {
	return d.deploymentData.CreatedAt
}

func NewDeployment(address common.Address, txHash common.Hash, createdAt *big.Int, parsedABI abi.ABI) Deployment {
	return Deployment{deploymentData{
		Address:   address,
		TxHash:    txHash,
		CreatedAt: createdAt,
		ParsedABI: parsedABI,
	}}
}

type Deployments map[string]deploymentData

func getDeploymentsFromFile(path string) (Deployments, error) {
	deployments := Deployments{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&deployments); err != nil {
		return nil, err
	}
	return deployments, nil
}

func getDeploymentsFromWeb(path string) (Deployments, error) {
	deployments := Deployments{}

	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&deployments); err != nil {
		return nil, err
	}
	return deployments, nil
}

func GetDeploymentsFrom(path string) (Deployments, error) {
	if path == "" {
		return nil, nil
	}

	if strings.HasPrefix(path, "http://") ||
		strings.HasPrefix(path, "https://") {
		return getDeploymentsFromWeb(path)
	} else {
		return getDeploymentsFromFile(path)
	}
}

func (ds Deployments) Get(contract string) (Deployment, bool) {
	d, ok := ds[contract]
	return Deployment{d}, ok
}
