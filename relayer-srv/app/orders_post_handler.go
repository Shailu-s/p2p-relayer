package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/volmexfinance/relayers/relayer-srv/big_ext"
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"github.com/volmexfinance/relayers/relayer-srv/validation"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

type reqAsset struct {
	VirtualToken string `json:"virtual_token"`
	Value        string `json:"value"`
}

type reqOrder struct {
	OrderType    string   `json:"order_type"`
	Deadline     string   `json:"deadline"`
	Trader       string   `json:"trader"`
	MakeAsset    reqAsset `json:"make_asset"`
	TakeAsset    reqAsset `json:"take_asset"`
	Salt         string   `json:"salt"`
	TriggerPrice string   `json:"trigger_price"`
	IsShort      bool     `json:"is_short"`
	Signature    string   `json:"signature"`
}

// toDBAssets validates and converts an order asset received via REST API to the data model used in DB.
func toDbAssets(reqAsset reqAsset) (models.Assets, error) {
	var dbAssets models.Assets
	if reqAsset.VirtualToken == "" {
		return dbAssets, errors.New("empty virtual token")
	}

	if reqAsset.Value == "" {
		return dbAssets, errors.New("empty value")
	}

	value, err := big_ext.CheckAndParseBigFloat(reqAsset.Value)
	if err != nil {
		return dbAssets, err
	}

	if big_ext.FloatLessThanOrEqual(value, big.NewFloat(0)) {
		return dbAssets, errors.New("value must be greater than 0")
	}

	dbAssets = models.Assets{
		VirtualToken: reqAsset.VirtualToken,
		Value:        reqAsset.Value,
	}
	return dbAssets, nil
}

// toDBOrder converts an order received via REST API to the data model used in DB.
func toDBOrder(req reqOrder, chain string) (*models.Order, error) {
	// check and parse request signature
	if req.Signature == "" {
		return nil, errors.New("missing signature")
	}

	// check and parse request trader
	if req.Trader == "" {
		return nil, errors.New("invalid trader")
	}

	// check and parse request order type
	orderType, err := models.CheckAndParseOrderType(req.OrderType)
	if err != nil {
		return nil, fmt.Errorf("invalid order type: %w", err)
	}

	// check and parse request order salt
	if req.Salt == "" {
		return nil, errors.New("invalid salt")
	}

	salt, err := big_ext.CheckAndParseBigInt(req.Salt)
	if err != nil {
		return nil, fmt.Errorf("invalid salt: %w", err)
	}

	if big_ext.LessThanOrEqual(salt, big.NewInt(0)) {
		return nil, errors.New("invalid salt: must be greater than 0")
	}

	// check and parse take and make assets
	if req.TakeAsset.VirtualToken == req.MakeAsset.VirtualToken {
		return nil, errors.New("make and take asset virtual tokens are the same")
	}

	orderMakeAsset, err := toDbAssets(req.MakeAsset)
	if err != nil {
		return nil, fmt.Errorf("invalid make asset: %w", err)
	}

	orderTakeAsset, err := toDbAssets(req.TakeAsset)
	if err != nil {
		return nil, fmt.Errorf("invalid take asset: %w", err)
	}

	// check and parse request order deadline
	deadline, err := strconv.ParseUint(req.Deadline, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid deadline value: %w", err)
	}

	// Map request order to DB order
	order := models.NewOrder(
		req.Trader,
		req.Salt,
		chain,
		orderType,
		req.TriggerPrice,
		req.Signature,
		deadline,
		req.IsShort,
		orderMakeAsset,
		orderTakeAsset,
	)

	return order, nil
}

// Endpoint: POST /insertOrder/{chain}
func (a *App) ordersPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Parameters
	chain, ok := vars["chain"]
	if !ok || chain == "" {
		a.logger.Errorf("invalid 'chain' parameter value")
		responError(w, http.StatusBadRequest, "")
		return
	}

	// Body
	body := r.Body
	if body == nil {
		a.logger.Errorf("cannot decode 'nil' body")
		responError(w, http.StatusBadRequest, "")
		return
	}

	var order reqOrder
	if err := json.NewDecoder(body).Decode(&order); err != nil {
		a.logger.Errorf("cannot decode request's JSON body: %v", err)
		responError(w, http.StatusBadRequest, "")
		return
	}

	dbOrder, err := toDBOrder(order, strings.ToUpper(chain))
	if err != nil {
		a.logger.Errorf("failed to conver request to db order: %v", err)
		responError(w, http.StatusBadRequest, "")
		return
	}

	//TODO: to be completed with security tasks
	if err = validation.VerifyOrderSignature(*dbOrder, a.relayer.GetChainID(chain), a.relayer.GetPositioningContract(chain)); err != nil {
		a.logger.Errorf("err in validation: %v", err)
		responError(w, http.StatusBadRequest, "")
		return
	}

	// insert order in the database
	if err = a.relayer.InsertOrder(dbOrder); err != nil {
		a.logger.Errorf("order instertion failed: %v", err)
		responError(w, http.StatusInternalServerError, "")
		return
	}

	a.logger.Infof("added new order to db, order_id: %s", dbOrder.OrderID)

	// signature, err := a.relayer.CreateSignature(amount, recipientAddress, destinationChainID)
	// if err != nil {
	// 	common.ResponError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	responOk(w, "")
}
