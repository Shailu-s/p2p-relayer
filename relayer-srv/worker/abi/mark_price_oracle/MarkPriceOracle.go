// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MarkPriceOracle

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

// MarkPriceOracleMetaData contains all meta data concerning the MarkPriceOracle contract.
var MarkPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"underlyingPrices\",\"type\":\"uint256[]\"}],\"name\":\"AssetsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"markPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ObservationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"matchingEngine\",\"type\":\"address\"}],\"name\":\"ObservationAdderSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADD_OBSERVATION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_ORACLE_ADMIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TWAP_INTERVAL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_underlyingPrice\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_asset\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofHash\",\"type\":\"bytes32[]\"}],\"name\":\"addAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_underlyingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_proofHash\",\"type\":\"bytes32\"}],\"name\":\"addObservation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"baseTokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCustomMarkTwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceCumulative\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCustomUnderlyingTwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceCumulative\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIndexCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingLastPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLastUpdatedTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastUpdatedTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_twInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMarkTwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceCumulative\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_positioningConfig\",\"type\":\"address\"}],\"name\":\"grantTwapIntervalRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexByBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexOracle\",\"outputs\":[{\"internalType\":\"contractIIndexPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexTwInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_priceCumulative\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_asset\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"markTwInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"observationsByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"markPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positioning\",\"outputs\":[{\"internalType\":\"contractIPositioning\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIndexPriceOracle\",\"name\":\"_indexOracle\",\"type\":\"address\"}],\"name\":\"setIndexOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_indexTwInterval\",\"type\":\"uint256\"}],\"name\":\"setIndexTwInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_markTwInterval\",\"type\":\"uint256\"}],\"name\":\"setMarkTwInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adder\",\"type\":\"address\"}],\"name\":\"setObservationAdder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPositioning\",\"name\":\"_positioning\",\"type\":\"address\"}],\"name\":\"setPositioning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MarkPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use MarkPriceOracleMetaData.ABI instead.
var MarkPriceOracleABI = MarkPriceOracleMetaData.ABI

// MarkPriceOracle is an auto generated Go binding around an Ethereum contract.
type MarkPriceOracle struct {
	MarkPriceOracleCaller     // Read-only binding to the contract
	MarkPriceOracleTransactor // Write-only binding to the contract
	MarkPriceOracleFilterer   // Log filterer for contract events
}

// MarkPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarkPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarkPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarkPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarkPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarkPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarkPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarkPriceOracleSession struct {
	Contract     *MarkPriceOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarkPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarkPriceOracleCallerSession struct {
	Contract *MarkPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MarkPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarkPriceOracleTransactorSession struct {
	Contract     *MarkPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MarkPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarkPriceOracleRaw struct {
	Contract *MarkPriceOracle // Generic contract binding to access the raw methods on
}

// MarkPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarkPriceOracleCallerRaw struct {
	Contract *MarkPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// MarkPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarkPriceOracleTransactorRaw struct {
	Contract *MarkPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarkPriceOracle creates a new instance of MarkPriceOracle, bound to a specific deployed contract.
func NewMarkPriceOracle(address common.Address, backend bind.ContractBackend) (*MarkPriceOracle, error) {
	contract, err := bindMarkPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracle{MarkPriceOracleCaller: MarkPriceOracleCaller{contract: contract}, MarkPriceOracleTransactor: MarkPriceOracleTransactor{contract: contract}, MarkPriceOracleFilterer: MarkPriceOracleFilterer{contract: contract}}, nil
}

// NewMarkPriceOracleCaller creates a new read-only instance of MarkPriceOracle, bound to a specific deployed contract.
func NewMarkPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*MarkPriceOracleCaller, error) {
	contract, err := bindMarkPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleCaller{contract: contract}, nil
}

// NewMarkPriceOracleTransactor creates a new write-only instance of MarkPriceOracle, bound to a specific deployed contract.
func NewMarkPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*MarkPriceOracleTransactor, error) {
	contract, err := bindMarkPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleTransactor{contract: contract}, nil
}

// NewMarkPriceOracleFilterer creates a new log filterer instance of MarkPriceOracle, bound to a specific deployed contract.
func NewMarkPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*MarkPriceOracleFilterer, error) {
	contract, err := bindMarkPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleFilterer{contract: contract}, nil
}

// bindMarkPriceOracle binds a generic wrapper to an already deployed contract.
func bindMarkPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarkPriceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarkPriceOracle *MarkPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarkPriceOracle.Contract.MarkPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarkPriceOracle *MarkPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.MarkPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarkPriceOracle *MarkPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.MarkPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarkPriceOracle *MarkPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarkPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarkPriceOracle *MarkPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarkPriceOracle *MarkPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// ADDOBSERVATIONROLE is a free data retrieval call binding the contract method 0x0db0f0f5.
//
// Solidity: function ADD_OBSERVATION_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCaller) ADDOBSERVATIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "ADD_OBSERVATION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDOBSERVATIONROLE is a free data retrieval call binding the contract method 0x0db0f0f5.
//
// Solidity: function ADD_OBSERVATION_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleSession) ADDOBSERVATIONROLE() ([32]byte, error) {
	return _MarkPriceOracle.Contract.ADDOBSERVATIONROLE(&_MarkPriceOracle.CallOpts)
}

