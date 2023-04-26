package models

import (
	"github.com/lib/pq"
)

type TransactionStatusType uint
type TransactionLogStatusType uint

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

const (
	TransactionLogStatusTypeInit      TransactionLogStatusType = 1
	TransactionLogStatusTypePending   TransactionLogStatusType = 2
	TransactionLogStatusTypeConfirmed TransactionLogStatusType = 3
)

// TODO: Move this table into postgresDB
type TransactionSent struct {
	ID                int                   `gorm:"primaryKey"`
	TransactionHash   string                `json:"transaction_hash"`
	OrderID           pq.StringArray        `gorm:"type:text[]"`
	TransactionStatus TransactionStatusType `gorm:"type:BIGINT"`
	Error             string                `json:"error"`
	CreatedAt         int64                 `json:"created_at"`
	UpdatedAt         int64                 `json:"updated_at"`
	Leader            string                `json:"leader"`
	LeaderPublicKey   string                `json:"leader_public_key"`
	ChainName         string                `json:"chain_name"`
}

type TransactionLog struct {
	ID              int                      `gorm:"primaryKey"`
	Traders         pq.StringArray           `json:"traders" gorm:"type:TEXT[]"`
	Deadline        pq.Int64Array            `json:"deadline" gorm:"type:BIGINT[]"`
	Salt            pq.StringArray           `json:"salt" gorm:"type:TEXT[]"`
	NewLeftFill     string                   `json:"new_left_fill" gorm:"type:TEXT"`
	NewRightFill    string                   `json:"new_right_fill" gorm:"type:TEXT"`
	CreatedAt       int64                    `json:"createdAt"`
	UpdatedAt       int64                    `json:"updatedAt"`
	OrderID         pq.StringArray           `json:"order_id" gorm:"type:TEXT[]"`
	Status          TransactionLogStatusType `json:"status" gorm:"type:BIGINT"`
	BlockHeight     uint64                   `json:"block_height" gorm:"type:BIGINT"`
	ChainName       string                   `json:"chain_name"`
	TransactionHash string                   `json:"transaction_hash" gorm:"type:TEXT"`
}
