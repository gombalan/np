package np

import (
	"testing"
)

func TestReshapeFloat64OneD(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if res.Arr[i][j] != arr.Arr[i*3+j] {
				t.Errorf("Element at row %d colum %d must be %f", i, j, arr.Arr[i*3+j])
			}
		}
	}

	arr = Linspace(3, 3, 3, false)
	if arr.Reshape(2, 3).Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	arr = Linspace(3, 4, 6, false)
	if arr.Reshape(-2, -3).Err == nil {
		t.Errorf("There must be ErrNegativeValue")
	}

	arr = Linspace(3, 4, 6, false)
	if arr.Reshape(2, 2).Err == nil {
		t.Errorf("There must be ErrSizeNotMatch")
	}
}

func TestFlattenFloat64(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3).Flatten()

	if len, _ := res.Len(); len != nil && *len != 6 {
		t.Errorf("Len of the flattened array must be %d", 6)
	}

	for i := range res.Arr {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element of flattened array at index %d must be %f", i, arr.Arr[i])
		}
	}

	res = arr.Reshape(3, 3).Flatten()
	if _, err := res.Len(); err == nil {
		t.Errorf("There must be ErrSizeNotMatch")
	}
}

func TestReshapeFloat64TwoD(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)
	res = res.Reshape(3, 2)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if res.Arr[i][j] != arr.Arr[i*2+j] {
				t.Errorf("Element at row %d colum %d must be %f", i, j, arr.Arr[i*2+j])
			}
		}
	}

	res = arr.Reshape(-1, -2)
	if res.Reshape(2, 3).Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	res = arr.Reshape(2, 3)
	if res.Reshape(-2, -3).Err == nil {
		t.Errorf("There must be ErrNegativeValue")
	}

	if res.Reshape(3, 3).Err == nil {
		t.Errorf("There must be ErrNegativeValue")
	}
}

func TestTransposeFloat64TwoD(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3).Transpose().Flatten()
	arr = Float64OneDArray{Arr: []float64{1, 3, 1, 4, 1, 7}}

	for i := 0; i < len(arr.Arr); i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	res = arr.Reshape(3, 3).Transpose().Flatten()
	if res.Err == nil {
		t.Errorf("There must be ErrSizeNotMatch")
	}
}
