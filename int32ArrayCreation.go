package np

import "fmt"

func Arange(start int32, stop int32, step int32) Int32OneDArray {
	if step == 0 {
		return Int32OneDArray{nil, newError(ErrZeroValue, fmt.Sprintf("but, value of step is %v", step))}
	}

	if start == stop {
		return Int32OneDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of start: %v and stop: %v must not be equal, but they are", start, stop))}
	}

	size := (stop - start) / step
	if size < 1 {
		return Int32OneDArray{nil, newError(ErrInvalidParameter, fmt.Sprintf("value of start: %v, stop: %v and step: %v result in non-positive array size", start, stop, step))}
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