// ADDOBSERVATIONROLE is a free data retrieval call binding the contract method 0x0db0f0f5.
//
// Solidity: function ADD_OBSERVATION_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) ADDOBSERVATIONROLE() ([32]byte, error) {
	return _MarkPriceOracle.Contract.ADDOBSERVATIONROLE(&_MarkPriceOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MarkPriceOracle.Contract.DEFAULTADMINROLE(&_MarkPriceOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MarkPriceOracle.Contract.DEFAULTADMINROLE(&_MarkPriceOracle.CallOpts)
}

// PRICEORACLEADMIN is a free data retrieval call binding the contract method 0xa1f68193.
//
// Solidity: function PRICE_ORACLE_ADMIN() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCaller) PRICEORACLEADMIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "PRICE_ORACLE_ADMIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRICEORACLEADMIN is a free data retrieval call binding the contract method 0xa1f68193.
//
// Solidity: function PRICE_ORACLE_ADMIN() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleSession) PRICEORACLEADMIN() ([32]byte, error) {
	return _MarkPriceOracle.Contract.PRICEORACLEADMIN(&_MarkPriceOracle.CallOpts)
}

// PRICEORACLEADMIN is a free data retrieval call binding the contract method 0xa1f68193.
//
// Solidity: function PRICE_ORACLE_ADMIN() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) PRICEORACLEADMIN() ([32]byte, error) {
	return _MarkPriceOracle.Contract.PRICEORACLEADMIN(&_MarkPriceOracle.CallOpts)
}

// TWAPINTERVALROLE is a free data retrieval call binding the contract method 0x026828aa.
//
// Solidity: function TWAP_INTERVAL_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCaller) TWAPINTERVALROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "TWAP_INTERVAL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TWAPINTERVALROLE is a free data retrieval call binding the contract method 0x026828aa.
//
// Solidity: function TWAP_INTERVAL_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleSession) TWAPINTERVALROLE() ([32]byte, error) {
	return _MarkPriceOracle.Contract.TWAPINTERVALROLE(&_MarkPriceOracle.CallOpts)
}

// TWAPINTERVALROLE is a free data retrieval call binding the contract method 0x026828aa.
//
// Solidity: function TWAP_INTERVAL_ROLE() view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) TWAPINTERVALROLE() ([32]byte, error) {
	return _MarkPriceOracle.Contract.TWAPINTERVALROLE(&_MarkPriceOracle.CallOpts)
}

// BaseTokenByIndex is a free data retrieval call binding the contract method 0x599a945c.
//
// Solidity: function baseTokenByIndex(uint256 ) view returns(address)
func (_MarkPriceOracle *MarkPriceOracleCaller) BaseTokenByIndex(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "baseTokenByIndex", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseTokenByIndex is a free data retrieval call binding the contract method 0x599a945c.
//
// Solidity: function baseTokenByIndex(uint256 ) view returns(address)
func (_MarkPriceOracle *MarkPriceOracleSession) BaseTokenByIndex(arg0 *big.Int) (common.Address, error) {
	return _MarkPriceOracle.Contract.BaseTokenByIndex(&_MarkPriceOracle.CallOpts, arg0)
}

// BaseTokenByIndex is a free data retrieval call binding the contract method 0x599a945c.
//
// Solidity: function baseTokenByIndex(uint256 ) view returns(address)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) BaseTokenByIndex(arg0 *big.Int) (common.Address, error) {
	return _MarkPriceOracle.Contract.BaseTokenByIndex(&_MarkPriceOracle.CallOpts, arg0)
}

// GetCustomMarkTwap is a free data retrieval call binding the contract method 0x7d6f94a5.
//
// Solidity: function getCustomMarkTwap(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleCaller) GetCustomMarkTwap(opts *bind.CallOpts, _index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "getCustomMarkTwap", _index, _startTimestamp, _endTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCustomMarkTwap is a free data retrieval call binding the contract method 0x7d6f94a5.
//
// Solidity: function getCustomMarkTwap(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleSession) GetCustomMarkTwap(_index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetCustomMarkTwap(&_MarkPriceOracle.CallOpts, _index, _startTimestamp, _endTimestamp)
}

// GetCustomMarkTwap is a free data retrieval call binding the contract method 0x7d6f94a5.
//
// Solidity: function getCustomMarkTwap(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) GetCustomMarkTwap(_index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetCustomMarkTwap(&_MarkPriceOracle.CallOpts, _index, _startTimestamp, _endTimestamp)
}

// GetCustomUnderlyingTwap is a free data retrieval call binding the contract method 0xb8d423a5.
//
// Solidity: function getCustomUnderlyingTwap(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleCaller) GetCustomUnderlyingTwap(opts *bind.CallOpts, _index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "getCustomUnderlyingTwap", _index, _startTimestamp, _endTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCustomUnderlyingTwap is a free data retrieval call binding the contract method 0xb8d423a5.
//
// Solidity: function getCustomUnderlyingTwap(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleSession) GetCustomUnderlyingTwap(_index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetCustomUnderlyingTwap(&_MarkPriceOracle.CallOpts, _index, _startTimestamp, _endTimestamp)
}

// GetCustomUnderlyingTwap is a free data retrieval call binding the contract method 0xb8d423a5.
//
// Solidity: function getCustomUnderlyingTwap(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) GetCustomUnderlyingTwap(_index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetCustomUnderlyingTwap(&_MarkPriceOracle.CallOpts, _index, _startTimestamp, _endTimestamp)
}

// GetIndexCount is a free data retrieval call binding the contract method 0xec244370.
//
// Solidity: function getIndexCount() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleCaller) GetIndexCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "getIndexCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexCount is a free data retrieval call binding the contract method 0xec244370.
//
// Solidity: function getIndexCount() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleSession) GetIndexCount() (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetIndexCount(&_MarkPriceOracle.CallOpts)
}

