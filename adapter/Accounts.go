// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AccountsABI is the input ABI used to generate the binding from.
const AccountsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"numberOfAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSignedUp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"accounts\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"},{\"name\":\"passwordSalt\",\"type\":\"bytes4\"},{\"name\":\"identityHashLock\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"SignUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"createTemporary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"name\":\"createUsingProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getAccountId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getAccountIdFromSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"},{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"passwordProof\",\"type\":\"address\"}],\"name\":\"setPassword\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountId\",\"type\":\"bytes8\"}],\"name\":\"isTemporary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccountsBin is the compiled bytecode used for deploying new contracts.
const AccountsBin = `0x6080604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630f03e4c3146100ca57806322b6ffca146100f55780636b88688814610150578063715018a6146101b05780638da5cb5b146101c75780638e3700741461021e578063bb28d52f146102a1578063bef7a6651461039a578063e0b490f7146103dd578063efc81a8c1461046a578063f1cb50f014610481578063f2fde38b14610509578063f4a3fad51461054c575b600080fd5b3480156100d657600080fd5b506100df6106a3565b6040518082815260200191505060405180910390f35b34801561010157600080fd5b50610136600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106a9565b604051808215151515815260200191505060405180910390f35b34801561015c57600080fd5b50610196600480360381019080803577ffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291905050506106c9565b604051808215151515815260200191505060405180910390f35b3480156101bc57600080fd5b506101c5610745565b005b3480156101d357600080fd5b506101dc610847565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561022a57600080fd5b5061029f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061086c565b005b3480156102ad57600080fd5b5061034e600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506109e7565b604051808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b3480156103a657600080fd5b506103db600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bb7565b005b3480156103e957600080fd5b5061041e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cc0565b604051808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34801561047657600080fd5b5061047f610e1b565b005b34801561048d57600080fd5b50610507600480360381019080803577ffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f89565b005b34801561051557600080fd5b5061054a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110da565b005b34801561055857600080fd5b50610592600480360381019080803577ffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050611141565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018660028111156105d457fe5b60ff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020018260001916600019168152602001965050505050505060405180910390f35b60055481565b60026020528060005260406000206000915054906101000a900460ff1681565b6000600160028111156106d857fe5b600160008477ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160149054906101000a900460ff16600281111561073d57fe5b149050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107a057600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515610930576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6163636f756e7420616c7265616479206578697374730000000000000000000081525060200191505060405180910390fd5b61093933611213565b9050610946818484610f89565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fdc0e23ec13ae3f3516bc4813ae5996e1fb88c4d2e57ffcc46aa94fb846e70de783604051808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390a350505050565b600080600080856040518082805190602001908083835b602083101515610a2357805182526020820191506020810190506020830392506109fe565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250610a5c83866113f6565b9150600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a9004780100000000000000000000000000000000000000000000000002905060006002811115610ad257fe5b600160008377ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160149054906101000a900460ff166002811115610b3757fe5b1415610bab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f70617373776f7264206d69736d6174636800000000000000000000000000000081525060200191505060405180910390fd5b80935050505092915050565b6000610bc233611213565b905081600160008377ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008377ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160146101000a81548160ff02191690836002811115610cb757fe5b02179055505050565b600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a9004780100000000000000000000000000000000000000000000000002905060006002811115610d3757fe5b600160008377ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160149054906101000a900460ff166002811115610d9c57fe5b14151515610e12576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f756e6b6e6f776e2061646472657373000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515610edf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6163636f756e7420616c7265616479206578697374730000000000000000000081525060200191505060405180910390fd5b610ee833611213565b9050600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fdc0e23ec13ae3f3516bc4813ae5996e1fb88c4d2e57ffcc46aa94fb846e70de783604051808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390a350565b6000600160008577ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff191681526020019081526020016000209050828160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550818160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908378010000000000000000000000000000000000000000000000009004021790555050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561113557600080fd5b61113e816114ee565b50565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a90047c010000000000000000000000000000000000000000000000000000000002908060030154905086565b6000606060008343604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040529150816040518082805190602001908083835b6020831015156112b4578051825260208201915060208101905060208303925061128f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250600160008477ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff191681526020019081526020016000209050838160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060028160000160146101000a81548160ff0219169083600281111561139257fe5b02179055506001600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050919050565b6000806000806041855114151561141057600093506114e5565b6020850151925060408501519150606085015160001a9050601b8160ff16101561143b57601b810190505b601b8160ff16141580156114535750601c8160ff1614155b1561146157600093506114e5565b600186828585604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af11580156114d8573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561152a57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820a456c88d67710338efd2e94154803cba2729c96b554e54c2159019c8ae7eea190029`

// DeployAccounts deploys a new Ethereum contract, binding an instance of Accounts to it.
func DeployAccounts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Accounts, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}, AccountsFilterer: AccountsFilterer{contract: contract}}, nil
}

