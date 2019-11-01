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

//go:generate mockgen -source data_type_registry.go -destination ./mocks/mock_data_type_registry.go -package mocks IDataTypeRegistryContract

type IDataTypeRegistryCalls interface {
	Exists(
		name string,
	) (
		bool,
		error,
	)
	Get(
		name string,
	) (
		types.DataType,
		error,
	)
	IsOwner(
		name string,
		owner common.Address,
	) (
		bool,
		error,
	)
}

type IDataTypeRegistryTransacts interface {
	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		name string,
		schemaHash common.Hash,
	) (*chainTypes.Receipt, error)
	Unregister(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		name string,
	) (*chainTypes.Receipt, error)
}

type IDataTypeRegistryEvents interface {
	IDataTypeRegistryFilterer
	IDataTypeRegistryParser
	IDataTypeRegistryWatcher
}

type IDataTypeRegistryFilterer interface {
	FilterRegistration(
		opts *bind.FilterOpts,

	) (ablbind.EventIterator, error)
	FilterUnregistration(
		opts *bind.FilterOpts,

	) (ablbind.EventIterator, error)
}

type IDataTypeRegistryParser interface {
	ParseRegistration(log chainTypes.Log) (*contracts.DataTypeRegistryRegistration, error)
	ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.DataTypeRegistryRegistration, error)
	ParseUnregistration(log chainTypes.Log) (*contracts.DataTypeRegistryUnregistration, error)
	ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.DataTypeRegistryUnregistration, error)
}

type IDataTypeRegistryWatcher interface {
	WatchRegistration(
		opts *bind.WatchOpts,
		sink chan<- *contracts.DataTypeRegistryRegistration,

	) (event.Subscription, error)
	WatchUnregistration(
		opts *bind.WatchOpts,
		sink chan<- *contracts.DataTypeRegistryUnregistration,

	) (event.Subscription, error)
}

type IDataTypeRegistryContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IDataTypeRegistryCalls
	IDataTypeRegistryTransacts
	IDataTypeRegistryEvents
}

// Manager is contract wrapper struct
type DataTypeRegistryContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	contracts.DataTypeRegistryCaller
	contracts.DataTypeRegistryFilterer
	contracts.DataTypeRegistryTransactor
}

func NewDataTypeRegistryContract(deployment ablbind.Deployment, backend ablbind.ContractBackend) interface{} {
	if deployment.Address() == (common.Address{}) {
		evmABI, err := abi.JSON(strings.NewReader(contracts.DataTypeRegistryABI))
		if err != nil {
			panic(err)
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(contracts.DataTypeRegistryAddress),
			common.HexToHash(contracts.DataTypeRegistryTxHash),
			new(big.Int).SetBytes(common.HexToHash(contracts.DataTypeRegistryCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, backend)

	contract := &DataTypeRegistryContract{
		Deployment: deployment,
		client:     backend,

		DataTypeRegistryCaller:     contracts.NewDataTypeRegistryCaller(base),
		DataTypeRegistryTransactor: contracts.NewDataTypeRegistryTransactor(base),
		DataTypeRegistryFilterer:   contracts.NewDataTypeRegistryFilterer(base),
	}

	return contract
}

func (c *DataTypeRegistryContract) GetSelectors() map[string]string {
	selectors := make(map[string]string)
	selectors["0x261a323e"] = "exists(string)"
	selectors["0x693ec85e"] = "get(string)"
	selectors["0xbde1eee7"] = "isOwner(string,address)"

	selectors["0x656afdee"] = "register(string,bytes32)"
	selectors["0x6598a1ae"] = "unregister(string)"

	return selectors
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string name) constant returns(bool)
func (c *DataTypeRegistryContract) Exists(
	name string,
) (

	bool,
	error,
) {
	return c.DataTypeRegistryCaller.Exists(nil, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (c *DataTypeRegistryContract) Get(
	name string,
) (

	types.DataType,
	error,
) {
	return c.DataTypeRegistryCaller.Get(nil, name)
}

// IsOwner is a free data retrieval call binding the contract method 0xbde1eee7.
//
// Solidity: function isOwner(string name, address owner) constant returns(bool)
func (c *DataTypeRegistryContract) IsOwner(
	name string,
	owner common.Address,
) (

	bool,
	error,
) {
	return c.DataTypeRegistryCaller.IsOwner(nil, name, owner)
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (c *DataTypeRegistryContract) Register(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	name string,
	schemaHash common.Hash,
) (*chainTypes.Receipt, error) {
	tx, err := c.DataTypeRegistryTransactor.Register(c.client.Transactor(ctx, opts), name, schemaHash)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (c *DataTypeRegistryContract) Unregister(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	name string,
) (*chainTypes.Receipt, error) {
	tx, err := c.DataTypeRegistryTransactor.Unregister(c.client.Transactor(ctx, opts), name)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
