// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MatchingEngine

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
	_ = abi.ConvertType
)

// LibAssetAsset is an auto generated low-level Go binding around an user-defined struct.
type LibAssetAsset struct {
	VirtualToken common.Address
	Value        *big.Int
}

// LibFillFillResult is an auto generated low-level Go binding around an user-defined struct.
type LibFillFillResult struct {
	LeftValue  *big.Int
	RightValue *big.Int
}

// LibOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type LibOrderOrder struct {
	OrderType              [4]byte
	Deadline               uint64
	Trader                 common.Address
	MakeAsset              LibAssetAsset
	TakeAsset              LibAssetAsset
	Salt                   *big.Int
	LimitOrderTriggerPrice *big.Int
	IsShort                bool
}

// MatchingEngineMetaData contains all meta data concerning the MatchingEngine contract.
var MatchingEngineMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minSalt\",\"type\":\"uint256\"}],\"name\":\"CanceledAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[2]\",\"name\":\"traders\",\"type\":\"address[2]\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"deadline\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"salt\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLeftFill\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRightFill\",\"type\":\"uint256\"}],\"name\":\"Matched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[2]\",\"name\":\"traders\",\"type\":\"address[2]\"},{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"salts\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"fills\",\"type\":\"uint256[2]\"}],\"name\":\"OrdersFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CAN_MATCH_ORDERS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MATCHING_ENGINE_CORE_ADMIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minSalt\",\"type\":\"uint256\"}],\"name\":\"cancelAllOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrdersInBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fills\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantMatchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"makerMinSalt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"markPriceOracle\",\"outputs\":[{\"internalType\":\"contractIMarkPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"ordersLeft\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"ordersRight\",\"type\":\"tuple[]\"}],\"name\":\"matchOrderInBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"orderLeft\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"orderRight\",\"type\":\"tuple\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"leftValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightValue\",\"type\":\"uint256\"}],\"internalType\":\"structLibFill.FillResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MatchingEngineABI is the input ABI used to generate the binding from.
// Deprecated: Use MatchingEngineMetaData.ABI instead.
var MatchingEngineABI = MatchingEngineMetaData.ABI

// MatchingEngine is an auto generated Go binding around an Ethereum contract.
type MatchingEngine struct {
	MatchingEngineCaller     // Read-only binding to the contract
	MatchingEngineTransactor // Write-only binding to the contract
	MatchingEngineFilterer   // Log filterer for contract events
}

// MatchingEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type MatchingEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchingEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MatchingEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchingEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MatchingEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchingEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MatchingEngineSession struct {
	Contract     *MatchingEngine   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MatchingEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MatchingEngineCallerSession struct {
	Contract *MatchingEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MatchingEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MatchingEngineTransactorSession struct {
	Contract     *MatchingEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MatchingEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type MatchingEngineRaw struct {
	Contract *MatchingEngine // Generic contract binding to access the raw methods on
}

// MatchingEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MatchingEngineCallerRaw struct {
	Contract *MatchingEngineCaller // Generic read-only contract binding to access the raw methods on
}

// MatchingEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MatchingEngineTransactorRaw struct {
	Contract *MatchingEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMatchingEngine creates a new instance of MatchingEngine, bound to a specific deployed contract.
func NewMatchingEngine(address common.Address, backend bind.ContractBackend) (*MatchingEngine, error) {
	contract, err := bindMatchingEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MatchingEngine{MatchingEngineCaller: MatchingEngineCaller{contract: contract}, MatchingEngineTransactor: MatchingEngineTransactor{contract: contract}, MatchingEngineFilterer: MatchingEngineFilterer{contract: contract}}, nil
}

// NewMatchingEngineCaller creates a new read-only instance of MatchingEngine, bound to a specific deployed contract.
func NewMatchingEngineCaller(address common.Address, caller bind.ContractCaller) (*MatchingEngineCaller, error) {
	contract, err := bindMatchingEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MatchingEngineCaller{contract: contract}, nil
}

// NewMatchingEngineTransactor creates a new write-only instance of MatchingEngine, bound to a specific deployed contract.
func NewMatchingEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*MatchingEngineTransactor, error) {
	contract, err := bindMatchingEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MatchingEngineTransactor{contract: contract}, nil
}