// GetIndexCount is a free data retrieval call binding the contract method 0xec244370.
//
// Solidity: function getIndexCount() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) GetIndexCount() (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetIndexCount(&_MarkPriceOracle.CallOpts)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x65fa2f7f.
//
// Solidity: function getLastPrice(uint256 _index) view returns(uint256 underlyingLastPrice)
func (_MarkPriceOracle *MarkPriceOracleCaller) GetLastPrice(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "getLastPrice", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPrice is a free data retrieval call binding the contract method 0x65fa2f7f.
//
// Solidity: function getLastPrice(uint256 _index) view returns(uint256 underlyingLastPrice)
func (_MarkPriceOracle *MarkPriceOracleSession) GetLastPrice(_index *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetLastPrice(&_MarkPriceOracle.CallOpts, _index)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x65fa2f7f.
//
// Solidity: function getLastPrice(uint256 _index) view returns(uint256 underlyingLastPrice)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) GetLastPrice(_index *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetLastPrice(&_MarkPriceOracle.CallOpts, _index)
}

// GetLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xa367c24d.
//
// Solidity: function getLastUpdatedTimestamp(uint256 _index) view returns(uint256 lastUpdatedTimestamp)
func (_MarkPriceOracle *MarkPriceOracleCaller) GetLastUpdatedTimestamp(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "getLastUpdatedTimestamp", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xa367c24d.
//
// Solidity: function getLastUpdatedTimestamp(uint256 _index) view returns(uint256 lastUpdatedTimestamp)
func (_MarkPriceOracle *MarkPriceOracleSession) GetLastUpdatedTimestamp(_index *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetLastUpdatedTimestamp(&_MarkPriceOracle.CallOpts, _index)
}

// GetLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xa367c24d.
//
// Solidity: function getLastUpdatedTimestamp(uint256 _index) view returns(uint256 lastUpdatedTimestamp)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) GetLastUpdatedTimestamp(_index *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetLastUpdatedTimestamp(&_MarkPriceOracle.CallOpts, _index)
}

// GetMarkTwap is a free data retrieval call binding the contract method 0xf6f94975.
//
// Solidity: function getMarkTwap(uint256 _twInterval, uint256 _index) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleCaller) GetMarkTwap(opts *bind.CallOpts, _twInterval *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "getMarkTwap", _twInterval, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarkTwap is a free data retrieval call binding the contract method 0xf6f94975.
//
// Solidity: function getMarkTwap(uint256 _twInterval, uint256 _index) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleSession) GetMarkTwap(_twInterval *big.Int, _index *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetMarkTwap(&_MarkPriceOracle.CallOpts, _twInterval, _index)
}

// GetMarkTwap is a free data retrieval call binding the contract method 0xf6f94975.
//
// Solidity: function getMarkTwap(uint256 _twInterval, uint256 _index) view returns(uint256 priceCumulative)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) GetMarkTwap(_twInterval *big.Int, _index *big.Int) (*big.Int, error) {
	return _MarkPriceOracle.Contract.GetMarkTwap(&_MarkPriceOracle.CallOpts, _twInterval, _index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MarkPriceOracle.Contract.GetRoleAdmin(&_MarkPriceOracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MarkPriceOracle.Contract.GetRoleAdmin(&_MarkPriceOracle.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MarkPriceOracle *MarkPriceOracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MarkPriceOracle *MarkPriceOracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MarkPriceOracle.Contract.HasRole(&_MarkPriceOracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MarkPriceOracle.Contract.HasRole(&_MarkPriceOracle.CallOpts, role, account)
}

// IndexByBaseToken is a free data retrieval call binding the contract method 0xac416c17.
//
// Solidity: function indexByBaseToken(address ) view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleCaller) IndexByBaseToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "indexByBaseToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexByBaseToken is a free data retrieval call binding the contract method 0xac416c17.
//
// Solidity: function indexByBaseToken(address ) view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleSession) IndexByBaseToken(arg0 common.Address) (*big.Int, error) {
	return _MarkPriceOracle.Contract.IndexByBaseToken(&_MarkPriceOracle.CallOpts, arg0)
}

// IndexByBaseToken is a free data retrieval call binding the contract method 0xac416c17.
//
// Solidity: function indexByBaseToken(address ) view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) IndexByBaseToken(arg0 common.Address) (*big.Int, error) {
	return _MarkPriceOracle.Contract.IndexByBaseToken(&_MarkPriceOracle.CallOpts, arg0)
}

// IndexOracle is a free data retrieval call binding the contract method 0x2e5e0b18.
//
// Solidity: function indexOracle() view returns(address)
func (_MarkPriceOracle *MarkPriceOracleCaller) IndexOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "indexOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexOracle is a free data retrieval call binding the contract method 0x2e5e0b18.
//
// Solidity: function indexOracle() view returns(address)
func (_MarkPriceOracle *MarkPriceOracleSession) IndexOracle() (common.Address, error) {
	return _MarkPriceOracle.Contract.IndexOracle(&_MarkPriceOracle.CallOpts)
}

// IndexOracle is a free data retrieval call binding the contract method 0x2e5e0b18.
//
// Solidity: function indexOracle() view returns(address)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) IndexOracle() (common.Address, error) {
	return _MarkPriceOracle.Contract.IndexOracle(&_MarkPriceOracle.CallOpts)
}

// IndexTwInterval is a free data retrieval call binding the contract method 0x872d775c.
//
// Solidity: function indexTwInterval() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleCaller) IndexTwInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "indexTwInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexTwInterval is a free data retrieval call binding the contract method 0x872d775c.
//
// Solidity: function indexTwInterval() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleSession) IndexTwInterval() (*big.Int, error) {
	return _MarkPriceOracle.Contract.IndexTwInterval(&_MarkPriceOracle.CallOpts)
}

