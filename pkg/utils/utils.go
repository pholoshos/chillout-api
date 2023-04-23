// internal/utils/utils.go

package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSONResponse writes a JSON response to the HTTP response writer.
func WriteJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// WriteJSONError writes a JSON error response to the HTTP response writer.
func WriteJSONError(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

// DecodeJSONBody decodes the JSON request body into the provided struct.
func DecodeJSONBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
