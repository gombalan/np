package np

import "errors"

func Arange(start int32, stop int32, step int32) Int32OneDArray {
	if step == 0 {
		return Int32OneDArray{nil, errors.New(ErrZeroValue)}
	}

	if start == stop {
		return Int32OneDArray{nil, errors.New(ErrInvalidParameter)}
	}

	size := (stop - start) / step
	if size < 1 {
		return Int32OneDArray{nil, errors.New(ErrInvalidParameter)}
	}

	result := make([]int32, 0, size)
	if stop > start {
		for i := start; i < stop; i += step {
			result = append(result, i)
		}
	} else {
		for i := start; i > stop; i += step {
			result = append(result, i)
		}
	}

	return Int32OneDArray{result, nil}
}
