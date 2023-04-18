package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func (a *App) ordersQueueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Parameters
	chain, ok := vars["chain"]
	if !ok || chain == "" {
		a.logger.Errorf("invalid 'chain' parameter value")
		responError(w, http.StatusBadRequest, "")
		return
	}

	orders, err := a.relayer.GetOrderQueue(strings.ToUpper(chain))
	if err != nil {
		a.logger.Errorf("failed to get order queue: %v", err)
		responError(w, http.StatusInternalServerError, "")
		return
	}

	// Here we are returning directly from the DB the order queue model
	// TODO: Define an app specific response serializable struct to decouple from DB model
	responOk(w, orders)
}
