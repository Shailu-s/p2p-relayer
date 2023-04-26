package models

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/volmexfinance/relayers/relayer-srv/big_ext"
	"math/big"
)

type MatchedStatus uint
type OrderType string

const (
	MatchedStatusZero                  MatchedStatus = 1
	MatchedStatusInit                  MatchedStatus = 2
	MatchedStatusFullMatchFound        MatchedStatus = 3
	MatchedStatusPartialMatchFound     MatchedStatus = 4
	MatchedStatusValidated             MatchedStatus = 5
	MatchedStatusValidationConfirmed   MatchedStatus = 6
	MatchedStatusSentToContract        MatchedStatus = 7
	MatchedStatusPartialMatchConfirmed MatchedStatus = 8
	MatchedStatusFullMatchConfirmed    MatchedStatus = 9
	MatchedStatusSentFailed            MatchedStatus = 10
	MatchedStatusFailedConfirmed       MatchedStatus = 11
	MatchedStatusBlocked               MatchedStatus = 12
	Canceled                           MatchedStatus = 13
)

func (m MatchedStatus) String() string {
	switch m {
	case MatchedStatusZero:
		return "zero"
	case MatchedStatusInit:
		return "init"
	case MatchedStatusFullMatchFound:
		return "full_match_found"
	case MatchedStatusPartialMatchFound:
		return "partial_match_found"
	case MatchedStatusValidated:
		return "validated"
	case MatchedStatusValidationConfirmed:
		return "validation_confirmed"
	case MatchedStatusSentToContract:
		return "sent_to_contract"
	case MatchedStatusPartialMatchConfirmed:
		return "partial_match_confirmed"
	case MatchedStatusFullMatchConfirmed:
		return "full_match_confirmed"
	case MatchedStatusSentFailed:
		return "sent_failed"
	case MatchedStatusFailedConfirmed:
		return "failed_confirmed"
	case MatchedStatusBlocked:
		return "blocked"
	case Canceled:
		return "canceled"
	default:
		return "unknown"
	}
}

const (
	OrderTypeOrder                OrderType = "0xf555eb98" // bytes4(keccack256(abi.encodePacked("Order")))
	OrderTypeStopLossIndexPrice   OrderType = "0x835d5c1e" // bytes4(keccak256(abi.encodePacked("StopLossIndexPrice")));
	OrderTypeStopLossLastPrice    OrderType = "0xd9ed8042" // bytes4(keccak256(abi.encodePacked("StopLossLastPrice")));
	OrderTypeStopLossMarkPrice    OrderType = "0xe144c7ec" // bytes4(keccak256(abi.encodePacked("StopLossMarkPrice")));
	OrderTypeTakeProfitIndexPrice OrderType = "0x67393efa" // bytes4(keccak256(abi.encodePacked("TakeProfitIndexPrice")));
	OrderTypeTakeProfitLastPrice  OrderType = "0xc7dc86f6" // bytes4(keccak256(abi.encodePacked("TakeProfitLastPrice")));
	OrderTypeTakeProfitMarkPrice  OrderType = "0xb6d64e04" // bytes4(keccak256(abi.encodePacked("TakeProfitMarkPrice")));
)

func CheckAndParseOrderType(orderType string) (OrderType, error) {
	switch OrderType(orderType) {
	case OrderTypeOrder,
		OrderTypeStopLossIndexPrice,
		OrderTypeStopLossLastPrice,
		OrderTypeStopLossMarkPrice,
		OrderTypeTakeProfitIndexPrice,
		OrderTypeTakeProfitLastPrice,
		OrderTypeTakeProfitMarkPrice:
		return OrderType(orderType), nil
	default:
		return "", fmt.Errorf("unknown order type: %s", orderType)
	}
}

func (o OrderType) String() string {
	switch o {
	case OrderTypeOrder:
		return "order"
	case OrderTypeStopLossIndexPrice:
		return "stop_loss_index_price"
	case OrderTypeStopLossLastPrice:
		return "stop_loss_last_price"
	case OrderTypeStopLossMarkPrice:
		return "stop_loss_mark_price"
	case OrderTypeTakeProfitIndexPrice:
		return "take_profit_index_price"
	case OrderTypeTakeProfitLastPrice:
		return "take_profit_last_price"
	case OrderTypeTakeProfitMarkPrice:
		return "take_profit_mark_price"
	default:
		return "unknown"
	}
}

func (o OrderType) Bytes() [4]byte {
	var byteArr [4]byte
	copy(byteArr[:], common.FromHex(string(o)))
	return byteArr
}

// TODO: change price here from float to string
type Order struct {
	OrderID      string        `gorm:"primaryKey" json:"order_id"`
	OrderType    OrderType     `json:"order_type"`
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
func NewOrder(trader, salt, chain string, orderType OrderType, triggerPrice, signature string, deadline uint64, isShort bool, makeAsset, takeAsset Assets) *Order {
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
