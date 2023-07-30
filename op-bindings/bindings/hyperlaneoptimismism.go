// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HyperlaneOptimismISMMetaData contains all meta data concerning the HyperlaneOptimismISM contract.
var HyperlaneOptimismISMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"InvalidCrossChainSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCrossChainCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCrossChainCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"ReceivedMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"l1Hook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moduleType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Hook\",\"type\":\"address\"}],\"name\":\"setOptimismHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"verifiedMessageIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"verifyMessageId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161094138038061094183398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516108af6100926000396000818161058c01526105e901526108af6000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c80636465e69f116100505780636465e69f1461011457806386f8011f1461012e578063f7e83aee1461014157600080fd5b8063106ddb18146100775780635161ad8b146100c1578063529fe082146100d6575b600080fd5b6002546100979073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100d46100cf3660046106ff565b610154565b005b6101046100e43660046106ff565b600160209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100b8565b61011c600681565b60405160ff90911681526020016100b8565b6100d461013c366004610746565b61028e565b61010461014f3660046107b3565b6104fb565b60025473ffffffffffffffffffffffffffffffffffffffff16610175610573565b73ffffffffffffffffffffffffffffffffffffffff161461021d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f48797065726c616e654f7074696d69736d49534d3a2073656e6465722069732060448201527f6e6f742074686520686f6f6b000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600081815260016020818152604080842086855290915280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690921790915551829184917ff9347035c6654873e4aa34118c96b53189dc92edcb628c7bad1f9bce9089ecfe9190a35050565b600054610100900460ff16158080156102ae5750600054600160ff909116105b806102c85750303b1580156102c8575060005460ff166001145b610354576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610214565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156103b257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8216610454576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f48797065726c616e654f7074696d69736d49534d3a20696e76616c6964206c3160448201527f486f6f6b000000000000000000000000000000000000000000000000000000006064820152608401610214565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841617905580156104f757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b60008061053d84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061061292505050565b9050600061054b858561061d565b60009283526001602090815260408085209285529190529091205460ff169695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633146105e4576040517f4a74df9200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61060d7f000000000000000000000000000000000000000000000000000000000000000061063f565b905090565b805160209091012090565b600061062d60296009848661081f565b61063691610849565b90505b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82163314610690576040517f4a74df9200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106399190610885565b6000806040838503121561071257600080fd5b50508035926020909101359150565b73ffffffffffffffffffffffffffffffffffffffff8116811461074357600080fd5b50565b60006020828403121561075857600080fd5b813561076381610721565b9392505050565b60008083601f84011261077c57600080fd5b50813567ffffffffffffffff81111561079457600080fd5b6020830191508360208285010111156107ac57600080fd5b9250929050565b600080600080604085870312156107c957600080fd5b843567ffffffffffffffff808211156107e157600080fd5b6107ed8883890161076a565b9096509450602087013591508082111561080657600080fd5b506108138782880161076a565b95989497509550505050565b6000808585111561082f57600080fd5b8386111561083c57600080fd5b5050820193919092039150565b80356020831015610639577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b1692915050565b60006020828403121561089757600080fd5b81516107638161072156fea164736f6c634300080f000a",
}

// HyperlaneOptimismISMABI is the input ABI used to generate the binding from.
// Deprecated: Use HyperlaneOptimismISMMetaData.ABI instead.
var HyperlaneOptimismISMABI = HyperlaneOptimismISMMetaData.ABI

// HyperlaneOptimismISMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HyperlaneOptimismISMMetaData.Bin instead.
var HyperlaneOptimismISMBin = HyperlaneOptimismISMMetaData.Bin

// DeployHyperlaneOptimismISM deploys a new Ethereum contract, binding an instance of HyperlaneOptimismISM to it.
func DeployHyperlaneOptimismISM(auth *bind.TransactOpts, backend bind.ContractBackend, _l2Messenger common.Address) (common.Address, *types.Transaction, *HyperlaneOptimismISM, error) {
	parsed, err := HyperlaneOptimismISMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HyperlaneOptimismISMBin), backend, _l2Messenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HyperlaneOptimismISM{HyperlaneOptimismISMCaller: HyperlaneOptimismISMCaller{contract: contract}, HyperlaneOptimismISMTransactor: HyperlaneOptimismISMTransactor{contract: contract}, HyperlaneOptimismISMFilterer: HyperlaneOptimismISMFilterer{contract: contract}}, nil
}

// HyperlaneOptimismISM is an auto generated Go binding around an Ethereum contract.
type HyperlaneOptimismISM struct {
	HyperlaneOptimismISMCaller     // Read-only binding to the contract
	HyperlaneOptimismISMTransactor // Write-only binding to the contract
	HyperlaneOptimismISMFilterer   // Log filterer for contract events
}

// HyperlaneOptimismISMCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyperlaneOptimismISMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneOptimismISMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperlaneOptimismISMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneOptimismISMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperlaneOptimismISMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneOptimismISMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperlaneOptimismISMSession struct {
	Contract     *HyperlaneOptimismISM // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HyperlaneOptimismISMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperlaneOptimismISMCallerSession struct {
	Contract *HyperlaneOptimismISMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// HyperlaneOptimismISMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperlaneOptimismISMTransactorSession struct {
	Contract     *HyperlaneOptimismISMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// HyperlaneOptimismISMRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyperlaneOptimismISMRaw struct {
	Contract *HyperlaneOptimismISM // Generic contract binding to access the raw methods on
}

// HyperlaneOptimismISMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperlaneOptimismISMCallerRaw struct {
	Contract *HyperlaneOptimismISMCaller // Generic read-only contract binding to access the raw methods on
}

// HyperlaneOptimismISMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperlaneOptimismISMTransactorRaw struct {
	Contract *HyperlaneOptimismISMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperlaneOptimismISM creates a new instance of HyperlaneOptimismISM, bound to a specific deployed contract.
func NewHyperlaneOptimismISM(address common.Address, backend bind.ContractBackend) (*HyperlaneOptimismISM, error) {
	contract, err := bindHyperlaneOptimismISM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HyperlaneOptimismISM{HyperlaneOptimismISMCaller: HyperlaneOptimismISMCaller{contract: contract}, HyperlaneOptimismISMTransactor: HyperlaneOptimismISMTransactor{contract: contract}, HyperlaneOptimismISMFilterer: HyperlaneOptimismISMFilterer{contract: contract}}, nil
}

// NewHyperlaneOptimismISMCaller creates a new read-only instance of HyperlaneOptimismISM, bound to a specific deployed contract.
func NewHyperlaneOptimismISMCaller(address common.Address, caller bind.ContractCaller) (*HyperlaneOptimismISMCaller, error) {
	contract, err := bindHyperlaneOptimismISM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperlaneOptimismISMCaller{contract: contract}, nil
}

// NewHyperlaneOptimismISMTransactor creates a new write-only instance of HyperlaneOptimismISM, bound to a specific deployed contract.
func NewHyperlaneOptimismISMTransactor(address common.Address, transactor bind.ContractTransactor) (*HyperlaneOptimismISMTransactor, error) {
	contract, err := bindHyperlaneOptimismISM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperlaneOptimismISMTransactor{contract: contract}, nil
}

// NewHyperlaneOptimismISMFilterer creates a new log filterer instance of HyperlaneOptimismISM, bound to a specific deployed contract.
func NewHyperlaneOptimismISMFilterer(address common.Address, filterer bind.ContractFilterer) (*HyperlaneOptimismISMFilterer, error) {
	contract, err := bindHyperlaneOptimismISM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperlaneOptimismISMFilterer{contract: contract}, nil
}

// bindHyperlaneOptimismISM binds a generic wrapper to an already deployed contract.
func bindHyperlaneOptimismISM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HyperlaneOptimismISMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperlaneOptimismISM *HyperlaneOptimismISMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperlaneOptimismISM.Contract.HyperlaneOptimismISMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperlaneOptimismISM *HyperlaneOptimismISMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.Contract.HyperlaneOptimismISMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperlaneOptimismISM *HyperlaneOptimismISMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.Contract.HyperlaneOptimismISMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperlaneOptimismISM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperlaneOptimismISM *HyperlaneOptimismISMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperlaneOptimismISM *HyperlaneOptimismISMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.Contract.contract.Transact(opts, method, params...)
}

// L1Hook is a free data retrieval call binding the contract method 0x106ddb18.
//
// Solidity: function l1Hook() view returns(address)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCaller) L1Hook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperlaneOptimismISM.contract.Call(opts, &out, "l1Hook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Hook is a free data retrieval call binding the contract method 0x106ddb18.
//
// Solidity: function l1Hook() view returns(address)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMSession) L1Hook() (common.Address, error) {
	return _HyperlaneOptimismISM.Contract.L1Hook(&_HyperlaneOptimismISM.CallOpts)
}

