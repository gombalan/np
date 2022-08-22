package np

import (
	"errors"
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
