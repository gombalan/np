package np

import (
	"fmt"
	"testing"
)

func TestArangeInt32(t *testing.T) {
	res := ArangeInt32(0, 4, 1)
	for i := range res.arr {
		if res.arr[int32(i)] != int32(i) {
			t.Errorf("Element at %d != %d", i, i)
		}
	}

	res = ArangeInt32(1, 4, 1)
	for i := range res.arr {
		if res.arr[int32(i)] != 1+int32(i) {
			t.Errorf("Element at %d != %d", i, i+1)
		}
	}

	res = ArangeInt32(1, 4, 2)
	for i := range res.arr {
		if res.arr[int32(i)] != 1+2*int32(i) {
			t.Errorf("Element at %d != %d", i, 2*i+1)
		}
	}

	res = ArangeInt32(1, 1, 0)
	if res.err == nil {
		t.Errorf("There must be an ErrZeroValue")
	}

	res = ArangeInt32(1, 1, 2)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	res = ArangeInt32(1, -3, 2)
	if res.err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestLenInt32(t *testing.T) {
	res, _ := ArangeInt32(0, 4, 1).Len()
	if res != int32(4) {
		t.Errorf("Length of array must be %d", 4)
	}
}

func TestMaxInt32(t *testing.T) {
	_, err := ArangeInt32(7, 4, 5).Max()
	if err == nil {
		t.Errorf("There must be an ErrEmptyArray")
	}

	max, _ := ArangeInt32(0, 4, 1).Max()
	if max != int32(3) {
		t.Errorf("Maximum value in array must be %d", 3)
	}
}

func TestMinInt32(t *testing.T) {
	_, err := ArangeInt32(4, 4, 1).Min()
	if err == nil {
		t.Errorf("There must be an ErrEmptyArray")
	}

	min, _ := ArangeInt32(4, 0, -1).Min()
	if min != int32(1) {
		fmt.Println(min)
		t.Errorf("Minimum value in array must be %d", 1)
	}
}

func TestSumInt32(t *testing.T) {
	sum, _ := ArangeInt32(0, 4, 1).Sum()
	if sum != int32(6) {
		t.Errorf("Sum of all value in array must be %d", 6)
	}
}

func TestMeanInt32(t *testing.T) {
	mean, _ := ArangeInt32(0, 4, 1).Mean()
	if mean != float64(1.5) {
		t.Errorf("Mean of all value in array must be %f", float64(1.5))
	}
}

func TestMedianInt32(t *testing.T) {
	median, _ := ArangeInt32(0, 4, 1).Median()
	if median != float64(1.5) {
		t.Errorf("Median of all value in array must be %f", float64(1.5))
	}

	median, _ = ArangeInt32(0, 5, 1).Median()
	if median != float64(2) {
		t.Errorf("Median of all value in array must be %f", float64(2))
	}
}
