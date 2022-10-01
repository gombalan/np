package np

import "testing"

func TestMaxInt32(t *testing.T) {
	max, _ := Arange(0, 4, 1).Max()
	if *max != 3 {
		t.Errorf("Maximum value in array must be %d", 3)
	}

	_, err := Arange(7, 4, 5).Max()
	if err == nil {
		t.Errorf("There must be an ErrEmptyArray")
	}
}

func TestMinInt32(t *testing.T) {
	min, _ := Arange(4, 0, -1).Min()
	if *min != 1 {
		t.Errorf("Minimum value in array must be %d", 1)
	}

	_, err := Arange(4, 4, 1).Min()
	if err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestPtpInt32(t *testing.T) {
	ptp, _ := Arange(0, 4, 1).Ptp()
	if *ptp != 3 {
		t.Errorf("Point to point value in array must be %d", 3)
	}

	_, err := Arange(4, 4, 1).Ptp()
	if err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestSumInt32(t *testing.T) {
	sum, _ := Arange(0, 4, 1).Sum()
	if *sum != 6 {
		t.Errorf("Sum of all value in array must be %d", 6)
	}

	_, err := Arange(4, 4, 1).Sum()
	if err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestMeanInt32(t *testing.T) {
	mean, _ := Arange(0, 4, 1).Mean()
	if *mean != 1.5 {
		t.Errorf("Mean of all value in array must be %f", float64(1.5))
	}

	_, err := Arange(4, 4, 1).Mean()
	if err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestMedianInt32(t *testing.T) {
	median, _ := Arange(0, 4, 1).Median()
	if *median != 1.5 {
		t.Errorf("Median of all value in array must be %f", float64(1.5))
	}

	median, _ = Arange(0, 5, 1).Median()
	if *median != 2 {
		t.Errorf("Median of all value in array must be %f", float64(2))
	}

	_, err := Arange(4, 4, 1).Median()
	if err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}

func TestModeInt32(t *testing.T) {
	mode, _ := Float64OneDArray{Arr: []float64{1, 1, 1, 3, 4}}.Mode()
	if *mode != 1 {
		t.Errorf("Mode of all value in array must be %d", 1)
	}

	_, err := Arange(4, 4, 1).Mode()
	if err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}
