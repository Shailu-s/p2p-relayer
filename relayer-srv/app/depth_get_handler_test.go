package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/volmexfinance/relayers/internal/testlib"
	"github.com/volmexfinance/relayers/relayer-srv/big_ext"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"testing"
)

func fakeDbAsset(token string, value uint64) testlib.FakeDbAssetFn {
	return testlib.FakeDbAsset(token, value)
}

func fakeAnyValueDbAsset(token string) testlib.FakeDbAssetFn {
	return testlib.FakeAnyValueDbAsset(token)
}

// fakeBuyOrder returns a fake buy order with the given price, make asset, take asset, and fills.
func fakeBuyOrder(price float64, makeAssetFn, takeAssetFn testlib.FakeDbAssetFn, fills int64) *db.Order {
	return testlib.FakeDbOrder(
		false,
		testlib.WithPrice(price),
		testlib.WithAssets(makeAssetFn, takeAssetFn),
		testlib.WithFills(fills),
	)
}

// fakeSellOrder returns a fake sell order with the given price, make asset, take asset, and fills.
func fakeSellOrder(price float64, makeAssetFn, takeAssetFn testlib.FakeDbAssetFn, fills int64) *db.Order {
	return testlib.FakeDbOrder(
		true,
		testlib.WithPrice(price),
		testlib.WithAssets(makeAssetFn, takeAssetFn),
		testlib.WithFills(fills),
	)
}

func TestAggregateOrderDepthDetails_bidOrders(t *testing.T) {
	/// Given
	// Base token = EVIV
	// Quote token = USDT
	orderDetails := []*db.Order{
		fakeBuyOrder(30.000, fakeDbAsset("USDT", 20), fakeAnyValueDbAsset("EVIV"), 0),
		fakeBuyOrder(30.001, fakeDbAsset("USDT", 40), fakeAnyValueDbAsset("EVIV"), 20),
		fakeBuyOrder(40.000, fakeDbAsset("USDT", 60), fakeAnyValueDbAsset("EVIV"), 20),
		fakeBuyOrder(40.001, fakeDbAsset("USDT", 80), fakeAnyValueDbAsset("EVIV"), 40),
	}

	/// When
	bidTokens, keyBidTokens, askTokens, keyAskTokens := aggregateOrderDepthDetails(orderDetails)

	/// Then
	assert.Equal(t, []float64{30, 40}, keyBidTokens) // USDT/EVIV
	assert.Len(t, bidTokens, 2)
	assert.EqualValues(t, 40, big_ext.ToFloat64(bidTokens[30])) // USDT
	assert.EqualValues(t, 80, big_ext.ToFloat64(bidTokens[40])) // USDT

	assert.Len(t, keyAskTokens, 0)
	assert.Len(t, askTokens, 0)
}

func TestAggregateOrderDepthDetails_askOrders(t *testing.T) {
	/// Given
	// Base token = EVIV
	// Quote token = USDT
	orderDetails := []*db.Order{
		fakeSellOrder(10.000, fakeDbAsset("EVIV", 10), fakeAnyValueDbAsset("USDT"), 0),
		fakeSellOrder(10.001, fakeDbAsset("EVIV", 30), fakeAnyValueDbAsset("USDT"), 20),
		fakeSellOrder(20.000, fakeDbAsset("EVIV", 60), fakeAnyValueDbAsset("USDT"), 20),
		fakeSellOrder(20.001, fakeDbAsset("EVIV", 80), fakeAnyValueDbAsset("USDT"), 40),

		// A very large number
		fakeSellOrder(130, fakeDbAsset("EVIV", 769230769230769230), fakeAnyValueDbAsset("USDT"), 0),
	}

	/// When
	bidTokens, keyBidTokens, askTokens, keyAskTokens := aggregateOrderDepthDetails(orderDetails)

	/// Then
	assert.Len(t, keyBidTokens, 0)
	assert.Len(t, bidTokens, 0)

	assert.Equal(t, []float64{10, 20, 130}, keyAskTokens) // USDT/EVIV
	assert.Len(t, askTokens, 3)
	assert.EqualValues(t, 200.01, big_ext.ToFloat64(askTokens[10]))  // USDT
	assert.EqualValues(t, 1600.04, big_ext.ToFloat64(askTokens[20])) // USDT

	assert.EqualValues(t, 1e20, big_ext.ToFloat64(askTokens[130])) // USDT
}

func TestAggregateOrderDepthDetails_priceRounding(t *testing.T) {
	/// Given
	// Base token = EVIV
	// Quote token = USDT
	orderDetails := []*db.Order{
		fakeBuyOrder(50.005, fakeDbAsset("USDT", 10), fakeAnyValueDbAsset("EVIV"), 0),
		fakeSellOrder(50.004, fakeDbAsset("EVIV", 30), fakeAnyValueDbAsset("USDT"), 0),
	}

	/// When
	bidTokens, keyBidTokens, askTokens, keyAskTokens := aggregateOrderDepthDetails(orderDetails)

	/// Then
	assert.Equal(t, []float64{50}, keyAskTokens)
	assert.Len(t, askTokens, 1)
	assert.EqualValues(t, 1500.12, big_ext.ToFloat64(askTokens[50]))

	assert.Equal(t, []float64{50.01}, keyBidTokens)
	assert.Len(t, bidTokens, 1)
	assert.EqualValues(t, 10, big_ext.ToFloat64(bidTokens[50.01]))
}
