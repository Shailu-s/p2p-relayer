// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IndexPriceOracle

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

// IndexPriceOracleMetaData contains all meta data concerning the IndexPriceOracle contract.
var IndexPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"underlyingPrices\",\"type\":\"uint256[]\"}],\"name\":\"AssetsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"index\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"underlyingPrice\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ObservationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"matchingEngine\",\"type\":\"address\"}],\"name\":\"ObservationAdderSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADD_OBSERVATION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_TIMESTAMP_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_ORACLE_ADMIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_underlyingPrice\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_asset\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_capRatio\",\"type\":\"uint256[]\"}],\"name\":\"addAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_underlyingPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofHashes\",\"type\":\"bytes32[]\"}],\"name\":\"addObservation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"baseTokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cardinality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCustomEpochPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCustomIndexSma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceCumulative\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIndexCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getIndexObservation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getIndexPriceByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_smInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getIndexSma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"volatilityTokenSma\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iVolatilityTokenSma\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLastEpochPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingLastPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_smInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLastSma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceCumulative\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLastUpdatedTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastUpdatedTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"grantInitialTimestampRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexByBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"indexPriceAtEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexSmInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_volatilityPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_volatilityIndex\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_capRatio\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_smInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"answer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"observationsByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_smInterval\",\"type\":\"uint256\"}],\"name\":\"setIndexSmInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"setInitialTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adder\",\"type\":\"address\"}],\"name\":\"setObservationAdder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"volatilityCapRatioByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IndexPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IndexPriceOracleMetaData.ABI instead.
var IndexPriceOracleABI = IndexPriceOracleMetaData.ABI

// IndexPriceOracle is an auto generated Go binding around an Ethereum contract.
type IndexPriceOracle struct {
	IndexPriceOracleCaller     // Read-only binding to the contract
	IndexPriceOracleTransactor // Write-only binding to the contract
	IndexPriceOracleFilterer   // Log filterer for contract events
}

// IndexPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndexPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndexPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IndexPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IndexPriceOracleSession struct {
	Contract     *IndexPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IndexPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IndexPriceOracleCallerSession struct {
	Contract *IndexPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IndexPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IndexPriceOracleTransactorSession struct {
	Contract     *IndexPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IndexPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IndexPriceOracleRaw struct {
	Contract *IndexPriceOracle // Generic contract binding to access the raw methods on
}

// IndexPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IndexPriceOracleCallerRaw struct {
	Contract *IndexPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IndexPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IndexPriceOracleTransactorRaw struct {
	Contract *IndexPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndexPriceOracle creates a new instance of IndexPriceOracle, bound to a specific deployed contract.
func NewIndexPriceOracle(address common.Address, backend bind.ContractBackend) (*IndexPriceOracle, error) {
	contract, err := bindIndexPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracle{IndexPriceOracleCaller: IndexPriceOracleCaller{contract: contract}, IndexPriceOracleTransactor: IndexPriceOracleTransactor{contract: contract}, IndexPriceOracleFilterer: IndexPriceOracleFilterer{contract: contract}}, nil
}

// NewIndexPriceOracleCaller creates a new read-only instance of IndexPriceOracle, bound to a specific deployed contract.
func NewIndexPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*IndexPriceOracleCaller, error) {
	contract, err := bindIndexPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleCaller{contract: contract}, nil
}

// NewIndexPriceOracleTransactor creates a new write-only instance of IndexPriceOracle, bound to a specific deployed contract.
func NewIndexPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IndexPriceOracleTransactor, error) {
	contract, err := bindIndexPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleTransactor{contract: contract}, nil
}

// NewIndexPriceOracleFilterer creates a new log filterer instance of IndexPriceOracle, bound to a specific deployed contract.
func NewIndexPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IndexPriceOracleFilterer, error) {
	contract, err := bindIndexPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleFilterer{contract: contract}, nil
}

// bindIndexPriceOracle binds a generic wrapper to an already deployed contract.
func bindIndexPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IndexPriceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IndexPriceOracle *IndexPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IndexPriceOracle.Contract.IndexPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IndexPriceOracle *IndexPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.IndexPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IndexPriceOracle *IndexPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.IndexPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IndexPriceOracle *IndexPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IndexPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IndexPriceOracle *IndexPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IndexPriceOracle *IndexPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// ADDOBSERVATIONROLE is a free data retrieval call binding the contract method 0x0db0f0f5.
//
// Solidity: function ADD_OBSERVATION_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCaller) ADDOBSERVATIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "ADD_OBSERVATION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDOBSERVATIONROLE is a free data retrieval call binding the contract method 0x0db0f0f5.
//
// Solidity: function ADD_OBSERVATION_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleSession) ADDOBSERVATIONROLE() ([32]byte, error) {
	return _IndexPriceOracle.Contract.ADDOBSERVATIONROLE(&_IndexPriceOracle.CallOpts)
}

// ADDOBSERVATIONROLE is a free data retrieval call binding the contract method 0x0db0f0f5.
//
// Solidity: function ADD_OBSERVATION_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) ADDOBSERVATIONROLE() ([32]byte, error) {
	return _IndexPriceOracle.Contract.ADDOBSERVATIONROLE(&_IndexPriceOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IndexPriceOracle.Contract.DEFAULTADMINROLE(&_IndexPriceOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IndexPriceOracle.Contract.DEFAULTADMINROLE(&_IndexPriceOracle.CallOpts)
}

// INITIALTIMESTAMPROLE is a free data retrieval call binding the contract method 0x304e8491.
//
// Solidity: function INITIAL_TIMESTAMP_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCaller) INITIALTIMESTAMPROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "INITIAL_TIMESTAMP_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIALTIMESTAMPROLE is a free data retrieval call binding the contract method 0x304e8491.
//
// Solidity: function INITIAL_TIMESTAMP_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleSession) INITIALTIMESTAMPROLE() ([32]byte, error) {
	return _IndexPriceOracle.Contract.INITIALTIMESTAMPROLE(&_IndexPriceOracle.CallOpts)
}

