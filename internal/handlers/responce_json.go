package handlers

import (
	"encoding/json"
	"net/http"
)

func responseJson(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]any{
		"error": err.Error(),
	})
}
