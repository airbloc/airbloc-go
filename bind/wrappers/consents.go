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

//go:generate mockgen -source consents.go -destination ./mocks/mock_consents.go -package mocks IConsentsContract

type IConsentsCalls interface {
	IsAllowed(
		userId types.ID,
		appName string,
		action uint8,
		dataType string,
	) (
		bool,
		error,
	)
	IsAllowedAt(
		userId types.ID,
		appName string,
		action uint8,
		dataType string,
		blockNumber *big.Int,
	) (
		bool,
		error,
	)
}

type IConsentsTransacts interface {
	Consent(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
		consentData types.ConsentData,
	) (*chainTypes.Receipt, error)
	ConsentByController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData types.ConsentData,
	) (*chainTypes.Receipt, error)
	ConsentMany(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		appName string,
		consentData []types.ConsentData,
	) (*chainTypes.Receipt, error)
	ConsentManyByController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData []types.ConsentData,
	) (*chainTypes.Receipt, error)
	ModifyConsentByController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData types.ConsentData,
		passwordSignature []byte,
	) (*chainTypes.Receipt, error)
	ModifyConsentManyByController(
		ctx context.Context,
		opts *ablbind.TransactOpts,
		userId types.ID,
		appName string,
		consentData []types.ConsentData,
		passwordSignature []byte,
	) (*chainTypes.Receipt, error)
}

type IConsentsEvents interface {
	IConsentsFilterer
	IConsentsParser
	IConsentsWatcher
}

type IConsentsFilterer interface {
	FilterConsented(
		opts *bind.FilterOpts,
		action []uint8,
		userId []types.ID,
		appAddr []common.Address,

	) (ablbind.EventIterator, error)
}

type IConsentsParser interface {
	ParseConsented(log chainTypes.Log) (*contracts.ConsentsConsented, error)
	ParseConsentedFromReceipt(receipt *chainTypes.Receipt) ([]*contracts.ConsentsConsented, error)
}

type IConsentsWatcher interface {
	WatchConsented(
		opts *bind.WatchOpts,
		sink chan<- *contracts.ConsentsConsented,
		action []uint8,
		userId []types.ID,
		appAddr []common.Address,

	) (event.Subscription, error)
}

type IConsentsContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IConsentsCalls
	IConsentsTransacts
	IConsentsEvents
}

// Manager is contract wrapper struct
type ConsentsContract struct {
	ablbind.Deployment
	client ablbind.ContractBackend

	contracts.ConsentsCaller
	contracts.ConsentsFilterer
	contracts.ConsentsTransactor
}

func NewConsentsContract(deployment ablbind.Deployment, backend ablbind.ContractBackend) interface{} {
	if deployment.Address() == (common.Address{}) {
		evmABI, err := abi.JSON(strings.NewReader(contracts.ConsentsABI))
		if err != nil {
			panic(err)
		}

		deployment = ablbind.NewDeployment(
			common.HexToAddress(contracts.ConsentsAddress),
			common.HexToHash(contracts.ConsentsTxHash),
			new(big.Int).SetBytes(common.HexToHash(contracts.ConsentsCreatedAt).Bytes()),
			evmABI,
		)
	}

	base := ablbind.NewBoundContract(deployment.Address(), deployment.ParsedABI, backend)

	contract := &ConsentsContract{
		Deployment: deployment,
		client:     backend,

		ConsentsCaller:     contracts.NewConsentsCaller(base),
		ConsentsTransactor: contracts.NewConsentsTransactor(base),
		ConsentsFilterer:   contracts.NewConsentsFilterer(base),
	}

	return contract
}

