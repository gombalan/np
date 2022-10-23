package np

import (
	"fmt"
)

func (a Float64OneDArray) Flip() Float64OneDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64OneDArray{nil, propagateError(err, "failed to roll array")}
	}

	len, _ := a.Len()
	result := make([]float64, *len)
	for i := 0; i < *len; i++ {
		result[i] = a.Arr[*len-1-i]
	}

	return Float64OneDArray{Arr: result}
}

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

func (a Float64OneDArray) Roll(k int) Float64OneDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64OneDArray{nil, propagateError(err, "failed to roll array")}
	}

	len, _ := a.Len()
	k = k % (*len)
	if k == 0 {
		return a
	}

	result := make([]float64, *len)
	for i := 0; i < *len; i++ {
		if i-k < 0 {
			result[i] = a.Arr[(i-k+*len)%(*len)]
		} else {
			result[i] = a.Arr[(i-k)%(*len)]
		}
	}

	return Float64OneDArray{Arr: result}
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

func (a Float64TwoDArray) Diag(k int) Float64OneDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64OneDArray{nil, propagateError(err, "failed to extract array's diagonal")}
	}

	shape, _ := a.Shape()
	if *shape.NCol != *shape.NRow {
		return Float64OneDArray{nil, newError(ErrNonRectangularArray, fmt.Sprintf("array should be rectangular, but value of nRow: %v and nCol: %v are different", *shape.NRow, *shape.NCol))}
	}

	start, stop, size := 0, *shape.NRow, *shape.NRow

	if k > 0 {
		if k >= *shape.NCol {
			return Float64OneDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of k: %v should be less than nRow: %v", k, *shape.NRow))}
		}

		start, stop, size = 0, *shape.NRow-k, *shape.NRow-k
	}

	if k < 0 {
		if -k >= *shape.NCol {
			return Float64OneDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of -k: %v should be less than nRow: %v", k, -1*(*shape.NRow)))}
		}

		start, stop, size = -k, *shape.NRow, *shape.NRow+k
	}

	res := make([]float64, size)
	for i := start; i < stop; i++ {
		res[i-start] = a.Arr[i][i+k]
	}

	return Float64OneDArray{res, nil}
}

func (a Float64TwoDArray) Triu(k int) Float64TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64TwoDArray{nil, propagateError(err, "failed to create upper triangle of array")}
	}

	shape, _ := a.Shape()
	if *shape.NCol != *shape.NRow {
		return Float64TwoDArray{nil, newError(ErrNonRectangularArray, fmt.Sprintf("array should be rectangular, but value of nRow: %v and nCol: %v are different", *shape.NRow, *shape.NCol))}
	}

	if k > 0 {
		if k >= *shape.NCol {
			return Float64TwoDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of k: %v should be less than nRow: %v", k, *shape.NRow))}
		}
	}

	if k < 0 {
		if -k >= *shape.NCol {
			return Float64TwoDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of -k: %v should be less than nRow: %v", k, -1*(*shape.NRow)))}
		}
	}

	result := make([][]float64, *shape.NRow)
	for i := 0; i < *shape.NRow; i++ {
		result[i] = make([]float64, *shape.NCol)
		for j := 0; j < *shape.NCol; j++ {
			if (i-k >= 0 || i == 0) && j < i+k {
				result[i][j] = 0
			} else {
				result[i][j] = a.Arr[i][j]
			}
		}
	}

	return Float64TwoDArray{result, nil}
}