// IndexTwInterval is a free data retrieval call binding the contract method 0x872d775c.
//
// Solidity: function indexTwInterval() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) IndexTwInterval() (*big.Int, error) {
	return _MarkPriceOracle.Contract.IndexTwInterval(&_MarkPriceOracle.CallOpts)
}

// MarkTwInterval is a free data retrieval call binding the contract method 0x26759a97.
//
// Solidity: function markTwInterval() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleCaller) MarkTwInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "markTwInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarkTwInterval is a free data retrieval call binding the contract method 0x26759a97.
//
// Solidity: function markTwInterval() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleSession) MarkTwInterval() (*big.Int, error) {
	return _MarkPriceOracle.Contract.MarkTwInterval(&_MarkPriceOracle.CallOpts)
}

// MarkTwInterval is a free data retrieval call binding the contract method 0x26759a97.
//
// Solidity: function markTwInterval() view returns(uint256)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) MarkTwInterval() (*big.Int, error) {
	return _MarkPriceOracle.Contract.MarkTwInterval(&_MarkPriceOracle.CallOpts)
}

// ObservationsByIndex is a free data retrieval call binding the contract method 0x80319d88.
//
// Solidity: function observationsByIndex(uint256 , uint256 ) view returns(uint256 timestamp, uint256 underlyingPrice, bytes32 proofHash, uint256 markPrice)
func (_MarkPriceOracle *MarkPriceOracleCaller) ObservationsByIndex(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp       *big.Int
	UnderlyingPrice *big.Int
	ProofHash       [32]byte
	MarkPrice       *big.Int
}, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "observationsByIndex", arg0, arg1)

	outstruct := new(struct {
		Timestamp       *big.Int
		UnderlyingPrice *big.Int
		ProofHash       [32]byte
		MarkPrice       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnderlyingPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProofHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.MarkPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ObservationsByIndex is a free data retrieval call binding the contract method 0x80319d88.
//
// Solidity: function observationsByIndex(uint256 , uint256 ) view returns(uint256 timestamp, uint256 underlyingPrice, bytes32 proofHash, uint256 markPrice)
func (_MarkPriceOracle *MarkPriceOracleSession) ObservationsByIndex(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp       *big.Int
	UnderlyingPrice *big.Int
	ProofHash       [32]byte
	MarkPrice       *big.Int
}, error) {
	return _MarkPriceOracle.Contract.ObservationsByIndex(&_MarkPriceOracle.CallOpts, arg0, arg1)
}

// ObservationsByIndex is a free data retrieval call binding the contract method 0x80319d88.
//
// Solidity: function observationsByIndex(uint256 , uint256 ) view returns(uint256 timestamp, uint256 underlyingPrice, bytes32 proofHash, uint256 markPrice)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) ObservationsByIndex(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp       *big.Int
	UnderlyingPrice *big.Int
	ProofHash       [32]byte
	MarkPrice       *big.Int
}, error) {
	return _MarkPriceOracle.Contract.ObservationsByIndex(&_MarkPriceOracle.CallOpts, arg0, arg1)
}

// Positioning is a free data retrieval call binding the contract method 0xee1d32f8.
//
// Solidity: function positioning() view returns(address)
func (_MarkPriceOracle *MarkPriceOracleCaller) Positioning(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "positioning")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Positioning is a free data retrieval call binding the contract method 0xee1d32f8.
//
// Solidity: function positioning() view returns(address)
func (_MarkPriceOracle *MarkPriceOracleSession) Positioning() (common.Address, error) {
	return _MarkPriceOracle.Contract.Positioning(&_MarkPriceOracle.CallOpts)
}

// Positioning is a free data retrieval call binding the contract method 0xee1d32f8.
//
// Solidity: function positioning() view returns(address)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) Positioning() (common.Address, error) {
	return _MarkPriceOracle.Contract.Positioning(&_MarkPriceOracle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MarkPriceOracle *MarkPriceOracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MarkPriceOracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MarkPriceOracle *MarkPriceOracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MarkPriceOracle.Contract.SupportsInterface(&_MarkPriceOracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MarkPriceOracle *MarkPriceOracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MarkPriceOracle.Contract.SupportsInterface(&_MarkPriceOracle.CallOpts, interfaceId)
}

// AddAssets is a paid mutator transaction binding the contract method 0xb4fd50d6.
//
// Solidity: function addAssets(uint256[] _underlyingPrice, address[] _asset, bytes32[] _proofHash) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) AddAssets(opts *bind.TransactOpts, _underlyingPrice []*big.Int, _asset []common.Address, _proofHash [][32]byte) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "addAssets", _underlyingPrice, _asset, _proofHash)
}

// AddAssets is a paid mutator transaction binding the contract method 0xb4fd50d6.
//
// Solidity: function addAssets(uint256[] _underlyingPrice, address[] _asset, bytes32[] _proofHash) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) AddAssets(_underlyingPrice []*big.Int, _asset []common.Address, _proofHash [][32]byte) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.AddAssets(&_MarkPriceOracle.TransactOpts, _underlyingPrice, _asset, _proofHash)
}

// AddAssets is a paid mutator transaction binding the contract method 0xb4fd50d6.
//
// Solidity: function addAssets(uint256[] _underlyingPrice, address[] _asset, bytes32[] _proofHash) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) AddAssets(_underlyingPrice []*big.Int, _asset []common.Address, _proofHash [][32]byte) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.AddAssets(&_MarkPriceOracle.TransactOpts, _underlyingPrice, _asset, _proofHash)
}

