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

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_WHITELISTED\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"OfferPresented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_offeree\",\"type\":\"address\"}],\"name\":\"OfferSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_offeree\",\"type\":\"address\"}],\"name\":\"OfferRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"OfferOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_offerId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_offeror\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_offeree\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_reverted\",\"type\":\"bool\"}],\"name\":\"OfferClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offeror\",\"type\":\"address\"},{\"name\":\"_offeree\",\"type\":\"address\"},{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"},{\"name\":\"_opt\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"}],\"name\":\"close\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"}],\"name\":\"getOffer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"}],\"name\":\"getParticipants\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offerId\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ExchangeBin is the compiled bytecode used for deploying new contracts.
const ExchangeBin = `0x608060405260043610610107576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630988ca8c1461010c5780630b0235be1461019557806318b919e914610281578063217fe6c61461031157806324953eaa146103b2578063286dd3f51461041857806339c79e0c1461045b5780635778472a146104ab578063715018a61461058d5780637b9417c8146105a45780638221d46f146105e75780638da5cb5b1461066a5780639185d08d146106c1578063987757dd146107495780639b19251a1461077a578063cabd734f146107d5578063d597b8c814610806578063e2ec6ec314610933578063f2fde38b14610999575b600080fd5b34801561011857600080fd5b50610193600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506109dc565b005b3480156101a157600080fd5b506101c46004803603810190808035600019169060200190929190505050610a5d565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019060200280838360005b8381101561026b578082015181840152602081019050610250565b5050505090500194505050505060405180910390f35b34801561028d57600080fd5b50610296610b61565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102d65780820151818401526020810190506102bb565b50505050905090810190601f1680156103035780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561031d57600080fd5b50610398600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610b9a565b604051808215151515815260200191505060405180910390f35b3480156103be57600080fd5b5061041660048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610c21565b005b34801561042457600080fd5b50610459600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cbd565b005b34801561046757600080fd5b5061048a6004803603810190808035600019169060200190929190505050610d5a565b60405180831515151581526020018281526020019250505060405180910390f35b3480156104b757600080fd5b506104da600480360381019080803560001916906020019092919050505061101f565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018215151515815260200194505050505060405180910390f35b34801561059957600080fd5b506105a26110ce565b005b3480156105b057600080fd5b506105e5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111d0565b005b3480156105f357600080fd5b50610668600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061126d565b005b34801561067657600080fd5b5061067f6115da565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106cd57600080fd5b5061074760048036038101908080356000191690602001909291908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190929190803590602001909291905050506115ff565b005b34801561075557600080fd5b506107786004803603810190808035600019169060200190929190505050611a7d565b005b34801561078657600080fd5b506107bb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b70565b604051808215151515815260200191505060405180910390f35b3480156107e157600080fd5b506108046004803603810190808035600019169060200190929190505050611bb8565b005b34801561081257600080fd5b506108356004803603810190808035600019169060200190929190505050611cab565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825286818151815260200191508051906020019060200280838360005b8381101561091a5780820151818401526020810190506108ff565b5050505090500197505050505050505060405180910390f35b34801561093f57600080fd5b5061099760048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050611df0565b005b3480156109a557600080fd5b506109da600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e8c565b005b610a59826001836040518082805190602001908083835b602083101515610a1857805182526020820191506020810190506020830392506109f3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020611ef390919063ffffffff16565b5050565b60008060606000610a78856002611f0c90919063ffffffff16565b60000190508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260020180805480602002602001604051908101604052809291908181526020018280548015610b4c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610b02575b50505050509050935093509350509193909250565b6040805190810160405280600981526020017f77686974656c697374000000000000000000000000000000000000000000000081525081565b6000610c19836001846040518082805190602001908083835b602083101515610bd85780518252602082019150602081019050602083039250610bb3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020611f3390919063ffffffff16565b905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c7e57600080fd5b600090505b8151811015610cb957610cac8282815181101515610c9d57fe5b90602001906020020151610cbd565b8080600101915050610c83565b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d1857600080fd5b610d57816040805190810160405280600981526020017f77686974656c6973740000000000000000000000000000000000000000000000815250611f8c565b50565b6000806000806000806000610d79886002611f0c90919063ffffffff16565b9450610d84856120c0565b9350935093506000856003015414151515610e07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6f72646572206e6f7420666f756e64000000000000000000000000000000000081525060200191505060405180910390fd5b8460050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ece576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642073656e64657200000000000000000000000000000000000081525060200191505060405180910390fd5b846004015443119050610eeb88600261212790919063ffffffff16565b1515610f5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f636c6f7365206572726f7200000000000000000000000000000000000000000081525060200191505060405180910390fd5b87600019167fcccbb373df8fdcf2da39cc161a0fe9a0839ceb353ef23e1b4da1e6f4d70c5c5a85858585604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018215151515815260200194505050505060405180910390a28082965096505050505050915091565b600080600080600061103b8660026121ef90919063ffffffff16565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360020160149054906101000a900460ff169450945094509450509193509193565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561112957600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561122b57600080fd5b61126a816040805190810160405280600981526020017f77686974656c6973740000000000000000000000000000000000000000000000815250612216565b50565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614151515611313576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f696e76616c6964206f666665726f72206164647265737300000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156113b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f696e76616c6964206f666665726520616464726573730000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561145d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f696e76616c696420636f6e74726163742061646472657373000000000000000081525060200191505060405180910390fd5b61147c8273ffffffffffffffffffffffffffffffffffffffff1661234a565b15156114f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f696e76616c696420636f6e74726163742061646472657373000000000000000081525060200191505060405180910390fd5b61156a6080604051908101604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200160011515815250600261235d90919063ffffffff16565b905080600019167fb3dcdc967620c0dc51d36ab62fbfc84a0affcf92a63d94af5ee8d834ba73b94b83604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a250505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060006116188760026121ef90919063ffffffff16565b9250600073ffffffffffffffffffffffffffffffffffffffff168360020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156116e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6f72646572206e6f7420666f756e64000000000000000000000000000000000081525060200191505060405180910390fd5b8260020160149054906101000a900460ff1615151561176a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6f72646572206e6f7420736574746c656400000000000000000000000000000081525060200191505060405180910390fd5b8260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611831576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642073656e64657200000000000000000000000000000000000081525060200191505060405180910390fd5b600a8651111515156118ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6f7074206c656e67746820746f6f206c6f6e670000000000000000000000000081525060200191505060405180910390fd5b83915060008414156118bc57606491505b43820191506119c1876080604051908101604052806060604051908101604052808860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018860010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018b81525081526020018881526020018581526020018660020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525060026125ea9092919063ffffffff16565b9050801515611a38576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f6f70656e206572726f720000000000000000000000000000000000000000000081525060200191505060405180910390fd5b86600019167ffd2bfccc4b855ca8a74ab03e9139b37f94ed13cceddc7d1fd65b866f50e4d7b9836040518082815260200191505060405180910390a250505050505050565b611a9181600261272990919063ffffffff16565b1515611b05576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f736574746c65206572726f72000000000000000000000000000000000000000081525060200191505060405180910390fd5b80600019167f0fe6f606b7157a05fc0e4584759dad9e197f98e663fc52151e3d7d933d8b30f733604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a250565b6000611bb1826040805190810160405280600981526020017f77686974656c6973740000000000000000000000000000000000000000000000815250610b9a565b9050919050565b611bcc8160026127d690919063ffffffff16565b1515611c40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f72656a656374206572726f72000000000000000000000000000000000000000081525060200191505060405180910390fd5b80600019167f9a14811a3c7c99539946cbc19e1cbda74d422f25aa745041d5be47fdb866886633604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a250565b60008060606000806000806000611ccc896002611f0c90919063ffffffff16565b91508160000190508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682600201846003015485600401548660050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683805480602002602001604051908101604052809291908181526020018280548015611dd257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611d88575b50505050509350975097509750975097509750505091939550919395565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611e4d57600080fd5b600090505b8151811015611e8857611e7b8282815181101515611e6c57fe5b906020019060200201516111d0565b8080600101915050611e52565b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611ee757600080fd5b611ef08161290f565b50565b611efd8282611f33565b1515611f0857600080fd5b5050565b60008260000160008360001916600019168152602001908152602001600020905092915050565b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b612009826001836040518082805190602001908083835b602083101515611fc85780518252602082019150602081019050602083039250611fa3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020612a0990919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff167fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a826040518080602001828103825283818151815260200191508051906020019080838360005b83811015612082578082015181840152602081019050612067565b50505050905090810190601f1680156120af5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b60008060008360000160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168460000160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600301549250925092509193909250565b600082600001600083600019166000191681526020019081526020016000206000808201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006121aa9190612ac5565b5050600382016000905560048201600090556005820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550506001905092915050565b60008260010160008360001916600019168152602001908152602001600020905092915050565b612293826001836040518082805190602001908083835b602083101515612252578051825260208201915060208101905060208303925061222d565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020612a6790919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b8381101561230c5780820151818401526020810190506122f1565b50505050905090810190601f1680156123395780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b600080823b905060008111915050919050565b60008082600001518360200151846040015143604051602001808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018281526020019450505050506040516020818303038152906040526040518082805190602001908083835b60208310151561248a5780518252602082019150602081019050602083039250612465565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050600183606001901515908115158152505082846001016000836000191660001916815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff0219169083151502179055509050508091505092915050565b600081846000016000856000191660001916815260200190815260200160002060008201518160000160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906126bd929190612ae6565b505050602082015181600301556040820151816004015560608201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050600190509392505050565b600080836001016000846000191660001916815260200190815260200160002090508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156127ad57600091506127cf565b60008160020160146101000a81548160ff021916908315150217905550600191505b5092915050565b600080836001016000846000191660001916815260200190815260200160002090508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561285a5760009150612908565b8360010160008460001916600019168152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549060ff02191690555050600191505b5092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561294b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b5080546000825590600052602060002090810190612ae39190612b70565b50565b828054828255906000526020600020908101928215612b5f579160200282015b82811115612b5e5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612b06565b5b509050612b6c9190612b95565b5090565b612b9291905b80821115612b8e576000816000905550600101612b76565b5090565b90565b612bd591905b80821115612bd157600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101612b9b565b5090565b905600a165627a7a72305820c8f6b69319f169d39e69892a1514445d9ee359b99ef46fe6fa2566886b3c41240029`

