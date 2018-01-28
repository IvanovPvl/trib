package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const validErrMessage = "number argument of type int is needed"

func tribHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/trib", handleGetTrib)

	return mux
}

func handleGetTrib(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	numberParam := r.URL.Query().Get("number")
	if numberParam == "" {
		sendError(w, errors.New(validErrMessage))
		return
	}

	number, err := strconv.Atoi(numberParam)
	if err != nil {
		sendError(w, err)
		return
	}

	tribb, err := trib(number)
	if err != nil {
		sendError(w, err)
		return
	}

	resp := SuccessResponse{tribb}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func sendError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusBadRequest)
	errResp := ErrorResponse{e.Error()}
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		panic(err)
	}
}