// AddObservation is a paid mutator transaction binding the contract method 0xa8c81560.
//
// Solidity: function addObservation(uint256 _underlyingPrice, uint256 _index, bytes32 _proofHash) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) AddObservation(opts *bind.TransactOpts, _underlyingPrice *big.Int, _index *big.Int, _proofHash [32]byte) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "addObservation", _underlyingPrice, _index, _proofHash)
}

// AddObservation is a paid mutator transaction binding the contract method 0xa8c81560.
//
// Solidity: function addObservation(uint256 _underlyingPrice, uint256 _index, bytes32 _proofHash) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) AddObservation(_underlyingPrice *big.Int, _index *big.Int, _proofHash [32]byte) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.AddObservation(&_MarkPriceOracle.TransactOpts, _underlyingPrice, _index, _proofHash)
}

// AddObservation is a paid mutator transaction binding the contract method 0xa8c81560.
//
// Solidity: function addObservation(uint256 _underlyingPrice, uint256 _index, bytes32 _proofHash) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) AddObservation(_underlyingPrice *big.Int, _index *big.Int, _proofHash [32]byte) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.AddObservation(&_MarkPriceOracle.TransactOpts, _underlyingPrice, _index, _proofHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.GrantRole(&_MarkPriceOracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.GrantRole(&_MarkPriceOracle.TransactOpts, role, account)
}

// GrantTwapIntervalRole is a paid mutator transaction binding the contract method 0x84c7a620.
//
// Solidity: function grantTwapIntervalRole(address _positioningConfig) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) GrantTwapIntervalRole(opts *bind.TransactOpts, _positioningConfig common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "grantTwapIntervalRole", _positioningConfig)
}

// GrantTwapIntervalRole is a paid mutator transaction binding the contract method 0x84c7a620.
//
// Solidity: function grantTwapIntervalRole(address _positioningConfig) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) GrantTwapIntervalRole(_positioningConfig common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.GrantTwapIntervalRole(&_MarkPriceOracle.TransactOpts, _positioningConfig)
}

// GrantTwapIntervalRole is a paid mutator transaction binding the contract method 0x84c7a620.
//
// Solidity: function grantTwapIntervalRole(address _positioningConfig) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) GrantTwapIntervalRole(_positioningConfig common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.GrantTwapIntervalRole(&_MarkPriceOracle.TransactOpts, _positioningConfig)
}


// Initialize is a paid mutator transaction binding the contract method 0x86e553cf.
//
// Solidity: function initialize(uint256[] _priceCumulative, address[] _asset, bytes32[] _proofHash, address _admin) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) Initialize(opts *bind.TransactOpts, _priceCumulative []*big.Int, _asset []common.Address, _proofHash [][32]byte, _admin common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "initialize", _priceCumulative, _asset, _proofHash, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x86e553cf.
//
// Solidity: function initialize(uint256[] _priceCumulative, address[] _asset, bytes32[] _proofHash, address _admin) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) Initialize(_priceCumulative []*big.Int, _asset []common.Address, _proofHash [][32]byte, _admin common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.Initialize(&_MarkPriceOracle.TransactOpts, _priceCumulative, _asset, _proofHash, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x86e553cf.
//
// Solidity: function initialize(uint256[] _priceCumulative, address[] _asset, bytes32[] _proofHash, address _admin) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) Initialize(_priceCumulative []*big.Int, _asset []common.Address, _proofHash [][32]byte, _admin common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.Initialize(&_MarkPriceOracle.TransactOpts, _priceCumulative, _asset, _proofHash, _admin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.RenounceRole(&_MarkPriceOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.RenounceRole(&_MarkPriceOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.RevokeRole(&_MarkPriceOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.RevokeRole(&_MarkPriceOracle.TransactOpts, role, account)
}

// SetIndexOracle is a paid mutator transaction binding the contract method 0x102a46a9.
//
// Solidity: function setIndexOracle(address _indexOracle) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) SetIndexOracle(opts *bind.TransactOpts, _indexOracle common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "setIndexOracle", _indexOracle)
}

// SetIndexOracle is a paid mutator transaction binding the contract method 0x102a46a9.
//
// Solidity: function setIndexOracle(address _indexOracle) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) SetIndexOracle(_indexOracle common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetIndexOracle(&_MarkPriceOracle.TransactOpts, _indexOracle)
}

// SetIndexOracle is a paid mutator transaction binding the contract method 0x102a46a9.
//
// Solidity: function setIndexOracle(address _indexOracle) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) SetIndexOracle(_indexOracle common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetIndexOracle(&_MarkPriceOracle.TransactOpts, _indexOracle)
}

// SetIndexTwInterval is a paid mutator transaction binding the contract method 0x66a25830.
//
// Solidity: function setIndexTwInterval(uint256 _indexTwInterval) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) SetIndexTwInterval(opts *bind.TransactOpts, _indexTwInterval *big.Int) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "setIndexTwInterval", _indexTwInterval)
}

// SetIndexTwInterval is a paid mutator transaction binding the contract method 0x66a25830.
//
// Solidity: function setIndexTwInterval(uint256 _indexTwInterval) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) SetIndexTwInterval(_indexTwInterval *big.Int) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetIndexTwInterval(&_MarkPriceOracle.TransactOpts, _indexTwInterval)
}

// SetIndexTwInterval is a paid mutator transaction binding the contract method 0x66a25830.
//
// Solidity: function setIndexTwInterval(uint256 _indexTwInterval) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) SetIndexTwInterval(_indexTwInterval *big.Int) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetIndexTwInterval(&_MarkPriceOracle.TransactOpts, _indexTwInterval)
}

