package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/gavv/httpexpect.v1"
)

var successResponses = map[string]map[string]int{
	"0":  {"tribonacci_number": 0},
	"20": {"tribonacci_number": 35890},
}

var errResponse = map[string]string{
	"message": validErrMessage,
}

var negResponse = map[string]string{
	"message": negErrMessage,
}

func TestHandler(t *testing.T) {
	handler := tribHandler()
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	for k, v := range successResponses {
		e.GET("/trib").
			WithQuery("number", k).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			Equal(v)
	}

	e.GET("/trib").
		WithQuery("number", "").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(errResponse)

	e.GET("/trib").
		WithQuery("number", "-10").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(negResponse)
}
