package np

import "testing"

func TestMaxInt32(t *testing.T) {
	_, err := Arange(7, 4, 5).Max()
	if err == nil {
		t.Errorf("There must be an ErrEmptyArray")
	}

	max, _ := Arange(0, 4, 1).Max()
	if max != 3 {
		t.Errorf("Maximum value in array must be %d", 3)
	}
}

func TestMinInt32(t *testing.T) {
	_, err := Arange(4, 4, 1).Min()
	if err == nil {
		t.Errorf("There must be an ErrEmptyArray")
	}

	min, _ := Arange(4, 0, -1).Min()
	if min != 1 {
		t.Errorf("Minimum value in array must be %d", 1)
	}
}

func TestPtpInt32(t *testing.T) {
	_, err := Arange(4, 4, 1).Ptp()
	if err == nil {
		t.Errorf("There must be an ErrEmptyArray")
	}

	ptp, _ := Arange(0, 4, 1).Ptp()
	if ptp != 3 {
		t.Errorf("Point to point value in array must be %d", 3)
	}
}

func TestSumInt32(t *testing.T) {
	sum, _ := Arange(0, 4, 1).Sum()
	if sum != 6 {
		t.Errorf("Sum of all value in array must be %d", 6)
	}
}

func TestMeanInt32(t *testing.T) {
	mean, _ := Arange(0, 4, 1).Mean()
	if mean != 1.5 {
		t.Errorf("Mean of all value in array must be %f", float64(1.5))
	}
}

func TestMedianInt32(t *testing.T) {
	median, _ := Arange(0, 4, 1).Median()
	if median != 1.5 {
		t.Errorf("Median of all value in array must be %f", float64(1.5))
	}

	median, _ = Arange(0, 5, 1).Median()
	if median != 2 {
		t.Errorf("Median of all value in array must be %f", float64(2))
	}
}

func TestModeInt32(t *testing.T) {
	mode, _ := Int32OneDArray{arr: []int32{1, 1, 1, 3, 4}}.Mode()
	if mode != 1 {
		t.Errorf("Mode of all value in array must be %d", 1)
	}
}
