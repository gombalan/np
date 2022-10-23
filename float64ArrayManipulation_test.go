package np

import (
	"fmt"
	"math"
	"testing"
)

func TestTrimZerosFloat64OneD(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{0, 0, 1, 3, 4, 7}}
	res := Float64OneDArray{Arr: []float64{1, 3, 4, 7}}
	arr = arr.TrimZeros("f")

	for i := 0; i < len(res.Arr); i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	arr = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 0, 0}}
	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3}}
	arr = arr.TrimZeros("b")

	for i := 0; i < len(res.Arr); i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	arr = Float64OneDArray{Arr: []float64{0, 0, 1, 3, 0, 0}}
	res = Float64OneDArray{Arr: []float64{1, 3}}
	arr = arr.TrimZeros("fb")

	for i := 0; i < len(res.Arr); i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	if arr.TrimZeros("k").Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	arr = Float64OneDArray{}
	if arr.TrimZeros("f").Err == nil {
		t.Errorf("There must be ErrEmptyArray")
	}
}

func TestFlipFloat64OneD(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Flip()
	arr = Float64OneDArray{Arr: []float64{7, 4, 3, 1, 1, 1}}

	for i := 0; i < 6; i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	arr = Float64OneDArray{}
	if arr.Flip().Err == nil {
		t.Errorf("There must be ErrEmptyArray")
	}
}

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

func TestRollFloat64OneD(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 2, 3, 4, 5, 6}}
	res := arr.Roll(2)
	arr = Float64OneDArray{Arr: []float64{5, 6, 1, 2, 3, 4}}
	for i := 0; i < 6; i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	arr = Float64OneDArray{Arr: []float64{1, 2, 3, 4, 5, 6}}
	arr = arr.Roll(8)
	for i := 0; i < 6; i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	res = arr.Roll(-2)
	arr = Float64OneDArray{Arr: []float64{1, 2, 3, 4, 5, 6}}
	for i := 0; i < 6; i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	res = arr.Roll(0)
	arr = arr.Roll(6)
	for i := 0; i < 6; i++ {
		if res.Arr[i] != arr.Arr[i] {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	arr = Float64OneDArray{}
	if arr.Roll(3).Err == nil {
		t.Errorf("There must be ErrEmptyArray")
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

func TestDiagFloat64TwoD(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}
	res := arr.Reshape(3, 3).Diag(0)
	arr = Float64OneDArray{Arr: []float64{1, 4, 4}}

	for i := 0; i < len(arr.Arr); i++ {
		if math.Abs(res.Arr[i]-arr.Arr[i]) > 1e-9 {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	arr = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}
	res = arr.Reshape(3, 3).Diag(1)
	arr = Float64OneDArray{Arr: []float64{1, 7}}

	for i := 0; i < len(arr.Arr); i++ {
		if math.Abs(res.Arr[i]-arr.Arr[i]) > 1e-9 {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	arr = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}
	res = arr.Reshape(3, 3).Diag(-1)
	arr = Float64OneDArray{Arr: []float64{3, 6}}

	for i := 0; i < len(arr.Arr); i++ {
		if math.Abs(res.Arr[i]-arr.Arr[i]) > 1e-9 {
			t.Errorf("Element at index %d must be %f", i, arr.Arr[i])
		}
	}

	arr = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}
	res = arr.Reshape(4, 3).Diag(0)
	if res.Err == nil {
		t.Errorf("There must be ErrSizeNotMatch")
	}

	res = arr.Reshape(3, 3).Diag(4)
	if res.Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	res = arr.Reshape(3, 3).Diag(-4)
	if res.Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	arr = Float64OneDArray{Arr: []float64{1, 7, 1, 8, 1, 0, 3, 4, 7, 2, 6, 4}}
	res = arr.Reshape(4, 3).Diag(0)
	if res.Err == nil {
		t.Errorf("There must be ErrNonRectangularArray")
	}
}

func TestTriuFloat64TwoD(t *testing.T) {
	res := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Triu(0)
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 0, 4, 7, 0, 0, 4}}.Reshape(3, 3)

	nRow, nCol := len(arr.Arr), len(arr.Arr)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Triu(1)
	arr = Float64OneDArray{Arr: []float64{0, 1, 1, 0, 0, 7, 0, 0, 0}}.Reshape(3, 3)

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res, arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Triu(-1)
	arr = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 0, 6, 4}}.Reshape(3, 3)

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res, arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(4, 3).Triu(0)
	if res.Err == nil {
		t.Errorf("There must be ErrSizeNotMatch")
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4, 7, 5, 8}}.Reshape(4, 3).Triu(0)
	if res.Err == nil {
		t.Errorf("There must be ErrNonRectangularArray")
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Triu(-5)
	if res.Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Triu(5)
	if res.Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}
}

