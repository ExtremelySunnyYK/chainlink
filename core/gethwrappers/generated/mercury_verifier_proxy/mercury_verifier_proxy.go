// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mercury_verifier_proxy

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

var MercuryVerifierProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"accessController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessForbidden\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"ConfigDigestAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VerifierInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"VerifierNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAccessController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAccessController\",\"type\":\"address\"}],\"name\":\"AccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldConfigDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newConfigDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"VerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"VerifierUnset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"accessController\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"initializeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"accessController\",\"type\":\"address\"}],\"name\":\"setAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currentConfigDigest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newConfigDigest\",\"type\":\"bytes32\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"unsetVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signedReport\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"verifierResponse\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516111c83803806111c883398101604081905261002f91610188565b33806000816100855760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b5576100b5816100de565b5050600380546001600160a01b0319166001600160a01b039390931692909217909155506101b8565b6001600160a01b0381163314156101375760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007c565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561019a57600080fd5b81516001600160a01b03811681146101b157600080fd5b9392505050565b611001806101c76000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806379ba509711610081578063eeb7b2481161005b578063eeb7b248146101c8578063f08391d8146101fe578063f2fde38b1461021157600080fd5b806379ba50971461018f5780638da5cb5b146101975780638e760afe146101b557600080fd5b80632cc99477116100b25780632cc9947714610154578063674145f6146101695780636e9140941461017c57600080fd5b806316d6b5f6146100ce578063181f5a7714610112575b600080fd5b60035473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b60408051808201909152601381527f566572696669657250726f787920302e302e310000000000000000000000000060208201525b6040516101099190610f23565b610167610162366004610cbe565b610224565b005b610167610177366004610c8e565b610383565b61016761018a366004610c75565b6105de565b6101676106d6565b60005473ffffffffffffffffffffffffffffffffffffffff166100e8565b6101476101c3366004610ce0565b6107d3565b6100e86101d6366004610c75565b60009081526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b61016761020c366004610c2f565b610a1b565b61016761021f366004610c2f565b610aa2565b600081815260026020526040902054819073ffffffffffffffffffffffffffffffffffffffff1680156102a7576040517f375d1fe60000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff821660248201526044015b60405180910390fd5b60008481526002602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610304576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838152600260209081526040918290208054337fffffffffffffffffffffffff0000000000000000000000000000000000000000909116811790915582518781529182018690528183015290517fbeb513e532542a562ac35699e7cd9ae7d198dcd3eee15bada6c857d28ceaddcf9181900360600190a150505050565b61038b610ab6565b8073ffffffffffffffffffffffffffffffffffffffff81166103d9576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3d3ac1b500000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a79060240160206040518083038186803b15801561045e57600080fd5b505afa158015610472573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104969190610c53565b6104cc576040517f75b0527a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600083815260026020526040902054839073ffffffffffffffffffffffffffffffffffffffff16801561054a576040517f375d1fe60000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8216602482015260440161029e565b600085815260026020908152604080832080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff891690811790915581519384529183018890528201527fbeb513e532542a562ac35699e7cd9ae7d198dcd3eee15bada6c857d28ceaddcf9060600160405180910390a15050505050565b6105e6610ab6565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610645576040517fb151802b0000000000000000000000000000000000000000000000000000000081526004810183905260240161029e565b6000828152600260205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055517f11dc15c4b8ac2b183166cc8427e5385a5ece8308217a4217338c6a7614845c4c906106ca908490849091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60405180910390a15050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610757576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161029e565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60035460609073ffffffffffffffffffffffffffffffffffffffff1680158015906108a257506040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690636b14daf8906108509033906000903690600401610eb0565b60206040518083038186803b15801561086857600080fd5b505afa15801561087c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a09190610c53565b155b156108d9576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006108e58486610f36565b60008181526002602052604090205490915073ffffffffffffffffffffffffffffffffffffffff1680610947576040517fb151802b0000000000000000000000000000000000000000000000000000000081526004810183905260240161029e565b6040517f3d3ac1b500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690633d3ac1b59061099d90899089903390600401610ee9565b600060405180830381600087803b1580156109b757600080fd5b505af11580156109cb573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610a119190810190610d52565b9695505050505050565b610a23610ab6565b6003805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f953e92b1a6442e9c3242531154a3f6f6eb00b4e9c719ba8118fa6235e4ce89b691016106ca565b610aaa610ab6565b610ab381610b39565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161029e565b565b73ffffffffffffffffffffffffffffffffffffffff8116331415610bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161029e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215610c4157600080fd5b8135610c4c81610fd2565b9392505050565b600060208284031215610c6557600080fd5b81518015158114610c4c57600080fd5b600060208284031215610c8757600080fd5b5035919050565b60008060408385031215610ca157600080fd5b823591506020830135610cb381610fd2565b809150509250929050565b60008060408385031215610cd157600080fd5b50508035926020909101359150565b60008060208385031215610cf357600080fd5b823567ffffffffffffffff80821115610d0b57600080fd5b818501915085601f830112610d1f57600080fd5b813581811115610d2e57600080fd5b866020828501011115610d4057600080fd5b60209290920196919550909350505050565b600060208284031215610d6457600080fd5b815167ffffffffffffffff80821115610d7c57600080fd5b818401915084601f830112610d9057600080fd5b815181811115610da257610da2610fa3565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610de857610de8610fa3565b81604052828152876020848701011115610e0157600080fd5b610e12836020830160208801610f73565b979650505050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60008151808452610e7e816020860160208601610f73565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff84168152604060208201526000610ee0604083018486610e1d565b95945050505050565b604081526000610efd604083018587610e1d565b905073ffffffffffffffffffffffffffffffffffffffff83166020830152949350505050565b602081526000610c4c6020830184610e66565b80356020831015610f6d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b165b92915050565b60005b83811015610f8e578181015183820152602001610f76565b83811115610f9d576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81168114610ab357600080fdfea164736f6c6343000806000a",
}

