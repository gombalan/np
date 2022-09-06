package np

func Arange(start int32, stop int32, step int32) Int32OneDArray {
	if step == 0 {
		return Int32OneDArray{nil, newError(ErrZeroValue, arg{key: "step", value: 0})}
	}

	if start == stop {
		return Int32OneDArray{nil, newError(ErrInvalidParameter, arg{key: "start", value: start}, arg{key: "stop", value: stop})}
	}

	size := (stop - start) / step
	if size < 1 {
		return Int32OneDArray{nil, newError(ErrInvalidParameter, arg{key: "size", value: size}, arg{key: "step", value: 0}, arg{key: "start", value: start}, arg{key: "stop", value: stop})}
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
