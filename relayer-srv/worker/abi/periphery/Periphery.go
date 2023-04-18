// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Periphery

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

// PeripheryMetaData contains all meta data concerning the Periphery contract.
var PeripheryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRelayerAddress\",\"type\":\"address\"}],\"name\":\"RelayerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWhitelist\",\"type\":\"bool\"}],\"name\":\"TraderWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWhitelist\",\"type\":\"bool\"}],\"name\":\"VaultWhitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_MULTISIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRADER_WHITELISTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOLMEX_PERP_PERIPHERY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"_ordersLeft\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signaturesLeft\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"_ordersRight\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signaturesRight\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"liquidator\",\"type\":\"bytes\"}],\"name\":\"batchOpenPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositToVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexPriceOracle\",\"outputs\":[{\"internalType\":\"contractIIndexPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVolmexPerpView\",\"name\":\"_perpView\",\"type\":\"address\"},{\"internalType\":\"contractIMarkPriceOracle\",\"name\":\"_markPriceOracle\",\"type\":\"address\"},{\"internalType\":\"contractIIndexPriceOracle\",\"name\":\"_indexPriceOracle\",\"type\":\"address\"},{\"internalType\":\"address[2]\",\"name\":\"_vaults\",\"type\":\"address[2]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTraderWhitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTraderWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"markPriceOracle\",\"outputs\":[{\"internalType\":\"contractIMarkPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"_orderLeft\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signatureLeft\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"_orderRight\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signatureRight\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"liquidator\",\"type\":\"bytes\"}],\"name\":\"openPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"perpView\",\"outputs\":[{\"internalType\":\"contractIVolmexPerpView\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMarkPriceOracle\",\"name\":\"_markPriceOracle\",\"type\":\"address\"}],\"name\":\"setMarkPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleTraderWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferToVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isWhitelist\",\"type\":\"bool\"}],\"name\":\"whitelistTrader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isWhitelist\",\"type\":\"bool\"}],\"name\":\"whitelistVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PeripheryABI is the input ABI used to generate the binding from.
// Deprecated: Use PeripheryMetaData.ABI instead.
var PeripheryABI = PeripheryMetaData.ABI

// Periphery is an auto generated Go binding around an Ethereum contract.
type Periphery struct {
	PeripheryCaller     // Read-only binding to the contract
	PeripheryTransactor // Write-only binding to the contract
	PeripheryFilterer   // Log filterer for contract events
}

// PeripheryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeripheryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeripheryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeripheryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripherySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeripherySession struct {
	Contract     *Periphery        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeripheryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeripheryCallerSession struct {
	Contract *PeripheryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PeripheryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeripheryTransactorSession struct {
	Contract     *PeripheryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PeripheryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeripheryRaw struct {
	Contract *Periphery // Generic contract binding to access the raw methods on
}

// PeripheryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeripheryCallerRaw struct {
	Contract *PeripheryCaller // Generic read-only contract binding to access the raw methods on
}

// PeripheryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeripheryTransactorRaw struct {
	Contract *PeripheryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeriphery creates a new instance of Periphery, bound to a specific deployed contract.
func NewPeriphery(address common.Address, backend bind.ContractBackend) (*Periphery, error) {
	contract, err := bindPeriphery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Periphery{PeripheryCaller: PeripheryCaller{contract: contract}, PeripheryTransactor: PeripheryTransactor{contract: contract}, PeripheryFilterer: PeripheryFilterer{contract: contract}}, nil
}

// NewPeripheryCaller creates a new read-only instance of Periphery, bound to a specific deployed contract.
func NewPeripheryCaller(address common.Address, caller bind.ContractCaller) (*PeripheryCaller, error) {
	contract, err := bindPeriphery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryCaller{contract: contract}, nil
}

// NewPeripheryTransactor creates a new write-only instance of Periphery, bound to a specific deployed contract.
func NewPeripheryTransactor(address common.Address, transactor bind.ContractTransactor) (*PeripheryTransactor, error) {
	contract, err := bindPeriphery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryTransactor{contract: contract}, nil
}

// NewPeripheryFilterer creates a new log filterer instance of Periphery, bound to a specific deployed contract.
func NewPeripheryFilterer(address common.Address, filterer bind.ContractFilterer) (*PeripheryFilterer, error) {
	contract, err := bindPeriphery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeripheryFilterer{contract: contract}, nil
}

// bindPeriphery binds a generic wrapper to an already deployed contract.
func bindPeriphery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PeripheryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Periphery *PeripheryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Periphery.Contract.PeripheryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Periphery *PeripheryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Periphery.Contract.PeripheryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Periphery *PeripheryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Periphery.Contract.PeripheryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Periphery *PeripheryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Periphery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Periphery *PeripheryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Periphery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Periphery *PeripheryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Periphery.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Periphery *PeripheryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Periphery *PeripherySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Periphery.Contract.DEFAULTADMINROLE(&_Periphery.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Periphery *PeripheryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Periphery.Contract.DEFAULTADMINROLE(&_Periphery.CallOpts)
}

// RELAYERMULTISIG is a free data retrieval call binding the contract method 0x4a2718a1.
//
// Solidity: function RELAYER_MULTISIG() view returns(bytes32)
func (_Periphery *PeripheryCaller) RELAYERMULTISIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "RELAYER_MULTISIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERMULTISIG is a free data retrieval call binding the contract method 0x4a2718a1.
//
// Solidity: function RELAYER_MULTISIG() view returns(bytes32)
func (_Periphery *PeripherySession) RELAYERMULTISIG() ([32]byte, error) {
	return _Periphery.Contract.RELAYERMULTISIG(&_Periphery.CallOpts)
}

// RELAYERMULTISIG is a free data retrieval call binding the contract method 0x4a2718a1.
//
// Solidity: function RELAYER_MULTISIG() view returns(bytes32)
func (_Periphery *PeripheryCallerSession) RELAYERMULTISIG() ([32]byte, error) {
	return _Periphery.Contract.RELAYERMULTISIG(&_Periphery.CallOpts)
}

// TRADERWHITELISTER is a free data retrieval call binding the contract method 0xcb2bba5f.
//
// Solidity: function TRADER_WHITELISTER() view returns(bytes32)
func (_Periphery *PeripheryCaller) TRADERWHITELISTER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "TRADER_WHITELISTER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRADERWHITELISTER is a free data retrieval call binding the contract method 0xcb2bba5f.
//
// Solidity: function TRADER_WHITELISTER() view returns(bytes32)
func (_Periphery *PeripherySession) TRADERWHITELISTER() ([32]byte, error) {
	return _Periphery.Contract.TRADERWHITELISTER(&_Periphery.CallOpts)
}

// TRADERWHITELISTER is a free data retrieval call binding the contract method 0xcb2bba5f.
//
// Solidity: function TRADER_WHITELISTER() view returns(bytes32)
func (_Periphery *PeripheryCallerSession) TRADERWHITELISTER() ([32]byte, error) {
	return _Periphery.Contract.TRADERWHITELISTER(&_Periphery.CallOpts)
}

// VOLMEXPERPPERIPHERY is a free data retrieval call binding the contract method 0x67735165.
//
// Solidity: function VOLMEX_PERP_PERIPHERY() view returns(bytes32)
func (_Periphery *PeripheryCaller) VOLMEXPERPPERIPHERY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "VOLMEX_PERP_PERIPHERY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VOLMEXPERPPERIPHERY is a free data retrieval call binding the contract method 0x67735165.
//
// Solidity: function VOLMEX_PERP_PERIPHERY() view returns(bytes32)
func (_Periphery *PeripherySession) VOLMEXPERPPERIPHERY() ([32]byte, error) {
	return _Periphery.Contract.VOLMEXPERPPERIPHERY(&_Periphery.CallOpts)
}

// VOLMEXPERPPERIPHERY is a free data retrieval call binding the contract method 0x67735165.
//
// Solidity: function VOLMEX_PERP_PERIPHERY() view returns(bytes32)
func (_Periphery *PeripheryCallerSession) VOLMEXPERPPERIPHERY() ([32]byte, error) {
	return _Periphery.Contract.VOLMEXPERPPERIPHERY(&_Periphery.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Periphery *PeripheryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Periphery *PeripherySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Periphery.Contract.GetRoleAdmin(&_Periphery.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Periphery *PeripheryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Periphery.Contract.GetRoleAdmin(&_Periphery.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Periphery *PeripheryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Periphery *PeripherySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Periphery.Contract.HasRole(&_Periphery.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Periphery *PeripheryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Periphery.Contract.HasRole(&_Periphery.CallOpts, role, account)
}

// IndexPriceOracle is a free data retrieval call binding the contract method 0x5434a1df.
//
// Solidity: function indexPriceOracle() view returns(address)
func (_Periphery *PeripheryCaller) IndexPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "indexPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexPriceOracle is a free data retrieval call binding the contract method 0x5434a1df.
//
// Solidity: function indexPriceOracle() view returns(address)
func (_Periphery *PeripherySession) IndexPriceOracle() (common.Address, error) {
	return _Periphery.Contract.IndexPriceOracle(&_Periphery.CallOpts)
}

// IndexPriceOracle is a free data retrieval call binding the contract method 0x5434a1df.
//
// Solidity: function indexPriceOracle() view returns(address)
func (_Periphery *PeripheryCallerSession) IndexPriceOracle() (common.Address, error) {
	return _Periphery.Contract.IndexPriceOracle(&_Periphery.CallOpts)
}

// IsTraderWhitelistEnabled is a free data retrieval call binding the contract method 0xc0b0082d.
//
// Solidity: function isTraderWhitelistEnabled() view returns(bool)
func (_Periphery *PeripheryCaller) IsTraderWhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "isTraderWhitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTraderWhitelistEnabled is a free data retrieval call binding the contract method 0xc0b0082d.
//
// Solidity: function isTraderWhitelistEnabled() view returns(bool)
func (_Periphery *PeripherySession) IsTraderWhitelistEnabled() (bool, error) {
	return _Periphery.Contract.IsTraderWhitelistEnabled(&_Periphery.CallOpts)
}

// IsTraderWhitelistEnabled is a free data retrieval call binding the contract method 0xc0b0082d.
//
// Solidity: function isTraderWhitelistEnabled() view returns(bool)
func (_Periphery *PeripheryCallerSession) IsTraderWhitelistEnabled() (bool, error) {
	return _Periphery.Contract.IsTraderWhitelistEnabled(&_Periphery.CallOpts)
}

// IsTraderWhitelisted is a free data retrieval call binding the contract method 0x6b5a27d1.
//
// Solidity: function isTraderWhitelisted(address ) view returns(bool)
func (_Periphery *PeripheryCaller) IsTraderWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "isTraderWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTraderWhitelisted is a free data retrieval call binding the contract method 0x6b5a27d1.
//
// Solidity: function isTraderWhitelisted(address ) view returns(bool)
func (_Periphery *PeripherySession) IsTraderWhitelisted(arg0 common.Address) (bool, error) {
	return _Periphery.Contract.IsTraderWhitelisted(&_Periphery.CallOpts, arg0)
}

// IsTraderWhitelisted is a free data retrieval call binding the contract method 0x6b5a27d1.
//
// Solidity: function isTraderWhitelisted(address ) view returns(bool)
func (_Periphery *PeripheryCallerSession) IsTraderWhitelisted(arg0 common.Address) (bool, error) {
	return _Periphery.Contract.IsTraderWhitelisted(&_Periphery.CallOpts, arg0)
}

// MarkPriceOracle is a free data retrieval call binding the contract method 0x00079bb2.
//
// Solidity: function markPriceOracle() view returns(address)
func (_Periphery *PeripheryCaller) MarkPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "markPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarkPriceOracle is a free data retrieval call binding the contract method 0x00079bb2.
//
// Solidity: function markPriceOracle() view returns(address)
func (_Periphery *PeripherySession) MarkPriceOracle() (common.Address, error) {
	return _Periphery.Contract.MarkPriceOracle(&_Periphery.CallOpts)
}

// MarkPriceOracle is a free data retrieval call binding the contract method 0x00079bb2.
//
// Solidity: function markPriceOracle() view returns(address)
func (_Periphery *PeripheryCallerSession) MarkPriceOracle() (common.Address, error) {
	return _Periphery.Contract.MarkPriceOracle(&_Periphery.CallOpts)
}

// PerpView is a free data retrieval call binding the contract method 0xa9ed5454.
//
// Solidity: function perpView() view returns(address)
func (_Periphery *PeripheryCaller) PerpView(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "perpView")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PerpView is a free data retrieval call binding the contract method 0xa9ed5454.
//
// Solidity: function perpView() view returns(address)
func (_Periphery *PeripherySession) PerpView() (common.Address, error) {
	return _Periphery.Contract.PerpView(&_Periphery.CallOpts)
}

// PerpView is a free data retrieval call binding the contract method 0xa9ed5454.
//
// Solidity: function perpView() view returns(address)
func (_Periphery *PeripheryCallerSession) PerpView() (common.Address, error) {
	return _Periphery.Contract.PerpView(&_Periphery.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Periphery *PeripheryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Periphery *PeripherySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Periphery.Contract.SupportsInterface(&_Periphery.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Periphery *PeripheryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Periphery.Contract.SupportsInterface(&_Periphery.CallOpts, interfaceId)
}

// BatchOpenPosition is a paid mutator transaction binding the contract method 0x8cce9776.
//
// Solidity: function batchOpenPosition(uint256 _index, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] _ordersLeft, bytes[] _signaturesLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] _ordersRight, bytes[] _signaturesRight, bytes liquidator) returns()
func (_Periphery *PeripheryTransactor) BatchOpenPosition(opts *bind.TransactOpts, _index *big.Int, _ordersLeft []LibOrderOrder, _signaturesLeft [][]byte, _ordersRight []LibOrderOrder, _signaturesRight [][]byte, liquidator []byte) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "batchOpenPosition", _index, _ordersLeft, _signaturesLeft, _ordersRight, _signaturesRight, liquidator)
}

// BatchOpenPosition is a paid mutator transaction binding the contract method 0x8cce9776.
//
// Solidity: function batchOpenPosition(uint256 _index, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] _ordersLeft, bytes[] _signaturesLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] _ordersRight, bytes[] _signaturesRight, bytes liquidator) returns()
func (_Periphery *PeripherySession) BatchOpenPosition(_index *big.Int, _ordersLeft []LibOrderOrder, _signaturesLeft [][]byte, _ordersRight []LibOrderOrder, _signaturesRight [][]byte, liquidator []byte) (*types.Transaction, error) {
	return _Periphery.Contract.BatchOpenPosition(&_Periphery.TransactOpts, _index, _ordersLeft, _signaturesLeft, _ordersRight, _signaturesRight, liquidator)
}

// BatchOpenPosition is a paid mutator transaction binding the contract method 0x8cce9776.
//
// Solidity: function batchOpenPosition(uint256 _index, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] _ordersLeft, bytes[] _signaturesLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool)[] _ordersRight, bytes[] _signaturesRight, bytes liquidator) returns()
func (_Periphery *PeripheryTransactorSession) BatchOpenPosition(_index *big.Int, _ordersLeft []LibOrderOrder, _signaturesLeft [][]byte, _ordersRight []LibOrderOrder, _signaturesRight [][]byte, liquidator []byte) (*types.Transaction, error) {
	return _Periphery.Contract.BatchOpenPosition(&_Periphery.TransactOpts, _index, _ordersLeft, _signaturesLeft, _ordersRight, _signaturesRight, liquidator)
}

// DepositToVault is a paid mutator transaction binding the contract method 0x7457d619.
//
// Solidity: function depositToVault(uint256 _index, address _token, uint256 _amount) payable returns()
func (_Periphery *PeripheryTransactor) DepositToVault(opts *bind.TransactOpts, _index *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "depositToVault", _index, _token, _amount)
}

// DepositToVault is a paid mutator transaction binding the contract method 0x7457d619.
//
// Solidity: function depositToVault(uint256 _index, address _token, uint256 _amount) payable returns()
func (_Periphery *PeripherySession) DepositToVault(_index *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.Contract.DepositToVault(&_Periphery.TransactOpts, _index, _token, _amount)
}

// DepositToVault is a paid mutator transaction binding the contract method 0x7457d619.
//
// Solidity: function depositToVault(uint256 _index, address _token, uint256 _amount) payable returns()
func (_Periphery *PeripheryTransactorSession) DepositToVault(_index *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.Contract.DepositToVault(&_Periphery.TransactOpts, _index, _token, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Periphery *PeripheryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Periphery *PeripherySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.GrantRole(&_Periphery.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Periphery *PeripheryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.GrantRole(&_Periphery.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x11426333.
//
// Solidity: function initialize(address _perpView, address _markPriceOracle, address _indexPriceOracle, address[2] _vaults, address _owner, address _relayer) returns()
func (_Periphery *PeripheryTransactor) Initialize(opts *bind.TransactOpts, _perpView common.Address, _markPriceOracle common.Address, _indexPriceOracle common.Address, _vaults [2]common.Address, _owner common.Address, _relayer common.Address) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "initialize", _perpView, _markPriceOracle, _indexPriceOracle, _vaults, _owner, _relayer)
}

// Initialize is a paid mutator transaction binding the contract method 0x11426333.
//
// Solidity: function initialize(address _perpView, address _markPriceOracle, address _indexPriceOracle, address[2] _vaults, address _owner, address _relayer) returns()
func (_Periphery *PeripherySession) Initialize(_perpView common.Address, _markPriceOracle common.Address, _indexPriceOracle common.Address, _vaults [2]common.Address, _owner common.Address, _relayer common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.Initialize(&_Periphery.TransactOpts, _perpView, _markPriceOracle, _indexPriceOracle, _vaults, _owner, _relayer)
}

// Initialize is a paid mutator transaction binding the contract method 0x11426333.
//
// Solidity: function initialize(address _perpView, address _markPriceOracle, address _indexPriceOracle, address[2] _vaults, address _owner, address _relayer) returns()
func (_Periphery *PeripheryTransactorSession) Initialize(_perpView common.Address, _markPriceOracle common.Address, _indexPriceOracle common.Address, _vaults [2]common.Address, _owner common.Address, _relayer common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.Initialize(&_Periphery.TransactOpts, _perpView, _markPriceOracle, _indexPriceOracle, _vaults, _owner, _relayer)
}

// OpenPosition is a paid mutator transaction binding the contract method 0xeda94035.
//
// Solidity: function openPosition(uint256 _index, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) _orderLeft, bytes _signatureLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) _orderRight, bytes _signatureRight, bytes liquidator) returns()
func (_Periphery *PeripheryTransactor) OpenPosition(opts *bind.TransactOpts, _index *big.Int, _orderLeft LibOrderOrder, _signatureLeft []byte, _orderRight LibOrderOrder, _signatureRight []byte, liquidator []byte) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "openPosition", _index, _orderLeft, _signatureLeft, _orderRight, _signatureRight, liquidator)
}

// OpenPosition is a paid mutator transaction binding the contract method 0xeda94035.
//
// Solidity: function openPosition(uint256 _index, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) _orderLeft, bytes _signatureLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) _orderRight, bytes _signatureRight, bytes liquidator) returns()
func (_Periphery *PeripherySession) OpenPosition(_index *big.Int, _orderLeft LibOrderOrder, _signatureLeft []byte, _orderRight LibOrderOrder, _signatureRight []byte, liquidator []byte) (*types.Transaction, error) {
	return _Periphery.Contract.OpenPosition(&_Periphery.TransactOpts, _index, _orderLeft, _signatureLeft, _orderRight, _signatureRight, liquidator)
}

// OpenPosition is a paid mutator transaction binding the contract method 0xeda94035.
//
// Solidity: function openPosition(uint256 _index, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) _orderLeft, bytes _signatureLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) _orderRight, bytes _signatureRight, bytes liquidator) returns()
func (_Periphery *PeripheryTransactorSession) OpenPosition(_index *big.Int, _orderLeft LibOrderOrder, _signatureLeft []byte, _orderRight LibOrderOrder, _signatureRight []byte, liquidator []byte) (*types.Transaction, error) {
	return _Periphery.Contract.OpenPosition(&_Periphery.TransactOpts, _index, _orderLeft, _signatureLeft, _orderRight, _signatureRight, liquidator)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Periphery *PeripheryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Periphery *PeripherySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.RenounceRole(&_Periphery.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Periphery *PeripheryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.RenounceRole(&_Periphery.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Periphery *PeripheryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Periphery *PeripherySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.RevokeRole(&_Periphery.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Periphery *PeripheryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.RevokeRole(&_Periphery.TransactOpts, role, account)
}

// SetMarkPriceOracle is a paid mutator transaction binding the contract method 0x38710272.
//
// Solidity: function setMarkPriceOracle(address _markPriceOracle) returns()
func (_Periphery *PeripheryTransactor) SetMarkPriceOracle(opts *bind.TransactOpts, _markPriceOracle common.Address) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "setMarkPriceOracle", _markPriceOracle)
}

// SetMarkPriceOracle is a paid mutator transaction binding the contract method 0x38710272.
//
// Solidity: function setMarkPriceOracle(address _markPriceOracle) returns()
func (_Periphery *PeripherySession) SetMarkPriceOracle(_markPriceOracle common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.SetMarkPriceOracle(&_Periphery.TransactOpts, _markPriceOracle)
}

// SetMarkPriceOracle is a paid mutator transaction binding the contract method 0x38710272.
//
// Solidity: function setMarkPriceOracle(address _markPriceOracle) returns()
func (_Periphery *PeripheryTransactorSession) SetMarkPriceOracle(_markPriceOracle common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.SetMarkPriceOracle(&_Periphery.TransactOpts, _markPriceOracle)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _relayer) returns()
func (_Periphery *PeripheryTransactor) SetRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "setRelayer", _relayer)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _relayer) returns()
func (_Periphery *PeripherySession) SetRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.SetRelayer(&_Periphery.TransactOpts, _relayer)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _relayer) returns()
func (_Periphery *PeripheryTransactorSession) SetRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.SetRelayer(&_Periphery.TransactOpts, _relayer)
}

// ToggleTraderWhitelistEnabled is a paid mutator transaction binding the contract method 0x8afa6e75.
//
// Solidity: function toggleTraderWhitelistEnabled() returns()
func (_Periphery *PeripheryTransactor) ToggleTraderWhitelistEnabled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "toggleTraderWhitelistEnabled")
}

