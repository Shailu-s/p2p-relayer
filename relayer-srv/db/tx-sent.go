package db

import (
	"fmt"
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"strings"
	"time"
)

// CreateTxSent creates a transaction sent in DB
func (db *DataBase) CreateTxSent(txn *models.TransactionSent, chainName string) error {
	txn.ChainName = chainName
	txn.CreatedAt = time.Now().Unix()
	if result := db.DB.Model(models.TransactionSent{}).Create(&txn); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// GetTxnByID finds and returns transaction sent with given ID from DB, error if not found
func (db *DataBase) GetTxnByID(ID int64, chain string) (*models.TransactionSent, error) {
	var txnSent models.TransactionSent
	if result := db.DB.Model(models.TransactionSent{}).Where("id = ? AND chain_name = ?", ID, chain).Find(&txnSent); result.Error != nil {
		return nil, result.Error
	} else {
		return &txnSent, nil
	}
}

// GetTxnByID finds and returns transaction sent with given order ID from DB, error if not found
func (db *DataBase) GetTxnByOrderID(orderID string) ([]models.TransactionSent, error) {
	var txnSent []models.TransactionSent
	if result := db.DB.Model(models.TransactionSent{}).Where("order_id LIKE ?", fmt.Sprintf("%%%s%%", orderID)).Find(&txnSent); result.Error != nil {
		return nil, result.Error
	} else {
		return txnSent, nil
	}
}

// UpdateTxnSent takes transaction sent and update it on DB
func (db *DataBase) UpdateTxnSent(updatedTxn *models.TransactionSent, chain string) (*models.TransactionSent, error) {
	updatedTxn.UpdatedAt = time.Now().Unix()
	if result := db.DB.Model(models.TransactionSent{}).Where("id = ? AND chain_name = ?", updatedTxn.ID, chain).Order("created_at desc").Updates(updatedTxn); result.Error != nil {
		return nil, result.Error
	} else {
		return updatedTxn, nil
	}
}

// FindLastTxnSent returns last transaction sent in DB
func (db *DataBase) FindLastTxnSent(chain string) (*models.TransactionSent, error) {
	var lastTx models.TransactionSent
	if result := db.DB.Model(models.TransactionSent{}).Where("chain_name = ?", chain).Order("id desc").Limit(1).First(&lastTx); result.Error != nil {
		return nil, result.Error
	}

	return &lastTx, nil
}

// FindSecondLastTxnSent returns second last transaction sent in DB
func (db *DataBase) FindSecondLastTxnSent(chain string) (*models.TransactionSent, error) {
	var lastTx []models.TransactionSent
	if result := db.DB.Model(models.TransactionSent{}).Where("chain_name = ?", chain).Order("id desc").Limit(2).Find(&lastTx); result.Error != nil {
		return nil, result.Error
	}
	if len(lastTx) > 1 {
		return &lastTx[1], nil
	}
	return nil, fmt.Errorf("Unable to find second last txn")
}

// HandleTxnSentByLeader creates or update given transaction sent on DB, only leader can use it
func (db *DataBase) HandleTxnSentByLeader(txn models.TransactionSent, leader string, chain string) (*models.TransactionSent, error) {
	updateTx, err := db.FindLastTxnSent(chain)
	if err != nil {
		if err := db.CreateTxSent(&txn, chain); err != nil {
			return nil, err
		}
		return &txn, nil
	} else if strings.EqualFold(updateTx.LeaderPublicKey, leader) {
		updateTx.TransactionHash = txn.TransactionHash
		updateTx.OrderID = txn.OrderID
		txn1, err := db.UpdateTxnSent(updateTx, chain)
		if err != nil {
			return nil, err
		}
		return txn1, nil
	} else {
		if err := db.CreateTxSent(&txn, chain); err != nil {
			return nil, err
		}
		return &txn, nil
	}
}

// UpdateTxnStatus updates transaction sent and its staus on DB
func (db *DataBase) UpdateTxnStatus(txn *models.TransactionSent, newStatus models.TransactionStatusType) error {
	txn.UpdatedAt = time.Now().Unix()
	txn.TransactionStatus = newStatus
	if result := db.DB.Save(&txn); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetTotalTransactionCount returns total transaction sent stored in DB
func (db *DataBase) GetTotalTransactionCount(chain string) (int64, error) {
	var txnSent []models.TransactionSent
	var count int64
	if result := db.DB.Model(&txnSent).Where("chain_name = ?", chain).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// GetTxnsOnStatus returns transaction sent stored on DB on given status
func (db *DataBase) GetTxnsOnStatus(transactionStatusType []models.TransactionStatusType, chain string) (*[]models.TransactionSent, error) {
	var txnSent []models.TransactionSent
	if result := db.DB.Where("transaction_status in (?) AND chain_name = ?", transactionStatusType, chain).Find(&txnSent); result.Error != nil {
		return nil, result.Error
	} else {
		return &txnSent, nil
	}
}