// INITIALTIMESTAMPROLE is a free data retrieval call binding the contract method 0x304e8491.
//
// Solidity: function INITIAL_TIMESTAMP_ROLE() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) INITIALTIMESTAMPROLE() ([32]byte, error) {
	return _IndexPriceOracle.Contract.INITIALTIMESTAMPROLE(&_IndexPriceOracle.CallOpts)
}

// PRICEORACLEADMIN is a free data retrieval call binding the contract method 0xa1f68193.
//
// Solidity: function PRICE_ORACLE_ADMIN() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCaller) PRICEORACLEADMIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "PRICE_ORACLE_ADMIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRICEORACLEADMIN is a free data retrieval call binding the contract method 0xa1f68193.
//
// Solidity: function PRICE_ORACLE_ADMIN() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleSession) PRICEORACLEADMIN() ([32]byte, error) {
	return _IndexPriceOracle.Contract.PRICEORACLEADMIN(&_IndexPriceOracle.CallOpts)
}

// PRICEORACLEADMIN is a free data retrieval call binding the contract method 0xa1f68193.
//
// Solidity: function PRICE_ORACLE_ADMIN() view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) PRICEORACLEADMIN() ([32]byte, error) {
	return _IndexPriceOracle.Contract.PRICEORACLEADMIN(&_IndexPriceOracle.CallOpts)
}

// BaseTokenByIndex is a free data retrieval call binding the contract method 0x599a945c.
//
// Solidity: function baseTokenByIndex(uint256 ) view returns(address)
func (_IndexPriceOracle *IndexPriceOracleCaller) BaseTokenByIndex(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "baseTokenByIndex", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseTokenByIndex is a free data retrieval call binding the contract method 0x599a945c.
//
// Solidity: function baseTokenByIndex(uint256 ) view returns(address)
func (_IndexPriceOracle *IndexPriceOracleSession) BaseTokenByIndex(arg0 *big.Int) (common.Address, error) {
	return _IndexPriceOracle.Contract.BaseTokenByIndex(&_IndexPriceOracle.CallOpts, arg0)
}

// BaseTokenByIndex is a free data retrieval call binding the contract method 0x599a945c.
//
// Solidity: function baseTokenByIndex(uint256 ) view returns(address)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) BaseTokenByIndex(arg0 *big.Int) (common.Address, error) {
	return _IndexPriceOracle.Contract.BaseTokenByIndex(&_IndexPriceOracle.CallOpts, arg0)
}

// Cardinality is a free data retrieval call binding the contract method 0xdbffe9ad.
//
// Solidity: function cardinality() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCaller) Cardinality(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "cardinality")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cardinality is a free data retrieval call binding the contract method 0xdbffe9ad.
//
// Solidity: function cardinality() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleSession) Cardinality() (*big.Int, error) {
	return _IndexPriceOracle.Contract.Cardinality(&_IndexPriceOracle.CallOpts)
}

// Cardinality is a free data retrieval call binding the contract method 0xdbffe9ad.
//
// Solidity: function cardinality() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) Cardinality() (*big.Int, error) {
	return _IndexPriceOracle.Contract.Cardinality(&_IndexPriceOracle.CallOpts)
}

// GetCustomEpochPrice is a free data retrieval call binding the contract method 0xa819f650.
//
// Solidity: function getCustomEpochPrice(uint256 _index, uint256 _epochTimestamp) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetCustomEpochPrice(opts *bind.CallOpts, _index *big.Int, _epochTimestamp *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getCustomEpochPrice", _index, _epochTimestamp)

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCustomEpochPrice is a free data retrieval call binding the contract method 0xa819f650.
//
// Solidity: function getCustomEpochPrice(uint256 _index, uint256 _epochTimestamp) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleSession) GetCustomEpochPrice(_index *big.Int, _epochTimestamp *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.GetCustomEpochPrice(&_IndexPriceOracle.CallOpts, _index, _epochTimestamp)
}

// GetCustomEpochPrice is a free data retrieval call binding the contract method 0xa819f650.
//
// Solidity: function getCustomEpochPrice(uint256 _index, uint256 _epochTimestamp) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetCustomEpochPrice(_index *big.Int, _epochTimestamp *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.GetCustomEpochPrice(&_IndexPriceOracle.CallOpts, _index, _epochTimestamp)
}

// GetCustomIndexSma is a free data retrieval call binding the contract method 0x447f18ee.
//
// Solidity: function getCustomIndexSma(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetCustomIndexSma(opts *bind.CallOpts, _index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getCustomIndexSma", _index, _startTimestamp, _endTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCustomIndexSma is a free data retrieval call binding the contract method 0x447f18ee.
//
// Solidity: function getCustomIndexSma(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_IndexPriceOracle *IndexPriceOracleSession) GetCustomIndexSma(_index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetCustomIndexSma(&_IndexPriceOracle.CallOpts, _index, _startTimestamp, _endTimestamp)
}

// GetCustomIndexSma is a free data retrieval call binding the contract method 0x447f18ee.
//
// Solidity: function getCustomIndexSma(uint256 _index, uint256 _startTimestamp, uint256 _endTimestamp) view returns(uint256 priceCumulative)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetCustomIndexSma(_index *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetCustomIndexSma(&_IndexPriceOracle.CallOpts, _index, _startTimestamp, _endTimestamp)
}

// GetIndexCount is a free data retrieval call binding the contract method 0xec244370.
//
// Solidity: function getIndexCount() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetIndexCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getIndexCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexCount is a free data retrieval call binding the contract method 0xec244370.
//
// Solidity: function getIndexCount() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleSession) GetIndexCount() (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetIndexCount(&_IndexPriceOracle.CallOpts)
}

// GetIndexCount is a free data retrieval call binding the contract method 0xec244370.
//
// Solidity: function getIndexCount() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetIndexCount() (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetIndexCount(&_IndexPriceOracle.CallOpts)
}

// GetIndexObservation is a free data retrieval call binding the contract method 0xf1658e0c.
//
// Solidity: function getIndexObservation(uint256 _index) view returns(uint256 length)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetIndexObservation(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getIndexObservation", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexObservation is a free data retrieval call binding the contract method 0xf1658e0c.
//
// Solidity: function getIndexObservation(uint256 _index) view returns(uint256 length)
func (_IndexPriceOracle *IndexPriceOracleSession) GetIndexObservation(_index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetIndexObservation(&_IndexPriceOracle.CallOpts, _index)
}

