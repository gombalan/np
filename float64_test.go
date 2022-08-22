package np

import (
	"math"
	"testing"
)

func TestLinspace(t *testing.T) {
	res := Linspace(0, 4, 5, false)
	array := []float64{0, 0.8, 1.6, 2.4, 3.2}
	for i := range res.arr {
		if math.Abs(res.arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Linspace(1, 4, 4, true)
	array = []float64{1, 2, 3, 4}
	for i := range res.arr {
		if math.Abs(res.arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Linspace(4, 4, 2, true)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	res = Linspace(4, 4, 0, false)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestGeomspace(t *testing.T) {
	res := Geomspace(1, 16, 5, true)
	array := []float64{1, 2, 4, 8, 16}
	for i := range res.arr {
		if math.Abs(res.arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Geomspace(1, 16, 4, false)
	array = []float64{1, 2, 4, 8}
	for i := range res.arr {
		if math.Abs(res.arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Geomspace(0, 4, 2, true)
	if res.err == nil {
		t.Errorf("There must be an ErrZeroValue")
	}

	res = Geomspace(-4, 4, 2, false)
	if res.err == nil {
		t.Errorf("There must be an ErrNegativeValue")
	}

	res = Geomspace(4, -4, 2, false)
	if res.err == nil {
		t.Errorf("There must be an ErrNegativeValue")
	}

	res = Geomspace(4, 4, 2, false)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	res = Geomspace(4, 256, 1, false)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestLogspace(t *testing.T) {
	res := Logspace(1, 4, 4, 10, true)
	array := []float64{10, 100, 1000, 10000}
	for i := range res.arr {
		if math.Abs(res.arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Logspace(0, 3, 4, 2, true)
	array = []float64{1, 2, 4, 8}
	for i := range res.arr {
		if math.Abs(res.arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Logspace(5, 4, 2, 3, true)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestIdentity(t *testing.T) {
	res := Identity(2)
	array := [][]float64{{1, 0}, {0, 1}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i, array[i])
			}
		}
	}

	res = Identity(-2)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestOnes(t *testing.T) {
	res := Ones(2, 3)
	array := [][]float64{{1, 1, 1}, {1, 1, 1}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i, array[i])
			}
		}
	}

	res = Ones(-2, 2)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestZeros(t *testing.T) {
	res := Zeros(3, 2)
	array := [][]float64{{0, 0}, {0, 0}, {0, 0}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i, array[i])
			}
		}
	}

	res = Zeros(-2, 2)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestFull(t *testing.T) {
	res := Full(3, 2, 5)
	array := [][]float64{{5, 5}, {5, 5}, {5, 5}}
	for i := range array {
		for j := range array[i] {
			if math.Abs(res.arr[i][j]-array[i][j]) > 1e-9 {
				t.Errorf("Element at %d != %f", i, array[i])
			}
		}
	}

	res = Full(-2, 2, 3)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}
