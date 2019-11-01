package wrappers

import (
	"context"
	"math/big"
	"strings"

	ablbind "github.com/airbloc/airbloc-go/bind"
	contracts "github.com/airbloc/airbloc-go/bind/contracts"
	types "github.com/airbloc/airbloc-go/bind/types"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source app_registry.go -destination ./mocks/mock_app_registry.go -package mocks IAppRegistryContract

type IAppRegistryCalls interface {
	Exists(
		appName string,
	) (
		bool,
		error,
	)
	Get(
		appName string,
	) (
		types.App,
		error,
	)
	IsOwner(
		appName string,
		owner common.Address,
	) (
		bool,
		error,
	)
}

type IAppRegistryTransacts interface {
	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
	) (*chainTypes.Receipt, error)
	TransferAppOwner(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
		newOwner common.Address,
	) (*chainTypes.Receipt, error)
	Unregister(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
	) (*chainTypes.Receipt, error)
}

type IAppRegistryEvents interface {
	IAppRegistryFilterer
	IAppRegistryParser
	IAppRegistryWatcher
}

type IAppRegistryFilterer interface {
	FilterAppOwnerTransferred(
		opts *bind.FilterOpts,
		appAddr []common.Address,

		oldOwner []common.Address,

	) (ablbind.EventIterator, error)
	FilterRegistration(
		opts *bind.FilterOpts,
		appAddr []common.Address,

	) (ablbind.EventIterator, error)
	FilterUnregistration(
		opts *bind.FilterOpts,
		appAddr []common.Address,

	) (ablbind.EventIterator, error)
}

type IAppRegistryParser interface {
	ParseAppOwnerTransferred(log chainTypes.Log) (*contracts.AppRegistryAppOwnerTransferred, error)
	ParseAppOwnerTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.AppRegistryAppOwnerTransferred, error)
	ParseRegistration(log chainTypes.Log) (*contracts.AppRegistryRegistration, error)
	ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.AppRegistryRegistration, error)
	ParseUnregistration(log chainTypes.Log) (*contracts.AppRegistryUnregistration, error)
	ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.AppRegistryUnregistration, error)
}

type IAppRegistryWatcher interface {
	WatchAppOwnerTransferred(
		opts *bind.WatchOpts,
		sink chan<- *contracts.AppRegistryAppOwnerTransferred,
		appAddr []common.Address,

		oldOwner []common.Address,

	) (event.Subscription, error)
	WatchRegistration(
		opts *bind.WatchOpts,
		sink chan<- *contracts.AppRegistryRegistration,
		appAddr []common.Address,

	) (event.Subscription, error)
	WatchUnregistration(
		opts *bind.WatchOpts,
		sink chan<- *contracts.AppRegistryUnregistration,
		appAddr []common.Address,

	) (event.Subscription, error)
}

type IAppRegistryContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IAppRegistryCalls
	IAppRegistryTransacts
	IAppRegistryEvents
}

// Manager is contract wrapper struct
type AppRegistryContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	contracts.AppRegistryCaller
	contracts.AppRegistryFilterer
	contracts.AppRegistryTransactor
}

func NewAppRegistryContract(deployment ablbind.Deployment, backend ablbind.ContractBackend) interface{} {
	if deployment.Address() == (common.Address{}) {
		evmABI, err := abi.JSON(strings.NewReader(contracts.AppRegistryABI))
		if err != nil {
			panic(err)
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(contracts.AppRegistryAddress),
			common.HexToHash(contracts.AppRegistryTxHash),
			new(big.Int).SetBytes(common.HexToHash(contracts.AppRegistryCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, backend)

	contract := &AppRegistryContract{
		Deployment: deployment,
		client:     backend,

		AppRegistryCaller:     contracts.NewAppRegistryCaller(base),
		AppRegistryTransactor: contracts.NewAppRegistryTransactor(base),
		AppRegistryFilterer:   contracts.NewAppRegistryFilterer(base),
	}

	return contract
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string appName) constant returns(bool)
func (c *AppRegistryContract) Exists(
	appName string,
) (

	bool,
	error,
) {
	return c.AppRegistryCaller.Exists(nil, appName)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,address))
func (c *AppRegistryContract) Get(
	appName string,
) (

	types.App,
	error,
) {
	return c.AppRegistryCaller.Get(nil, appName)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string appName, address owner) constant returns(bool)
func (c *AppRegistryContract) IsOwner(
	appName string,
	owner common.Address,
) (

	bool,
	error,
) {
	return c.AppRegistryCaller.IsOwner(nil, appName, owner)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (c *AppRegistryContract) Register(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
) (*chainTypes.Receipt, error) {
	tx, err := c.AppRegistryTransactor.Register(c.client.Transactor(ctx, opts), appName)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (c *AppRegistryContract) TransferAppOwner(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
	newOwner common.Address,
) (*chainTypes.Receipt, error) {
	tx, err := c.AppRegistryTransactor.TransferAppOwner(c.client.Transactor(ctx, opts), appName, newOwner)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (c *AppRegistryContract) Unregister(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
) (*chainTypes.Receipt, error) {
	tx, err := c.AppRegistryTransactor.Unregister(c.client.Transactor(ctx, opts), appName)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
