package np

import (
	"testing"
)

func TestLenFloat64(t *testing.T) {
	res, _ := Linspace(0, 4, 5, false).Len()
	if res != 5 {
		t.Errorf("Length of array must be %d", 5)
	}

	if _, err := Linspace(4, 4, 3, false).Len(); err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}
}

func TestShapeFloat64(t *testing.T) {
	arr := Float64OneDArray{arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)

	shape, _ := res.Shape()
	if shape.NCol != 3 || shape.NRow != 2 {
		t.Errorf("Shape of array must be [%d, %d]", shape.NCol, shape.NRow)
	}

	res = arr.Reshape(3, 3)
	if _, err := res.Shape(); err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}
}

func TestSizeFloat64(t *testing.T) {
	arr := Float64OneDArray{arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)
	if size, _ := res.Size(); size != 6 {
		t.Errorf("Size should be %d", 6)
	}

	res = arr.Reshape(4, 4)
	if _, err := res.Size(); err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}
}

func TestReshapeFloat64OneD(t *testing.T) {
	arr := Float64OneDArray{arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if res.arr[i][j] != arr.arr[i*3+j] {
				t.Errorf("Element at row %d colum %d must be %f", i, j, arr.arr[i*3+j])
			}
		}
	}

	arr = Linspace(3, 3, 3, false)
	if arr.Reshape(2, 3).err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	arr = Linspace(3, 4, 6, false)
	if arr.Reshape(-2, -3).err == nil {
		t.Errorf("There must be ErrNegativeValue")
	}

	arr = Linspace(3, 4, 6, false)
	if arr.Reshape(2, 2).err == nil {
		t.Errorf("There must be ErrSizeNotMatch")
	}
}

func TestFlattenFloat64(t *testing.T) {
	arr := Float64OneDArray{arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3).Flatten()

	if len, _ := res.Len(); len != 6 {
		t.Errorf("Len of the flattened array must be %d", 6)
	}

	for i := range res.arr {
		if res.arr[i] != arr.arr[i] {
			t.Errorf("Element of flattened array at index %d must be %f", i, arr.arr[i])
		}
	}

	res = arr.Reshape(3, 3).Flatten()
	if _, err := res.Len(); err == nil {
		t.Errorf("There must be ErrSizeNotMatch")
	}
}

func TestReshapeFloat64TwoDay(t *testing.T) {
	arr := Float64OneDArray{arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)
	res = res.Reshape(3, 2)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if res.arr[i][j] != arr.arr[i*2+j] {
				t.Errorf("Element at row %d colum %d must be %f", i, j, arr.arr[i*2+j])
			}
		}
	}

	res = arr.Reshape(-1, -2)
	if res.Reshape(2, 3).err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	res = arr.Reshape(2, 3)
	if res.Reshape(-2, -3).err == nil {
		t.Errorf("There must be ErrNegativeValue")
	}

	if res.Reshape(3, 3).err == nil {
		t.Errorf("There must be ErrNegativeValue")
	}
}