// Accounts is an auto generated Go binding around an Ethereum contract.
type Accounts struct {
	AccountsCaller     // Read-only binding to the contract
	AccountsTransactor // Write-only binding to the contract
	AccountsFilterer   // Log filterer for contract events
}

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsSession struct {
	Contract     *Accounts         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsCallerSession struct {
	Contract *AccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsTransactorSession struct {
	Contract     *AccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsRaw struct {
	Contract *Accounts // Generic contract binding to access the raw methods on
}

// AccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsCallerRaw struct {
	Contract *AccountsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsTransactorRaw struct {
	Contract *AccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccounts creates a new instance of Accounts, bound to a specific deployed contract.
func NewAccounts(address common.Address, backend bind.ContractBackend) (*Accounts, error) {
	contract, err := bindAccounts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}, AccountsFilterer: AccountsFilterer{contract: contract}}, nil
}

// NewAccountsCaller creates a new read-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsCaller(address common.Address, caller bind.ContractCaller) (*AccountsCaller, error) {
	contract, err := bindAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsCaller{contract: contract}, nil
}

// NewAccountsTransactor creates a new write-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsTransactor, error) {
	contract, err := bindAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsTransactor{contract: contract}, nil
}

// NewAccountsFilterer creates a new log filterer instance of Accounts, bound to a specific deployed contract.
func NewAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountsFilterer, error) {
	contract, err := bindAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountsFilterer{contract: contract}, nil
}