// GetIndexObservation is a free data retrieval call binding the contract method 0xf1658e0c.
//
// Solidity: function getIndexObservation(uint256 _index) view returns(uint256 length)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetIndexObservation(_index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetIndexObservation(&_IndexPriceOracle.CallOpts, _index)
}

// GetIndexPriceByEpoch is a free data retrieval call binding the contract method 0x4fe09b2a.
//
// Solidity: function getIndexPriceByEpoch(uint256 _index) view returns(uint256 length)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetIndexPriceByEpoch(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getIndexPriceByEpoch", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexPriceByEpoch is a free data retrieval call binding the contract method 0x4fe09b2a.
//
// Solidity: function getIndexPriceByEpoch(uint256 _index) view returns(uint256 length)
func (_IndexPriceOracle *IndexPriceOracleSession) GetIndexPriceByEpoch(_index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetIndexPriceByEpoch(&_IndexPriceOracle.CallOpts, _index)
}

// GetIndexPriceByEpoch is a free data retrieval call binding the contract method 0x4fe09b2a.
//
// Solidity: function getIndexPriceByEpoch(uint256 _index) view returns(uint256 length)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetIndexPriceByEpoch(_index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetIndexPriceByEpoch(&_IndexPriceOracle.CallOpts, _index)
}

// GetIndexSma is a free data retrieval call binding the contract method 0xc3b621b9.
//
// Solidity: function getIndexSma(uint256 _smInterval, uint256 _index) view returns(uint256 volatilityTokenSma, uint256 iVolatilityTokenSma, uint256 lastUpdateTimestamp)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetIndexSma(opts *bind.CallOpts, _smInterval *big.Int, _index *big.Int) (struct {
	VolatilityTokenSma  *big.Int
	IVolatilityTokenSma *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getIndexSma", _smInterval, _index)

	outstruct := new(struct {
		VolatilityTokenSma  *big.Int
		IVolatilityTokenSma *big.Int
		LastUpdateTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VolatilityTokenSma = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IVolatilityTokenSma = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetIndexSma is a free data retrieval call binding the contract method 0xc3b621b9.
//
// Solidity: function getIndexSma(uint256 _smInterval, uint256 _index) view returns(uint256 volatilityTokenSma, uint256 iVolatilityTokenSma, uint256 lastUpdateTimestamp)
func (_IndexPriceOracle *IndexPriceOracleSession) GetIndexSma(_smInterval *big.Int, _index *big.Int) (struct {
	VolatilityTokenSma  *big.Int
	IVolatilityTokenSma *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.GetIndexSma(&_IndexPriceOracle.CallOpts, _smInterval, _index)
}

// GetIndexSma is a free data retrieval call binding the contract method 0xc3b621b9.
//
// Solidity: function getIndexSma(uint256 _smInterval, uint256 _index) view returns(uint256 volatilityTokenSma, uint256 iVolatilityTokenSma, uint256 lastUpdateTimestamp)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetIndexSma(_smInterval *big.Int, _index *big.Int) (struct {
	VolatilityTokenSma  *big.Int
	IVolatilityTokenSma *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.GetIndexSma(&_IndexPriceOracle.CallOpts, _smInterval, _index)
}

// GetLastEpochPrice is a free data retrieval call binding the contract method 0x217b14d9.
//
// Solidity: function getLastEpochPrice(uint256 _index) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetLastEpochPrice(opts *bind.CallOpts, _index *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getLastEpochPrice", _index)

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLastEpochPrice is a free data retrieval call binding the contract method 0x217b14d9.
//
// Solidity: function getLastEpochPrice(uint256 _index) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleSession) GetLastEpochPrice(_index *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.GetLastEpochPrice(&_IndexPriceOracle.CallOpts, _index)
}

// GetLastEpochPrice is a free data retrieval call binding the contract method 0x217b14d9.
//
// Solidity: function getLastEpochPrice(uint256 _index) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetLastEpochPrice(_index *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.GetLastEpochPrice(&_IndexPriceOracle.CallOpts, _index)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x65fa2f7f.
//
// Solidity: function getLastPrice(uint256 _index) view returns(uint256 underlyingLastPrice)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetLastPrice(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getLastPrice", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPrice is a free data retrieval call binding the contract method 0x65fa2f7f.
//
// Solidity: function getLastPrice(uint256 _index) view returns(uint256 underlyingLastPrice)
func (_IndexPriceOracle *IndexPriceOracleSession) GetLastPrice(_index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetLastPrice(&_IndexPriceOracle.CallOpts, _index)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x65fa2f7f.
//
// Solidity: function getLastPrice(uint256 _index) view returns(uint256 underlyingLastPrice)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetLastPrice(_index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetLastPrice(&_IndexPriceOracle.CallOpts, _index)
}

// GetLastSma is a free data retrieval call binding the contract method 0xed63b347.
//
// Solidity: function getLastSma(uint256 _smInterval, uint256 _index) view returns(uint256 priceCumulative)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetLastSma(opts *bind.CallOpts, _smInterval *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getLastSma", _smInterval, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastSma is a free data retrieval call binding the contract method 0xed63b347.
//
// Solidity: function getLastSma(uint256 _smInterval, uint256 _index) view returns(uint256 priceCumulative)
func (_IndexPriceOracle *IndexPriceOracleSession) GetLastSma(_smInterval *big.Int, _index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetLastSma(&_IndexPriceOracle.CallOpts, _smInterval, _index)
}

// GetLastSma is a free data retrieval call binding the contract method 0xed63b347.
//
// Solidity: function getLastSma(uint256 _smInterval, uint256 _index) view returns(uint256 priceCumulative)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetLastSma(_smInterval *big.Int, _index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetLastSma(&_IndexPriceOracle.CallOpts, _smInterval, _index)
}

// GetLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xa367c24d.
//
// Solidity: function getLastUpdatedTimestamp(uint256 _index) view returns(uint256 lastUpdatedTimestamp)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetLastUpdatedTimestamp(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getLastUpdatedTimestamp", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xa367c24d.
//
// Solidity: function getLastUpdatedTimestamp(uint256 _index) view returns(uint256 lastUpdatedTimestamp)
func (_IndexPriceOracle *IndexPriceOracleSession) GetLastUpdatedTimestamp(_index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetLastUpdatedTimestamp(&_IndexPriceOracle.CallOpts, _index)
}

// GetLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xa367c24d.
//
// Solidity: function getLastUpdatedTimestamp(uint256 _index) view returns(uint256 lastUpdatedTimestamp)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetLastUpdatedTimestamp(_index *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.GetLastUpdatedTimestamp(&_IndexPriceOracle.CallOpts, _index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IndexPriceOracle.Contract.GetRoleAdmin(&_IndexPriceOracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IndexPriceOracle.Contract.GetRoleAdmin(&_IndexPriceOracle.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IndexPriceOracle *IndexPriceOracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IndexPriceOracle *IndexPriceOracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IndexPriceOracle.Contract.HasRole(&_IndexPriceOracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IndexPriceOracle.Contract.HasRole(&_IndexPriceOracle.CallOpts, role, account)
}

// IndexByBaseToken is a free data retrieval call binding the contract method 0xac416c17.
//
// Solidity: function indexByBaseToken(address ) view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCaller) IndexByBaseToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "indexByBaseToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexByBaseToken is a free data retrieval call binding the contract method 0xac416c17.
//
// Solidity: function indexByBaseToken(address ) view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleSession) IndexByBaseToken(arg0 common.Address) (*big.Int, error) {
	return _IndexPriceOracle.Contract.IndexByBaseToken(&_IndexPriceOracle.CallOpts, arg0)
}

// IndexByBaseToken is a free data retrieval call binding the contract method 0xac416c17.
//
// Solidity: function indexByBaseToken(address ) view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) IndexByBaseToken(arg0 common.Address) (*big.Int, error) {
	return _IndexPriceOracle.Contract.IndexByBaseToken(&_IndexPriceOracle.CallOpts, arg0)
}

// IndexPriceAtEpochs is a free data retrieval call binding the contract method 0xfb7566c5.
//
// Solidity: function indexPriceAtEpochs(uint256 , uint256 ) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleCaller) IndexPriceAtEpochs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "indexPriceAtEpochs", arg0, arg1)

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IndexPriceAtEpochs is a free data retrieval call binding the contract method 0xfb7566c5.
//
// Solidity: function indexPriceAtEpochs(uint256 , uint256 ) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleSession) IndexPriceAtEpochs(arg0 *big.Int, arg1 *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.IndexPriceAtEpochs(&_IndexPriceOracle.CallOpts, arg0, arg1)
}

// IndexPriceAtEpochs is a free data retrieval call binding the contract method 0xfb7566c5.
//
// Solidity: function indexPriceAtEpochs(uint256 , uint256 ) view returns(uint256 price, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) IndexPriceAtEpochs(arg0 *big.Int, arg1 *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.IndexPriceAtEpochs(&_IndexPriceOracle.CallOpts, arg0, arg1)
}

// IndexSmInterval is a free data retrieval call binding the contract method 0xab57dfa5.
//
// Solidity: function indexSmInterval() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCaller) IndexSmInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "indexSmInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexSmInterval is a free data retrieval call binding the contract method 0xab57dfa5.
//
// Solidity: function indexSmInterval() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleSession) IndexSmInterval() (*big.Int, error) {
	return _IndexPriceOracle.Contract.IndexSmInterval(&_IndexPriceOracle.CallOpts)
}

// IndexSmInterval is a free data retrieval call binding the contract method 0xab57dfa5.
//
// Solidity: function indexSmInterval() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) IndexSmInterval() (*big.Int, error) {
	return _IndexPriceOracle.Contract.IndexSmInterval(&_IndexPriceOracle.CallOpts)
}

// InitialTimestamp is a free data retrieval call binding the contract method 0xd6d14171.
//
// Solidity: function initialTimestamp() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCaller) InitialTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "initialTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialTimestamp is a free data retrieval call binding the contract method 0xd6d14171.
//
// Solidity: function initialTimestamp() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleSession) InitialTimestamp() (*big.Int, error) {
	return _IndexPriceOracle.Contract.InitialTimestamp(&_IndexPriceOracle.CallOpts)
}

