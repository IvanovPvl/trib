package main

import (
	"errors"
	"math/big"
)

const negErrMessage = "number must be positive"

var cache = make(map[int]*big.Int)

func trib(n int) (*big.Int, error) {
	if r, ok := cache[n]; ok {
		return r, nil
	}

	if n < 0 {
		return nil, errors.New(negErrMessage)
	}

	for i := 0; i <= n; i++ {
		var f = big.NewInt(0)
		if i <= 1 {
			f.SetUint64(0)
		} else if i == 2 {
			f.SetUint64(1)
		} else {
			s := f.Add(cache[i-1], cache[i-2])
			f = f.Add(cache[i-3], s)
		}

		cache[i] = f
	}

	return cache[n], nil
}