// NewMatchingEngineFilterer creates a new log filterer instance of MatchingEngine, bound to a specific deployed contract.
func NewMatchingEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*MatchingEngineFilterer, error) {
	contract, err := bindMatchingEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MatchingEngineFilterer{contract: contract}, nil
}

// bindMatchingEngine binds a generic wrapper to an already deployed contract.
func bindMatchingEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MatchingEngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MatchingEngine *MatchingEngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MatchingEngine.Contract.MatchingEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MatchingEngine *MatchingEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MatchingEngine.Contract.MatchingEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MatchingEngine *MatchingEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MatchingEngine.Contract.MatchingEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MatchingEngine *MatchingEngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MatchingEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MatchingEngine *MatchingEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MatchingEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MatchingEngine *MatchingEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MatchingEngine.Contract.contract.Transact(opts, method, params...)
}

// CANMATCHORDERS is a free data retrieval call binding the contract method 0xe7ddc9d0.
//
// Solidity: function CAN_MATCH_ORDERS() view returns(bytes32)
func (_MatchingEngine *MatchingEngineCaller) CANMATCHORDERS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "CAN_MATCH_ORDERS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANMATCHORDERS is a free data retrieval call binding the contract method 0xe7ddc9d0.
//
// Solidity: function CAN_MATCH_ORDERS() view returns(bytes32)
func (_MatchingEngine *MatchingEngineSession) CANMATCHORDERS() ([32]byte, error) {
	return _MatchingEngine.Contract.CANMATCHORDERS(&_MatchingEngine.CallOpts)
}