// InitialTimestamp is a free data retrieval call binding the contract method 0xd6d14171.
//
// Solidity: function initialTimestamp() view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) InitialTimestamp() (*big.Int, error) {
	return _IndexPriceOracle.Contract.InitialTimestamp(&_IndexPriceOracle.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xca92f873.
//
// Solidity: function latestRoundData(uint256 _smInterval, uint256 _index) view returns(uint256 answer, uint256 lastUpdateTimestamp)
func (_IndexPriceOracle *IndexPriceOracleCaller) LatestRoundData(opts *bind.CallOpts, _smInterval *big.Int, _index *big.Int) (struct {
	Answer              *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "latestRoundData", _smInterval, _index)

	outstruct := new(struct {
		Answer              *big.Int
		LastUpdateTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Answer = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xca92f873.
//
// Solidity: function latestRoundData(uint256 _smInterval, uint256 _index) view returns(uint256 answer, uint256 lastUpdateTimestamp)
func (_IndexPriceOracle *IndexPriceOracleSession) LatestRoundData(_smInterval *big.Int, _index *big.Int) (struct {
	Answer              *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.LatestRoundData(&_IndexPriceOracle.CallOpts, _smInterval, _index)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xca92f873.
//
// Solidity: function latestRoundData(uint256 _smInterval, uint256 _index) view returns(uint256 answer, uint256 lastUpdateTimestamp)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) LatestRoundData(_smInterval *big.Int, _index *big.Int) (struct {
	Answer              *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _IndexPriceOracle.Contract.LatestRoundData(&_IndexPriceOracle.CallOpts, _smInterval, _index)
}

// ObservationsByIndex is a free data retrieval call binding the contract method 0x80319d88.
//
// Solidity: function observationsByIndex(uint256 , uint256 ) view returns(uint256 timestamp, uint256 underlyingPrice, bytes32 proofHash)
func (_IndexPriceOracle *IndexPriceOracleCaller) ObservationsByIndex(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp       *big.Int
	UnderlyingPrice *big.Int
	ProofHash       [32]byte
}, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "observationsByIndex", arg0, arg1)

	outstruct := new(struct {
		Timestamp       *big.Int
		UnderlyingPrice *big.Int
		ProofHash       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnderlyingPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProofHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// ObservationsByIndex is a free data retrieval call binding the contract method 0x80319d88.
//
// Solidity: function observationsByIndex(uint256 , uint256 ) view returns(uint256 timestamp, uint256 underlyingPrice, bytes32 proofHash)
func (_IndexPriceOracle *IndexPriceOracleSession) ObservationsByIndex(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp       *big.Int
	UnderlyingPrice *big.Int
	ProofHash       [32]byte
}, error) {
	return _IndexPriceOracle.Contract.ObservationsByIndex(&_IndexPriceOracle.CallOpts, arg0, arg1)
}

// ObservationsByIndex is a free data retrieval call binding the contract method 0x80319d88.
//
// Solidity: function observationsByIndex(uint256 , uint256 ) view returns(uint256 timestamp, uint256 underlyingPrice, bytes32 proofHash)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) ObservationsByIndex(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp       *big.Int
	UnderlyingPrice *big.Int
	ProofHash       [32]byte
}, error) {
	return _IndexPriceOracle.Contract.ObservationsByIndex(&_IndexPriceOracle.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IndexPriceOracle *IndexPriceOracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IndexPriceOracle *IndexPriceOracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IndexPriceOracle.Contract.SupportsInterface(&_IndexPriceOracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IndexPriceOracle.Contract.SupportsInterface(&_IndexPriceOracle.CallOpts, interfaceId)
}

// VolatilityCapRatioByIndex is a free data retrieval call binding the contract method 0x4704c96b.
//
// Solidity: function volatilityCapRatioByIndex(uint256 ) view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCaller) VolatilityCapRatioByIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IndexPriceOracle.contract.Call(opts, &out, "volatilityCapRatioByIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VolatilityCapRatioByIndex is a free data retrieval call binding the contract method 0x4704c96b.
//
// Solidity: function volatilityCapRatioByIndex(uint256 ) view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleSession) VolatilityCapRatioByIndex(arg0 *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.VolatilityCapRatioByIndex(&_IndexPriceOracle.CallOpts, arg0)
}

// VolatilityCapRatioByIndex is a free data retrieval call binding the contract method 0x4704c96b.
//
// Solidity: function volatilityCapRatioByIndex(uint256 ) view returns(uint256)
func (_IndexPriceOracle *IndexPriceOracleCallerSession) VolatilityCapRatioByIndex(arg0 *big.Int) (*big.Int, error) {
	return _IndexPriceOracle.Contract.VolatilityCapRatioByIndex(&_IndexPriceOracle.CallOpts, arg0)
}

// AddAssets is a paid mutator transaction binding the contract method 0x021adf31.
//
// Solidity: function addAssets(uint256[] _underlyingPrice, address[] _asset, bytes32[] _proofHash, uint256[] _capRatio) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) AddAssets(opts *bind.TransactOpts, _underlyingPrice []*big.Int, _asset []common.Address, _proofHash [][32]byte, _capRatio []*big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "addAssets", _underlyingPrice, _asset, _proofHash, _capRatio)
}

// AddAssets is a paid mutator transaction binding the contract method 0x021adf31.
//
// Solidity: function addAssets(uint256[] _underlyingPrice, address[] _asset, bytes32[] _proofHash, uint256[] _capRatio) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) AddAssets(_underlyingPrice []*big.Int, _asset []common.Address, _proofHash [][32]byte, _capRatio []*big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.AddAssets(&_IndexPriceOracle.TransactOpts, _underlyingPrice, _asset, _proofHash, _capRatio)
}

// AddAssets is a paid mutator transaction binding the contract method 0x021adf31.
//
// Solidity: function addAssets(uint256[] _underlyingPrice, address[] _asset, bytes32[] _proofHash, uint256[] _capRatio) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) AddAssets(_underlyingPrice []*big.Int, _asset []common.Address, _proofHash [][32]byte, _capRatio []*big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.AddAssets(&_IndexPriceOracle.TransactOpts, _underlyingPrice, _asset, _proofHash, _capRatio)
}

// AddObservation is a paid mutator transaction binding the contract method 0x1acb349b.
//
// Solidity: function addObservation(uint256[] _underlyingPrices, uint256[] _indexes, bytes32[] _proofHashes) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) AddObservation(opts *bind.TransactOpts, _underlyingPrices []*big.Int, _indexes []*big.Int, _proofHashes [][32]byte) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "addObservation", _underlyingPrices, _indexes, _proofHashes)
}

// AddObservation is a paid mutator transaction binding the contract method 0x1acb349b.
//
// Solidity: function addObservation(uint256[] _underlyingPrices, uint256[] _indexes, bytes32[] _proofHashes) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) AddObservation(_underlyingPrices []*big.Int, _indexes []*big.Int, _proofHashes [][32]byte) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.AddObservation(&_IndexPriceOracle.TransactOpts, _underlyingPrices, _indexes, _proofHashes)
}

// AddObservation is a paid mutator transaction binding the contract method 0x1acb349b.
//
// Solidity: function addObservation(uint256[] _underlyingPrices, uint256[] _indexes, bytes32[] _proofHashes) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) AddObservation(_underlyingPrices []*big.Int, _indexes []*big.Int, _proofHashes [][32]byte) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.AddObservation(&_IndexPriceOracle.TransactOpts, _underlyingPrices, _indexes, _proofHashes)
}

// GrantInitialTimestampRole is a paid mutator transaction binding the contract method 0x9fadf603.
//
// Solidity: function grantInitialTimestampRole(address _account) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) GrantInitialTimestampRole(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "grantInitialTimestampRole", _account)
}

