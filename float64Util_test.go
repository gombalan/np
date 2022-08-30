package np

import (
	"testing"
)

func TestLenFloat64(t *testing.T) {
	len, _ := Linspace(0, 4, 5, false).Len()
	if *len != 5 {
		t.Errorf("Length of array must be %d", 5)
	}

	if _, err := Linspace(4, 4, 3, false).Len(); err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}
}

func TestShapeFloat64(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)

	shape, _ := res.Shape()
	if *shape.NCol != 3 || *shape.NRow != 2 {
		t.Errorf("Shape of array must be [%d, %d]", shape.NCol, shape.NRow)
	}

	res = arr.Reshape(3, 3)
	if _, err := res.Shape(); err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}
}

func TestSizeFloat64(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)
	if size, _ := res.Size(); size != nil && *size != 6 {
		t.Errorf("Size should be %d", 6)
	}

	res = arr.Reshape(4, 4)
	if _, err := res.Size(); err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}
}
