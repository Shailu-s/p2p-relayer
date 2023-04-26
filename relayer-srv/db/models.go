package db

import (
	"math/big"

	"github.com/lib/pq"
	"github.com/volmexfinance/relayers/relayer-srv/big_ext"
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

// TODO: change price here from float to string
type Order struct {
	OrderID      string        `gorm:"primaryKey" json:"order_id"`
	OrderType    string        `json:"order_type"`
	Trader       string        `json:"trader"`
	Deadline     uint64        `json:"deadline"`
	IsShort      bool          `json:"is_short"`
	Assets       []Assets      `gorm:"foreignKey:OrderbookID"`
	Status       MatchedStatus `json:"status"`
	Price        float64       `json:"price"`
	Salt         string        `json:"salt"`
	TriggerPrice string        `json:"trigger_price"`
	Sign         string        `json:"sign"`
	FailCount    int64         `json:"fail_count"`
	Fills        string        `json:"fills"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
	ChainName    string        `json:"chain_name"`
}

// CreateOrderID create order id by concatenating trader, salt and chain fields.
func CreateOrderID(trader, salt, chain string) string {
	return trader + salt + chain
}

// CalculatePrice calculates price from order assets.
func CalculatePrice(isShort bool, makeValue, takeValue *big.Float) float64 {
	var price float64
	if isShort {
		// Sell order price = takeAsset / makeAsset (e.g., USDT/EVIV)
		price, _ = new(big.Float).Quo(takeValue, makeValue).Float64()
	} else {
		// Buy order price = makeAsset / takeAsset (e.g., USDT/EVIV)
		price, _ = new(big.Float).Quo(makeValue, takeValue).Float64()
	}
	return price
}

// NewOrder creates a new order instance.
func NewOrder(trader, salt, chain, orderType, triggerPrice, signature string, deadline uint64, isShort bool, makeAsset, takeAsset Assets) *Order {
	orderID := CreateOrderID(trader, salt, chain)
	price := CalculatePrice(isShort, makeAsset.ValueAsBigFloat(), takeAsset.ValueAsBigFloat())

	return &Order{
		OrderID:      orderID,
		Salt:         salt,
		OrderType:    orderType,
		Trader:       trader,
		ChainName:    chain,
		Deadline:     deadline,
		TriggerPrice: triggerPrice,
		Sign:         signature,
		FailCount:    0,
		IsShort:      isShort,
		Assets:       []Assets{makeAsset, takeAsset},
		Price:        price,
		Fills:        "0",
		Status:       MatchedStatusZero,
	}
}

// MakeAsset getter method to access the order's make asset token-value pair
func (o *Order) MakeAsset() *Assets {
	return &o.Assets[OrderAssetMake]
}

// TakeAsset getter method to access the order's take asset token-value pair
func (o *Order) TakeAsset() *Assets {
	return &o.Assets[OrderAssetTake]
}

func (o *Order) PriceAsBigFloat() *big.Float {
	return big.NewFloat(o.Price)
}

func (o *Order) OrderFills() *big.Int {
	return big_ext.ParseBigInt(o.Fills)
}

func (o *Order) OrderFillsAsBigFloat() *big.Float {
	return big_ext.ParseBigFloat(o.Fills)
}

func (o *Order) SetOrderFills(value *big.Int) {
	o.Fills = value.String()
}

func (o *Order) OrderSalt() *big.Int {
	return big_ext.ParseBigInt(o.Salt)
}

func (o *Order) OrderTriggerPrice() *big.Int {
	return big_ext.ParseBigInt(o.TriggerPrice)
}

// TODO: Move this table into postgresDB
type TransactionSent struct {
	ID                int                   `gorm:"primaryKey"`
	TransactionHash   string                `json:"transaction_hash"`
	OrderID           pq.StringArray        `gorm:"type:text[]"`
	TransactionStatus TransactionStatusType `gorm:"type:BIGINT"`
	Error             string                `json:"error"`
	CreatedAt         int64                 `json:"created_at"`
	UpdatedAt         int64                 `json:"updated_at"`
	Leader            string                `json:"leader"`
	LeaderPublicKey   string                `json:"leader_public_key"`
	ChainName         string                `json:"chain_name"`
}

type TransactionLog struct {
	ID              int                      `gorm:"primaryKey"`
	Traders         pq.StringArray           `json:"traders" gorm:"type:TEXT[]"`
	Deadline        pq.Int64Array            `json:"deadline" gorm:"type:BIGINT[]"`
	Salt            pq.StringArray           `json:"salt" gorm:"type:TEXT[]"`
	NewLeftFill     string                   `json:"new_left_fill" gorm:"type:TEXT"`
	NewRightFill    string                   `json:"new_right_fill" gorm:"type:TEXT"`
	CreatedAt       int64                    `json:"createdAt"`
	UpdatedAt       int64                    `json:"updatedAt"`
	OrderID         pq.StringArray           `json:"order_id" gorm:"type:TEXT[]"`
	Status          TransactionLogStatusType `json:"status" gorm:"type:BIGINT"`
	BlockHeight     uint64                   `json:"block_height" gorm:"type:BIGINT"`
	ChainName       string                   `json:"chain_name"`
	TransactionHash string                   `json:"transaction_hash" gorm:"type:TEXT"`
}