// bindAccounts binds a generic wrapper to an already deployed contract.
func bindAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.AccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts( bytes8) constant returns(owner address, status uint8, proxy address, passwordProof address, passwordSalt bytes4, identityHashLock bytes32)
func (_Accounts *AccountsCaller) Accounts(opts *bind.CallOpts, arg0 [8]byte) (struct {
	Owner            common.Address
	Status           uint8
	Proxy            common.Address
	PasswordProof    common.Address
	PasswordSalt     [4]byte
	IdentityHashLock [32]byte
}, error) {
	ret := new(struct {
		Owner            common.Address
		Status           uint8
		Proxy            common.Address
		PasswordProof    common.Address
		PasswordSalt     [4]byte
		IdentityHashLock [32]byte
	})
	out := ret
	err := _Accounts.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts( bytes8) constant returns(owner address, status uint8, proxy address, passwordProof address, passwordSalt bytes4, identityHashLock bytes32)
func (_Accounts *AccountsSession) Accounts(arg0 [8]byte) (struct {
	Owner            common.Address
	Status           uint8
	Proxy            common.Address
	PasswordProof    common.Address
	PasswordSalt     [4]byte
	IdentityHashLock [32]byte
}, error) {
	return _Accounts.Contract.Accounts(&_Accounts.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf4a3fad5.
//
// Solidity: function accounts( bytes8) constant returns(owner address, status uint8, proxy address, passwordProof address, passwordSalt bytes4, identityHashLock bytes32)
func (_Accounts *AccountsCallerSession) Accounts(arg0 [8]byte) (struct {
	Owner            common.Address
	Status           uint8
	Proxy            common.Address
	PasswordProof    common.Address
	PasswordSalt     [4]byte
	IdentityHashLock [32]byte
}, error) {
	return _Accounts.Contract.Accounts(&_Accounts.CallOpts, arg0)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(sender address) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountId(opts *bind.CallOpts, sender common.Address) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccountId", sender)
	return *ret0, err
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(sender address) constant returns(bytes8)
func (_Accounts *AccountsSession) GetAccountId(sender common.Address) ([8]byte, error) {
	return _Accounts.Contract.GetAccountId(&_Accounts.CallOpts, sender)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(sender address) constant returns(bytes8)
func (_Accounts *AccountsCallerSession) GetAccountId(sender common.Address) ([8]byte, error) {
	return _Accounts.Contract.GetAccountId(&_Accounts.CallOpts, sender)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0xbb28d52f.
//
// Solidity: function getAccountIdFromSignature(message bytes, signature bytes) constant returns(bytes8)
func (_Accounts *AccountsCaller) GetAccountIdFromSignature(opts *bind.CallOpts, message []byte, signature []byte) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAccountIdFromSignature", message, signature)
	return *ret0, err
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0xbb28d52f.
//
// Solidity: function getAccountIdFromSignature(message bytes, signature bytes) constant returns(bytes8)
func (_Accounts *AccountsSession) GetAccountIdFromSignature(message []byte, signature []byte) ([8]byte, error) {
	return _Accounts.Contract.GetAccountIdFromSignature(&_Accounts.CallOpts, message, signature)
}

// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0xbb28d52f.
//
// Solidity: function getAccountIdFromSignature(message bytes, signature bytes) constant returns(bytes8)
func (_Accounts *AccountsCallerSession) GetAccountIdFromSignature(message []byte, signature []byte) ([8]byte, error) {
	return _Accounts.Contract.GetAccountIdFromSignature(&_Accounts.CallOpts, message, signature)
}

// IsSignedUp is a free data retrieval call binding the contract method 0x22b6ffca.
//
// Solidity: function isSignedUp( address) constant returns(bool)
func (_Accounts *AccountsCaller) IsSignedUp(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isSignedUp", arg0)
	return *ret0, err
}

// IsSignedUp is a free data retrieval call binding the contract method 0x22b6ffca.
//
// Solidity: function isSignedUp( address) constant returns(bool)
func (_Accounts *AccountsSession) IsSignedUp(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.IsSignedUp(&_Accounts.CallOpts, arg0)
}

// IsSignedUp is a free data retrieval call binding the contract method 0x22b6ffca.
//
// Solidity: function isSignedUp( address) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsSignedUp(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.IsSignedUp(&_Accounts.CallOpts, arg0)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(accountId bytes8) constant returns(bool)
func (_Accounts *AccountsCaller) IsTemporary(opts *bind.CallOpts, accountId [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isTemporary", accountId)
	return *ret0, err
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(accountId bytes8) constant returns(bool)
func (_Accounts *AccountsSession) IsTemporary(accountId [8]byte) (bool, error) {
	return _Accounts.Contract.IsTemporary(&_Accounts.CallOpts, accountId)
}

// IsTemporary is a free data retrieval call binding the contract method 0x6b886888.
//
// Solidity: function isTemporary(accountId bytes8) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsTemporary(accountId [8]byte) (bool, error) {
	return _Accounts.Contract.IsTemporary(&_Accounts.CallOpts, accountId)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (_Accounts *AccountsCaller) NumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "numberOfAccounts")
	return *ret0, err
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (_Accounts *AccountsSession) NumberOfAccounts() (*big.Int, error) {
	return _Accounts.Contract.NumberOfAccounts(&_Accounts.CallOpts)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() constant returns(uint256)
func (_Accounts *AccountsCallerSession) NumberOfAccounts() (*big.Int, error) {
	return _Accounts.Contract.NumberOfAccounts(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsCallerSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Accounts *AccountsTransactor) Create(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "create")
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Accounts *AccountsSession) Create() (*types.Transaction, error) {
	return _Accounts.Contract.Create(&_Accounts.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Accounts *AccountsTransactorSession) Create() (*types.Transaction, error) {
	return _Accounts.Contract.Create(&_Accounts.TransactOpts)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0xbef7a665.
//
// Solidity: function createTemporary(proxy address) returns()
func (_Accounts *AccountsTransactor) CreateTemporary(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createTemporary", proxy)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0xbef7a665.
//
// Solidity: function createTemporary(proxy address) returns()
func (_Accounts *AccountsSession) CreateTemporary(proxy common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.CreateTemporary(&_Accounts.TransactOpts, proxy)
}

// CreateTemporary is a paid mutator transaction binding the contract method 0xbef7a665.
//
// Solidity: function createTemporary(proxy address) returns()
func (_Accounts *AccountsTransactorSession) CreateTemporary(proxy common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.CreateTemporary(&_Accounts.TransactOpts, proxy)
}

// CreateUsingProxy is a paid mutator transaction binding the contract method 0x8e370074.
//
// Solidity: function createUsingProxy(owner address, proxy address, passwordProof address) returns()
func (_Accounts *AccountsTransactor) CreateUsingProxy(opts *bind.TransactOpts, owner common.Address, proxy common.Address, passwordProof common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createUsingProxy", owner, proxy, passwordProof)
}

// CreateUsingProxy is a paid mutator transaction binding the contract method 0x8e370074.
//
// Solidity: function createUsingProxy(owner address, proxy address, passwordProof address) returns()
func (_Accounts *AccountsSession) CreateUsingProxy(owner common.Address, proxy common.Address, passwordProof common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.CreateUsingProxy(&_Accounts.TransactOpts, owner, proxy, passwordProof)
}

// CreateUsingProxy is a paid mutator transaction binding the contract method 0x8e370074.
//
// Solidity: function createUsingProxy(owner address, proxy address, passwordProof address) returns()
func (_Accounts *AccountsTransactorSession) CreateUsingProxy(owner common.Address, proxy common.Address, passwordProof common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.CreateUsingProxy(&_Accounts.TransactOpts, owner, proxy, passwordProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// SetPassword is a paid mutator transaction binding the contract method 0xf1cb50f0.
//
// Solidity: function setPassword(accountId bytes8, proxy address, passwordProof address) returns()
func (_Accounts *AccountsTransactor) SetPassword(opts *bind.TransactOpts, accountId [8]byte, proxy common.Address, passwordProof common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setPassword", accountId, proxy, passwordProof)
}

// SetPassword is a paid mutator transaction binding the contract method 0xf1cb50f0.
//
// Solidity: function setPassword(accountId bytes8, proxy address, passwordProof address) returns()
func (_Accounts *AccountsSession) SetPassword(accountId [8]byte, proxy common.Address, passwordProof common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetPassword(&_Accounts.TransactOpts, accountId, proxy, passwordProof)
}

// SetPassword is a paid mutator transaction binding the contract method 0xf1cb50f0.
//
// Solidity: function setPassword(accountId bytes8, proxy address, passwordProof address) returns()
func (_Accounts *AccountsTransactorSession) SetPassword(accountId [8]byte, proxy common.Address, passwordProof common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetPassword(&_Accounts.TransactOpts, accountId, proxy, passwordProof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Accounts *AccountsTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Accounts *AccountsSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Accounts *AccountsTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, _newOwner)
}

// AccountsOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Accounts contract.
type AccountsOwnershipRenouncedIterator struct {
	Event *AccountsOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOwnershipRenounced)
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
		it.Event = new(AccountsOwnershipRenounced)
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
func (it *AccountsOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOwnershipRenounced represents a OwnershipRenounced event raised by the Accounts contract.
type AccountsOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Accounts *AccountsFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*AccountsOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOwnershipRenouncedIterator{contract: _Accounts.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Accounts *AccountsFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *AccountsOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOwnershipRenounced)
				if err := _Accounts.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// AccountsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Accounts contract.
type AccountsOwnershipTransferredIterator struct {
	Event *AccountsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOwnershipTransferred)
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
		it.Event = new(AccountsOwnershipTransferred)
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
func (it *AccountsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOwnershipTransferred represents a OwnershipTransferred event raised by the Accounts contract.
type AccountsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Accounts *AccountsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOwnershipTransferredIterator{contract: _Accounts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Accounts *AccountsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOwnershipTransferred)
				if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// AccountsSignUpIterator is returned from FilterSignUp and is used to iterate over the raw logs and unpacked data for SignUp events raised by the Accounts contract.
type AccountsSignUpIterator struct {
	Event *AccountsSignUp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsSignUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsSignUp)
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
		it.Event = new(AccountsSignUp)
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
func (it *AccountsSignUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsSignUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsSignUp represents a SignUp event raised by the Accounts contract.
type AccountsSignUp struct {
	Owner     common.Address
	Proxy     common.Address
	AccountId [8]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignUp is a free log retrieval operation binding the contract event 0xdc0e23ec13ae3f3516bc4813ae5996e1fb88c4d2e57ffcc46aa94fb846e70de7.
//
// Solidity: e SignUp(owner indexed address, proxy indexed address, accountId bytes8)
func (_Accounts *AccountsFilterer) FilterSignUp(opts *bind.FilterOpts, owner []common.Address, proxy []common.Address) (*AccountsSignUpIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "SignUp", ownerRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return &AccountsSignUpIterator{contract: _Accounts.contract, event: "SignUp", logs: logs, sub: sub}, nil
}

// WatchSignUp is a free log subscription operation binding the contract event 0xdc0e23ec13ae3f3516bc4813ae5996e1fb88c4d2e57ffcc46aa94fb846e70de7.
//
// Solidity: e SignUp(owner indexed address, proxy indexed address, accountId bytes8)
func (_Accounts *AccountsFilterer) WatchSignUp(opts *bind.WatchOpts, sink chan<- *AccountsSignUp, owner []common.Address, proxy []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "SignUp", ownerRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsSignUp)
				if err := _Accounts.contract.UnpackLog(event, "SignUp", log); err != nil {
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
