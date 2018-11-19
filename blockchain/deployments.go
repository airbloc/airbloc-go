package blockchain

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
)

type Deployments struct {
	Accounts           *adapter.Accounts
	AppRegistry        *adapter.AppRegistry
	CollectionRegistry *adapter.CollectionRegistry
	DataRegistry       *adapter.DataRegistry
	SchemaRegistry     *adapter.SchemaRegistry
	Exchange           *adapter.Exchange
}

func DeployAll(client *Client) (*Deployments, error) {
	ctx := context.Background()

	_, tx, accounts, err := adapter.DeployAccounts(client.Account(), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deploy contract Accounts")
	}
	if receipt, err := client.WaitDeployed(ctx, tx); err != nil {
		return nil, errors.Wrap(err, "failed to wait deployment of contract Accounts")
	} else {
		log.Info("Account contract deployed", "address", receipt.ContractAddress, "gasUsed", receipt.GasUsed)
	}
	// TODO: too many :(
	return &Deployments{
		Accounts: accounts,
	}, nil
}

func LoadDeployments(path string, client *Client) (*Deployments, error) {
	var deployedAddressInfo struct {
		Accounts           string
		AppRegistry        string
		CollectionRegistry string
		DataRegistry       string
		SchemaRegistry     string
		Exchange           string
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load deployments from %s", path)
	}

	if err := json.Unmarshal(data, &deployedAddressInfo); err != nil {
		return nil, errors.Wrap(err, "failed to parse JSON")
	}

	accounts, err := adapter.NewAccounts(common.HexToAddress(deployedAddressInfo.Accounts), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind to contract Accounts")
	}

	appRegistry, err := adapter.NewAppRegistry(common.HexToAddress(deployedAddressInfo.AppRegistry), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind to contract AppRegistry")
	}

	collectionRegistry, err := adapter.NewCollectionRegistry(common.HexToAddress(deployedAddressInfo.CollectionRegistry), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind to contract CollectionRegistry")
	}

	dataRegistry, err := adapter.NewDataRegistry(common.HexToAddress(deployedAddressInfo.DataRegistry), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind to contract DataRegistry")
	}

	schemaRegistry, err := adapter.NewSchemaRegistry(common.HexToAddress(deployedAddressInfo.SchemaRegistry), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind to contract SchemaRegistry")
	}

	exchange, err := adapter.NewExchange(common.HexToAddress(deployedAddressInfo.Exchange), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind to contract Exchange")
	}

	return &Deployments{
		Accounts:           accounts,
		AppRegistry:        appRegistry,
		CollectionRegistry: collectionRegistry,
		DataRegistry:       dataRegistry,
		SchemaRegistry:     schemaRegistry,
		Exchange:           exchange,
	}, nil
}
