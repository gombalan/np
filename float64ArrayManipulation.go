package np

import (
	"fmt"
)

func (a Float64OneDArray) Reshape(nRow int, nCol int) Float64TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64TwoDArray{nil, propagateError(err, "failed to reshape array")}
	}

	if nRow <= 0 || nCol <= 0 {
		return Float64TwoDArray{nil, newError(ErrZeroOrNegativeValue, fmt.Sprintf("value of nRow: %v and nCol: %v should be greater than zero", nRow, nCol))}
	}

	if len, _ := a.Len(); len != nil && *len != nRow*nCol {
		return Float64TwoDArray{nil, newError(ErrSizeNotMatch, fmt.Sprintf("value of nRow: %v times nCol: %v should be equal to original array's dimension", nRow, nCol))}
	}

	result := make([][]float64, nRow)
	for i := 0; i < nRow; i++ {
		result[i] = make([]float64, nCol)
		for j := 0; j < nCol; j++ {
			result[i][j] = a.Arr[i*nCol+j]
		}
	}

	return Float64TwoDArray{result, nil}
}

func (a Float64TwoDArray) Flatten() Float64OneDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64OneDArray{nil, propagateError(err, "failed to flatten array")}
	}

	shape, _ := a.Shape()
	size, _ := a.Size()

	result := make([]float64, *size)
	for i := 0; i < *size; i++ {
		result[i] = a.Arr[i/(*shape.NCol)][i%(*shape.NCol)]
	}

	return Float64OneDArray{result, nil}
}

func (a Float64TwoDArray) Reshape(nRow int, nCol int) Float64TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64TwoDArray{nil, propagateError(err, "failed to reshape array")}
	}

	if nRow <= 0 || nCol <= 0 {
		return Float64TwoDArray{nil, newError(ErrZeroOrNegativeValue, fmt.Sprintf("value of nRow: %v and nCol: %v should be greater than zero", nRow, nCol))}
	}

	if size, _ := a.Size(); size != nil && *size != nRow*nCol {
		return Float64TwoDArray{nil, newError(ErrSizeNotMatch, fmt.Sprintf("value of nRow: %v times nCol: %v should be equal to original array's dimension", nRow, nCol))}
	}

	result := a.Flatten().Reshape(nRow, nCol)

	return result
}

func (a Float64TwoDArray) Transpose() Float64TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64TwoDArray{nil, propagateError(err, "failed to transpose array")}
	}

	shape, _ := a.Shape()
	result := make([][]float64, *shape.NCol)
	for i := 0; i < int(*shape.NCol); i++ {
		result[i] = make([]float64, *shape.NRow)
		for j := 0; j < *shape.NRow; j++ {
			result[i][j] = a.Arr[j][i]
		}
	}

	return Float64TwoDArray{result, nil}
}
