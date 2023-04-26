package db_test

// import (
// 	"fmt"
// 	"github.com/volmexfinance/relayers/internal/testlib"
// 	"testing"

// 	"github.com/volmexfinance/relayers/relayer-srv/db"
// )

// func TestOrder(t *testing.T) {
// 	var order = db.Order{
// 		OrderType:    "0xf555eb98",
// 		Trader:       "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
// 		Deadline:     87654321987654,
// 		Assets:       []db.Assets{{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}, {VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}},
// 		Price:        1,
// 		Salt:         "44",
// 		TriggerPrice: "0",
// 		Sign:         "",
// 		IsShort:      false,
// 		Status:       db.MatchedStatusInit,
// 		ChainName:    "ETH",
// 	}

// 	var order2 = db.Order{
// 		OrderType:    "0xf555eb98",
// 		Trader:       "0x74bC67ed6948f0a4C387C353975F142Dc640537a",
// 		Deadline:     87654321987654,
// 		Assets:       []db.Assets{{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}, {VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}},
// 		Price:        1,
// 		Salt:         "45",
// 		TriggerPrice: "0",
// 		Sign:         "",
// 		IsShort:      false,
// 		Status:       db.MatchedStatusInit,
// 		ChainName:    "ETH",
// 	}

// 	testDb := testlib.NewTestDB(t)

// 	order.OrderID = db.CreateOrderID(order.Trader, order.Salt, order.ChainName)
// 	order2.OrderID = db.CreateOrderID(order2.Trader, order2.Salt, order2.ChainName)

// 	er := testDb.CreateOrder(&order)
// 	if er != nil {
// 		t.Errorf("Failed to create order%q", er)
// 	}

// 	er = testDb.CreateOrder(&order2)
// 	if er != nil {
// 		t.Errorf("Failed to create order%q", er)
// 	}

// 	var txnLog = db.TransactionLog{
// 		Traders:         []string{"", ""},
// 		NewLeftFill:     "1000000000",
// 		NewRightFill:    "200000000000",
// 		OrderID:         []string{order.OrderID, order2.OrderID},
// 		ChainName:       "ETH",
// 		TransactionHash: "sshshs",
// 	}

// 	var txnLog2 = db.TransactionLog{
// 		Traders:         []string{"", ""},
// 		NewLeftFill:     "3000000000",
// 		NewRightFill:    "400000000000",
// 		OrderID:         []string{order2.OrderID, order.OrderID},
// 		ChainName:       "ETH",
// 		TransactionHash: "sshshs",
// 	}

// 	err := testDb.CreateTransactionLog(&txnLog)
// 	if err != nil {
// 		t.Errorf("Failed to create order%q", er)
// 	}

// 	err = testDb.CreateTransactionLog(&txnLog2)
// 	if err != nil {
// 		t.Errorf("Failed to create order%q", er)
// 	}

// 	err = testDb.UpdateFillAndStatusByTxnLog([]*db.Order{&order, &order2}, db.MatchedStatusSentFailed)
// 	if err != nil {
// 		t.Errorf("Failed to create order%q", er)
// 	}

// 	// find order
// 	foundOrder, erFind := testDb.FindOrder(order.OrderID)
// 	if erFind != nil {
// 		t.Errorf("FindOrder failed %q", erFind)
// 	}

// 	foundOrder2, erFind := testDb.FindOrder(order2.OrderID)
// 	if erFind != nil {
// 		t.Errorf("FindOrder failed %q", erFind)
// 	}

// 	fmt.Println("Fills and Status", foundOrder.Status, foundOrder.Fills, foundOrder2.Status, foundOrder2.Fills)
// 	got := foundOrder.Fills
// 	want := "400000000000"
// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}

// 	got = foundOrder2.Fills
// 	want = "3000000000"
// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }
