package np

type Shape2D struct {
	NRow *int
	NCol *int
}

func (a Float64OneDArray) Len() (*int, error) {
	if a.Err != nil {
		return nil, propagateError(a.Err, "failed to determine array's length")
	}
	length := len(a.Arr)

	return &length, nil
}

func (a Float64TwoDArray) Shape() (Shape2D, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Shape2D{}, propagateError(err, "failed to determine array's shape")
	}

	shape := Shape2D{}
	nCol, nRow := len(a.Arr[0]), len(a.Arr)
	shape.NCol = &nCol
	shape.NRow = &nRow

	return shape, nil
}

func (a Float64TwoDArray) Size() (*int, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, propagateError(err, "failed to determine array's size")
	}

	shape, _ := a.Shape()
	size := *shape.NCol * *shape.NRow

	return &size, nil
}
