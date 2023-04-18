package db

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type MatchedStatus uint
type TransactionStatusType uint
type TransactionLogStatusType uint

const (
	MatchedStatusZero                  MatchedStatus = 1
	MatchedStatusInit                  MatchedStatus = 2
	MatchedStatusFullMatchFound        MatchedStatus = 3
	MatchedStatusPartialMatchFound     MatchedStatus = 4
	MatchedStatusValidated             MatchedStatus = 5
	MatchedStatusValidationConfirmed   MatchedStatus = 6
	MatchedStatusSentToContract        MatchedStatus = 7
	MatchedStatusPartialMatchConfirmed MatchedStatus = 8
	MatchedStatusFullMatchConfirmed    MatchedStatus = 9
	MatchedStatusSentFailed            MatchedStatus = 10
	MatchedStatusFailedConfirmed       MatchedStatus = 11
	MatchedStatusBlocked               MatchedStatus = 12
	Canceled                           MatchedStatus = 13
)

const (
	TransactionStatusTypeLost             TransactionStatusType = 1
	TransactionStatusTypeInit             TransactionStatusType = 2
	TransactionStatusTypeNotFound         TransactionStatusType = 3
	TransactionStatusTypePending          TransactionStatusType = 4
	TransactionStatusTypeSent             TransactionStatusType = 5
	TransactionStatusTypeSuccess          TransactionStatusType = 6
	TransactionStatusTypeSuccessConfirmed TransactionStatusType = 7
	TransactionStatusTypeSentFailed       TransactionStatusType = 8
	TransactionStatusTypeFailed           TransactionStatusType = 9
	TransactionStatusTypeFailedConfirmed  TransactionStatusType = 10
)

// LibAssetAsset is an auto generated low-level Go binding around an user-defined struct.
type LibAssetAsset struct {
	VirtualToken common.Address `json:"virtualToken"`
	Value        *big.Int       `json:"value"`
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

const (
	TransactionLogStatusTypeInit      TransactionLogStatusType = 1
	TransactionLogStatusTypePending   TransactionLogStatusType = 2
	TransactionLogStatusTypeConfirmed TransactionLogStatusType = 3
)

const ORDER = "0xf555eb98"                   // bytes4(keccack256(abi.encodePacked("Order")))
const STOP_LOSS_INDEX_PRICE = "0x835d5c1e"   // bytes4(keccak256(abi.encodePacked("StopLossIndexPrice")));
const STOP_LOSS_LAST_PRICE = "0xd9ed8042"    //bytes4(keccak256(abi.encodePacked("StopLossLastPrice")));
const STOP_LOSS_MARK_PRICE = "0xe144c7ec"    //bytes4(keccak256(abi.encodePacked("StopLossMarkPrice")));
const TAKE_PROFIT_INDEX_PRICE = "0x67393efa" //bytes4(keccak256(abi.encodePacked("TakeProfitIndexPrice")));
const TAKE_PROFIT_LAST_PRICE = "0xc7dc86f6"  //bytes4(keccak256(abi.encodePacked("TakeProfitLastPrice")));
const TAKE_PROFIT_MARK_PRICE = "0xb6d64e04"  //bytes4(keccak256(abi.encodePacked("TakeProfitMarkPrice")));