// CANMATCHORDERS is a free data retrieval call binding the contract method 0xe7ddc9d0.
//
// Solidity: function CAN_MATCH_ORDERS() view returns(bytes32)
func (_MatchingEngine *MatchingEngineCallerSession) CANMATCHORDERS() ([32]byte, error) {
	return _MatchingEngine.Contract.CANMATCHORDERS(&_MatchingEngine.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MatchingEngine *MatchingEngineCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MatchingEngine *MatchingEngineSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MatchingEngine.Contract.DEFAULTADMINROLE(&_MatchingEngine.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MatchingEngine *MatchingEngineCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MatchingEngine.Contract.DEFAULTADMINROLE(&_MatchingEngine.CallOpts)
}

// MATCHINGENGINECOREADMIN is a free data retrieval call binding the contract method 0xbf611dc6.
//
// Solidity: function MATCHING_ENGINE_CORE_ADMIN() view returns(bytes32)
func (_MatchingEngine *MatchingEngineCaller) MATCHINGENGINECOREADMIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "MATCHING_ENGINE_CORE_ADMIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MATCHINGENGINECOREADMIN is a free data retrieval call binding the contract method 0xbf611dc6.
//
// Solidity: function MATCHING_ENGINE_CORE_ADMIN() view returns(bytes32)
func (_MatchingEngine *MatchingEngineSession) MATCHINGENGINECOREADMIN() ([32]byte, error) {
	return _MatchingEngine.Contract.MATCHINGENGINECOREADMIN(&_MatchingEngine.CallOpts)
}

// MATCHINGENGINECOREADMIN is a free data retrieval call binding the contract method 0xbf611dc6.
//
// Solidity: function MATCHING_ENGINE_CORE_ADMIN() view returns(bytes32)
func (_MatchingEngine *MatchingEngineCallerSession) MATCHINGENGINECOREADMIN() ([32]byte, error) {
	return _MatchingEngine.Contract.MATCHINGENGINECOREADMIN(&_MatchingEngine.CallOpts)
}

// Fills is a free data retrieval call binding the contract method 0x20158c44.
//
// Solidity: function fills(bytes32 ) view returns(uint256)
func (_MatchingEngine *MatchingEngineCaller) Fills(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "fills", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fills is a free data retrieval call binding the contract method 0x20158c44.
//
// Solidity: function fills(bytes32 ) view returns(uint256)
func (_MatchingEngine *MatchingEngineSession) Fills(arg0 [32]byte) (*big.Int, error) {
	return _MatchingEngine.Contract.Fills(&_MatchingEngine.CallOpts, arg0)
}

// Fills is a free data retrieval call binding the contract method 0x20158c44.
//
// Solidity: function fills(bytes32 ) view returns(uint256)
func (_MatchingEngine *MatchingEngineCallerSession) Fills(arg0 [32]byte) (*big.Int, error) {
	return _MatchingEngine.Contract.Fills(&_MatchingEngine.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MatchingEngine *MatchingEngineCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MatchingEngine *MatchingEngineSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MatchingEngine.Contract.GetRoleAdmin(&_MatchingEngine.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MatchingEngine *MatchingEngineCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MatchingEngine.Contract.GetRoleAdmin(&_MatchingEngine.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MatchingEngine *MatchingEngineCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MatchingEngine *MatchingEngineSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MatchingEngine.Contract.HasRole(&_MatchingEngine.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MatchingEngine *MatchingEngineCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MatchingEngine.Contract.HasRole(&_MatchingEngine.CallOpts, role, account)
}

// MakerMinSalt is a free data retrieval call binding the contract method 0x624d7600.
//
// Solidity: function makerMinSalt(address ) view returns(uint256)
func (_MatchingEngine *MatchingEngineCaller) MakerMinSalt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "makerMinSalt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MakerMinSalt is a free data retrieval call binding the contract method 0x624d7600.
//
// Solidity: function makerMinSalt(address ) view returns(uint256)
func (_MatchingEngine *MatchingEngineSession) MakerMinSalt(arg0 common.Address) (*big.Int, error) {
	return _MatchingEngine.Contract.MakerMinSalt(&_MatchingEngine.CallOpts, arg0)
}

// MakerMinSalt is a free data retrieval call binding the contract method 0x624d7600.
//
// Solidity: function makerMinSalt(address ) view returns(uint256)
func (_MatchingEngine *MatchingEngineCallerSession) MakerMinSalt(arg0 common.Address) (*big.Int, error) {
	return _MatchingEngine.Contract.MakerMinSalt(&_MatchingEngine.CallOpts, arg0)
}

// MarkPriceOracle is a free data retrieval call binding the contract method 0x00079bb2.
//
// Solidity: function markPriceOracle() view returns(address)
func (_MatchingEngine *MatchingEngineCaller) MarkPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "markPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarkPriceOracle is a free data retrieval call binding the contract method 0x00079bb2.
//
// Solidity: function markPriceOracle() view returns(address)
func (_MatchingEngine *MatchingEngineSession) MarkPriceOracle() (common.Address, error) {
	return _MatchingEngine.Contract.MarkPriceOracle(&_MatchingEngine.CallOpts)
}

// MarkPriceOracle is a free data retrieval call binding the contract method 0x00079bb2.
//
// Solidity: function markPriceOracle() view returns(address)
func (_MatchingEngine *MatchingEngineCallerSession) MarkPriceOracle() (common.Address, error) {
	return _MatchingEngine.Contract.MarkPriceOracle(&_MatchingEngine.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MatchingEngine *MatchingEngineCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MatchingEngine *MatchingEngineSession) Paused() (bool, error) {
	return _MatchingEngine.Contract.Paused(&_MatchingEngine.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MatchingEngine *MatchingEngineCallerSession) Paused() (bool, error) {
	return _MatchingEngine.Contract.Paused(&_MatchingEngine.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MatchingEngine *MatchingEngineCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MatchingEngine.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MatchingEngine *MatchingEngineSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MatchingEngine.Contract.SupportsInterface(&_MatchingEngine.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MatchingEngine *MatchingEngineCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MatchingEngine.Contract.SupportsInterface(&_MatchingEngine.CallOpts, interfaceId)
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0xbd545f53.
//
// Solidity: function cancelAllOrders(uint256 minSalt) returns()
func (_MatchingEngine *MatchingEngineTransactor) CancelAllOrders(opts *bind.TransactOpts, minSalt *big.Int) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "cancelAllOrders", minSalt)
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0xbd545f53.
//
// Solidity: function cancelAllOrders(uint256 minSalt) returns()
func (_MatchingEngine *MatchingEngineSession) CancelAllOrders(minSalt *big.Int) (*types.Transaction, error) {
	return _MatchingEngine.Contract.CancelAllOrders(&_MatchingEngine.TransactOpts, minSalt)
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0xbd545f53.
//
// Solidity: function cancelAllOrders(uint256 minSalt) returns()
func (_MatchingEngine *MatchingEngineTransactorSession) CancelAllOrders(minSalt *big.Int) (*types.Transaction, error) {
	return _MatchingEngine.Contract.CancelAllOrders(&_MatchingEngine.TransactOpts, minSalt)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf505f6ec.
//
// Solidity: function cancelOrder((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) order) returns()
func (_MatchingEngine *MatchingEngineTransactor) CancelOrder(opts *bind.TransactOpts, order LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf505f6ec.
//
// Solidity: function cancelOrder((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) order) returns()
func (_MatchingEngine *MatchingEngineSession) CancelOrder(order LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.Contract.CancelOrder(&_MatchingEngine.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf505f6ec.
//
// Solidity: function cancelOrder((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) order) returns()
func (_MatchingEngine *MatchingEngineTransactorSession) CancelOrder(order LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.Contract.CancelOrder(&_MatchingEngine.TransactOpts, order)
}

// CancelOrdersInBatch is a paid mutator transaction binding the contract method 0xb107dd86.
//
// Solidity: function cancelOrdersInBatch((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] orders) returns()
func (_MatchingEngine *MatchingEngineTransactor) CancelOrdersInBatch(opts *bind.TransactOpts, orders []LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "cancelOrdersInBatch", orders)
}

// CancelOrdersInBatch is a paid mutator transaction binding the contract method 0xb107dd86.
//
// Solidity: function cancelOrdersInBatch((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] orders) returns()
func (_MatchingEngine *MatchingEngineSession) CancelOrdersInBatch(orders []LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.Contract.CancelOrdersInBatch(&_MatchingEngine.TransactOpts, orders)
}

// CancelOrdersInBatch is a paid mutator transaction binding the contract method 0xb107dd86.
//
// Solidity: function cancelOrdersInBatch((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] orders) returns()
func (_MatchingEngine *MatchingEngineTransactorSession) CancelOrdersInBatch(orders []LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.Contract.CancelOrdersInBatch(&_MatchingEngine.TransactOpts, orders)
}

// GrantMatchOrders is a paid mutator transaction binding the contract method 0xabd13933.
//
// Solidity: function grantMatchOrders(address account) returns()
func (_MatchingEngine *MatchingEngineTransactor) GrantMatchOrders(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "grantMatchOrders", account)
}

// GrantMatchOrders is a paid mutator transaction binding the contract method 0xabd13933.
//
// Solidity: function grantMatchOrders(address account) returns()
func (_MatchingEngine *MatchingEngineSession) GrantMatchOrders(account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.Contract.GrantMatchOrders(&_MatchingEngine.TransactOpts, account)
}

// GrantMatchOrders is a paid mutator transaction binding the contract method 0xabd13933.
//
// Solidity: function grantMatchOrders(address account) returns()
func (_MatchingEngine *MatchingEngineTransactorSession) GrantMatchOrders(account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.Contract.GrantMatchOrders(&_MatchingEngine.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.Contract.GrantRole(&_MatchingEngine.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.Contract.GrantRole(&_MatchingEngine.TransactOpts, role, account)
}

// MatchOrderInBatch is a paid mutator transaction binding the contract method 0x873852dc.
//
// Solidity: function matchOrderInBatch((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] ordersLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] ordersRight) returns()
func (_MatchingEngine *MatchingEngineTransactor) MatchOrderInBatch(opts *bind.TransactOpts, ordersLeft []LibOrderOrder, ordersRight []LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "matchOrderInBatch", ordersLeft, ordersRight)
}

// MatchOrderInBatch is a paid mutator transaction binding the contract method 0x873852dc.
//
// Solidity: function matchOrderInBatch((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] ordersLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] ordersRight) returns()
func (_MatchingEngine *MatchingEngineSession) MatchOrderInBatch(ordersLeft []LibOrderOrder, ordersRight []LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.Contract.MatchOrderInBatch(&_MatchingEngine.TransactOpts, ordersLeft, ordersRight)
}

// MatchOrderInBatch is a paid mutator transaction binding the contract method 0x873852dc.
//
// Solidity: function matchOrderInBatch((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] ordersLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] ordersRight) returns()
func (_MatchingEngine *MatchingEngineTransactorSession) MatchOrderInBatch(ordersLeft []LibOrderOrder, ordersRight []LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.Contract.MatchOrderInBatch(&_MatchingEngine.TransactOpts, ordersLeft, ordersRight)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x6ab6a266.
//
// Solidity: function matchOrders((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderRight) returns((uint256,uint256))
func (_MatchingEngine *MatchingEngineTransactor) MatchOrders(opts *bind.TransactOpts, orderLeft LibOrderOrder, orderRight LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "matchOrders", orderLeft, orderRight)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x6ab6a266.
//
// Solidity: function matchOrders((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderRight) returns((uint256,uint256))
func (_MatchingEngine *MatchingEngineSession) MatchOrders(orderLeft LibOrderOrder, orderRight LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.Contract.MatchOrders(&_MatchingEngine.TransactOpts, orderLeft, orderRight)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x6ab6a266.
//
// Solidity: function matchOrders((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderRight) returns((uint256,uint256))
func (_MatchingEngine *MatchingEngineTransactorSession) MatchOrders(orderLeft LibOrderOrder, orderRight LibOrderOrder) (*types.Transaction, error) {
	return _MatchingEngine.Contract.MatchOrders(&_MatchingEngine.TransactOpts, orderLeft, orderRight)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.Contract.RenounceRole(&_MatchingEngine.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.Contract.RenounceRole(&_MatchingEngine.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.Contract.RevokeRole(&_MatchingEngine.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MatchingEngine *MatchingEngineTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MatchingEngine.Contract.RevokeRole(&_MatchingEngine.TransactOpts, role, account)
}

// MatchingEngineCanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the MatchingEngine contract.
type MatchingEngineCanceledIterator struct {
	Event *MatchingEngineCanceled // Event containing the contract specifics and raw log

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
func (it *MatchingEngineCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineCanceled)
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
		it.Event = new(MatchingEngineCanceled)
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
func (it *MatchingEngineCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineCanceled represents a Canceled event raised by the MatchingEngine contract.
type MatchingEngineCanceled struct {
	Hash      [32]byte
	Trader    common.Address
	BaseToken common.Address
	Amount    *big.Int
	Salt      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0x5e1752996b0d2eba5c21c58babc2ffd99710c0f531315cfd7bec221ef831c3dd.
//
// Solidity: event Canceled(bytes32 indexed hash, address trader, address baseToken, uint256 amount, uint256 salt)
func (_MatchingEngine *MatchingEngineFilterer) FilterCanceled(opts *bind.FilterOpts, hash [][32]byte) (*MatchingEngineCanceledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "Canceled", hashRule)
	if err != nil {
		return nil, err
	}
	return &MatchingEngineCanceledIterator{contract: _MatchingEngine.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0x5e1752996b0d2eba5c21c58babc2ffd99710c0f531315cfd7bec221ef831c3dd.
//
// Solidity: event Canceled(bytes32 indexed hash, address trader, address baseToken, uint256 amount, uint256 salt)
func (_MatchingEngine *MatchingEngineFilterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *MatchingEngineCanceled, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "Canceled", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineCanceled)
				if err := _MatchingEngine.contract.UnpackLog(event, "Canceled", log); err != nil {
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

// ParseCanceled is a log parse operation binding the contract event 0x5e1752996b0d2eba5c21c58babc2ffd99710c0f531315cfd7bec221ef831c3dd.
//
// Solidity: event Canceled(bytes32 indexed hash, address trader, address baseToken, uint256 amount, uint256 salt)
func (_MatchingEngine *MatchingEngineFilterer) ParseCanceled(log types.Log) (*MatchingEngineCanceled, error) {
	event := new(MatchingEngineCanceled)
	if err := _MatchingEngine.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEngineCanceledAllIterator is returned from FilterCanceledAll and is used to iterate over the raw logs and unpacked data for CanceledAll events raised by the MatchingEngine contract.
type MatchingEngineCanceledAllIterator struct {
	Event *MatchingEngineCanceledAll // Event containing the contract specifics and raw log

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
func (it *MatchingEngineCanceledAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineCanceledAll)
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
		it.Event = new(MatchingEngineCanceledAll)
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
func (it *MatchingEngineCanceledAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineCanceledAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineCanceledAll represents a CanceledAll event raised by the MatchingEngine contract.
type MatchingEngineCanceledAll struct {
	Trader  common.Address
	MinSalt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCanceledAll is a free log retrieval operation binding the contract event 0x8297ed9ba0d1dbf6f3de44e79b70af432d565b076408158b5fc59ebc58cc771a.
//
// Solidity: event CanceledAll(address indexed trader, uint256 minSalt)
func (_MatchingEngine *MatchingEngineFilterer) FilterCanceledAll(opts *bind.FilterOpts, trader []common.Address) (*MatchingEngineCanceledAllIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "CanceledAll", traderRule)
	if err != nil {
		return nil, err
	}
	return &MatchingEngineCanceledAllIterator{contract: _MatchingEngine.contract, event: "CanceledAll", logs: logs, sub: sub}, nil
}

// WatchCanceledAll is a free log subscription operation binding the contract event 0x8297ed9ba0d1dbf6f3de44e79b70af432d565b076408158b5fc59ebc58cc771a.
//
// Solidity: event CanceledAll(address indexed trader, uint256 minSalt)
func (_MatchingEngine *MatchingEngineFilterer) WatchCanceledAll(opts *bind.WatchOpts, sink chan<- *MatchingEngineCanceledAll, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "CanceledAll", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineCanceledAll)
				if err := _MatchingEngine.contract.UnpackLog(event, "CanceledAll", log); err != nil {
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

// ParseCanceledAll is a log parse operation binding the contract event 0x8297ed9ba0d1dbf6f3de44e79b70af432d565b076408158b5fc59ebc58cc771a.
//
// Solidity: event CanceledAll(address indexed trader, uint256 minSalt)
func (_MatchingEngine *MatchingEngineFilterer) ParseCanceledAll(log types.Log) (*MatchingEngineCanceledAll, error) {
	event := new(MatchingEngineCanceledAll)
	if err := _MatchingEngine.contract.UnpackLog(event, "CanceledAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEngineInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MatchingEngine contract.
type MatchingEngineInitializedIterator struct {
	Event *MatchingEngineInitialized // Event containing the contract specifics and raw log

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
func (it *MatchingEngineInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineInitialized)
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
		it.Event = new(MatchingEngineInitialized)
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
func (it *MatchingEngineInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineInitialized represents a Initialized event raised by the MatchingEngine contract.
type MatchingEngineInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MatchingEngine *MatchingEngineFilterer) FilterInitialized(opts *bind.FilterOpts) (*MatchingEngineInitializedIterator, error) {

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MatchingEngineInitializedIterator{contract: _MatchingEngine.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MatchingEngine *MatchingEngineFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MatchingEngineInitialized) (event.Subscription, error) {

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineInitialized)
				if err := _MatchingEngine.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MatchingEngine *MatchingEngineFilterer) ParseInitialized(log types.Log) (*MatchingEngineInitialized, error) {
	event := new(MatchingEngineInitialized)
	if err := _MatchingEngine.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEngineMatchedIterator is returned from FilterMatched and is used to iterate over the raw logs and unpacked data for Matched events raised by the MatchingEngine contract.
type MatchingEngineMatchedIterator struct {
	Event *MatchingEngineMatched // Event containing the contract specifics and raw log

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
func (it *MatchingEngineMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineMatched)
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
		it.Event = new(MatchingEngineMatched)
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
func (it *MatchingEngineMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineMatched represents a Matched event raised by the MatchingEngine contract.
type MatchingEngineMatched struct {
	Traders      [2]common.Address
	Deadline     [2]uint64
	Salt         [2]*big.Int
	NewLeftFill  *big.Int
	NewRightFill *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMatched is a free log retrieval operation binding the contract event 0x48341121016d0377a93805184a648921d7cf78e9378e868e3f16dc500d770a48.
//
// Solidity: event Matched(address[2] traders, uint64[2] deadline, uint256[2] salt, uint256 newLeftFill, uint256 newRightFill)
func (_MatchingEngine *MatchingEngineFilterer) FilterMatched(opts *bind.FilterOpts) (*MatchingEngineMatchedIterator, error) {

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "Matched")
	if err != nil {
		return nil, err
	}
	return &MatchingEngineMatchedIterator{contract: _MatchingEngine.contract, event: "Matched", logs: logs, sub: sub}, nil
}

// WatchMatched is a free log subscription operation binding the contract event 0x48341121016d0377a93805184a648921d7cf78e9378e868e3f16dc500d770a48.
//
// Solidity: event Matched(address[2] traders, uint64[2] deadline, uint256[2] salt, uint256 newLeftFill, uint256 newRightFill)
func (_MatchingEngine *MatchingEngineFilterer) WatchMatched(opts *bind.WatchOpts, sink chan<- *MatchingEngineMatched) (event.Subscription, error) {

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "Matched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineMatched)
				if err := _MatchingEngine.contract.UnpackLog(event, "Matched", log); err != nil {
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

// ParseMatched is a log parse operation binding the contract event 0x48341121016d0377a93805184a648921d7cf78e9378e868e3f16dc500d770a48.
//
// Solidity: event Matched(address[2] traders, uint64[2] deadline, uint256[2] salt, uint256 newLeftFill, uint256 newRightFill)
func (_MatchingEngine *MatchingEngineFilterer) ParseMatched(log types.Log) (*MatchingEngineMatched, error) {
	event := new(MatchingEngineMatched)
	if err := _MatchingEngine.contract.UnpackLog(event, "Matched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEngineOrdersFilledIterator is returned from FilterOrdersFilled and is used to iterate over the raw logs and unpacked data for OrdersFilled events raised by the MatchingEngine contract.
type MatchingEngineOrdersFilledIterator struct {
	Event *MatchingEngineOrdersFilled // Event containing the contract specifics and raw log

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
func (it *MatchingEngineOrdersFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineOrdersFilled)
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
		it.Event = new(MatchingEngineOrdersFilled)
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
func (it *MatchingEngineOrdersFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineOrdersFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineOrdersFilled represents a OrdersFilled event raised by the MatchingEngine contract.
type MatchingEngineOrdersFilled struct {
	Traders [2]common.Address
	Salts   [2]*big.Int
	Fills   [2]*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrdersFilled is a free log retrieval operation binding the contract event 0xb9d2f2ffd6d04a7450b31bb6d20022f66a28e2219e8ed64ba56e199316530129.
//
// Solidity: event OrdersFilled(address[2] traders, uint256[2] salts, uint256[2] fills)
func (_MatchingEngine *MatchingEngineFilterer) FilterOrdersFilled(opts *bind.FilterOpts) (*MatchingEngineOrdersFilledIterator, error) {

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "OrdersFilled")
	if err != nil {
		return nil, err
	}
	return &MatchingEngineOrdersFilledIterator{contract: _MatchingEngine.contract, event: "OrdersFilled", logs: logs, sub: sub}, nil
}

// WatchOrdersFilled is a free log subscription operation binding the contract event 0xb9d2f2ffd6d04a7450b31bb6d20022f66a28e2219e8ed64ba56e199316530129.
//
// Solidity: event OrdersFilled(address[2] traders, uint256[2] salts, uint256[2] fills)
func (_MatchingEngine *MatchingEngineFilterer) WatchOrdersFilled(opts *bind.WatchOpts, sink chan<- *MatchingEngineOrdersFilled) (event.Subscription, error) {

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "OrdersFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineOrdersFilled)
				if err := _MatchingEngine.contract.UnpackLog(event, "OrdersFilled", log); err != nil {
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

// ParseOrdersFilled is a log parse operation binding the contract event 0xb9d2f2ffd6d04a7450b31bb6d20022f66a28e2219e8ed64ba56e199316530129.
//
// Solidity: event OrdersFilled(address[2] traders, uint256[2] salts, uint256[2] fills)
func (_MatchingEngine *MatchingEngineFilterer) ParseOrdersFilled(log types.Log) (*MatchingEngineOrdersFilled, error) {
	event := new(MatchingEngineOrdersFilled)
	if err := _MatchingEngine.contract.UnpackLog(event, "OrdersFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEnginePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MatchingEngine contract.
type MatchingEnginePausedIterator struct {
	Event *MatchingEnginePaused // Event containing the contract specifics and raw log

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
func (it *MatchingEnginePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEnginePaused)
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
		it.Event = new(MatchingEnginePaused)
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
func (it *MatchingEnginePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEnginePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEnginePaused represents a Paused event raised by the MatchingEngine contract.
type MatchingEnginePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MatchingEngine *MatchingEngineFilterer) FilterPaused(opts *bind.FilterOpts) (*MatchingEnginePausedIterator, error) {

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MatchingEnginePausedIterator{contract: _MatchingEngine.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MatchingEngine *MatchingEngineFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MatchingEnginePaused) (event.Subscription, error) {

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEnginePaused)
				if err := _MatchingEngine.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MatchingEngine *MatchingEngineFilterer) ParsePaused(log types.Log) (*MatchingEnginePaused, error) {
	event := new(MatchingEnginePaused)
	if err := _MatchingEngine.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEngineRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MatchingEngine contract.
type MatchingEngineRoleAdminChangedIterator struct {
	Event *MatchingEngineRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MatchingEngineRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineRoleAdminChanged)
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
		it.Event = new(MatchingEngineRoleAdminChanged)
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
func (it *MatchingEngineRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineRoleAdminChanged represents a RoleAdminChanged event raised by the MatchingEngine contract.
type MatchingEngineRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MatchingEngine *MatchingEngineFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MatchingEngineRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MatchingEngineRoleAdminChangedIterator{contract: _MatchingEngine.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MatchingEngine *MatchingEngineFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MatchingEngineRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineRoleAdminChanged)
				if err := _MatchingEngine.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MatchingEngine *MatchingEngineFilterer) ParseRoleAdminChanged(log types.Log) (*MatchingEngineRoleAdminChanged, error) {
	event := new(MatchingEngineRoleAdminChanged)
	if err := _MatchingEngine.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEngineRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MatchingEngine contract.
type MatchingEngineRoleGrantedIterator struct {
	Event *MatchingEngineRoleGranted // Event containing the contract specifics and raw log

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
func (it *MatchingEngineRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineRoleGranted)
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
		it.Event = new(MatchingEngineRoleGranted)
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
func (it *MatchingEngineRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineRoleGranted represents a RoleGranted event raised by the MatchingEngine contract.
type MatchingEngineRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MatchingEngine *MatchingEngineFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MatchingEngineRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MatchingEngineRoleGrantedIterator{contract: _MatchingEngine.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MatchingEngine *MatchingEngineFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MatchingEngineRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineRoleGranted)
				if err := _MatchingEngine.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MatchingEngine *MatchingEngineFilterer) ParseRoleGranted(log types.Log) (*MatchingEngineRoleGranted, error) {
	event := new(MatchingEngineRoleGranted)
	if err := _MatchingEngine.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEngineRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MatchingEngine contract.
type MatchingEngineRoleRevokedIterator struct {
	Event *MatchingEngineRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MatchingEngineRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineRoleRevoked)
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
		it.Event = new(MatchingEngineRoleRevoked)
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
func (it *MatchingEngineRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineRoleRevoked represents a RoleRevoked event raised by the MatchingEngine contract.
type MatchingEngineRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MatchingEngine *MatchingEngineFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MatchingEngineRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MatchingEngineRoleRevokedIterator{contract: _MatchingEngine.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MatchingEngine *MatchingEngineFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MatchingEngineRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineRoleRevoked)
				if err := _MatchingEngine.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MatchingEngine *MatchingEngineFilterer) ParseRoleRevoked(log types.Log) (*MatchingEngineRoleRevoked, error) {
	event := new(MatchingEngineRoleRevoked)
	if err := _MatchingEngine.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchingEngineUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MatchingEngine contract.
type MatchingEngineUnpausedIterator struct {
	Event *MatchingEngineUnpaused // Event containing the contract specifics and raw log

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
func (it *MatchingEngineUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchingEngineUnpaused)
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
		it.Event = new(MatchingEngineUnpaused)
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
func (it *MatchingEngineUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchingEngineUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchingEngineUnpaused represents a Unpaused event raised by the MatchingEngine contract.
type MatchingEngineUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MatchingEngine *MatchingEngineFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MatchingEngineUnpausedIterator, error) {

	logs, sub, err := _MatchingEngine.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MatchingEngineUnpausedIterator{contract: _MatchingEngine.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MatchingEngine *MatchingEngineFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MatchingEngineUnpaused) (event.Subscription, error) {

	logs, sub, err := _MatchingEngine.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchingEngineUnpaused)
				if err := _MatchingEngine.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MatchingEngine *MatchingEngineFilterer) ParseUnpaused(log types.Log) (*MatchingEngineUnpaused, error) {
	event := new(MatchingEngineUnpaused)
	if err := _MatchingEngine.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
