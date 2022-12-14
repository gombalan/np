package np

import (
	"fmt"
	"math"
	"testing"
)

func TestArange(t *testing.T) {
	res := Arange(0, 4, 1)
	for i := range res.Arr {
		if int(res.Arr[i]) != i {
			t.Errorf("Element at %d != %d", i, i)
		}
	}

	res = Arange(1, 4, 1)
	for i := range res.Arr {
		if int(res.Arr[i]) != i+1 {
			t.Errorf("Element at %d != %d", i, i+1)
		}
	}

	res = Arange(1, 4, 2)
	for i := range res.Arr {
		if int(res.Arr[i]) != i*2+1 {
			t.Errorf("Element at %d != %d", i, 2*i+1)
		}
	}

	res = Arange(1, 1, 0)
	if res.Err == nil {
		t.Errorf("There must be an ErrZeroValue")
	}

	res = Arange(1, 1, 2)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	res = Arange(1, -3, 2)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestLinspace(t *testing.T) {
	res := Linspace(0, 4, 5, false)
	array := []float64{0, 0.8, 1.6, 2.4, 3.2}
	for i := range res.Arr {
		if math.Abs(res.Arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Linspace(1, 4, 4, true)
	array = []float64{1, 2, 3, 4}
	for i := range res.Arr {
		if math.Abs(res.Arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Linspace(4, 4, 2, true)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	res = Linspace(4, 4, 0, false)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestGeomspace(t *testing.T) {
	res := Geomspace(1, 16, 5, true)
	array := []float64{1, 2, 4, 8, 16}
	for i := range res.Arr {
		if math.Abs(res.Arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Geomspace(1, 16, 4, false)
	array = []float64{1, 2, 4, 8}
	for i := range res.Arr {
		if math.Abs(res.Arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Geomspace(0, 4, 2, true)
	if res.Err == nil {
		t.Errorf("There must be an ErrZeroValue")
	}

	res = Geomspace(-4, 4, 2, false)
	if res.Err == nil {
		t.Errorf("There must be an ErrNegativeValue")
	}

	res = Geomspace(4, -4, 2, false)
	if res.Err == nil {
		t.Errorf("There must be an ErrNegativeValue")
	}

	res = Geomspace(4, 4, 2, false)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	res = Geomspace(4, 256, 1, false)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestLogspace(t *testing.T) {
	res := Logspace(1, 4, 4, 10, true)
	array := []float64{10, 100, 1000, 10000}
	for i := range res.Arr {
		if math.Abs(res.Arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Logspace(0, 3, 4, 2, true)
	array = []float64{1, 2, 4, 8}
	for i := range res.Arr {
		if math.Abs(res.Arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Logspace(5, 4, 2, 3, true)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestEmpty(t *testing.T) {
	res := Empty(2)
	array := [][]float64{{0, 0}, {0, 0}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	res = Empty(-2)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestEmptyLike(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)

	array := EmptyLike(res).Arr

	if len(array) != 2 || len(array[0]) != 3 {
		t.Errorf("Shape of array must be (%d, %d)", 2, 3)
	}

	for i := range array {
		for j := range array[i] {
			if array[i][j] > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	err := EmptyLike(arr.Reshape(3, 3)).Err
	if err == nil {
		t.Error("There must be ErrSizeNotMatch")
	}
}

func TestIdentity(t *testing.T) {
	res := Identity(2)
	array := [][]float64{{1, 0}, {0, 1}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	res = Identity(-2)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestEye(t *testing.T) {
	res := Eye(2)
	fmt.Println(res.Arr)
	array := [][]float64{{1, 1}, {1, 1}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	res = Eye(-2)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestFull(t *testing.T) {
	res := Full(3, 2, 5)
	array := [][]float64{{5, 5}, {5, 5}, {5, 5}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	res = Full(-2, 2, 3)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestOnes(t *testing.T) {
	res := Ones(2, 3)
	array := [][]float64{{1, 1, 1}, {1, 1, 1}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*3+j, array[i][j])
			}
		}
	}

	res = Ones(-2, 2)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestZeros(t *testing.T) {
	res := Zeros(3, 2)
	array := [][]float64{{0, 0}, {0, 0}, {0, 0}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	res = Zeros(-2, 2)
	if res.Err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestFullLike(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}.Reshape(2, 3)
	res := FullLike(arr, 5)
	array := [][]float64{{5, 5, 5}, {5, 5, 5}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	arr = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}.Reshape(3, 3)
	res = FullLike(arr, 5)
	if res.Err == nil {
		t.Errorf("There must be an ErrSizeNotMatch")
	}
}

func TestOnesLike(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}.Reshape(2, 3)
	res := OnesLike(arr)
	array := [][]float64{{1, 1, 1}, {1, 1, 1}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	arr = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}.Reshape(3, 3)
	res = OnesLike(arr)
	if res.Err == nil {
		t.Errorf("There must be an ErrSizeNotMatch")
	}
}

func TestZerosLike(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}.Reshape(2, 3)
	res := ZerosLike(arr)
	array := [][]float64{{0, 0, 0}, {0, 0, 0}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i*2+j, array[i][j])
			}
		}
	}

	arr = Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4, 7}}.Reshape(3, 3)
	res = ZerosLike(arr)
	if res.Err == nil {
		t.Errorf("There must be an ErrSizeNotMatch")
	}
}

func TestDiag(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{3, 4, 7}}
	res := Diag(arr)
	array := [][]float64{{3, 0, 0}, {0, 4, 0}, {0, 0, 7}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and col %d != %f", i, j, array[i][j])
			}
		}
	}

	arr = Float64OneDArray{}
	res = Diag(arr)
	if res.Err == nil {
		t.Errorf("There must be an ErrEmptyArray")
	}
}

func TestVander(t *testing.T) {
	arr := Float64OneDArray{Arr: []float64{2, 3, 4}}
	res := Vander(arr, false)
	array := [][]float64{{4, 2, 1}, {9, 3, 1}, {16, 4, 1}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and col %d != %f", i, j, array[i][j])
			}
		}
	}

	res = Vander(arr, true)
	array = [][]float64{{1, 2, 4}, {1, 3, 9}, {1, 4, 16}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.Arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at row %d and col %d != %f", i, j, array[i][j])
			}
		}
	}
}