// SetMarkTwInterval is a paid mutator transaction binding the contract method 0x849f4373.
//
// Solidity: function setMarkTwInterval(uint256 _markTwInterval) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) SetMarkTwInterval(opts *bind.TransactOpts, _markTwInterval *big.Int) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "setMarkTwInterval", _markTwInterval)
}

// SetMarkTwInterval is a paid mutator transaction binding the contract method 0x849f4373.
//
// Solidity: function setMarkTwInterval(uint256 _markTwInterval) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) SetMarkTwInterval(_markTwInterval *big.Int) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetMarkTwInterval(&_MarkPriceOracle.TransactOpts, _markTwInterval)
}

// SetMarkTwInterval is a paid mutator transaction binding the contract method 0x849f4373.
//
// Solidity: function setMarkTwInterval(uint256 _markTwInterval) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) SetMarkTwInterval(_markTwInterval *big.Int) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetMarkTwInterval(&_MarkPriceOracle.TransactOpts, _markTwInterval)
}

// SetObservationAdder is a paid mutator transaction binding the contract method 0x7b72a641.
//
// Solidity: function setObservationAdder(address _adder) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) SetObservationAdder(opts *bind.TransactOpts, _adder common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "setObservationAdder", _adder)
}

// SetObservationAdder is a paid mutator transaction binding the contract method 0x7b72a641.
//
// Solidity: function setObservationAdder(address _adder) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) SetObservationAdder(_adder common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetObservationAdder(&_MarkPriceOracle.TransactOpts, _adder)
}

// SetObservationAdder is a paid mutator transaction binding the contract method 0x7b72a641.
//
// Solidity: function setObservationAdder(address _adder) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) SetObservationAdder(_adder common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetObservationAdder(&_MarkPriceOracle.TransactOpts, _adder)
}

// SetPositioning is a paid mutator transaction binding the contract method 0xfa70d29e.
//
// Solidity: function setPositioning(address _positioning) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactor) SetPositioning(opts *bind.TransactOpts, _positioning common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.contract.Transact(opts, "setPositioning", _positioning)
}

// SetPositioning is a paid mutator transaction binding the contract method 0xfa70d29e.
//
// Solidity: function setPositioning(address _positioning) returns()
func (_MarkPriceOracle *MarkPriceOracleSession) SetPositioning(_positioning common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetPositioning(&_MarkPriceOracle.TransactOpts, _positioning)
}

// SetPositioning is a paid mutator transaction binding the contract method 0xfa70d29e.
//
// Solidity: function setPositioning(address _positioning) returns()
func (_MarkPriceOracle *MarkPriceOracleTransactorSession) SetPositioning(_positioning common.Address) (*types.Transaction, error) {
	return _MarkPriceOracle.Contract.SetPositioning(&_MarkPriceOracle.TransactOpts, _positioning)
}

