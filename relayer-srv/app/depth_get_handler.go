package app

import (
	"github.com/gorilla/mux"
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"math"
	"math/big"
	"net/http"
)

type depthOrderDetail struct {
	Price float64 `json:"price"`
	Size  string  `json:"size"`
}

type depthOrderBookDetails struct {
	Bid []depthOrderDetail `json:"bid"`
	Ask []depthOrderDetail `json:"ask"`
}

func roundWithDecimals(value float64, decimals int) float64 {
	pow := math.Pow(10, float64(decimals))
	return math.Round(value*pow) / pow
}

func aggregateOrderDepthDetails(orderDetails []*models.Order) (map[float64]*big.Float, []float64, map[float64]*big.Float, []float64) {
	var bidTokens = make(map[float64]*big.Float)
	var keyBidTokens = make([]float64, 0, len(bidTokens))
	var askTokens = make(map[float64]*big.Float)
	var keyAskTokens = make([]float64, 0, len(askTokens))

	for _, order := range orderDetails {
		price := roundWithDecimals(order.Price, 2)
		if order.IsShort {
			// Sell order:
			//  - baseToken = Make asset (e.g., EVIV)
			//  - quoteToken = Take asset (e.g., USDT)
			makeAsset := order.MakeAsset().ValueAsBigFloat()

			// value = (makeAsset - fillsMakeAsset) * price (e.g., (EVIV - EVIV) * USDT/EVIV = USDT)
			value := new(big.Float).Sub(makeAsset, order.OrderFillsAsBigFloat())
			value = value.Mul(value, order.PriceAsBigFloat())

			if askTokens[price] != nil {
				askTokens[price] = askTokens[price].Add(askTokens[price], value)
			} else {
				keyAskTokens = append(keyAskTokens, price)
				askTokens[price] = value
			}
		} else {
			// Buy order:
			//  - baseToken = Take asset (e.g., EVIV)
			//  - quoteToken = Make asset (e.g., USDT)
			makeAsset := order.MakeAsset().ValueAsBigFloat()

			// value = makeAsset - fillsMakeAsset (e.g., USDT - USDT = USDT)
			value := new(big.Float).Sub(makeAsset, order.OrderFillsAsBigFloat())

			if bidTokens[price] != nil {
				bidTokens[price] = bidTokens[price].Add(bidTokens[price], value)
			} else {
				keyBidTokens = append(keyBidTokens, price)
				bidTokens[price] = value
			}
		}
	}

	return bidTokens, keyBidTokens, askTokens, keyAskTokens
}

func toDepthResponse(bidTokens map[float64]*big.Float, keyBidTokens []float64, askTokens map[float64]*big.Float, keyAskTokens []float64) depthOrderBookDetails {
	bids := make([]depthOrderDetail, 0, len(keyBidTokens))
	for i := range keyBidTokens {
		price := keyBidTokens[i]
		size := bidTokens[price]
		bids = append(bids, depthOrderDetail{Price: price, Size: size.Text('f', 0)})
	}

	asks := make([]depthOrderDetail, 0, len(keyAskTokens))
	for i := len(keyAskTokens) - 1; i >= 0; i-- {
		price := keyAskTokens[i]
		size := askTokens[price]
		asks = append(asks, depthOrderDetail{Price: price, Size: size.Text('f', 0)})
	}

	return depthOrderBookDetails{Bid: bids, Ask: asks}
}

// Endpoint: GET /depth/{token}/{chain}
func (a *App) depthGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	token, ok := vars["token"]
	if !ok || token == "" {
		a.logger.Errorf("invalid 'token' parameter value")
		responError(w, http.StatusBadRequest, "")
		return
	}
	chain, ok := vars["chain"]
	if !ok || chain == "" {
		a.logger.Errorf("invalid 'chain' parameter value")
		responError(w, http.StatusBadRequest, "")
		return
	}

	// check if base token exists
	hasBaseToken, err := a.relayer.CheckHasBaseToken(token, chain)
	if err != nil {
		a.logger.Errorf("check has base token failed: %v", err)
		responError(w, http.StatusInternalServerError, "")
		return
	}
	if !hasBaseToken {
		a.logger.Errorf("base token not found in chain %s: %s", chain, token)
		responError(w, http.StatusNotFound, "")
		return
	}

	// get depth details for base token
	orderDetails, err := a.relayer.GetDepthOrderDetails(token, chain)
	if err != nil {
		a.logger.Errorf("get depth order details failed: %v", err)
		responError(w, http.StatusInternalServerError, "")
		return
	}

	// curate depth details
	bidTokens, keyBidTokens, askTokens, keyAskTokens := aggregateOrderDepthDetails(orderDetails)
	responseDepthDetails := toDepthResponse(bidTokens, keyBidTokens, askTokens, keyAskTokens)

	responOk(w, responseDepthDetails)
}
