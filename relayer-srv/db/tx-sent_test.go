package db_test

import (
	"fmt"
	"github.com/volmexfinance/relayers/internal/testlib"
	"testing"

	"github.com/volmexfinance/relayers/relayer-srv/db"
)

func TestTxSent(t *testing.T) {
	testDb := testlib.NewTestDB(t)

	transactionHash := "0xd16d24998bcb3cd53c428f79596caf28dde19079bdea8b50746d155f7903c745" // TODO : Put txn hash here

	txSent := db.TransactionSent{
		ID:                1,
		TransactionHash:   transactionHash,
		TransactionStatus: db.TransactionStatusTypeInit,
		Error:             "",
	}

	// Create TxnSent
	er := testDb.CreateTxSent(&txSent, "ARB")
	if er != nil {
		t.Errorf("CreateTxnFailed:%q", er)
	}

	// Get TxnSnt by orderId
	txSentGotGetTxnByOrderID, erGetTxnByOrderId := testDb.GetTxnByOrderID("test123")
	if erGetTxnByOrderId != nil {
		t.Errorf("GetTxnByHash:%q", erGetTxnByOrderId)
	}

	fmt.Println(txSentGotGetTxnByOrderID)
}