// L1Hook is a free data retrieval call binding the contract method 0x106ddb18.
//
// Solidity: function l1Hook() view returns(address)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCallerSession) L1Hook() (common.Address, error) {
	return _HyperlaneOptimismISM.Contract.L1Hook(&_HyperlaneOptimismISM.CallOpts)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCaller) ModuleType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HyperlaneOptimismISM.contract.Call(opts, &out, "moduleType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMSession) ModuleType() (uint8, error) {
	return _HyperlaneOptimismISM.Contract.ModuleType(&_HyperlaneOptimismISM.CallOpts)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCallerSession) ModuleType() (uint8, error) {
	return _HyperlaneOptimismISM.Contract.ModuleType(&_HyperlaneOptimismISM.CallOpts)
}

// VerifiedMessageIds is a free data retrieval call binding the contract method 0x529fe082.
//
// Solidity: function verifiedMessageIds(bytes32 , bytes32 ) view returns(bool)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCaller) VerifiedMessageIds(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _HyperlaneOptimismISM.contract.Call(opts, &out, "verifiedMessageIds", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifiedMessageIds is a free data retrieval call binding the contract method 0x529fe082.
//
// Solidity: function verifiedMessageIds(bytes32 , bytes32 ) view returns(bool)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMSession) VerifiedMessageIds(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _HyperlaneOptimismISM.Contract.VerifiedMessageIds(&_HyperlaneOptimismISM.CallOpts, arg0, arg1)
}

// VerifiedMessageIds is a free data retrieval call binding the contract method 0x529fe082.
//
// Solidity: function verifiedMessageIds(bytes32 , bytes32 ) view returns(bool)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCallerSession) VerifiedMessageIds(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _HyperlaneOptimismISM.Contract.VerifiedMessageIds(&_HyperlaneOptimismISM.CallOpts, arg0, arg1)
}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes , bytes _message) view returns(bool)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCaller) Verify(opts *bind.CallOpts, arg0 []byte, _message []byte) (bool, error) {
	var out []interface{}
	err := _HyperlaneOptimismISM.contract.Call(opts, &out, "verify", arg0, _message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes , bytes _message) view returns(bool)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMSession) Verify(arg0 []byte, _message []byte) (bool, error) {
	return _HyperlaneOptimismISM.Contract.Verify(&_HyperlaneOptimismISM.CallOpts, arg0, _message)
}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes , bytes _message) view returns(bool)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMCallerSession) Verify(arg0 []byte, _message []byte) (bool, error) {
	return _HyperlaneOptimismISM.Contract.Verify(&_HyperlaneOptimismISM.CallOpts, arg0, _message)
}

// SetOptimismHook is a paid mutator transaction binding the contract method 0x86f8011f.
//
// Solidity: function setOptimismHook(address _l1Hook) returns()
func (_HyperlaneOptimismISM *HyperlaneOptimismISMTransactor) SetOptimismHook(opts *bind.TransactOpts, _l1Hook common.Address) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.contract.Transact(opts, "setOptimismHook", _l1Hook)
}

// SetOptimismHook is a paid mutator transaction binding the contract method 0x86f8011f.
//
// Solidity: function setOptimismHook(address _l1Hook) returns()
func (_HyperlaneOptimismISM *HyperlaneOptimismISMSession) SetOptimismHook(_l1Hook common.Address) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.Contract.SetOptimismHook(&_HyperlaneOptimismISM.TransactOpts, _l1Hook)
}

// SetOptimismHook is a paid mutator transaction binding the contract method 0x86f8011f.
//
// Solidity: function setOptimismHook(address _l1Hook) returns()
func (_HyperlaneOptimismISM *HyperlaneOptimismISMTransactorSession) SetOptimismHook(_l1Hook common.Address) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.Contract.SetOptimismHook(&_HyperlaneOptimismISM.TransactOpts, _l1Hook)
}

// VerifyMessageId is a paid mutator transaction binding the contract method 0x5161ad8b.
//
// Solidity: function verifyMessageId(bytes32 _sender, bytes32 _messageId) returns()
func (_HyperlaneOptimismISM *HyperlaneOptimismISMTransactor) VerifyMessageId(opts *bind.TransactOpts, _sender [32]byte, _messageId [32]byte) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.contract.Transact(opts, "verifyMessageId", _sender, _messageId)
}

// VerifyMessageId is a paid mutator transaction binding the contract method 0x5161ad8b.
//
// Solidity: function verifyMessageId(bytes32 _sender, bytes32 _messageId) returns()
func (_HyperlaneOptimismISM *HyperlaneOptimismISMSession) VerifyMessageId(_sender [32]byte, _messageId [32]byte) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.Contract.VerifyMessageId(&_HyperlaneOptimismISM.TransactOpts, _sender, _messageId)
}

