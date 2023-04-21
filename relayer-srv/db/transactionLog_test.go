package db

import (
	"testing"
)

// var tLog *TransactionLog

func (db *PostgresDataBase) TestCreateTransactionLog(t *testing.T) {
	tLog := TransactionLog{
		Traders:      []string{"0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9", "0x77a18F00CE53c004337b4A8b7A9a585AFFDEeD5e"},
		Deadline:     []int64{87654321987654, 87654321987654},
		Salt:         []string{"1", "2"},
		NewLeftFill:  "NewLeftFill",
		NewRightFill: "NewRightFill",
		OrderID:      []string{"", ""},
		Status:       TransactionLogStatusTypeInit,
	}
	err := db.CreateTransactionLog(&tLog)

	if err != nil {
		t.Errorf("TestCreateTransactionLog: failed with error %s", err.Error())
	}

	tLog.OrderID[0] = "123Test"
}
