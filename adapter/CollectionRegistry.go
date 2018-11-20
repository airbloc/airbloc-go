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

// CollectionRegistryABI is the input ABI used to generate the binding from.
const CollectionRegistryABI = "[{\"inputs\":[{\"name\":\"_accounts\",\"type\":\"address\"},{\"name\":\"_appReg\",\"type\":\"address\"},{\"name\":\"_schemaReg\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registrar\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"appId\",\"type\":\"bytes8\"},{\"indexed\":false,\"name\":\"collectionId\",\"type\":\"bytes8\"}],\"name\":\"Registration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"appId\",\"type\":\"bytes8\"}],\"name\":\"Unregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"Allowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes8\"}],\"name\":\"Denied\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_appId\",\"type\":\"bytes8\"},{\"name\":\"_schemaId\",\"type\":\"bytes8\"},{\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"appId\",\"type\":\"bytes8\"},{\"name\":\"schemaId\",\"type\":\"bytes8\"},{\"name\":\"incentiveRatioSelf\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"allowByPassword\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"},{\"name\":\"passwordSignature\",\"type\":\"bytes\"}],\"name\":\"denyByPassword\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes8\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"name\":\"user\",\"type\":\"bytes8\"}],\"name\":\"isCollectionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collectionId\",\"type\":\"bytes8\"},{\"name\":\"user\",\"type\":\"bytes8\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"isCollectionAllowedAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CollectionRegistryBin is the compiled bytecode used for deploying new contracts.
const CollectionRegistryBin = `0x6080604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630c9bb7d281146100a85780631885669414610112578063260a818e1461013457806347ba65d2146101565780634a91ee2a146101a35780634c9b30b41461020b57806397e4fea71461022d578063a3b42cba14610263578063f53fb0cb1461028e578063f8907491146102b9575b600080fd5b3480156100b457600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610110958335600160c060020a0319169536956044949193909101919081908401838280828437509497506102e19650505050505050565b005b34801561011e57600080fd5b50610110600160c060020a0319600435166104dc565b34801561014057600080fd5b50610110600160c060020a0319600435166105b9565b34801561016257600080fd5b50610178600160c060020a0319600435166107be565b60408051600160c060020a031994851681529290931660208301528183015290519081900360600190f35b3480156101af57600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610110958335600160c060020a03191695369560449491939091019190819084018382808284375094975061085f9650505050505050565b34801561021757600080fd5b50610110600160c060020a031960043516610a5a565b34801561023957600080fd5b5061024f600160c060020a031960043516610b37565b604080519115158252519081900360200190f35b34801561026f57600080fd5b50610110600160c060020a031960043581169060243516604435610b5e565b34801561029a57600080fd5b5061024f600160c060020a031960043581169060243516604435610efd565b3480156102c557600080fd5b5061024f600160c060020a031960043581169060243516610f69565b60008083604051602001808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff191681526008019150506040516020818303038152906040526040518082805190602001908083835b602083106103695780518252601f19909201916020918201910161034a565b51815160209384036101000a6000190180199092169116179052604080519290940182900382206001547f23d0601d00000000000000000000000000000000000000000000000000000000845260048401828152602485019687528b5160448601528b51929a50600160a060020a0390911697506323d0601d965089958b9550909390926064909101919085019080838360005b838110156104155781810151838201526020016103fd565b50505050905090810190601f1680156104425780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561046257600080fd5b505af1158015610476573d6000803e3d6000fd5b505050506040513d602081101561048c57600080fd5b5051905061049c84826000610f7d565b604051600160c060020a031980831691908616907f1a57e3d69528db9b16115c4ff4339d855e8468ce95579571daa74bd206d2430390600090a350505050565b600154604080517fe0b490f70000000000000000000000000000000000000000000000000000000081523360048201529051600092600160a060020a03169163e0b490f791602480830192602092919082900301818787803b15801561054157600080fd5b505af1158015610555573d6000803e3d6000fd5b505050506040513d602081101561056b57600080fd5b5051905061057b82826000610f7d565b604051600160c060020a031980831691908416907f1a57e3d69528db9b16115c4ff4339d855e8468ce95579571daa74bd206d2430390600090a35050565b60006105c482610b37565b151561061a576040805160e560020a62461bcd02815260206004820152601960248201527f636f6c6c656374696f6e20646f6573206e6f7420657869737400000000000000604482015290519081900360640190fd5b50600160c060020a03198082166000908152602081815260408083205460025482517f672b7beb00000000000000000000000000000000000000000000000000000000815260c060020a90920295861660048301523360248301529151600160a060020a039092169363672b7beb9360448084019491939192918390030190829087803b1580156106aa57600080fd5b505af11580156106be573d6000803e3d6000fd5b505050506040513d60208110156106d457600080fd5b50511515610752576040805160e560020a62461bcd02815260206004820152602360248201527f6f6e6c79206f776e65722063616e20726567697374657220636f6c6c6563746960448201527f6f6e2e0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160c060020a031980831660008181526020819052604080822080546fffffffffffffffffffffffffffffffff191681556001810183905560020182905551928416927f88bf0005675630b29e5b698355f1c09cabdf78e912367fc1850c1d8b33366f2f9190a35050565b60008060006107cc84610b37565b1515610822576040805160e560020a62461bcd02815260206004820152601960248201527f636f6c6c656374696f6e20646f6573206e6f7420657869737400000000000000604482015290519081900360640190fd5b505050600160c060020a0319166000908152602081905260409020805460019091015460c060020a80830293680100000000000000009093040291565b60008083604051602001808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff191681526008019150506040516020818303038152906040526040518082805190602001908083835b602083106108e75780518252601f1990920191602091820191016108c8565b51815160209384036101000a6000190180199092169116179052604080519290940182900382206001547f23d0601d00000000000000000000000000000000000000000000000000000000845260048401828152602485019687528b5160448601528b51929a50600160a060020a0390911697506323d0601d965089958b9550909390926064909101919085019080838360005b8381101561099357818101518382015260200161097b565b50505050905090810190601f1680156109c05780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156109e057600080fd5b505af11580156109f4573d6000803e3d6000fd5b505050506040513d6020811015610a0a57600080fd5b50519050610a1a84826001610f7d565b604051600160c060020a031980831691908616907f2575002f9c19a89406e73df97a2c23c867221b5aa503bd19f5fdc8798f00930690600090a350505050565b600154604080517fe0b490f70000000000000000000000000000000000000000000000000000000081523360048201529051600092600160a060020a03169163e0b490f791602480830192602092919082900301818787803b158015610abf57600080fd5b505af1158015610ad3573d6000803e3d6000fd5b505050506040513d6020811015610ae957600080fd5b50519050610af982826001610f7d565b604051600160c060020a031980831691908416907f2575002f9c19a89406e73df97a2c23c867221b5aa503bd19f5fdc8798f00930690600090a35050565b600160c060020a031990811660009081526020819052604090205460c060020a0216151590565b600254604080517f672b7beb000000000000000000000000000000000000000000000000000000008152600160c060020a031986166004820152336024820152905160009283928392600160a060020a039092169163672b7beb9160448082019260209290919082900301818787803b158015610bda57600080fd5b505af1158015610bee573d6000803e3d6000fd5b505050506040513d6020811015610c0457600080fd5b50511515610c82576040805160e560020a62461bcd02815260206004820152602360248201527f6f6e6c79206f776e65722063616e20726567697374657220636f6c6c6563746960448201527f6f6e2e0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600354604080517f97e4fea7000000000000000000000000000000000000000000000000000000008152600160c060020a0319881660048201529051600160a060020a03909216916397e4fea7916024808201926020929091908290030181600087803b158015610cf257600080fd5b505af1158015610d06573d6000803e3d6000fd5b505050506040513d6020811015610d1c57600080fd5b50511515610d74576040805160e560020a62461bcd02815260206004820152601b60248201527f676976656e20736368656d6120646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b60408051600160c060020a03198089166020808401919091529088166028830152603080830188905283518084039091018152605090920192839052815191929182918401908083835b60208310610ddd5780518252601f199092019160209182019101610dbe565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250610e168333611121565b600160c060020a0319811660009081526020818152604091829020805460c060020a808b0468010000000000000000026fffffffffffffffff000000000000000019918d0467ffffffffffffffff1990931692909217161781558251808401909352878352929450919250908101610e9d68056bc75e2d631000008763ffffffff6111cd16565b905280516001830155602090810151600283015560408051600160c060020a0319858116825291519189169233927fed612afce4032a5821a725a428005a5afc5c47bff7cc9c9b8d0d69e078b133fb9281900390910190a3505050505050565b600160c060020a03198084166000908152602081815260408083209386168352600390930190529081205460ff168015610f615750600160c060020a0319808516600090815260208181526040808320938716835260039093019052206001015482115b949350505050565b6000610f76838343610efd565b9392505050565b6000610f8884610b37565b1515610fde576040805160e560020a62461bcd02815260206004820152601a60248201527f436f6c6c656374696f6e20646f6573206e6f742065786973742e000000000000604482015290519081900360640190fd5b50600160c060020a0319808416600090815260208181526040808320938616835260039093019052206001810154158015906110b25750600154604080517f6b886888000000000000000000000000000000000000000000000000000000008152600160c060020a0319861660048201529051600160a060020a0390921691636b886888916024808201926020929091908290030181600087803b15801561108557600080fd5b505af1158015611099573d6000803e3d6000fd5b505050506040513d60208110156110af57600080fd5b50515b15611107576040805160e560020a62461bcd02815260206004820181905260248201527f546865206163636f756e742069732063757272656e746c79206c6f636b65642e604482015290519081900360640190fd5b805460ff1916911515919091178155436001909101555050565b604080516c01000000000000000000000000600160a060020a0384160260208083019190915243603483015260548083018690528351808403909101815260749092019283905281516000938392909182918401908083835b602083106111995780518252601f19909201916020918201910161117a565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120979650505050505050565b6000828211156111d957fe5b509003905600a165627a7a723058201bc2f26a65ac5af137c5ef4f2a47fd01e5dcd8e4a21bf2f75b42d97be60016460029`