// MarkPriceOracleAssetsAddedIterator is returned from FilterAssetsAdded and is used to iterate over the raw logs and unpacked data for AssetsAdded events raised by the MarkPriceOracle contract.
type MarkPriceOracleAssetsAddedIterator struct {
	Event *MarkPriceOracleAssetsAdded // Event containing the contract specifics and raw log

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
func (it *MarkPriceOracleAssetsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarkPriceOracleAssetsAdded)
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
		it.Event = new(MarkPriceOracleAssetsAdded)
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
func (it *MarkPriceOracleAssetsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarkPriceOracleAssetsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarkPriceOracleAssetsAdded represents a AssetsAdded event raised by the MarkPriceOracle contract.
type MarkPriceOracleAssetsAdded struct {
	LastIndex        *big.Int
	Assets           []common.Address
	UnderlyingPrices []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAssetsAdded is a free log retrieval operation binding the contract event 0xd54746bed1984556cd87b76fc3078d2670182f150532649f82479b40b3d9ab2d.
//
// Solidity: event AssetsAdded(uint256 indexed lastIndex, address[] assets, uint256[] underlyingPrices)
func (_MarkPriceOracle *MarkPriceOracleFilterer) FilterAssetsAdded(opts *bind.FilterOpts, lastIndex []*big.Int) (*MarkPriceOracleAssetsAddedIterator, error) {

	var lastIndexRule []interface{}
	for _, lastIndexItem := range lastIndex {
		lastIndexRule = append(lastIndexRule, lastIndexItem)
	}

	logs, sub, err := _MarkPriceOracle.contract.FilterLogs(opts, "AssetsAdded", lastIndexRule)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleAssetsAddedIterator{contract: _MarkPriceOracle.contract, event: "AssetsAdded", logs: logs, sub: sub}, nil
}

// WatchAssetsAdded is a free log subscription operation binding the contract event 0xd54746bed1984556cd87b76fc3078d2670182f150532649f82479b40b3d9ab2d.
//
// Solidity: event AssetsAdded(uint256 indexed lastIndex, address[] assets, uint256[] underlyingPrices)
func (_MarkPriceOracle *MarkPriceOracleFilterer) WatchAssetsAdded(opts *bind.WatchOpts, sink chan<- *MarkPriceOracleAssetsAdded, lastIndex []*big.Int) (event.Subscription, error) {

	var lastIndexRule []interface{}
	for _, lastIndexItem := range lastIndex {
		lastIndexRule = append(lastIndexRule, lastIndexItem)
	}

	logs, sub, err := _MarkPriceOracle.contract.WatchLogs(opts, "AssetsAdded", lastIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarkPriceOracleAssetsAdded)
				if err := _MarkPriceOracle.contract.UnpackLog(event, "AssetsAdded", log); err != nil {
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

// ParseAssetsAdded is a log parse operation binding the contract event 0xd54746bed1984556cd87b76fc3078d2670182f150532649f82479b40b3d9ab2d.
//
// Solidity: event AssetsAdded(uint256 indexed lastIndex, address[] assets, uint256[] underlyingPrices)
func (_MarkPriceOracle *MarkPriceOracleFilterer) ParseAssetsAdded(log types.Log) (*MarkPriceOracleAssetsAdded, error) {
	event := new(MarkPriceOracleAssetsAdded)
	if err := _MarkPriceOracle.contract.UnpackLog(event, "AssetsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarkPriceOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MarkPriceOracle contract.
type MarkPriceOracleInitializedIterator struct {
	Event *MarkPriceOracleInitialized // Event containing the contract specifics and raw log

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
func (it *MarkPriceOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarkPriceOracleInitialized)
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
		it.Event = new(MarkPriceOracleInitialized)
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
func (it *MarkPriceOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarkPriceOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarkPriceOracleInitialized represents a Initialized event raised by the MarkPriceOracle contract.
type MarkPriceOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MarkPriceOracle *MarkPriceOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarkPriceOracleInitializedIterator, error) {

	logs, sub, err := _MarkPriceOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleInitializedIterator{contract: _MarkPriceOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MarkPriceOracle *MarkPriceOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarkPriceOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _MarkPriceOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarkPriceOracleInitialized)
				if err := _MarkPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MarkPriceOracle *MarkPriceOracleFilterer) ParseInitialized(log types.Log) (*MarkPriceOracleInitialized, error) {
	event := new(MarkPriceOracleInitialized)
	if err := _MarkPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarkPriceOracleObservationAddedIterator is returned from FilterObservationAdded and is used to iterate over the raw logs and unpacked data for ObservationAdded events raised by the MarkPriceOracle contract.
type MarkPriceOracleObservationAddedIterator struct {
	Event *MarkPriceOracleObservationAdded // Event containing the contract specifics and raw log

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
func (it *MarkPriceOracleObservationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarkPriceOracleObservationAdded)
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
		it.Event = new(MarkPriceOracleObservationAdded)
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
func (it *MarkPriceOracleObservationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarkPriceOracleObservationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarkPriceOracleObservationAdded represents a ObservationAdded event raised by the MarkPriceOracle contract.
type MarkPriceOracleObservationAdded struct {
	Index           *big.Int
	UnderlyingPrice *big.Int
	MarkPrice       *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterObservationAdded is a free log retrieval operation binding the contract event 0xde2416383e7ff4b4db023368ad9bb5b4dac6b5a2151fdb4ca12ed35c9da916b2.
//
// Solidity: event ObservationAdded(uint256 indexed index, uint256 underlyingPrice, uint256 markPrice, uint256 timestamp)
func (_MarkPriceOracle *MarkPriceOracleFilterer) FilterObservationAdded(opts *bind.FilterOpts, index []*big.Int) (*MarkPriceOracleObservationAddedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _MarkPriceOracle.contract.FilterLogs(opts, "ObservationAdded", indexRule)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleObservationAddedIterator{contract: _MarkPriceOracle.contract, event: "ObservationAdded", logs: logs, sub: sub}, nil
}

// WatchObservationAdded is a free log subscription operation binding the contract event 0xde2416383e7ff4b4db023368ad9bb5b4dac6b5a2151fdb4ca12ed35c9da916b2.
//
// Solidity: event ObservationAdded(uint256 indexed index, uint256 underlyingPrice, uint256 markPrice, uint256 timestamp)
func (_MarkPriceOracle *MarkPriceOracleFilterer) WatchObservationAdded(opts *bind.WatchOpts, sink chan<- *MarkPriceOracleObservationAdded, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _MarkPriceOracle.contract.WatchLogs(opts, "ObservationAdded", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarkPriceOracleObservationAdded)
				if err := _MarkPriceOracle.contract.UnpackLog(event, "ObservationAdded", log); err != nil {
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

// ParseObservationAdded is a log parse operation binding the contract event 0xde2416383e7ff4b4db023368ad9bb5b4dac6b5a2151fdb4ca12ed35c9da916b2.
//
// Solidity: event ObservationAdded(uint256 indexed index, uint256 underlyingPrice, uint256 markPrice, uint256 timestamp)
func (_MarkPriceOracle *MarkPriceOracleFilterer) ParseObservationAdded(log types.Log) (*MarkPriceOracleObservationAdded, error) {
	event := new(MarkPriceOracleObservationAdded)
	if err := _MarkPriceOracle.contract.UnpackLog(event, "ObservationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarkPriceOracleObservationAdderSetIterator is returned from FilterObservationAdderSet and is used to iterate over the raw logs and unpacked data for ObservationAdderSet events raised by the MarkPriceOracle contract.
type MarkPriceOracleObservationAdderSetIterator struct {
	Event *MarkPriceOracleObservationAdderSet // Event containing the contract specifics and raw log

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
func (it *MarkPriceOracleObservationAdderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarkPriceOracleObservationAdderSet)
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
		it.Event = new(MarkPriceOracleObservationAdderSet)
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
func (it *MarkPriceOracleObservationAdderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarkPriceOracleObservationAdderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarkPriceOracleObservationAdderSet represents a ObservationAdderSet event raised by the MarkPriceOracle contract.
type MarkPriceOracleObservationAdderSet struct {
	MatchingEngine common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterObservationAdderSet is a free log retrieval operation binding the contract event 0x98b318ceaa8936e25855a68d4223b61f7cd0daa0c56d82005ca97c3ecef0ea94.
//
// Solidity: event ObservationAdderSet(address indexed matchingEngine)
func (_MarkPriceOracle *MarkPriceOracleFilterer) FilterObservationAdderSet(opts *bind.FilterOpts, matchingEngine []common.Address) (*MarkPriceOracleObservationAdderSetIterator, error) {

	var matchingEngineRule []interface{}
	for _, matchingEngineItem := range matchingEngine {
		matchingEngineRule = append(matchingEngineRule, matchingEngineItem)
	}

	logs, sub, err := _MarkPriceOracle.contract.FilterLogs(opts, "ObservationAdderSet", matchingEngineRule)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleObservationAdderSetIterator{contract: _MarkPriceOracle.contract, event: "ObservationAdderSet", logs: logs, sub: sub}, nil
}

// WatchObservationAdderSet is a free log subscription operation binding the contract event 0x98b318ceaa8936e25855a68d4223b61f7cd0daa0c56d82005ca97c3ecef0ea94.
//
// Solidity: event ObservationAdderSet(address indexed matchingEngine)
func (_MarkPriceOracle *MarkPriceOracleFilterer) WatchObservationAdderSet(opts *bind.WatchOpts, sink chan<- *MarkPriceOracleObservationAdderSet, matchingEngine []common.Address) (event.Subscription, error) {

	var matchingEngineRule []interface{}
	for _, matchingEngineItem := range matchingEngine {
		matchingEngineRule = append(matchingEngineRule, matchingEngineItem)
	}

	logs, sub, err := _MarkPriceOracle.contract.WatchLogs(opts, "ObservationAdderSet", matchingEngineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarkPriceOracleObservationAdderSet)
				if err := _MarkPriceOracle.contract.UnpackLog(event, "ObservationAdderSet", log); err != nil {
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

// ParseObservationAdderSet is a log parse operation binding the contract event 0x98b318ceaa8936e25855a68d4223b61f7cd0daa0c56d82005ca97c3ecef0ea94.
//
// Solidity: event ObservationAdderSet(address indexed matchingEngine)
func (_MarkPriceOracle *MarkPriceOracleFilterer) ParseObservationAdderSet(log types.Log) (*MarkPriceOracleObservationAdderSet, error) {
	event := new(MarkPriceOracleObservationAdderSet)
	if err := _MarkPriceOracle.contract.UnpackLog(event, "ObservationAdderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarkPriceOracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MarkPriceOracle contract.
type MarkPriceOracleRoleAdminChangedIterator struct {
	Event *MarkPriceOracleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MarkPriceOracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarkPriceOracleRoleAdminChanged)
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
		it.Event = new(MarkPriceOracleRoleAdminChanged)
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
func (it *MarkPriceOracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarkPriceOracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarkPriceOracleRoleAdminChanged represents a RoleAdminChanged event raised by the MarkPriceOracle contract.
type MarkPriceOracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MarkPriceOracle *MarkPriceOracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MarkPriceOracleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _MarkPriceOracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleRoleAdminChangedIterator{contract: _MarkPriceOracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MarkPriceOracle *MarkPriceOracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MarkPriceOracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _MarkPriceOracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarkPriceOracleRoleAdminChanged)
				if err := _MarkPriceOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_MarkPriceOracle *MarkPriceOracleFilterer) ParseRoleAdminChanged(log types.Log) (*MarkPriceOracleRoleAdminChanged, error) {
	event := new(MarkPriceOracleRoleAdminChanged)
	if err := _MarkPriceOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarkPriceOracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MarkPriceOracle contract.
type MarkPriceOracleRoleGrantedIterator struct {
	Event *MarkPriceOracleRoleGranted // Event containing the contract specifics and raw log

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
func (it *MarkPriceOracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarkPriceOracleRoleGranted)
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
		it.Event = new(MarkPriceOracleRoleGranted)
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
func (it *MarkPriceOracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarkPriceOracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarkPriceOracleRoleGranted represents a RoleGranted event raised by the MarkPriceOracle contract.
type MarkPriceOracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MarkPriceOracle *MarkPriceOracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarkPriceOracleRoleGrantedIterator, error) {

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

	logs, sub, err := _MarkPriceOracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleRoleGrantedIterator{contract: _MarkPriceOracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MarkPriceOracle *MarkPriceOracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MarkPriceOracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MarkPriceOracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarkPriceOracleRoleGranted)
				if err := _MarkPriceOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_MarkPriceOracle *MarkPriceOracleFilterer) ParseRoleGranted(log types.Log) (*MarkPriceOracleRoleGranted, error) {
	event := new(MarkPriceOracleRoleGranted)
	if err := _MarkPriceOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarkPriceOracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MarkPriceOracle contract.
type MarkPriceOracleRoleRevokedIterator struct {
	Event *MarkPriceOracleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MarkPriceOracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarkPriceOracleRoleRevoked)
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
		it.Event = new(MarkPriceOracleRoleRevoked)
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
func (it *MarkPriceOracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarkPriceOracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarkPriceOracleRoleRevoked represents a RoleRevoked event raised by the MarkPriceOracle contract.
type MarkPriceOracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MarkPriceOracle *MarkPriceOracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarkPriceOracleRoleRevokedIterator, error) {

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

	logs, sub, err := _MarkPriceOracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarkPriceOracleRoleRevokedIterator{contract: _MarkPriceOracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MarkPriceOracle *MarkPriceOracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MarkPriceOracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MarkPriceOracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarkPriceOracleRoleRevoked)
				if err := _MarkPriceOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_MarkPriceOracle *MarkPriceOracleFilterer) ParseRoleRevoked(log types.Log) (*MarkPriceOracleRoleRevoked, error) {
	event := new(MarkPriceOracleRoleRevoked)
	if err := _MarkPriceOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
