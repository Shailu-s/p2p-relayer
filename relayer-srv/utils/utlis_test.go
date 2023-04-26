package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	periphery "github.com/volmexfinance/relayers/relayer-srv/worker/abi/periphery"

	protocols_p2p "github.com/volmexfinance/relayers/relayer-srv/chat"
	"github.com/volmexfinance/relayers/relayer-srv/db"
)

func TestOrderToLibOrder(t *testing.T) {
	// Given
	newOrder := db.Order{
		OrderType: "0xf555eb98",
		Trader:    "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
		Deadline:  87654321987654,
		Assets: []db.Assets{
			{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"},
			{VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"},
		},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      false,
		Status:       db.MatchedStatusInit,
	}

	expectedType := reflect.TypeOf(periphery.LibOrderOrder{})

	// When
	libOrderGot := OrderToLibOrder(&newOrder)

	// Then
	assert.Equal(t, expectedType, reflect.TypeOf(libOrderGot))
}

func TestP2POrderToDBOrder(t *testing.T) {

	newOrder := protocols_p2p.Order{ // TODO : fill values
		// state:         "",
		// sizeCache:     "",
		// unknownFields: "",
		OrderType:              "0xf555eb98",
		Deadline:               87654321987654,
		Trader:                 "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
		MakeAsset:              &protocols_p2p.Asset{VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"},
		TakeAsset:              &protocols_p2p.Asset{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"},
		Salt:                   "44",
		LimitOrderTriggerPrice: "0",
		IsShort:                false,
		Sign:                   "",
		Fills:                  "",
	}

	DBOrderWant := &db.Order{}
	DBOrderGot, err := P2POrderToDBOrder(&newOrder, "")
	if err != nil {
		t.Errorf("got error %v on converting p2p order to db order", err)
	}
	fmt.Println(reflect.TypeOf(DBOrderGot))
	if reflect.TypeOf(DBOrderGot) != reflect.TypeOf(DBOrderWant) {
		t.Errorf("got %q, wanted db.Order", reflect.TypeOf(DBOrderGot))
	}
}

func TestEncodePositionContractData(t *testing.T) {

	order1 := db.Order{
		OrderType:    "0xf555eb98",
		Trader:       "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
		Deadline:     87654321987654,
		Assets:       []db.Assets{{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}, {VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      false,
		Status:       db.MatchedStatusInit,
	}
	order2 := db.Order{
		OrderType:    "0xf555eb98",
		Trader:       "0x77a18F00CE53c004337b4A8b7A9a585AFFDEeD5e",
		Deadline:     87654321987654,
		Assets:       []db.Assets{{VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}, {VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      true,
		Status:       db.MatchedStatusInit,
	}

	positionContractAdd := "" // TODO: put contract address
	_, er := EncodePeripheryContractData([]*db.Order{&order1}, []*db.Order{&order2}, positionContractAdd)
	if er != nil {
		t.Errorf("Failed to Encode order")
	}
}

func TestEncodeOrderStruct(t *testing.T) {
	// Given
	order1 := db.Order{
		OrderType: "0xf555eb98",
		Trader:    "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
		Deadline:  87654321987654,
		Assets: []db.Assets{
			{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"},
			{VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"},
		},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      false,
		Status:       db.MatchedStatusInit,
	}

	// Test vector from: https://eips.ethereum.org/EIPS/eip-712
	positionContractAdd := "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC"
	chainId := int64(4)

	// When
	_, er := EncodeOrderStruct(order1, chainId, positionContractAdd)

	// Then
	assert.NoErrorf(t, er, "Failed to Encode order struct")
}

func TestGetOrderIdsFromTLog(t *testing.T) {
	// Given

	tLog := db.TransactionLog{ // TODO : put values here
		Traders: []string{"0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9", "0x77a18F00CE53c004337b4A8b7A9a585AFFDEeD5e"},
		Salt:    []string{"44", "44"},
	}

	expectedOrderId1 := "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb944"
	expectedOrderId2 := "0x77a18F00CE53c004337b4A8b7A9a585AFFDEeD5e44"

	// When
	orderId1Got, orderId2Got := GetOrderIdsFromTLog(&tLog, "")

	// Then
	assert.Equal(t, expectedOrderId1, orderId1Got)
	assert.Equal(t, expectedOrderId2, orderId2Got)
}

// func TestCreateTransactionMsgP2P(t *testing.T) {
// 	txSent := db.TransactionSent{
// 		ID:                1,
// 		TransactionHash:   "",
// 		OrderID:           []string{},
// 		TransactionStatus: 0,
// 		Error:             "",
// 		CreatedAt:         1,
// 		UpdatedAt:         1,
// 		Nonce:             1,
// 	}
// 	_, er := CreateTransactionMsgP2P(&txSent)
// 	if er != nil {
// 		t.Errorf("failed to CreateTransactionMsgP2P%q", er)
// 	}
// }
