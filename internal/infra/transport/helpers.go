package transport

import (
	"encoding/json"
	"net/http"
)

// writeResponse helper function for writing JSON responses to the http.ResponseWriter
// sets content type to json and set header with correct status code.
// handle error if json encoding fails.
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "binding error message failed", http.StatusInternalServerError)
	}
}
