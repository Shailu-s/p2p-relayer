package models

import (
	"github.com/volmexfinance/relayers/relayer-srv/big_ext"
	"math/big"
)

type OrderAssetType uint

const (
	OrderAssetMake OrderAssetType = 0
	OrderAssetTake OrderAssetType = 1
)

type Assets struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	OrderbookID  string `json:"orderbook_id"`
	VirtualToken string `json:"virtual_token"`
	Value        string `json:"value"`
}

func (a *Assets) ValueAsBigInt() *big.Int {
	return big_ext.ParseBigInt(a.Value)
}

func (a *Assets) ValueAsBigFloat() *big.Float {
	return big_ext.ParseBigFloat(a.Value)
}