var MercuryVerifierProxyABI = MercuryVerifierProxyMetaData.ABI

var MercuryVerifierProxyBin = MercuryVerifierProxyMetaData.Bin

func DeployMercuryVerifierProxy(auth *bind.TransactOpts, backend bind.ContractBackend, accessController common.Address) (common.Address, *types.Transaction, *MercuryVerifierProxy, error) {
	parsed, err := MercuryVerifierProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MercuryVerifierProxyBin), backend, accessController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MercuryVerifierProxy{MercuryVerifierProxyCaller: MercuryVerifierProxyCaller{contract: contract}, MercuryVerifierProxyTransactor: MercuryVerifierProxyTransactor{contract: contract}, MercuryVerifierProxyFilterer: MercuryVerifierProxyFilterer{contract: contract}}, nil
}

type MercuryVerifierProxy struct {
	address common.Address
	abi     abi.ABI
	MercuryVerifierProxyCaller
	MercuryVerifierProxyTransactor
	MercuryVerifierProxyFilterer
}

type MercuryVerifierProxyCaller struct {
	contract *bind.BoundContract
}

type MercuryVerifierProxyTransactor struct {
	contract *bind.BoundContract
}

type MercuryVerifierProxyFilterer struct {
	contract *bind.BoundContract
}