func (c *ConsentsContract) GetSelectors() map[string]string {
	selectors := make(map[string]string)
	selectors["0x50615985"] = "isAllowed(bytes8,string,uint8,string)"
	selectors["0x7cdda67c"] = "isAllowedAt(bytes8,string,uint8,string,uint256)"

	selectors["0xcd4dc804"] = "consent(string,(uint8,string,bool))"
	selectors["0xf573f89a"] = "consentByController(bytes8,string,(uint8,string,bool))"
	selectors["0xdd43ad05"] = "consentMany(string,(uint8,string,bool)[])"
	selectors["0xae6d5034"] = "consentManyByController(bytes8,string,(uint8,string,bool)[])"
	selectors["0x0bfec389"] = "modifyConsentByController(bytes8,string,(uint8,string,bool),bytes)"
	selectors["0xe031b1cf"] = "modifyConsentManyByController(bytes8,string,(uint8,string,bool)[],bytes)"

	return selectors
}

// IsAllowed is a free data retrieval call binding the contract method 0x50615985.
//
// Solidity: function isAllowed(bytes8 userId, string appName, uint8 action, string dataType) constant returns(bool)
func (c *ConsentsContract) IsAllowed(
	userId types.ID,
	appName string,
	action uint8,
	dataType string,
) (

	bool,
	error,
) {
	return c.ConsentsCaller.IsAllowed(nil, userId, appName, action, dataType)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x7cdda67c.
//
// Solidity: function isAllowedAt(bytes8 userId, string appName, uint8 action, string dataType, uint256 blockNumber) constant returns(bool)
func (c *ConsentsContract) IsAllowedAt(
	userId types.ID,
	appName string,
	action uint8,
	dataType string,
	blockNumber *big.Int,
) (

	bool,
	error,
) {
	return c.ConsentsCaller.IsAllowedAt(nil, userId, appName, action, dataType, blockNumber)
}

// Consent is a paid mutator transaction binding the contract method 0xcd4dc804.
//
// Solidity: function consent(string appName, (uint8,string,bool) consentData) returns()
func (c *ConsentsContract) Consent(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
	consentData types.ConsentData,
) (*chainTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.Consent(c.client.Transactor(ctx, opts), appName, consentData)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf573f89a.
//
// Solidity: function consentByController(bytes8 userId, string appName, (uint8,string,bool) consentData) returns()
func (c *ConsentsContract) ConsentByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
) (*chainTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ConsentByController(c.client.Transactor(ctx, opts), userId, appName, consentData)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ConsentMany is a paid mutator transaction binding the contract method 0xdd43ad05.
//
// Solidity: function consentMany(string appName, (uint8,string,bool)[] consentData) returns()
func (c *ConsentsContract) ConsentMany(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	appName string,
	consentData []types.ConsentData,
) (*chainTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ConsentMany(c.client.Transactor(ctx, opts), appName, consentData)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ConsentManyByController is a paid mutator transaction binding the contract method 0xae6d5034.
//
// Solidity: function consentManyByController(bytes8 userId, string appName, (uint8,string,bool)[] consentData) returns()
func (c *ConsentsContract) ConsentManyByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData []types.ConsentData,
) (*chainTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ConsentManyByController(c.client.Transactor(ctx, opts), userId, appName, consentData)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ModifyConsentByController is a paid mutator transaction binding the contract method 0x0bfec389.
//
// Solidity: function modifyConsentByController(bytes8 userId, string appName, (uint8,string,bool) consentData, bytes passwordSignature) returns()
func (c *ConsentsContract) ModifyConsentByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData types.ConsentData,
	passwordSignature []byte,
) (*chainTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ModifyConsentByController(c.client.Transactor(ctx, opts), userId, appName, consentData, passwordSignature)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// ModifyConsentManyByController is a paid mutator transaction binding the contract method 0xe031b1cf.
//
// Solidity: function modifyConsentManyByController(bytes8 userId, string appName, (uint8,string,bool)[] consentData, bytes passwordSignature) returns()
func (c *ConsentsContract) ModifyConsentManyByController(
	ctx context.Context,
	opts *ablbind.TransactOpts,
	userId types.ID,
	appName string,
	consentData []types.ConsentData,
	passwordSignature []byte,
) (*chainTypes.Receipt, error) {
	tx, err := c.ConsentsTransactor.ModifyConsentManyByController(c.client.Transactor(ctx, opts), userId, appName, consentData, passwordSignature)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
