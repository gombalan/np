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
