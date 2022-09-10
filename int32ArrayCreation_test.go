package np

import (
	"testing"
)

func TestArange(t *testing.T) {
	res := Arange(0, 4, 1)
	for i := range res.Arr {
		if res.Arr[i] != int32(i) {
			t.Errorf("Element at %d != %d", i, i)
		}
	}

	res = Arange(1, 4, 1)
	for i := range res.Arr {
		if res.Arr[i] != 1+int32(i) {
			t.Errorf("Element at %d != %d", i, i+1)
		}
	}

	res = Arange(1, 4, 2)
	for i := range res.Arr {
		if res.Arr[i] != 1+2*int32(i) {
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
