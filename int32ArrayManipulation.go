package np

import "errors"

func (a Int32OneDArray) Reshape(nRow int, nCol int) Int32TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Int32TwoDArray{nil, err}
	}

	if nRow < 1 || nCol < 1 {
		return Int32TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	if dim, _ := a.Len(); dim != nil && *dim != nRow*nCol {
		return Int32TwoDArray{nil, errors.New(ErrSizeNotMatch)}
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
