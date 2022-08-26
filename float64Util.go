package np

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
