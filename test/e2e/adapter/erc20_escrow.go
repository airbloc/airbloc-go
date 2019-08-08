// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testadapter

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

// ERC20EscrowABI is the input ABI used to generate the binding from.
const ERC20EscrowABI = "[{\"inputs\":[{\"name\":\"exchangeContract\",\"type\":\"address\"}],\"payable\":false,\"signature\":\"constructor\",\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"sign\",\"type\":\"bytes4\"},{\"name\":\"args\",\"type\":\"bytes\"},{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"signature\":\"0xf8411fa9\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"offerId\",\"type\":\"bytes8\"}],\"name\":\"transact\",\"outputs\":[],\"payable\":false,\"signature\":\"0x0bd9e0f8\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransactSelector\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"signature\":\"0xc0a79b5b\",\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ERC20Escrow is an auto generated Go binding around an Ethereum contract.
type ERC20Escrow struct {
	address               common.Address
	txHash                common.Hash
	createdAt             *big.Int
	ERC20EscrowCaller     // Read-only binding to the contract
	ERC20EscrowTransactor // Write-only binding to the contract
	ERC20EscrowFilterer   // Log filterer for contract events
}

// Address is getter method of ERC20Escrow.address
func (_ERC20Escrow *ERC20Escrow) Address() common.Address {
	return _ERC20Escrow.address
}

// TxHash is getter method of ERC20Escrow.txHash
func (_ERC20Escrow *ERC20Escrow) TxHash() common.Hash {
	return _ERC20Escrow.txHash
}

// CreatedAt is getter method of ERC20Escrow.createdAt
func (_ERC20Escrow *ERC20Escrow) CreatedAt() *big.Int {
	return _ERC20Escrow.createdAt
}

// ERC20EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20EscrowSession struct {
	Contract     *ERC20Escrow      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20EscrowRaw struct {
	Contract *ERC20Escrow // Generic contract binding to access the raw methods on
}

// NewERC20Escrow creates a new instance of ERC20Escrow, bound to a specific deployed contract.
func NewERC20Escrow(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (*ERC20Escrow, error) {
	contract, err := bindERC20Escrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Escrow{
		address:               address,
		txHash:                txHash,
		createdAt:             createdAt,
		ERC20EscrowCaller:     ERC20EscrowCaller{contract: contract},
		ERC20EscrowTransactor: ERC20EscrowTransactor{contract: contract},
		ERC20EscrowFilterer:   ERC20EscrowFilterer{contract: contract},
	}, nil
}

// bindERC20Escrow binds a generic wrapper to an already deployed contract.
func bindERC20Escrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Escrow *ERC20EscrowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Escrow.Contract.ERC20EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Escrow *ERC20EscrowRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.ERC20EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Escrow *ERC20EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.ERC20EscrowTransactor.contract.Transact(opts, method, params...)
}

// ERC20EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20EscrowCallerSession struct {
	Contract *ERC20EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20EscrowCallerRaw struct {
	Contract *ERC20EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// NewERC20EscrowCaller creates a new read-only instance of ERC20Escrow, bound to a specific deployed contract.
func NewERC20EscrowCaller(address common.Address, caller bind.ContractCaller) (*ERC20EscrowCaller, error) {
	contract, err := bindERC20Escrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20EscrowCaller{contract: contract}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Escrow *ERC20EscrowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Escrow.Contract.contract.Call(opts, result, method, params...)
}

// ERC20EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20EscrowTransactorSession struct {
	Contract     *ERC20EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20EscrowTransactorRaw struct {
	Contract *ERC20EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20EscrowTransactor creates a new write-only instance of ERC20Escrow, bound to a specific deployed contract.
func NewERC20EscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20EscrowTransactor, error) {
	contract, err := bindERC20Escrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20EscrowTransactor{contract: contract}, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Escrow *ERC20EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Escrow *ERC20EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.contract.Transact(opts, method, params...)
}

// ERC20EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewERC20EscrowFilterer creates a new log filterer instance of ERC20Escrow, bound to a specific deployed contract.
func NewERC20EscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20EscrowFilterer, error) {
	contract, err := bindERC20Escrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20EscrowFilterer{contract: contract}, nil
}

//go:generate mockgen -source erc20_escrow.go -destination ./mocks/mock_erc20_escrow.go -package mocks IERC20EscrowManager,IERC20EscrowContract
type IERC20EscrowManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error)
	GetTransactSelector() ([4]byte, error)

	// Transact methods
	Transact(ctx context.Context, token common.Address, amount *big.Int, offerId types.ID) error
}

type IERC20EscrowContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int
	Filterer() ERC20EscrowFilterer

	IERC20EscrowCalls
	IERC20EscrowTransacts
	IERC20EscrowEvents
}

// ERC20EscrowContract is contract wrapper struct
type ERC20EscrowContract struct {
	client   blockchain.TxClient
	contract *ERC20Escrow
	ERC20EscrowFilterer
}

// Address is getter method of ERC20Escrow.address
func (c *ERC20EscrowContract) Address() common.Address {
	return c.contract.Address()
}

