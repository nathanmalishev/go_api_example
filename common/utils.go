package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"code"`
	}
	AppResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

var (
	InvalidData  = "Could not decode data"
	FetchError   = "Could not fetch data"
	JwtHTTPError = "No valid JWT attached"
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errObj); err == nil {
		w.Write(j)
	}
}

func WriteJson(w http.ResponseWriter, message string, data interface{}, code int) {

	mess := AppResponse{
		Message: message,
		Data:    data,
	}
	j, err := json.Marshal(mess)
	if err != nil {
		DisplayAppError(w, err, InvalidData, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(j)
}
