package app

import (
	"encoding/json"
	"net/http"
)

type errorResp struct {
	Error string `json:"error"`
}

// responJSON makes the response with payload as json format
func responJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// responError makes the error response with payload as json format
func responError(w http.ResponseWriter, code int, message string) {
	responJSON(w, code, errorResp{Error: message})
}

func responOk(w http.ResponseWriter, payload interface{}) {
	responJSON(w, http.StatusOK, payload)
}
