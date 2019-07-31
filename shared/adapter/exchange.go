// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.NewKeyedTransactor
	_ = types.HexToID
	_ = common.Big1
	_ = ethTypes.BloomLookup
	_ = event.NewSubscription
)

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "{\"Constructor\":{\"Name\":\"\",\"Const\":false,\"Inputs\":[{\"Name\":\"appReg\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":null},\"Methods\":{\"addDataIds\":{\"Name\":\"addDataIds\",\"Const\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataIds\",\"Type\":{\"Elem\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":4,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"cancel\":{\"Name\":\"cancel\",\"Const\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"getOffer\":{\"Name\":\"getOffer\",\"Const\":true,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":25,\"Type\":{},\"Size\":0,\"T\":6,\"TupleElems\":[{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":4,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":22,\"Type\":{},\"Size\":256,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":22,\"Type\":{},\"Size\":256,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":25,\"Type\":{},\"Size\":0,\"T\":6,\"TupleElems\":[{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":4,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null}],\"TupleRawNames\":[\"addr\",\"sign\",\"args\"]},{\"Elem\":null,\"Kind\":8,\"Type\":{},\"Size\":8,\"T\":1,\"TupleElems\":null,\"TupleRawNames\":null}],\"TupleRawNames\":[\"provider\",\"consumer\",\"dataIds\",\"at\",\"until\",\"escrow\",\"status\"]},\"Indexed\":false}]},\"getOfferMembers\":{\"Name\":\"getOfferMembers\",\"Const\":true,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"offerExists\":{\"Name\":\"offerExists\",\"Const\":true,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":1,\"Type\":{},\"Size\":0,\"T\":2,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"order\":{\"Name\":\"order\",\"Const\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"prepare\":{\"Name\":\"prepare\",\"Const\":false,\"Inputs\":[{\"Name\":\"provider\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"consumer\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"escrow\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"escrowSign\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":4,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"escrowArgs\",\"Type\":{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"dataIds\",\"Type\":{\"Elem\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":4,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[{\"Name\":\"\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"reject\":{\"Name\":\"reject\",\"Const\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]},\"settle\":{\"Name\":\"settle\",\"Const\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}],\"Outputs\":[]}},\"Events\":{\"EscrowExecutionFailed\":{\"Name\":\"EscrowExecutionFailed\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"reason\",\"Type\":{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"OfferCanceled\":{\"Name\":\"OfferCanceled\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"providerAppName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"OfferPrepared\":{\"Name\":\"OfferPrepared\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"providerAppName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"OfferPresented\":{\"Name\":\"OfferPresented\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"providerAppName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"OfferReceipt\":{\"Name\":\"OfferReceipt\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"providerAppName\",\"Type\":{\"Elem\":null,\"Kind\":24,\"Type\":{},\"Size\":0,\"T\":3,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false},{\"Name\":\"consumer\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"result\",\"Type\":{\"Elem\":null,\"Kind\":23,\"Type\":{},\"Size\":0,\"T\":9,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":false}]},\"OfferRejected\":{\"Name\":\"OfferRejected\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"consumer\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true}]},\"OfferSettled\":{\"Name\":\"OfferSettled\",\"Anonymous\":false,\"Inputs\":[{\"Name\":\"offerId\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":8,\"T\":8,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true},{\"Name\":\"consumer\",\"Type\":{\"Elem\":null,\"Kind\":17,\"Type\":{},\"Size\":20,\"T\":7,\"TupleElems\":null,\"TupleRawNames\":null},\"Indexed\":true}]}}}"

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	address            common.Address
	txHash             common.Hash
	createdAt          *big.Int
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// Address is getter method of Exchange.address
func (_Exchange *Exchange) Address() common.Address {
	return _Exchange.address
}

// TxHash is getter method of Exchange.txHash
func (_Exchange *Exchange) TxHash() common.Hash {
	return _Exchange.txHash
}