// DeployCollectionRegistry deploys a new Ethereum contract, binding an instance of CollectionRegistry to it.
func DeployCollectionRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _accounts common.Address, _appReg common.Address, _schemaReg common.Address) (common.Address, *types.Transaction, *CollectionRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CollectionRegistryBin), backend, _accounts, _appReg, _schemaReg)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CollectionRegistry{CollectionRegistryCaller: CollectionRegistryCaller{contract: contract}, CollectionRegistryTransactor: CollectionRegistryTransactor{contract: contract}, CollectionRegistryFilterer: CollectionRegistryFilterer{contract: contract}}, nil
}

// CollectionRegistry is an auto generated Go binding around an Ethereum contract.
type CollectionRegistry struct {
	CollectionRegistryCaller     // Read-only binding to the contract
	CollectionRegistryTransactor // Write-only binding to the contract
	CollectionRegistryFilterer   // Log filterer for contract events
}

// CollectionRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectionRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectionRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectionRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectionRegistrySession struct {
	Contract     *CollectionRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CollectionRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectionRegistryCallerSession struct {
	Contract *CollectionRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CollectionRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectionRegistryTransactorSession struct {
	Contract     *CollectionRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CollectionRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectionRegistryRaw struct {
	Contract *CollectionRegistry // Generic contract binding to access the raw methods on
}

// CollectionRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectionRegistryCallerRaw struct {
	Contract *CollectionRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CollectionRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectionRegistryTransactorRaw struct {
	Contract *CollectionRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollectionRegistry creates a new instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistry(address common.Address, backend bind.ContractBackend) (*CollectionRegistry, error) {
	contract, err := bindCollectionRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistry{CollectionRegistryCaller: CollectionRegistryCaller{contract: contract}, CollectionRegistryTransactor: CollectionRegistryTransactor{contract: contract}, CollectionRegistryFilterer: CollectionRegistryFilterer{contract: contract}}, nil
}

// NewCollectionRegistryCaller creates a new read-only instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryCaller(address common.Address, caller bind.ContractCaller) (*CollectionRegistryCaller, error) {
	contract, err := bindCollectionRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryCaller{contract: contract}, nil
}

// NewCollectionRegistryTransactor creates a new write-only instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectionRegistryTransactor, error) {
	contract, err := bindCollectionRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryTransactor{contract: contract}, nil
}

// NewCollectionRegistryFilterer creates a new log filterer instance of CollectionRegistry, bound to a specific deployed contract.
func NewCollectionRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectionRegistryFilterer, error) {
	contract, err := bindCollectionRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryFilterer{contract: contract}, nil
}