func (a Float64TwoDArray) Tril(k int) Float64TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64TwoDArray{nil, propagateError(err, "failed to create upper triangle of array")}
	}

	shape, _ := a.Shape()
	if *shape.NCol != *shape.NRow {
		return Float64TwoDArray{nil, newError(ErrNonRectangularArray, fmt.Sprintf("array should be rectangular, but value of nRow: %v and nCol: %v are different", *shape.NRow, *shape.NCol))}
	}

	if k > 0 && k >= *shape.NCol {
		return Float64TwoDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of k: %v should be less than nRow: %v", k, *shape.NRow))}
	}

	if k < 0 && -k >= *shape.NCol {
		return Float64TwoDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of -k: %v should be less than nRow: %v", k, -1*(*shape.NRow)))}
	}

	result := make([][]float64, *shape.NRow)
	for i := 0; i < *shape.NRow; i++ {
		result[i] = make([]float64, *shape.NCol)
		for j := 0; j < *shape.NCol; j++ {
			if (i-k >= 0 || i == 0) && j <= i+k {
				result[i][j] = a.Arr[i][j]
			} else {
				result[i][j] = 0
			}
		}
	}

	return Float64TwoDArray{result, nil}
}

func (a Float64TwoDArray) Roll(j int, k int) Float64TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64TwoDArray{nil, propagateError(err, "failed to roll array")}
	}

	shape, _ := a.Shape()
	j, k = j%(*shape.NRow), k%(*shape.NCol)
	if j == 0 && k == 0 {
		return a
	}

	result := make([][]float64, *shape.NRow)
	for m := 0; m < *shape.NRow; m++ {
		result[m] = make([]float64, *shape.NCol)
		for n := 0; n < *shape.NCol; n++ {
			if m-j < 0 {
				if n-k < 0 {
					result[m][n] = a.Arr[(m-j+*shape.NRow)%(*shape.NRow)][(n-k+*shape.NCol)%(*shape.NCol)]
				} else {
					result[m][n] = a.Arr[(m-j+*shape.NRow)%(*shape.NRow)][(n-k)%(*shape.NCol)]
				}
			} else {
				if n-k < 0 {
					result[m][n] = a.Arr[(m-j)%(*shape.NRow)][(n-k+*shape.NCol)%(*shape.NCol)]
				} else {
					result[m][n] = a.Arr[(m-j)%(*shape.NRow)][(n-k)%(*shape.NCol)]
				}
			}
		}
	}

	return Float64TwoDArray{result, nil}
}

func (a Float64TwoDArray) Rot90(k int, fromAxis int, toAxis int) Float64TwoDArray {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Float64TwoDArray{nil, propagateError(err, "failed to rotate array")}
	}

	if fromAxis == toAxis || fromAxis < 0 || fromAxis > 1 || toAxis < 0 || toAxis > 1 {
		return Float64TwoDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of fromAxis: %v and toAxis: %v should be different and be either 0 or 1", fromAxis, toAxis))}
	}

	shape, _ := a.Shape()
	nRow, nCol := *shape.NRow, *shape.NCol
	if k%2 == 1 || k%2 == -1 {
		nRow, nCol = *shape.NCol, *shape.NRow
	}

	if fromAxis > toAxis {
		k = -k
	}

	k = k % 4
	if k < 0 {
		k = k + 4
	}

	switch k {
	case 1:
		result := make([][]float64, nRow)
		for i := 0; i < nRow; i++ {
			result[i] = make([]float64, nCol)
			for j := 0; j < nCol; j++ {
				result[i][j] = a.Arr[j][nCol-i]
			}
		}

		return Float64TwoDArray{result, nil}
	case 2:
		result := make([][]float64, nRow)
		for i := 0; i < nRow; i++ {
			result[i] = make([]float64, nCol)
			for j := 0; j < nCol; j++ {
				result[i][j] = a.Arr[nRow-i-1][nCol-j-1]
			}
		}

		return Float64TwoDArray{result, nil}
	case 3:
		result := make([][]float64, nRow)
		for i := 0; i < nRow; i++ {
			result[i] = make([]float64, nCol)
			for j := 0; j < nCol; j++ {
				result[i][j] = a.Arr[nCol-j-1][i]
			}
		}

		return Float64TwoDArray{result, nil}
	default:
		return a
	}
}