// CreatedAt is getter method of Exchange.createdAt
func (_Exchange *Exchange) CreatedAt() *big.Int {
	return _Exchange.createdAt
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{
		address:            address,
		txHash:             txHash,
		createdAt:          createdAt,
		ExchangeCaller:     ExchangeCaller{contract: contract},
		ExchangeTransactor: ExchangeTransactor{contract: contract},
		ExchangeFilterer:   ExchangeFilterer{contract: contract},
	}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

//go:generate mockgen -source exchange.go -destination ./mocks/mock_exchange.go -package mocks IExchangeManager,IExchangeContract
type IExchangeManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	GetOffer(offerId types.ID) (types.Offer, error)
	GetOfferMembers(offerId types.ID) (common.Address, common.Address, error)
	OfferExists(offerId types.ID) (bool, error)

	// Transact methods
	AddDataIds(ctx context.Context, offerId types.ID, dataIds []types.DataId) error
	Cancel(ctx context.Context, offerId types.ID) error
	Order(ctx context.Context, offerId types.ID) error
	Prepare(ctx context.Context, provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (types.ID, error)
	Reject(ctx context.Context, offerId types.ID) error
	Settle(ctx context.Context, offerId types.ID) error

	FilterEscrowExecutionFailed(opts *bind.FilterOpts) (*ExchangeEscrowExecutionFailedIterator, error)
	WatchEscrowExecutionFailed(opts *bind.WatchOpts, sink chan<- *ExchangeEscrowExecutionFailed) (event.Subscription, error)

	FilterOfferCanceled(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferCanceledIterator, error)
	WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferCanceled, offerId []types.ID) (event.Subscription, error)

	FilterOfferPrepared(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPreparedIterator, error)
	WatchOfferPrepared(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPrepared, offerId []types.ID) (event.Subscription, error)

	FilterOfferPresented(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPresentedIterator, error)
	WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPresented, offerId []types.ID) (event.Subscription, error)

	FilterOfferReceipt(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferReceiptIterator, error)
	WatchOfferReceipt(opts *bind.WatchOpts, sink chan<- *ExchangeOfferReceipt, offerId []types.ID, consumer []common.Address) (event.Subscription, error)

	FilterOfferRejected(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferRejectedIterator, error)
	WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *ExchangeOfferRejected, offerId []types.ID, consumer []common.Address) (event.Subscription, error)

	FilterOfferSettled(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferSettledIterator, error)
	WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferSettled, offerId []types.ID, consumer []common.Address) (event.Subscription, error)
}

type IExchangeContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IExchangeCalls
	IExchangeTransacts
	IExchangeEvents
}

// Manager is contract wrapper struct
type ExchangeContract struct {
	client   blockchain.TxClient
	contract *Exchange
	ExchangeFilterer
}

// Address is getter method of Exchange.address
func (c *ExchangeContract) Address() common.Address {
	return c.contract.Address()
}

// TxHash is getter method of Exchange.txHash
func (c *ExchangeContract) TxHash() common.Hash {
	return c.contract.TxHash()
}

// CreatedAt is getter method of Exchange.createdAt
func (c *ExchangeContract) CreatedAt() *big.Int {
	return c.contract.CreatedAt()
}