// bindCollectionRegistry binds a generic wrapper to an already deployed contract.
func bindCollectionRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionRegistry *CollectionRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CollectionRegistry.Contract.CollectionRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionRegistry *CollectionRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.CollectionRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionRegistry *CollectionRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.CollectionRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionRegistry *CollectionRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CollectionRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionRegistry *CollectionRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionRegistry *CollectionRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) Exists(opts *bind.CallOpts, _id [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "exists", _id)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) Exists(_id [8]byte) (bool, error) {
	return _CollectionRegistry.Contract.Exists(&_CollectionRegistry.CallOpts, _id)
}

// Exists is a free data retrieval call binding the contract method 0x97e4fea7.
//
// Solidity: function exists(_id bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) Exists(_id [8]byte) (bool, error) {
	return _CollectionRegistry.Contract.Exists(&_CollectionRegistry.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x47ba65d2.
//
// Solidity: function get(_id bytes8) constant returns(appId bytes8, schemaId bytes8, incentiveRatioSelf uint256)
func (_CollectionRegistry *CollectionRegistryCaller) Get(opts *bind.CallOpts, _id [8]byte) (struct {
	AppId              [8]byte
	SchemaId           [8]byte
	IncentiveRatioSelf *big.Int
}, error) {
	ret := new(struct {
		AppId              [8]byte
		SchemaId           [8]byte
		IncentiveRatioSelf *big.Int
	})
	out := ret
	err := _CollectionRegistry.contract.Call(opts, out, "get", _id)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0x47ba65d2.
//
// Solidity: function get(_id bytes8) constant returns(appId bytes8, schemaId bytes8, incentiveRatioSelf uint256)
func (_CollectionRegistry *CollectionRegistrySession) Get(_id [8]byte) (struct {
	AppId              [8]byte
	SchemaId           [8]byte
	IncentiveRatioSelf *big.Int
}, error) {
	return _CollectionRegistry.Contract.Get(&_CollectionRegistry.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x47ba65d2.
//
// Solidity: function get(_id bytes8) constant returns(appId bytes8, schemaId bytes8, incentiveRatioSelf uint256)
func (_CollectionRegistry *CollectionRegistryCallerSession) Get(_id [8]byte) (struct {
	AppId              [8]byte
	SchemaId           [8]byte
	IncentiveRatioSelf *big.Int
}, error) {
	return _CollectionRegistry.Contract.Get(&_CollectionRegistry.CallOpts, _id)
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xf8907491.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) IsCollectionAllowed(opts *bind.CallOpts, collectionId [8]byte, user [8]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "isCollectionAllowed", collectionId, user)
	return *ret0, err
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xf8907491.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) IsCollectionAllowed(collectionId [8]byte, user [8]byte) (bool, error) {
	return _CollectionRegistry.Contract.IsCollectionAllowed(&_CollectionRegistry.CallOpts, collectionId, user)
}

// IsCollectionAllowed is a free data retrieval call binding the contract method 0xf8907491.
//
// Solidity: function isCollectionAllowed(collectionId bytes8, user bytes8) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) IsCollectionAllowed(collectionId [8]byte, user [8]byte) (bool, error) {
	return _CollectionRegistry.Contract.IsCollectionAllowed(&_CollectionRegistry.CallOpts, collectionId, user)
}

// IsCollectionAllowedAt is a free data retrieval call binding the contract method 0xf53fb0cb.
//
// Solidity: function isCollectionAllowedAt(collectionId bytes8, user bytes8, blockNumber uint256) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCaller) IsCollectionAllowedAt(opts *bind.CallOpts, collectionId [8]byte, user [8]byte, blockNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CollectionRegistry.contract.Call(opts, out, "isCollectionAllowedAt", collectionId, user, blockNumber)
	return *ret0, err
}

// IsCollectionAllowedAt is a free data retrieval call binding the contract method 0xf53fb0cb.
//
// Solidity: function isCollectionAllowedAt(collectionId bytes8, user bytes8, blockNumber uint256) constant returns(bool)
func (_CollectionRegistry *CollectionRegistrySession) IsCollectionAllowedAt(collectionId [8]byte, user [8]byte, blockNumber *big.Int) (bool, error) {
	return _CollectionRegistry.Contract.IsCollectionAllowedAt(&_CollectionRegistry.CallOpts, collectionId, user, blockNumber)
}

// IsCollectionAllowedAt is a free data retrieval call binding the contract method 0xf53fb0cb.
//
// Solidity: function isCollectionAllowedAt(collectionId bytes8, user bytes8, blockNumber uint256) constant returns(bool)
func (_CollectionRegistry *CollectionRegistryCallerSession) IsCollectionAllowedAt(collectionId [8]byte, user [8]byte, blockNumber *big.Int) (bool, error) {
	return _CollectionRegistry.Contract.IsCollectionAllowedAt(&_CollectionRegistry.CallOpts, collectionId, user, blockNumber)
}

// Allow is a paid mutator transaction binding the contract method 0x4c9b30b4.
//
// Solidity: function allow(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Allow(opts *bind.TransactOpts, _id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "allow", _id)
}

