package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (a *App) ordersGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Parameters
	trader, ok := vars["trader"]
	if !ok || trader == "" {
		a.logger.Errorf("invalid 'trader' parameter value")
		responError(w, http.StatusBadRequest, "")
		return
	}

	chain, ok := vars["chain"]
	if !ok || chain == "" {
		a.logger.Errorf("invalid 'chain' parameter value")
		responJSON(w, http.StatusBadRequest, "")
		return
	}

	orders, err := a.relayer.GetAllOrders(trader, chain)
	if err != nil {
		a.logger.Errorf("failed to get orders: %v", err)
		responError(w, http.StatusInternalServerError, "")
		return
	}

	// Here we are returning directly from the DB the order queue model
	// TODO: Define an app specific response serializable struct to decouple from DB model
	responOk(w, orders)
}
