package np

import (
	"errors"
	"sort"
)

type Int32OneDArray struct {
	arr []int32
	err error
}

func ArangeInt32(start int32, stop int32, step int32) Int32OneDArray {
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

func (a Int32OneDArray) Len() (int32, error) {
	if a.err != nil {
		return 0, a.err
	}

	return int32(len(a.arr)), nil
}

func (a Int32OneDArray) Max() (int32, error) {
	if err := validateArray(a); err != nil {
		return 0, err
	}

	max := a.arr[0]
	for _, value := range a.arr {
		if value > max {
			max = value
		}
	}

	return max, nil
}

func (a Int32OneDArray) Min() (int32, error) {
	if err := validateArray(a); err != nil {
		return 0, err
	}

	min := a.arr[0]
	for _, value := range a.arr {
		if value < min {
			min = value
		}
	}

	return min, nil
}

func (a Int32OneDArray) Ptp() (int32, error) {
	if err := validateArray(a); err != nil {
		return 0, err
	}

	max, _ := a.Max()
	min, _ := a.Min()

	return max - min, nil
}

func (a Int32OneDArray) Sum() (int32, error) {
	if err := validateArray(a); err != nil {
		return 0, err
	}

	sum := int32(0)
	for _, value := range a.arr {
		sum += value
	}

	return sum, nil
}

func (a Int32OneDArray) Mean() (float64, error) {
	if err := validateArray(a); err != nil {
		return 0, err
	}

	sum, _ := a.Sum()

	return float64(sum) / float64(len(a.arr)), nil
}

func (a Int32OneDArray) Median() (float64, error) {
	if err := validateArray(a); err != nil {
		return 0, err
	}

	sort.SliceStable(a.arr, func(i, j int) bool { return a.arr[i] < a.arr[j] })

	if len(a.arr)%2 != 0 {
		return float64(a.arr[len(a.arr)/2]), nil
	}

	return float64(a.arr[len(a.arr)/2]+a.arr[len(a.arr)/2-1]) / 2, nil
}

func (a Int32OneDArray) Mode() (int32, error) {
	if err := validateArray(a); err != nil {
		return 0, err
	}

	numberMap := make(map[int32]int)
	count := 0
	mode := a.arr[0]
	for _, value := range a.arr {
		if _, exist := numberMap[value]; !exist {
			numberMap[value] = 0
		}
		numberMap[value]++

		if numberMap[value] > count {
			count = numberMap[value]
			mode = value
		}
	}

	return mode, nil
}