// ToggleTraderWhitelistEnabled is a paid mutator transaction binding the contract method 0x8afa6e75.
//
// Solidity: function toggleTraderWhitelistEnabled() returns()
func (_Periphery *PeripherySession) ToggleTraderWhitelistEnabled() (*types.Transaction, error) {
	return _Periphery.Contract.ToggleTraderWhitelistEnabled(&_Periphery.TransactOpts)
}

// ToggleTraderWhitelistEnabled is a paid mutator transaction binding the contract method 0x8afa6e75.
//
// Solidity: function toggleTraderWhitelistEnabled() returns()
func (_Periphery *PeripheryTransactorSession) ToggleTraderWhitelistEnabled() (*types.Transaction, error) {
	return _Periphery.Contract.ToggleTraderWhitelistEnabled(&_Periphery.TransactOpts)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x03ee2295.
//
// Solidity: function transferToVault(address _token, address _from, uint256 _amount) returns()
func (_Periphery *PeripheryTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "transferToVault", _token, _from, _amount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x03ee2295.
//
// Solidity: function transferToVault(address _token, address _from, uint256 _amount) returns()
func (_Periphery *PeripherySession) TransferToVault(_token common.Address, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.Contract.TransferToVault(&_Periphery.TransactOpts, _token, _from, _amount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x03ee2295.
//
// Solidity: function transferToVault(address _token, address _from, uint256 _amount) returns()
func (_Periphery *PeripheryTransactorSession) TransferToVault(_token common.Address, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.Contract.TransferToVault(&_Periphery.TransactOpts, _token, _from, _amount)
}

// WhitelistTrader is a paid mutator transaction binding the contract method 0xeed1cebd.
//
// Solidity: function whitelistTrader(address _trader, bool _isWhitelist) returns()
func (_Periphery *PeripheryTransactor) WhitelistTrader(opts *bind.TransactOpts, _trader common.Address, _isWhitelist bool) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "whitelistTrader", _trader, _isWhitelist)
}

// WhitelistTrader is a paid mutator transaction binding the contract method 0xeed1cebd.
//
// Solidity: function whitelistTrader(address _trader, bool _isWhitelist) returns()
func (_Periphery *PeripherySession) WhitelistTrader(_trader common.Address, _isWhitelist bool) (*types.Transaction, error) {
	return _Periphery.Contract.WhitelistTrader(&_Periphery.TransactOpts, _trader, _isWhitelist)
}

// WhitelistTrader is a paid mutator transaction binding the contract method 0xeed1cebd.
//
// Solidity: function whitelistTrader(address _trader, bool _isWhitelist) returns()
func (_Periphery *PeripheryTransactorSession) WhitelistTrader(_trader common.Address, _isWhitelist bool) (*types.Transaction, error) {
	return _Periphery.Contract.WhitelistTrader(&_Periphery.TransactOpts, _trader, _isWhitelist)
}

// WhitelistVault is a paid mutator transaction binding the contract method 0xe6e66c68.
//
// Solidity: function whitelistVault(address _vault, bool _isWhitelist) returns()
func (_Periphery *PeripheryTransactor) WhitelistVault(opts *bind.TransactOpts, _vault common.Address, _isWhitelist bool) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "whitelistVault", _vault, _isWhitelist)
}