// GrantInitialTimestampRole is a paid mutator transaction binding the contract method 0x9fadf603.
//
// Solidity: function grantInitialTimestampRole(address _account) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) GrantInitialTimestampRole(_account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.GrantInitialTimestampRole(&_IndexPriceOracle.TransactOpts, _account)
}

// GrantInitialTimestampRole is a paid mutator transaction binding the contract method 0x9fadf603.
//
// Solidity: function grantInitialTimestampRole(address _account) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) GrantInitialTimestampRole(_account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.GrantInitialTimestampRole(&_IndexPriceOracle.TransactOpts, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.GrantRole(&_IndexPriceOracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.GrantRole(&_IndexPriceOracle.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e4efa1e.
//
// Solidity: function initialize(address _admin, uint256[] _volatilityPrices, address[] _volatilityIndex, bytes32[] _proofHash, uint256[] _capRatio) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _volatilityPrices []*big.Int, _volatilityIndex []common.Address, _proofHash [][32]byte, _capRatio []*big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "initialize", _admin, _volatilityPrices, _volatilityIndex, _proofHash, _capRatio)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e4efa1e.
//
// Solidity: function initialize(address _admin, uint256[] _volatilityPrices, address[] _volatilityIndex, bytes32[] _proofHash, uint256[] _capRatio) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) Initialize(_admin common.Address, _volatilityPrices []*big.Int, _volatilityIndex []common.Address, _proofHash [][32]byte, _capRatio []*big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.Initialize(&_IndexPriceOracle.TransactOpts, _admin, _volatilityPrices, _volatilityIndex, _proofHash, _capRatio)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e4efa1e.
//
// Solidity: function initialize(address _admin, uint256[] _volatilityPrices, address[] _volatilityIndex, bytes32[] _proofHash, uint256[] _capRatio) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) Initialize(_admin common.Address, _volatilityPrices []*big.Int, _volatilityIndex []common.Address, _proofHash [][32]byte, _capRatio []*big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.Initialize(&_IndexPriceOracle.TransactOpts, _admin, _volatilityPrices, _volatilityIndex, _proofHash, _capRatio)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.RenounceRole(&_IndexPriceOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.RenounceRole(&_IndexPriceOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.RevokeRole(&_IndexPriceOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.RevokeRole(&_IndexPriceOracle.TransactOpts, role, account)
}

// SetIndexSmInterval is a paid mutator transaction binding the contract method 0xfe4ab28a.
//
// Solidity: function setIndexSmInterval(uint256 _smInterval) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) SetIndexSmInterval(opts *bind.TransactOpts, _smInterval *big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "setIndexSmInterval", _smInterval)
}

// SetIndexSmInterval is a paid mutator transaction binding the contract method 0xfe4ab28a.
//
// Solidity: function setIndexSmInterval(uint256 _smInterval) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) SetIndexSmInterval(_smInterval *big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.SetIndexSmInterval(&_IndexPriceOracle.TransactOpts, _smInterval)
}

// SetIndexSmInterval is a paid mutator transaction binding the contract method 0xfe4ab28a.
//
// Solidity: function setIndexSmInterval(uint256 _smInterval) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) SetIndexSmInterval(_smInterval *big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.SetIndexSmInterval(&_IndexPriceOracle.TransactOpts, _smInterval)
}

// SetInitialTimestamp is a paid mutator transaction binding the contract method 0x21f926db.
//
// Solidity: function setInitialTimestamp(uint256 _timestamp) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) SetInitialTimestamp(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "setInitialTimestamp", _timestamp)
}

// SetInitialTimestamp is a paid mutator transaction binding the contract method 0x21f926db.
//
// Solidity: function setInitialTimestamp(uint256 _timestamp) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) SetInitialTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.SetInitialTimestamp(&_IndexPriceOracle.TransactOpts, _timestamp)
}

// SetInitialTimestamp is a paid mutator transaction binding the contract method 0x21f926db.
//
// Solidity: function setInitialTimestamp(uint256 _timestamp) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) SetInitialTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.SetInitialTimestamp(&_IndexPriceOracle.TransactOpts, _timestamp)
}

// SetObservationAdder is a paid mutator transaction binding the contract method 0x7b72a641.
//
// Solidity: function setObservationAdder(address _adder) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactor) SetObservationAdder(opts *bind.TransactOpts, _adder common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.contract.Transact(opts, "setObservationAdder", _adder)
}

// SetObservationAdder is a paid mutator transaction binding the contract method 0x7b72a641.
//
// Solidity: function setObservationAdder(address _adder) returns()
func (_IndexPriceOracle *IndexPriceOracleSession) SetObservationAdder(_adder common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.SetObservationAdder(&_IndexPriceOracle.TransactOpts, _adder)
}

// SetObservationAdder is a paid mutator transaction binding the contract method 0x7b72a641.
//
// Solidity: function setObservationAdder(address _adder) returns()
func (_IndexPriceOracle *IndexPriceOracleTransactorSession) SetObservationAdder(_adder common.Address) (*types.Transaction, error) {
	return _IndexPriceOracle.Contract.SetObservationAdder(&_IndexPriceOracle.TransactOpts, _adder)
}

