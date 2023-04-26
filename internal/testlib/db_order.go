package testlib

import (
	"github.com/go-faker/faker/v4"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"math/big"
)

type FakeDbOrderModifierFn func(order *db.Order)
type FakeDbAssetFn func() db.Assets

func RandomDbAsset() FakeDbAssetFn {
	return func() db.Assets {
		var asset db.Assets
		if err := faker.FakeData(&asset); err != nil {
			panic(err)
		}

		// Random asset "token" name
		asset.VirtualToken = faker.Currency()

		// Random asset value
		value, err := faker.RandomInt(1, 100, 1)
		if err != nil {
			panic(err)
		}

		asset.Value = big.NewInt(int64(value[0])).String()

		return asset
	}
}

func FakeAnyValueDbAsset(token string) FakeDbAssetFn {
	return func() db.Assets {
		var asset db.Assets
		if err := faker.FakeData(&asset); err != nil {
			panic(err)
		}

		asset.VirtualToken = token

		// Random asset value
		value, err := faker.RandomInt(1, 100, 1)
		if err != nil {
			panic(err)
		}

		asset.Value = big.NewInt(int64(value[0])).String()

		return asset
	}
}

func FakeDbAsset(token string, value uint64) FakeDbAssetFn {
	return func() db.Assets {
		var asset db.Assets
		if err := faker.FakeData(&asset); err != nil {
			panic(err)
		}

		asset.VirtualToken = token
		asset.Value = big.NewInt(int64(value)).String()

		return asset
	}
}

func WithAssets(makeAsset, takeAsset FakeDbAssetFn) FakeDbOrderModifierFn {
	return func(order *db.Order) {
		order.Assets = []db.Assets{makeAsset(), takeAsset()}
	}
}

func WithPrice(price float64) FakeDbOrderModifierFn {
	return func(order *db.Order) {
		order.Price = price
	}
}

func WithFills(fills int64) FakeDbOrderModifierFn {
	return func(order *db.Order) {
		order.Fills = big.NewInt(fills).String()
	}
}

func WithStatus(status db.MatchedStatus) FakeDbOrderModifierFn {
	return func(order *db.Order) {
		order.Status = status
	}
}

func FakeDbOrder(isShort bool, modifiers ...FakeDbOrderModifierFn) *db.Order {
	var order *db.Order
	if err := faker.FakeData(&order); err != nil {
		panic(err)
	}

	order.IsShort = isShort

	for _, modify := range modifiers {
		modify(order)
	}

	return order
}