// Allow is a paid mutator transaction binding the contract method 0x4c9b30b4.
//
// Solidity: function allow(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistrySession) Allow(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Allow(&_CollectionRegistry.TransactOpts, _id)
}

// Allow is a paid mutator transaction binding the contract method 0x4c9b30b4.
//
// Solidity: function allow(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Allow(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Allow(&_CollectionRegistry.TransactOpts, _id)
}

// AllowByPassword is a paid mutator transaction binding the contract method 0x4a91ee2a.
//
// Solidity: function allowByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) AllowByPassword(opts *bind.TransactOpts, _id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "allowByPassword", _id, passwordSignature)
}

// AllowByPassword is a paid mutator transaction binding the contract method 0x4a91ee2a.
//
// Solidity: function allowByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistrySession) AllowByPassword(_id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.AllowByPassword(&_CollectionRegistry.TransactOpts, _id, passwordSignature)
}

// AllowByPassword is a paid mutator transaction binding the contract method 0x4a91ee2a.
//
// Solidity: function allowByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) AllowByPassword(_id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.AllowByPassword(&_CollectionRegistry.TransactOpts, _id, passwordSignature)
}

// Deny is a paid mutator transaction binding the contract method 0x18856694.
//
// Solidity: function deny(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Deny(opts *bind.TransactOpts, _id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "deny", _id)
}

// Deny is a paid mutator transaction binding the contract method 0x18856694.
//
// Solidity: function deny(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistrySession) Deny(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Deny(&_CollectionRegistry.TransactOpts, _id)
}

// Deny is a paid mutator transaction binding the contract method 0x18856694.
//
// Solidity: function deny(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Deny(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Deny(&_CollectionRegistry.TransactOpts, _id)
}

// DenyByPassword is a paid mutator transaction binding the contract method 0x0c9bb7d2.
//
// Solidity: function denyByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) DenyByPassword(opts *bind.TransactOpts, _id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "denyByPassword", _id, passwordSignature)
}

// DenyByPassword is a paid mutator transaction binding the contract method 0x0c9bb7d2.
//
// Solidity: function denyByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistrySession) DenyByPassword(_id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.DenyByPassword(&_CollectionRegistry.TransactOpts, _id, passwordSignature)
}

// DenyByPassword is a paid mutator transaction binding the contract method 0x0c9bb7d2.
//
// Solidity: function denyByPassword(_id bytes8, passwordSignature bytes) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) DenyByPassword(_id [8]byte, passwordSignature []byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.DenyByPassword(&_CollectionRegistry.TransactOpts, _id, passwordSignature)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Register(opts *bind.TransactOpts, _appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "register", _appId, _schemaId, _ratio)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistrySession) Register(_appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Register(&_CollectionRegistry.TransactOpts, _appId, _schemaId, _ratio)
}

// Register is a paid mutator transaction binding the contract method 0xa3b42cba.
//
// Solidity: function register(_appId bytes8, _schemaId bytes8, _ratio uint256) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Register(_appId [8]byte, _schemaId [8]byte, _ratio *big.Int) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Register(&_CollectionRegistry.TransactOpts, _appId, _schemaId, _ratio)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactor) Unregister(opts *bind.TransactOpts, _id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.contract.Transact(opts, "unregister", _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistrySession) Unregister(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Unregister(&_CollectionRegistry.TransactOpts, _id)
}

