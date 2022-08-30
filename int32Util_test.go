package np

import "testing"

func TestLenInt32(t *testing.T) {
	res, _ := Arange(0, 4, 1).Len()
	if *res != 4 {
		t.Errorf("Length of array must be %d", 4)
	}

	_, err := Arange(4, 4, 1).Len()
	if err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}
}
