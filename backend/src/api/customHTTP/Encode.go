package customHTTP

import (
	"encoding/json"
	"mines/src/api/errorHandler"
	"net/http"
)

// Encode any struct given to it
func Encode(w http.ResponseWriter, data interface{}) bool {
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		errorHandler.Err(w, "Error during encoding of data", http.StatusInternalServerError)
		return false
	}
	return true
}