// Unregister is a paid mutator transaction binding the contract method 0x260a818e.
//
// Solidity: function unregister(_id bytes8) returns()
func (_CollectionRegistry *CollectionRegistryTransactorSession) Unregister(_id [8]byte) (*types.Transaction, error) {
	return _CollectionRegistry.Contract.Unregister(&_CollectionRegistry.TransactOpts, _id)
}

// CollectionRegistryAllowedIterator is returned from FilterAllowed and is used to iterate over the raw logs and unpacked data for Allowed events raised by the CollectionRegistry contract.
type CollectionRegistryAllowedIterator struct {
	Event *CollectionRegistryAllowed // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryAllowed)
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
		it.Event = new(CollectionRegistryAllowed)
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
func (it *CollectionRegistryAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryAllowed represents a Allowed event raised by the CollectionRegistry contract.
type CollectionRegistryAllowed struct {
	CollectionId [8]byte
	UserId       [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowed is a free log retrieval operation binding the contract event 0x2575002f9c19a89406e73df97a2c23c867221b5aa503bd19f5fdc8798f009306.
//
// Solidity: e Allowed(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterAllowed(opts *bind.FilterOpts, collectionId [][8]byte, userId [][8]byte) (*CollectionRegistryAllowedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Allowed", collectionIdRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryAllowedIterator{contract: _CollectionRegistry.contract, event: "Allowed", logs: logs, sub: sub}, nil
}

// FilterAllowed parses the event from given transaction receipt.
//
// Solidity: e Allowed(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseAllowedFromReceipt(receipt *types.Receipt) (*CollectionRegistryAllowed, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x2575002f9c19a89406e73df97a2c23c867221b5aa503bd19f5fdc8798f009306") {
			event := new(CollectionRegistryAllowed)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Allowed", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Allowed event not found")
}

// WatchAllowed is a free log subscription operation binding the contract event 0x2575002f9c19a89406e73df97a2c23c867221b5aa503bd19f5fdc8798f009306.
//
// Solidity: e Allowed(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchAllowed(opts *bind.WatchOpts, sink chan<- *CollectionRegistryAllowed, collectionId [][8]byte, userId [][8]byte) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Allowed", collectionIdRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryAllowed)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Allowed", log); err != nil {
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

// CollectionRegistryDeniedIterator is returned from FilterDenied and is used to iterate over the raw logs and unpacked data for Denied events raised by the CollectionRegistry contract.
type CollectionRegistryDeniedIterator struct {
	Event *CollectionRegistryDenied // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryDenied)
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
		it.Event = new(CollectionRegistryDenied)
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
func (it *CollectionRegistryDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryDenied represents a Denied event raised by the CollectionRegistry contract.
type CollectionRegistryDenied struct {
	CollectionId [8]byte
	UserId       [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDenied is a free log retrieval operation binding the contract event 0x1a57e3d69528db9b16115c4ff4339d855e8468ce95579571daa74bd206d24303.
//
// Solidity: e Denied(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterDenied(opts *bind.FilterOpts, collectionId [][8]byte, userId [][8]byte) (*CollectionRegistryDeniedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Denied", collectionIdRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryDeniedIterator{contract: _CollectionRegistry.contract, event: "Denied", logs: logs, sub: sub}, nil
}

// FilterDenied parses the event from given transaction receipt.
//
// Solidity: e Denied(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseDeniedFromReceipt(receipt *types.Receipt) (*CollectionRegistryDenied, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x1a57e3d69528db9b16115c4ff4339d855e8468ce95579571daa74bd206d24303") {
			event := new(CollectionRegistryDenied)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Denied", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Denied event not found")
}

// WatchDenied is a free log subscription operation binding the contract event 0x1a57e3d69528db9b16115c4ff4339d855e8468ce95579571daa74bd206d24303.
//
// Solidity: e Denied(collectionId indexed bytes8, userId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchDenied(opts *bind.WatchOpts, sink chan<- *CollectionRegistryDenied, collectionId [][8]byte, userId [][8]byte) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Denied", collectionIdRule, userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryDenied)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Denied", log); err != nil {
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

// CollectionRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the CollectionRegistry contract.
type CollectionRegistryRegistrationIterator struct {
	Event *CollectionRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryRegistration)
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
		it.Event = new(CollectionRegistryRegistration)
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
func (it *CollectionRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryRegistration represents a Registration event raised by the CollectionRegistry contract.
type CollectionRegistryRegistration struct {
	Registrar    common.Address
	AppId        [8]byte
	CollectionId [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xed612afce4032a5821a725a428005a5afc5c47bff7cc9c9b8d0d69e078b133fb.
//
// Solidity: e Registration(registrar indexed address, appId indexed bytes8, collectionId bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, registrar []common.Address, appId [][8]byte) (*CollectionRegistryRegistrationIterator, error) {

	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Registration", registrarRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryRegistrationIterator{contract: _CollectionRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// FilterRegistration parses the event from given transaction receipt.
//
// Solidity: e Registration(registrar indexed address, appId indexed bytes8, collectionId bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseRegistrationFromReceipt(receipt *types.Receipt) (*CollectionRegistryRegistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0xed612afce4032a5821a725a428005a5afc5c47bff7cc9c9b8d0d69e078b133fb") {
			event := new(CollectionRegistryRegistration)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Registration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Registration event not found")
}

// WatchRegistration is a free log subscription operation binding the contract event 0xed612afce4032a5821a725a428005a5afc5c47bff7cc9c9b8d0d69e078b133fb.
//
// Solidity: e Registration(registrar indexed address, appId indexed bytes8, collectionId bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *CollectionRegistryRegistration, registrar []common.Address, appId [][8]byte) (event.Subscription, error) {

	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Registration", registrarRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryRegistration)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
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

// CollectionRegistryUnregistrationIterator is returned from FilterUnregistration and is used to iterate over the raw logs and unpacked data for Unregistration events raised by the CollectionRegistry contract.
type CollectionRegistryUnregistrationIterator struct {
	Event *CollectionRegistryUnregistration // Event containing the contract specifics and raw log

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
func (it *CollectionRegistryUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionRegistryUnregistration)
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
		it.Event = new(CollectionRegistryUnregistration)
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
func (it *CollectionRegistryUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionRegistryUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionRegistryUnregistration represents a Unregistration event raised by the CollectionRegistry contract.
type CollectionRegistryUnregistration struct {
	CollectionId [8]byte
	AppId        [8]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnregistration is a free log retrieval operation binding the contract event 0x88bf0005675630b29e5b698355f1c09cabdf78e912367fc1850c1d8b33366f2f.
//
// Solidity: e Unregistration(collectionId indexed bytes8, appId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) FilterUnregistration(opts *bind.FilterOpts, collectionId [][8]byte, appId [][8]byte) (*CollectionRegistryUnregistrationIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.FilterLogs(opts, "Unregistration", collectionIdRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionRegistryUnregistrationIterator{contract: _CollectionRegistry.contract, event: "Unregistration", logs: logs, sub: sub}, nil
}

// FilterUnregistration parses the event from given transaction receipt.
//
// Solidity: e Unregistration(collectionId indexed bytes8, appId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) ParseUnregistrationFromReceipt(receipt *types.Receipt) (*CollectionRegistryUnregistration, error) {
	for _, log := range receipt.Logs {
		if log.Topics[0] == common.HexToHash("0x88bf0005675630b29e5b698355f1c09cabdf78e912367fc1850c1d8b33366f2f") {
			event := new(CollectionRegistryUnregistration)
			if err := _CollectionRegistry.contract.UnpackLog(event, "Unregistration", *log); err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("Unregistration event not found")
}

// WatchUnregistration is a free log subscription operation binding the contract event 0x88bf0005675630b29e5b698355f1c09cabdf78e912367fc1850c1d8b33366f2f.
//
// Solidity: e Unregistration(collectionId indexed bytes8, appId indexed bytes8)
func (_CollectionRegistry *CollectionRegistryFilterer) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *CollectionRegistryUnregistration, collectionId [][8]byte, appId [][8]byte) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var appIdRule []interface{}
	for _, appIdItem := range appId {
		appIdRule = append(appIdRule, appIdItem)
	}

	logs, sub, err := _CollectionRegistry.contract.WatchLogs(opts, "Unregistration", collectionIdRule, appIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionRegistryUnregistration)
				if err := _CollectionRegistry.contract.UnpackLog(event, "Unregistration", log); err != nil {
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