// VerifyMessageId is a paid mutator transaction binding the contract method 0x5161ad8b.
//
// Solidity: function verifyMessageId(bytes32 _sender, bytes32 _messageId) returns()
func (_HyperlaneOptimismISM *HyperlaneOptimismISMTransactorSession) VerifyMessageId(_sender [32]byte, _messageId [32]byte) (*types.Transaction, error) {
	return _HyperlaneOptimismISM.Contract.VerifyMessageId(&_HyperlaneOptimismISM.TransactOpts, _sender, _messageId)
}

// HyperlaneOptimismISMInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HyperlaneOptimismISM contract.
type HyperlaneOptimismISMInitializedIterator struct {
	Event *HyperlaneOptimismISMInitialized // Event containing the contract specifics and raw log

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
func (it *HyperlaneOptimismISMInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneOptimismISMInitialized)
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
		it.Event = new(HyperlaneOptimismISMInitialized)
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
func (it *HyperlaneOptimismISMInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneOptimismISMInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneOptimismISMInitialized represents a Initialized event raised by the HyperlaneOptimismISM contract.
type HyperlaneOptimismISMInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMFilterer) FilterInitialized(opts *bind.FilterOpts) (*HyperlaneOptimismISMInitializedIterator, error) {

	logs, sub, err := _HyperlaneOptimismISM.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HyperlaneOptimismISMInitializedIterator{contract: _HyperlaneOptimismISM.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HyperlaneOptimismISMInitialized) (event.Subscription, error) {

	logs, sub, err := _HyperlaneOptimismISM.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneOptimismISMInitialized)
				if err := _HyperlaneOptimismISM.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMFilterer) ParseInitialized(log types.Log) (*HyperlaneOptimismISMInitialized, error) {
	event := new(HyperlaneOptimismISMInitialized)
	if err := _HyperlaneOptimismISM.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneOptimismISMReceivedMessageIterator is returned from FilterReceivedMessage and is used to iterate over the raw logs and unpacked data for ReceivedMessage events raised by the HyperlaneOptimismISM contract.
type HyperlaneOptimismISMReceivedMessageIterator struct {
	Event *HyperlaneOptimismISMReceivedMessage // Event containing the contract specifics and raw log

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
func (it *HyperlaneOptimismISMReceivedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneOptimismISMReceivedMessage)
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
		it.Event = new(HyperlaneOptimismISMReceivedMessage)
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
func (it *HyperlaneOptimismISMReceivedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneOptimismISMReceivedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneOptimismISMReceivedMessage represents a ReceivedMessage event raised by the HyperlaneOptimismISM contract.
type HyperlaneOptimismISMReceivedMessage struct {
	Sender    [32]byte
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceivedMessage is a free log retrieval operation binding the contract event 0xf9347035c6654873e4aa34118c96b53189dc92edcb628c7bad1f9bce9089ecfe.
//
// Solidity: event ReceivedMessage(bytes32 indexed sender, bytes32 indexed messageId)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMFilterer) FilterReceivedMessage(opts *bind.FilterOpts, sender [][32]byte, messageId [][32]byte) (*HyperlaneOptimismISMReceivedMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _HyperlaneOptimismISM.contract.FilterLogs(opts, "ReceivedMessage", senderRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneOptimismISMReceivedMessageIterator{contract: _HyperlaneOptimismISM.contract, event: "ReceivedMessage", logs: logs, sub: sub}, nil
}

// WatchReceivedMessage is a free log subscription operation binding the contract event 0xf9347035c6654873e4aa34118c96b53189dc92edcb628c7bad1f9bce9089ecfe.
//
// Solidity: event ReceivedMessage(bytes32 indexed sender, bytes32 indexed messageId)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMFilterer) WatchReceivedMessage(opts *bind.WatchOpts, sink chan<- *HyperlaneOptimismISMReceivedMessage, sender [][32]byte, messageId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _HyperlaneOptimismISM.contract.WatchLogs(opts, "ReceivedMessage", senderRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneOptimismISMReceivedMessage)
				if err := _HyperlaneOptimismISM.contract.UnpackLog(event, "ReceivedMessage", log); err != nil {
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

// ParseReceivedMessage is a log parse operation binding the contract event 0xf9347035c6654873e4aa34118c96b53189dc92edcb628c7bad1f9bce9089ecfe.
//
// Solidity: event ReceivedMessage(bytes32 indexed sender, bytes32 indexed messageId)
func (_HyperlaneOptimismISM *HyperlaneOptimismISMFilterer) ParseReceivedMessage(log types.Log) (*HyperlaneOptimismISMReceivedMessage, error) {
	event := new(HyperlaneOptimismISMReceivedMessage)
	if err := _HyperlaneOptimismISM.contract.UnpackLog(event, "ReceivedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
