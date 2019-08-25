// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"context"
	"math/big"

	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source erc20_escrow_wrapper.go -destination ./mocks/mock_erc20_escrow.go -package mocks IERC20EscrowManager,IERC20EscrowContract
type IERC20EscrowManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IERC20EscrowCalls

	// Transact methods
	Transact(ctx context.Context, token common.Address, amount *big.Int, offerId types.ID) error

	// Event methods
	IERC20EscrowFilterer
	IERC20EscrowWatcher
}

type IERC20EscrowCalls interface {
	Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error)
	GetTransactSelector() ([4]byte, error)
}

type IERC20EscrowTransacts interface {
	Transact(ctx context.Context, token common.Address, amount *big.Int, offerId types.ID) (*klayTypes.Receipt, error)
}

type IERC20EscrowEvents interface {
	IERC20EscrowFilterer
	IERC20EscrowParser
	IERC20EscrowWatcher
}

type IERC20EscrowFilterer interface {
}

type IERC20EscrowParser interface {
}

type IERC20EscrowWatcher interface {
}

type IERC20EscrowContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IERC20EscrowCalls
	IERC20EscrowTransacts
	IERC20EscrowEvents
}

// Manager is contract wrapper struct
type ERC20EscrowContract struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int
	client    blockchain.TxClient

	ERC20EscrowCaller
	ERC20EscrowFilterer
	ERC20EscrowTransactor
}

// Address is getter method of Accounts.address
func (c *ERC20EscrowContract) Address() common.Address {
	return c.address
}

// TxHash is getter method of Accounts.txHash
func (c *ERC20EscrowContract) TxHash() common.Hash {
	return c.txHash
}

// CreatedAt is getter method of Accounts.createdAt
func (c *ERC20EscrowContract) CreatedAt() *big.Int {
	return c.createdAt
}

func newERC20EscrowContract(address common.Address, txHash common.Hash, createdAt *big.Int, parsedABI abi.ABI, backend bind.ContractBackend) interface{} {
	contract := bind.NewBoundContract(address, parsedABI, backend, backend, backend)

	return &ERC20EscrowContract{
		address:   address,
		txHash:    txHash,
		createdAt: createdAt,
		client:    backend.(blockchain.TxClient),

		ERC20EscrowCaller:     ERC20EscrowCaller{contract: contract},
		ERC20EscrowTransactor: ERC20EscrowTransactor{contract: contract},
		ERC20EscrowFilterer:   ERC20EscrowFilterer{contract: contract},
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("ERC20Escrow", newERC20EscrowContract)
	blockchain.RegisterSelector("0x0bd9e0f8", "transact(address,uint256,bytes8)")
}

// Convert is a free data retrieval call binding the contract method 0xf8411fa9.
//
// Solidity: function convert(bytes4 sign, bytes args, bytes8 offerId) constant returns(bytes)
func (c *ERC20EscrowContract) Convert(sign [4]byte, args []byte, offerId types.ID) ([]byte, error) {
	return c.ERC20EscrowCaller.Convert(nil, sign, args, offerId)
}

// GetTransactSelector is a free data retrieval call binding the contract method 0xc0a79b5b.
//
// Solidity: function getTransactSelector() constant returns(bytes4)
func (c *ERC20EscrowContract) GetTransactSelector() ([4]byte, error) {
	return c.ERC20EscrowCaller.GetTransactSelector(nil)
}

// Transact is a paid mutator transaction binding the contract method 0x0bd9e0f8.
//
// Solidity: function transact(address token, uint256 amount, bytes8 offerId) returns()
func (c *ERC20EscrowContract) Transact(ctx context.Context, token common.Address, amount *big.Int, offerId types.ID) (*klayTypes.Receipt, error) {
	tx, err := c.ERC20EscrowTransactor.Transact(c.client.Account(), token, amount, offerId)
	if err != nil {
		return nil, err
	}
	return c.client.WaitMined(ctx, tx)
}
