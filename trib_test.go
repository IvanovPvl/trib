package main

import (
	"math/big"
	"testing"
)

type testPair struct {
	value  int
	result *big.Int
}

var tests = []testPair{
	{0, big.NewInt(0)},
	{1, big.NewInt(0)},
	{2, big.NewInt(1)},
	{20, big.NewInt(35890)},
}

func TestTrib(t *testing.T) {
	for _, pair := range tests {
		v, _ := trib(pair.value)
		if v.Cmp(pair.result) != 0 {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}

	v, err := trib(-10)
	if err == nil {
		t.Error(
			"For", -10,
			"expected", err,
			"got", v,
		)
	}
}
