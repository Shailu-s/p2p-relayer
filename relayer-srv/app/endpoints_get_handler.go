package app

import (
	"net/http"
)

type endpoints struct {
	Endpoints []string `json:"endpoints"`
}

// Endpoint: GET /
func (a *App) endpointsGetHandler(w http.ResponseWriter, r *http.Request) {
	resp := endpoints{
		Endpoints: []string{
			"/insertOrder/{chain}",
			"/getOrders/{trader}/{chain}",
			"/getOrderQueue/{chain}",
			"/depth/{token}/{chain}",
		},
	}

	responOk(w, resp)
}
