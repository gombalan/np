package np

import "testing"

func TestReshapeInt32(t *testing.T) {
	arr := Int32OneDArray{Arr: []int32{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if res.Arr[i][j] != arr.Arr[i*3+j] {
				t.Errorf("Element at row %d colum %d must be %d", i, j, arr.Arr[i*3+j])
			}
		}
	}

	arr = Int32OneDArray{}
	if err := arr.Reshape(2, 3).Err; err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	if err := Arange(0, 4, 1).Reshape(-1, 3).Err; err == nil {
		t.Errorf("There must be an ErrInvalidParameter")
	}

	if err := Arange(0, 4, 1).Reshape(3, 3).Err; err == nil {
		t.Errorf("There must be an ErrSizeNotMatch")
	}
}