func TestTrilFloat64TwoD(t *testing.T) {
	res := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Tril(0)
	arr := Float64OneDArray{Arr: []float64{1, 0, 0, 3, 4, 0, 2, 6, 4}}.Reshape(3, 3)

	nRow, nCol := len(arr.Arr), len(arr.Arr)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Tril(1)
	arr = Float64OneDArray{Arr: []float64{1, 1, 0, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3)

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Tril(-1)
	arr = Float64OneDArray{Arr: []float64{0, 0, 0, 3, 0, 0, 2, 6, 0}}.Reshape(3, 3)

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(4, 3).Tril(0)
	if res.Err == nil {
		t.Errorf("There must be ErrSizeNotMatch")
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4, 7, 5, 8}}.Reshape(4, 3).Tril(0)
	if res.Err == nil {
		t.Errorf("There must be ErrNonRectangularArray")
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Tril(-5)
	if res.Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	res = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7, 2, 6, 4}}.Reshape(3, 3).Tril(5)
	if res.Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}
}

func TestFlipFloat64TwoD(t *testing.T) {
	res := Float64OneDArray{Arr: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}.Reshape(3, 3).Flip(-1)
	arr := Float64OneDArray{Arr: []float64{9, 8, 7, 6, 5, 4, 3, 2, 1}}.Reshape(3, 3)

	nRow, nCol := len(arr.Arr), len(arr.Arr)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = arr.Flip(0)
	arr = Float64OneDArray{Arr: []float64{3, 2, 1, 6, 5, 4, 9, 8, 7}}.Reshape(3, 3)

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = arr.Flip(1)
	arr = Float64OneDArray{Arr: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}.Reshape(3, 3)

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	if arr.Flip(2).Err == nil {
		t.Errorf("There must be ErrInvalidParameter")
	}

	arr = Float64TwoDArray{}
	if arr.Flip(0).Err == nil {
		t.Errorf("There must be ErrEmptyArray")
	}
}

func TestFliplrFloat64TwoD(t *testing.T) {
	res := Float64OneDArray{Arr: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}.Reshape(3, 3).Fliplr()
	arr := Float64OneDArray{Arr: []float64{3, 2, 1, 6, 5, 4, 9, 8, 7}}.Reshape(3, 3)

	nRow, nCol := len(arr.Arr), len(arr.Arr)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}
}

func TestFlipudFloat64TwoD(t *testing.T) {
	res := Float64OneDArray{Arr: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}.Reshape(3, 3).Flipud()
	arr := Float64OneDArray{Arr: []float64{7, 8, 9, 4, 5, 6, 1, 2, 3}}.Reshape(3, 3)

	nRow, nCol := len(arr.Arr), len(arr.Arr)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}
}

func TestRollFloat64TwoD(t *testing.T) {
	res := Float64OneDArray{Arr: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}.Reshape(3, 3).Roll(1, 1)
	arr := Float64OneDArray{Arr: []float64{9, 7, 8, 3, 1, 2, 6, 4, 5}}.Reshape(3, 3)

	nRow, nCol := len(arr.Arr), len(arr.Arr)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = arr.Roll(3, 6)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				fmt.Println(res.Arr, arr.Arr)
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	arr = Float64TwoDArray{}
	if arr.Roll(1, 2).Err == nil {
		t.Errorf("There must be ErrEmptyArray")
	}
}

func TestRot90TwoD(t *testing.T) {
	res := Float64OneDArray{Arr: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}.Reshape(3, 4).Rot90(1, 0, 1)
	res1 := Float64OneDArray{Arr: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}.Reshape(3, 4).Rot90(3, 1, 0)
	arr := Float64OneDArray{Arr: []float64{3, 7, 11, 2, 6, 10, 1, 5, 9, 0, 4, 8}}.Reshape(4, 3)

	shape, _ := arr.Shape()
	nRow, nCol := *shape.NRow, *shape.NCol

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}

			if math.Abs(res1.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = Float64OneDArray{Arr: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}.Reshape(3, 4).Rot90(2, 0, 1)
	res1 = Float64OneDArray{Arr: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}.Reshape(3, 4).Rot90(2, 1, 0)
	arr = Float64OneDArray{Arr: []float64{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}.Reshape(3, 4)

	shape, _ = arr.Shape()
	nRow, nCol = *shape.NRow, *shape.NCol

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}

			if math.Abs(res1.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}

	res = Float64OneDArray{Arr: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}.Reshape(3, 4).Rot90(3, 0, 1)
	res1 = Float64OneDArray{Arr: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}.Reshape(3, 4).Rot90(1, 1, 0)
	arr = Float64OneDArray{Arr: []float64{8, 4, 0, 9, 5, 1, 10, 6, 2, 11, 7, 3}}.Reshape(4, 3)

	shape, _ = arr.Shape()
	nRow, nCol = *shape.NRow, *shape.NCol

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if math.Abs(res.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}

			if math.Abs(res1.Arr[i][j]-arr.Arr[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and column %d must be %f", i, j, arr.Arr[i][j])
			}
		}
	}
}
