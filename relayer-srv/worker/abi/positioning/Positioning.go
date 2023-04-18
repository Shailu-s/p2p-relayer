// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Positioning

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

// IPositioningRealizePnlParams is an auto generated low-level Go binding around an user-defined struct.
type IPositioningRealizePnlParams struct {
	Trader    common.Address
	BaseToken common.Address
	Base      *big.Int
	Quote     *big.Int
}

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

// PositioningMetaData contains all meta data concerning the Positioning contract.
var PositioningMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"defaultFeeReceiver\",\"type\":\"address\"}],\"name\":\"DefaultFeeReceiverChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingPayment\",\"type\":\"int256\"}],\"name\":\"FundingPaymentSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"}],\"name\":\"FundingPeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"markTwap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexTwap\",\"type\":\"uint256\"}],\"name\":\"FundingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexPriceOracle\",\"type\":\"address\"}],\"name\":\"IndexPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWhitelist\",\"type\":\"bool\"}],\"name\":\"LiquidatorWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"exchangedPositionSize\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"exchangedPositionNotional\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderIndexPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"name\":\"PositionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionNotional\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"PositionLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"positioningCallee\",\"type\":\"address\"}],\"name\":\"PositioningCalleeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"TrustedForwarderChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POSITIONING_ADMIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POSITIONING_CALLEE_ADMIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountBalance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getAccountValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"accountValue\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getAllPendingFundingPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"pendingFundingPayment\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fundingPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"getIndexPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"getLastFundingRate\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"lastFundingRate\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"getLiquidatablePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"getNextFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nextFunding\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderValidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"getPendingFundingPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"base\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"quote\",\"type\":\"int256\"}],\"internalType\":\"structIPositioning.RealizePnlParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getPnlToBeRealized\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositioning\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositioningConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"positioningConfigArg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultControllerArg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountBalanceArg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"matchingEngineArg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"markPriceArg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexPriceArg\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPriceIndex\",\"type\":\"uint256\"},{\"internalType\":\"address[2]\",\"name\":\"liquidators\",\"type\":\"address[2]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"isAccountLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLiquidatorWhitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isLiquidatorWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"positionSize\",\"type\":\"int256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"liquidateFullPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"makerMinSalt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"orderLeft\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatureLeft\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"orderType\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"makeAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"virtualToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"takeAsset\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"limitOrderTriggerPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isShort\",\"type\":\"bool\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"orderRight\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatureRight\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"liquidator\",\"type\":\"bytes\"}],\"name\":\"openPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDefaultFeeReceiver\",\"type\":\"address\"}],\"name\":\"setDefaultFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setFundingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"indexPriceOracle\",\"type\":\"address\"}],\"name\":\"setIndexPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketRegistryArg\",\"type\":\"address\"}],\"name\":\"setMarketRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"positioningArg\",\"type\":\"address\"}],\"name\":\"setPositioning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"settleAllFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"settleFunding\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"fundingPayment\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"globalTwPremiumGrowth\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleLiquidatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWhitelist\",\"type\":\"bool\"}],\"name\":\"whitelistLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PositioningABI is the input ABI used to generate the binding from.
// Deprecated: Use PositioningMetaData.ABI instead.
var PositioningABI = PositioningMetaData.ABI

// Positioning is an auto generated Go binding around an Ethereum contract.
type Positioning struct {
	PositioningCaller     // Read-only binding to the contract
	PositioningTransactor // Write-only binding to the contract
	PositioningFilterer   // Log filterer for contract events
}

// PositioningCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositioningCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositioningTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositioningTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositioningFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositioningFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositioningSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositioningSession struct {
	Contract     *Positioning      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PositioningCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositioningCallerSession struct {
	Contract *PositioningCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PositioningTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositioningTransactorSession struct {
	Contract     *PositioningTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PositioningRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositioningRaw struct {
	Contract *Positioning // Generic contract binding to access the raw methods on
}

// PositioningCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositioningCallerRaw struct {
	Contract *PositioningCaller // Generic read-only contract binding to access the raw methods on
}

// PositioningTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositioningTransactorRaw struct {
	Contract *PositioningTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositioning creates a new instance of Positioning, bound to a specific deployed contract.
func NewPositioning(address common.Address, backend bind.ContractBackend) (*Positioning, error) {
	contract, err := bindPositioning(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Positioning{PositioningCaller: PositioningCaller{contract: contract}, PositioningTransactor: PositioningTransactor{contract: contract}, PositioningFilterer: PositioningFilterer{contract: contract}}, nil
}

// NewPositioningCaller creates a new read-only instance of Positioning, bound to a specific deployed contract.
func NewPositioningCaller(address common.Address, caller bind.ContractCaller) (*PositioningCaller, error) {
	contract, err := bindPositioning(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositioningCaller{contract: contract}, nil
}

// NewPositioningTransactor creates a new write-only instance of Positioning, bound to a specific deployed contract.
func NewPositioningTransactor(address common.Address, transactor bind.ContractTransactor) (*PositioningTransactor, error) {
	contract, err := bindPositioning(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositioningTransactor{contract: contract}, nil
}

// NewPositioningFilterer creates a new log filterer instance of Positioning, bound to a specific deployed contract.
func NewPositioningFilterer(address common.Address, filterer bind.ContractFilterer) (*PositioningFilterer, error) {
	contract, err := bindPositioning(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositioningFilterer{contract: contract}, nil
}

// bindPositioning binds a generic wrapper to an already deployed contract.
func bindPositioning(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositioningMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Positioning *PositioningRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Positioning.Contract.PositioningCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Positioning *PositioningRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Positioning.Contract.PositioningTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Positioning *PositioningRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Positioning.Contract.PositioningTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Positioning *PositioningCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Positioning.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Positioning *PositioningTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Positioning.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Positioning *PositioningTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Positioning.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Positioning *PositioningCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Positioning *PositioningSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Positioning.Contract.DEFAULTADMINROLE(&_Positioning.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Positioning *PositioningCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Positioning.Contract.DEFAULTADMINROLE(&_Positioning.CallOpts)
}

// POSITIONINGADMIN is a free data retrieval call binding the contract method 0x5ced6587.
//
// Solidity: function POSITIONING_ADMIN() view returns(bytes32)
func (_Positioning *PositioningCaller) POSITIONINGADMIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "POSITIONING_ADMIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSITIONINGADMIN is a free data retrieval call binding the contract method 0x5ced6587.
//
// Solidity: function POSITIONING_ADMIN() view returns(bytes32)
func (_Positioning *PositioningSession) POSITIONINGADMIN() ([32]byte, error) {
	return _Positioning.Contract.POSITIONINGADMIN(&_Positioning.CallOpts)
}

// POSITIONINGADMIN is a free data retrieval call binding the contract method 0x5ced6587.
//
// Solidity: function POSITIONING_ADMIN() view returns(bytes32)
func (_Positioning *PositioningCallerSession) POSITIONINGADMIN() ([32]byte, error) {
	return _Positioning.Contract.POSITIONINGADMIN(&_Positioning.CallOpts)
}

// POSITIONINGCALLEEADMIN is a free data retrieval call binding the contract method 0x3c640d6b.
//
// Solidity: function POSITIONING_CALLEE_ADMIN() view returns(bytes32)
func (_Positioning *PositioningCaller) POSITIONINGCALLEEADMIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "POSITIONING_CALLEE_ADMIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSITIONINGCALLEEADMIN is a free data retrieval call binding the contract method 0x3c640d6b.
//
// Solidity: function POSITIONING_CALLEE_ADMIN() view returns(bytes32)
func (_Positioning *PositioningSession) POSITIONINGCALLEEADMIN() ([32]byte, error) {
	return _Positioning.Contract.POSITIONINGCALLEEADMIN(&_Positioning.CallOpts)
}

// POSITIONINGCALLEEADMIN is a free data retrieval call binding the contract method 0x3c640d6b.
//
// Solidity: function POSITIONING_CALLEE_ADMIN() view returns(bytes32)
func (_Positioning *PositioningCallerSession) POSITIONINGCALLEEADMIN() ([32]byte, error) {
	return _Positioning.Contract.POSITIONINGCALLEEADMIN(&_Positioning.CallOpts)
}

// DefaultFeeReceiver is a free data retrieval call binding the contract method 0x3abf6fd4.
//
// Solidity: function defaultFeeReceiver() view returns(address)
func (_Positioning *PositioningCaller) DefaultFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "defaultFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultFeeReceiver is a free data retrieval call binding the contract method 0x3abf6fd4.
//
// Solidity: function defaultFeeReceiver() view returns(address)
func (_Positioning *PositioningSession) DefaultFeeReceiver() (common.Address, error) {
	return _Positioning.Contract.DefaultFeeReceiver(&_Positioning.CallOpts)
}

// DefaultFeeReceiver is a free data retrieval call binding the contract method 0x3abf6fd4.
//
// Solidity: function defaultFeeReceiver() view returns(address)
func (_Positioning *PositioningCallerSession) DefaultFeeReceiver() (common.Address, error) {
	return _Positioning.Contract.DefaultFeeReceiver(&_Positioning.CallOpts)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(address)
func (_Positioning *PositioningCaller) GetAccountBalance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getAccountBalance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(address)
func (_Positioning *PositioningSession) GetAccountBalance() (common.Address, error) {
	return _Positioning.Contract.GetAccountBalance(&_Positioning.CallOpts)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(address)
func (_Positioning *PositioningCallerSession) GetAccountBalance() (common.Address, error) {
	return _Positioning.Contract.GetAccountBalance(&_Positioning.CallOpts)
}

// GetAccountValue is a free data retrieval call binding the contract method 0x5ae80951.
//
// Solidity: function getAccountValue(address trader) view returns(int256 accountValue)
func (_Positioning *PositioningCaller) GetAccountValue(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getAccountValue", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountValue is a free data retrieval call binding the contract method 0x5ae80951.
//
// Solidity: function getAccountValue(address trader) view returns(int256 accountValue)
func (_Positioning *PositioningSession) GetAccountValue(trader common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetAccountValue(&_Positioning.CallOpts, trader)
}

// GetAccountValue is a free data retrieval call binding the contract method 0x5ae80951.
//
// Solidity: function getAccountValue(address trader) view returns(int256 accountValue)
func (_Positioning *PositioningCallerSession) GetAccountValue(trader common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetAccountValue(&_Positioning.CallOpts, trader)
}

// GetAllPendingFundingPayment is a free data retrieval call binding the contract method 0x80f57f11.
//
// Solidity: function getAllPendingFundingPayment(address trader) view returns(int256 pendingFundingPayment)
func (_Positioning *PositioningCaller) GetAllPendingFundingPayment(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getAllPendingFundingPayment", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllPendingFundingPayment is a free data retrieval call binding the contract method 0x80f57f11.
//
// Solidity: function getAllPendingFundingPayment(address trader) view returns(int256 pendingFundingPayment)
func (_Positioning *PositioningSession) GetAllPendingFundingPayment(trader common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetAllPendingFundingPayment(&_Positioning.CallOpts, trader)
}

// GetAllPendingFundingPayment is a free data retrieval call binding the contract method 0x80f57f11.
//
// Solidity: function getAllPendingFundingPayment(address trader) view returns(int256 pendingFundingPayment)
func (_Positioning *PositioningCallerSession) GetAllPendingFundingPayment(trader common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetAllPendingFundingPayment(&_Positioning.CallOpts, trader)
}

// GetFundingPeriod is a free data retrieval call binding the contract method 0xb5e2b047.
//
// Solidity: function getFundingPeriod() view returns(uint256 fundingPeriod)
func (_Positioning *PositioningCaller) GetFundingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getFundingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingPeriod is a free data retrieval call binding the contract method 0xb5e2b047.
//
// Solidity: function getFundingPeriod() view returns(uint256 fundingPeriod)
func (_Positioning *PositioningSession) GetFundingPeriod() (*big.Int, error) {
	return _Positioning.Contract.GetFundingPeriod(&_Positioning.CallOpts)
}

// GetFundingPeriod is a free data retrieval call binding the contract method 0xb5e2b047.
//
// Solidity: function getFundingPeriod() view returns(uint256 fundingPeriod)
func (_Positioning *PositioningCallerSession) GetFundingPeriod() (*big.Int, error) {
	return _Positioning.Contract.GetFundingPeriod(&_Positioning.CallOpts)
}

// GetIndexPrice is a free data retrieval call binding the contract method 0xb263e010.
//
// Solidity: function getIndexPrice(address baseToken) view returns(uint256 indexPrice)
func (_Positioning *PositioningCaller) GetIndexPrice(opts *bind.CallOpts, baseToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getIndexPrice", baseToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexPrice is a free data retrieval call binding the contract method 0xb263e010.
//
// Solidity: function getIndexPrice(address baseToken) view returns(uint256 indexPrice)
func (_Positioning *PositioningSession) GetIndexPrice(baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetIndexPrice(&_Positioning.CallOpts, baseToken)
}

// GetIndexPrice is a free data retrieval call binding the contract method 0xb263e010.
//
// Solidity: function getIndexPrice(address baseToken) view returns(uint256 indexPrice)
func (_Positioning *PositioningCallerSession) GetIndexPrice(baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetIndexPrice(&_Positioning.CallOpts, baseToken)
}

// GetLastFundingRate is a free data retrieval call binding the contract method 0xed061b45.
//
// Solidity: function getLastFundingRate(address baseToken) view returns(int256 lastFundingRate)
func (_Positioning *PositioningCaller) GetLastFundingRate(opts *bind.CallOpts, baseToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getLastFundingRate", baseToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastFundingRate is a free data retrieval call binding the contract method 0xed061b45.
//
// Solidity: function getLastFundingRate(address baseToken) view returns(int256 lastFundingRate)
func (_Positioning *PositioningSession) GetLastFundingRate(baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetLastFundingRate(&_Positioning.CallOpts, baseToken)
}

// GetLastFundingRate is a free data retrieval call binding the contract method 0xed061b45.
//
// Solidity: function getLastFundingRate(address baseToken) view returns(int256 lastFundingRate)
func (_Positioning *PositioningCallerSession) GetLastFundingRate(baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetLastFundingRate(&_Positioning.CallOpts, baseToken)
}

// GetLiquidatablePosition is a free data retrieval call binding the contract method 0x6702a33f.
//
// Solidity: function getLiquidatablePosition(address trader, address baseToken) view returns(uint256)
func (_Positioning *PositioningCaller) GetLiquidatablePosition(opts *bind.CallOpts, trader common.Address, baseToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getLiquidatablePosition", trader, baseToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidatablePosition is a free data retrieval call binding the contract method 0x6702a33f.
//
// Solidity: function getLiquidatablePosition(address trader, address baseToken) view returns(uint256)
func (_Positioning *PositioningSession) GetLiquidatablePosition(trader common.Address, baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetLiquidatablePosition(&_Positioning.CallOpts, trader, baseToken)
}

// GetLiquidatablePosition is a free data retrieval call binding the contract method 0x6702a33f.
//
// Solidity: function getLiquidatablePosition(address trader, address baseToken) view returns(uint256)
func (_Positioning *PositioningCallerSession) GetLiquidatablePosition(trader common.Address, baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetLiquidatablePosition(&_Positioning.CallOpts, trader, baseToken)
}

// GetNextFunding is a free data retrieval call binding the contract method 0x794ead7a.
//
// Solidity: function getNextFunding(address baseToken) view returns(uint256 nextFunding)
func (_Positioning *PositioningCaller) GetNextFunding(opts *bind.CallOpts, baseToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getNextFunding", baseToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFunding is a free data retrieval call binding the contract method 0x794ead7a.
//
// Solidity: function getNextFunding(address baseToken) view returns(uint256 nextFunding)
func (_Positioning *PositioningSession) GetNextFunding(baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetNextFunding(&_Positioning.CallOpts, baseToken)
}

// GetNextFunding is a free data retrieval call binding the contract method 0x794ead7a.
//
// Solidity: function getNextFunding(address baseToken) view returns(uint256 nextFunding)
func (_Positioning *PositioningCallerSession) GetNextFunding(baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetNextFunding(&_Positioning.CallOpts, baseToken)
}

// GetOrderValidate is a free data retrieval call binding the contract method 0x8462f997.
//
// Solidity: function getOrderValidate((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) order) view returns(bool)
func (_Positioning *PositioningCaller) GetOrderValidate(opts *bind.CallOpts, order LibOrderOrder) (bool, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getOrderValidate", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOrderValidate is a free data retrieval call binding the contract method 0x8462f997.
//
// Solidity: function getOrderValidate((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) order) view returns(bool)
func (_Positioning *PositioningSession) GetOrderValidate(order LibOrderOrder) (bool, error) {
	return _Positioning.Contract.GetOrderValidate(&_Positioning.CallOpts, order)
}

// GetOrderValidate is a free data retrieval call binding the contract method 0x8462f997.
//
// Solidity: function getOrderValidate((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) order) view returns(bool)
func (_Positioning *PositioningCallerSession) GetOrderValidate(order LibOrderOrder) (bool, error) {
	return _Positioning.Contract.GetOrderValidate(&_Positioning.CallOpts, order)
}

// GetPendingFundingPayment is a free data retrieval call binding the contract method 0xcb379aa2.
//
// Solidity: function getPendingFundingPayment(address trader, address baseToken) view returns(int256)
func (_Positioning *PositioningCaller) GetPendingFundingPayment(opts *bind.CallOpts, trader common.Address, baseToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getPendingFundingPayment", trader, baseToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingFundingPayment is a free data retrieval call binding the contract method 0xcb379aa2.
//
// Solidity: function getPendingFundingPayment(address trader, address baseToken) view returns(int256)
func (_Positioning *PositioningSession) GetPendingFundingPayment(trader common.Address, baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetPendingFundingPayment(&_Positioning.CallOpts, trader, baseToken)
}

// GetPendingFundingPayment is a free data retrieval call binding the contract method 0xcb379aa2.
//
// Solidity: function getPendingFundingPayment(address trader, address baseToken) view returns(int256)
func (_Positioning *PositioningCallerSession) GetPendingFundingPayment(trader common.Address, baseToken common.Address) (*big.Int, error) {
	return _Positioning.Contract.GetPendingFundingPayment(&_Positioning.CallOpts, trader, baseToken)
}

// GetPnlToBeRealized is a free data retrieval call binding the contract method 0x1393a469.
//
// Solidity: function getPnlToBeRealized((address,address,int256,int256) params) view returns(int256)
func (_Positioning *PositioningCaller) GetPnlToBeRealized(opts *bind.CallOpts, params IPositioningRealizePnlParams) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getPnlToBeRealized", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPnlToBeRealized is a free data retrieval call binding the contract method 0x1393a469.
//
// Solidity: function getPnlToBeRealized((address,address,int256,int256) params) view returns(int256)
func (_Positioning *PositioningSession) GetPnlToBeRealized(params IPositioningRealizePnlParams) (*big.Int, error) {
	return _Positioning.Contract.GetPnlToBeRealized(&_Positioning.CallOpts, params)
}

// GetPnlToBeRealized is a free data retrieval call binding the contract method 0x1393a469.
//
// Solidity: function getPnlToBeRealized((address,address,int256,int256) params) view returns(int256)
func (_Positioning *PositioningCallerSession) GetPnlToBeRealized(params IPositioningRealizePnlParams) (*big.Int, error) {
	return _Positioning.Contract.GetPnlToBeRealized(&_Positioning.CallOpts, params)
}

// GetPositioning is a free data retrieval call binding the contract method 0x561838b2.
//
// Solidity: function getPositioning() view returns(address)
func (_Positioning *PositioningCaller) GetPositioning(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getPositioning")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPositioning is a free data retrieval call binding the contract method 0x561838b2.
//
// Solidity: function getPositioning() view returns(address)
func (_Positioning *PositioningSession) GetPositioning() (common.Address, error) {
	return _Positioning.Contract.GetPositioning(&_Positioning.CallOpts)
}

// GetPositioning is a free data retrieval call binding the contract method 0x561838b2.
//
// Solidity: function getPositioning() view returns(address)
func (_Positioning *PositioningCallerSession) GetPositioning() (common.Address, error) {
	return _Positioning.Contract.GetPositioning(&_Positioning.CallOpts)
}

// GetPositioningConfig is a free data retrieval call binding the contract method 0x73c9439a.
//
// Solidity: function getPositioningConfig() view returns(address)
func (_Positioning *PositioningCaller) GetPositioningConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getPositioningConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPositioningConfig is a free data retrieval call binding the contract method 0x73c9439a.
//
// Solidity: function getPositioningConfig() view returns(address)
func (_Positioning *PositioningSession) GetPositioningConfig() (common.Address, error) {
	return _Positioning.Contract.GetPositioningConfig(&_Positioning.CallOpts)
}

// GetPositioningConfig is a free data retrieval call binding the contract method 0x73c9439a.
//
// Solidity: function getPositioningConfig() view returns(address)
func (_Positioning *PositioningCallerSession) GetPositioningConfig() (common.Address, error) {
	return _Positioning.Contract.GetPositioningConfig(&_Positioning.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Positioning *PositioningCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Positioning *PositioningSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Positioning.Contract.GetRoleAdmin(&_Positioning.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Positioning *PositioningCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Positioning.Contract.GetRoleAdmin(&_Positioning.CallOpts, role)
}

// GetVaultController is a free data retrieval call binding the contract method 0x77520b02.
//
// Solidity: function getVaultController() view returns(address)
func (_Positioning *PositioningCaller) GetVaultController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "getVaultController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultController is a free data retrieval call binding the contract method 0x77520b02.
//
// Solidity: function getVaultController() view returns(address)
func (_Positioning *PositioningSession) GetVaultController() (common.Address, error) {
	return _Positioning.Contract.GetVaultController(&_Positioning.CallOpts)
}

// GetVaultController is a free data retrieval call binding the contract method 0x77520b02.
//
// Solidity: function getVaultController() view returns(address)
func (_Positioning *PositioningCallerSession) GetVaultController() (common.Address, error) {
	return _Positioning.Contract.GetVaultController(&_Positioning.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Positioning *PositioningCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Positioning *PositioningSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Positioning.Contract.HasRole(&_Positioning.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Positioning *PositioningCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Positioning.Contract.HasRole(&_Positioning.CallOpts, role, account)
}

// IsAccountLiquidatable is a free data retrieval call binding the contract method 0xbe2801b5.
//
// Solidity: function isAccountLiquidatable(address trader) view returns(bool)
func (_Positioning *PositioningCaller) IsAccountLiquidatable(opts *bind.CallOpts, trader common.Address) (bool, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "isAccountLiquidatable", trader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountLiquidatable is a free data retrieval call binding the contract method 0xbe2801b5.
//
// Solidity: function isAccountLiquidatable(address trader) view returns(bool)
func (_Positioning *PositioningSession) IsAccountLiquidatable(trader common.Address) (bool, error) {
	return _Positioning.Contract.IsAccountLiquidatable(&_Positioning.CallOpts, trader)
}

// IsAccountLiquidatable is a free data retrieval call binding the contract method 0xbe2801b5.
//
// Solidity: function isAccountLiquidatable(address trader) view returns(bool)
func (_Positioning *PositioningCallerSession) IsAccountLiquidatable(trader common.Address) (bool, error) {
	return _Positioning.Contract.IsAccountLiquidatable(&_Positioning.CallOpts, trader)
}

// IsLiquidatorWhitelistEnabled is a free data retrieval call binding the contract method 0x7810d854.
//
// Solidity: function isLiquidatorWhitelistEnabled() view returns(bool)
func (_Positioning *PositioningCaller) IsLiquidatorWhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "isLiquidatorWhitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatorWhitelistEnabled is a free data retrieval call binding the contract method 0x7810d854.
//
// Solidity: function isLiquidatorWhitelistEnabled() view returns(bool)
func (_Positioning *PositioningSession) IsLiquidatorWhitelistEnabled() (bool, error) {
	return _Positioning.Contract.IsLiquidatorWhitelistEnabled(&_Positioning.CallOpts)
}

// IsLiquidatorWhitelistEnabled is a free data retrieval call binding the contract method 0x7810d854.
//
// Solidity: function isLiquidatorWhitelistEnabled() view returns(bool)
func (_Positioning *PositioningCallerSession) IsLiquidatorWhitelistEnabled() (bool, error) {
	return _Positioning.Contract.IsLiquidatorWhitelistEnabled(&_Positioning.CallOpts)
}

// IsLiquidatorWhitelisted is a free data retrieval call binding the contract method 0xe319fccc.
//
// Solidity: function isLiquidatorWhitelisted(address ) view returns(bool)
func (_Positioning *PositioningCaller) IsLiquidatorWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "isLiquidatorWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatorWhitelisted is a free data retrieval call binding the contract method 0xe319fccc.
//
// Solidity: function isLiquidatorWhitelisted(address ) view returns(bool)
func (_Positioning *PositioningSession) IsLiquidatorWhitelisted(arg0 common.Address) (bool, error) {
	return _Positioning.Contract.IsLiquidatorWhitelisted(&_Positioning.CallOpts, arg0)
}

// IsLiquidatorWhitelisted is a free data retrieval call binding the contract method 0xe319fccc.
//
// Solidity: function isLiquidatorWhitelisted(address ) view returns(bool)
func (_Positioning *PositioningCallerSession) IsLiquidatorWhitelisted(arg0 common.Address) (bool, error) {
	return _Positioning.Contract.IsLiquidatorWhitelisted(&_Positioning.CallOpts, arg0)
}

// MakerMinSalt is a free data retrieval call binding the contract method 0x624d7600.
//
// Solidity: function makerMinSalt(address ) view returns(uint256)
func (_Positioning *PositioningCaller) MakerMinSalt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "makerMinSalt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MakerMinSalt is a free data retrieval call binding the contract method 0x624d7600.
//
// Solidity: function makerMinSalt(address ) view returns(uint256)
func (_Positioning *PositioningSession) MakerMinSalt(arg0 common.Address) (*big.Int, error) {
	return _Positioning.Contract.MakerMinSalt(&_Positioning.CallOpts, arg0)
}

// MakerMinSalt is a free data retrieval call binding the contract method 0x624d7600.
//
// Solidity: function makerMinSalt(address ) view returns(uint256)
func (_Positioning *PositioningCallerSession) MakerMinSalt(arg0 common.Address) (*big.Int, error) {
	return _Positioning.Contract.MakerMinSalt(&_Positioning.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Positioning *PositioningCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Positioning *PositioningSession) Owner() (common.Address, error) {
	return _Positioning.Contract.Owner(&_Positioning.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Positioning *PositioningCallerSession) Owner() (common.Address, error) {
	return _Positioning.Contract.Owner(&_Positioning.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Positioning *PositioningCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Positioning *PositioningSession) Paused() (bool, error) {
	return _Positioning.Contract.Paused(&_Positioning.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Positioning *PositioningCallerSession) Paused() (bool, error) {
	return _Positioning.Contract.Paused(&_Positioning.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Positioning *PositioningCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Positioning.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Positioning *PositioningSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Positioning.Contract.SupportsInterface(&_Positioning.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Positioning *PositioningCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Positioning.Contract.SupportsInterface(&_Positioning.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Positioning *PositioningTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Positioning *PositioningSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.GrantRole(&_Positioning.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Positioning *PositioningTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.GrantRole(&_Positioning.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xbf3cc7cc.
//
// Solidity: function initialize(address positioningConfigArg, address vaultControllerArg, address accountBalanceArg, address matchingEngineArg, address markPriceArg, address indexPriceArg, uint256 underlyingPriceIndex, address[2] liquidators) returns()
func (_Positioning *PositioningTransactor) Initialize(opts *bind.TransactOpts, positioningConfigArg common.Address, vaultControllerArg common.Address, accountBalanceArg common.Address, matchingEngineArg common.Address, markPriceArg common.Address, indexPriceArg common.Address, underlyingPriceIndex *big.Int, liquidators [2]common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "initialize", positioningConfigArg, vaultControllerArg, accountBalanceArg, matchingEngineArg, markPriceArg, indexPriceArg, underlyingPriceIndex, liquidators)
}

// Initialize is a paid mutator transaction binding the contract method 0xbf3cc7cc.
//
// Solidity: function initialize(address positioningConfigArg, address vaultControllerArg, address accountBalanceArg, address matchingEngineArg, address markPriceArg, address indexPriceArg, uint256 underlyingPriceIndex, address[2] liquidators) returns()
func (_Positioning *PositioningSession) Initialize(positioningConfigArg common.Address, vaultControllerArg common.Address, accountBalanceArg common.Address, matchingEngineArg common.Address, markPriceArg common.Address, indexPriceArg common.Address, underlyingPriceIndex *big.Int, liquidators [2]common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.Initialize(&_Positioning.TransactOpts, positioningConfigArg, vaultControllerArg, accountBalanceArg, matchingEngineArg, markPriceArg, indexPriceArg, underlyingPriceIndex, liquidators)
}

// Initialize is a paid mutator transaction binding the contract method 0xbf3cc7cc.
//
// Solidity: function initialize(address positioningConfigArg, address vaultControllerArg, address accountBalanceArg, address matchingEngineArg, address markPriceArg, address indexPriceArg, uint256 underlyingPriceIndex, address[2] liquidators) returns()
func (_Positioning *PositioningTransactorSession) Initialize(positioningConfigArg common.Address, vaultControllerArg common.Address, accountBalanceArg common.Address, matchingEngineArg common.Address, markPriceArg common.Address, indexPriceArg common.Address, underlyingPriceIndex *big.Int, liquidators [2]common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.Initialize(&_Positioning.TransactOpts, positioningConfigArg, vaultControllerArg, accountBalanceArg, matchingEngineArg, markPriceArg, indexPriceArg, underlyingPriceIndex, liquidators)
}

// Liquidate is a paid mutator transaction binding the contract method 0x65d461bd.
//
// Solidity: function liquidate(address trader, address baseToken, int256 positionSize) returns()
func (_Positioning *PositioningTransactor) Liquidate(opts *bind.TransactOpts, trader common.Address, baseToken common.Address, positionSize *big.Int) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "liquidate", trader, baseToken, positionSize)
}

// Liquidate is a paid mutator transaction binding the contract method 0x65d461bd.
//
// Solidity: function liquidate(address trader, address baseToken, int256 positionSize) returns()
func (_Positioning *PositioningSession) Liquidate(trader common.Address, baseToken common.Address, positionSize *big.Int) (*types.Transaction, error) {
	return _Positioning.Contract.Liquidate(&_Positioning.TransactOpts, trader, baseToken, positionSize)
}

// Liquidate is a paid mutator transaction binding the contract method 0x65d461bd.
//
// Solidity: function liquidate(address trader, address baseToken, int256 positionSize) returns()
func (_Positioning *PositioningTransactorSession) Liquidate(trader common.Address, baseToken common.Address, positionSize *big.Int) (*types.Transaction, error) {
	return _Positioning.Contract.Liquidate(&_Positioning.TransactOpts, trader, baseToken, positionSize)
}

// LiquidateFullPosition is a paid mutator transaction binding the contract method 0x57958824.
//
// Solidity: function liquidateFullPosition(address trader, address baseToken) returns()
func (_Positioning *PositioningTransactor) LiquidateFullPosition(opts *bind.TransactOpts, trader common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "liquidateFullPosition", trader, baseToken)
}

// LiquidateFullPosition is a paid mutator transaction binding the contract method 0x57958824.
//
// Solidity: function liquidateFullPosition(address trader, address baseToken) returns()
func (_Positioning *PositioningSession) LiquidateFullPosition(trader common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.LiquidateFullPosition(&_Positioning.TransactOpts, trader, baseToken)
}

// LiquidateFullPosition is a paid mutator transaction binding the contract method 0x57958824.
//
// Solidity: function liquidateFullPosition(address trader, address baseToken) returns()
func (_Positioning *PositioningTransactorSession) LiquidateFullPosition(trader common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.LiquidateFullPosition(&_Positioning.TransactOpts, trader, baseToken)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x18001816.
//
// Solidity: function openPosition((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderLeft, bytes signatureLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderRight, bytes signatureRight, bytes liquidator) returns()
func (_Positioning *PositioningTransactor) OpenPosition(opts *bind.TransactOpts, orderLeft LibOrderOrder, signatureLeft []byte, orderRight LibOrderOrder, signatureRight []byte, liquidator []byte) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "openPosition", orderLeft, signatureLeft, orderRight, signatureRight, liquidator)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x18001816.
//
// Solidity: function openPosition((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderLeft, bytes signatureLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderRight, bytes signatureRight, bytes liquidator) returns()
func (_Positioning *PositioningSession) OpenPosition(orderLeft LibOrderOrder, signatureLeft []byte, orderRight LibOrderOrder, signatureRight []byte, liquidator []byte) (*types.Transaction, error) {
	return _Positioning.Contract.OpenPosition(&_Positioning.TransactOpts, orderLeft, signatureLeft, orderRight, signatureRight, liquidator)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x18001816.
//
// Solidity: function openPosition((bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderLeft, bytes signatureLeft, (bytes4,uint64,address,(address,uint256),(address,uint256),uint256,uint128,bool) orderRight, bytes signatureRight, bytes liquidator) returns()
func (_Positioning *PositioningTransactorSession) OpenPosition(orderLeft LibOrderOrder, signatureLeft []byte, orderRight LibOrderOrder, signatureRight []byte, liquidator []byte) (*types.Transaction, error) {
	return _Positioning.Contract.OpenPosition(&_Positioning.TransactOpts, orderLeft, signatureLeft, orderRight, signatureRight, liquidator)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Positioning *PositioningTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Positioning *PositioningSession) Pause() (*types.Transaction, error) {
	return _Positioning.Contract.Pause(&_Positioning.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Positioning *PositioningTransactorSession) Pause() (*types.Transaction, error) {
	return _Positioning.Contract.Pause(&_Positioning.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Positioning *PositioningTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Positioning *PositioningSession) RenounceOwnership() (*types.Transaction, error) {
	return _Positioning.Contract.RenounceOwnership(&_Positioning.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Positioning *PositioningTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Positioning.Contract.RenounceOwnership(&_Positioning.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Positioning *PositioningTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Positioning *PositioningSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.RenounceRole(&_Positioning.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Positioning *PositioningTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.RenounceRole(&_Positioning.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Positioning *PositioningTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Positioning *PositioningSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.RevokeRole(&_Positioning.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Positioning *PositioningTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.RevokeRole(&_Positioning.TransactOpts, role, account)
}

// SetDefaultFeeReceiver is a paid mutator transaction binding the contract method 0x1cdfe3d8.
//
// Solidity: function setDefaultFeeReceiver(address newDefaultFeeReceiver) returns()
func (_Positioning *PositioningTransactor) SetDefaultFeeReceiver(opts *bind.TransactOpts, newDefaultFeeReceiver common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "setDefaultFeeReceiver", newDefaultFeeReceiver)
}

// SetDefaultFeeReceiver is a paid mutator transaction binding the contract method 0x1cdfe3d8.
//
// Solidity: function setDefaultFeeReceiver(address newDefaultFeeReceiver) returns()
func (_Positioning *PositioningSession) SetDefaultFeeReceiver(newDefaultFeeReceiver common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SetDefaultFeeReceiver(&_Positioning.TransactOpts, newDefaultFeeReceiver)
}

// SetDefaultFeeReceiver is a paid mutator transaction binding the contract method 0x1cdfe3d8.
//
// Solidity: function setDefaultFeeReceiver(address newDefaultFeeReceiver) returns()
func (_Positioning *PositioningTransactorSession) SetDefaultFeeReceiver(newDefaultFeeReceiver common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SetDefaultFeeReceiver(&_Positioning.TransactOpts, newDefaultFeeReceiver)
}

// SetFundingPeriod is a paid mutator transaction binding the contract method 0x2a5aa292.
//
// Solidity: function setFundingPeriod(uint256 period) returns()
func (_Positioning *PositioningTransactor) SetFundingPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "setFundingPeriod", period)
}

// SetFundingPeriod is a paid mutator transaction binding the contract method 0x2a5aa292.
//
// Solidity: function setFundingPeriod(uint256 period) returns()
func (_Positioning *PositioningSession) SetFundingPeriod(period *big.Int) (*types.Transaction, error) {
	return _Positioning.Contract.SetFundingPeriod(&_Positioning.TransactOpts, period)
}

// SetFundingPeriod is a paid mutator transaction binding the contract method 0x2a5aa292.
//
// Solidity: function setFundingPeriod(uint256 period) returns()
func (_Positioning *PositioningTransactorSession) SetFundingPeriod(period *big.Int) (*types.Transaction, error) {
	return _Positioning.Contract.SetFundingPeriod(&_Positioning.TransactOpts, period)
}

// SetIndexPriceOracle is a paid mutator transaction binding the contract method 0x30a81a21.
//
// Solidity: function setIndexPriceOracle(address indexPriceOracle) returns()
func (_Positioning *PositioningTransactor) SetIndexPriceOracle(opts *bind.TransactOpts, indexPriceOracle common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "setIndexPriceOracle", indexPriceOracle)
}

// SetIndexPriceOracle is a paid mutator transaction binding the contract method 0x30a81a21.
//
// Solidity: function setIndexPriceOracle(address indexPriceOracle) returns()
func (_Positioning *PositioningSession) SetIndexPriceOracle(indexPriceOracle common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SetIndexPriceOracle(&_Positioning.TransactOpts, indexPriceOracle)
}

// SetIndexPriceOracle is a paid mutator transaction binding the contract method 0x30a81a21.
//
// Solidity: function setIndexPriceOracle(address indexPriceOracle) returns()
func (_Positioning *PositioningTransactorSession) SetIndexPriceOracle(indexPriceOracle common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SetIndexPriceOracle(&_Positioning.TransactOpts, indexPriceOracle)
}

// SetMarketRegistry is a paid mutator transaction binding the contract method 0xd8579704.
//
// Solidity: function setMarketRegistry(address marketRegistryArg) returns()
func (_Positioning *PositioningTransactor) SetMarketRegistry(opts *bind.TransactOpts, marketRegistryArg common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "setMarketRegistry", marketRegistryArg)
}

// SetMarketRegistry is a paid mutator transaction binding the contract method 0xd8579704.
//
// Solidity: function setMarketRegistry(address marketRegistryArg) returns()
func (_Positioning *PositioningSession) SetMarketRegistry(marketRegistryArg common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SetMarketRegistry(&_Positioning.TransactOpts, marketRegistryArg)
}

// SetMarketRegistry is a paid mutator transaction binding the contract method 0xd8579704.
//
// Solidity: function setMarketRegistry(address marketRegistryArg) returns()
func (_Positioning *PositioningTransactorSession) SetMarketRegistry(marketRegistryArg common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SetMarketRegistry(&_Positioning.TransactOpts, marketRegistryArg)
}

// SetPositioning is a paid mutator transaction binding the contract method 0xfa70d29e.
//
// Solidity: function setPositioning(address positioningArg) returns()
func (_Positioning *PositioningTransactor) SetPositioning(opts *bind.TransactOpts, positioningArg common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "setPositioning", positioningArg)
}

// SetPositioning is a paid mutator transaction binding the contract method 0xfa70d29e.
//
// Solidity: function setPositioning(address positioningArg) returns()
func (_Positioning *PositioningSession) SetPositioning(positioningArg common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SetPositioning(&_Positioning.TransactOpts, positioningArg)
}

// SetPositioning is a paid mutator transaction binding the contract method 0xfa70d29e.
//
// Solidity: function setPositioning(address positioningArg) returns()
func (_Positioning *PositioningTransactorSession) SetPositioning(positioningArg common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SetPositioning(&_Positioning.TransactOpts, positioningArg)
}

// SettleAllFunding is a paid mutator transaction binding the contract method 0xeb9b912e.
//
// Solidity: function settleAllFunding(address trader) returns()
func (_Positioning *PositioningTransactor) SettleAllFunding(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "settleAllFunding", trader)
}

// SettleAllFunding is a paid mutator transaction binding the contract method 0xeb9b912e.
//
// Solidity: function settleAllFunding(address trader) returns()
func (_Positioning *PositioningSession) SettleAllFunding(trader common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SettleAllFunding(&_Positioning.TransactOpts, trader)
}

// SettleAllFunding is a paid mutator transaction binding the contract method 0xeb9b912e.
//
// Solidity: function settleAllFunding(address trader) returns()
func (_Positioning *PositioningTransactorSession) SettleAllFunding(trader common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SettleAllFunding(&_Positioning.TransactOpts, trader)
}

// SettleFunding is a paid mutator transaction binding the contract method 0x1e81ac33.
//
// Solidity: function settleFunding(address trader, address baseToken) returns(int256 fundingPayment, int256 globalTwPremiumGrowth)
func (_Positioning *PositioningTransactor) SettleFunding(opts *bind.TransactOpts, trader common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "settleFunding", trader, baseToken)
}

// SettleFunding is a paid mutator transaction binding the contract method 0x1e81ac33.
//
// Solidity: function settleFunding(address trader, address baseToken) returns(int256 fundingPayment, int256 globalTwPremiumGrowth)
func (_Positioning *PositioningSession) SettleFunding(trader common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SettleFunding(&_Positioning.TransactOpts, trader, baseToken)
}

// SettleFunding is a paid mutator transaction binding the contract method 0x1e81ac33.
//
// Solidity: function settleFunding(address trader, address baseToken) returns(int256 fundingPayment, int256 globalTwPremiumGrowth)
func (_Positioning *PositioningTransactorSession) SettleFunding(trader common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.SettleFunding(&_Positioning.TransactOpts, trader, baseToken)
}

// ToggleLiquidatorWhitelist is a paid mutator transaction binding the contract method 0x34e97093.
//
// Solidity: function toggleLiquidatorWhitelist() returns()
func (_Positioning *PositioningTransactor) ToggleLiquidatorWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "toggleLiquidatorWhitelist")
}

// ToggleLiquidatorWhitelist is a paid mutator transaction binding the contract method 0x34e97093.
//
// Solidity: function toggleLiquidatorWhitelist() returns()
func (_Positioning *PositioningSession) ToggleLiquidatorWhitelist() (*types.Transaction, error) {
	return _Positioning.Contract.ToggleLiquidatorWhitelist(&_Positioning.TransactOpts)
}

// ToggleLiquidatorWhitelist is a paid mutator transaction binding the contract method 0x34e97093.
//
// Solidity: function toggleLiquidatorWhitelist() returns()
func (_Positioning *PositioningTransactorSession) ToggleLiquidatorWhitelist() (*types.Transaction, error) {
	return _Positioning.Contract.ToggleLiquidatorWhitelist(&_Positioning.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Positioning *PositioningTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Positioning *PositioningSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.TransferOwnership(&_Positioning.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Positioning *PositioningTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Positioning.Contract.TransferOwnership(&_Positioning.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Positioning *PositioningTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Positioning *PositioningSession) Unpause() (*types.Transaction, error) {
	return _Positioning.Contract.Unpause(&_Positioning.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Positioning *PositioningTransactorSession) Unpause() (*types.Transaction, error) {
	return _Positioning.Contract.Unpause(&_Positioning.TransactOpts)
}

// WhitelistLiquidator is a paid mutator transaction binding the contract method 0x2bf9e786.
//
// Solidity: function whitelistLiquidator(address liquidator, bool isWhitelist) returns()
func (_Positioning *PositioningTransactor) WhitelistLiquidator(opts *bind.TransactOpts, liquidator common.Address, isWhitelist bool) (*types.Transaction, error) {
	return _Positioning.contract.Transact(opts, "whitelistLiquidator", liquidator, isWhitelist)
}

// WhitelistLiquidator is a paid mutator transaction binding the contract method 0x2bf9e786.
//
// Solidity: function whitelistLiquidator(address liquidator, bool isWhitelist) returns()
func (_Positioning *PositioningSession) WhitelistLiquidator(liquidator common.Address, isWhitelist bool) (*types.Transaction, error) {
	return _Positioning.Contract.WhitelistLiquidator(&_Positioning.TransactOpts, liquidator, isWhitelist)
}

// WhitelistLiquidator is a paid mutator transaction binding the contract method 0x2bf9e786.
//
// Solidity: function whitelistLiquidator(address liquidator, bool isWhitelist) returns()
func (_Positioning *PositioningTransactorSession) WhitelistLiquidator(liquidator common.Address, isWhitelist bool) (*types.Transaction, error) {
	return _Positioning.Contract.WhitelistLiquidator(&_Positioning.TransactOpts, liquidator, isWhitelist)
}

// PositioningDefaultFeeReceiverChangedIterator is returned from FilterDefaultFeeReceiverChanged and is used to iterate over the raw logs and unpacked data for DefaultFeeReceiverChanged events raised by the Positioning contract.
type PositioningDefaultFeeReceiverChangedIterator struct {
	Event *PositioningDefaultFeeReceiverChanged // Event containing the contract specifics and raw log

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
func (it *PositioningDefaultFeeReceiverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningDefaultFeeReceiverChanged)
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
		it.Event = new(PositioningDefaultFeeReceiverChanged)
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
func (it *PositioningDefaultFeeReceiverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningDefaultFeeReceiverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningDefaultFeeReceiverChanged represents a DefaultFeeReceiverChanged event raised by the Positioning contract.
type PositioningDefaultFeeReceiverChanged struct {
	DefaultFeeReceiver common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDefaultFeeReceiverChanged is a free log retrieval operation binding the contract event 0x7329c1c48aea91720917e2650906f276d3f08e189bd646d117f71cce9fa9f389.
//
// Solidity: event DefaultFeeReceiverChanged(address defaultFeeReceiver)
func (_Positioning *PositioningFilterer) FilterDefaultFeeReceiverChanged(opts *bind.FilterOpts) (*PositioningDefaultFeeReceiverChangedIterator, error) {

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "DefaultFeeReceiverChanged")
	if err != nil {
		return nil, err
	}
	return &PositioningDefaultFeeReceiverChangedIterator{contract: _Positioning.contract, event: "DefaultFeeReceiverChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultFeeReceiverChanged is a free log subscription operation binding the contract event 0x7329c1c48aea91720917e2650906f276d3f08e189bd646d117f71cce9fa9f389.
//
// Solidity: event DefaultFeeReceiverChanged(address defaultFeeReceiver)
func (_Positioning *PositioningFilterer) WatchDefaultFeeReceiverChanged(opts *bind.WatchOpts, sink chan<- *PositioningDefaultFeeReceiverChanged) (event.Subscription, error) {

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "DefaultFeeReceiverChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningDefaultFeeReceiverChanged)
				if err := _Positioning.contract.UnpackLog(event, "DefaultFeeReceiverChanged", log); err != nil {
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

// ParseDefaultFeeReceiverChanged is a log parse operation binding the contract event 0x7329c1c48aea91720917e2650906f276d3f08e189bd646d117f71cce9fa9f389.
//
// Solidity: event DefaultFeeReceiverChanged(address defaultFeeReceiver)
func (_Positioning *PositioningFilterer) ParseDefaultFeeReceiverChanged(log types.Log) (*PositioningDefaultFeeReceiverChanged, error) {
	event := new(PositioningDefaultFeeReceiverChanged)
	if err := _Positioning.contract.UnpackLog(event, "DefaultFeeReceiverChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningFundingPaymentSettledIterator is returned from FilterFundingPaymentSettled and is used to iterate over the raw logs and unpacked data for FundingPaymentSettled events raised by the Positioning contract.
type PositioningFundingPaymentSettledIterator struct {
	Event *PositioningFundingPaymentSettled // Event containing the contract specifics and raw log

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
func (it *PositioningFundingPaymentSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningFundingPaymentSettled)
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
		it.Event = new(PositioningFundingPaymentSettled)
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
func (it *PositioningFundingPaymentSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningFundingPaymentSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningFundingPaymentSettled represents a FundingPaymentSettled event raised by the Positioning contract.
type PositioningFundingPaymentSettled struct {
	Trader         common.Address
	BaseToken      common.Address
	FundingPayment *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundingPaymentSettled is a free log retrieval operation binding the contract event 0x733330d4aad1a878654bf888817b79bc6478013399be29fa3b8845c81305249e.
//
// Solidity: event FundingPaymentSettled(address indexed trader, address indexed baseToken, int256 fundingPayment)
func (_Positioning *PositioningFilterer) FilterFundingPaymentSettled(opts *bind.FilterOpts, trader []common.Address, baseToken []common.Address) (*PositioningFundingPaymentSettledIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "FundingPaymentSettled", traderRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &PositioningFundingPaymentSettledIterator{contract: _Positioning.contract, event: "FundingPaymentSettled", logs: logs, sub: sub}, nil
}

// WatchFundingPaymentSettled is a free log subscription operation binding the contract event 0x733330d4aad1a878654bf888817b79bc6478013399be29fa3b8845c81305249e.
//
// Solidity: event FundingPaymentSettled(address indexed trader, address indexed baseToken, int256 fundingPayment)
func (_Positioning *PositioningFilterer) WatchFundingPaymentSettled(opts *bind.WatchOpts, sink chan<- *PositioningFundingPaymentSettled, trader []common.Address, baseToken []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "FundingPaymentSettled", traderRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningFundingPaymentSettled)
				if err := _Positioning.contract.UnpackLog(event, "FundingPaymentSettled", log); err != nil {
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

// ParseFundingPaymentSettled is a log parse operation binding the contract event 0x733330d4aad1a878654bf888817b79bc6478013399be29fa3b8845c81305249e.
//
// Solidity: event FundingPaymentSettled(address indexed trader, address indexed baseToken, int256 fundingPayment)
func (_Positioning *PositioningFilterer) ParseFundingPaymentSettled(log types.Log) (*PositioningFundingPaymentSettled, error) {
	event := new(PositioningFundingPaymentSettled)
	if err := _Positioning.contract.UnpackLog(event, "FundingPaymentSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningFundingPeriodSetIterator is returned from FilterFundingPeriodSet and is used to iterate over the raw logs and unpacked data for FundingPeriodSet events raised by the Positioning contract.
type PositioningFundingPeriodSetIterator struct {
	Event *PositioningFundingPeriodSet // Event containing the contract specifics and raw log

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
func (it *PositioningFundingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningFundingPeriodSet)
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
		it.Event = new(PositioningFundingPeriodSet)
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
func (it *PositioningFundingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningFundingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningFundingPeriodSet represents a FundingPeriodSet event raised by the Positioning contract.
type PositioningFundingPeriodSet struct {
	FundingInterval *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFundingPeriodSet is a free log retrieval operation binding the contract event 0xb80da7d4baf212353dabcadf9e435b9371f56b8e9e58408066239128dd596381.
//
// Solidity: event FundingPeriodSet(uint256 fundingInterval)
func (_Positioning *PositioningFilterer) FilterFundingPeriodSet(opts *bind.FilterOpts) (*PositioningFundingPeriodSetIterator, error) {

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "FundingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &PositioningFundingPeriodSetIterator{contract: _Positioning.contract, event: "FundingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchFundingPeriodSet is a free log subscription operation binding the contract event 0xb80da7d4baf212353dabcadf9e435b9371f56b8e9e58408066239128dd596381.
//
// Solidity: event FundingPeriodSet(uint256 fundingInterval)
func (_Positioning *PositioningFilterer) WatchFundingPeriodSet(opts *bind.WatchOpts, sink chan<- *PositioningFundingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "FundingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningFundingPeriodSet)
				if err := _Positioning.contract.UnpackLog(event, "FundingPeriodSet", log); err != nil {
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

// ParseFundingPeriodSet is a log parse operation binding the contract event 0xb80da7d4baf212353dabcadf9e435b9371f56b8e9e58408066239128dd596381.
//
// Solidity: event FundingPeriodSet(uint256 fundingInterval)
func (_Positioning *PositioningFilterer) ParseFundingPeriodSet(log types.Log) (*PositioningFundingPeriodSet, error) {
	event := new(PositioningFundingPeriodSet)
	if err := _Positioning.contract.UnpackLog(event, "FundingPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningFundingUpdatedIterator is returned from FilterFundingUpdated and is used to iterate over the raw logs and unpacked data for FundingUpdated events raised by the Positioning contract.
type PositioningFundingUpdatedIterator struct {
	Event *PositioningFundingUpdated // Event containing the contract specifics and raw log

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
func (it *PositioningFundingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningFundingUpdated)
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
		it.Event = new(PositioningFundingUpdated)
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
func (it *PositioningFundingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningFundingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningFundingUpdated represents a FundingUpdated event raised by the Positioning contract.
type PositioningFundingUpdated struct {
	BaseToken common.Address
	MarkTwap  *big.Int
	IndexTwap *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundingUpdated is a free log retrieval operation binding the contract event 0x54e4482fe1d38392effe5d53f0e9e72f60221a75a10cea7abbb684bfb03519bf.
//
// Solidity: event FundingUpdated(address indexed baseToken, uint256 markTwap, uint256 indexTwap)
func (_Positioning *PositioningFilterer) FilterFundingUpdated(opts *bind.FilterOpts, baseToken []common.Address) (*PositioningFundingUpdatedIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "FundingUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &PositioningFundingUpdatedIterator{contract: _Positioning.contract, event: "FundingUpdated", logs: logs, sub: sub}, nil
}

// WatchFundingUpdated is a free log subscription operation binding the contract event 0x54e4482fe1d38392effe5d53f0e9e72f60221a75a10cea7abbb684bfb03519bf.
//
// Solidity: event FundingUpdated(address indexed baseToken, uint256 markTwap, uint256 indexTwap)
func (_Positioning *PositioningFilterer) WatchFundingUpdated(opts *bind.WatchOpts, sink chan<- *PositioningFundingUpdated, baseToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "FundingUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningFundingUpdated)
				if err := _Positioning.contract.UnpackLog(event, "FundingUpdated", log); err != nil {
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

// ParseFundingUpdated is a log parse operation binding the contract event 0x54e4482fe1d38392effe5d53f0e9e72f60221a75a10cea7abbb684bfb03519bf.
//
// Solidity: event FundingUpdated(address indexed baseToken, uint256 markTwap, uint256 indexTwap)
func (_Positioning *PositioningFilterer) ParseFundingUpdated(log types.Log) (*PositioningFundingUpdated, error) {
	event := new(PositioningFundingUpdated)
	if err := _Positioning.contract.UnpackLog(event, "FundingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningIndexPriceSetIterator is returned from FilterIndexPriceSet and is used to iterate over the raw logs and unpacked data for IndexPriceSet events raised by the Positioning contract.
type PositioningIndexPriceSetIterator struct {
	Event *PositioningIndexPriceSet // Event containing the contract specifics and raw log

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
func (it *PositioningIndexPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningIndexPriceSet)
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
		it.Event = new(PositioningIndexPriceSet)
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
func (it *PositioningIndexPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningIndexPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningIndexPriceSet represents a IndexPriceSet event raised by the Positioning contract.
type PositioningIndexPriceSet struct {
	IndexPriceOracle common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIndexPriceSet is a free log retrieval operation binding the contract event 0x6d07c0605722f31411d7e5480cf1e07b76d8812ee79f3ab05082e589008816f0.
//
// Solidity: event IndexPriceSet(address indexed indexPriceOracle)
func (_Positioning *PositioningFilterer) FilterIndexPriceSet(opts *bind.FilterOpts, indexPriceOracle []common.Address) (*PositioningIndexPriceSetIterator, error) {

	var indexPriceOracleRule []interface{}
	for _, indexPriceOracleItem := range indexPriceOracle {
		indexPriceOracleRule = append(indexPriceOracleRule, indexPriceOracleItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "IndexPriceSet", indexPriceOracleRule)
	if err != nil {
		return nil, err
	}
	return &PositioningIndexPriceSetIterator{contract: _Positioning.contract, event: "IndexPriceSet", logs: logs, sub: sub}, nil
}

// WatchIndexPriceSet is a free log subscription operation binding the contract event 0x6d07c0605722f31411d7e5480cf1e07b76d8812ee79f3ab05082e589008816f0.
//
// Solidity: event IndexPriceSet(address indexed indexPriceOracle)
func (_Positioning *PositioningFilterer) WatchIndexPriceSet(opts *bind.WatchOpts, sink chan<- *PositioningIndexPriceSet, indexPriceOracle []common.Address) (event.Subscription, error) {

	var indexPriceOracleRule []interface{}
	for _, indexPriceOracleItem := range indexPriceOracle {
		indexPriceOracleRule = append(indexPriceOracleRule, indexPriceOracleItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "IndexPriceSet", indexPriceOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningIndexPriceSet)
				if err := _Positioning.contract.UnpackLog(event, "IndexPriceSet", log); err != nil {
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

// ParseIndexPriceSet is a log parse operation binding the contract event 0x6d07c0605722f31411d7e5480cf1e07b76d8812ee79f3ab05082e589008816f0.
//
// Solidity: event IndexPriceSet(address indexed indexPriceOracle)
func (_Positioning *PositioningFilterer) ParseIndexPriceSet(log types.Log) (*PositioningIndexPriceSet, error) {
	event := new(PositioningIndexPriceSet)
	if err := _Positioning.contract.UnpackLog(event, "IndexPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Positioning contract.
type PositioningInitializedIterator struct {
	Event *PositioningInitialized // Event containing the contract specifics and raw log

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
func (it *PositioningInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningInitialized)
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
		it.Event = new(PositioningInitialized)
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
func (it *PositioningInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningInitialized represents a Initialized event raised by the Positioning contract.
type PositioningInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Positioning *PositioningFilterer) FilterInitialized(opts *bind.FilterOpts) (*PositioningInitializedIterator, error) {

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PositioningInitializedIterator{contract: _Positioning.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Positioning *PositioningFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PositioningInitialized) (event.Subscription, error) {

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningInitialized)
				if err := _Positioning.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Positioning *PositioningFilterer) ParseInitialized(log types.Log) (*PositioningInitialized, error) {
	event := new(PositioningInitialized)
	if err := _Positioning.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningLiquidatorWhitelistedIterator is returned from FilterLiquidatorWhitelisted and is used to iterate over the raw logs and unpacked data for LiquidatorWhitelisted events raised by the Positioning contract.
type PositioningLiquidatorWhitelistedIterator struct {
	Event *PositioningLiquidatorWhitelisted // Event containing the contract specifics and raw log

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
func (it *PositioningLiquidatorWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningLiquidatorWhitelisted)
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
		it.Event = new(PositioningLiquidatorWhitelisted)
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
func (it *PositioningLiquidatorWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningLiquidatorWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningLiquidatorWhitelisted represents a LiquidatorWhitelisted event raised by the Positioning contract.
type PositioningLiquidatorWhitelisted struct {
	Liquidator  common.Address
	IsWhitelist bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLiquidatorWhitelisted is a free log retrieval operation binding the contract event 0xbbda333b832934c97c0f8d2dd25f2fb172ba2cb951d01ae173d6c458ebb61606.
//
// Solidity: event LiquidatorWhitelisted(address indexed liquidator, bool isWhitelist)
func (_Positioning *PositioningFilterer) FilterLiquidatorWhitelisted(opts *bind.FilterOpts, liquidator []common.Address) (*PositioningLiquidatorWhitelistedIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "LiquidatorWhitelisted", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &PositioningLiquidatorWhitelistedIterator{contract: _Positioning.contract, event: "LiquidatorWhitelisted", logs: logs, sub: sub}, nil
}

// WatchLiquidatorWhitelisted is a free log subscription operation binding the contract event 0xbbda333b832934c97c0f8d2dd25f2fb172ba2cb951d01ae173d6c458ebb61606.
//
// Solidity: event LiquidatorWhitelisted(address indexed liquidator, bool isWhitelist)
func (_Positioning *PositioningFilterer) WatchLiquidatorWhitelisted(opts *bind.WatchOpts, sink chan<- *PositioningLiquidatorWhitelisted, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "LiquidatorWhitelisted", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningLiquidatorWhitelisted)
				if err := _Positioning.contract.UnpackLog(event, "LiquidatorWhitelisted", log); err != nil {
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

// ParseLiquidatorWhitelisted is a log parse operation binding the contract event 0xbbda333b832934c97c0f8d2dd25f2fb172ba2cb951d01ae173d6c458ebb61606.
//
// Solidity: event LiquidatorWhitelisted(address indexed liquidator, bool isWhitelist)
func (_Positioning *PositioningFilterer) ParseLiquidatorWhitelisted(log types.Log) (*PositioningLiquidatorWhitelisted, error) {
	event := new(PositioningLiquidatorWhitelisted)
	if err := _Positioning.contract.UnpackLog(event, "LiquidatorWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Positioning contract.
type PositioningOwnershipTransferredIterator struct {
	Event *PositioningOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PositioningOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningOwnershipTransferred)
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
		it.Event = new(PositioningOwnershipTransferred)
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
func (it *PositioningOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningOwnershipTransferred represents a OwnershipTransferred event raised by the Positioning contract.
type PositioningOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Positioning *PositioningFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PositioningOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PositioningOwnershipTransferredIterator{contract: _Positioning.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Positioning *PositioningFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PositioningOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningOwnershipTransferred)
				if err := _Positioning.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Positioning *PositioningFilterer) ParseOwnershipTransferred(log types.Log) (*PositioningOwnershipTransferred, error) {
	event := new(PositioningOwnershipTransferred)
	if err := _Positioning.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Positioning contract.
type PositioningPausedIterator struct {
	Event *PositioningPaused // Event containing the contract specifics and raw log

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
func (it *PositioningPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningPaused)
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
		it.Event = new(PositioningPaused)
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
func (it *PositioningPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningPaused represents a Paused event raised by the Positioning contract.
type PositioningPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Positioning *PositioningFilterer) FilterPaused(opts *bind.FilterOpts) (*PositioningPausedIterator, error) {

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PositioningPausedIterator{contract: _Positioning.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Positioning *PositioningFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PositioningPaused) (event.Subscription, error) {

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningPaused)
				if err := _Positioning.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Positioning *PositioningFilterer) ParsePaused(log types.Log) (*PositioningPaused, error) {
	event := new(PositioningPaused)
	if err := _Positioning.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningPositionChangedIterator is returned from FilterPositionChanged and is used to iterate over the raw logs and unpacked data for PositionChanged events raised by the Positioning contract.
type PositioningPositionChangedIterator struct {
	Event *PositioningPositionChanged // Event containing the contract specifics and raw log

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
func (it *PositioningPositionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningPositionChanged)
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
		it.Event = new(PositioningPositionChanged)
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
func (it *PositioningPositionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningPositionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningPositionChanged represents a PositionChanged event raised by the Positioning contract.
type PositioningPositionChanged struct {
	Trader                    common.Address
	BaseToken                 common.Address
	ExchangedPositionSize     *big.Int
	ExchangedPositionNotional *big.Int
	Fee                       *big.Int
	OrderIndexPrice           *big.Int
	OrderType                 [4]byte
	IsShort                   bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPositionChanged is a free log retrieval operation binding the contract event 0xf37b987cb17b588f526af5efc02fc0f2191f6b6015836d54f3bf393aaa98076f.
//
// Solidity: event PositionChanged(address indexed trader, address indexed baseToken, int256 exchangedPositionSize, int256 exchangedPositionNotional, uint256 fee, uint256 orderIndexPrice, bytes4 orderType, bool isShort)
func (_Positioning *PositioningFilterer) FilterPositionChanged(opts *bind.FilterOpts, trader []common.Address, baseToken []common.Address) (*PositioningPositionChangedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "PositionChanged", traderRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &PositioningPositionChangedIterator{contract: _Positioning.contract, event: "PositionChanged", logs: logs, sub: sub}, nil
}

// WatchPositionChanged is a free log subscription operation binding the contract event 0xf37b987cb17b588f526af5efc02fc0f2191f6b6015836d54f3bf393aaa98076f.
//
// Solidity: event PositionChanged(address indexed trader, address indexed baseToken, int256 exchangedPositionSize, int256 exchangedPositionNotional, uint256 fee, uint256 orderIndexPrice, bytes4 orderType, bool isShort)
func (_Positioning *PositioningFilterer) WatchPositionChanged(opts *bind.WatchOpts, sink chan<- *PositioningPositionChanged, trader []common.Address, baseToken []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "PositionChanged", traderRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningPositionChanged)
				if err := _Positioning.contract.UnpackLog(event, "PositionChanged", log); err != nil {
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

// ParsePositionChanged is a log parse operation binding the contract event 0xf37b987cb17b588f526af5efc02fc0f2191f6b6015836d54f3bf393aaa98076f.
//
// Solidity: event PositionChanged(address indexed trader, address indexed baseToken, int256 exchangedPositionSize, int256 exchangedPositionNotional, uint256 fee, uint256 orderIndexPrice, bytes4 orderType, bool isShort)
func (_Positioning *PositioningFilterer) ParsePositionChanged(log types.Log) (*PositioningPositionChanged, error) {
	event := new(PositioningPositionChanged)
	if err := _Positioning.contract.UnpackLog(event, "PositionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningPositionLiquidatedIterator is returned from FilterPositionLiquidated and is used to iterate over the raw logs and unpacked data for PositionLiquidated events raised by the Positioning contract.
type PositioningPositionLiquidatedIterator struct {
	Event *PositioningPositionLiquidated // Event containing the contract specifics and raw log

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
func (it *PositioningPositionLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningPositionLiquidated)
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
		it.Event = new(PositioningPositionLiquidated)
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
func (it *PositioningPositionLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningPositionLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningPositionLiquidated represents a PositionLiquidated event raised by the Positioning contract.
type PositioningPositionLiquidated struct {
	Trader           common.Address
	BaseToken        common.Address
	PositionNotional *big.Int
	PositionSize     *big.Int
	LiquidationFee   *big.Int
	Liquidator       common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPositionLiquidated is a free log retrieval operation binding the contract event 0xd9aced30440caca81570436bc942f816cfd95a3f08f700a2aeb6334c7cb5b497.
//
// Solidity: event PositionLiquidated(address indexed trader, address indexed baseToken, uint256 positionNotional, uint256 positionSize, uint256 liquidationFee, address liquidator)
func (_Positioning *PositioningFilterer) FilterPositionLiquidated(opts *bind.FilterOpts, trader []common.Address, baseToken []common.Address) (*PositioningPositionLiquidatedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "PositionLiquidated", traderRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &PositioningPositionLiquidatedIterator{contract: _Positioning.contract, event: "PositionLiquidated", logs: logs, sub: sub}, nil
}

// WatchPositionLiquidated is a free log subscription operation binding the contract event 0xd9aced30440caca81570436bc942f816cfd95a3f08f700a2aeb6334c7cb5b497.
//
// Solidity: event PositionLiquidated(address indexed trader, address indexed baseToken, uint256 positionNotional, uint256 positionSize, uint256 liquidationFee, address liquidator)
func (_Positioning *PositioningFilterer) WatchPositionLiquidated(opts *bind.WatchOpts, sink chan<- *PositioningPositionLiquidated, trader []common.Address, baseToken []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "PositionLiquidated", traderRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningPositionLiquidated)
				if err := _Positioning.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
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

// ParsePositionLiquidated is a log parse operation binding the contract event 0xd9aced30440caca81570436bc942f816cfd95a3f08f700a2aeb6334c7cb5b497.
//
// Solidity: event PositionLiquidated(address indexed trader, address indexed baseToken, uint256 positionNotional, uint256 positionSize, uint256 liquidationFee, address liquidator)
func (_Positioning *PositioningFilterer) ParsePositionLiquidated(log types.Log) (*PositioningPositionLiquidated, error) {
	event := new(PositioningPositionLiquidated)
	if err := _Positioning.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningPositioningCalleeChangedIterator is returned from FilterPositioningCalleeChanged and is used to iterate over the raw logs and unpacked data for PositioningCalleeChanged events raised by the Positioning contract.
type PositioningPositioningCalleeChangedIterator struct {
	Event *PositioningPositioningCalleeChanged // Event containing the contract specifics and raw log

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
func (it *PositioningPositioningCalleeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningPositioningCalleeChanged)
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
		it.Event = new(PositioningPositioningCalleeChanged)
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
func (it *PositioningPositioningCalleeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningPositioningCalleeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningPositioningCalleeChanged represents a PositioningCalleeChanged event raised by the Positioning contract.
type PositioningPositioningCalleeChanged struct {
	PositioningCallee common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPositioningCalleeChanged is a free log retrieval operation binding the contract event 0x7e4d746f11e5ad965b922917270e2ac76abafff78d7f9fe274de3b4353ae7a1f.
//
// Solidity: event PositioningCalleeChanged(address indexed positioningCallee)
func (_Positioning *PositioningFilterer) FilterPositioningCalleeChanged(opts *bind.FilterOpts, positioningCallee []common.Address) (*PositioningPositioningCalleeChangedIterator, error) {

	var positioningCalleeRule []interface{}
	for _, positioningCalleeItem := range positioningCallee {
		positioningCalleeRule = append(positioningCalleeRule, positioningCalleeItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "PositioningCalleeChanged", positioningCalleeRule)
	if err != nil {
		return nil, err
	}
	return &PositioningPositioningCalleeChangedIterator{contract: _Positioning.contract, event: "PositioningCalleeChanged", logs: logs, sub: sub}, nil
}

// WatchPositioningCalleeChanged is a free log subscription operation binding the contract event 0x7e4d746f11e5ad965b922917270e2ac76abafff78d7f9fe274de3b4353ae7a1f.
//
// Solidity: event PositioningCalleeChanged(address indexed positioningCallee)
func (_Positioning *PositioningFilterer) WatchPositioningCalleeChanged(opts *bind.WatchOpts, sink chan<- *PositioningPositioningCalleeChanged, positioningCallee []common.Address) (event.Subscription, error) {

	var positioningCalleeRule []interface{}
	for _, positioningCalleeItem := range positioningCallee {
		positioningCalleeRule = append(positioningCalleeRule, positioningCalleeItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "PositioningCalleeChanged", positioningCalleeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningPositioningCalleeChanged)
				if err := _Positioning.contract.UnpackLog(event, "PositioningCalleeChanged", log); err != nil {
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

// ParsePositioningCalleeChanged is a log parse operation binding the contract event 0x7e4d746f11e5ad965b922917270e2ac76abafff78d7f9fe274de3b4353ae7a1f.
//
// Solidity: event PositioningCalleeChanged(address indexed positioningCallee)
func (_Positioning *PositioningFilterer) ParsePositioningCalleeChanged(log types.Log) (*PositioningPositioningCalleeChanged, error) {
	event := new(PositioningPositioningCalleeChanged)
	if err := _Positioning.contract.UnpackLog(event, "PositioningCalleeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Positioning contract.
type PositioningRoleAdminChangedIterator struct {
	Event *PositioningRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PositioningRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningRoleAdminChanged)
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
		it.Event = new(PositioningRoleAdminChanged)
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
func (it *PositioningRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningRoleAdminChanged represents a RoleAdminChanged event raised by the Positioning contract.
type PositioningRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Positioning *PositioningFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PositioningRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PositioningRoleAdminChangedIterator{contract: _Positioning.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Positioning *PositioningFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PositioningRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningRoleAdminChanged)
				if err := _Positioning.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Positioning *PositioningFilterer) ParseRoleAdminChanged(log types.Log) (*PositioningRoleAdminChanged, error) {
	event := new(PositioningRoleAdminChanged)
	if err := _Positioning.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Positioning contract.
type PositioningRoleGrantedIterator struct {
	Event *PositioningRoleGranted // Event containing the contract specifics and raw log

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
func (it *PositioningRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningRoleGranted)
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
		it.Event = new(PositioningRoleGranted)
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
func (it *PositioningRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningRoleGranted represents a RoleGranted event raised by the Positioning contract.
type PositioningRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Positioning *PositioningFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PositioningRoleGrantedIterator, error) {

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

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PositioningRoleGrantedIterator{contract: _Positioning.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Positioning *PositioningFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PositioningRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningRoleGranted)
				if err := _Positioning.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Positioning *PositioningFilterer) ParseRoleGranted(log types.Log) (*PositioningRoleGranted, error) {
	event := new(PositioningRoleGranted)
	if err := _Positioning.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Positioning contract.
type PositioningRoleRevokedIterator struct {
	Event *PositioningRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PositioningRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningRoleRevoked)
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
		it.Event = new(PositioningRoleRevoked)
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
func (it *PositioningRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningRoleRevoked represents a RoleRevoked event raised by the Positioning contract.
type PositioningRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Positioning *PositioningFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PositioningRoleRevokedIterator, error) {

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

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PositioningRoleRevokedIterator{contract: _Positioning.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Positioning *PositioningFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PositioningRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningRoleRevoked)
				if err := _Positioning.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Positioning *PositioningFilterer) ParseRoleRevoked(log types.Log) (*PositioningRoleRevoked, error) {
	event := new(PositioningRoleRevoked)
	if err := _Positioning.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningTrustedForwarderChangedIterator is returned from FilterTrustedForwarderChanged and is used to iterate over the raw logs and unpacked data for TrustedForwarderChanged events raised by the Positioning contract.
type PositioningTrustedForwarderChangedIterator struct {
	Event *PositioningTrustedForwarderChanged // Event containing the contract specifics and raw log

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
func (it *PositioningTrustedForwarderChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningTrustedForwarderChanged)
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
		it.Event = new(PositioningTrustedForwarderChanged)
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
func (it *PositioningTrustedForwarderChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningTrustedForwarderChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningTrustedForwarderChanged represents a TrustedForwarderChanged event raised by the Positioning contract.
type PositioningTrustedForwarderChanged struct {
	Forwarder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedForwarderChanged is a free log retrieval operation binding the contract event 0x871264f4293af7d2865ae7eae628b228f4991c57cb45b39c99f0b774ebe29018.
//
// Solidity: event TrustedForwarderChanged(address indexed forwarder)
func (_Positioning *PositioningFilterer) FilterTrustedForwarderChanged(opts *bind.FilterOpts, forwarder []common.Address) (*PositioningTrustedForwarderChangedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "TrustedForwarderChanged", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &PositioningTrustedForwarderChangedIterator{contract: _Positioning.contract, event: "TrustedForwarderChanged", logs: logs, sub: sub}, nil
}

// WatchTrustedForwarderChanged is a free log subscription operation binding the contract event 0x871264f4293af7d2865ae7eae628b228f4991c57cb45b39c99f0b774ebe29018.
//
// Solidity: event TrustedForwarderChanged(address indexed forwarder)
func (_Positioning *PositioningFilterer) WatchTrustedForwarderChanged(opts *bind.WatchOpts, sink chan<- *PositioningTrustedForwarderChanged, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "TrustedForwarderChanged", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningTrustedForwarderChanged)
				if err := _Positioning.contract.UnpackLog(event, "TrustedForwarderChanged", log); err != nil {
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

// ParseTrustedForwarderChanged is a log parse operation binding the contract event 0x871264f4293af7d2865ae7eae628b228f4991c57cb45b39c99f0b774ebe29018.
//
// Solidity: event TrustedForwarderChanged(address indexed forwarder)
func (_Positioning *PositioningFilterer) ParseTrustedForwarderChanged(log types.Log) (*PositioningTrustedForwarderChanged, error) {
	event := new(PositioningTrustedForwarderChanged)
	if err := _Positioning.contract.UnpackLog(event, "TrustedForwarderChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositioningUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Positioning contract.
type PositioningUnpausedIterator struct {
	Event *PositioningUnpaused // Event containing the contract specifics and raw log

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
func (it *PositioningUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositioningUnpaused)
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
		it.Event = new(PositioningUnpaused)
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
func (it *PositioningUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositioningUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositioningUnpaused represents a Unpaused event raised by the Positioning contract.
type PositioningUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Positioning *PositioningFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PositioningUnpausedIterator, error) {

	logs, sub, err := _Positioning.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PositioningUnpausedIterator{contract: _Positioning.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Positioning *PositioningFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PositioningUnpaused) (event.Subscription, error) {

	logs, sub, err := _Positioning.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositioningUnpaused)
				if err := _Positioning.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Positioning *PositioningFilterer) ParseUnpaused(log types.Log) (*PositioningUnpaused, error) {
	event := new(PositioningUnpaused)
	if err := _Positioning.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