// DeployExchange deploys a new Ethereum contract, binding an instance of Exchange to it.
func DeployExchange(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Exchange, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExchangeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
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
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_Exchange *ExchangeCaller) ROLEWHITELISTED(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "ROLE_WHITELISTED")
	return *ret0, err
}

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_Exchange *ExchangeSession) ROLEWHITELISTED() (string, error) {
	return _Exchange.Contract.ROLEWHITELISTED(&_Exchange.CallOpts)
}

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_Exchange *ExchangeCallerSession) ROLEWHITELISTED() (string, error) {
	return _Exchange.Contract.ROLEWHITELISTED(&_Exchange.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Exchange *ExchangeCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _Exchange.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Exchange *ExchangeSession) CheckRole(_operator common.Address, _role string) error {
	return _Exchange.Contract.CheckRole(&_Exchange.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Exchange *ExchangeCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _Exchange.Contract.CheckRole(&_Exchange.CallOpts, _operator, _role)
}

// GetOffer is a free data retrieval call binding the contract method 0xd597b8c8.
//
// Solidity: function getOffer(_offerId bytes32) constant returns(address, address, address[], uint256, uint256, address)
func (_Exchange *ExchangeCaller) GetOffer(opts *bind.CallOpts, _offerId [32]byte) (common.Address, common.Address, []common.Address, *big.Int, *big.Int, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new([]common.Address)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _Exchange.contract.Call(opts, out, "getOffer", _offerId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetOffer is a free data retrieval call binding the contract method 0xd597b8c8.
//
// Solidity: function getOffer(_offerId bytes32) constant returns(address, address, address[], uint256, uint256, address)
func (_Exchange *ExchangeSession) GetOffer(_offerId [32]byte) (common.Address, common.Address, []common.Address, *big.Int, *big.Int, common.Address, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, _offerId)
}

// GetOffer is a free data retrieval call binding the contract method 0xd597b8c8.
//
// Solidity: function getOffer(_offerId bytes32) constant returns(address, address, address[], uint256, uint256, address)
func (_Exchange *ExchangeCallerSession) GetOffer(_offerId [32]byte) (common.Address, common.Address, []common.Address, *big.Int, *big.Int, common.Address, error) {
	return _Exchange.Contract.GetOffer(&_Exchange.CallOpts, _offerId)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(_offerId bytes32) constant returns(address, address, address, bool)
func (_Exchange *ExchangeCaller) GetOrder(opts *bind.CallOpts, _offerId [32]byte) (common.Address, common.Address, common.Address, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
		ret3 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Exchange.contract.Call(opts, out, "getOrder", _offerId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(_offerId bytes32) constant returns(address, address, address, bool)
func (_Exchange *ExchangeSession) GetOrder(_offerId [32]byte) (common.Address, common.Address, common.Address, bool, error) {
	return _Exchange.Contract.GetOrder(&_Exchange.CallOpts, _offerId)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(_offerId bytes32) constant returns(address, address, address, bool)
func (_Exchange *ExchangeCallerSession) GetOrder(_offerId [32]byte) (common.Address, common.Address, common.Address, bool, error) {
	return _Exchange.Contract.GetOrder(&_Exchange.CallOpts, _offerId)
}

// GetParticipants is a free data retrieval call binding the contract method 0x0b0235be.
//
// Solidity: function getParticipants(_offerId bytes32) constant returns(address, address, address[])
func (_Exchange *ExchangeCaller) GetParticipants(opts *bind.CallOpts, _offerId [32]byte) (common.Address, common.Address, []common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Exchange.contract.Call(opts, out, "getParticipants", _offerId)
	return *ret0, *ret1, *ret2, err
}

// GetParticipants is a free data retrieval call binding the contract method 0x0b0235be.
//
// Solidity: function getParticipants(_offerId bytes32) constant returns(address, address, address[])
func (_Exchange *ExchangeSession) GetParticipants(_offerId [32]byte) (common.Address, common.Address, []common.Address, error) {
	return _Exchange.Contract.GetParticipants(&_Exchange.CallOpts, _offerId)
}

// GetParticipants is a free data retrieval call binding the contract method 0x0b0235be.
//
// Solidity: function getParticipants(_offerId bytes32) constant returns(address, address, address[])
func (_Exchange *ExchangeCallerSession) GetParticipants(_offerId [32]byte) (common.Address, common.Address, []common.Address, error) {
	return _Exchange.Contract.GetParticipants(&_Exchange.CallOpts, _offerId)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Exchange *ExchangeCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Exchange *ExchangeSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Exchange.Contract.HasRole(&_Exchange.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Exchange *ExchangeCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Exchange.Contract.HasRole(&_Exchange.CallOpts, _operator, _role)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(_operator address) constant returns(bool)
func (_Exchange *ExchangeCaller) Whitelist(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "whitelist", _operator)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(_operator address) constant returns(bool)
func (_Exchange *ExchangeSession) Whitelist(_operator common.Address) (bool, error) {
	return _Exchange.Contract.Whitelist(&_Exchange.CallOpts, _operator)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(_operator address) constant returns(bool)
func (_Exchange *ExchangeCallerSession) Whitelist(_operator common.Address) (bool, error) {
	return _Exchange.Contract.Whitelist(&_Exchange.CallOpts, _operator)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_operator address) returns()
func (_Exchange *ExchangeTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addAddressToWhitelist", _operator)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_operator address) returns()
func (_Exchange *ExchangeSession) AddAddressToWhitelist(_operator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAddressToWhitelist(&_Exchange.TransactOpts, _operator)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_operator address) returns()
func (_Exchange *ExchangeTransactorSession) AddAddressToWhitelist(_operator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAddressToWhitelist(&_Exchange.TransactOpts, _operator)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addAddressesToWhitelist", _operators)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeSession) AddAddressesToWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAddressesToWhitelist(&_Exchange.TransactOpts, _operators)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeTransactorSession) AddAddressesToWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAddressesToWhitelist(&_Exchange.TransactOpts, _operators)
}

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(_offerId bytes32) returns(bool, uint256)
func (_Exchange *ExchangeTransactor) Close(opts *bind.TransactOpts, _offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "close", _offerId)
}

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(_offerId bytes32) returns(bool, uint256)
func (_Exchange *ExchangeSession) Close(_offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Close(&_Exchange.TransactOpts, _offerId)
}

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(_offerId bytes32) returns(bool, uint256)
func (_Exchange *ExchangeTransactorSession) Close(_offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Close(&_Exchange.TransactOpts, _offerId)
}

// Open is a paid mutator transaction binding the contract method 0x9185d08d.
//
// Solidity: function open(_offerId bytes32, _opt address[], _amount uint256, _timeout uint256) returns()
func (_Exchange *ExchangeTransactor) Open(opts *bind.TransactOpts, _offerId [32]byte, _opt []common.Address, _amount *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "open", _offerId, _opt, _amount, _timeout)
}

// Open is a paid mutator transaction binding the contract method 0x9185d08d.
//
// Solidity: function open(_offerId bytes32, _opt address[], _amount uint256, _timeout uint256) returns()
func (_Exchange *ExchangeSession) Open(_offerId [32]byte, _opt []common.Address, _amount *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Open(&_Exchange.TransactOpts, _offerId, _opt, _amount, _timeout)
}

// Open is a paid mutator transaction binding the contract method 0x9185d08d.
//
// Solidity: function open(_offerId bytes32, _opt address[], _amount uint256, _timeout uint256) returns()
func (_Exchange *ExchangeTransactorSession) Open(_offerId [32]byte, _opt []common.Address, _amount *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Open(&_Exchange.TransactOpts, _offerId, _opt, _amount, _timeout)
}

// Order is a paid mutator transaction binding the contract method 0x8221d46f.
//
// Solidity: function order(_offeror address, _offeree address, _contract address) returns()
func (_Exchange *ExchangeTransactor) Order(opts *bind.TransactOpts, _offeror common.Address, _offeree common.Address, _contract common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "order", _offeror, _offeree, _contract)
}

// Order is a paid mutator transaction binding the contract method 0x8221d46f.
//
// Solidity: function order(_offeror address, _offeree address, _contract address) returns()
func (_Exchange *ExchangeSession) Order(_offeror common.Address, _offeree common.Address, _contract common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, _offeror, _offeree, _contract)
}

// Order is a paid mutator transaction binding the contract method 0x8221d46f.
//
// Solidity: function order(_offeror address, _offeree address, _contract address) returns()
func (_Exchange *ExchangeTransactorSession) Order(_offeror common.Address, _offeree common.Address, _contract common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, _offeror, _offeree, _contract)
}

// Reject is a paid mutator transaction binding the contract method 0xcabd734f.
//
// Solidity: function reject(_offerId bytes32) returns()
func (_Exchange *ExchangeTransactor) Reject(opts *bind.TransactOpts, _offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "reject", _offerId)
}

// Reject is a paid mutator transaction binding the contract method 0xcabd734f.
//
// Solidity: function reject(_offerId bytes32) returns()
func (_Exchange *ExchangeSession) Reject(_offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Reject(&_Exchange.TransactOpts, _offerId)
}

// Reject is a paid mutator transaction binding the contract method 0xcabd734f.
//
// Solidity: function reject(_offerId bytes32) returns()
func (_Exchange *ExchangeTransactorSession) Reject(_offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Reject(&_Exchange.TransactOpts, _offerId)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_operator address) returns()
func (_Exchange *ExchangeTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeAddressFromWhitelist", _operator)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_operator address) returns()
func (_Exchange *ExchangeSession) RemoveAddressFromWhitelist(_operator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAddressFromWhitelist(&_Exchange.TransactOpts, _operator)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_operator address) returns()
func (_Exchange *ExchangeTransactorSession) RemoveAddressFromWhitelist(_operator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAddressFromWhitelist(&_Exchange.TransactOpts, _operator)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeAddressesFromWhitelist", _operators)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeSession) RemoveAddressesFromWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAddressesFromWhitelist(&_Exchange.TransactOpts, _operators)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(_operators address[]) returns()
func (_Exchange *ExchangeTransactorSession) RemoveAddressesFromWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAddressesFromWhitelist(&_Exchange.TransactOpts, _operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(_offerId bytes32) returns()
func (_Exchange *ExchangeTransactor) Settle(opts *bind.TransactOpts, _offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "settle", _offerId)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(_offerId bytes32) returns()
func (_Exchange *ExchangeSession) Settle(_offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Settle(&_Exchange.TransactOpts, _offerId)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(_offerId bytes32) returns()
func (_Exchange *ExchangeTransactorSession) Settle(_offerId [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Settle(&_Exchange.TransactOpts, _offerId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Exchange *ExchangeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, _newOwner)
}

// ExchangeOfferClosedIterator is returned from FilterOfferClosed and is used to iterate over the raw logs and unpacked data for OfferClosed events raised by the Exchange contract.
type ExchangeOfferClosedIterator struct {
	Event *ExchangeOfferClosed // Event containing the contract specifics and raw log

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
func (it *ExchangeOfferClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferClosed)
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
		it.Event = new(ExchangeOfferClosed)
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
func (it *ExchangeOfferClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferClosed represents a OfferClosed event raised by the Exchange contract.
type ExchangeOfferClosed struct {
	OfferId  [32]byte
	Offeror  common.Address
	Offeree  common.Address
	Amount   *big.Int
	Reverted bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferClosed is a free log retrieval operation binding the contract event 0xcccbb373df8fdcf2da39cc161a0fe9a0839ceb353ef23e1b4da1e6f4d70c5c5a.
//
// Solidity: e OfferClosed(_offerId indexed bytes32, _offeror address, _offeree address, _amount uint256, _reverted bool)
func (_Exchange *ExchangeFilterer) FilterOfferClosed(opts *bind.FilterOpts, _offerId [][32]byte) (*ExchangeOfferClosedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferClosed", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferClosedIterator{contract: _Exchange.contract, event: "OfferClosed", logs: logs, sub: sub}, nil
}

// WatchOfferClosed is a free log subscription operation binding the contract event 0xcccbb373df8fdcf2da39cc161a0fe9a0839ceb353ef23e1b4da1e6f4d70c5c5a.
//
// Solidity: e OfferClosed(_offerId indexed bytes32, _offeror address, _offeree address, _amount uint256, _reverted bool)
func (_Exchange *ExchangeFilterer) WatchOfferClosed(opts *bind.WatchOpts, sink chan<- *ExchangeOfferClosed, _offerId [][32]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferClosed", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferClosed)
				if err := _Exchange.contract.UnpackLog(event, "OfferClosed", log); err != nil {
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

// ExchangeOfferOpenedIterator is returned from FilterOfferOpened and is used to iterate over the raw logs and unpacked data for OfferOpened events raised by the Exchange contract.
type ExchangeOfferOpenedIterator struct {
	Event *ExchangeOfferOpened // Event containing the contract specifics and raw log

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
func (it *ExchangeOfferOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOfferOpened)
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
		it.Event = new(ExchangeOfferOpened)
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
func (it *ExchangeOfferOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOfferOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOfferOpened represents a OfferOpened event raised by the Exchange contract.
type ExchangeOfferOpened struct {
	OfferId [32]byte
	Timeout *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferOpened is a free log retrieval operation binding the contract event 0xfd2bfccc4b855ca8a74ab03e9139b37f94ed13cceddc7d1fd65b866f50e4d7b9.
//
// Solidity: e OfferOpened(_offerId indexed bytes32, _timeout uint256)
func (_Exchange *ExchangeFilterer) FilterOfferOpened(opts *bind.FilterOpts, _offerId [][32]byte) (*ExchangeOfferOpenedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferOpened", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferOpenedIterator{contract: _Exchange.contract, event: "OfferOpened", logs: logs, sub: sub}, nil
}

// WatchOfferOpened is a free log subscription operation binding the contract event 0xfd2bfccc4b855ca8a74ab03e9139b37f94ed13cceddc7d1fd65b866f50e4d7b9.
//
// Solidity: e OfferOpened(_offerId indexed bytes32, _timeout uint256)
func (_Exchange *ExchangeFilterer) WatchOfferOpened(opts *bind.WatchOpts, sink chan<- *ExchangeOfferOpened, _offerId [][32]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferOpened", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOfferOpened)
				if err := _Exchange.contract.UnpackLog(event, "OfferOpened", log); err != nil {
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

	logs chan types.Log        // Log channel receiving the found contract events
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
	OfferId  [32]byte
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferPresented is a free log retrieval operation binding the contract event 0xb3dcdc967620c0dc51d36ab62fbfc84a0affcf92a63d94af5ee8d834ba73b94b.
//
// Solidity: e OfferPresented(_offerId indexed bytes32, _contract address)
func (_Exchange *ExchangeFilterer) FilterOfferPresented(opts *bind.FilterOpts, _offerId [][32]byte) (*ExchangeOfferPresentedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferPresented", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferPresentedIterator{contract: _Exchange.contract, event: "OfferPresented", logs: logs, sub: sub}, nil
}

// WatchOfferPresented is a free log subscription operation binding the contract event 0xb3dcdc967620c0dc51d36ab62fbfc84a0affcf92a63d94af5ee8d834ba73b94b.
//
// Solidity: e OfferPresented(_offerId indexed bytes32, _contract address)
func (_Exchange *ExchangeFilterer) WatchOfferPresented(opts *bind.WatchOpts, sink chan<- *ExchangeOfferPresented, _offerId [][32]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferPresented", _offerIdRule)
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

// ExchangeOfferRejectedIterator is returned from FilterOfferRejected and is used to iterate over the raw logs and unpacked data for OfferRejected events raised by the Exchange contract.
type ExchangeOfferRejectedIterator struct {
	Event *ExchangeOfferRejected // Event containing the contract specifics and raw log

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
	OfferId [32]byte
	Offeree common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferRejected is a free log retrieval operation binding the contract event 0x9a14811a3c7c99539946cbc19e1cbda74d422f25aa745041d5be47fdb8668866.
//
// Solidity: e OfferRejected(_offerId indexed bytes32, _offeree address)
func (_Exchange *ExchangeFilterer) FilterOfferRejected(opts *bind.FilterOpts, _offerId [][32]byte) (*ExchangeOfferRejectedIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferRejected", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferRejectedIterator{contract: _Exchange.contract, event: "OfferRejected", logs: logs, sub: sub}, nil
}

// WatchOfferRejected is a free log subscription operation binding the contract event 0x9a14811a3c7c99539946cbc19e1cbda74d422f25aa745041d5be47fdb8668866.
//
// Solidity: e OfferRejected(_offerId indexed bytes32, _offeree address)
func (_Exchange *ExchangeFilterer) WatchOfferRejected(opts *bind.WatchOpts, sink chan<- *ExchangeOfferRejected, _offerId [][32]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferRejected", _offerIdRule)
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

	logs chan types.Log        // Log channel receiving the found contract events
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
	OfferId [32]byte
	Offeree common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferSettled is a free log retrieval operation binding the contract event 0x0fe6f606b7157a05fc0e4584759dad9e197f98e663fc52151e3d7d933d8b30f7.
//
// Solidity: e OfferSettled(_offerId indexed bytes32, _offeree address)
func (_Exchange *ExchangeFilterer) FilterOfferSettled(opts *bind.FilterOpts, _offerId [][32]byte) (*ExchangeOfferSettledIterator, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OfferSettled", _offerIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOfferSettledIterator{contract: _Exchange.contract, event: "OfferSettled", logs: logs, sub: sub}, nil
}

// WatchOfferSettled is a free log subscription operation binding the contract event 0x0fe6f606b7157a05fc0e4584759dad9e197f98e663fc52151e3d7d933d8b30f7.
//
// Solidity: e OfferSettled(_offerId indexed bytes32, _offeree address)
func (_Exchange *ExchangeFilterer) WatchOfferSettled(opts *bind.WatchOpts, sink chan<- *ExchangeOfferSettled, _offerId [][32]byte) (event.Subscription, error) {

	var _offerIdRule []interface{}
	for _, _offerIdItem := range _offerId {
		_offerIdRule = append(_offerIdRule, _offerIdItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OfferSettled", _offerIdRule)
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

// ExchangeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Exchange contract.
type ExchangeOwnershipRenouncedIterator struct {
	Event *ExchangeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipRenounced)
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
		it.Event = new(ExchangeOwnershipRenounced)
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
func (it *ExchangeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipRenounced represents a OwnershipRenounced event raised by the Exchange contract.
type ExchangeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Exchange *ExchangeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ExchangeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipRenouncedIterator{contract: _Exchange.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Exchange *ExchangeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipRenounced)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ExchangeRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the Exchange contract.
type ExchangeRoleAddedIterator struct {
	Event *ExchangeRoleAdded // Event containing the contract specifics and raw log

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
func (it *ExchangeRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRoleAdded)
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
		it.Event = new(ExchangeRoleAdded)
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
func (it *ExchangeRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRoleAdded represents a RoleAdded event raised by the Exchange contract.
type ExchangeRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*ExchangeRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeRoleAddedIterator{contract: _Exchange.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *ExchangeRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRoleAdded)
				if err := _Exchange.contract.UnpackLog(event, "RoleAdded", log); err != nil {
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

// ExchangeRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the Exchange contract.
type ExchangeRoleRemovedIterator struct {
	Event *ExchangeRoleRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRoleRemoved)
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
		it.Event = new(ExchangeRoleRemoved)
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
func (it *ExchangeRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRoleRemoved represents a RoleRemoved event raised by the Exchange contract.
type ExchangeRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*ExchangeRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeRoleRemovedIterator{contract: _Exchange.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Exchange *ExchangeFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRoleRemoved)
				if err := _Exchange.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
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
