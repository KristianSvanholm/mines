package errorHandler

import (
	"log"
	"net/http"
)

// Err is a flexible error handler that chooses to print error to web-request or console depending on w's existence
func Err(w http.ResponseWriter, errorTxt string, errorCode int) {
	if w != nil {
		http.Error(w, errorTxt, errorCode)
	} else {
		log.Print(errorTxt)
	}
}
