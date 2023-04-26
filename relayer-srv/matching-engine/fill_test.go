package matching_engine

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volmexfinance/relayers/internal/testlib"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"testing"
)

func fakeDbAsset(token string, value uint64) testlib.FakeDbAssetFn {
	return testlib.FakeDbAsset(token, value)
}

func fakeDbOrder(isShort bool, status db.MatchedStatus, fills int64, makeAssetFn, takeAssetFn testlib.FakeDbAssetFn) *db.Order {
	return testlib.FakeDbOrder(
		isShort,
		testlib.WithAssets(makeAssetFn, takeAssetFn),
		testlib.WithFills(fills),
		testlib.WithStatus(status),
	)
}

func TestMatchAndUpdateOrders_StatusInitToFull(t *testing.T) {
	/// Given
	// Sell order giving 100 EVIV in exchange for 200 USDT
	leftOrder := fakeDbOrder(true, db.MatchedStatusInit, 0,
		fakeDbAsset("EVIV", 100), // Make asset
		fakeDbAsset("USDT", 200), // Take asset
	)
	// Buy order giving 200 USDT in exchange for 100 EVIV
	rightOrder := fakeDbOrder(false, db.MatchedStatusInit, 0,
		fakeDbAsset("USDT", 200), // Make asset
		fakeDbAsset("EVIV", 100), // Take asset
	)

	/// When
	err := matchAndUpdateOrders(leftOrder, rightOrder)

	/// Then
	require.NoError(t, err)

	// The amount of 100 EVIV (Make assets) has been FULLY matched, 0 left
	assert.Equal(t, "100", leftOrder.Fills)
	assert.Equal(t, db.MatchedStatusFullMatchFound, leftOrder.Status)

	// The amount of 200 USDT (Make assets) has been FULLY matched, 0 left
	assert.Equal(t, "200", rightOrder.Fills)
	assert.Equal(t, db.MatchedStatusFullMatchFound, rightOrder.Status)
}

func TestMatchAndUpdateOrders_StatusInitToPartial(t *testing.T) {
	/// Given
	// Sell order giving 400 EVIV in exchange for 200 USDT
	leftOrder := fakeDbOrder(true, db.MatchedStatusInit, 0,
		fakeDbAsset("EVIV", 400), // Make asset
		fakeDbAsset("USDT", 200), // Take asset
	)
	// Buy order giving 200 USDT in exchange for 100 EVIV
	rightOrder := fakeDbOrder(false, db.MatchedStatusInit, 0,
		fakeDbAsset("USDT", 100), // Make asset
		fakeDbAsset("EVIV", 200), // Take asset
	)

	/// When
	err := matchAndUpdateOrders(leftOrder, rightOrder)

	/// Then
	require.NoError(t, err)

	// The amount of 200 EVIV (Make assets) has been PARTIALLY matched, 200 EVIV left
	assert.Equal(t, "200", leftOrder.Fills)
	assert.Equal(t, db.MatchedStatusPartialMatchFound, leftOrder.Status)

	// The amount of 100 USDT (Make assets) has been FULLY matched, 0 USDT left
	assert.Equal(t, "100", rightOrder.Fills)
	assert.Equal(t, db.MatchedStatusFullMatchFound, rightOrder.Status)
}

func TestMatchAndUpdateOrders_StatusPartialToFull(t *testing.T) {
	/// Given
	// Sell order giving 400 EVIV in exchange for 200 USDT
	leftOrder := fakeDbOrder(true, db.MatchedStatusInit, 50,
		fakeDbAsset("EVIV", 400), // Make asset
		fakeDbAsset("USDT", 200), // Take asset
	)
	// Buy order giving 200 USDT in exchange for 100 EVIV
	rightOrder := fakeDbOrder(false, db.MatchedStatusInit, 20,
		fakeDbAsset("USDT", 100), // Make asset
		fakeDbAsset("EVIV", 200), // Take asset
	)

	/// When
	err := matchAndUpdateOrders(leftOrder, rightOrder)

	/// Then
	require.NoError(t, err)

	// The amount of 210 EVIV (Make assets) has been PARTIALLY matched, 190 EVIV left
	assert.Equal(t, "210", leftOrder.Fills)
	assert.Equal(t, db.MatchedStatusPartialMatchFound, leftOrder.Status)

	// The amount of 100 USDT (Make assets) has been FULLY matched, 0 USDT left
	assert.Equal(t, "100", rightOrder.Fills)
	assert.Equal(t, db.MatchedStatusFullMatchFound, rightOrder.Status)
}