// NewManager makes new *Manager struct
func NewExchangeContract(client blockchain.TxClient) IExchangeContract {
	contract := client.GetContract(&Exchange{}).(*Exchange)
	return &ExchangeContract{
		client:           client,
		contract:         contract,
		ExchangeFilterer: contract.ExchangeFilterer,
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("Exchange", (&Exchange{}).new)
	blockchain.RegisterSelector("0x367a9005", "addDataIds(bytes8,bytes20[])")
	blockchain.RegisterSelector("0xb2d9ba39", "cancel(bytes8)")
	blockchain.RegisterSelector("0x0cf833fb", "order(bytes8)")
	blockchain.RegisterSelector("0x77e61c33", "prepare(string,address,address,bytes4,bytes,bytes20[])")
	blockchain.RegisterSelector("0x6622e153", "reject(bytes8)")
	blockchain.RegisterSelector("0xa60d9b5f", "settle(bytes8)")
}

func (_Exchange *Exchange) new(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (interface{}, error) {
	return NewExchange(address, txHash, createdAt, backend)
}

type IExchangeCalls interface {
	GetOffer(offerId types.ID) (types.Offer, error)
	GetOfferMembers(offerId types.ID) (common.Address, common.Address, error)
	OfferExists(offerId types.ID) (bool, error)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (c *ExchangeContract) GetOffer(offerId types.ID) (types.Offer, error) {
	return c.contract.GetOffer(nil, offerId)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (_Exchange *ExchangeCaller) GetOffer(opts *bind.CallOpts, offerId types.ID) (types.Offer, error) {
	ret := new(types.Offer)

	out := ret
	err := _Exchange.contract.Call(opts, out, "getOffer", offerId)
	return *ret, err
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (_Exchange *ExchangeSession) GetOffer(offerId types.ID) (types.Offer, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, offerId)
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (_Exchange *ExchangeCallerSession) GetOffer(offerId types.ID) (types.Offer, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, offerId)
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (c *ExchangeContract) GetOfferMembers(offerId types.ID) (common.Address, common.Address, error) {
	return c.contract.GetOfferMembers(nil, offerId)
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (_Exchange *ExchangeCaller) GetOfferMembers(opts *bind.CallOpts, offerId types.ID) (common.Address, common.Address, error) {
	var (
		ret0, ret1 = new(common.Address), new(common.Address)
	)
	out := &[]interface{}{ret0, ret1}
	err := _Exchange.contract.Call(opts, out, "getOfferMembers", offerId)
	return *ret0, *ret1, err
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (_Exchange *ExchangeSession) GetOfferMembers(offerId types.ID) (common.Address, common.Address, error) {
	return _Exchange.Contract.GetOfferMembers(&_Exchange.CallOpts, offerId)
}

// GetOfferMembers is a free data retrieval call binding the contract method 0x72dfa465.
//
// Solidity: function getOfferMembers(bytes8 offerId) constant returns(address, address)
func (_Exchange *ExchangeCallerSession) GetOfferMembers(offerId types.ID) (common.Address, common.Address, error) {
	return _Exchange.Contract.GetOfferMembers(&_Exchange.CallOpts, offerId)
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (c *ExchangeContract) OfferExists(offerId types.ID) (bool, error) {
	return c.contract.OfferExists(nil, offerId)
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (_Exchange *ExchangeCaller) OfferExists(opts *bind.CallOpts, offerId types.ID) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{ret0}
	err := _Exchange.contract.Call(opts, out, "offerExists", offerId)
	return *ret0, err
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (_Exchange *ExchangeSession) OfferExists(offerId types.ID) (bool, error) {
	return _Exchange.Contract.OfferExists(&_Exchange.CallOpts, offerId)
}

// OfferExists is a free data retrieval call binding the contract method 0xc4a03da9.
//
// Solidity: function offerExists(bytes8 offerId) constant returns(bool)
func (_Exchange *ExchangeCallerSession) OfferExists(offerId types.ID) (bool, error) {
	return _Exchange.Contract.OfferExists(&_Exchange.CallOpts, offerId)
}

type IExchangeTransacts interface {
	AddDataIds(ctx context.Context, offerId types.ID, dataIds []types.DataId) (*ethTypes.Receipt, error)
	Cancel(ctx context.Context, offerId types.ID) (*ethTypes.Receipt, error)
	Order(ctx context.Context, offerId types.ID) (*ethTypes.Receipt, error)
	Prepare(ctx context.Context, provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*ethTypes.Receipt, error)
	Reject(ctx context.Context, offerId types.ID) (*ethTypes.Receipt, error)
	Settle(ctx context.Context, offerId types.ID) (*ethTypes.Receipt, error)
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (c *ExchangeContract) AddDataIds(ctx context.Context, offerId types.ID, dataIds []types.DataId) (*ethTypes.Receipt, error) {
	tx, err := c.contract.AddDataIds(c.client.Account(), offerId, dataIds)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (_Exchange *ExchangeTransactor) AddDataIds(opts *bind.TransactOpts, offerId types.ID, dataIds []types.DataId) (*ethTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addDataIds", offerId, dataIds)
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (_Exchange *ExchangeSession) AddDataIds(offerId types.ID, dataIds []types.DataId) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.AddDataIds(&_Exchange.TransactOpts, offerId, dataIds)
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (_Exchange *ExchangeTransactorSession) AddDataIds(offerId types.ID, dataIds []types.DataId) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.AddDataIds(&_Exchange.TransactOpts, offerId, dataIds)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (c *ExchangeContract) Cancel(ctx context.Context, offerId types.ID) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Cancel(c.client.Account(), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactor) Cancel(opts *bind.TransactOpts, offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancel", offerId)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (_Exchange *ExchangeSession) Cancel(offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, offerId)
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactorSession) Cancel(offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (c *ExchangeContract) Order(ctx context.Context, offerId types.ID) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Order(c.client.Account(), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactor) Order(opts *bind.TransactOpts, offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "order", offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (_Exchange *ExchangeSession) Order(offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, offerId)
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactorSession) Order(offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, offerId)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (c *ExchangeContract) Prepare(ctx context.Context, provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Prepare(c.client.Account(), provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (_Exchange *ExchangeTransactor) Prepare(opts *bind.TransactOpts, provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*ethTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "prepare", provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (_Exchange *ExchangeSession) Prepare(provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Prepare(&_Exchange.TransactOpts, provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (_Exchange *ExchangeTransactorSession) Prepare(provider string, consumer common.Address, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Prepare(&_Exchange.TransactOpts, provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (c *ExchangeContract) Reject(ctx context.Context, offerId types.ID) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Reject(c.client.Account(), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactor) Reject(opts *bind.TransactOpts, offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "reject", offerId)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (_Exchange *ExchangeSession) Reject(offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Reject(&_Exchange.TransactOpts, offerId)
}

// Reject is a paid mutator transaction binding the contract method 0x6622e153.
//
// Solidity: function reject(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactorSession) Reject(offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Reject(&_Exchange.TransactOpts, offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (c *ExchangeContract) Settle(ctx context.Context, offerId types.ID) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Settle(c.client.Account(), offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactor) Settle(opts *bind.TransactOpts, offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.contract.Transact(opts, "settle", offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (_Exchange *ExchangeSession) Settle(offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Settle(&_Exchange.TransactOpts, offerId)
}

// Settle is a paid mutator transaction binding the contract method 0xa60d9b5f.
//
// Solidity: function settle(bytes8 offerId) returns()
func (_Exchange *ExchangeTransactorSession) Settle(offerId types.ID) (*ethTypes.Transaction, error) {
	return _Exchange.Contract.Settle(&_Exchange.TransactOpts, offerId)
}

type IExchangeEvents interface {
	FilterEscrowExecutionFailed(opts *bind.FilterOpts) (*ExchangeEscrowExecutionFailedIterator, error)
	ParseEscrowExecutionFailedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeEscrowExecutionFailed, error)
	WatchEscrowExecutionFailed(opts *bind.WatchOpts, sink chan<- *ExchangeEscrowExecutionFailed) (event.Subscription, error)

	FilterOfferCanceled(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferCanceledIterator, error)
	ParseOfferCanceledFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferCanceled, error)
	WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferCanceled, offerId []types.ID) (event.Subscription, error)

	FilterOfferPrepared(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPreparedIterator, error)
	ParseOfferPreparedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferPrepared, error)
	WatchOfferPrepared(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPrepared, offerId []types.ID) (event.Subscription, error)

	FilterOfferPresented(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPresentedIterator, error)
	ParseOfferPresentedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferPresented, error)
	WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPresented, offerId []types.ID) (event.Subscription, error)

	FilterOfferReceipt(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferReceiptIterator, error)
	ParseOfferReceiptFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferReceipt, error)
	WatchOfferReceipt(opts *bind.WatchOpts, sink chan<- *ExchangeOfferReceipt, offerId []types.ID, consumer []common.Address) (event.Subscription, error)

	FilterOfferRejected(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferRejectedIterator, error)
	ParseOfferRejectedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferRejected, error)
	WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *ExchangeOfferRejected, offerId []types.ID, consumer []common.Address) (event.Subscription, error)

	FilterOfferSettled(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferSettledIterator, error)
	ParseOfferSettledFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferSettled, error)
	WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferSettled, offerId []types.ID, consumer []common.Address) (event.Subscription, error)
}

// ExchangeEscrowExecutionFailedIterator is returned from FilterEscrowExecutionFailed and is used to iterate over the raw logs and unpacked data for EscrowExecutionFailed events raised by the Exchange contract.
type ExchangeEscrowExecutionFailedIterator struct {
	Event *ExchangeEscrowExecutionFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeEscrowExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEscrowExecutionFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeEscrowExecutionFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeEscrowExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEscrowExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEscrowExecutionFailed represents a EscrowExecutionFailed event raised by the Exchange contract.
type ExchangeEscrowExecutionFailed struct {
	Reason []byte
	Raw    ethTypes.Log // Blockchain specific contextual infos
}

// FilterEscrowExecutionFailed is a free log retrieval operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *ExchangeFilterer) FilterEscrowExecutionFailed(opts *bind.FilterOpts) (*ExchangeEscrowExecutionFailedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EscrowExecutionFailed")
	if err != nil {
		return nil, err
	}
	return &ExchangeEscrowExecutionFailedIterator{contract: _Exchange.contract, event: "EscrowExecutionFailed", logs: logs, sub: sub}, nil
}

// FilterEscrowExecutionFailed parses the event from given transaction receipt.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (manager *ExchangeContract) ParseEscrowExecutionFailedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeEscrowExecutionFailed, error) {
	return manager.contract.ParseEscrowExecutionFailedFromReceipt(receipt)
}

// FilterEscrowExecutionFailed parses the event from given transaction receipt.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *ExchangeFilterer) ParseEscrowExecutionFailedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeEscrowExecutionFailed, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7") {
			event := new(ExchangeEscrowExecutionFailed)
			if err := _Exchange.contract.UnpackLog(event, "EscrowExecutionFailed", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("EscrowExecutionFailed event not found")
}

// WatchEscrowExecutionFailed is a free log subscription operation binding the contract event 0x40e7fa7728ad0189a69a1f7d9b3b202f751810b2be48db0b9224d7f81cd232f7.
//
// Solidity: event EscrowExecutionFailed(bytes reason)
func (_Exchange *ExchangeFilterer) WatchEscrowExecutionFailed(opts *bind.WatchOpts, sink chan<- *ExchangeEscrowExecutionFailed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EscrowExecutionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEscrowExecutionFailed)
				if err := _Exchange.contract.UnpackLog(event, "EscrowExecutionFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the Exchange contract.
type ExchangeOfferCanceledIterator struct {
	Event *ExchangeOfferCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeOfferCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferCanceled represents a OfferCanceled event raised by the Exchange contract.
type ExchangeOfferCanceled struct {
	OfferId         types.ID
	ProviderAppName string
	Raw             ethTypes.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) FilterOfferCanceled(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferCanceledIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferCanceled", offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferCanceledIterator{contract: _Exchange.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// FilterOfferCanceled parses the event from given transaction receipt.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (manager *ExchangeContract) ParseOfferCanceledFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferCanceled, error) {
	return manager.contract.ParseOfferCanceledFromReceipt(receipt)
}

// FilterOfferCanceled parses the event from given transaction receipt.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferCanceledFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferCanceled, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a") {
			event := new(ExchangeOfferCanceled)
			if err := _Exchange.contract.UnpackLog(event, "OfferCanceled", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferCanceled event not found")
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0x05b47b0f8bd37a836f7a5c080cb883841c1282c69dd1874a46d4fafc7e8aa27a.
//
// Solidity: event OfferCanceled(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferCanceled, offerId []types.ID) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferCanceled", offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferCanceled)
				if err := _Exchange.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeOfferPreparedIterator is returned from FilterOfferPrepared and is used to iterate over the raw logs and unpacked data for OfferPrepared events raised by the Exchange contract.
type ExchangeOfferPreparedIterator struct {
	Event *ExchangeOfferPrepared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferPrepared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeOfferPrepared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferPrepared represents a OfferPrepared event raised by the Exchange contract.
type ExchangeOfferPrepared struct {
	OfferId         types.ID
	ProviderAppName string
	Raw             ethTypes.Log // Blockchain specific contextual infos
}

// FilterOfferPrepared is a free log retrieval operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) FilterOfferPrepared(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPreparedIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferPrepared", offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferPreparedIterator{contract: _Exchange.contract, event: "OfferPrepared", logs: logs, sub: sub}, nil
}

// FilterOfferPrepared parses the event from given transaction receipt.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (manager *ExchangeContract) ParseOfferPreparedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferPrepared, error) {
	return manager.contract.ParseOfferPreparedFromReceipt(receipt)
}

// FilterOfferPrepared parses the event from given transaction receipt.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferPreparedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferPrepared, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5") {
			event := new(ExchangeOfferPrepared)
			if err := _Exchange.contract.UnpackLog(event, "OfferPrepared", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferPrepared event not found")
}

// WatchOfferPrepared is a free log subscription operation binding the contract event 0x821d45f3b8db50a4777ad807928db085f0c986433cf51c2afdc8c6af90d1aef5.
//
// Solidity: event OfferPrepared(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) WatchOfferPrepared(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPrepared, offerId []types.ID) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferPrepared", offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferPrepared)
				if err := _Exchange.contract.UnpackLog(event, "OfferPrepared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeOfferPresentedIterator is returned from FilterOfferPresented and is used to iterate over the raw logs and unpacked data for OfferPresented events raised by the Exchange contract.
type ExchangeOfferPresentedIterator struct {
	Event *ExchangeOfferPresented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferPresentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferPresented)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeOfferPresented)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferPresentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferPresentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferPresented represents a OfferPresented event raised by the Exchange contract.
type ExchangeOfferPresented struct {
	OfferId         types.ID
	ProviderAppName string
	Raw             ethTypes.Log // Blockchain specific contextual infos
}

// FilterOfferPresented is a free log retrieval operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) FilterOfferPresented(opts *bind.FilterOpts, offerId []types.ID) (*ExchangeOfferPresentedIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferPresented", offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferPresentedIterator{contract: _Exchange.contract, event: "OfferPresented", logs: logs, sub: sub}, nil
}

// FilterOfferPresented parses the event from given transaction receipt.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (manager *ExchangeContract) ParseOfferPresentedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferPresented, error) {
	return manager.contract.ParseOfferPresentedFromReceipt(receipt)
}

// FilterOfferPresented parses the event from given transaction receipt.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) ParseOfferPresentedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferPresented, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa") {
			event := new(ExchangeOfferPresented)
			if err := _Exchange.contract.UnpackLog(event, "OfferPresented", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferPresented event not found")
}

// WatchOfferPresented is a free log subscription operation binding the contract event 0x198eb5e3b4b2cd8cca381c07c5696b7caffe2c775d93f75d0053073e36a865fa.
//
// Solidity: event OfferPresented(bytes8 indexed offerId, string providerAppName)
func (_Exchange *ExchangeFilterer) WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPresented, offerId []types.ID) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferPresented", offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferPresented)
				if err := _Exchange.contract.UnpackLog(event, "OfferPresented", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeOfferReceiptIterator is returned from FilterOfferReceipt and is used to iterate over the raw logs and unpacked data for OfferReceipt events raised by the Exchange contract.
type ExchangeOfferReceiptIterator struct {
	Event *ExchangeOfferReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeOfferReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferReceipt represents a OfferReceipt event raised by the Exchange contract.
type ExchangeOfferReceipt struct {
	OfferId         types.ID
	ProviderAppName string
	Consumer        common.Address
	Result          []byte
	Raw             ethTypes.Log // Blockchain specific contextual infos
}

// FilterOfferReceipt is a free log retrieval operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *ExchangeFilterer) FilterOfferReceipt(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferReceiptIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferReceipt", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferReceiptIterator{contract: _Exchange.contract, event: "OfferReceipt", logs: logs, sub: sub}, nil
}

// FilterOfferReceipt parses the event from given transaction receipt.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (manager *ExchangeContract) ParseOfferReceiptFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferReceipt, error) {
	return manager.contract.ParseOfferReceiptFromReceipt(receipt)
}

// FilterOfferReceipt parses the event from given transaction receipt.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *ExchangeFilterer) ParseOfferReceiptFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferReceipt, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa") {
			event := new(ExchangeOfferReceipt)
			if err := _Exchange.contract.UnpackLog(event, "OfferReceipt", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferReceipt event not found")
}

// WatchOfferReceipt is a free log subscription operation binding the contract event 0x7a2b40d55d10a35fd97231e1d36fc9df7c48361f16299086103e0712135c59fa.
//
// Solidity: event OfferReceipt(bytes8 indexed offerId, string providerAppName, address indexed consumer, bytes result)
func (_Exchange *ExchangeFilterer) WatchOfferReceipt(opts *bind.WatchOpts, sink chan<- *ExchangeOfferReceipt, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferReceipt", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferReceipt)
				if err := _Exchange.contract.UnpackLog(event, "OfferReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeOfferRejectedIterator is returned from FilterOfferRejected and is used to iterate over the raw logs and unpacked data for OfferRejected events raised by the Exchange contract.
type ExchangeOfferRejectedIterator struct {
	Event *ExchangeOfferRejected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferRejected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeOfferRejected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferRejected represents a OfferRejected event raised by the Exchange contract.
type ExchangeOfferRejected struct {
	OfferId  types.ID
	Consumer common.Address
	Raw      ethTypes.Log // Blockchain specific contextual infos
}

// FilterOfferRejected is a free log retrieval operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) FilterOfferRejected(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferRejectedIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferRejected", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferRejectedIterator{contract: _Exchange.contract, event: "OfferRejected", logs: logs, sub: sub}, nil
}

// FilterOfferRejected parses the event from given transaction receipt.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (manager *ExchangeContract) ParseOfferRejectedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferRejected, error) {
	return manager.contract.ParseOfferRejectedFromReceipt(receipt)
}

// FilterOfferRejected parses the event from given transaction receipt.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) ParseOfferRejectedFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferRejected, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843") {
			event := new(ExchangeOfferRejected)
			if err := _Exchange.contract.UnpackLog(event, "OfferRejected", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferRejected event not found")
}

// WatchOfferRejected is a free log subscription operation binding the contract event 0x94c89cb0104a1fa8726bf8a9e9151423d67ff6f8eb09ed7392386649655c6843.
//
// Solidity: event OfferRejected(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *ExchangeOfferRejected, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferRejected", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferRejected)
				if err := _Exchange.contract.UnpackLog(event, "OfferRejected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeOfferSettledIterator is returned from FilterOfferSettled and is used to iterate over the raw logs and unpacked data for OfferSettled events raised by the Exchange contract.
type ExchangeOfferSettledIterator struct {
	Event *ExchangeOfferSettled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan ethTypes.Log     // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeOfferSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferSettled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeOfferSettled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeOfferSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferSettled represents a OfferSettled event raised by the Exchange contract.
type ExchangeOfferSettled struct {
	OfferId  types.ID
	Consumer common.Address
	Raw      ethTypes.Log // Blockchain specific contextual infos
}

// FilterOfferSettled is a free log retrieval operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) FilterOfferSettled(opts *bind.FilterOpts, offerId []types.ID, consumer []common.Address) (*ExchangeOfferSettledIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferSettled", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferSettledIterator{contract: _Exchange.contract, event: "OfferSettled", logs: logs, sub: sub}, nil
}

// FilterOfferSettled parses the event from given transaction receipt.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (manager *ExchangeContract) ParseOfferSettledFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferSettled, error) {
	return manager.contract.ParseOfferSettledFromReceipt(receipt)
}

// FilterOfferSettled parses the event from given transaction receipt.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) ParseOfferSettledFromReceipt(receipt *ethTypes.Receipt) (*ExchangeOfferSettled, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e") {
			event := new(ExchangeOfferSettled)
			if err := _Exchange.contract.UnpackLog(event, "OfferSettled", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OfferSettled event not found")
}

// WatchOfferSettled is a free log subscription operation binding the contract event 0xb37cb3a83f4f40ee469256bdfc4a2881c9ce188960c87bf11359151a461b723e.
//
// Solidity: event OfferSettled(bytes8 indexed offerId, address indexed consumer)
func (_Exchange *ExchangeFilterer) WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferSettled, offerId []types.ID, consumer []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferSettled", offerIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferSettled)
				if err := _Exchange.contract.UnpackLog(event, "OfferSettled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
