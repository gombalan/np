package np

import "testing"

func TestLenInt32(t *testing.T) {
	res, _ := Arange(0, 4, 1).Len()
	if res != int32(4) {
		t.Errorf("Length of array must be %d", 4)
	}
}

func TestReshapeInt32(t *testing.T) {
	arr := Int32OneDArray{arr: []int32{1, 1, 1, 3, 4, 7}}
	res := arr.Reshape(2, 3)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if res.arr[i][j] != arr.arr[i*3+j] {
				t.Errorf("Element at row %d colum %d must be %d", i, j, arr.arr[i*3+j])
			}
		}
	}
}