// WhitelistVault is a paid mutator transaction binding the contract method 0xe6e66c68.
//
// Solidity: function whitelistVault(address _vault, bool _isWhitelist) returns()
func (_Periphery *PeripherySession) WhitelistVault(_vault common.Address, _isWhitelist bool) (*types.Transaction, error) {
	return _Periphery.Contract.WhitelistVault(&_Periphery.TransactOpts, _vault, _isWhitelist)
}

// WhitelistVault is a paid mutator transaction binding the contract method 0xe6e66c68.
//
// Solidity: function whitelistVault(address _vault, bool _isWhitelist) returns()
func (_Periphery *PeripheryTransactorSession) WhitelistVault(_vault common.Address, _isWhitelist bool) (*types.Transaction, error) {
	return _Periphery.Contract.WhitelistVault(&_Periphery.TransactOpts, _vault, _isWhitelist)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0x0426a61d.
//
// Solidity: function withdrawFromVault(uint256 _index, address _token, address _to, uint256 _amount) returns()
func (_Periphery *PeripheryTransactor) WithdrawFromVault(opts *bind.TransactOpts, _index *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "withdrawFromVault", _index, _token, _to, _amount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0x0426a61d.
//
// Solidity: function withdrawFromVault(uint256 _index, address _token, address _to, uint256 _amount) returns()
func (_Periphery *PeripherySession) WithdrawFromVault(_index *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.Contract.WithdrawFromVault(&_Periphery.TransactOpts, _index, _token, _to, _amount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0x0426a61d.
//
// Solidity: function withdrawFromVault(uint256 _index, address _token, address _to, uint256 _amount) returns()
func (_Periphery *PeripheryTransactorSession) WithdrawFromVault(_index *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Periphery.Contract.WithdrawFromVault(&_Periphery.TransactOpts, _index, _token, _to, _amount)
}

// PeripheryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Periphery contract.
type PeripheryInitializedIterator struct {
	Event *PeripheryInitialized // Event containing the contract specifics and raw log

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
func (it *PeripheryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeripheryInitialized)
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
		it.Event = new(PeripheryInitialized)
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
func (it *PeripheryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeripheryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeripheryInitialized represents a Initialized event raised by the Periphery contract.
type PeripheryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Periphery *PeripheryFilterer) FilterInitialized(opts *bind.FilterOpts) (*PeripheryInitializedIterator, error) {

	logs, sub, err := _Periphery.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PeripheryInitializedIterator{contract: _Periphery.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Periphery *PeripheryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PeripheryInitialized) (event.Subscription, error) {

	logs, sub, err := _Periphery.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeripheryInitialized)
				if err := _Periphery.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Periphery *PeripheryFilterer) ParseInitialized(log types.Log) (*PeripheryInitialized, error) {
	event := new(PeripheryInitialized)
	if err := _Periphery.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeripheryRelayerUpdatedIterator is returned from FilterRelayerUpdated and is used to iterate over the raw logs and unpacked data for RelayerUpdated events raised by the Periphery contract.
type PeripheryRelayerUpdatedIterator struct {
	Event *PeripheryRelayerUpdated // Event containing the contract specifics and raw log

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
func (it *PeripheryRelayerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeripheryRelayerUpdated)
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
		it.Event = new(PeripheryRelayerUpdated)
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
func (it *PeripheryRelayerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeripheryRelayerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeripheryRelayerUpdated represents a RelayerUpdated event raised by the Periphery contract.
type PeripheryRelayerUpdated struct {
	NewRelayerAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRelayerUpdated is a free log retrieval operation binding the contract event 0x09bf35b7806e61e7baf86e6547c4221f69d130ccc0a19215d064685a89464a70.
//
// Solidity: event RelayerUpdated(address indexed newRelayerAddress)
func (_Periphery *PeripheryFilterer) FilterRelayerUpdated(opts *bind.FilterOpts, newRelayerAddress []common.Address) (*PeripheryRelayerUpdatedIterator, error) {

	var newRelayerAddressRule []interface{}
	for _, newRelayerAddressItem := range newRelayerAddress {
		newRelayerAddressRule = append(newRelayerAddressRule, newRelayerAddressItem)
	}

	logs, sub, err := _Periphery.contract.FilterLogs(opts, "RelayerUpdated", newRelayerAddressRule)
	if err != nil {
		return nil, err
	}
	return &PeripheryRelayerUpdatedIterator{contract: _Periphery.contract, event: "RelayerUpdated", logs: logs, sub: sub}, nil
}

// WatchRelayerUpdated is a free log subscription operation binding the contract event 0x09bf35b7806e61e7baf86e6547c4221f69d130ccc0a19215d064685a89464a70.
//
// Solidity: event RelayerUpdated(address indexed newRelayerAddress)
func (_Periphery *PeripheryFilterer) WatchRelayerUpdated(opts *bind.WatchOpts, sink chan<- *PeripheryRelayerUpdated, newRelayerAddress []common.Address) (event.Subscription, error) {

	var newRelayerAddressRule []interface{}
	for _, newRelayerAddressItem := range newRelayerAddress {
		newRelayerAddressRule = append(newRelayerAddressRule, newRelayerAddressItem)
	}

	logs, sub, err := _Periphery.contract.WatchLogs(opts, "RelayerUpdated", newRelayerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeripheryRelayerUpdated)
				if err := _Periphery.contract.UnpackLog(event, "RelayerUpdated", log); err != nil {
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

// ParseRelayerUpdated is a log parse operation binding the contract event 0x09bf35b7806e61e7baf86e6547c4221f69d130ccc0a19215d064685a89464a70.
//
// Solidity: event RelayerUpdated(address indexed newRelayerAddress)
func (_Periphery *PeripheryFilterer) ParseRelayerUpdated(log types.Log) (*PeripheryRelayerUpdated, error) {
	event := new(PeripheryRelayerUpdated)
	if err := _Periphery.contract.UnpackLog(event, "RelayerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeripheryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Periphery contract.
type PeripheryRoleAdminChangedIterator struct {
	Event *PeripheryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PeripheryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeripheryRoleAdminChanged)
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
		it.Event = new(PeripheryRoleAdminChanged)
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
func (it *PeripheryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeripheryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeripheryRoleAdminChanged represents a RoleAdminChanged event raised by the Periphery contract.
type PeripheryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Periphery *PeripheryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PeripheryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Periphery.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PeripheryRoleAdminChangedIterator{contract: _Periphery.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Periphery *PeripheryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PeripheryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Periphery.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeripheryRoleAdminChanged)
				if err := _Periphery.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Periphery *PeripheryFilterer) ParseRoleAdminChanged(log types.Log) (*PeripheryRoleAdminChanged, error) {
	event := new(PeripheryRoleAdminChanged)
	if err := _Periphery.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeripheryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Periphery contract.
type PeripheryRoleGrantedIterator struct {
	Event *PeripheryRoleGranted // Event containing the contract specifics and raw log

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
func (it *PeripheryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeripheryRoleGranted)
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
		it.Event = new(PeripheryRoleGranted)
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
func (it *PeripheryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeripheryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeripheryRoleGranted represents a RoleGranted event raised by the Periphery contract.
type PeripheryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Periphery *PeripheryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PeripheryRoleGrantedIterator, error) {

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

	logs, sub, err := _Periphery.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PeripheryRoleGrantedIterator{contract: _Periphery.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Periphery *PeripheryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PeripheryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Periphery.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeripheryRoleGranted)
				if err := _Periphery.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Periphery *PeripheryFilterer) ParseRoleGranted(log types.Log) (*PeripheryRoleGranted, error) {
	event := new(PeripheryRoleGranted)
	if err := _Periphery.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeripheryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Periphery contract.
type PeripheryRoleRevokedIterator struct {
	Event *PeripheryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PeripheryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeripheryRoleRevoked)
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
		it.Event = new(PeripheryRoleRevoked)
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
func (it *PeripheryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeripheryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeripheryRoleRevoked represents a RoleRevoked event raised by the Periphery contract.
type PeripheryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Periphery *PeripheryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PeripheryRoleRevokedIterator, error) {

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

	logs, sub, err := _Periphery.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PeripheryRoleRevokedIterator{contract: _Periphery.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Periphery *PeripheryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PeripheryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Periphery.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeripheryRoleRevoked)
				if err := _Periphery.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Periphery *PeripheryFilterer) ParseRoleRevoked(log types.Log) (*PeripheryRoleRevoked, error) {
	event := new(PeripheryRoleRevoked)
	if err := _Periphery.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeripheryTraderWhitelistedIterator is returned from FilterTraderWhitelisted and is used to iterate over the raw logs and unpacked data for TraderWhitelisted events raised by the Periphery contract.
type PeripheryTraderWhitelistedIterator struct {
	Event *PeripheryTraderWhitelisted // Event containing the contract specifics and raw log

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
func (it *PeripheryTraderWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeripheryTraderWhitelisted)
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
		it.Event = new(PeripheryTraderWhitelisted)
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
func (it *PeripheryTraderWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeripheryTraderWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeripheryTraderWhitelisted represents a TraderWhitelisted event raised by the Periphery contract.
type PeripheryTraderWhitelisted struct {
	Account     common.Address
	IsWhitelist bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTraderWhitelisted is a free log retrieval operation binding the contract event 0x08f7f1f75f58cc1c49f314a967261548d102e1ce5415459a17e184cf501bd21a.
//
// Solidity: event TraderWhitelisted(address indexed account, bool isWhitelist)
func (_Periphery *PeripheryFilterer) FilterTraderWhitelisted(opts *bind.FilterOpts, account []common.Address) (*PeripheryTraderWhitelistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Periphery.contract.FilterLogs(opts, "TraderWhitelisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &PeripheryTraderWhitelistedIterator{contract: _Periphery.contract, event: "TraderWhitelisted", logs: logs, sub: sub}, nil
}

// WatchTraderWhitelisted is a free log subscription operation binding the contract event 0x08f7f1f75f58cc1c49f314a967261548d102e1ce5415459a17e184cf501bd21a.
//
// Solidity: event TraderWhitelisted(address indexed account, bool isWhitelist)
func (_Periphery *PeripheryFilterer) WatchTraderWhitelisted(opts *bind.WatchOpts, sink chan<- *PeripheryTraderWhitelisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Periphery.contract.WatchLogs(opts, "TraderWhitelisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeripheryTraderWhitelisted)
				if err := _Periphery.contract.UnpackLog(event, "TraderWhitelisted", log); err != nil {
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

// ParseTraderWhitelisted is a log parse operation binding the contract event 0x08f7f1f75f58cc1c49f314a967261548d102e1ce5415459a17e184cf501bd21a.
//
// Solidity: event TraderWhitelisted(address indexed account, bool isWhitelist)
func (_Periphery *PeripheryFilterer) ParseTraderWhitelisted(log types.Log) (*PeripheryTraderWhitelisted, error) {
	event := new(PeripheryTraderWhitelisted)
	if err := _Periphery.contract.UnpackLog(event, "TraderWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeripheryVaultWhitelistedIterator is returned from FilterVaultWhitelisted and is used to iterate over the raw logs and unpacked data for VaultWhitelisted events raised by the Periphery contract.
type PeripheryVaultWhitelistedIterator struct {
	Event *PeripheryVaultWhitelisted // Event containing the contract specifics and raw log

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
func (it *PeripheryVaultWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeripheryVaultWhitelisted)
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
		it.Event = new(PeripheryVaultWhitelisted)
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
func (it *PeripheryVaultWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeripheryVaultWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeripheryVaultWhitelisted represents a VaultWhitelisted event raised by the Periphery contract.
type PeripheryVaultWhitelisted struct {
	Vault       common.Address
	IsWhitelist bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVaultWhitelisted is a free log retrieval operation binding the contract event 0x9652e0d255e1984241157f0e5f65c662604b493e181fde308bc5b66d116d63fe.
//
// Solidity: event VaultWhitelisted(address indexed vault, bool isWhitelist)
func (_Periphery *PeripheryFilterer) FilterVaultWhitelisted(opts *bind.FilterOpts, vault []common.Address) (*PeripheryVaultWhitelistedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Periphery.contract.FilterLogs(opts, "VaultWhitelisted", vaultRule)
	if err != nil {
		return nil, err
	}
	return &PeripheryVaultWhitelistedIterator{contract: _Periphery.contract, event: "VaultWhitelisted", logs: logs, sub: sub}, nil
}

// WatchVaultWhitelisted is a free log subscription operation binding the contract event 0x9652e0d255e1984241157f0e5f65c662604b493e181fde308bc5b66d116d63fe.
//
// Solidity: event VaultWhitelisted(address indexed vault, bool isWhitelist)
func (_Periphery *PeripheryFilterer) WatchVaultWhitelisted(opts *bind.WatchOpts, sink chan<- *PeripheryVaultWhitelisted, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Periphery.contract.WatchLogs(opts, "VaultWhitelisted", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeripheryVaultWhitelisted)
				if err := _Periphery.contract.UnpackLog(event, "VaultWhitelisted", log); err != nil {
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

// ParseVaultWhitelisted is a log parse operation binding the contract event 0x9652e0d255e1984241157f0e5f65c662604b493e181fde308bc5b66d116d63fe.
//
// Solidity: event VaultWhitelisted(address indexed vault, bool isWhitelist)
func (_Periphery *PeripheryFilterer) ParseVaultWhitelisted(log types.Log) (*PeripheryVaultWhitelisted, error) {
	event := new(PeripheryVaultWhitelisted)
	if err := _Periphery.contract.UnpackLog(event, "VaultWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
