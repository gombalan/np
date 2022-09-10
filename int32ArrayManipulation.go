package np

import (
	"fmt"
)

func (a Int32OneDArray) Reshape(nRow int, nCol int) Int32TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Int32TwoDArray{nil, propagateError(err, "failed to reshape array")}
	}

	if nRow <= 0 || nCol <= 0 {
		return Int32TwoDArray{nil, newError(ErrZeroOrNegativeValue, fmt.Sprintf("both value of nRow: %v and nCol: %v should be positive", nRow, nCol))}
	}

	if dim, _ := a.Len(); dim != nil && *dim != nRow*nCol {
		return Int32TwoDArray{nil, newError(ErrSizeNotMatch, fmt.Sprintf("value of nRow: %v times nCol: %v should be equal to original array's dimension", nRow, nCol))}
	}

	result := make([][]int32, nRow)
	for i := 0; i < nRow; i++ {
		row := make([]int32, nCol)
		for j := 0; j < nCol; j++ {
			row[j] = a.Arr[i*nCol+j]
		}
		result[i] = row
	}

	return Int32TwoDArray{result, nil}
}
