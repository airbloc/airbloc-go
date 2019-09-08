// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"math/big"

	blockchain "github.com/airbloc/airbloc-go/shared/blockchain"
	abi "github.com/klaytn/klaytn/accounts/abi"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	common "github.com/klaytn/klaytn/common"
)

//go:generate mockgen -source erc20_escrow_wrapper.go -destination ./mocks/mock_erc20_escrow.go -package mocks IErc20EscrowManager,IErc20EscrowContract
type IErc20EscrowManager interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	// Call methods
	IErc20EscrowCalls

	// Transact methods

	// Event methods
	IErc20EscrowFilterer
	IErc20EscrowWatcher
}

type IErc20EscrowCalls interface {
}

type IErc20EscrowTransacts interface {
}

type IErc20EscrowEvents interface {
	IErc20EscrowFilterer
	IErc20EscrowParser
	IErc20EscrowWatcher
}

type IErc20EscrowFilterer interface {
}

type IErc20EscrowParser interface {
}

type IErc20EscrowWatcher interface {
}

type IErc20EscrowContract interface {
	Address() common.Address
	TxHash() common.Hash
	CreatedAt() *big.Int

	IErc20EscrowCalls
	IErc20EscrowTransacts
	IErc20EscrowEvents
}

// Manager is contract wrapper struct
type Erc20EscrowContract struct {
	address   common.Address
	txHash    common.Hash
	createdAt *big.Int
	client    blockchain.TxClient

	Erc20EscrowCaller
	Erc20EscrowFilterer
	Erc20EscrowTransactor
}

// Address is getter method of Accounts.address
func (c *Erc20EscrowContract) Address() common.Address {
	return c.address
}

// TxHash is getter method of Accounts.txHash
func (c *Erc20EscrowContract) TxHash() common.Hash {
	return c.txHash
}

// CreatedAt is getter method of Accounts.createdAt
func (c *Erc20EscrowContract) CreatedAt() *big.Int {
	return c.createdAt
}

func newErc20EscrowContract(address common.Address, txHash common.Hash, createdAt *big.Int, parsedABI abi.ABI, backend bind.ContractBackend) interface{} {
	contract := bind.NewBoundContract(address, parsedABI, backend, backend, backend)

	return &Erc20EscrowContract{
		address:   address,
		txHash:    txHash,
		createdAt: createdAt,
		client:    backend.(blockchain.TxClient),

		Erc20EscrowCaller:     Erc20EscrowCaller{contract: contract},
		Erc20EscrowTransactor: Erc20EscrowTransactor{contract: contract},
		Erc20EscrowFilterer:   Erc20EscrowFilterer{contract: contract},
	}
}

// convenient hacks for blockchain.Client
func init() {
	blockchain.AddContractConstructor("Erc20Escrow", newErc20EscrowContract)
}
