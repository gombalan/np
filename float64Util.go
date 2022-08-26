package np

import (
	"errors"
)

type Shape2D struct {
	NRow int32
	NCol int32
}

func (a Float64OneDArray) Len() (int32, error) {
	if a.err != nil {
		return 0, a.err
	}

	return int32(len(a.arr)), nil
}

func (a Float64TwoDArray) Shape() (Shape2D, error) {
	if err := validateArray(a.err, len(a.arr)); err != nil {
		return Shape2D{}, err
	}

	shape := Shape2D{}
	shape.NCol = int32(len(a.arr[0]))
	shape.NRow = int32(len(a.arr))

	return shape, nil
}

func (a Float64TwoDArray) Size() (int32, error) {
	if err := validateArray(a.err, len(a.arr)); err != nil {
		return 0, err
	}

	shape, _ := a.Shape()
	size := shape.NCol * shape.NRow

	return size, nil
}

func (a Float64OneDArray) Reshape(nRow int, nCol int) Float64TwoDArray {
	if err := validateArray(a.err, len(a.arr)); err != nil {
		return Float64TwoDArray{nil, err}
	}

	if nRow < 1 || nCol < 1 {
		return Float64TwoDArray{nil, errors.New(ErrNegativeValue)}
	}

	if dim, _ := a.Len(); int(dim) != nRow*nCol {
		return Float64TwoDArray{nil, errors.New(ErrSizeNotMatch)}
	}

	result := make([][]float64, nRow)
	for i := 0; i < nRow; i++ {
		row := make([]float64, nCol)
		for j := 0; j < nCol; j++ {
			row[j] = a.arr[i*nCol+j]
		}
		result[i] = row
	}

	return Float64TwoDArray{result, nil}
}

func (a Float64TwoDArray) Flatten() Float64OneDArray {
	if err := validateArray(a.err, len(a.arr)); err != nil {
		return Float64OneDArray{nil, nil, err}
	}

	shape, _ := a.Shape()
	size, _ := a.Size()

	result := make([]float64, size)
	for i := 0; i < int(size); i++ {
		result[i] = a.arr[i/int(shape.NCol)][i%int(shape.NCol)]
	}

	return Float64OneDArray{result, nil, nil}
}

func (a Float64TwoDArray) Reshape(nRow int, nCol int) Float64TwoDArray {
	if err := validateArray(a.err, len(a.arr)); err != nil {
		return Float64TwoDArray{nil, err}
	}

	if nRow < 1 || nCol < 1 {
		return Float64TwoDArray{nil, errors.New(ErrInvalidParameter)}
	}

	if size, _ := a.Size(); int(size) != nRow*nCol {
		return Float64TwoDArray{nil, errors.New(ErrSizeNotMatch)}
	}

	result := a.Flatten().Reshape(nRow, nCol)

	return result
}
