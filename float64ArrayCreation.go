package np

import (
	"errors"
	"math"
)

func Linspace(start float64, stop float64, num int, stopIncluded bool) Float64OneDArray {
	if num < 2 {
		return Float64OneDArray{nil, nil, errors.New(ErrInvalidParameter)}
	}

	if start == stop {
		return Float64OneDArray{nil, nil, errors.New(ErrInvalidParameter)}
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

	return Float64OneDArray{result, float64Pointer(step), nil}
}

func Geomspace(start float64, stop float64, num int, stopIncluded bool) Float64OneDArray {
	if start == 0 {
		return Float64OneDArray{nil, nil, errors.New(ErrZeroValue)}
	}

	if start < 0 || stop < 0 {
		return Float64OneDArray{nil, nil, errors.New(ErrNegativeValue)}
	}

	if num < 2 {
		return Float64OneDArray{nil, nil, errors.New(ErrInvalidParameter)}
	}

	if start == stop {
		return Float64OneDArray{nil, nil, errors.New(ErrInvalidParameter)}
	}

	result := make([]float64, num)
	n := num
	if stopIncluded {
		n -= 1
	}
	step := root(stop/start, n)

	for i := 0; i < num; i++ {
		result[i] = start
		start *= step
	}

	return Float64OneDArray{result, float64Pointer(step), nil}
}

func Logspace(start float64, stop float64, num int, base float64, stopIncluded bool) Float64OneDArray {
	if stop <= start {
		return Float64OneDArray{nil, nil, errors.New(ErrInvalidParameter)}
	}

	start = math.Pow(base, start)
	stop = math.Pow(base, stop)

	return Geomspace(start, stop, num, stopIncluded)
}

func Empty(n int) Float64TwoDArray {
	if n <= 0 {
		return Float64TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
	}

	return Float64TwoDArray{result, nil}
}

func EmptyLike(a Float64TwoDArray) Float64TwoDArray {
	if err := validateArray(a.err, len(a.arr)); err != nil {
		return Float64TwoDArray{nil, err}
	}

	shape, _ := a.Shape()
	result := make([][]float64, shape.NRow)
	for i := 0; i < len(result); i++ {
		result[i] = make([]float64, shape.NCol)
	}

	return Float64TwoDArray{result, nil}
}

func Identity(n int) Float64TwoDArray {
	if n <= 0 {
		return Float64TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = 0
			if j == i {
				result[i][j] = 1
			}
		}
	}

	return Float64TwoDArray{result, nil}
}

func Eye(n int) Float64TwoDArray {
	if n <= 0 {
		return Float64TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = 0
			if j == i || j == n-1-i {
				result[i][j] = 1
			}
		}
	}

	return Float64TwoDArray{result, nil}
}

func Full(nRow int, nCol int, fill_value float64) Float64TwoDArray {
	if nRow <= 0 || nCol <= 0 {
		return Float64TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	result := make([][]float64, nRow)
	for i := 0; i < nRow; i++ {
		result[i] = make([]float64, nCol)
		for j := 0; j < nCol; j++ {
			result[i][j] = fill_value
		}
	}

	return Float64TwoDArray{result, nil}
}

func FullLike(a Float64TwoDArray, fill_value float64) Float64TwoDArray {
	if err := validateArray(a.err, len(a.arr)); err != nil {
		return Float64TwoDArray{nil, err}
	}

	shape, _ := a.Shape()
	result := make([][]float64, shape.NRow)
	for i := 0; i < len(result); i++ {
		result[i] = make([]float64, shape.NCol)
		for j := 0; j < len(result[i]); j++ {
			result[i][j] = fill_value
		}
	}

	return Float64TwoDArray{result, nil}
}

func Ones(nRow int, nCol int) Float64TwoDArray {
	return Full(nRow, nCol, 1)
}

func OnesLike(a Float64TwoDArray) Float64TwoDArray {
	return FullLike(a, 1)
}

func Zeros(nRow int, nCol int) Float64TwoDArray {
	return Full(nRow, nCol, 0)
}

func ZerosLike(a Float64TwoDArray) Float64TwoDArray {
	return FullLike(a, 0)
}
