package handlers

import (
	"mines/src/api/customHTTP"
	"mines/src/mines"
	"net/http"
)

func GetField(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	customHTTP.Encode(w, mines.Field)
}
