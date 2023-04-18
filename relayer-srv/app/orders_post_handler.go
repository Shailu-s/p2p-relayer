package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/volmexfinance/relayers/relayer-srv/big_ext"
	"github.com/volmexfinance/relayers/relayer-srv/db"
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

func checkReqOrder(req reqOrder) error {
	if req.Signature == "" {
		return errors.New("missing signature")
	}

	if req.Salt == "" {
		return errors.New("invalid salt")
	}

	if req.Trader == "" {
		return errors.New("invalid trader")
	}

	if req.MakeAsset.VirtualToken == "" {
		return errors.New("invalid make asset virtual token")
	}

	if req.TakeAsset.VirtualToken == "" {
		return errors.New("invalid take asset virtual token")
	}

	if req.MakeAsset.Value == "" {
		return errors.New("invalid make asset value")
	}

	if req.TakeAsset.Value == "" {
		return errors.New("invalid take asset value")
	}

	return nil
}

// toDBAssets converts an order asset received via REST API to the data model used in DB.
func toDbAssets(assets reqAsset) db.Assets {
	return db.Assets{
		VirtualToken: assets.VirtualToken,
		Value:        assets.Value,
	}
}

// toDBOrder converts an order received via REST API to the data model used in DB.
func toDBOrder(req reqOrder, chain string) (*db.Order, error) {

	// check and parse request order salt
	salt, err := big_ext.CheckAndParseBigInt(req.Salt)
	if err != nil {
		return nil, fmt.Errorf("invalid salt: %w", err)
	}

	if big_ext.LessThanOrEqual(salt, big.NewInt(0)) {
		return nil, errors.New("invalid salt: must be greater than 0")
	}

	// check take and make asset virtual tokens are not the same
	if req.TakeAsset.VirtualToken == req.MakeAsset.VirtualToken {
		return nil, errors.New("make and take asset virtual tokens are the same")
	}

	// check and parse request order make asset
	makeValue, err := big_ext.CheckAndParseBigFloat(req.MakeAsset.Value)
	if err != nil {
		return nil, fmt.Errorf("invalid make asset value: %w", err)
	}

	if big_ext.FloatLessThanOrEqual(makeValue, big.NewFloat(0)) {
		return nil, errors.New("invalid make asset value: must be greater than 0")
	}

	// check and parse request order take asset
	takeValue, err := big_ext.CheckAndParseBigFloat(req.TakeAsset.Value)
	if err != nil {
		return nil, fmt.Errorf("invalid take asset value: %w", err)
	}

	if big_ext.FloatLessThanOrEqual(takeValue, big.NewFloat(0)) {
		return nil, errors.New("invalid take asset value: must be greater than 0")
	}

	// check and parse request order deadline
	deadline, err := strconv.ParseUint(req.Deadline, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid deadline value: %w", err)
	}

	// Map request order to DB order
	orderMakeAsset := toDbAssets(req.MakeAsset)
	orderTakeAsset := toDbAssets(req.TakeAsset)

	order := db.NewOrder(
		req.Trader,
		req.Salt,
		chain,
		req.OrderType,
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

	if err := checkReqOrder(order); err != nil {
		a.logger.Errorf("bad insert order request: %s", err.Error())
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