type MercuryVerifierProxySession struct {
	Contract     *MercuryVerifierProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MercuryVerifierProxyCallerSession struct {
	Contract *MercuryVerifierProxyCaller
	CallOpts bind.CallOpts
}

type MercuryVerifierProxyTransactorSession struct {
	Contract     *MercuryVerifierProxyTransactor
	TransactOpts bind.TransactOpts
}

type MercuryVerifierProxyRaw struct {
	Contract *MercuryVerifierProxy
}

type MercuryVerifierProxyCallerRaw struct {
	Contract *MercuryVerifierProxyCaller
}

type MercuryVerifierProxyTransactorRaw struct {
	Contract *MercuryVerifierProxyTransactor
}

func NewMercuryVerifierProxy(address common.Address, backend bind.ContractBackend) (*MercuryVerifierProxy, error) {
	abi, err := abi.JSON(strings.NewReader(MercuryVerifierProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMercuryVerifierProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxy{address: address, abi: abi, MercuryVerifierProxyCaller: MercuryVerifierProxyCaller{contract: contract}, MercuryVerifierProxyTransactor: MercuryVerifierProxyTransactor{contract: contract}, MercuryVerifierProxyFilterer: MercuryVerifierProxyFilterer{contract: contract}}, nil
}

func NewMercuryVerifierProxyCaller(address common.Address, caller bind.ContractCaller) (*MercuryVerifierProxyCaller, error) {
	contract, err := bindMercuryVerifierProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxyCaller{contract: contract}, nil
}

func NewMercuryVerifierProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*MercuryVerifierProxyTransactor, error) {
	contract, err := bindMercuryVerifierProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxyTransactor{contract: contract}, nil
}

func NewMercuryVerifierProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*MercuryVerifierProxyFilterer, error) {
	contract, err := bindMercuryVerifierProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxyFilterer{contract: contract}, nil
}

func bindMercuryVerifierProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MercuryVerifierProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MercuryVerifierProxy *MercuryVerifierProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MercuryVerifierProxy.Contract.MercuryVerifierProxyCaller.contract.Call(opts, result, method, params...)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.MercuryVerifierProxyTransactor.contract.Transfer(opts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.MercuryVerifierProxyTransactor.contract.Transact(opts, method, params...)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MercuryVerifierProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.contract.Transfer(opts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.contract.Transact(opts, method, params...)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCaller) GetAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MercuryVerifierProxy.contract.Call(opts, &out, "getAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) GetAccessController() (common.Address, error) {
	return _MercuryVerifierProxy.Contract.GetAccessController(&_MercuryVerifierProxy.CallOpts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCallerSession) GetAccessController() (common.Address, error) {
	return _MercuryVerifierProxy.Contract.GetAccessController(&_MercuryVerifierProxy.CallOpts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCaller) GetVerifier(opts *bind.CallOpts, configDigest [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MercuryVerifierProxy.contract.Call(opts, &out, "getVerifier", configDigest)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) GetVerifier(configDigest [32]byte) (common.Address, error) {
	return _MercuryVerifierProxy.Contract.GetVerifier(&_MercuryVerifierProxy.CallOpts, configDigest)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCallerSession) GetVerifier(configDigest [32]byte) (common.Address, error) {
	return _MercuryVerifierProxy.Contract.GetVerifier(&_MercuryVerifierProxy.CallOpts, configDigest)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MercuryVerifierProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) Owner() (common.Address, error) {
	return _MercuryVerifierProxy.Contract.Owner(&_MercuryVerifierProxy.CallOpts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCallerSession) Owner() (common.Address, error) {
	return _MercuryVerifierProxy.Contract.Owner(&_MercuryVerifierProxy.CallOpts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MercuryVerifierProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) TypeAndVersion() (string, error) {
	return _MercuryVerifierProxy.Contract.TypeAndVersion(&_MercuryVerifierProxy.CallOpts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyCallerSession) TypeAndVersion() (string, error) {
	return _MercuryVerifierProxy.Contract.TypeAndVersion(&_MercuryVerifierProxy.CallOpts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MercuryVerifierProxy.contract.Transact(opts, "acceptOwnership")
}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.AcceptOwnership(&_MercuryVerifierProxy.TransactOpts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.AcceptOwnership(&_MercuryVerifierProxy.TransactOpts)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactor) InitializeVerifier(opts *bind.TransactOpts, configDigest [32]byte, verifierAddress common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.contract.Transact(opts, "initializeVerifier", configDigest, verifierAddress)
}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) InitializeVerifier(configDigest [32]byte, verifierAddress common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.InitializeVerifier(&_MercuryVerifierProxy.TransactOpts, configDigest, verifierAddress)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorSession) InitializeVerifier(configDigest [32]byte, verifierAddress common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.InitializeVerifier(&_MercuryVerifierProxy.TransactOpts, configDigest, verifierAddress)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactor) SetAccessController(opts *bind.TransactOpts, accessController common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.contract.Transact(opts, "setAccessController", accessController)
}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) SetAccessController(accessController common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.SetAccessController(&_MercuryVerifierProxy.TransactOpts, accessController)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorSession) SetAccessController(accessController common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.SetAccessController(&_MercuryVerifierProxy.TransactOpts, accessController)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactor) SetVerifier(opts *bind.TransactOpts, currentConfigDigest [32]byte, newConfigDigest [32]byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.contract.Transact(opts, "setVerifier", currentConfigDigest, newConfigDigest)
}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) SetVerifier(currentConfigDigest [32]byte, newConfigDigest [32]byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.SetVerifier(&_MercuryVerifierProxy.TransactOpts, currentConfigDigest, newConfigDigest)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorSession) SetVerifier(currentConfigDigest [32]byte, newConfigDigest [32]byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.SetVerifier(&_MercuryVerifierProxy.TransactOpts, currentConfigDigest, newConfigDigest)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.TransferOwnership(&_MercuryVerifierProxy.TransactOpts, to)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.TransferOwnership(&_MercuryVerifierProxy.TransactOpts, to)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactor) UnsetVerifier(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.contract.Transact(opts, "unsetVerifier", configDigest)
}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) UnsetVerifier(configDigest [32]byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.UnsetVerifier(&_MercuryVerifierProxy.TransactOpts, configDigest)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorSession) UnsetVerifier(configDigest [32]byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.UnsetVerifier(&_MercuryVerifierProxy.TransactOpts, configDigest)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactor) Verify(opts *bind.TransactOpts, signedReport []byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.contract.Transact(opts, "verify", signedReport)
}

func (_MercuryVerifierProxy *MercuryVerifierProxySession) Verify(signedReport []byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.Verify(&_MercuryVerifierProxy.TransactOpts, signedReport)
}

func (_MercuryVerifierProxy *MercuryVerifierProxyTransactorSession) Verify(signedReport []byte) (*types.Transaction, error) {
	return _MercuryVerifierProxy.Contract.Verify(&_MercuryVerifierProxy.TransactOpts, signedReport)
}

type MercuryVerifierProxyAccessControllerSetIterator struct {
	Event *MercuryVerifierProxyAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MercuryVerifierProxyAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MercuryVerifierProxyAccessControllerSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(MercuryVerifierProxyAccessControllerSet)
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

func (it *MercuryVerifierProxyAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *MercuryVerifierProxyAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MercuryVerifierProxyAccessControllerSet struct {
	OldAccessController common.Address
	NewAccessController common.Address
	Raw                 types.Log
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) FilterAccessControllerSet(opts *bind.FilterOpts) (*MercuryVerifierProxyAccessControllerSetIterator, error) {

	logs, sub, err := _MercuryVerifierProxy.contract.FilterLogs(opts, "AccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxyAccessControllerSetIterator{contract: _MercuryVerifierProxy.contract, event: "AccessControllerSet", logs: logs, sub: sub}, nil
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) WatchAccessControllerSet(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _MercuryVerifierProxy.contract.WatchLogs(opts, "AccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MercuryVerifierProxyAccessControllerSet)
				if err := _MercuryVerifierProxy.contract.UnpackLog(event, "AccessControllerSet", log); err != nil {
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

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) ParseAccessControllerSet(log types.Log) (*MercuryVerifierProxyAccessControllerSet, error) {
	event := new(MercuryVerifierProxyAccessControllerSet)
	if err := _MercuryVerifierProxy.contract.UnpackLog(event, "AccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MercuryVerifierProxyOwnershipTransferRequestedIterator struct {
	Event *MercuryVerifierProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MercuryVerifierProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MercuryVerifierProxyOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(MercuryVerifierProxyOwnershipTransferRequested)
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

func (it *MercuryVerifierProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MercuryVerifierProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MercuryVerifierProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MercuryVerifierProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MercuryVerifierProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxyOwnershipTransferRequestedIterator{contract: _MercuryVerifierProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MercuryVerifierProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MercuryVerifierProxyOwnershipTransferRequested)
				if err := _MercuryVerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*MercuryVerifierProxyOwnershipTransferRequested, error) {
	event := new(MercuryVerifierProxyOwnershipTransferRequested)
	if err := _MercuryVerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MercuryVerifierProxyOwnershipTransferredIterator struct {
	Event *MercuryVerifierProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MercuryVerifierProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MercuryVerifierProxyOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(MercuryVerifierProxyOwnershipTransferred)
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

func (it *MercuryVerifierProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MercuryVerifierProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MercuryVerifierProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MercuryVerifierProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MercuryVerifierProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxyOwnershipTransferredIterator{contract: _MercuryVerifierProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MercuryVerifierProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MercuryVerifierProxyOwnershipTransferred)
				if err := _MercuryVerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) ParseOwnershipTransferred(log types.Log) (*MercuryVerifierProxyOwnershipTransferred, error) {
	event := new(MercuryVerifierProxyOwnershipTransferred)
	if err := _MercuryVerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MercuryVerifierProxyVerifierSetIterator struct {
	Event *MercuryVerifierProxyVerifierSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MercuryVerifierProxyVerifierSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MercuryVerifierProxyVerifierSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(MercuryVerifierProxyVerifierSet)
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

func (it *MercuryVerifierProxyVerifierSetIterator) Error() error {
	return it.fail
}

func (it *MercuryVerifierProxyVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MercuryVerifierProxyVerifierSet struct {
	OldConfigDigest [32]byte
	NewConfigDigest [32]byte
	VerifierAddress common.Address
	Raw             types.Log
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) FilterVerifierSet(opts *bind.FilterOpts) (*MercuryVerifierProxyVerifierSetIterator, error) {

	logs, sub, err := _MercuryVerifierProxy.contract.FilterLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxyVerifierSetIterator{contract: _MercuryVerifierProxy.contract, event: "VerifierSet", logs: logs, sub: sub}, nil
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyVerifierSet) (event.Subscription, error) {

	logs, sub, err := _MercuryVerifierProxy.contract.WatchLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MercuryVerifierProxyVerifierSet)
				if err := _MercuryVerifierProxy.contract.UnpackLog(event, "VerifierSet", log); err != nil {
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

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) ParseVerifierSet(log types.Log) (*MercuryVerifierProxyVerifierSet, error) {
	event := new(MercuryVerifierProxyVerifierSet)
	if err := _MercuryVerifierProxy.contract.UnpackLog(event, "VerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MercuryVerifierProxyVerifierUnsetIterator struct {
	Event *MercuryVerifierProxyVerifierUnset

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MercuryVerifierProxyVerifierUnsetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MercuryVerifierProxyVerifierUnset)
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

	select {
	case log := <-it.logs:
		it.Event = new(MercuryVerifierProxyVerifierUnset)
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

func (it *MercuryVerifierProxyVerifierUnsetIterator) Error() error {
	return it.fail
}

func (it *MercuryVerifierProxyVerifierUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MercuryVerifierProxyVerifierUnset struct {
	ConfigDigest    [32]byte
	VerifierAddress common.Address
	Raw             types.Log
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) FilterVerifierUnset(opts *bind.FilterOpts) (*MercuryVerifierProxyVerifierUnsetIterator, error) {

	logs, sub, err := _MercuryVerifierProxy.contract.FilterLogs(opts, "VerifierUnset")
	if err != nil {
		return nil, err
	}
	return &MercuryVerifierProxyVerifierUnsetIterator{contract: _MercuryVerifierProxy.contract, event: "VerifierUnset", logs: logs, sub: sub}, nil
}

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) WatchVerifierUnset(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyVerifierUnset) (event.Subscription, error) {

	logs, sub, err := _MercuryVerifierProxy.contract.WatchLogs(opts, "VerifierUnset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MercuryVerifierProxyVerifierUnset)
				if err := _MercuryVerifierProxy.contract.UnpackLog(event, "VerifierUnset", log); err != nil {
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

func (_MercuryVerifierProxy *MercuryVerifierProxyFilterer) ParseVerifierUnset(log types.Log) (*MercuryVerifierProxyVerifierUnset, error) {
	event := new(MercuryVerifierProxyVerifierUnset)
	if err := _MercuryVerifierProxy.contract.UnpackLog(event, "VerifierUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MercuryVerifierProxy *MercuryVerifierProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MercuryVerifierProxy.abi.Events["AccessControllerSet"].ID:
		return _MercuryVerifierProxy.ParseAccessControllerSet(log)
	case _MercuryVerifierProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _MercuryVerifierProxy.ParseOwnershipTransferRequested(log)
	case _MercuryVerifierProxy.abi.Events["OwnershipTransferred"].ID:
		return _MercuryVerifierProxy.ParseOwnershipTransferred(log)
	case _MercuryVerifierProxy.abi.Events["VerifierSet"].ID:
		return _MercuryVerifierProxy.ParseVerifierSet(log)
	case _MercuryVerifierProxy.abi.Events["VerifierUnset"].ID:
		return _MercuryVerifierProxy.ParseVerifierUnset(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MercuryVerifierProxyAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x953e92b1a6442e9c3242531154a3f6f6eb00b4e9c719ba8118fa6235e4ce89b6")
}

func (MercuryVerifierProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MercuryVerifierProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MercuryVerifierProxyVerifierSet) Topic() common.Hash {
	return common.HexToHash("0xbeb513e532542a562ac35699e7cd9ae7d198dcd3eee15bada6c857d28ceaddcf")
}

func (MercuryVerifierProxyVerifierUnset) Topic() common.Hash {
	return common.HexToHash("0x11dc15c4b8ac2b183166cc8427e5385a5ece8308217a4217338c6a7614845c4c")
}

func (_MercuryVerifierProxy *MercuryVerifierProxy) Address() common.Address {
	return _MercuryVerifierProxy.address
}

type MercuryVerifierProxyInterface interface {
	GetAccessController(opts *bind.CallOpts) (common.Address, error)

	GetVerifier(opts *bind.CallOpts, configDigest [32]byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	InitializeVerifier(opts *bind.TransactOpts, configDigest [32]byte, verifierAddress common.Address) (*types.Transaction, error)

	SetAccessController(opts *bind.TransactOpts, accessController common.Address) (*types.Transaction, error)

	SetVerifier(opts *bind.TransactOpts, currentConfigDigest [32]byte, newConfigDigest [32]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnsetVerifier(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error)

	Verify(opts *bind.TransactOpts, signedReport []byte) (*types.Transaction, error)

	FilterAccessControllerSet(opts *bind.FilterOpts) (*MercuryVerifierProxyAccessControllerSetIterator, error)

	WatchAccessControllerSet(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyAccessControllerSet) (event.Subscription, error)

	ParseAccessControllerSet(log types.Log) (*MercuryVerifierProxyAccessControllerSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MercuryVerifierProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MercuryVerifierProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MercuryVerifierProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MercuryVerifierProxyOwnershipTransferred, error)

	FilterVerifierSet(opts *bind.FilterOpts) (*MercuryVerifierProxyVerifierSetIterator, error)

	WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyVerifierSet) (event.Subscription, error)

	ParseVerifierSet(log types.Log) (*MercuryVerifierProxyVerifierSet, error)

	FilterVerifierUnset(opts *bind.FilterOpts) (*MercuryVerifierProxyVerifierUnsetIterator, error)

	WatchVerifierUnset(opts *bind.WatchOpts, sink chan<- *MercuryVerifierProxyVerifierUnset) (event.Subscription, error)

	ParseVerifierUnset(log types.Log) (*MercuryVerifierProxyVerifierUnset, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
