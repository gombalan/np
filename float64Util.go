package np

type Shape2D struct {
	NRow *int
	NCol *int
}

func (a Float64OneDArray) Len() (*int, error) {
	if a.Err != nil {
		return nil, a.Err
	}

	return intPointer(len(a.Arr)), nil
}

func (a Float64TwoDArray) Shape() (Shape2D, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return Shape2D{}, err
	}

	shape := Shape2D{}
	shape.NCol = intPointer(len(a.Arr[0]))
	shape.NRow = intPointer(len(a.Arr))

	return shape, nil
}

func (a Float64TwoDArray) Size() (*int, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	shape, _ := a.Shape()
	size := *shape.NCol * *shape.NRow

	return intPointer(size), nil
}
