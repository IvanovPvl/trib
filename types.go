package main

import "math/big"

type SuccessResponse struct {
	Tribonacci *big.Int `json:"tribonacci_number"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
