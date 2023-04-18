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
