// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"errors"
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
	_ = errors.New
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
const AccountsBin = `0x6080604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630f03e4c381146100c957806322b6ffca146100f05780636b88688814610125578063715018a6146101475780638da5cb5b1461015e5780638e3700741461018f578063bb28d52f146101bc578063bef7a66514610270578063e0b490f714610291578063efc81a8c146102b2578063f1cb50f0146102c7578063f2fde38b146102fb578063f4a3fad51461031c575b600080fd5b3480156100d557600080fd5b506100de6103bc565b60408051918252519081900360200190f35b3480156100fc57600080fd5b50610111600160a060020a03600435166103c2565b604080519115158252519081900360200190f35b34801561013157600080fd5b50610111600160c060020a0319600435166103d7565b34801561015357600080fd5b5061015c610411565b005b34801561016a57600080fd5b5061017361047d565b60408051600160a060020a039092168252519081900360200190f35b34801561019b57600080fd5b5061015c600160a060020a036004358116906024358116906044351661048c565b3480156101c857600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261025394369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061055d9650505050505050565b60408051600160c060020a03199092168252519081900360200190f35b34801561027c57600080fd5b5061015c600160a060020a0360043516610699565b34801561029d57600080fd5b50610253600160a060020a0360043516610717565b3480156102be57600080fd5b5061015c6107d7565b3480156102d357600080fd5b5061015c600160c060020a031960043516600160a060020a0360243581169060443516610891565b34801561030757600080fd5b5061015c600160a060020a0360043516610929565b34801561032857600080fd5b5061033e600160c060020a03196004351661094c565b604051600160a060020a03871681526020810186600281111561035d57fe5b60ff168152600160a060020a039586166020820152939094166040808501919091527fffffffff0000000000000000000000000000000000000000000000000000000090921660608401526080830152519081900360a0019350915050f35b60055481565b60026020526000908152604090205460ff1681565b60006001600160c060020a0319831660009081526001602052604090205460a060020a900460ff16600281111561040a57fe5b1492915050565b600054600160a060020a0316331461042857600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b3360009081526002602052604081205460ff16156104f4576040805160e560020a62461bcd02815260206004820152601660248201527f6163636f756e7420616c72656164792065786973747300000000000000000000604482015290519081900360640190fd5b6104fd336109b4565b905061050a818484610891565b60408051600160c060020a0319831681529051600160a060020a0380861692908716917fdc0e23ec13ae3f3516bc4813ae5996e1fb88c4d2e57ffcc46aa94fb846e70de79181900360200190a350505050565b600080600080856040518082805190602001908083835b602083106105935780518252601f199092019160209182019101610574565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092506105cc8386610aef565b600160a060020a03811660009081526003602052604081205491935078010000000000000000000000000000000000000000000000009091029150600160c060020a0319821660009081526001602052604090205460a060020a900460ff16600281111561063657fe5b141561068c576040805160e560020a62461bcd02815260206004820152601160248201527f70617373776f7264206d69736d61746368000000000000000000000000000000604482015290519081900360640190fd5b8093505b50505092915050565b60006106a4336109b4565b600160c060020a031981166000908152600160208190526040909120808201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387161790558054929350909174ff0000000000000000000000000000000000000000191660a060020a8302179055505050565b600160a060020a03811660009081526004602052604081205478010000000000000000000000000000000000000000000000000281600160c060020a0319821660009081526001602052604090205460a060020a900460ff16600281111561077b57fe5b14156107d1576040805160e560020a62461bcd02815260206004820152600f60248201527f756e6b6e6f776e20616464726573730000000000000000000000000000000000604482015290519081900360640190fd5b92915050565b3360009081526002602052604081205460ff161561083f576040805160e560020a62461bcd02815260206004820152601660248201527f6163636f756e7420616c72656164792065786973747300000000000000000000604482015290519081900360640190fd5b610848336109b4565b60408051600160c060020a031983168152905191925060009133917fdc0e23ec13ae3f3516bc4813ae5996e1fb88c4d2e57ffcc46aa94fb846e70de7919081900360200190a350565b600160c060020a031983166000908152600160208181526040808420928301805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03988916179091556002909301805490931694909516938417909155918152600390915220805467ffffffffffffffff19167801000000000000000000000000000000000000000000000000909204919091179055565b600054600160a060020a0316331461094057600080fd5b61094981610bbf565b50565b60016020819052600091825260409091208054918101546002820154600390920154600160a060020a038085169460a060020a9081900460ff1694938216939182169291047c0100000000000000000000000000000000000000000000000000000000029086565b60006060600083436040516020018083600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040529150816040518082805190602001908083835b60208310610a355780518252601f199092019160209182019101610a16565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120600160c060020a0319811660009081526001808452858220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039d909d169c8d1774ff00000000000000000000000000000000000000001916740200000000000000000000000000000000000000001790559a81526002909252929020805460ff19169098179097559695505050505050565b60008060008084516041141515610b095760009350610690565b50505060208201516040830151606084015160001a601b60ff82161015610b2e57601b015b8060ff16601b14158015610b4657508060ff16601c14155b15610b545760009350610690565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015610bae573d6000803e3d6000fd5b505050602060405103519350610690565b600160a060020a0381161515610bd457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582024fbfa0cc047beade69c7486f66ac5dbb397229e46c9f2de852a1ef0c91c645f0029`

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

// FilterOwnershipRenounced parses the event from given transaction receipt.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Accounts *AccountsFilterer) ParseOwnershipRenouncedFromReceipt(receipt *types.Receipt) (*AccountsOwnershipRenounced, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820") {
			event := new(AccountsOwnershipRenounced)
			if err := _Accounts.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OwnershipRenounced event not found")
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

// FilterOwnershipTransferred parses the event from given transaction receipt.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Accounts *AccountsFilterer) ParseOwnershipTransferredFromReceipt(receipt *types.Receipt) (*AccountsOwnershipTransferred, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0") {
			event := new(AccountsOwnershipTransferred)
			if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("OwnershipTransferred event not found")
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

// FilterSignUp parses the event from given transaction receipt.
//
// Solidity: e SignUp(owner indexed address, proxy indexed address, accountId bytes8)
func (_Accounts *AccountsFilterer) ParseSignUpFromReceipt(receipt *types.Receipt) (*AccountsSignUp, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xdc0e23ec13ae3f3516bc4813ae5996e1fb88c4d2e57ffcc46aa94fb846e70de7") {
			event := new(AccountsSignUp)
			if err := _Accounts.contract.UnpackLog(event, "SignUp", log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("SignUp event not found")
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