// TxHash is getter method of ERC20Escrow.txHash
func (c *ERC20EscrowContract) TxHash() common.Hash {
	return c.contract.TxHash()
}

// CreatedAt is getter method of ERC20Escrow.createdAt
func (c *ERC20EscrowContract) CreatedAt() *big.Int {
	return c.contract.CreatedAt()
}

// Filterer is getter method of ERC20Escrow.ERC20EscrowFilterer
func (c *ERC20EscrowContract) Filterer() ERC20EscrowFilterer {
	return c.ERC20EscrowFilterer
}

// NewERC20EscrowContract makes new *ERC20EscrowContract struct
func NewERC20EscrowContract(client blockchain.TxClient) IERC20EscrowContract {
	contract := client.GetContract(&ERC20Escrow{}).(*ERC20Escrow)
	return &ERC20EscrowContract{
		client:              client,
		contract:            contract,
		ERC20EscrowFilterer: contract.ERC20EscrowFilterer,
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("ERC20Escrow", (&ERC20Escrow{}).new)
	blockchain.RegisterSelector("0x0bd9e0f8", "transact(address,uint256,bytes8)")
}

func (_ERC20Escrow *ERC20Escrow) new(address common.Address, txHash common.Hash, createdAt *big.Int, backend bind.ContractBackend) (interface{}, error) {
	return NewERC20Escrow(address, txHash, createdAt, backend)
}

type IERC20EscrowCalls interface {
	Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error)
	GetTransactSelector() ([4]byte, error)
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (c *ERC20EscrowContract) Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error) {
	return c.contract.Convert(nil, sign, args, offerId)
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (_ERC20Escrow *ERC20EscrowCaller) Convert(opts *bind.CallOpts, sign [4]byte, args []byte, offerId types.ID) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := &[]interface{}{ret0}
	err := _ERC20Escrow.contract.Call(opts, out, "convert", sign, args, offerId)
	return *ret0, err
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (_ERC20Escrow *ERC20EscrowSession) Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error) {
	return _ERC20Escrow.Contract.Convert(&_ERC20Escrow.CallOpts, sign, args, offerId)
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (_ERC20Escrow *ERC20EscrowCallerSession) Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error) {
	return _ERC20Escrow.Contract.Convert(&_ERC20Escrow.CallOpts, sign, args, offerId)
}

// GetTransactSelector is a free data retrieval call binding the contract method 0xc0a79b5b.
//
// Solidity: function getTransactSelector() constant returns(bytes4)
func (c *ERC20EscrowContract) GetTransactSelector() ([4]byte, error) {
	return c.contract.GetTransactSelector(nil)
}

// GetTransactSelector is a free data retrieval call binding the contract method 0xc0a79b5b.
//
// Solidity: function getTransactSelector() constant returns(bytes4)
func (_ERC20Escrow *ERC20EscrowCaller) GetTransactSelector(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := &[]interface{}{ret0}
	err := _ERC20Escrow.contract.Call(opts, out, "getTransactSelector")
	return *ret0, err
}

// GetTransactSelector is a free data retrieval call binding the contract method 0xc0a79b5b.
//
// Solidity: function getTransactSelector() constant returns(bytes4)
func (_ERC20Escrow *ERC20EscrowSession) GetTransactSelector() ([4]byte, error) {
	return _ERC20Escrow.Contract.GetTransactSelector(&_ERC20Escrow.CallOpts)
}

// GetTransactSelector is a free data retrieval call binding the contract method 0xc0a79b5b.
//
// Solidity: function getTransactSelector() constant returns(bytes4)
func (_ERC20Escrow *ERC20EscrowCallerSession) GetTransactSelector() ([4]byte, error) {
	return _ERC20Escrow.Contract.GetTransactSelector(&_ERC20Escrow.CallOpts)
}

type IERC20EscrowTransacts interface {
	Transact(ctx context.Context, token common.Address, amount *big.Int, offerId types.ID) (*ethTypes.Receipt, error)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (c *ERC20EscrowContract) Transact(ctx context.Context, token common.Address, amount *big.Int, offerId types.ID) (*ethTypes.Receipt, error) {
	tx, err := c.contract.Transact(c.client.Account(), token, amount, offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (_ERC20Escrow *ERC20EscrowTransactor) Transact(opts *bind.TransactOpts, token common.Address, amount *big.Int, offerId types.ID) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.contract.Transact(opts, "transact", token, amount, offerId)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (_ERC20Escrow *ERC20EscrowSession) Transact(token common.Address, amount *big.Int, offerId types.ID) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.Transact(&_ERC20Escrow.TransactOpts, token, amount, offerId)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (_ERC20Escrow *ERC20EscrowTransactorSession) Transact(token common.Address, amount *big.Int, offerId types.ID) (*ethTypes.Transaction, error) {
	return _ERC20Escrow.Contract.Transact(&_ERC20Escrow.TransactOpts, token, amount, offerId)
}

type IERC20EscrowEvents interface{}
