package np

import (
	"errors"
	"math"
)

type Float64OneDArray struct {
	arr  []float64
	step float64
	err  error
}

func Linspace(start float64, stop float64, num int, stopIncluded bool) Float64OneDArray {
	if num < 2 {
		return Float64OneDArray{nil, 0, errors.New(ErrInvalidParameter)}
	}

	if start == stop {
		return Float64OneDArray{nil, 0, errors.New(ErrInvalidParameter)}
	}

	result := make([]float64, 0, num)
	step := (stop - start) / float64(num)
	if stopIncluded {
		step = (stop - start) / float64(num-1)
	}

	for i := 0; i < num; i++ {
		result = append(result, start)
		start += step
	}

	return Float64OneDArray{result, step, nil}
}

func Logspace(start float64, stop float64, num int, base int, stopIncluded bool) Float64OneDArray {
	return Float64OneDArray{nil, 0, nil}
}

func Geomspace(start float64, stop float64, num int, stopIncluded bool) Float64OneDArray {
	if start == 0 {
		return Float64OneDArray{nil, 0, errors.New(ErrZeroValue)}
	}

	if start < 0 || stop < 0 {
		return Float64OneDArray{nil, 0, errors.New(ErrNegativeValue)}
	}

	if num < 2 {
		return Float64OneDArray{nil, 0, errors.New(ErrInvalidParameter)}
	}

	if start == stop {
		return Float64OneDArray{nil, 0, errors.New(ErrInvalidParameter)}
	}

	result := make([]float64, 0, num)
	step := root(stop/start, num)
	if stopIncluded {
		step = root(stop/start, num-1)
	}

	for i := 0; i < num; i++ {
		result = append(result, start)
		start *= step
	}

	return Float64OneDArray{result, step, nil}
}

func root(a float64, n int) float64 {
	n1 := n - 1
	n1f, rn := float64(n1), 1/float64(n)
	x, x0 := 1., 0.
	for {
		potx, t2 := 1/x, a
		for b := n1; b > 0; b >>= 1 {
			if b&1 == 1 {
				t2 *= potx
			}
			potx *= potx
		}
		x0, x = x, rn*(n1f*x+t2)
		if math.Abs(x-x0)*1e15 < x {
			break
		}
	}
	return x
}
