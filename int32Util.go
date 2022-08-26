package np

import "errors"

func (a Int32OneDArray) Len() (int32, error) {
	if a.err != nil {
		return 0, a.err
	}

	return int32(len(a.arr)), nil
}

func (a Int32OneDArray) Reshape(nRow int, nCol int) Int32TwoDArray {
	if err := validateArray(a.err, len(a.arr)); err != nil {
		return Int32TwoDArray{nil, err}
	}

	if nRow < 1 || nCol < 1 {
		return Int32TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	if dim, _ := a.Len(); int(dim) != nRow*nCol {
		return Int32TwoDArray{nil, errors.New(ErrSizeNotMatch)}
	}

	result := make([][]int32, nRow)
	for i := 0; i < nRow; i++ {
		row := make([]int32, nCol)
		for j := 0; j < nCol; j++ {
			row[j] = a.arr[i*nCol+j]
		}
		result[i] = row
	}

	return Int32TwoDArray{result, nil}
}
