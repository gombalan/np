package np

import (
	"fmt"
	"math"
	"testing"
)

func TestLinspace(t *testing.T) {
	res := Linspace(0, 4, 5, false)
	array := []float64{0, 0.8, 1.6, 2.4, 3.2}
	for i := range res.arr {
		if math.Abs(res.arr[i]-array[i]) > 1e-9 {
			fmt.Println(res.arr[i])
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
			fmt.Println(res.arr[i])
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Linspace(1, 16, 4, false)
	array = []float64{1, 2, 4, 8}
	for i := range res.arr {
		if math.Abs(res.arr[i]-array[i]) > 1e-9 {
			t.Errorf("Element at %d != %f", i, array[i])
		}
	}

	res = Linspace(0, 4, 2, true)
	if res.err == nil {
		t.Errorf("There must be an ErrZeroValue")
	}

	res = Linspace(-4, 4, 2, false)
	if res.err == nil {
		t.Errorf("There must be an ErrNegativeValue")
	}

	res = Linspace(4, -4, 2, false)
	if res.err == nil {
		t.Errorf("There must be an ErrNegativeValue")
	}

	res = Linspace(4, 4, 2, false)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	res = Linspace(4, 256, 1, false)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}