package models

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestOrderModel_CreateOrderId(t *testing.T) {
	/// Given
	trader := "0x401f0b1c51a7048d3db9a8ca4e9a370e563e0fb9"
	salt := "44"
	chainName := "ARB"

	/// When
	orderId := CreateOrderID(trader, salt, chainName)

	/// Then
	assert.Equal(t, "0x401f0b1c51a7048d3db9a8ca4e9a370e563e0fb944ARB", orderId)
}

func TestOrderModel_CalculatePrice_BuyOrder(t *testing.T) {
	/// Given
	// A buy order (long order), giving 300 USDT in exchange for 600 EVIV
	makeAssetValue := big.NewFloat(300)
	takeAssetValue := big.NewFloat(600)

	/// When
	price := CalculatePrice(false, makeAssetValue, takeAssetValue)

	/// Then
	// Price units: USDT/EVIV
	assert.EqualValues(t, 0.5, price)
}

func TestOrderModel_CalculatePrice_SellOrder(t *testing.T) {
	/// Given
	// A sell order (short order), giving 100 EVIV in exchange for 200 USDT
	makeAssetValue := big.NewFloat(100)
	takeAssetValue := big.NewFloat(200)

	/// When
	price := CalculatePrice(true, makeAssetValue, takeAssetValue)

	/// Then
	// Price units: USDT/EVIV
	assert.EqualValues(t, 2, price)
}
