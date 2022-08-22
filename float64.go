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

type Float64TwoDArray struct {
	arr [][]float64
	err error
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
	n := num
	if stopIncluded {
		n -= 1
	}
	step := root(stop/start, n)

	for i := 0; i < num; i++ {
		result = append(result, start)
		start *= step
	}

	return Float64OneDArray{result, step, nil}
}

func Logspace(start float64, stop float64, num int, base float64, stopIncluded bool) Float64OneDArray {
	if stop <= start {
		return Float64OneDArray{nil, 0, errors.New(ErrInvalidParameter)}
	}

	start = math.Pow(base, start)
	stop = math.Pow(base, stop)

	return Geomspace(start, stop, num, stopIncluded)
}

func Identity(n int) Float64TwoDArray {
	if n <= 0 {
		return Float64TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	result := make([][]float64, 0, n)

	for i := 0; i < n; i++ {
		row := make([]float64, n)
		for j := 0; j < n; j++ {
			row[j] = 0
			if j == i {
				row[j] = 1
			}
		}

		result = append(result, row)
	}

	return Float64TwoDArray{result, nil}
}

func Full(nRow int, nCol int, num float64) Float64TwoDArray {
	if nRow <= 0 || nCol <= 0 {
		return Float64TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	result := make([][]float64, 0, nRow)

	for i := 0; i < nRow; i++ {
		row := make([]float64, nCol)
		for j := 0; j < nCol; j++ {
			row[j] = num
		}

		result = append(result, row)
	}

	return Float64TwoDArray{result, nil}
}

func Ones(nRow int, nCol int) Float64TwoDArray {
	return Full(nRow, nCol, 1)
}

func Zeros(nRow int, nCol int) Float64TwoDArray {
	return Full(nRow, nCol, 0)
}