// IndexPriceOracleAssetsAddedIterator is returned from FilterAssetsAdded and is used to iterate over the raw logs and unpacked data for AssetsAdded events raised by the IndexPriceOracle contract.
type IndexPriceOracleAssetsAddedIterator struct {
	Event *IndexPriceOracleAssetsAdded // Event containing the contract specifics and raw log

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
func (it *IndexPriceOracleAssetsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexPriceOracleAssetsAdded)
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
		it.Event = new(IndexPriceOracleAssetsAdded)
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
func (it *IndexPriceOracleAssetsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexPriceOracleAssetsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexPriceOracleAssetsAdded represents a AssetsAdded event raised by the IndexPriceOracle contract.
type IndexPriceOracleAssetsAdded struct {
	LastIndex        *big.Int
	Assets           []common.Address
	UnderlyingPrices []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAssetsAdded is a free log retrieval operation binding the contract event 0xd54746bed1984556cd87b76fc3078d2670182f150532649f82479b40b3d9ab2d.
//
// Solidity: event AssetsAdded(uint256 indexed lastIndex, address[] assets, uint256[] underlyingPrices)
func (_IndexPriceOracle *IndexPriceOracleFilterer) FilterAssetsAdded(opts *bind.FilterOpts, lastIndex []*big.Int) (*IndexPriceOracleAssetsAddedIterator, error) {

	var lastIndexRule []interface{}
	for _, lastIndexItem := range lastIndex {
		lastIndexRule = append(lastIndexRule, lastIndexItem)
	}

	logs, sub, err := _IndexPriceOracle.contract.FilterLogs(opts, "AssetsAdded", lastIndexRule)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleAssetsAddedIterator{contract: _IndexPriceOracle.contract, event: "AssetsAdded", logs: logs, sub: sub}, nil
}

// WatchAssetsAdded is a free log subscription operation binding the contract event 0xd54746bed1984556cd87b76fc3078d2670182f150532649f82479b40b3d9ab2d.
//
// Solidity: event AssetsAdded(uint256 indexed lastIndex, address[] assets, uint256[] underlyingPrices)
func (_IndexPriceOracle *IndexPriceOracleFilterer) WatchAssetsAdded(opts *bind.WatchOpts, sink chan<- *IndexPriceOracleAssetsAdded, lastIndex []*big.Int) (event.Subscription, error) {

	var lastIndexRule []interface{}
	for _, lastIndexItem := range lastIndex {
		lastIndexRule = append(lastIndexRule, lastIndexItem)
	}

	logs, sub, err := _IndexPriceOracle.contract.WatchLogs(opts, "AssetsAdded", lastIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexPriceOracleAssetsAdded)
				if err := _IndexPriceOracle.contract.UnpackLog(event, "AssetsAdded", log); err != nil {
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
func (_IndexPriceOracle *IndexPriceOracleFilterer) ParseAssetsAdded(log types.Log) (*IndexPriceOracleAssetsAdded, error) {
	event := new(IndexPriceOracleAssetsAdded)
	if err := _IndexPriceOracle.contract.UnpackLog(event, "AssetsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexPriceOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IndexPriceOracle contract.
type IndexPriceOracleInitializedIterator struct {
	Event *IndexPriceOracleInitialized // Event containing the contract specifics and raw log

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
func (it *IndexPriceOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexPriceOracleInitialized)
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
		it.Event = new(IndexPriceOracleInitialized)
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
func (it *IndexPriceOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexPriceOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexPriceOracleInitialized represents a Initialized event raised by the IndexPriceOracle contract.
type IndexPriceOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IndexPriceOracle *IndexPriceOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*IndexPriceOracleInitializedIterator, error) {

	logs, sub, err := _IndexPriceOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleInitializedIterator{contract: _IndexPriceOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IndexPriceOracle *IndexPriceOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IndexPriceOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _IndexPriceOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexPriceOracleInitialized)
				if err := _IndexPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IndexPriceOracle *IndexPriceOracleFilterer) ParseInitialized(log types.Log) (*IndexPriceOracleInitialized, error) {
	event := new(IndexPriceOracleInitialized)
	if err := _IndexPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexPriceOracleObservationAddedIterator is returned from FilterObservationAdded and is used to iterate over the raw logs and unpacked data for ObservationAdded events raised by the IndexPriceOracle contract.
type IndexPriceOracleObservationAddedIterator struct {
	Event *IndexPriceOracleObservationAdded // Event containing the contract specifics and raw log

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
func (it *IndexPriceOracleObservationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexPriceOracleObservationAdded)
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
		it.Event = new(IndexPriceOracleObservationAdded)
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
func (it *IndexPriceOracleObservationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexPriceOracleObservationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexPriceOracleObservationAdded represents a ObservationAdded event raised by the IndexPriceOracle contract.
type IndexPriceOracleObservationAdded struct {
	Index           []*big.Int
	UnderlyingPrice []*big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterObservationAdded is a free log retrieval operation binding the contract event 0x30ff7507500c6f8ce27c81a024f60f8d17c748cf31e268a17d9f1d3af70ec996.
//
// Solidity: event ObservationAdded(uint256[] index, uint256[] underlyingPrice, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleFilterer) FilterObservationAdded(opts *bind.FilterOpts) (*IndexPriceOracleObservationAddedIterator, error) {

	logs, sub, err := _IndexPriceOracle.contract.FilterLogs(opts, "ObservationAdded")
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleObservationAddedIterator{contract: _IndexPriceOracle.contract, event: "ObservationAdded", logs: logs, sub: sub}, nil
}

// WatchObservationAdded is a free log subscription operation binding the contract event 0x30ff7507500c6f8ce27c81a024f60f8d17c748cf31e268a17d9f1d3af70ec996.
//
// Solidity: event ObservationAdded(uint256[] index, uint256[] underlyingPrice, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleFilterer) WatchObservationAdded(opts *bind.WatchOpts, sink chan<- *IndexPriceOracleObservationAdded) (event.Subscription, error) {

	logs, sub, err := _IndexPriceOracle.contract.WatchLogs(opts, "ObservationAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexPriceOracleObservationAdded)
				if err := _IndexPriceOracle.contract.UnpackLog(event, "ObservationAdded", log); err != nil {
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

// ParseObservationAdded is a log parse operation binding the contract event 0x30ff7507500c6f8ce27c81a024f60f8d17c748cf31e268a17d9f1d3af70ec996.
//
// Solidity: event ObservationAdded(uint256[] index, uint256[] underlyingPrice, uint256 timestamp)
func (_IndexPriceOracle *IndexPriceOracleFilterer) ParseObservationAdded(log types.Log) (*IndexPriceOracleObservationAdded, error) {
	event := new(IndexPriceOracleObservationAdded)
	if err := _IndexPriceOracle.contract.UnpackLog(event, "ObservationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexPriceOracleObservationAdderSetIterator is returned from FilterObservationAdderSet and is used to iterate over the raw logs and unpacked data for ObservationAdderSet events raised by the IndexPriceOracle contract.
type IndexPriceOracleObservationAdderSetIterator struct {
	Event *IndexPriceOracleObservationAdderSet // Event containing the contract specifics and raw log

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
func (it *IndexPriceOracleObservationAdderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexPriceOracleObservationAdderSet)
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
		it.Event = new(IndexPriceOracleObservationAdderSet)
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
func (it *IndexPriceOracleObservationAdderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexPriceOracleObservationAdderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexPriceOracleObservationAdderSet represents a ObservationAdderSet event raised by the IndexPriceOracle contract.
type IndexPriceOracleObservationAdderSet struct {
	MatchingEngine common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterObservationAdderSet is a free log retrieval operation binding the contract event 0x98b318ceaa8936e25855a68d4223b61f7cd0daa0c56d82005ca97c3ecef0ea94.
//
// Solidity: event ObservationAdderSet(address indexed matchingEngine)
func (_IndexPriceOracle *IndexPriceOracleFilterer) FilterObservationAdderSet(opts *bind.FilterOpts, matchingEngine []common.Address) (*IndexPriceOracleObservationAdderSetIterator, error) {

	var matchingEngineRule []interface{}
	for _, matchingEngineItem := range matchingEngine {
		matchingEngineRule = append(matchingEngineRule, matchingEngineItem)
	}

	logs, sub, err := _IndexPriceOracle.contract.FilterLogs(opts, "ObservationAdderSet", matchingEngineRule)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleObservationAdderSetIterator{contract: _IndexPriceOracle.contract, event: "ObservationAdderSet", logs: logs, sub: sub}, nil
}

// WatchObservationAdderSet is a free log subscription operation binding the contract event 0x98b318ceaa8936e25855a68d4223b61f7cd0daa0c56d82005ca97c3ecef0ea94.
//
// Solidity: event ObservationAdderSet(address indexed matchingEngine)
func (_IndexPriceOracle *IndexPriceOracleFilterer) WatchObservationAdderSet(opts *bind.WatchOpts, sink chan<- *IndexPriceOracleObservationAdderSet, matchingEngine []common.Address) (event.Subscription, error) {

	var matchingEngineRule []interface{}
	for _, matchingEngineItem := range matchingEngine {
		matchingEngineRule = append(matchingEngineRule, matchingEngineItem)
	}

	logs, sub, err := _IndexPriceOracle.contract.WatchLogs(opts, "ObservationAdderSet", matchingEngineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexPriceOracleObservationAdderSet)
				if err := _IndexPriceOracle.contract.UnpackLog(event, "ObservationAdderSet", log); err != nil {
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
func (_IndexPriceOracle *IndexPriceOracleFilterer) ParseObservationAdderSet(log types.Log) (*IndexPriceOracleObservationAdderSet, error) {
	event := new(IndexPriceOracleObservationAdderSet)
	if err := _IndexPriceOracle.contract.UnpackLog(event, "ObservationAdderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexPriceOracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IndexPriceOracle contract.
type IndexPriceOracleRoleAdminChangedIterator struct {
	Event *IndexPriceOracleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IndexPriceOracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexPriceOracleRoleAdminChanged)
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
		it.Event = new(IndexPriceOracleRoleAdminChanged)
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
func (it *IndexPriceOracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexPriceOracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexPriceOracleRoleAdminChanged represents a RoleAdminChanged event raised by the IndexPriceOracle contract.
type IndexPriceOracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IndexPriceOracle *IndexPriceOracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IndexPriceOracleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _IndexPriceOracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleRoleAdminChangedIterator{contract: _IndexPriceOracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IndexPriceOracle *IndexPriceOracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IndexPriceOracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IndexPriceOracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexPriceOracleRoleAdminChanged)
				if err := _IndexPriceOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_IndexPriceOracle *IndexPriceOracleFilterer) ParseRoleAdminChanged(log types.Log) (*IndexPriceOracleRoleAdminChanged, error) {
	event := new(IndexPriceOracleRoleAdminChanged)
	if err := _IndexPriceOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexPriceOracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IndexPriceOracle contract.
type IndexPriceOracleRoleGrantedIterator struct {
	Event *IndexPriceOracleRoleGranted // Event containing the contract specifics and raw log

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
func (it *IndexPriceOracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexPriceOracleRoleGranted)
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
		it.Event = new(IndexPriceOracleRoleGranted)
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
func (it *IndexPriceOracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexPriceOracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexPriceOracleRoleGranted represents a RoleGranted event raised by the IndexPriceOracle contract.
type IndexPriceOracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IndexPriceOracle *IndexPriceOracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IndexPriceOracleRoleGrantedIterator, error) {

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

	logs, sub, err := _IndexPriceOracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleRoleGrantedIterator{contract: _IndexPriceOracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IndexPriceOracle *IndexPriceOracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IndexPriceOracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IndexPriceOracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexPriceOracleRoleGranted)
				if err := _IndexPriceOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_IndexPriceOracle *IndexPriceOracleFilterer) ParseRoleGranted(log types.Log) (*IndexPriceOracleRoleGranted, error) {
	event := new(IndexPriceOracleRoleGranted)
	if err := _IndexPriceOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexPriceOracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IndexPriceOracle contract.
type IndexPriceOracleRoleRevokedIterator struct {
	Event *IndexPriceOracleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IndexPriceOracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexPriceOracleRoleRevoked)
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
		it.Event = new(IndexPriceOracleRoleRevoked)
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
func (it *IndexPriceOracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexPriceOracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexPriceOracleRoleRevoked represents a RoleRevoked event raised by the IndexPriceOracle contract.
type IndexPriceOracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IndexPriceOracle *IndexPriceOracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IndexPriceOracleRoleRevokedIterator, error) {

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

	logs, sub, err := _IndexPriceOracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IndexPriceOracleRoleRevokedIterator{contract: _IndexPriceOracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IndexPriceOracle *IndexPriceOracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IndexPriceOracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IndexPriceOracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexPriceOracleRoleRevoked)
				if err := _IndexPriceOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_IndexPriceOracle *IndexPriceOracleFilterer) ParseRoleRevoked(log types.Log) (*IndexPriceOracleRoleRevoked, error) {
	event := new(IndexPriceOracleRoleRevoked)
	if err := _IndexPriceOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
