package wrappers

import (
	"context"
	"math/big"
	"strings"

	ablbind "github.com/airbloc/airbloc-go/shared/adapter"
	contracts "github.com/airbloc/airbloc-go/shared/adapter/contracts"
	types "github.com/airbloc/airbloc-go/shared/adapter/types"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	chainTypes "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

//go:generate mockgen -source controller_registry.go -destination ./mocks/mock_controller_registry.go -package mocks IControllerRegistryContract

type IControllerRegistryCalls interface {
	Exists(
		controller common.Address,
	) (
		bool,
		error,
	)
	Get(
		controller common.Address,
	) (
		types.DataController,
		error,
	)
}

type IControllerRegistryTransacts interface {
	Register(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		controllerAddr common.Address,
	) (*chainTypes.Receipt, error)
}

type IControllerRegistryEvents interface {
	IControllerRegistryFilterer
	IControllerRegistryParser
	IControllerRegistryWatcher
}

type IControllerRegistryFilterer interface {
	FilterOwnershipTransferred(
		opts *bind.FilterOpts,
		previousOwner []common.Address,
		newOwner []common.Address,
	) (ablbind.EventIterator, error)
	FilterRegistration(
		opts *bind.FilterOpts,
		controller []common.Address,
	) (ablbind.EventIterator, error)
	FilterUnregistration(
		opts *bind.FilterOpts,
		controller []common.Address,
	) (ablbind.EventIterator, error)
}

type IControllerRegistryParser interface {
	ParseOwnershipTransferred(log chainTypes.Log) (*contracts.ControllerRegistryOwnershipTransferred, error)
	ParseOwnershipTransferredFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ControllerRegistryOwnershipTransferred, error)
	ParseRegistration(log chainTypes.Log) (*contracts.ControllerRegistryRegistration, error)
	ParseRegistrationFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ControllerRegistryRegistration, error)
	ParseUnregistration(log chainTypes.Log) (*contracts.ControllerRegistryUnregistration, error)
	ParseUnregistrationFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ControllerRegistryUnregistration, error)
}

type IControllerRegistryWatcher interface {
	WatchOwnershipTransferred(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ControllerRegistryOwnershipTransferred,
		previousOwner []common.Address,
		newOwner []common.Address,
	) (event.Subscription, error)
	WatchRegistration(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ControllerRegistryRegistration,
		controller []common.Address,
	) (event.Subscription, error)
	WatchUnregistration(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ControllerRegistryUnregistration,
		controller []common.Address,
	) (event.Subscription, error)
}

type IControllerRegistryContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IControllerRegistryCalls
	IControllerRegistryTransacts
	IControllerRegistryEvents
}

// Manager is contract wrapper struct
type ControllerRegistryContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	contracts.ControllerRegistryCaller
	contracts.ControllerRegistryFilterer
	contracts.ControllerRegistryTransactor
}

func NewControllerRegistryContract(deployment ablbind.Deployment, backend ablbind.ContractBackend) interface{} {
	if deployment.Address() == (common.Address{}) {
		evmABI, err := abi.JSON(strings.NewReader(contracts.ControllerRegistryABI))
		if err != nil {
			panic(err)
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(contracts.ControllerRegistryAddress),
			common.HexToHash(contracts.ControllerRegistryTxHash),
			new(big.Int).SetBytes(common.HexToHash(contracts.ControllerRegistryCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, backend)

	contract := &ControllerRegistryContract{
		Deployment: deployment,
		client:     backend,

		ControllerRegistryCaller:     contracts.NewControllerRegistryCaller(base),
		ControllerRegistryTransactor: contracts.NewControllerRegistryTransactor(base),
		ControllerRegistryFilterer:   contracts.NewControllerRegistryFilterer(base),
	}

	return contract
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address controller) constant returns(bool)
func (c *ControllerRegistryContract) Exists(
	controller common.Address,
) (

	bool,
	error,
) {
	return c.ControllerRegistryCaller.Exists(nil, controller)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
func (c *ControllerRegistryContract) Get(
	controller common.Address,
) (

	types.DataController,
	error,
) {
	return c.ControllerRegistryCaller.Get(nil, controller)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (c *ControllerRegistryContract) Register(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	controllerAddr common.Address,
) (*chainTypes.Receipt, error) {
	tx, err := c.ControllerRegistryTransactor.Register(c.client.Transactor(ctx, opts), controllerAddr)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
